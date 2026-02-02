// ============================================
// API Types - Request/Response types
// ============================================

// ============================================
// Response Wrappers
// ============================================

/**
 * Standard API response format
 */
export interface BaseResponse<T = unknown> {
    code: number;
    message: string;
    data: T;
}

/**
 * Pagination metadata
 */
export interface Pagination {
    page: number;
    page_size: number;
    total: number;
    total_pages: number;
}

/**
 * Paginated API response format
 */
export interface PaginatedBaseResponse<T = unknown> {
    code: number;
    message: string;
    data: T;
    pagination?: Pagination;
}

// ============================================
// Frontend Result Wrappers
// ============================================

export interface ApiResult<T> {
    success: boolean;
    data?: T;
    error?: string;
    code?: number;
}

export interface ApiPaginatedResult<T> extends ApiResult<T> {
    pagination?: Pagination;
}

// ============================================
// Auth DTOs
// ============================================

export interface LoginRequest {
    username: string;
    password: string;
}

export interface LoginResponseDto {
    access_token: string;
    refresh_token: string;
}

export interface RegisterRequest {
    username: string;
    password: string;
    first_name: string;
    last_name: string;
    gender: number;
    birthday: string;
}

export interface RegisterResponseDto {
    success: boolean;
}

export interface RefreshTokenRequest {
    tenant_id: number;
}

export interface RefreshTokenResponseDto {
    access_token: string;
}

// User & Tenant DTOs
export interface TenantMembershipDto {
    tenant_id: string;
    tenant_name: string;
    tenant_slug: string;
    role: "owner" | "admin" | "member";
    joined_at: string;
}

export interface UserDto {
    id: string;
    username: string;
    email?: string;
    name: string;
    first_name?: string;
    last_name?: string;
    avatar?: string;
    role: "user" | "admin";
    created_at: string;
    updated_at: string;
    tenants: TenantMembershipDto[];
}

// ============================================
// Plan DTOs - Matched with Actual Backend Response
// ============================================

export interface PlanFeaturesDto {
    max_links: number;
    ttl: number;                // in seconds
    customer_domain?: boolean;  // boolean flag
}

export interface PlanDto {
    id: number;
    name: string;
    description: string;
    base_price: number;
    period: string;
    features: PlanFeaturesDto;
    is_active: boolean;
    created_at: string;
    updated_at: string;
}
