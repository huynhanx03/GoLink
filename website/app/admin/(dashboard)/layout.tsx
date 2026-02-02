"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { authService } from "@/lib/api";
import { UserWithTenants } from "@/lib/types";
import { AdminSidebar, AdminHeader } from "@/components/shared";
import { Loader2 } from "lucide-react";

export default function AdminLayout({
    children,
}: {
    children: React.ReactNode;
}) {
    const router = useRouter();
    const [user, setUser] = useState<UserWithTenants | null>(null);
    const [isLoading, setIsLoading] = useState(true);

    useEffect(() => {
        const checkAuth = async () => {
            const result = await authService.getCurrentUser();

            if (!result.success || !result.data) {
                router.push("/admin/login");
                return;
            }

            if (result.data.role !== "admin") {
                router.push("/admin/login");
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

    return (
        <div className="min-h-screen bg-background">
            <AdminSidebar
                user={{
                    name: user.name,
                    email: user.email || "",
                    avatar: user.avatar,
                }}
            />
            <AdminHeader />

            <main className="pl-64 pt-14">
                <div className="p-6">{children}</div>
            </main>
        </div>
    );
}
