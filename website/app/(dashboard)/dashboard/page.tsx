"use client";

import { useEffect, useState } from "react";
import Link from "next/link";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { authService, tenantService } from "@/lib/api";
import { TenantWithPlan, UserWithTenants } from "@/lib/types";
import {
    Building2,
    Plus,
    ArrowRight,
    Link2,
    Users,
    Loader2,
} from "lucide-react";

export default function DashboardPage() {
    const [user, setUser] = useState<UserWithTenants | null>(null);
    const [tenants, setTenants] = useState<TenantWithPlan[]>([]);
    const [isLoading, setIsLoading] = useState(true);

    useEffect(() => {
        const loadData = async () => {
            const userResult = await authService.getCurrentUser();
            if (userResult.success && userResult.data) {
                setUser(userResult.data);

                const tenantsResult = await tenantService.getUserTenants(userResult.data.id);
                if (tenantsResult.success && tenantsResult.data) {
                    setTenants(tenantsResult.data);
                }
            }
            setIsLoading(false);
        };

        loadData();
    }, []);

    if (isLoading) {
        return (
            <div className="flex items-center justify-center h-64">
                <Loader2 className="h-8 w-8 animate-spin text-muted-foreground" />
            </div>
        );
    }

    return (
        <div className="max-w-5xl">
            {/* Header */}
            <div className="mb-8">
                <h1 className="text-3xl font-bold mb-2">
                    Welcome back, {user?.name?.split(" ")[0]}!
                </h1>
                <p className="text-muted-foreground">
                    Select a workspace to manage your short links
                </p>
            </div>

            {/* Workspaces Grid */}
            <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
                {tenants.map((tenant) => (
                    <Link key={tenant.id} href={`/workspace/${tenant.slug}`}>
                        <Card className="p-6 h-full hover:shadow-lg hover:border-[var(--gold)]/50 transition-all duration-200 cursor-pointer group">
                            {/* Header */}
                            <div className="flex items-start justify-between mb-4">
                                <div className="h-12 w-12 rounded-xl bg-muted flex items-center justify-center group-hover:bg-[var(--gold)]/10 transition-colors">
                                    <Building2 className="h-6 w-6 text-foreground group-hover:text-[var(--gold)]" />
                                </div>
                                <Badge
                                    variant={tenant.plan.tier === "free" ? "secondary" : "default"}
                                    className={
                                        tenant.plan.tier === "pro"
                                            ? "bg-[var(--gold)] text-[var(--gold-foreground)]"
                                            : ""
                                    }
                                >
                                    {tenant.plan.name}
                                </Badge>
                            </div>

                            {/* Name */}
                            <h3 className="text-lg font-semibold mb-1">{tenant.name}</h3>
                            <p className="text-sm text-muted-foreground mb-4">
                                /{tenant.slug}
                            </p>

                            {/* Stats */}
                            <div className="flex items-center gap-4 text-sm text-muted-foreground">
                                <div className="flex items-center gap-1">
                                    <Link2 className="h-4 w-4" />
                                    <span>{tenant.linkCount} links</span>
                                </div>
                                <div className="flex items-center gap-1">
                                    <Users className="h-4 w-4" />
                                    <span>{tenant.memberCount} members</span>
                                </div>
                            </div>

                            {/* Arrow */}
                            <div className="mt-4 pt-4 border-t border-border flex items-center justify-between">
                                <span className="text-sm font-medium text-muted-foreground group-hover:text-foreground transition-colors">
                                    Open workspace
                                </span>
                                <ArrowRight className="h-4 w-4 text-muted-foreground group-hover:text-foreground group-hover:translate-x-1 transition-all" />
                            </div>
                        </Card>
                    </Link>
                ))}

                {/* Create New Workspace */}
                <Link href="/dashboard/new-workspace">
                    <Card className="p-6 h-full border-dashed hover:border-[var(--gold)] hover:bg-[var(--gold)]/5 transition-all duration-200 cursor-pointer flex flex-col items-center justify-center text-center min-h-[200px]">
                        <div className="h-12 w-12 rounded-xl bg-muted flex items-center justify-center mb-4">
                            <Plus className="h-6 w-6" />
                        </div>
                        <h3 className="font-semibold mb-1">Create Workspace</h3>
                        <p className="text-sm text-muted-foreground">
                            Start a new workspace for your team
                        </p>
                    </Card>
                </Link>
            </div>

            {/* Empty State */}
            {tenants.length === 0 && (
                <Card className="p-12 text-center border-dashed">
                    <Building2 className="h-12 w-12 mx-auto mb-4 text-muted-foreground" />
                    <h3 className="text-lg font-semibold mb-2">No workspaces yet</h3>
                    <p className="text-muted-foreground mb-6 max-w-md mx-auto">
                        Create your first workspace to start managing short links for your
                        team or project.
                    </p>
                    <Button asChild className="bg-[var(--gold)] hover:bg-[var(--gold)]/90 text-[var(--gold-foreground)]">
                        <Link href="/dashboard/new-workspace">
                            <Plus className="mr-2 h-4 w-4" />
                            Create Workspace
                        </Link>
                    </Button>
                </Card>
            )}
        </div>
    );
}
