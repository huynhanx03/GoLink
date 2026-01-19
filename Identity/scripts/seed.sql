-- create table

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    is_admin BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS tenants (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    tier_id INTEGER,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    level INTEGER DEFAULT 0,
    lft INTEGER DEFAULT 0,
    rgt INTEGER DEFAULT 0,
    parent_id INTEGER DEFAULT -1,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS resources (
    id SERIAL PRIMARY KEY,
    key VARCHAR(255) UNIQUE NOT NULL,
    description VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL PRIMARY KEY,
    role_id INTEGER NOT NULL REFERENCES roles(id),
    resource_id INTEGER NOT NULL REFERENCES resources(id),
    description VARCHAR(255),
    scopes INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS tenant_members (
    id SERIAL PRIMARY KEY,
    tenant_id INTEGER NOT NULL REFERENCES tenants(id),
    user_id INTEGER NOT NULL REFERENCES users(id),
    role_id INTEGER NOT NULL REFERENCES roles(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS domains (
    id SERIAL PRIMARY KEY,
    tenant_id INTEGER NOT NULL REFERENCES tenants(id),
    domain VARCHAR(255) UNIQUE NOT NULL,
    is_verified BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS credentials (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    type VARCHAR(255) NOT NULL,
    credential_data JSONB NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS federated_identities (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    provider VARCHAR(255) NOT NULL,
    provider_user_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(provider, provider_user_id)
);

CREATE TABLE IF NOT EXISTS attribute_definitions (
    id SERIAL PRIMARY KEY,
    key VARCHAR(255) UNIQUE NOT NULL,
    data_type VARCHAR(50) DEFAULT 'string',
    description VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_attribute_values (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    attribute_id INTEGER NOT NULL REFERENCES attribute_definitions(id),
    value TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(user_id, attribute_id)
);

-- insert data

INSERT INTO resources (id, key, description, created_at, updated_at) VALUES
(1, 'generation', 'Link generation management', NOW(), NOW()),
(2, 'domains', 'Domain management', NOW(), NOW())
ON CONFLICT (id) DO NOTHING;

INSERT INTO attribute_definitions (id, key, data_type, description, created_at, updated_at) VALUES
(1, 'first_name', 'string', 'First name', NOW(), NOW()),
(2, 'last_name', 'string', 'Last name', NOW(), NOW()),
(3, 'gender', 'int', 'Gender (0: Female, 1: Male, 2: Other)', NOW(), NOW()),
(4, 'date_of_birth', 'date', 'Date of birth', NOW(), NOW())
ON CONFLICT (id) DO NOTHING;

INSERT INTO roles (id, name, level, lft, rgt, parent_id, created_at, updated_at) VALUES
(1, 'owner',  100, 1, 6, -1, NOW(), NOW()),
(2, 'admin',   50, 2, 5,  1, NOW(), NOW()),
(3, 'member',  10, 3, 4,  2, NOW(), NOW())
ON CONFLICT (id) DO NOTHING;

INSERT INTO permissions (id, role_id, resource_id, description, scopes, created_at, updated_at) VALUES
(1, 1, 2, 'Owner: Create domain', 1, NOW(), NOW()),
(2, 1, 2, 'Owner: Read domain', 2, NOW(), NOW()),
(3, 1, 2, 'Owner: Update domain', 4, NOW(), NOW()),
(4, 1, 2, 'Owner: Delete domain', 8, NOW(), NOW()),
(5, 3, 1, 'Member: Create generation', 1, NOW(), NOW()),
(6, 3, 1, 'Member: Read generation', 2, NOW(), NOW()),
(7, 3, 1, 'Member: Delete generation', 8, NOW(), NOW())
ON CONFLICT (id) DO NOTHING;

-- reset sequences

SELECT setval('users_id_seq', (SELECT COALESCE(MAX(id), 1) FROM users));
SELECT setval('roles_id_seq', (SELECT COALESCE(MAX(id), 1) FROM roles));
SELECT setval('resources_id_seq', (SELECT COALESCE(MAX(id), 1) FROM resources));
SELECT setval('attribute_definitions_id_seq', (SELECT COALESCE(MAX(id), 1) FROM attribute_definitions));
SELECT setval('user_attribute_values_id_seq', (SELECT COALESCE(MAX(id), 1) FROM user_attribute_values));
SELECT setval('permissions_id_seq', (SELECT COALESCE(MAX(id), 1) FROM permissions));
