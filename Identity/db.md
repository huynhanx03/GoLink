erDiagram
    USERS {
        uuid id PK
        string email UK
        string default_tenant_id FK
        bool is_admin
    }
    
    TENANTS {
        string id PK
        string name
        uuid owner_id FK
        string tier_id
    }
    
    TENANT_MEMBERS {
        string tenant_id PK
        uuid user_id PK
        int role_id FK
    }
    
    ROLES {
        int id PK
        string name UK
    }
    
    ROLE_HIERARCHY {
        int parent_role_id PK
        int child_role_id PK
    }
    
    RESOURCES {
        string id PK
        string name UK
        string display_name
    }
    
    PERMISSIONS {
        int id PK
        int role_id FK
        string resource_id FK
        int scopes
    }
    
    CREDENTIALS {
        uuid id PK
        uuid user_id FK
        string type
        jsonb secret_data
    }
    
    USER_ATTRIBUTES {
        uuid id PK
        uuid user_id FK
        string name
        string value
    }
    
    FEDERATED_IDENTITIES {
        uuid user_id PK
        string provider PK
        string external_id
    }
    
    DOMAINS {
        uuid id PK
        string domain UK
        string tenant_id FK
    }
    
    USERS ||--o{ TENANT_MEMBERS : joins
    TENANTS ||--o{ TENANT_MEMBERS : has
    ROLES ||--o{ TENANT_MEMBERS : assigned
    ROLES ||--o{ PERMISSIONS : grants
    ROLES ||--o{ ROLE_HIERARCHY : parent
    ROLES ||--o{ ROLE_HIERARCHY : child
    RESOURCES ||--o{ PERMISSIONS : protects
    USERS ||--o{ CREDENTIALS : auth
    USERS ||--o{ USER_ATTRIBUTES : has
    USERS ||--o{ FEDERATED_IDENTITIES : social
    USERS ||--o| TENANTS : owns
    TENANTS ||--o{ DOMAINS : owns
