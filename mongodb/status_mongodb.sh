#!/bin/bash
ISMASTER=$(mongo --port 5000 --quiet --eval 'db.isMaster().ismaster')
if [ "$ISMASTER" = "true" ]
then
    echo "Mongodb is up and running"
    exit 0
else
    echo "Mongodb is down"
    exit 2
fi
