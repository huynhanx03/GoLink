// ============================================
// API Configuration
// ============================================

/**
 * Returns the base URL for API requests
 * Reads from NEXT_PUBLIC_API_URL environment variable
 */
export function getBaseUrl(): string {
    return process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";
}

/**
 * API Endpoints configuration
 * Centralized location for all API endpoints
 */
export const endpoints = {
    // Auth endpoints
    auth: {
        login: "/identity/auth/login",
        register: "/identity/auth/register",
        refresh: "/identity/auth/refresh",
        logout: "/identity/auth/logout",
        me: "/identity/users/me",
    },
    // Identity management endpoints (admin)
    identity: {
        roles: {
            find: "/identity/roles/find",
            get: (id: number) => `/identity/roles/${id}`,
            create: "/identity/roles",
            update: (id: number) => `/identity/roles/${id}`,
            delete: (id: number) => `/identity/roles/${id}`,
        },
        resources: {
            find: "/identity/resources/find",
            get: (id: number) => `/identity/resources/${id}`,
            create: "/identity/resources",
            update: (id: number) => `/identity/resources/${id}`,
            delete: (id: number) => `/identity/resources/${id}`,
        },
        permissions: {
            find: "/identity/permissions/find",
            get: (id: number) => `/identity/permissions/${id}`,
            create: "/identity/permissions",
            update: (id: number) => `/identity/permissions/${id}`,
            delete: (id: number) => `/identity/permissions/${id}`,
        },
    },
    // Billing endpoints
    billing: {
        activePlans: "/billing/plans/active",
    },
} as const;

/**
 * Build complete URL from endpoint
 */
export function buildUrl(endpoint: string): string {
    return `${getBaseUrl()}${endpoint}`;
}
