"use client";

import { useEffect, useState } from "react";
import { useRouter, usePathname } from "next/navigation";
import { authService } from "@/lib/api";
import { UserWithTenants } from "@/lib/types";
import { DashboardSidebar, DashboardHeader } from "@/components/shared";
import { Loader2 } from "lucide-react";

export default function DashboardLayout({
    children,
}: {
    children: React.ReactNode;
}) {
    const router = useRouter();
    const pathname = usePathname();
    const [user, setUser] = useState<UserWithTenants | null>(null);
    const [isLoading, setIsLoading] = useState(true);

    useEffect(() => {
        const checkAuth = async () => {
            const result = await authService.getCurrentUser();

            if (!result.success || !result.data) {
                router.push("/login");
                return;
            }

            setUser(result.data);
            setIsLoading(false);
        };

        checkAuth();
    }, [router]);

    if (isLoading) {
        return (
            <div className="min-h-screen flex items-center justify-center bg-background">
                <Loader2 className="h-8 w-8 animate-spin text-muted-foreground" />
            </div>
        );
    }

    if (!user) {
        return null;
    }

    const workspaces = user.tenants.map((t) => ({
        id: t.tenantId,
        name: t.tenantName,
        slug: t.tenantSlug,
    }));

    // Extract current workspace from URL
    const workspaceSlugMatch = pathname.match(/\/workspace\/([^/]+)/);
    const currentWorkspaceSlug = workspaceSlugMatch ? workspaceSlugMatch[1] : null;
    const currentWorkspace = currentWorkspaceSlug
        ? workspaces.find((w) => w.slug === currentWorkspaceSlug)
        : undefined;

    return (
        <div className="min-h-screen bg-background">
            <DashboardSidebar
                user={{
                    name: user.name,
                    email: user.email || "",
                    avatar: user.avatar,
                }}
                workspaces={workspaces}
                currentWorkspace={currentWorkspace}
            />
            <DashboardHeader />

            <main className="pl-64 pt-14 min-h-screen bg-white dark:bg-zinc-900">
                <div className="p-6">{children}</div>
            </main>
        </div>
    );
}
