"use client";

import Link from "next/link";
import { useTheme } from "next-themes";
import { useLanguage } from "@/lib/i18n";
import { Button } from "@/components/ui/button";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Zap, Moon, Sun, Globe } from "lucide-react";
import { useEffect, useState } from "react";

export function Navbar() {
    const { theme, setTheme } = useTheme();
    const { language, setLanguage, t } = useLanguage();
    const [mounted, setMounted] = useState(false);

    useEffect(() => {
        setMounted(true);
    }, []);

    return (
        <header className="fixed top-4 left-1/2 -translate-x-1/2 z-50 w-[95%] max-w-6xl">
            <nav className="flex items-center rounded-full border border-border/50 bg-background/80 dark:bg-zinc-900/80 backdrop-blur-xl px-4 py-2 shadow-lg">
                {/* Logo */}
                <Link href="/" className="flex items-center gap-2 group">
                    <div className="flex h-9 w-9 items-center justify-center rounded-full bg-gradient-to-br from-[var(--gold)] to-amber-500 text-zinc-900 transition-transform group-hover:scale-105">
                        <Zap className="h-5 w-5" />
                    </div>
                    <span className="text-lg font-bold tracking-tight">GoLink</span>
                </Link>

                {/* Spacer */}
                <div className="flex-1" />

                {/* Right - All navigation and actions */}
                <div className="flex items-center gap-1">
                    {/* Home */}
                    <Link
                        href="/"
                        className="px-4 py-2 text-sm font-medium text-muted-foreground hover:text-foreground transition-colors rounded-full hover:bg-muted"
                    >
                        {t("nav.home")}
                    </Link>

                    {/* Login - Special golden button */}
                    <Button
                        asChild
                        className="rounded-full px-5 font-semibold cursor-pointer bg-[var(--gold)] hover:bg-[var(--gold)]/90 text-[var(--gold-foreground)] ml-1"
                    >
                        <Link href="/auth">{t("nav.login")}</Link>
                    </Button>

                    {/* Divider */}
                    <div className="w-px h-5 bg-border mx-2" />

                    {/* Language Switcher */}
                    <DropdownMenu>
                        <DropdownMenuTrigger asChild>
                            <Button variant="ghost" size="icon" className="h-9 w-9 rounded-full cursor-pointer">
                                <Globe className="h-4 w-4" />
                            </Button>
                        </DropdownMenuTrigger>
                        <DropdownMenuContent align="end">
                            <DropdownMenuItem
                                onClick={() => setLanguage("en")}
                                className={`cursor-pointer ${language === "en" ? "bg-muted" : ""}`}
                            >
                                🇺🇸 English
                            </DropdownMenuItem>
                            <DropdownMenuItem
                                onClick={() => setLanguage("vi")}
                                className={`cursor-pointer ${language === "vi" ? "bg-muted" : ""}`}
                            >
                                🇻🇳 Tiếng Việt
                            </DropdownMenuItem>
                        </DropdownMenuContent>
                    </DropdownMenu>

                    {/* Theme Switcher */}
                    {mounted && (
                        <Button
                            variant="ghost"
                            size="icon"
                            className="h-9 w-9 rounded-full cursor-pointer"
                            onClick={() => setTheme(theme === "dark" ? "light" : "dark")}
                        >
                            {theme === "dark" ? (
                                <Sun className="h-4 w-4" />
                            ) : (
                                <Moon className="h-4 w-4" />
                            )}
                        </Button>
                    )}
                </div>
            </nav>
        </header>
    );
}
