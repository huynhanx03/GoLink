// ============================================
// Auth Service - Real API calls
// ============================================

import { http } from "../http";
import { endpoints } from "../config";
import {
    ApiResult,
    LoginRequest,
    LoginResponseDto,
    RegisterRequest,
    RegisterResponseDto,
    UserDto,
} from "../types";
import { setTokens, clearTokens, hasTokens } from "../token";
import { UserWithTenants, TenantMembership } from "@/lib/types";

class AuthService {
    /**
     * Login with username and password
     */
    async login(credentials: LoginRequest): Promise<ApiResult<LoginResponseDto>> {
        const result = await http.post<LoginResponseDto>(
            endpoints.auth.login,
            credentials,
            false // No auth required for login
        );

        if (result.success && result.data) {
            setTokens(result.data.access_token, result.data.refresh_token);
        }

        return result;
    }

    /**
     * Register new user
     */
    async register(data: RegisterRequest): Promise<ApiResult<RegisterResponseDto>> {
        return http.post<RegisterResponseDto>(
            endpoints.auth.register,
            data,
            false // No auth required for register
        );
    }

    /**
     * Logout - clear tokens
     */
    async logout(): Promise<ApiResult<void>> {
        clearTokens();
        return {
            success: true,
        };
    }

    /**
     * Check if user is authenticated
     */
    isAuthenticated(): boolean {
        return hasTokens();
    }

    /**
     * Get current user profile
     */
    async getCurrentUser(): Promise<ApiResult<UserWithTenants>> {
        const result = await http.get<UserDto>(endpoints.auth.me);

        if (!result.success || !result.data) {
            return {
                success: false,
                error: result.error,
                code: result.code,
            };
        }

        // Map DTO to domain model
        const userDto = result.data;
        const user: UserWithTenants = {
            id: userDto.id,
            username: userDto.username,
            email: userDto.email,
            name: userDto.name || `${userDto.first_name || ""} ${userDto.last_name || ""}`.trim(),
            avatar: userDto.avatar,
            role: userDto.role,
            createdAt: new Date(userDto.created_at),
            updatedAt: new Date(userDto.updated_at),
            tenants: (userDto.tenants || []).map((t): TenantMembership => ({
                tenantId: t.tenant_id,
                tenantName: t.tenant_name,
                tenantSlug: t.tenant_slug,
                role: t.role,
                joinedAt: new Date(t.joined_at),
            })),
        };

        return {
            success: true,
            data: user,
        };
    }
}

// Export singleton instance
export const authService = new AuthService();
