"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { getAccessToken, decodeTokenPayload } from "@/lib/api/token";
import { AdminSidebar, AdminHeader } from "@/components/shared";
import { Loader2 } from "lucide-react";

interface AdminClaims {
    is_admin: boolean;
    user_id: number;
    username: string;
}

export default function AdminLayout({
    children,
}: {
    children: React.ReactNode;
}) {
    const router = useRouter();
    const [claims, setClaims] = useState<AdminClaims | null>(null);
    const [isLoading, setIsLoading] = useState(true);

    useEffect(() => {
        const token = getAccessToken();
        if (!token) {
            router.push("/admin/login");
            return;
        }

        const decoded = decodeTokenPayload<AdminClaims>(token);
        if (!decoded?.is_admin) {
            router.push("/admin/login");
            return;
        }

        setClaims(decoded);
        setIsLoading(false);
    }, [router]);

    if (isLoading || !claims) {
        return (
            <div className="min-h-screen flex items-center justify-center bg-background">
                <Loader2 className="h-8 w-8 animate-spin text-muted-foreground" />
            </div>
        );
    }

    return (
        <div className="min-h-screen bg-background">
            <AdminSidebar
                user={{
                    name: claims.username,
                    email: "",
                    avatar: undefined,
                }}
            />
            <AdminHeader />

            <main className="pl-64 pt-14">
                <div className="p-6">{children}</div>
            </main>
        </div>
    );
}
