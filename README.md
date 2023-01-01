# Extractor

Microservice to extract data from online shops.

Available endpoints:
- [`/self`](https://multimo.ml/extractor/self): Liveliness and readiness check
- [`/v1/info`](https://multimo.ml/extractor/v1/info): Returns internal state of extractor
- [`/v1/extract`](https://multimo.ml/extractor/v1/extract): Signals extractor to initiate a new extraction

## Setup/installation

Prerequisites:
- [Go](https://go.dev/)
- [Docker](https://www.docker.com/)

Example usage:
- See all available options: `make help`
- Run microservice in a container: `make run`
- Release a new version: `make release ver=x.y.z`

All work should be done on `main`, `prod` should never be checked out or manually edited.
When releasing, the changes are merged into `prod` and both branches are pushed.
A GitHub Action workflow will then build and publish the image to GHCR, and deploy it to Kubernetes.

## License

Project is licensed under the [GNU AGPLv3 license](LICENSE).
