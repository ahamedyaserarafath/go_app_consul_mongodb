## Simple docker file which will copy the current directory
## and download the respective golang and its package necessary to run the server
## finally it will run the simple api server in golang
FROM golang:1.8


WORKDIR /app/
COPY start_consul_golangapp.sh /app/

## Below command will install the respective consul package
RUN set -x && apt-get update && apt-get install -y wget zip iputils-ping 

## Downloadn the Consul and move it to path
RUN set -x && wget https://releases.hashicorp.com/consul/1.4.0/consul_1.4.0_linux_amd64.zip && unzip consul_1.4.0_linux_amd64.zip && mv consul /usr/local/bin/ && rm consul_1.4.0_linux_amd64.zip


RUN mkdir /go/src/app
COPY app /go/src/app/

WORKDIR /go/src/app
RUN go-wrapper download   
RUN go-wrapper install    
RUN cd /go/src/app/ && go build

WORKDIR /app/

EXPOSE 8080
ENTRYPOINT ["/app/start_consul_golangapp.sh"]

