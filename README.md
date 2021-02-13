# Icon GraphQL API

Docker hub [image](https://hub.docker.com/r/pranavt61/kafka-websocket-server)

## Local build
```
go build -o main .
```

## Docker build
```
docker build . -t icon-graphql-api:latest --target prod
docker run \
  -p "8000:8000" \
  icon-graphql-api:latest
```

## Enviroment Variables

| Name | Description | Default | Required |
|------|-------------|---------|----------|
| ICON_GRAPHQL_API_PORT | Exposed port for the graphql server | 8000 | False |
