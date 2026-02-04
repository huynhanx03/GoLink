"use client";

import { useState } from "react";
import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Badge } from "@/components/ui/badge";
import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { fakeTenants } from "@/lib/data/fake-data";
import {
    Search,
    MoreHorizontal,
    Pencil,
    Trash2,
    Building2,
    Users,
    Link2,
    CreditCard,
} from "lucide-react";

export default function AdminTenantsPage() {
    const [searchQuery, setSearchQuery] = useState("");

    const filteredTenants = fakeTenants.filter(
        (tenant) =>
            tenant.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
            tenant.slug.toLowerCase().includes(searchQuery.toLowerCase())
    );

    return (
        <div className="max-w-6xl">
            {/* Header */}
            <div className="flex items-center justify-between mb-8">
                <div>
                    <h1 className="text-3xl font-bold text-white mb-2">Tenants</h1>
                    <p className="text-zinc-400">
                        Manage workspaces and their subscriptions
                    </p>
                </div>
            </div>

            {/* Search */}
            <Card className="p-4 bg-zinc-900 border-zinc-800 mb-6">
                <div className="relative">
                    <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-zinc-500" />
                    <Input
                        placeholder="Search tenants by name or slug..."
                        value={searchQuery}
                        onChange={(e) => setSearchQuery(e.target.value)}
                        className="pl-10 bg-zinc-800 border-zinc-700 text-white placeholder:text-zinc-500"
                    />
                </div>
            </Card>

            {/* Tenants Table */}
            <Card className="bg-zinc-900 border-zinc-800">
                <Table>
                    <TableHeader>
                        <TableRow className="border-zinc-800 hover:bg-transparent">
                            <TableHead className="text-zinc-400">Tenant</TableHead>
                            <TableHead className="text-zinc-400">Plan</TableHead>
                            <TableHead className="text-zinc-400">Members</TableHead>
                            <TableHead className="text-zinc-400">Links</TableHead>
                            <TableHead className="text-zinc-400">Created</TableHead>
                            <TableHead className="text-zinc-400 w-[100px]">Actions</TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {filteredTenants.map((tenant) => (
                            <TableRow
                                key={tenant.id}
                                className="border-zinc-800 hover:bg-zinc-800/50"
                            >
                                <TableCell>
                                    <div className="flex items-center gap-3">
                                        <div className="h-10 w-10 rounded-lg bg-zinc-800 flex items-center justify-center">
                                            <Building2 className="h-5 w-5 text-zinc-400" />
                                        </div>
                                        <div>
                                            <div className="font-medium text-white">{tenant.name}</div>
                                            <div className="text-sm text-zinc-400">/{tenant.slug}</div>
                                        </div>
                                    </div>
                                </TableCell>
                                <TableCell>
                                    <Badge
                                        className={
                                            tenant.plan.tier === "enterprise"
                                                ? "bg-purple-500/20 text-purple-400"
                                                : tenant.plan.tier === "pro"
                                                    ? "bg-[var(--gold)]/20 text-[var(--gold)]"
                                                    : "bg-zinc-700 text-zinc-300"
                                        }
                                    >
                                        {tenant.plan.name}
                                    </Badge>
                                </TableCell>
                                <TableCell>
                                    <div className="flex items-center gap-2 text-white">
                                        <Users className="h-4 w-4 text-zinc-500" />
                                        {tenant.memberCount}
                                    </div>
                                </TableCell>
                                <TableCell>
                                    <div className="flex items-center gap-2 text-white">
                                        <Link2 className="h-4 w-4 text-zinc-500" />
                                        {tenant.linkCount}
                                    </div>
                                </TableCell>
                                <TableCell>
                                    <span className="text-zinc-400">
                                        {new Date(tenant.createdAt).toLocaleDateString()}
                                    </span>
                                </TableCell>
                                <TableCell>
                                    <DropdownMenu>
                                        <DropdownMenuTrigger asChild>
                                            <Button
                                                variant="ghost"
                                                size="icon"
                                                className="h-8 w-8 text-zinc-400 hover:text-white hover:bg-zinc-800 cursor-pointer"
                                            >
                                                <MoreHorizontal className="h-4 w-4" />
                                            </Button>
                                        </DropdownMenuTrigger>
                                        <DropdownMenuContent
                                            align="end"
                                            className="bg-zinc-900 border-zinc-800"
                                        >
                                            <DropdownMenuItem className="text-zinc-300 hover:bg-zinc-800 cursor-pointer">
                                                <Pencil className="mr-2 h-4 w-4" />
                                                Edit
                                            </DropdownMenuItem>
                                            <DropdownMenuItem className="text-zinc-300 hover:bg-zinc-800 cursor-pointer">
                                                <CreditCard className="mr-2 h-4 w-4" />
                                                Change Plan
                                            </DropdownMenuItem>
                                            <DropdownMenuSeparator className="bg-zinc-800" />
                                            <DropdownMenuItem className="text-red-400 hover:bg-zinc-800 cursor-pointer">
                                                <Trash2 className="mr-2 h-4 w-4" />
                                                Delete
                                            </DropdownMenuItem>
                                        </DropdownMenuContent>
                                    </DropdownMenu>
                                </TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </Card>
        </div>
    );
}
