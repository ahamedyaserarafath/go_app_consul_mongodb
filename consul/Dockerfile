## Simple docker file which will copy the current directory
## and run the startup script
## finally it will run the consul in background

FROM debian:jessie-slim

WORKDIR /app
COPY start_consul.sh /app/

## Below command will install the respective consul package
RUN set -x && apt-get update && apt-get install -y wget zip iputils-ping 

## Downloadn the Consul and move it to path
RUN set -x && wget https://releases.hashicorp.com/consul/1.4.0/consul_1.4.0_linux_amd64.zip && unzip consul_1.4.0_linux_amd64.zip && mv consul /usr/local/bin/ && rm consul_1.4.0_linux_amd64.zip

EXPOSE 8500
ENTRYPOINT ["/app/start_consul.sh"]

