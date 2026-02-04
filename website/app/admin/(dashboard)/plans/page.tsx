"use client";

import { useState, useEffect } from "react";
import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Switch } from "@/components/ui/switch";
import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { planService } from "@/lib/api";
import { toast } from "sonner";
import { Plus, Pencil, Check, Star } from "lucide-react";

import { useLanguage } from "@/lib/i18n";
import { Feature, Plan } from "@/lib/types";

export default function AdminPlansPage() {
    const [plans, setPlans] = useState<Plan[]>([]);
    const [isLoading, setIsLoading] = useState(true);
    const [isCreateOpen, setIsCreateOpen] = useState(false);
    const { t } = useLanguage();

    useEffect(() => {
        const fetchPlans = async () => {
            try {
                const result = await planService.getActivePlans();
                if (result.success && result.data) {
                    setPlans(result.data);
                }
            } catch (error) {
                toast.error("Failed to load plans");
            } finally {
                setIsLoading(false);
            }
        };
        fetchPlans();
    }, []);

    const handleToggleActive = (planId: string) => {
        // Optimistic update
        setPlans(
            plans.map((p) =>
                p.id === planId ? { ...p, isActive: !p.isActive } : p
            )
        );
        toast.success("Plan status updated");
    };

    const formatFeature = (feature: Feature): string => {
        if (typeof feature === 'string') return feature;
        return t(feature.key, feature.params);
    };

    return (
        <div className="max-w-6xl">
            {/* Header */}
            <div className="flex items-center justify-between mb-8">
                <div>
                    <h1 className="text-3xl font-bold text-white mb-2">Plans</h1>
                    <p className="text-zinc-400">
                        Configure pricing plans and features
                    </p>
                </div>
                <Dialog open={isCreateOpen} onOpenChange={setIsCreateOpen}>
                    <DialogTrigger asChild>
                        <Button className="bg-[var(--gold)] hover:bg-[var(--gold)]/90 text-[var(--gold-foreground)] cursor-pointer">
                            <Plus className="mr-2 h-4 w-4" />
                            Add Plan
                        </Button>
                    </DialogTrigger>
                    <DialogContent className="bg-zinc-900 border-zinc-800">
                        <DialogHeader>
                            <DialogTitle className="text-white">Create New Plan</DialogTitle>
                            <DialogDescription className="text-zinc-400">
                                Add a new pricing plan for your customers.
                            </DialogDescription>
                        </DialogHeader>
                        <div className="space-y-4 py-4">
                            <div className="space-y-2">
                                <Label className="text-zinc-300">Plan Name</Label>
                                <Input
                                    placeholder="e.g., Business"
                                    className="bg-zinc-800 border-zinc-700 text-white"
                                />
                            </div>
                            <div className="grid grid-cols-2 gap-4">
                                <div className="space-y-2">
                                    <Label className="text-zinc-300">Price</Label>
                                    <Input
                                        type="number"
                                        placeholder="29"
                                        className="bg-zinc-800 border-zinc-700 text-white"
                                    />
                                </div>
                                <div className="space-y-2">
                                    <Label className="text-zinc-300">Links/month</Label>
                                    <Input
                                        type="number"
                                        placeholder="5000"
                                        className="bg-zinc-800 border-zinc-700 text-white"
                                    />
                                </div>
                            </div>
                        </div>
                        <DialogFooter>
                            <Button
                                variant="outline"
                                onClick={() => setIsCreateOpen(false)}
                                className="border-zinc-700 text-zinc-300 hover:bg-zinc-800 cursor-pointer"
                            >
                                Cancel
                            </Button>
                            <Button
                                onClick={() => {
                                    setIsCreateOpen(false);
                                    toast.success("Plan created");
                                }}
                                className="bg-[var(--gold)] hover:bg-[var(--gold)]/90 text-[var(--gold-foreground)] cursor-pointer"
                            >
                                Create Plan
                            </Button>
                        </DialogFooter>
                    </DialogContent>
                </Dialog>
            </div>

            {/* Plans Grid */}
            <div className="grid md:grid-cols-3 gap-6">
                {plans.map((plan) => (
                    <Card
                        key={plan.id}
                        className={`p-6 bg-zinc-900 border-zinc-800 relative ${plan.isPopular ? "ring-2 ring-[var(--gold)]" : ""
                            }`}
                    >
                        {plan.isPopular && (
                            <Badge className="absolute -top-3 left-1/2 -translate-x-1/2 bg-[var(--gold)] text-[var(--gold-foreground)]">
                                <Star className="mr-1 h-3 w-3" />
                                Popular
                            </Badge>
                        )}

                        {/* Header */}
                        <div className="flex items-start justify-between mb-6">
                            <div>
                                <h3 className="text-xl font-bold text-white">{plan.name}</h3>
                                <p className="text-sm text-zinc-400 mt-1">{plan.description}</p>
                            </div>
                            <div className="flex items-center gap-2">
                                <Switch
                                    checked={plan.isActive}
                                    onCheckedChange={() => handleToggleActive(plan.id)}
                                />
                            </div>
                        </div>

                        {/* Price */}
                        <div className="mb-6">
                            {plan.price === 99 ? (
                                <div className="text-3xl font-bold text-white">Custom</div>
                            ) : (
                                <div className="flex items-baseline gap-1">
                                    <span className="text-3xl font-bold text-white">
                                        ${plan.price}
                                    </span>
                                    <span className="text-zinc-400">/mo</span>
                                </div>
                            )}
                        </div>

                        {/* Features */}
                        <div className="space-y-2 mb-6 flex-1">
                            {plan.features.slice(0, 5).map((feature, index) => (
                                <div key={index} className="flex items-center gap-2 text-sm">
                                    <Check className="h-4 w-4 text-[var(--gold)] shrink-0" />
                                    <span className="text-zinc-300">
                                        {formatFeature(feature)}
                                    </span>
                                </div>
                            ))}
                            {plan.features.length > 5 && (
                                <p className="text-sm text-zinc-500 pl-6">
                                    +{plan.features.length - 5} more features
                                </p>
                            )}
                        </div>

                        {/* Actions */}
                        <Button
                            variant="outline"
                            className="w-full border-zinc-700 text-zinc-300 hover:bg-zinc-800 cursor-pointer"
                        >
                            <Pencil className="mr-2 h-4 w-4" />
                            Edit Plan
                        </Button>
                    </Card>
                ))}
            </div>
        </div>
    );
}
