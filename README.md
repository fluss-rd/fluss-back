# Fluss backend services

## Prerequisites

You need to have installed:

- [Docker](https://docs.docker.com/engine/install/ubuntu/)
- [docker-compose](https://docs.docker.com/compose/install/)

## Run the project

**Prerequisites**:

- Configure the variables as needed in `environments/`. Remember that the ports you use in the `.env` files are the corresponding ports in the container.
- Configure Influx DB:
  - Run `docker-compose up influxdb`.
  - Go to `localhost:8086`.
  - Fill the form as needed, and use "fluss" as organization name and create a bucket with the name "fluss".
  - Go to `reporting.local.env` and `telemetry.local.env` and replace `INFLUXDB_TOKEN` with the token get from `Load data -> Tokens` and clicking on `fluss's Token`.

To run the project just do:

```bash
docker-compose up
```
