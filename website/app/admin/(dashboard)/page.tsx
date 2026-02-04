"use client";

import { useEffect, useState } from "react";
import { useLanguage } from "@/lib/i18n";
import { Card } from "@/components/ui/card";
import { fakeAdminStats } from "@/lib/data/fake-data";
import { AdminStats } from "@/lib/types";
import {
    Users,
    Building2,
    Link2,
    MousePointerClick,
    CreditCard,
    DollarSign,
    TrendingUp,
} from "lucide-react";

export default function AdminDashboardPage() {
    const { t } = useLanguage();
    const [stats, setStats] = useState<AdminStats | null>(null);

    useEffect(() => {
        // Simulate loading
        setTimeout(() => {
            setStats(fakeAdminStats);
        }, 500);
    }, []);

    if (!stats) {
        return null;
    }

    const statCards = [
        {
            label: t("admin.totalUsers"),
            value: stats.totalUsers.toLocaleString(),
            icon: Users,
            color: "bg-blue-500/10 text-blue-500",
            change: "+12%",
        },
        {
            label: t("admin.totalTenants"),
            value: stats.totalTenants.toLocaleString(),
            icon: Building2,
            color: "bg-purple-500/10 text-purple-500",
            change: "+8%",
        },
        {
            label: t("admin.totalLinks"),
            value: stats.totalLinks.toLocaleString(),
            icon: Link2,
            color: "bg-green-500/10 text-green-500",
            change: "+24%",
        },
        {
            label: t("admin.totalClicks"),
            value: (stats.totalClicks / 1000000).toFixed(1) + "M",
            icon: MousePointerClick,
            color: "bg-orange-500/10 text-orange-500",
            change: "+18%",
        },
        {
            label: t("admin.activeSubscriptions"),
            value: stats.activeSubscriptions.toLocaleString(),
            icon: CreditCard,
            color: "bg-pink-500/10 text-pink-500",
            change: "+5%",
        },
        {
            label: t("admin.monthlyRevenue"),
            value: "$" + stats.monthlyRevenue.toLocaleString(),
            icon: DollarSign,
            color: "bg-[var(--gold)]/10 text-[var(--gold)]",
            change: "+15%",
        },
    ];

    return (
        <div className="max-w-6xl">
            {/* Header */}
            <div className="mb-6">
                <h1 className="text-3xl font-bold text-foreground mb-2">{t("admin.dashboard")}</h1>
                <p className="text-muted-foreground">
                    {t("admin.overview")}
                </p>
            </div>

            {/* Stats Grid */}
            <div className="grid sm:grid-cols-2 lg:grid-cols-3 gap-6">
                {statCards.map((stat) => (
                    <Card
                        key={stat.label}
                        className="p-6 bg-card border-border hover:border-muted-foreground/50 transition-colors"
                    >
                        <div className="flex items-start justify-between mb-4">
                            <div className={`h-12 w-12 rounded-xl ${stat.color} flex items-center justify-center`}>
                                <stat.icon className="h-6 w-6" />
                            </div>
                            <div className="flex items-center gap-1 text-green-500 text-sm">
                                <TrendingUp className="h-4 w-4" />
                                {stat.change}
                            </div>
                        </div>
                        <p className="text-sm text-muted-foreground mb-1">{stat.label}</p>
                        <p className="text-2xl font-bold text-foreground">{stat.value}</p>
                    </Card>
                ))}
            </div>
        </div>
    );
}
