#!/bin/bash
## Simple bash shell script will start the mongodb and insert the data 
## For now hardcoded the path of mongodb bin and data folder
set -e

echo "Consul agent starting in mygoapp container, please wait till it sync with another consul agents"
consul agent -dev -client=0.0.0.0 -bind=0.0.0.0 -retry-join=10.5.0.5  > consul.log &
sleep 40
/go/src/app/app
echo "Starting the mygoapp, please access http://localhost:8081 or http://<ip_address>:8081"
#Extra line added in the script to run all command line arguments
exec "$@";

