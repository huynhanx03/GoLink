// ============================================
// Token Management
// Handles access/refresh token storage and retrieval
// ============================================

const ACCESS_TOKEN_KEY = "access_token";
const REFRESH_TOKEN_KEY = "refresh_token";

/**
 * Checks if we're running in a browser environment
 */
function isBrowser(): boolean {
    return typeof window !== "undefined";
}

/**
 * Get access token from storage
 */
export function getAccessToken(): string | null {
    if (!isBrowser()) return null;
    return localStorage.getItem(ACCESS_TOKEN_KEY);
}

/**
 * Get refresh token from storage
 */
export function getRefreshToken(): string | null {
    if (!isBrowser()) return null;
    return localStorage.getItem(REFRESH_TOKEN_KEY);
}

/**
 * Get current tenant ID from URL path
 * URL format: /dashboard/{tenant_id}/...
 */
export function getTenantIdFromUrl(): number | null {
    if (!isBrowser()) return null;

    const pathParts = window.location.pathname.split("/");
    // Expected: ["", "dashboard", "{tenant_id}", ...]
    if (pathParts.length >= 3 && pathParts[1] === "dashboard") {
        const tenantId = parseInt(pathParts[2], 10);
        if (!isNaN(tenantId)) {
            return tenantId;
        }
    }
    return null;
}

/**
 * Store tokens after login
 */
export function setTokens(accessToken: string, refreshToken: string): void {
    if (!isBrowser()) return;
    localStorage.setItem(ACCESS_TOKEN_KEY, accessToken);
    localStorage.setItem(REFRESH_TOKEN_KEY, refreshToken);
}

/**
 * Update access token after refresh
 */
export function updateAccessToken(accessToken: string): void {
    if (!isBrowser()) return;
    localStorage.setItem(ACCESS_TOKEN_KEY, accessToken);
}

/**
 * Clear all auth tokens (logout)
 */
export function clearTokens(): void {
    if (!isBrowser()) return;
    localStorage.removeItem(ACCESS_TOKEN_KEY);
    localStorage.removeItem(REFRESH_TOKEN_KEY);
}

/**
 * Check if user has valid tokens
 */
export function hasTokens(): boolean {
    return !!getAccessToken() && !!getRefreshToken();
}

/**
 * Decode JWT payload without verification (client-side only)
 */
export function decodeTokenPayload<T = Record<string, unknown>>(token: string): T | null {
    try {
        const base64Url = token.split(".")[1];
        if (!base64Url) return null;
        const base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
        const json = decodeURIComponent(
            atob(base64)
                .split("")
                .map((c) => "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2))
                .join("")
        );
        return JSON.parse(json) as T;
    } catch {
        return null;
    }
}
