"use client";

import { Card } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Feature, Plan } from "@/lib/types";
import { Check } from "lucide-react";
import { useLanguage } from "@/lib/i18n";

interface PlanCardProps {
    plan: Plan;
    isCurrentPlan?: boolean;
}

export function PlanCard({
    plan,
    isCurrentPlan = false,
}: PlanCardProps) {
    const { t } = useLanguage();

    // Helper to render feature text
    const renderFeature = (feature: Feature): string => {
        if (typeof feature === 'string') return feature;
        if (feature.key === 'pricing.feature.backend_raw') {
            return feature.fallback || feature.params?.text?.toString() || "";
        }
        return t(feature.key, feature.params);
    };

    return (
        <Card
            className={`relative p-6 md:p-8 flex flex-col h-full transition-all duration-300 ${plan.isPopular
                ? "border-[var(--gold)] border-2 shadow-lg shadow-[var(--gold)]/10 scale-[1.02]"
                : "border-border hover:border-[var(--gold)]/50"
                } ${isCurrentPlan ? "bg-muted/50" : ""}`}
        >
            {plan.isPopular && (
                <Badge className="absolute -top-3 left-1/2 -translate-x-1/2 bg-[var(--gold)] text-[var(--gold-foreground)] hover:bg-[var(--gold)]">
                    {t("pricing.popular")}
                </Badge>
            )}

            {/* Header */}
            <div className="mb-6">
                <h3 className="text-xl font-semibold mb-2">{plan.name}</h3>
                <p className="text-sm text-muted-foreground">{plan.description}</p>
            </div>

            {/* Price */}
            <div className="mb-6">
                <div className="flex items-baseline gap-1">
                    <span className="text-4xl font-bold">
                        ${plan.price === 0 ? "0" : plan.price}
                    </span>
                    <span className="text-muted-foreground">/mo</span>
                </div>
            </div>

            {/* Features List */}
            <div className="flex-1">
                <h4 className="font-medium mb-4">{t("pricing.whatsIncluded")}</h4>
                <ul className="space-y-3">
                    {plan.features.map((feature, index) => (
                        <li key={index} className="flex items-start gap-3">
                            <Check className="h-5 w-5 text-[var(--gold)] shrink-0 mt-0.5" />
                            <span className="text-sm text-muted-foreground">{renderFeature(feature)}</span>
                        </li>
                    ))}
                </ul>
            </div>
        </Card>
    );
}
