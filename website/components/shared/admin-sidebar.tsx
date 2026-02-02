"use client";

import { useState } from "react";
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
    Shield,
    LayoutDashboard,
    Users,
    CreditCard,
    Settings,
    User,
    LogOut,
    ChevronDown,
    ChevronRight,
    Lock,
    Tags,
} from "lucide-react";

interface AdminSidebarProps {
    user: {
        name: string;
        email: string;
        avatar?: string;
    };
}

interface NavItem {
    name: string;
    href: string;
    icon: React.ComponentType<{ className?: string }>;
    children?: { name: string; href: string; icon?: React.ComponentType<{ className?: string }> }[];
}

export function AdminSidebar({ user }: AdminSidebarProps) {
    const pathname = usePathname();
    const { t } = useLanguage();
    const [expandedItems, setExpandedItems] = useState<string[]>(["Settings"]);

    const navigation: NavItem[] = [
        { name: t("admin.sidebar.dashboard"), href: "/admin", icon: LayoutDashboard },
        { name: t("admin.sidebar.users"), href: "/admin/users", icon: Users },
        { name: t("admin.sidebar.plans"), href: "/admin/plans", icon: CreditCard },
        {
            name: t("admin.sidebar.settings"),
            href: "/admin/settings",
            icon: Settings,
            children: [
                { name: t("admin.sidebar.permissions"), href: "/admin/settings/permissions", icon: Lock },
                { name: t("admin.sidebar.attributes"), href: "/admin/settings/attributes", icon: Tags },
            ]
        },
    ];

    const isActive = (href: string) => {
        if (href === "/admin") {
            return pathname === "/admin";
        }
        return pathname === href || pathname.startsWith(href + "/");
    };

    const isParentActive = (item: NavItem) => {
        if (item.children) {
            return item.children.some(child => isActive(child.href));
        }
        return isActive(item.href);
    };

    const toggleExpand = (name: string) => {
        setExpandedItems(prev =>
            prev.includes(name)
                ? prev.filter(n => n !== name)
                : [...prev, name]
        );
    };

    return (
        <aside className="fixed left-0 top-0 bottom-0 w-64 bg-card border-r border-border flex flex-col">
            {/* Logo */}
            <div className="p-6 border-b border-border">
                <Link href="/admin" className="flex items-center gap-3">
                    <div className="flex h-10 w-10 items-center justify-center rounded-xl bg-[var(--gold)] text-[var(--gold-foreground)]">
                        <Shield className="h-5 w-5" />
                    </div>
                    <div>
                        <span className="text-lg font-bold text-foreground">GoLink</span>
                        <div className="text-xs text-muted-foreground uppercase tracking-wider">Admin</div>
                    </div>
                </Link>
            </div>

            {/* Navigation */}
            <nav className="flex-1 p-4 space-y-1 overflow-y-auto">
                {navigation.map((item) => (
                    <div key={item.name}>
                        {item.children ? (
                            <>
                                <button
                                    onClick={() => toggleExpand(item.name)}
                                    className={`w-full flex items-center justify-between gap-3 px-4 py-3 rounded-lg text-sm font-medium transition-colors ${isParentActive(item)
                                            ? "bg-muted text-foreground"
                                            : "text-muted-foreground hover:bg-muted hover:text-foreground"
                                        }`}
                                >
                                    <div className="flex items-center gap-3">
                                        <item.icon className="h-5 w-5" />
                                        {item.name}
                                    </div>
                                    {expandedItems.includes(item.name) ? (
                                        <ChevronDown className="h-4 w-4" />
                                    ) : (
                                        <ChevronRight className="h-4 w-4" />
                                    )}
                                </button>
                                {expandedItems.includes(item.name) && (
                                    <div className="ml-4 mt-1 space-y-1 border-l border-border pl-4">
                                        {item.children.map((child) => (
                                            <Link
                                                key={child.href}
                                                href={child.href}
                                                className={`flex items-center gap-3 px-4 py-2 rounded-lg text-sm transition-colors ${isActive(child.href)
                                                        ? "bg-[var(--gold)] text-[var(--gold-foreground)]"
                                                        : "text-muted-foreground hover:bg-muted hover:text-foreground"
                                                    }`}
                                            >
                                                {child.icon && <child.icon className="h-4 w-4" />}
                                                {child.name}
                                            </Link>
                                        ))}
                                    </div>
                                )}
                            </>
                        ) : (
                            <Link
                                href={item.href}
                                className={`flex items-center gap-3 px-4 py-3 rounded-lg text-sm font-medium transition-colors ${isActive(item.href)
                                        ? "bg-[var(--gold)] text-[var(--gold-foreground)]"
                                        : "text-muted-foreground hover:bg-muted hover:text-foreground"
                                    }`}
                            >
                                <item.icon className="h-5 w-5" />
                                {item.name}
                            </Link>
                        )}
                    </div>
                ))}
            </nav>

            {/* User */}
            <div className="p-4 border-t border-border">
                <DropdownMenu>
                    <DropdownMenuTrigger asChild>
                        <Button
                            variant="ghost"
                            className="w-full justify-start h-auto py-3 px-4 text-foreground hover:bg-muted cursor-pointer"
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
                                <div className="text-xs text-muted-foreground truncate">{user.email}</div>
                            </div>
                        </Button>
                    </DropdownMenuTrigger>
                    <DropdownMenuContent className="w-56" align="end" side="top">
                        <DropdownMenuLabel className="text-muted-foreground">{t("admin.sidebar.account")}</DropdownMenuLabel>
                        <DropdownMenuSeparator />
                        <DropdownMenuItem asChild className="cursor-pointer">
                            <Link href="/admin/profile">
                                <User className="mr-2 h-4 w-4" />
                                {t("admin.sidebar.profile")}
                            </Link>
                        </DropdownMenuItem>
                        <DropdownMenuSeparator />
                        <DropdownMenuItem className="text-destructive cursor-pointer">
                            <LogOut className="mr-2 h-4 w-4" />
                            {t("admin.sidebar.logout")}
                        </DropdownMenuItem>
                    </DropdownMenuContent>
                </DropdownMenu>
            </div>
        </aside>
    );
}
