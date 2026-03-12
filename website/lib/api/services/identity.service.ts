// ============================================
// Identity Service - Roles, Resources, Permissions
// ============================================

import { http } from "../http";
import { endpoints } from "../config";
import {
    ApiResult,
    RoleDto,
    CreateRoleDto,
    ResourceDto,
    PermissionDto,
    CreatePermissionDto,
    UpdatePermissionDto,
    PaginatedDto,
    QueryOptionsDto,
} from "../types";

class IdentityService {
    // ---- Roles ----

    async findRoles(query?: QueryOptionsDto): Promise<ApiResult<PaginatedDto<RoleDto>>> {
        return http.post<PaginatedDto<RoleDto>>(
            endpoints.identity.roles.find,
            query ?? { pagination: { page: 1, page_size: 100 } }
        );
    }

    async createRole(data: CreateRoleDto): Promise<ApiResult<RoleDto>> {
        return http.post<RoleDto>(endpoints.identity.roles.create, data);
    }

    // ---- Resources ----

    async findResources(query?: QueryOptionsDto): Promise<ApiResult<PaginatedDto<ResourceDto>>> {
        return http.post<PaginatedDto<ResourceDto>>(
            endpoints.identity.resources.find,
            query ?? { pagination: { page: 1, page_size: 100 } }
        );
    }

    // ---- Permissions ----

    async findPermissions(query?: QueryOptionsDto): Promise<ApiResult<PaginatedDto<PermissionDto>>> {
        return http.post<PaginatedDto<PermissionDto>>(
            endpoints.identity.permissions.find,
            query ?? { pagination: { page: 1, page_size: 1000 } }
        );
    }

    async createPermission(data: CreatePermissionDto): Promise<ApiResult<PermissionDto>> {
        return http.post<PermissionDto>(endpoints.identity.permissions.create, data);
    }

    async updatePermission(id: number, data: UpdatePermissionDto): Promise<ApiResult<PermissionDto>> {
        return http.put<PermissionDto>(endpoints.identity.permissions.update(id), data);
    }

    async deletePermission(id: number): Promise<ApiResult<void>> {
        return http.delete<void>(endpoints.identity.permissions.delete(id));
    }
}

export const identityService = new IdentityService();
