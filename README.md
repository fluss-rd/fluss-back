# fluss-back
Fluss backend services

## How was the WQI calculated?
To know more about the formula and the underlying theory, see the following article:
[A comparison between weighted arithmetic and Canadian
methods for a drinking water quality index at selected
locations in shatt al-kufa](https://iopscience.iop.org/article/10.1088/1757-899X/433/1/012026/pdf)

## How to run the project?

1. Install docker: https://docs.docker.com/engine/install/ubuntu/ or https://linuxconfig.org/how-to-install-docker-on-ubuntu-20-04-lts-focal-fossa
2. Install docker-compose: https://docs.docker.com/compose/install/
3. Run docker-compose up

### Common errors solutions
- If docker-compose up says

```
docker.errors.DockerException: Error while fetching server API version: ('Connection aborted.', PermissionError(13, 'Permission denied'))
[100155] Failed to execute script docker-compose
```
Make sure docker is running or restart it in case it didn't start properly with `sudo service docker start` or `sudo service docker restart`.

To rebuild an image you must use `docker-compose build` or `docker-compose up --build`.

You could assume that is the main file of all app/cmd/api-gateway/main.go

# Common Errors

## problems with docker-component up rabbit

- make sure this file exists: /etc/rabbitmq/definitions.json
