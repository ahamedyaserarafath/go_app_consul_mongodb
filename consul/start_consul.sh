#!/bin/bash
## Simple bash shell script will start the mongodb and insert the data 
## For now hardcoded the path of mongodb bin and data folder
set -e

echo "Starting the consul agent in consul container, by joining the another mongodb consul agent.. Please wait"
sleep 40
consul agent -dev -client=0.0.0.0 -bind=0.0.0.0 -retry-join=10.5.0.5  > consul.log &

#Extra line added in the script to run all command line arguments
ping localhost > /dev/null
exec "$@";

