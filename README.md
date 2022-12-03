# Extractor-timer

## Project description

Extractor-timer extracts (scrapes) data from web stores and puts it into a Mongo database. It does that periodicaly. 

Author: [Miha Krumpestar](https://github.com/mk2376)

## Setup/installation

```
go mod download && go mod verify
```

Other usages are defined in [Makefile](Makefile)

## Endpoints

### extractor-timer (public)

Base URL: `https://multimo.ml/extractor`

Paths:

- `/self` Health check.
- `/v1/info` Server internal state.
- `/v1/extract` Manualy starts extraction.

### products-db (private)

MongoDB instance. `M_USERNAME` and `M_PASSWORD` are provided in `.env`.

URL: `mongodb://${M_USERNAME}:${M_PASSWORD}@products-db:27017/`

### products-db-ui (public, behind middleware auth)

Mongo-express instance, connected to `products-db`. Credentials  are provided in `.env` in [shared repository](https://github.com/MultimoML/shared).

URL: `https://products-db-ui.multimo.ml`

## License

RSOcena is licensed under the [GNU AGPLv3 license](LICENSE).
