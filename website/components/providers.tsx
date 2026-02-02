"use client";

import { ThemeProvider as NextThemesProvider } from "next-themes";
import { LanguageProvider } from "@/lib/i18n";

export function Providers({ children }: { children: React.ReactNode }) {
    return (
        <NextThemesProvider
            attribute="class"
            defaultTheme="system"
            enableSystem
            disableTransitionOnChange
        >
            <LanguageProvider>{children}</LanguageProvider>
        </NextThemesProvider>
    );
}
