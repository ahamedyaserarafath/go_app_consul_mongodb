## Readme file for building the docker and run the docker
## This docker is for installing the consul and running in the particular port

Steps to Run the simple consul server

1. Clean up the existing the docker images (Be caution it will remove every images in the server)

  'docker rm $(docker ps -a -q)'
  'docker rmi $(docker images -q)'

2. Build the docker container with the below command
  
   'docker build -t consul .'

3. Run the docker container with the below command

   'docker run -it -p 8500:8500 -d consul'
