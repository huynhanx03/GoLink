"use client";

import { useState, useEffect } from "react";
import { Switch } from "@/components/ui/switch";
import { Label } from "@/components/ui/label";
import { Badge } from "@/components/ui/badge";
import { PlanCard } from "./plan-card";
import { useLanguage } from "@/lib/i18n";
import { planService } from "@/lib/api";
import { Plan } from "@/lib/types";
import { Loader2 } from "lucide-react";

interface PricingSectionProps {
    onSelectPlan?: (plan: Plan) => void;
    currentPlanId?: string;
    showTitle?: boolean;
    className?: string;
}

export function PricingSection({
    onSelectPlan,
    currentPlanId,
    showTitle = true,
    className = "",
}: PricingSectionProps) {
    const { t } = useLanguage();
    const [plans, setPlans] = useState<Plan[]>([]);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchPlans = async () => {
            setIsLoading(true);
            setError(null);

            // Now returns Plan[] directly (Clean Architecture)
            const result = await planService.getActivePlans();

            if (result.success && result.data) {
                setPlans(result.data);
            } else {
                setError(result.error || "Failed to load plans");
            }

            setIsLoading(false);
        };

        fetchPlans();
    }, []);

    if (isLoading) {
        return (
            <section className={className}>
                <div className="flex justify-center items-center py-20">
                    <Loader2 className="h-8 w-8 animate-spin text-muted-foreground" />
                </div>
            </section>
        );
    }

    if (error) {
        return (
            <section className={className}>
                <div className="text-center py-20">
                    <p className="text-muted-foreground">{error}</p>
                </div>
            </section>
        );
    }

    return (
        <section className={className}>
            {showTitle && (
                <div className="text-center mb-12">
                    <h2 className="text-3xl md:text-4xl font-bold mb-4">
                        {t("pricing.title")}
                    </h2>
                    <p className="text-lg text-muted-foreground max-w-2xl mx-auto">
                        {t("pricing.subtitle")}
                    </p>
                </div>
            )}

            {/* Plans Grid */}
            <div className="grid md:grid-cols-3 gap-6 lg:gap-8">
                {plans.map((plan) => (
                    <PlanCard
                        key={plan.id}
                        plan={plan}
                        isCurrentPlan={currentPlanId === plan.id}
                    />
                ))}
            </div>
        </section>
    );
}
