import {
    Tenant,
    TenantWithPlan,
    DashboardStats,
    ApiResponse,
} from "@/lib/types";
import {
    fakeTenants,
    fakeDashboardStats,
    simulateDelay,
    findTenantBySlug,
} from "@/lib/data/fake-data";

// ============================================
// Tenant Service - Uses fake data for now
// ============================================

class TenantService {
    async getTenantBySlug(slug: string): Promise<ApiResponse<TenantWithPlan>> {
        await simulateDelay(400);

        const tenant = findTenantBySlug(slug);

        if (!tenant) {
            return {
                success: false,
                error: "Workspace not found",
            };
        }

        return {
            success: true,
            data: tenant,
        };
    }

    async getTenantById(id: string): Promise<ApiResponse<TenantWithPlan>> {
        await simulateDelay(400);

        const tenant = fakeTenants.find((t) => t.id === id);

        if (!tenant) {
            return {
                success: false,
                error: "Workspace not found",
            };
        }

        return {
            success: true,
            data: tenant,
        };
    }

    async getUserTenants(userId: string): Promise<ApiResponse<TenantWithPlan[]>> {
        await simulateDelay(500);

        // In fake mode, return tenants based on user's memberships
        // For now, return all tenants for demo
        return {
            success: true,
            data: fakeTenants,
        };
    }

    async getTenantStats(tenantId: string): Promise<ApiResponse<DashboardStats>> {
        await simulateDelay(300);

        const stats = fakeDashboardStats[tenantId] || {
            totalLinks: 0,
            totalClicks: 0,
            activeLinks: 0,
            linksThisMonth: 0,
        };

        return {
            success: true,
            data: stats,
        };
    }

    async createTenant(
        name: string,
        ownerId: string
    ): Promise<ApiResponse<TenantWithPlan>> {
        await simulateDelay(800);

        const slug = name.toLowerCase().replace(/\s+/g, "-");

        // Check if slug exists
        const existing = findTenantBySlug(slug);
        if (existing) {
            return {
                success: false,
                error: "A workspace with this name already exists",
            };
        }

        // In fake mode, just return a mock tenant
        const newTenant: TenantWithPlan = {
            id: `tenant-${Date.now()}`,
            name,
            slug,
            planId: "plan-free",
            ownerId,
            createdAt: new Date(),
            updatedAt: new Date(),
            plan: {
                id: "plan-free",
                name: "Free",
                tier: "free",
                description: "Perfect for personal use",
                price: 0,
                interval: "monthly",
                limits: {
                    linksPerMonth: 50,
                    clicksPerLink: 1000,
                    customDomains: 0,
                    teamMembers: 1,
                    analyticsRetentionDays: 7,
                },
                features: ["50 links per month", "Basic analytics"],
                isActive: true,
                createdAt: new Date(),
                updatedAt: new Date(),
            },
            memberCount: 1,
            linkCount: 0,
        };

        return {
            success: true,
            data: newTenant,
            message: "Workspace created successfully",
        };
    }

    async updateTenant(
        tenantId: string,
        updates: Partial<Tenant>
    ): Promise<ApiResponse<TenantWithPlan>> {
        await simulateDelay(500);

        const tenant = fakeTenants.find((t) => t.id === tenantId);

        if (!tenant) {
            return {
                success: false,
                error: "Workspace not found",
            };
        }

        // In fake mode, return merged data
        const updated: TenantWithPlan = {
            ...tenant,
            ...updates,
            updatedAt: new Date(),
        };

        return {
            success: true,
            data: updated,
            message: "Workspace updated successfully",
        };
    }

    async deleteTenant(tenantId: string): Promise<ApiResponse<void>> {
        await simulateDelay(600);

        const tenant = fakeTenants.find((t) => t.id === tenantId);

        if (!tenant) {
            return {
                success: false,
                error: "Workspace not found",
            };
        }

        return {
            success: true,
            message: "Workspace deleted successfully",
        };
    }
}

// Export singleton instance
export const tenantService = new TenantService();
