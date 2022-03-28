# Fluss backend services

## Prerequisites

You need to have installed:

- [Docker](https://docs.docker.com/engine/install/ubuntu/)
- [docker-compose](https://docs.docker.com/compose/install/)

## Run the project

**Configure the environment variables**:

- Make a copy of the content of `environment.example` into a new directory called `environment`
- Configure the variables as needed

**Run the services**:

```
docker-compose up
```

**Configure Influx DB**:

- Go to `localhost:8086`
- Fill the form as needed, and use "fluss" as organization name and create a bucket with the name "fluss" also
- Go to `reporting.local.env` and `telemetry.local.env` and replace `INFLUXDB_TOKEN` with a token you can find going to `Load data -> Tokens` and click on `fluss's Token`

