# Docker - Golang - Mongo - consul - docker compose
# Deploying the golang and mongo as a app docker. 
- [Introduction](#Introduction)
- [Pre-requisites](#pre-requisites)
- [Installation and configuration](#Installation-and-configuration)

# Introduction
In this post, we will deploy the golang, mongodb, consul as a docker compose as consul will discover the golang app services

* Docker-compose file, has the following docker containers:
    - Mongodb database engine
    - Consul service discovery tool
    - Golang application
* Golang Application will discover the database configuration on startup using the Consul service discovery tool with API
* Golang Application will connect to a database, read a greeting string and print it on the url.
* Golang Application will be served as HTTP and be available on host's 127.0.0.1:8080


# Pre-requisites
Before we get started installing the golang and mongo. 
* Ensure you install the latest version of docker and docker-compose.

# Installation and configuration

* Clean up the existing the docker-compose images if needed(Be caution it will remove every images in the server)
```
docker-compose stop
docker-compose rm
```
* Build the docker-compose with the below command
```  
docker-compose build 
```
* Run the docker-compose with the below command
```
docker-compose up
```
Port Forwarded information:
* 5000 port is exposed for mongodb
* 8500 port is exposed for accessing consul HTTP api
* 8080 port is exposed for mygoapp
