"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import Image from "next/image";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { authService } from "@/lib/api";
import { toast } from "sonner";
import { Loader2, Lock, User, Shield } from "lucide-react";
import { Navbar } from "@/components/shared";
import { useLanguage } from "@/lib/i18n";

export default function AdminLoginPage() {
    const router = useRouter();
    const { t } = useLanguage();
    const [isLoading, setIsLoading] = useState(false);
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();

        if (!username || !password) {
            toast.error("Please fill in all fields");
            return;
        }

        setIsLoading(true);

        const result = await authService.login({ username, password });

        if (result.success && result.data) {
            if (result.data.role !== "admin") {
                toast.error("Access denied. Admin credentials required.");
                setIsLoading(false);
                return;
            }
            toast.success("Welcome, Admin!");
            router.push("/admin");
        } else {
            toast.error(result.error || "Login failed");
        }

        setIsLoading(false);
    };

    return (
        <div className="min-h-screen relative">
            {/* Background Image */}
            <div className="fixed inset-0 -z-10">
                <Image
                    src="/admin-bg.png"
                    alt="Background"
                    fill
                    className="object-cover"
                    priority
                />
                <div className="absolute inset-0 bg-black/60 backdrop-blur-sm" />
            </div>

            {/* Navbar */}
            <Navbar />

            {/* Centered Admin Login Card */}
            <div className="min-h-screen flex items-start justify-center px-4 pt-32 pb-8">
                <div className="w-full max-w-md bg-background/95 backdrop-blur-xl rounded-3xl border border-border/50 shadow-2xl p-8">
                    {/* Admin Header */}
                    <div className="flex items-center justify-center gap-3 mb-8">
                        <div className="flex h-12 w-12 items-center justify-center rounded-xl bg-[var(--gold)] text-[var(--gold-foreground)]">
                            <Shield className="h-6 w-6" />
                        </div>
                        <div>
                            <span className="text-2xl font-bold tracking-tight">GoLink</span>
                            <div className="text-xs text-muted-foreground uppercase tracking-wider">{t("admin.portal")}</div>
                        </div>
                    </div>

                    {/* Title */}
                    <div className="text-center mb-8">
                        <h1 className="text-2xl font-bold mb-2">{t("admin.access")}</h1>
                        <p className="text-muted-foreground">
                            {t("admin.signInCredentials")}
                        </p>
                    </div>

                    {/* Login Form */}
                    <form onSubmit={handleSubmit} className="space-y-4">
                        <div className="space-y-2">
                            <Label htmlFor="username">{t("auth.username")}</Label>
                            <div className="relative">
                                <User className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
                                <Input
                                    id="username"
                                    type="text"
                                    placeholder={t("admin.enterUsername")}
                                    value={username}
                                    onChange={(e) => setUsername(e.target.value)}
                                    disabled={isLoading}
                                    className="pl-10"
                                    required
                                />
                            </div>
                        </div>

                        <div className="space-y-2">
                            <Label htmlFor="password">{t("auth.password")}</Label>
                            <div className="relative">
                                <Lock className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
                                <Input
                                    id="password"
                                    type="password"
                                    placeholder="••••••••"
                                    value={password}
                                    onChange={(e) => setPassword(e.target.value)}
                                    disabled={isLoading}
                                    className="pl-10"
                                    required
                                />
                            </div>
                        </div>

                        <Button
                            type="submit"
                            className="w-full bg-[var(--gold)] hover:bg-[var(--gold)]/90 text-[var(--gold-foreground)] font-semibold cursor-pointer mt-6"
                            disabled={isLoading}
                        >
                            {isLoading ? (
                                <Loader2 className="h-5 w-5 animate-spin" />
                            ) : (
                                t("admin.accessPanel")
                            )}
                        </Button>
                    </form>
                </div>
            </div>
        </div>
    );
}
