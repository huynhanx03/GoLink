// ============================================
// HTTP Client with Token Refresh
// Reusable HTTP client that handles auth automatically
// ============================================

import { getBaseUrl, endpoints } from "./config";
import {
    BaseResponse,
    PaginatedBaseResponse,
    ApiResult,
    ApiPaginatedResult,
    RefreshTokenRequest,
    RefreshTokenResponseDto,
} from "./types";
import {
    getAccessToken,
    getRefreshToken,
    updateAccessToken,
    clearTokens,
    getTenantIdFromUrl,
} from "./token";

// Flag to prevent multiple refresh requests
let isRefreshing = false;
let refreshPromise: Promise<boolean> | null = null;

/**
 * Parse backend response and extract error message
 */
function extractErrorMessage(response: BaseResponse): string {
    return response.message || "An error occurred";
}

/**
 * Attempt to refresh the access token
 */
async function refreshAccessToken(): Promise<boolean> {
    const refreshToken = getRefreshToken();
    const tenantId = getTenantIdFromUrl();

    if (!refreshToken || tenantId === null) {
        clearTokens();
        return false;
    }

    try {
        const response = await fetch(`${getBaseUrl()}${endpoints.auth.refresh}`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: `Bearer ${refreshToken}`,
            },
            body: JSON.stringify({ tenant_id: tenantId } as RefreshTokenRequest),
        });

        if (!response.ok) {
            clearTokens();
            return false;
        }

        const result: BaseResponse<RefreshTokenResponseDto> = await response.json();
        if (result.data?.access_token) {
            updateAccessToken(result.data.access_token);
            return true;
        }

        clearTokens();
        return false;
    } catch {
        clearTokens();
        return false;
    }
}

/**
 * Ensure only one refresh request at a time
 */
async function ensureTokenRefresh(): Promise<boolean> {
    if (isRefreshing && refreshPromise) {
        return refreshPromise;
    }

    isRefreshing = true;
    refreshPromise = refreshAccessToken().finally(() => {
        isRefreshing = false;
        refreshPromise = null;
    });

    return refreshPromise;
}

/**
 * Build request headers with auth token
 */
function buildHeaders(includeAuth: boolean = true): HeadersInit {
    const headers: HeadersInit = {
        "Content-Type": "application/json",
    };

    if (includeAuth) {
        const token = getAccessToken();
        if (token) {
            headers["Authorization"] = `Bearer ${token}`;
        }
    }

    return headers;
}

interface RequestOptions {
    method?: "GET" | "POST" | "PUT" | "PATCH" | "DELETE";
    body?: unknown;
    params?: Record<string, string | number | boolean | undefined>;
    requireAuth?: boolean;
}

/**
 * Build URL with query parameters
 */
function buildUrlWithParams(
    endpoint: string,
    params?: Record<string, string | number | boolean | undefined>
): string {
    const url = new URL(`${getBaseUrl()}${endpoint}`);

    if (params) {
        Object.entries(params).forEach(([key, value]) => {
            if (value !== undefined) {
                url.searchParams.append(key, String(value));
            }
        });
    }

    return url.toString();
}

/**
 * Core request function with automatic token refresh
 */
async function request<T>(
    endpoint: string,
    options: RequestOptions = {}
): Promise<ApiResult<T>> {
    const { method = "GET", body, params, requireAuth = true } = options;
    const url = buildUrlWithParams(endpoint, params);

    const makeRequest = async (): Promise<Response> => {
        return fetch(url, {
            method,
            headers: buildHeaders(requireAuth),
            body: body ? JSON.stringify(body) : undefined,
        });
    };

    try {
        let response = await makeRequest();

        // If 401, try to refresh token and retry
        if (response.status === 401 && requireAuth) {
            const refreshed = await ensureTokenRefresh();
            if (refreshed) {
                response = await makeRequest();
            } else {
                // Redirect to login
                if (typeof window !== "undefined") {
                    window.location.href = "/auth";
                }
                return {
                    success: false,
                    error: "Session expired. Please login again.",
                    code: 401,
                };
            }
        }

        const result: BaseResponse<T> = await response.json();

        if (!response.ok) {
            return {
                success: false,
                error: extractErrorMessage(result),
                code: result.code,
            };
        }

        return {
            success: true,
            data: result.data,
            code: result.code,
        };
    } catch (error) {
        return {
            success: false,
            error: error instanceof Error ? error.message : "Network error",
        };
    }
}

/**
 * Request with pagination support
 */
async function requestPaginated<T>(
    endpoint: string,
    options: RequestOptions = {}
): Promise<ApiPaginatedResult<T>> {
    const { method = "GET", body, params, requireAuth = true } = options;
    const url = buildUrlWithParams(endpoint, params);

    const makeRequest = async (): Promise<Response> => {
        return fetch(url, {
            method,
            headers: buildHeaders(requireAuth),
            body: body ? JSON.stringify(body) : undefined,
        });
    };

    try {
        let response = await makeRequest();

        // If 401, try to refresh token and retry
        if (response.status === 401 && requireAuth) {
            const refreshed = await ensureTokenRefresh();
            if (refreshed) {
                response = await makeRequest();
            } else {
                if (typeof window !== "undefined") {
                    window.location.href = "/auth";
                }
                return {
                    success: false,
                    error: "Session expired. Please login again.",
                    code: 401,
                };
            }
        }

        const result: PaginatedBaseResponse<T> = await response.json();

        if (!response.ok) {
            return {
                success: false,
                error: extractErrorMessage(result),
                code: result.code,
            };
        }

        return {
            success: true,
            data: result.data,
            code: result.code,
            pagination: result.pagination,
        };
    } catch (error) {
        return {
            success: false,
            error: error instanceof Error ? error.message : "Network error",
        };
    }
}

// ============================================
// HTTP Client Export
// ============================================

export const http = {
    /**
     * GET request
     */
    get: <T>(
        endpoint: string,
        params?: Record<string, string | number | boolean | undefined>,
        requireAuth: boolean = true
    ): Promise<ApiResult<T>> => {
        return request<T>(endpoint, { method: "GET", params, requireAuth });
    },

    /**
     * GET request with pagination
     */
    getPaginated: <T>(
        endpoint: string,
        params?: Record<string, string | number | boolean | undefined>,
        requireAuth: boolean = true
    ): Promise<ApiPaginatedResult<T>> => {
        return requestPaginated<T>(endpoint, { method: "GET", params, requireAuth });
    },

    /**
     * POST request
     */
    post: <T>(
        endpoint: string,
        body?: unknown,
        requireAuth: boolean = true
    ): Promise<ApiResult<T>> => {
        return request<T>(endpoint, { method: "POST", body, requireAuth });
    },

    /**
     * PUT request
     */
    put: <T>(
        endpoint: string,
        body?: unknown,
        requireAuth: boolean = true
    ): Promise<ApiResult<T>> => {
        return request<T>(endpoint, { method: "PUT", body, requireAuth });
    },

    /**
     * PATCH request
     */
    patch: <T>(
        endpoint: string,
        body?: unknown,
        requireAuth: boolean = true
    ): Promise<ApiResult<T>> => {
        return request<T>(endpoint, { method: "PATCH", body, requireAuth });
    },

    /**
     * DELETE request
     */
    delete: <T>(endpoint: string, requireAuth: boolean = true): Promise<ApiResult<T>> => {
        return request<T>(endpoint, { method: "DELETE", requireAuth });
    },
};
