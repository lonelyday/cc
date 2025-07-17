# Currency Service
## Getting Started

### Prerequisites

- Docker


### Running with Docker

#### With docker compose (A)

```bash
export OER_APP_ID=<my-app_id>
```

```bash
docker compose up -d
```


#### With docker build and run (B) 

Build the Docker image:

```bash
docker build . -f Dockerfile -t cc-rest-api:0.1
```

Run the container:

```bash
docker run -p 9090:9090 \
    -e OER_APP_ID=<App_id> \
    -d --rm cc-rest-api:0.1
```

#### Environment Variables

- `OER_APP_ID`: API key for external currency data provider (**required!**)

## API Endpoints

- `GET /rates?currencies=CUR1,CUR2` e.g. `curl -X GET 'http://localhost:9090/rates?currencies=PLN,JPY'`

## License

MIT