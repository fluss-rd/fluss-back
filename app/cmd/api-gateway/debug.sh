#!/bin/sh
cd $APP_HOME/app/api-gateway/router
dlv debug --headless --log -l 0.0.0.0:2345 --api-version=2