"use client";

import { useState } from "react";
import Link from "next/link";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useLanguage } from "@/lib/i18n";
import { linkService } from "@/lib/api";
import { toast } from "sonner";
import { Link2, ArrowRight, Copy, Check, Loader2, RefreshCw } from "lucide-react";

interface UrlShortenerWidgetProps {
    variant?: "hero" | "compact";
    className?: string;
}

export function UrlShortenerWidget({ variant = "hero", className = "" }: UrlShortenerWidgetProps) {
    const { t } = useLanguage();
    const [url, setUrl] = useState("");
    const [shortUrl, setShortUrl] = useState("");
    const [isLoading, setIsLoading] = useState(false);
    const [isCopied, setIsCopied] = useState(false);

    const handleShorten = async () => {
        if (!url) {
            toast.error("Please enter a URL");
            return;
        }

        // Basic URL validation
        try {
            new URL(url);
        } catch {
            toast.error("Please enter a valid URL");
            return;
        }

        setIsLoading(true);
        try {
            const result = await linkService.quickCreate(url);
            if (result.success && result.data) {
                setShortUrl(`golink.io/${result.data.shortCode}`);
                toast.success("URL shortened successfully!");
            } else {
                toast.error(result.error || "Failed to shorten URL");
            }
        } catch {
            toast.error("Something went wrong");
        } finally {
            setIsLoading(false);
        }
    };

    const handleCopy = async () => {
        try {
            await navigator.clipboard.writeText(`https://${shortUrl}`);
            setIsCopied(true);
            toast.success("Copied to clipboard!");
            setTimeout(() => setIsCopied(false), 2000);
        } catch {
            toast.error("Failed to copy");
        }
    };

    const handleReset = () => {
        setUrl("");
        setShortUrl("");
        setIsCopied(false);
    };

    const isHero = variant === "hero";

    return (
        <div className={className}>
            {/* Form container */}
            <div
                className={`relative rounded-2xl border border-border/50 bg-background/50 dark:bg-zinc-900/50 backdrop-blur-xl overflow-hidden ${isHero ? "p-4" : "p-3"
                    }`}
            >
                {/* Decorative accent line */}
                <div className="absolute top-0 left-0 right-0 h-px bg-gradient-to-r from-transparent via-[var(--gold)]/50 to-transparent" />

                <div className="flex flex-col sm:flex-row gap-3">
                    {/* Input - shows URL or shortUrl */}
                    <div className="flex-1 relative">
                        <Link2 className="absolute left-4 top-1/2 -translate-y-1/2 h-5 w-5 text-muted-foreground" />
                        {shortUrl ? (
                            <div className="flex items-center pl-12 pr-4 h-12 rounded-xl border border-border/50 bg-background/50 dark:bg-zinc-800/50 text-sm font-medium">
                                {shortUrl}
                            </div>
                        ) : (
                            <Input
                                placeholder={t("shortener.placeholder")}
                                value={url}
                                onChange={(e) => setUrl(e.target.value)}
                                onKeyDown={(e) => e.key === "Enter" && handleShorten()}
                                className="pl-12 h-12 text-sm rounded-xl border-border/50 bg-background/50 dark:bg-zinc-800/50"
                            />
                        )}
                    </div>

                    {/* Button - Shorten or Copy */}
                    {shortUrl ? (
                        <div className="flex gap-2">
                            <Button
                                onClick={handleCopy}
                                className="cursor-pointer bg-[var(--gold)] hover:bg-[var(--gold)]/90 text-[var(--gold-foreground)] font-medium h-12 px-5 rounded-xl"
                            >
                                {isCopied ? (
                                    <>
                                        <Check className="mr-2 h-4 w-4" />
                                        {t("shortener.copied")}
                                    </>
                                ) : (
                                    <>
                                        <Copy className="mr-2 h-4 w-4" />
                                        {t("shortener.copy")}
                                    </>
                                )}
                            </Button>
                            <Button
                                variant="outline"
                                onClick={handleReset}
                                className="cursor-pointer h-12 w-12 rounded-xl"
                            >
                                <RefreshCw className="h-4 w-4" />
                            </Button>
                        </div>
                    ) : (
                        <Button
                            onClick={handleShorten}
                            disabled={isLoading || !url}
                            className="cursor-pointer bg-[var(--gold)] hover:bg-[var(--gold)]/90 text-[var(--gold-foreground)] font-medium h-12 px-6 rounded-xl"
                        >
                            {isLoading ? (
                                <Loader2 className="h-5 w-5 animate-spin" />
                            ) : (
                                <>
                                    {t("shortener.button")}
                                    <ArrowRight className="ml-2 h-5 w-5" />
                                </>
                            )}
                        </Button>
                    )}
                </div>
            </div>

            {/* Free limit notice */}
            {isHero && (
                <p className="text-center text-sm text-muted-foreground mt-4">
                    {t("shortener.freeLimit")}{" "}
                    <Link
                        href="/auth"
                        className="text-[var(--gold)] hover:underline font-semibold"
                    >
                        {t("shortener.signUp")}
                    </Link>{" "}
                    {t("shortener.forUnlimited")}
                </p>
            )}
        </div>
    );
}
