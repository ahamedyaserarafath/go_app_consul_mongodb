## Simple docker file which will copy the current directory
## and run the startup script which will insert the data 
## finally it will run the mongodb in background

FROM debian:jessie-slim

WORKDIR /app
COPY start_mongodb.sh /app/
COPY mongodb_consul.json /app/
COPY status_mongodb.sh /app/

RUN mkdir -p /app/mongodb_data/

## Below command will install the respective package for mongodb
RUN set -x && apt-get update && apt-get install -y libssl1.0.0 libssl-dev apt-transport-https ca-certificates iputils-ping

## Install the mongodb for debian
RUN set -x && apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 9DA31620334BD75D9DCB49F368818C72E52529D4 && echo "deb http://repo.mongodb.org/apt/debian jessie/mongodb-org/4.0 main" | tee /etc/apt/sources.list.d/mongodb-org-4.0.list && apt-get update && apt-get install -y mongodb-org

## Below command will install the respective consul package
RUN set -x && apt-get install -y wget zip 

## Downloadn the Consul and move it to path
RUN set -x && wget https://releases.hashicorp.com/consul/1.4.0/consul_1.4.0_linux_amd64.zip && unzip consul_1.4.0_linux_amd64.zip && mv consul /usr/local/bin/ && rm consul_1.4.0_linux_amd64.zip


EXPOSE 5000
ENTRYPOINT ["/app/start_mongodb.sh"]
