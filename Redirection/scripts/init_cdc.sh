#!/bin/bash

# Configuration
CONNECT_HOST="localhost"
CONNECT_PORT="2010"
CONNECTOR_NAME="scylla-link-connector"

echo "Registering Debezium ScyllaDB Connector at http://$CONNECT_HOST:$CONNECT_PORT..."

# Register Connector
response=$(curl -s -o /dev/null -w "%{http_code}" -X POST -H "Accept:application/json" -H "Content-Type:application/json" http://$CONNECT_HOST:$CONNECT_PORT/connectors/ -d '{
  "name": "'"$CONNECTOR_NAME"'",
  "config": {
    "connector.class": "com.scylladb.cdc.debezium.connector.ScyllaConnector",
    "scylla.cluster.ip.addresses": "scylla-node1:9042",
    "scylla.name": "golink.links",
    "topic.prefix": "golink.links",
    "scylla.table.names": "generation_ks.links",
    "key.converter": "org.apache.kafka.connect.json.JsonConverter",
    "key.converter.schemas.enable": "false",
    "value.converter": "org.apache.kafka.connect.json.JsonConverter",
    "value.converter.schemas.enable": "false"
  }
}')

if [ "$response" -eq 201 ]; then
  echo "Connector registered successfully!"
elif [ "$response" -eq 409 ]; then
  echo "Connector already exists."
else
  echo "Failed to register connector. HTTP Status: $response"
  echo "Configuring failed? Check if Debezium is running: docker ps | grep debezium"
  exit 1
fi
