CREATE KEYSPACE IF NOT EXISTS generation_ks
WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 3};

USE generation_ks;

CREATE TABLE IF NOT EXISTS links (
    id text PRIMARY KEY,
    original_url text,
    created_at timestamp,
    updated_at timestamp
);
