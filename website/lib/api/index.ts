// ============================================
// API Module - Centralized exports
// ============================================

// Config
export { getBaseUrl, buildUrl, endpoints } from "./config";

// HTTP Client
export { http } from "./http";

// Token Management
export {
    getAccessToken,
    getRefreshToken,
    getTenantIdFromUrl,
    setTokens,
    updateAccessToken,
    clearTokens,
    hasTokens,
} from "./token";

// Services
export { authService } from "./services/auth.service";
export { planService } from "./services/plan.service";

// Keep existing services for backward compatibility
export { linkService } from "./services/link.service";
export { tenantService } from "./services/tenant.service";

// Types
export type {
    BaseResponse,
    PaginatedBaseResponse,
    Pagination,
    ApiResult,
    ApiPaginatedResult,
    LoginRequest,
    LoginResponseDto,
    RegisterRequest,
    RegisterResponseDto,
    RefreshTokenRequest,
    RefreshTokenResponseDto,
    PlanDto,
    UserDto,
    TenantMembershipDto,
} from "./types";
