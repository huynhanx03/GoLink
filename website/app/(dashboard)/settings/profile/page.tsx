"use client";

import { useEffect, useState } from "react";
import { useLanguage } from "@/lib/i18n";
import { User } from "@/lib/types";
import { authService } from "@/lib/api/services/auth.service";
import { ProfileForm } from "@/components/shared/profile-form";
import { Loader2 } from "lucide-react";

export default function UserProfilePage() {
    const { t } = useLanguage();
    const [user, setUser] = useState<User | null>(null);
    const [isLoading, setIsLoading] = useState(true);

    useEffect(() => {
        const loadUser = async () => {
            try {
                const response = await authService.getCurrentUser();
                if (response.success && response.data) {
                    setUser(response.data);
                }
            } catch (error) {
                console.error("Failed to load user", error);
            } finally {
                setIsLoading(false);
            }
        };

        loadUser();
    }, []);

    if (isLoading) {
        return (
            <div className="flex h-full items-center justify-center">
                <Loader2 className="h-8 w-8 animate-spin text-muted-foreground" />
            </div>
        );
    }

    if (!user) {
        return (
            <div className="rounded-lg border border-red-200 bg-red-50 p-4 text-red-800 dark:border-red-900/50 dark:bg-red-900/10 dark:text-red-200">
                {t("user.profile.error")}
            </div>
        );
    }

    return (
        <div className="space-y-4">
            <div className="space-y-1">
                <h1 className="text-3xl font-bold tracking-tight">{t("user.profile.title")}</h1>
                <p className="text-muted-foreground">
                    {t("user.profile.subtitle")}
                </p>
            </div>
            <ProfileForm user={user} />
        </div>
    );
}
