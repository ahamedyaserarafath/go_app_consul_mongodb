# Readme file for building the docker-compose and run the docker-compose

1. docker-compose file, has the following docker containers:
    - Mongodb database engine
    - Consul service discovery tool
    - Golang application
2. The Golang Application will discover the database configuration on startup using the Consul service discovery tool with API
3. The Golang Application will connect to a database, read a greeting string and print it on the url.
4. The Golang Application will be served as HTTP and be available on host's 127.0.0.1:8080


Prerequsite:

1. Docker and docker-compose should be installed on the machines

Steps to Run the docker-compose server

1. Clean up the existing the docker-compose images if needed(Be caution it will remove every images in the server)

  'docker-compose stop'
  'docker-compose rm'

2. Build the docker-compose with the below command
  
   'docker-compose build '

3. Run the docker-compose with the below command

   'docker-compose up'

Port Forwarded:
1. 5000 port is exposed for mongodb
2. 8500 port is exposed for accessing consul HTTP api
3. 8080 port is exposed for mygoapp
