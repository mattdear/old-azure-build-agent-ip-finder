# Azure Client IP Finder

## Setup
1. Open a terminal in the azure-build-agent-ip-finder folder.

2. Run this command to build the docker image. 

```
docker build -t azure-build-agent-ip-finder .
```

3. Run this command to start the container on a host.

```
docker run -dp 8080:8080 --name azure-build-agent-ip-finder azure-build-agent-ip-finder
```

## Run

Replace the \<IP> with the IP address of the container or, the IP of the host running the container. You will then recieve the IP address of your Azure Build Agent as a JSON response.

```
curl <IP>:8080/getazurebuildagentip
```