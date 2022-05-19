# clientip

## Setup

```shell
# Docker
docker build -t go-hostip .
docker run -dp 8080:8080 --name hostip go-hostip

# Verify container
curl $IP:8080/hostip
```
