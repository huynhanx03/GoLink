import { ApiResponse, PaginatedResponse } from "@/lib/types";

// ============================================
// API Configuration
// ============================================

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || "/api";

// ============================================
// Base API Client
// ============================================

interface RequestOptions extends RequestInit {
    params?: Record<string, string | number | boolean | undefined>;
}

class ApiClient {
    private baseUrl: string;

    constructor(baseUrl: string) {
        this.baseUrl = baseUrl;
    }

    private buildUrl(endpoint: string, params?: Record<string, string | number | boolean | undefined>): string {
        const url = new URL(`${this.baseUrl}${endpoint}`, window.location.origin);

        if (params) {
            Object.entries(params).forEach(([key, value]) => {
                if (value !== undefined) {
                    url.searchParams.append(key, String(value));
                }
            });
        }

        return url.toString();
    }

    async request<T>(endpoint: string, options: RequestOptions = {}): Promise<ApiResponse<T>> {
        const { params, ...fetchOptions } = options;
        const url = this.buildUrl(endpoint, params);

        const defaultHeaders: HeadersInit = {
            "Content-Type": "application/json",
        };

        try {
            const response = await fetch(url, {
                ...fetchOptions,
                headers: {
                    ...defaultHeaders,
                    ...fetchOptions.headers,
                },
            });

            const data = await response.json();

            if (!response.ok) {
                return {
                    success: false,
                    error: data.message || data.error || "An error occurred",
                };
            }

            return {
                success: true,
                data: data as T,
            };
        } catch (error) {
            return {
                success: false,
                error: error instanceof Error ? error.message : "Network error",
            };
        }
    }

    async get<T>(endpoint: string, params?: Record<string, string | number | boolean | undefined>): Promise<ApiResponse<T>> {
        return this.request<T>(endpoint, { method: "GET", params });
    }

    async post<T>(endpoint: string, body?: unknown): Promise<ApiResponse<T>> {
        return this.request<T>(endpoint, {
            method: "POST",
            body: body ? JSON.stringify(body) : undefined,
        });
    }

    async put<T>(endpoint: string, body?: unknown): Promise<ApiResponse<T>> {
        return this.request<T>(endpoint, {
            method: "PUT",
            body: body ? JSON.stringify(body) : undefined,
        });
    }

    async patch<T>(endpoint: string, body?: unknown): Promise<ApiResponse<T>> {
        return this.request<T>(endpoint, {
            method: "PATCH",
            body: body ? JSON.stringify(body) : undefined,
        });
    }

    async delete<T>(endpoint: string): Promise<ApiResponse<T>> {
        return this.request<T>(endpoint, { method: "DELETE" });
    }
}

// Export singleton instance
export const apiClient = new ApiClient(API_BASE_URL);

// Re-export for convenience
export type { ApiResponse, PaginatedResponse };
