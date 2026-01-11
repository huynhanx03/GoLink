FROM debezium/connect:2.7.3.Final

USER root
RUN microdnf install gzip tar && microdnf clean all

USER kafka

# Install ScyllaDB Connector
RUN curl -so /kafka/connect/scylla-cdc-source-connector-1.2.7-jar-with-dependencies.jar https://repo1.maven.org/maven2/com/scylladb/scylla-cdc-source-connector/1.2.7/scylla-cdc-source-connector-1.2.7-jar-with-dependencies.jar
