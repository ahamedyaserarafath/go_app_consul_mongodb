#!/bin/bash
## Simple bash shell script will start the mongodb and insert the data 
## For now hardcoded the path of mongodb bin and data folder
set -e

echo "Starting the mongodb, please wait"
mongod --fork --port 5000 --dbpath /app/mongodb_data --bind_ip 0.0.0.0 --replSet rpset --logpath db.log
mongo --port 5000 --eval "rs.initiate()" > /dev/null
echo "Waiting for mongodb to come up as master.."
sleep 30
echo "Inserting the greeting string as 'Hello, world'"
mongo --port 5000 --eval 'db.greeting.insert({"title": "Hello, world"})'

echo "Starting the consul agent in mongodb,please wait.."
consul agent -dev  -client=0.0.0.0 -bind=0.0.0.0 -enable-script-checks=true -config-file=mongodb_consul.json > consul.log &

#Extra line added in the script to run all command line arguments
ping localhost > /dev/null
exec "$@";

