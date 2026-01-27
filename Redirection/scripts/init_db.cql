CREATE KEYSPACE IF NOT EXISTS redirection_ks
WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

USE redirection_ks;

CREATE TABLE IF NOT EXISTS links (
    id text PRIMARY KEY,
    original_url text,
    user_id int,
    tenant_id int,
    created_at timestamp,
    updated_at timestamp
);
