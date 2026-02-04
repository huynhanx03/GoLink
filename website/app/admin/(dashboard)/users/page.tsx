"use client";

import { useState } from "react";
import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Badge } from "@/components/ui/badge";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
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
import { fakeUsers } from "@/lib/data/fake-data";
import { Search, MoreHorizontal, Pencil, Ban, Trash2, UserPlus } from "lucide-react";

export default function AdminUsersPage() {
    const [searchQuery, setSearchQuery] = useState("");

    const filteredUsers = fakeUsers.filter(
        (user) =>
            user.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
            user.email.toLowerCase().includes(searchQuery.toLowerCase())
    );

    return (
        <div className="max-w-6xl">
            {/* Header */}
            <div className="flex items-center justify-between mb-8">
                <div>
                    <h1 className="text-3xl font-bold text-white mb-2">Users</h1>
                    <p className="text-zinc-400">Manage platform users and their access</p>
                </div>
                <Button className="bg-[var(--gold)] hover:bg-[var(--gold)]/90 text-[var(--gold-foreground)] cursor-pointer">
                    <UserPlus className="mr-2 h-4 w-4" />
                    Add User
                </Button>
            </div>

            {/* Search */}
            <Card className="p-4 bg-zinc-900 border-zinc-800 mb-6">
                <div className="relative">
                    <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-zinc-500" />
                    <Input
                        placeholder="Search users by name or email..."
                        value={searchQuery}
                        onChange={(e) => setSearchQuery(e.target.value)}
                        className="pl-10 bg-zinc-800 border-zinc-700 text-white placeholder:text-zinc-500"
                    />
                </div>
            </Card>

            {/* Users Table */}
            <Card className="bg-zinc-900 border-zinc-800">
                <Table>
                    <TableHeader>
                        <TableRow className="border-zinc-800 hover:bg-transparent">
                            <TableHead className="text-zinc-400">User</TableHead>
                            <TableHead className="text-zinc-400">Role</TableHead>
                            <TableHead className="text-zinc-400">Workspaces</TableHead>
                            <TableHead className="text-zinc-400">Joined</TableHead>
                            <TableHead className="text-zinc-400 w-[100px]">Actions</TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {filteredUsers.map((user) => (
                            <TableRow key={user.id} className="border-zinc-800 hover:bg-zinc-800/50">
                                <TableCell>
                                    <div className="flex items-center gap-3">
                                        <Avatar className="h-10 w-10">
                                            <AvatarImage src={user.avatar} alt={user.name} />
                                            <AvatarFallback className="bg-zinc-700 text-zinc-300">
                                                {user.name.split(" ").map((n) => n[0]).join("")}
                                            </AvatarFallback>
                                        </Avatar>
                                        <div>
                                            <div className="font-medium text-white">{user.name}</div>
                                            <div className="text-sm text-zinc-400">{user.email}</div>
                                        </div>
                                    </div>
                                </TableCell>
                                <TableCell>
                                    <Badge
                                        variant={user.role === "admin" ? "default" : "secondary"}
                                        className={
                                            user.role === "admin"
                                                ? "bg-[var(--gold)] text-[var(--gold-foreground)]"
                                                : "bg-zinc-700 text-zinc-300"
                                        }
                                    >
                                        {user.role}
                                    </Badge>
                                </TableCell>
                                <TableCell>
                                    <span className="text-white">{user.tenants.length}</span>
                                </TableCell>
                                <TableCell>
                                    <span className="text-zinc-400">
                                        {new Date(user.createdAt).toLocaleDateString()}
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
                                                <Ban className="mr-2 h-4 w-4" />
                                                Suspend
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
