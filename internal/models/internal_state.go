package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CurrentState string

const (
	CurrentStateIdle    CurrentState = "idle"
	CurrentStateRunning CurrentState = "running"
)

type InternalState struct {
	Id               primitive.ObjectID `json:"id" bson:"_id" validate:"required"`
	RunInterval      time.Duration      `json:"run-interval" bson:"run-interval" validate:"required"`             // in nanoseconds
	LastRunTimestamp primitive.DateTime `json:"last-run-timestamp" bson:"last-run-timestamp" validate:"required"` // in milliseconds
	NextRunTimestamp primitive.DateTime `json:"next-run-timestamp" bson:"next-run-timestamp" validate:"required"` // in milliseconds
	CurrentState     CurrentState       `json:"current-state" bson:"current-state" validate:"required"`
	RunCount         uint64             `json:"run-count" bson:"run-count" validate:"required"`
}
