package internal_state

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/multimoml/extractor/internal/configs"
	"github.com/multimoml/extractor/internal/db_client"
	"github.com/multimoml/extractor/internal/models"
	"github.com/multimoml/extractor/internal/scraper"

	"github.com/jasonlvhit/gocron"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	once sync.Once

	ctx           context.Context
	internalState models.InternalState
	dbCollection  *mongo.Collection
	scheduler     *gocron.Scheduler

	filter = bson.M{
		"_id": id(),
	}
)

func InternalState(ctxIn ...context.Context) models.InternalState {

	once.Do(func() { // <-- atomic, does not allow repeating
		ctx = ctxIn[0]
		initInternalStateDB() // <-- thread safe
		setScheduler()
		fmt.Printf("InternalState set\n")

		if internalState.NextRunTimestamp.Time().Before(time.Now()) || internalState.RunCount == 0 {
			fmt.Printf("\nStarting first Scrape run...\n")
			internalState.LastRunTimestamp = primitive.DateTime(time.Now().UnixMilli())
			internalState.NextRunTimestamp = primitive.DateTime(internalState.LastRunTimestamp.Time().Add(internalState.RunInterval).UnixMilli())

			scheduler.RunAll()
		}
	})

	return internalState
}

// DO NOT USE YET! UNFINISHED! TODO:
func InternalStateHotReload() {
	fmt.Printf("\nReloading InternalState...\n")

	scheduler.Clear()
	setScheduler()
	fmt.Printf("InternalState reloaded\n")
}

func Scrape() {
	fmt.Printf("\nStarted Scrape...\n")
	start := time.Now()

	internalState.CurrentState = models.CurrentStateRunning
	applyToDB()

	products := scraper.ScrapeSpar(ctx)
	scraper.WriteProductsSpar(ctx, products)

	// prevent changing internalState time if function was triggered before internalState.NextRunTimestamp
	if internalState.NextRunTimestamp.Time().Before(time.Now()) {
		internalState.LastRunTimestamp = internalState.NextRunTimestamp
		internalState.NextRunTimestamp = primitive.DateTime(internalState.LastRunTimestamp.Time().Add(internalState.RunInterval).UnixMilli())
	}

	internalState.CurrentState = models.CurrentStateIdle
	internalState.RunCount++
	applyToDB()

	// calculate to exe time
	elapsed := time.Since(start)
	fmt.Printf("\nRun at %v finished. Total run time %v\n\n", start.Format("2006-01-02T15:04:05Z07:00"), elapsed)
}

func setScheduler() {
	fmt.Printf("Started setScheduler...\n")

	interval := uint64(internalState.RunInterval / time.Second)
	t := internalState.LastRunTimestamp.Time()

	scheduler = gocron.NewScheduler()
	scheduler.Every(interval).Seconds().From(&t).Do(Scrape)

	scheduler.Start()
	fmt.Printf("setScheduler set\n")
}

func initInternalStateDB() {
	fmt.Printf("Started initInternalStateDB...\n")

	dbClient := db_client.DBClient()
	dbCollection = db_client.GetCollectionInternalState(dbClient)

	runIntervalEnv, err := configs.GetEnv("RUN_INTERVAL") // in seconds
	if err != nil {
		panic(err)
	}

	runIntervalSeconds, err := strconv.Atoi(*runIntervalEnv)
	if err != nil {
		log.Fatal(err)
	}

	runInterval := time.Second * time.Duration(runIntervalSeconds)

	now := primitive.DateTime(time.Now().UnixMilli())
	nextRunTimestamp := primitive.DateTime(now.Time().Add(runInterval).UnixMilli())

	internalState = models.InternalState{
		Id:               id(),
		RunInterval:      runInterval,
		LastRunTimestamp: now,
		NextRunTimestamp: nextRunTimestamp,
		CurrentState:     models.CurrentStateIdle,
	}

	result, err := dbCollection.InsertOne(context.Background(), internalState)
	if err != nil {
		if !mongo.IsDuplicateKeyError(err) {
			fmt.Println(result)
			log.Fatal(err)
		}

		internalStateDB := getFromDB()

		internalState.LastRunTimestamp = internalStateDB.LastRunTimestamp
		internalState.NextRunTimestamp = primitive.DateTime(internalState.LastRunTimestamp.Time().Add(internalState.RunInterval).UnixMilli())
		internalState.CurrentState = models.CurrentStateIdle
		internalState.RunCount = internalStateDB.RunCount

		applyToDB()
	}

	// fmt.Printf("%+v", internalState)

	fmt.Printf("initInternalStateDB set\n")
}

func getFromDB() models.InternalState {
	var internalState models.InternalState

	err := dbCollection.FindOne(context.Background(), filter).Decode(&internalState)
	if err != nil {
		log.Fatal(err)
	}

	return internalState
}

func applyToDB() {
	dbCollection.UpdateOne(context.Background(), filter, bson.M{"$set": internalState})
}

func id() primitive.ObjectID {
	id := primitive.ObjectID{}
	err := id.UnmarshalText([]byte(fmt.Sprintf("%024d", 1)))
	if err != nil {
		log.Fatal(err)
	}

	return id
}
