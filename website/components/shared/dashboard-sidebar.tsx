"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { useLanguage } from "@/lib/i18n";
import { Button } from "@/components/ui/button";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import {
    Zap,
    LayoutDashboard,
    Link2,
    Settings,
    LogOut,
    Building2,
    User,
} from "lucide-react";

interface DashboardSidebarProps {
    user: {
        name: string;
        email: string;
        avatar?: string;
    };
    workspaces: {
        id: string;
        name: string;
        slug: string;
    }[];
    currentWorkspace?: {
        id: string;
        name: string;
        slug: string;
    };
}

export function DashboardSidebar({
    user,
    workspaces,
    currentWorkspace,
}: DashboardSidebarProps) {
    const pathname = usePathname();
    const { t } = useLanguage();

    const navigation = currentWorkspace
        ? [
            {
                name: t("user.sidebar.overview"),
                href: `/workspace/${currentWorkspace.slug}`,
                icon: LayoutDashboard,
            },
            {
                name: t("user.sidebar.links"),
                href: `/workspace/${currentWorkspace.slug}/links`,
                icon: Link2,
            },
            {
                name: t("user.sidebar.settings"),
                href: `/workspace/${currentWorkspace.slug}/settings`,
                icon: Settings,
            },
        ]
        : [];

    const isActive = (href: string) => pathname === href;
    const isWorkspaceSelected = (slug: string) => pathname.startsWith(`/workspace/${slug}`);

    return (
        <aside className="fixed left-0 top-0 bottom-0 w-64 border-r border-border bg-slate-100 dark:bg-zinc-950 flex flex-col">
            {/* Logo */}
            <div className="p-6 border-b border-border">
                <Link href="/dashboard" className="flex items-center gap-3 group">
                    <div className="relative flex h-10 w-10 items-center justify-center rounded-xl bg-gradient-to-br from-[var(--gold)] to-[#B8860B] shadow-lg shadow-[var(--gold)]/20 transition-all duration-300 group-hover:shadow-xl group-hover:shadow-[var(--gold)]/30 group-hover:scale-105">
                        <Zap className="h-5 w-5 text-[var(--gold-foreground)] transition-transform duration-300 group-hover:scale-110" />
                        <div className="absolute inset-0 rounded-xl bg-white/20 opacity-0 group-hover:opacity-100 transition-opacity duration-300" />
                    </div>
                    <span className="text-xl font-bold tracking-tight text-foreground">GoLink</span>
                </Link>
            </div>

            {/* Workspaces + Navigation Combined */}
            <nav className="flex-1 p-4 overflow-y-auto">
                {/* Workspaces Header */}
                <div className="px-3 py-2 mb-2">
                    <span className="text-xs font-semibold text-muted-foreground uppercase tracking-wider">
                        {t("user.sidebar.workspaces")}
                    </span>
                </div>

                <div className="space-y-1">
                    {workspaces.map((workspace) => {
                        const isSelected = isWorkspaceSelected(workspace.slug);
                        return (
                            <div key={workspace.id}>
                                <Link
                                    href={`/workspace/${workspace.slug}`}
                                    className={`flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm transition-colors ${isSelected
                                        ? "bg-[var(--gold)] text-[var(--gold-foreground)] font-semibold"
                                        : "text-muted-foreground hover:bg-muted hover:text-foreground font-medium"
                                        }`}
                                >
                                    <Building2 className="h-4 w-4" />
                                    <span className="truncate">{workspace.name}</span>
                                </Link>

                                {/* Show navigation under selected workspace */}
                                {isSelected && (
                                    <div className="mt-2 ml-7 space-y-0.5">
                                        {navigation.map((item) => (
                                            <Link
                                                key={item.name}
                                                href={item.href}
                                                className={`flex items-center gap-2.5 px-3 py-2 rounded-lg text-sm transition-colors ${isActive(item.href)
                                                    ? "bg-black/10 dark:bg-white/10 text-foreground font-medium"
                                                    : "text-muted-foreground hover:bg-black/5 dark:hover:bg-white/5 hover:text-foreground"
                                                    }`}
                                            >
                                                <item.icon className="h-4 w-4" />
                                                {item.name}
                                            </Link>
                                        ))}
                                    </div>
                                )}
                            </div>
                        );
                    })}
                </div>
            </nav>

            {/* User */}
            <div className="p-4 border-t border-border">
                <DropdownMenu>
                    <DropdownMenuTrigger asChild>
                        <Button
                            variant="ghost"
                            className="w-full justify-start h-auto py-3 px-4 cursor-pointer"
                        >
                            <Avatar className="h-8 w-8 mr-3">
                                <AvatarImage src={user.avatar} alt={user.name} />
                                <AvatarFallback className="bg-muted text-muted-foreground">
                                    {user.name
                                        .split(" ")
                                        .map((n) => n[0])
                                        .join("")}
                                </AvatarFallback>
                            </Avatar>
                            <div className="text-left flex-1 min-w-0">
                                <div className="font-medium text-sm truncate">{user.name}</div>
                                <div className="text-xs text-muted-foreground truncate">
                                    {user.email}
                                </div>
                            </div>
                        </Button>
                    </DropdownMenuTrigger>
                    <DropdownMenuContent className="w-56" align="end" side="top">
                        <DropdownMenuItem asChild className="cursor-pointer">
                            <Link href="/settings/profile">
                                <User className="mr-2 h-4 w-4" />
                                {t("user.sidebar.profile")}
                            </Link>
                        </DropdownMenuItem>
                        <DropdownMenuSeparator />
                        <DropdownMenuItem
                            className="text-destructive cursor-pointer"
                            onClick={() => {
                                // Handle logout
                            }}
                        >
                            <LogOut className="mr-2 h-4 w-4" />
                            {t("user.sidebar.logout")}
                        </DropdownMenuItem>
                    </DropdownMenuContent>
                </DropdownMenu>
            </div>
        </aside>
    );
}
