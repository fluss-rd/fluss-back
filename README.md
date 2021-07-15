# fluss-back
Fluss backend services


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