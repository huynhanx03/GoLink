// ============================================
// Plan Mapper - DTO to Domain Conversion
// ============================================

import { PlanDto } from "../types";
import { Feature, Plan } from "@/lib/types";

const DEFAULT_PLAN_DESCRIPTION = "Flexible plan for your needs";

export class PlanMapper {
    /**
     * Map DTO to Domain Model
     * Handles type conversions and formats feature strings directly
     */
    static toDomain(dto: PlanDto): Plan {
        return {
            id: dto.id?.toString() || "0",
            name: dto.name || "Unknown Plan",
            description: dto.description || DEFAULT_PLAN_DESCRIPTION,
            price: dto.base_price ?? 0,
            interval: (dto.period === "year" ? "yearly" : "monthly"),
            features: [
                ...this.mapFeaturesToDomain(dto),
            ],
            isActive: dto.is_active ?? false,
            isPopular: false,
            createdAt: dto.created_at ? new Date(dto.created_at) : new Date(),
            updatedAt: dto.updated_at ? new Date(dto.updated_at) : new Date(),
        };
    }

    /**
     * Convert features DTO into user-friendly feature objects
     */
    private static mapFeaturesToDomain(dto: PlanDto): Feature[] {
        const features: Feature[] = [];
        const dtoFeatures = dto.features;

        if (!dtoFeatures) return features;

        if (dtoFeatures.max_links > 0) {
            features.push({
                key: "pricing.feature.links",
                params: { count: dtoFeatures.max_links.toLocaleString() }
            });
        } else if (dtoFeatures.max_links === -1) {
            features.push({ key: "pricing.feature.unlimitedLinks" });
        }

        if (dtoFeatures.customer_domain) {
            features.push({ key: "pricing.feature.customDomains" });
        }
        if (dtoFeatures.ttl && dtoFeatures.ttl > 0) {
            const days = Math.round(dtoFeatures.ttl / 86400);
            features.push({
                key: "pricing.feature.retention",
                params: { days: days }
            });
        } else if (dtoFeatures.ttl === -1) {
            features.push({ key: "pricing.feature.unlimitedRetention" });
        }

        return features;
    }
}
