"use client";

import { useEffect, useState } from "react";
import { User } from "@/lib/types";
import { authService } from "@/lib/api/services/auth.service";
import { ProfileForm } from "@/components/shared/profile-form";
import { Loader2 } from "lucide-react";

export default function AdminProfilePage() {
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
            <div className="flex h-full items-center justify-center p-8">
                <Loader2 className="h-8 w-8 animate-spin text-muted-foreground" />
            </div>
        );
    }

    if (!user) {
        return (
            <div className="p-8">
                <div className="rounded-lg border border-red-200 bg-red-50 p-4 text-red-800 dark:border-red-900/50 dark:bg-red-900/10 dark:text-red-200">
                    Error loading user profile. Please try logging in again.
                </div>
            </div>
        );
    }

    return (
        <div className="space-y-4">
            <div className="flex items-center justify-between">
                <h2 className="text-3xl font-bold tracking-tight">Profile</h2>
            </div>
            <div className="space-y-4">
                <ProfileForm user={user} />
            </div>
        </div>
    );
}
