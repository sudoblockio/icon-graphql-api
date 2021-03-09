# Icon GraphQL API

Docker hub [image](https://hub.docker.com/r/pranavt61/icon-graphql-api)

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

## Usage

Once the container is up, visiting the `{prefix}/` endpoint on your browser will take you to the Graphql Playground.

Example Graphql query:
```
query Block {
  block(hash: "202d264fb85603ab19f747a60c0cf1aac53b6a15d9567ce7c8bd5e015c023296") {
    hash,
    number,
    transaction_count,
    peer_id
  }
}
```

### Schemas

Blocks:

| Field | Type | Parameter |
|------|-------------|---------|
| hash | String | True |
| number | Int | False |
| signature | String | False |
| item_id | String | False |
| transaction_count | Int | False |
| type | String | False |
| version | String | False |
| peer_id | String | False |
| merkle_root_hash | String | False |
| item_timestamp | String | False |
| parent_hash | String | False |
| timestamp | Int | False |

Transactions:

| Field | Type | Filterable |
|------|-------------|---------|
| hash | String | True |
| signature | String | False |
| fee | Int | False |
| block_number | Int | False |
| transaction_index | Int | False |
| type | String | False |
| receipt_step_price | Int | False |
| from_address | String | False |
| value | Int | False |
| timestamp | String | False |
| receipt_status | Int | False |
| item_id | String | False |
| receipt_logs | String | False |
| block_hash | String | False |
| to_address | String | False |
| version | String | False |
| nonce | Int | False |
| receipt_cumulative_step_used | Int | False |
| receipt_score_address | String | False |
| data_type | String | False |
| item_timestamp | String | False |

Logs:

| Field | Type | Filterable |
|------|-------------|---------|
| transaction_hash| String | True |
| address| String | False |
| data| [String] | False |
| indexed| [String] | False |
| item_id| String | False |
| block_timestamp| Int | False |
| block_number| Int | False |
| block_hash| String | False |
| transaction_index| Int | False |
| type| String | False |
| item_timestamp| String | False |


## Enviroment Variables

| Name | Description | Default | Required |
|------|-------------|---------|----------|
| ICON_GRAPHQL_API_MONGODB_URL | URL of hosted mongodb service | NULL | True |
| ICON_GRAPHQL_API_MONGODB_USER | Username of hosted mongodb service | NULL | True |
| ICON_GRAPHQL_API_MONGODB_PASS | Password of hosetd mongodb service | NULL | True |
| ICON_GRAPHQL_API_PORT | Exposed port for the graphql server | 8000 | False |
| ICON_GRAPHQL_API_PREFIX | Prefix for all endpoints exposed | "" | False |
