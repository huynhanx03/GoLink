// ============================================
// Plan Service - Real API calls
// ============================================

import { http } from "../http";
import { endpoints } from "../config";
import { ApiResult, PlanDto } from "../types";
import { Plan } from "@/lib/types";
import { PlanMapper } from "../mappers/plan.mapper";

class PlanService {
    /**
     * Get all active plans
     * Uses PlanMapper to convert DTOs to Domain Models
     */
    async getActivePlans(): Promise<ApiResult<Plan[]>> {
        const result = await http.get<PlanDto[]>(
            endpoints.billing.activePlans,
            undefined,
            false
        );

        if (result.success && result.data && Array.isArray(result.data)) {
            const plans = result.data.map((d) => PlanMapper.toDomain(d));
            plans[Math.floor(plans.length / 2)].isPopular = true;
            return { ...result, data: plans };
        }

        return {
            success: false,
            error: result.error,
            code: result.code,
            data: [],
        };
    }
}

// Export singleton instance
export const planService = new PlanService();
