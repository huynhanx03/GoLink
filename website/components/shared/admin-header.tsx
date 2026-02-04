"use client";

import { useEffect, useState } from "react";
import { useTheme } from "next-themes";
import { useLanguage } from "@/lib/i18n";
import { Button } from "@/components/ui/button";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Moon, Sun, Globe } from "lucide-react";

export function AdminHeader() {
    const { theme, setTheme } = useTheme();
    const { language, setLanguage } = useLanguage();
    const [mounted, setMounted] = useState(false);

    useEffect(() => {
        setMounted(true);
    }, []);

    return (
        <header className="fixed top-0 left-64 right-0 h-14 z-40 border-b border-border bg-background/80 backdrop-blur-xl">
            <div className="h-full flex items-center justify-end px-6 gap-2">
                {/* Language Switcher */}
                <DropdownMenu>
                    <DropdownMenuTrigger asChild>
                        <Button variant="ghost" size="sm" className="h-9 px-3 gap-2 cursor-pointer">
                            <Globe className="h-4 w-4" />
                            <span className="text-sm font-medium">
                                {language === "en" ? "EN" : "VI"}
                            </span>
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
                        className="h-9 w-9 cursor-pointer"
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
        </header>
    );
}
