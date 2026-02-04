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
