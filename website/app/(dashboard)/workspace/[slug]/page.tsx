"use client";

import { useEffect, useState } from "react";
import { useParams } from "next/navigation";
import { useLanguage } from "@/lib/i18n";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Input } from "@/components/ui/input";
import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog";
import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table";
import { Label } from "@/components/ui/label";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { tenantService, linkService } from "@/lib/api";
import { TenantWithPlan, ShortLinkWithCreator, DashboardStats } from "@/lib/types";
import { toast } from "sonner";
import {
    Plus,
    Copy,
    Check,
    BarChart3,
    Link2,
    MousePointerClick,
    TrendingUp,
    Loader2,
    ExternalLink,
    MoreHorizontal,
    Pencil,
    Trash2,
    Power,
    Globe,
    Users,
    ArrowRight,
} from "lucide-react";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";

// Fake members data
const fakeMembers = [
    { id: "1", name: "John Doe", email: "john@example.com", role: "Owner", avatar: "" },
    { id: "2", name: "Jane Smith", email: "jane@example.com", role: "Admin", avatar: "" },
    { id: "3", name: "Bob Wilson", email: "bob@example.com", role: "Member", avatar: "" },
];

export default function WorkspacePage() {
    const params = useParams();
    const slug = params.slug as string;
    const { t } = useLanguage();

    const [tenant, setTenant] = useState<TenantWithPlan | null>(null);
    const [stats, setStats] = useState<DashboardStats | null>(null);
    const [links, setLinks] = useState<ShortLinkWithCreator[]>([]);
    const [isLoading, setIsLoading] = useState(true);
    const [isCreateOpen, setIsCreateOpen] = useState(false);
    const [newUrl, setNewUrl] = useState("");
    const [newAlias, setNewAlias] = useState("");
    const [isCreating, setIsCreating] = useState(false);
    const [copiedId, setCopiedId] = useState<string | null>(null);

    useEffect(() => {
        const loadData = async () => {
            const tenantResult = await tenantService.getTenantBySlug(slug);
            if (tenantResult.success && tenantResult.data) {
                setTenant(tenantResult.data);

                const [statsResult, linksResult] = await Promise.all([
                    tenantService.getTenantStats(tenantResult.data.id),
                    linkService.getLinks(tenantResult.data.id),
                ]);

                if (statsResult.success && statsResult.data) {
                    setStats(statsResult.data);
                }
                if (linksResult.success && linksResult.data) {
                    setLinks(linksResult.data.items);
                }
            }
            setIsLoading(false);
        };

        loadData();
    }, [slug]);

    const handleCreateLink = async () => {
        if (!newUrl || !tenant) {
            toast.error("Please enter a URL");
            return;
        }

        setIsCreating(true);

        const result = await linkService.createLink(tenant.id, "user-1", {
            originalUrl: newUrl,
            customAlias: newAlias || undefined,
        });

        if (result.success && result.data) {
            setLinks([result.data, ...links]);
            setNewUrl("");
            setNewAlias("");
            setIsCreateOpen(false);
            toast.success("Link created successfully!");
        } else {
            toast.error(result.error || "Failed to create link");
        }

        setIsCreating(false);
    };

    const handleCopy = async (link: ShortLinkWithCreator) => {
        const baseUrl = window.location.origin;
        const shortUrl = `${baseUrl}/${link.shortCode}`;

        try {
            await navigator.clipboard.writeText(shortUrl);
            setCopiedId(link.id);
            toast.success("Copied to clipboard!");
            setTimeout(() => setCopiedId(null), 2000);
        } catch {
            toast.error("Failed to copy");
        }
    };

    const handleToggleStatus = async (link: ShortLinkWithCreator) => {
        const result = await linkService.toggleLinkStatus(link.id);
        if (result.success && result.data) {
            setLinks(links.map((l) => (l.id === link.id ? result.data! : l)));
            toast.success(result.data.isActive ? "Link activated" : "Link deactivated");
        }
    };

    const handleDelete = async (linkId: string) => {
        const result = await linkService.deleteLink(linkId);
        if (result.success) {
            setLinks(links.filter((l) => l.id !== linkId));
            toast.success("Link deleted");
        }
    };

    if (isLoading) {
        return (
            <div className="flex items-center justify-center h-64">
                <Loader2 className="h-8 w-8 animate-spin text-muted-foreground" />
            </div>
        );
    }

    if (!tenant) {
        return (
            <div className="text-center py-12">
                <h2 className="text-xl font-semibold mb-2">{t("workspace.notFound")}</h2>
                <p className="text-muted-foreground">
                    {t("workspace.notFoundDesc")}
                </p>
            </div>
        );
    }

    const baseUrl = typeof window !== "undefined" ? window.location.origin : "";

    return (
        <div className="space-y-6">
            {/* Header */}
            <div className="flex items-center justify-between">
                <div>
                    <div className="flex items-center gap-3 mb-1">
                        <h1 className="text-3xl font-bold">{tenant.name}</h1>
                        <Badge
                            variant={tenant.plan.tier === "free" ? "secondary" : "default"}
                            className={
                                tenant.plan.tier === "pro"
                                    ? "bg-[var(--gold)] text-[var(--gold-foreground)]"
                                    : ""
                            }
                        >
                            {tenant.plan.name}
                        </Badge>
                    </div>
                    <p className="text-muted-foreground">
                        {t("workspace.manageLinks")}
                    </p>
                </div>

                <Dialog open={isCreateOpen} onOpenChange={setIsCreateOpen}>
                    <DialogTrigger asChild>
                        <Button className="bg-[var(--gold)] hover:bg-[var(--gold)]/90 text-[var(--gold-foreground)] cursor-pointer">
                            <Plus className="mr-2 h-4 w-4" />
                            {t("workspace.createLink")}
                        </Button>
                    </DialogTrigger>
                    <DialogContent>
                        <DialogHeader>
                            <DialogTitle>{t("workspace.createShortLink")}</DialogTitle>
                            <DialogDescription>
                                {t("workspace.createShortLinkDesc")}
                            </DialogDescription>
                        </DialogHeader>
                        <div className="space-y-4 py-4">
                            <div className="space-y-2">
                                <Label htmlFor="url">{t("workspace.destinationUrl")}</Label>
                                <Input
                                    id="url"
                                    placeholder="https://example.com/very-long-url"
                                    value={newUrl}
                                    onChange={(e) => setNewUrl(e.target.value)}
                                />
                            </div>
                            <div className="space-y-2">
                                <Label htmlFor="alias">{t("workspace.customAlias")}</Label>
                                <div className="flex items-center gap-2">
                                    <span className="text-sm text-muted-foreground">
                                        {baseUrl}/
                                    </span>
                                    <Input
                                        id="alias"
                                        placeholder="my-link"
                                        value={newAlias}
                                        onChange={(e) => setNewAlias(e.target.value)}
                                        className="flex-1"
                                    />
                                </div>
                            </div>
                        </div>
                        <DialogFooter>
                            <Button
                                variant="outline"
                                onClick={() => setIsCreateOpen(false)}
                                className="cursor-pointer"
                            >
                                {t("common.cancel")}
                            </Button>
                            <Button
                                onClick={handleCreateLink}
                                disabled={isCreating}
                                className="bg-[var(--gold)] hover:bg-[var(--gold)]/90 text-[var(--gold-foreground)] cursor-pointer"
                            >
                                {isCreating ? (
                                    <Loader2 className="h-4 w-4 animate-spin" />
                                ) : (
                                    t("workspace.createLink")
                                )}
                            </Button>
                        </DialogFooter>
                    </DialogContent>
                </Dialog>
            </div>

            {/* Domain & Members Row */}
            <div className="grid md:grid-cols-2 gap-4">
                {/* Domain Card */}
                <Card className="p-5">
                    <div className="flex items-center justify-between">
                        <div className="flex items-center gap-3">
                            <div className="h-10 w-10 rounded-lg bg-purple-500/10 flex items-center justify-center">
                                <Globe className="h-5 w-5 text-purple-500" />
                            </div>
                            <div>
                                <p className="text-sm text-muted-foreground">{t("workspace.domain")}</p>
                                <p className="font-medium">{baseUrl}/{tenant.slug}</p>
                            </div>
                        </div>
                        <Button variant="outline" size="sm" className="cursor-pointer">
                            <Pencil className="h-3.5 w-3.5 mr-1.5" />
                            {t("workspace.edit")}
                        </Button>
                    </div>
                </Card>

                {/* Members Card */}
                <Card className="p-5">
                    <div className="flex items-center justify-between">
                        <div className="flex items-center gap-3">
                            <div className="h-10 w-10 rounded-lg bg-blue-500/10 flex items-center justify-center">
                                <Users className="h-5 w-5 text-blue-500" />
                            </div>
                            <div>
                                <p className="text-sm text-muted-foreground">{t("workspace.members")}</p>
                                <div className="flex items-center gap-2">
                                    <div className="flex -space-x-2">
                                        {fakeMembers.slice(0, 3).map((member) => (
                                            <Avatar key={member.id} className="h-6 w-6 border-2 border-background">
                                                <AvatarImage src={member.avatar} />
                                                <AvatarFallback className="text-[10px] bg-muted">
                                                    {member.name.split(" ").map(n => n[0]).join("")}
                                                </AvatarFallback>
                                            </Avatar>
                                        ))}
                                    </div>
                                    <span className="text-sm font-medium">{fakeMembers.length}</span>
                                </div>
                            </div>
                        </div>
                        <Button variant="outline" size="sm" className="cursor-pointer">
                            <Plus className="h-3.5 w-3.5 mr-1.5" />
                            {t("workspace.addMember")}
                        </Button>
                    </div>
                </Card>
            </div>

            {/* Create Link Widget - Same style as landing page */}
            <Card className="p-6 bg-card/50 backdrop-blur border-border/50">
                <div className="flex items-center gap-3">
                    <div className="flex-1 relative">
                        <Link2 className="absolute left-4 top-1/2 -translate-y-1/2 h-5 w-5 text-muted-foreground" />
                        <Input
                            placeholder={t("shortener.placeholder")}
                            value={newUrl}
                            onChange={(e) => setNewUrl(e.target.value)}
                            className="h-12 pl-12 pr-4 bg-muted/50 border-border/50 text-base"
                        />
                    </div>
                    <Button
                        onClick={handleCreateLink}
                        disabled={isCreating || !newUrl}
                        className="h-12 px-6 bg-[var(--gold)] hover:bg-[var(--gold)]/90 text-[var(--gold-foreground)] font-medium cursor-pointer"
                    >
                        {isCreating ? (
                            <Loader2 className="h-5 w-5 animate-spin" />
                        ) : (
                            <>
                                {t("shortener.button")}
                                <ArrowRight className="ml-2 h-4 w-4" />
                            </>
                        )}
                    </Button>
                </div>
            </Card>

            {/* Stats Cards */}
            {stats && (
                <div className="grid sm:grid-cols-2 lg:grid-cols-4 gap-4">
                    <Card className="p-5">
                        <div className="flex items-center gap-4">
                            <div className="h-11 w-11 rounded-xl bg-blue-500/10 flex items-center justify-center">
                                <Link2 className="h-5 w-5 text-blue-500" />
                            </div>
                            <div>
                                <p className="text-sm text-muted-foreground">{t("workspace.totalLinks")}</p>
                                <p className="text-2xl font-bold">{stats.totalLinks}</p>
                            </div>
                        </div>
                    </Card>

                    <Card className="p-5">
                        <div className="flex items-center gap-4">
                            <div className="h-11 w-11 rounded-xl bg-green-500/10 flex items-center justify-center">
                                <MousePointerClick className="h-5 w-5 text-green-500" />
                            </div>
                            <div>
                                <p className="text-sm text-muted-foreground">{t("workspace.totalClicks")}</p>
                                <p className="text-2xl font-bold">
                                    {stats.totalClicks.toLocaleString()}
                                </p>
                            </div>
                        </div>
                    </Card>

                    <Card className="p-5">
                        <div className="flex items-center gap-4">
                            <div className="h-11 w-11 rounded-xl bg-purple-500/10 flex items-center justify-center">
                                <BarChart3 className="h-5 w-5 text-purple-500" />
                            </div>
                            <div>
                                <p className="text-sm text-muted-foreground">{t("workspace.activeLinks")}</p>
                                <p className="text-2xl font-bold">{stats.activeLinks}</p>
                            </div>
                        </div>
                    </Card>

                    <Card className="p-5">
                        <div className="flex items-center gap-4">
                            <div className="h-11 w-11 rounded-xl bg-[var(--gold)]/10 flex items-center justify-center">
                                <TrendingUp className="h-5 w-5 text-[var(--gold)]" />
                            </div>
                            <div>
                                <p className="text-sm text-muted-foreground">{t("workspace.thisMonth")}</p>
                                <p className="text-2xl font-bold">{stats.linksThisMonth}</p>
                            </div>
                        </div>
                    </Card>
                </div>
            )}

            {/* Links Table */}
            <Card>
                <div className="p-5 border-b border-border">
                    <h2 className="text-lg font-semibold">{t("workspace.yourLinks")}</h2>
                </div>

                {links.length > 0 ? (
                    <Table>
                        <TableHeader>
                            <TableRow>
                                <TableHead>{t("workspace.shortLink")}</TableHead>
                                <TableHead>{t("workspace.destination")}</TableHead>
                                <TableHead>{t("workspace.clicks")}</TableHead>
                                <TableHead>{t("workspace.status")}</TableHead>
                                <TableHead>{t("workspace.created")}</TableHead>
                                <TableHead className="w-[100px]">{t("common.actions")}</TableHead>
                            </TableRow>
                        </TableHeader>
                        <TableBody>
                            {links.map((link) => (
                                <TableRow key={link.id}>
                                    <TableCell>
                                        <div className="flex items-center gap-2">
                                            <code className="text-sm font-medium">
                                                /{link.shortCode}
                                            </code>
                                            <Button
                                                variant="ghost"
                                                size="icon"
                                                className="h-7 w-7 cursor-pointer"
                                                onClick={() => handleCopy(link)}
                                            >
                                                {copiedId === link.id ? (
                                                    <Check className="h-3.5 w-3.5 text-green-500" />
                                                ) : (
                                                    <Copy className="h-3.5 w-3.5" />
                                                )}
                                            </Button>
                                        </div>
                                    </TableCell>
                                    <TableCell>
                                        <div className="flex items-center gap-2 max-w-xs">
                                            <span className="truncate text-sm text-muted-foreground">
                                                {link.originalUrl}
                                            </span>
                                            <a
                                                href={link.originalUrl}
                                                target="_blank"
                                                rel="noopener noreferrer"
                                                className="shrink-0"
                                            >
                                                <ExternalLink className="h-3 w-3 text-muted-foreground hover:text-foreground" />
                                            </a>
                                        </div>
                                    </TableCell>
                                    <TableCell>
                                        <span className="font-medium">
                                            {link.clicks.toLocaleString()}
                                        </span>
                                    </TableCell>
                                    <TableCell>
                                        <Badge variant={link.isActive ? "default" : "secondary"}>
                                            {link.isActive ? t("workspace.active") : t("workspace.inactive")}
                                        </Badge>
                                    </TableCell>
                                    <TableCell>
                                        <span className="text-sm text-muted-foreground">
                                            {new Date(link.createdAt).toLocaleDateString()}
                                        </span>
                                    </TableCell>
                                    <TableCell>
                                        <DropdownMenu>
                                            <DropdownMenuTrigger asChild>
                                                <Button
                                                    variant="ghost"
                                                    size="icon"
                                                    className="h-7 w-7 cursor-pointer"
                                                >
                                                    <MoreHorizontal className="h-4 w-4" />
                                                </Button>
                                            </DropdownMenuTrigger>
                                            <DropdownMenuContent align="end">
                                                <DropdownMenuItem className="cursor-pointer">
                                                    <Pencil className="mr-2 h-4 w-4" />
                                                    {t("workspace.edit")}
                                                </DropdownMenuItem>
                                                <DropdownMenuItem
                                                    className="cursor-pointer"
                                                    onClick={() => handleToggleStatus(link)}
                                                >
                                                    <Power className="mr-2 h-4 w-4" />
                                                    {link.isActive ? t("workspace.deactivate") : t("workspace.activate")}
                                                </DropdownMenuItem>
                                                <DropdownMenuSeparator />
                                                <DropdownMenuItem
                                                    className="text-destructive cursor-pointer"
                                                    onClick={() => handleDelete(link.id)}
                                                >
                                                    <Trash2 className="mr-2 h-4 w-4" />
                                                    {t("workspace.delete")}
                                                </DropdownMenuItem>
                                            </DropdownMenuContent>
                                        </DropdownMenu>
                                    </TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                ) : (
                    <div className="p-12 text-center">
                        <Link2 className="h-12 w-12 mx-auto mb-4 text-muted-foreground" />
                        <h3 className="text-lg font-semibold mb-2">{t("workspace.noLinks")}</h3>
                        <p className="text-muted-foreground mb-6">
                            {t("workspace.noLinksDesc")}
                        </p>
                        <Button
                            onClick={() => setIsCreateOpen(true)}
                            className="bg-[var(--gold)] hover:bg-[var(--gold)]/90 text-[var(--gold-foreground)] cursor-pointer"
                        >
                            <Plus className="mr-2 h-4 w-4" />
                            {t("workspace.createLink")}
                        </Button>
                    </div>
                )}
            </Card>
        </div>
    );
}
