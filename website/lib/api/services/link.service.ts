import {
    ShortLink,
    ShortLinkWithCreator,
    CreateShortLinkInput,
    ApiResponse,
    PaginatedResponse,
} from "@/lib/types";
import {
    fakeLinks,
    simulateDelay,
    generateShortCode,
    getLinksByTenantId,
} from "@/lib/data/fake-data";

// ============================================
// Link Service - Uses fake data for now
// ============================================

class LinkService {
    // In-memory store for new links during session
    private sessionLinks: ShortLinkWithCreator[] = [...fakeLinks];

    async getLinks(
        tenantId: string,
        page: number = 1,
        pageSize: number = 10
    ): Promise<ApiResponse<PaginatedResponse<ShortLinkWithCreator>>> {
        await simulateDelay(500);

        const allLinks = this.sessionLinks.filter((l) => l.tenantId === tenantId);
        const total = allLinks.length;
        const totalPages = Math.ceil(total / pageSize);
        const start = (page - 1) * pageSize;
        const items = allLinks.slice(start, start + pageSize);

        return {
            success: true,
            data: {
                items,
                total,
                page,
                pageSize,
                totalPages,
            },
        };
    }

    async getLinkById(linkId: string): Promise<ApiResponse<ShortLinkWithCreator>> {
        await simulateDelay(300);

        const link = this.sessionLinks.find((l) => l.id === linkId);

        if (!link) {
            return {
                success: false,
                error: "Link not found",
            };
        }

        return {
            success: true,
            data: link,
        };
    }

    async createLink(
        tenantId: string,
        userId: string,
        input: CreateShortLinkInput
    ): Promise<ApiResponse<ShortLinkWithCreator>> {
        await simulateDelay(600);

        // Validate URL
        try {
            new URL(input.originalUrl);
        } catch {
            return {
                success: false,
                error: "Invalid URL format",
            };
        }

        // Check if custom alias is taken
        if (input.customAlias) {
            const existing = this.sessionLinks.find(
                (l) => l.shortCode === input.customAlias || l.customAlias === input.customAlias
            );
            if (existing) {
                return {
                    success: false,
                    error: "This alias is already taken",
                };
            }
        }

        const newLink: ShortLinkWithCreator = {
            id: `link-${Date.now()}`,
            tenantId,
            originalUrl: input.originalUrl,
            shortCode: input.customAlias || generateShortCode(),
            customAlias: input.customAlias,
            clicks: 0,
            isActive: true,
            expiresAt: input.expiresAt,
            createdAt: new Date(),
            updatedAt: new Date(),
            createdBy: userId,
            creator: {
                id: userId,
                name: "Current User",
                email: "user@example.com",
                avatar: `https://api.dicebear.com/7.x/avataaars/svg?seed=${userId}`,
            },
        };

        this.sessionLinks.unshift(newLink);

        return {
            success: true,
            data: newLink,
            message: "Link created successfully",
        };
    }

    async updateLink(
        linkId: string,
        updates: Partial<ShortLink>
    ): Promise<ApiResponse<ShortLinkWithCreator>> {
        await simulateDelay(500);

        const index = this.sessionLinks.findIndex((l) => l.id === linkId);

        if (index === -1) {
            return {
                success: false,
                error: "Link not found",
            };
        }

        this.sessionLinks[index] = {
            ...this.sessionLinks[index],
            ...updates,
            updatedAt: new Date(),
        };

        return {
            success: true,
            data: this.sessionLinks[index],
            message: "Link updated successfully",
        };
    }

    async deleteLink(linkId: string): Promise<ApiResponse<void>> {
        await simulateDelay(400);

        const index = this.sessionLinks.findIndex((l) => l.id === linkId);

        if (index === -1) {
            return {
                success: false,
                error: "Link not found",
            };
        }

        this.sessionLinks.splice(index, 1);

        return {
            success: true,
            message: "Link deleted successfully",
        };
    }

    async toggleLinkStatus(linkId: string): Promise<ApiResponse<ShortLinkWithCreator>> {
        await simulateDelay(300);

        const link = this.sessionLinks.find((l) => l.id === linkId);

        if (!link) {
            return {
                success: false,
                error: "Link not found",
            };
        }

        link.isActive = !link.isActive;
        link.updatedAt = new Date();

        return {
            success: true,
            data: link,
        };
    }

    // Quick create for anonymous users (landing page)
    async quickCreate(originalUrl: string): Promise<ApiResponse<{ shortUrl: string; shortCode: string }>> {
        await simulateDelay(800);

        try {
            new URL(originalUrl);
        } catch {
            return {
                success: false,
                error: "Invalid URL format",
            };
        }

        const shortCode = generateShortCode();
        const baseUrl = typeof window !== "undefined" ? window.location.origin : "https://golink.io";

        return {
            success: true,
            data: {
                shortUrl: `${baseUrl}/${shortCode}`,
                shortCode,
            },
        };
    }
}

// Export singleton instance
export const linkService = new LinkService();
