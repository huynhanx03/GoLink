import {
    User,
    UserWithTenants,
    Tenant,
    TenantWithPlan,
    Plan,
    ShortLinkWithCreator,
    AdminStats,
    DashboardStats,
    TenantMembership,
    Role,
    Resource,
    Permission,
    AttributeDefinition,
} from "@/lib/types";

// User requested to remove fake data. 
// Exporting empty arrays to satisfy imports without providing data.

export const fakeUsers: UserWithTenants[] = [];

export const fakePlans: Plan[] = [];

export const fakeTenants: TenantWithPlan[] = [];

export const fakeLinks: ShortLinkWithCreator[] = [];

export const fakeDashboardStats: Record<string, DashboardStats> = {};

export const fakeAdminStats: AdminStats = {
    totalUsers: 0,
    totalTenants: 0,
    totalLinks: 0,
    totalClicks: 0,
    activeSubscriptions: 0,
    monthlyRevenue: 0,
};

export const fakeRoles: Role[] = [];

export const fakeResources: Resource[] = [];

export const fakePermissions: Permission[] = [];

export const fakeAttributeDefinitions: AttributeDefinition[] = [];

// Helper functions that return undefined or empty arrays
export function simulateDelay(ms: number = 500): Promise<void> {
    return new Promise((resolve) => setTimeout(resolve, ms));
}

export function generateShortCode(): string {
    return "abcdef";
}

export function findUserByUsername(username: string): UserWithTenants | undefined {
    return undefined;
}

export function findTenantBySlug(slug: string): TenantWithPlan | undefined {
    return undefined;
}

export function getLinksByTenantId(tenantId: string): ShortLinkWithCreator[] {
    return [];
}

export function getPlanById(planId: string): Plan | undefined {
    return undefined;
}
