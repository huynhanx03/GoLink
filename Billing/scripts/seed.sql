CREATE DATABASE billing;
\c billing;

-- create tables

CREATE TABLE IF NOT EXISTS plans (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(200) NULL,
    base_price DOUBLE PRECISION NOT NULL,
    period VARCHAR(20) NOT NULL,
    features JSONB NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP,
    deleted_by INTEGER
);

CREATE TABLE IF NOT EXISTS subscriptions (
    id SERIAL PRIMARY KEY,
    tenant_id INTEGER NOT NULL,
    plan_id INTEGER NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'PENDING',
    current_period_start TIMESTAMP WITH TIME ZONE NULL,
    current_period_end TIMESTAMP WITH TIME ZONE NULL,
    cancel_at_period_end BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP,
    deleted_by INTEGER
);

CREATE TABLE IF NOT EXISTS invoices (
    id SERIAL PRIMARY KEY,
    subscription_id INTEGER NOT NULL,
    tenant_id INTEGER NOT NULL,
    amount DOUBLE PRECISION NOT NULL,
    currency VARCHAR(3) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'PENDING',
    payment_id VARCHAR(255) NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP,
    deleted_by INTEGER
);

-- insert data

INSERT INTO plans (id, name, description, base_price, period, features, is_active, created_at, updated_at, deleted_at, deleted_by) VALUES
(1, 'Free', 'Perfect for getting started', 0.00, 'forever', '{"max_links": 10, "ttl": 604800}', true, NOW(), NOW(), NULL, NULL),
(2, 'Pro', 'For power users and creators', 9.99, 'month', '{"max_links": 100, "ttl": 2592000}', true, NOW(), NOW(), NULL, NULL),
(3, 'Enterprise', 'For large teams and businesses', 99.99, 'month', '{"max_links": 1000, "ttl": 31536000, "customer_domain": true}', true, NOW(), NOW(), NULL, NULL)
ON CONFLICT (id) DO NOTHING;

-- reset sequences

SELECT setval('plans_id_seq', (SELECT COALESCE(MAX(id), 1) FROM plans));

