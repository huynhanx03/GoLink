"use client";

import { useState } from "react";
import { Role } from "@/lib/types";
import { fakeRoles } from "@/lib/data/fake-data";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table";
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
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select";
import { Label } from "@/components/ui/label";
import { Plus, Trash2, Pencil, Search, ShieldCheck } from "lucide-react";
import { toast } from "sonner";

export default function RolesPage() {
    const [roles, setRoles] = useState<Role[]>(fakeRoles);
    const [searchQuery, setSearchQuery] = useState("");
    const [isOpen, setIsOpen] = useState(false);
    const [editingRole, setEditingRole] = useState<Role | null>(null);

    const [formData, setFormData] = useState({
        name: "",
        level: 0,
        parentId: "none",
    });

    const filteredRoles = roles.filter(
        (role) =>
            role.name.toLowerCase().includes(searchQuery.toLowerCase())
    );

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();

        const parentId = formData.parentId === "none" ? undefined : formData.parentId;

        if (editingRole) {
            // Update existing
            const updated = roles.map(role =>
                role.id === editingRole.id
                    ? { ...role, name: formData.name, level: Number(formData.level), parentId, updatedAt: new Date() }
                    : role
            );
            setRoles(updated);
            toast.success("Role updated successfully");
        } else {
            // Create new
            const newRole: Role = {
                id: `role-${Date.now()}`,
                name: formData.name,
                level: Number(formData.level),
                parentId,
                lft: 0, // Mock value
                rgt: 0, // Mock value
                createdAt: new Date(),
                updatedAt: new Date(),
            };
            setRoles([...roles, newRole]);
            toast.success("Role created successfully");
        }

        setIsOpen(false);
        resetForm();
    };

    const resetForm = () => {
        setFormData({ name: "", level: 0, parentId: "none" });
        setEditingRole(null);
    };

    const handleEdit = (role: Role) => {
        setEditingRole(role);
        setFormData({
            name: role.name,
            level: role.level,
            parentId: role.parentId || "none",
        });
        setIsOpen(true);
    };

    const handleDelete = (id: string) => {
        if (confirm("Are you sure you want to delete this role?")) {
            setRoles(roles.filter(role => role.id !== id));
            toast.success("Role deleted successfully");
        }
    };

    return (
        <div className="flex-1 space-y-4 p-8 pt-6">
            <div className="flex items-center justify-between space-y-2">
                <div>
                    <h2 className="text-3xl font-bold tracking-tight">Roles</h2>
                    <p className="text-muted-foreground">
                        Manage system roles and hierarchy.
                    </p>
                </div>
                <div className="flex items-center space-x-2">
                    <Dialog open={isOpen} onOpenChange={(open) => {
                        setIsOpen(open);
                        if (!open) resetForm();
                    }}>
                        <DialogTrigger asChild>
                            <Button>
                                <Plus className="mr-2 h-4 w-4" /> Add Role
                            </Button>
                        </DialogTrigger>
                        <DialogContent>
                            <DialogHeader>
                                <DialogTitle>{editingRole ? "Edit Role" : "Add Role"}</DialogTitle>
                                <DialogDescription>
                                    Define a role and its hierarchy level.
                                </DialogDescription>
                            </DialogHeader>
                            <form onSubmit={handleSubmit} className="space-y-4">
                                <div className="space-y-2">
                                    <Label htmlFor="name">Role Name</Label>
                                    <Input
                                        id="name"
                                        placeholder="e.g. Manager"
                                        value={formData.name}
                                        onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                                        required
                                    />
                                </div>
                                <div className="space-y-2">
                                    <Label htmlFor="level">Level</Label>
                                    <Input
                                        id="level"
                                        type="number"
                                        placeholder="e.g. 50"
                                        value={formData.level}
                                        onChange={(e) => setFormData({ ...formData, level: Number(e.target.value) })}
                                        required
                                    />
                                    <p className="text-[0.8rem] text-muted-foreground">
                                        Higher level indicates higher privilege.
                                    </p>
                                </div>
                                <div className="space-y-2">
                                    <Label htmlFor="parent">Parent Role</Label>
                                    <Select
                                        value={formData.parentId}
                                        onValueChange={(value) => setFormData({ ...formData, parentId: value })}
                                    >
                                        <SelectTrigger>
                                            <SelectValue placeholder="Select parent role" />
                                        </SelectTrigger>
                                        <SelectContent>
                                            <SelectItem value="none">None (Top Level)</SelectItem>
                                            {roles
                                                .filter(r => r.id !== editingRole?.id) // Prevent self-parenting loop
                                                .map((role) => (
                                                    <SelectItem key={role.id} value={role.id}>
                                                        {role.name}
                                                    </SelectItem>
                                                ))}
                                        </SelectContent>
                                    </Select>
                                </div>
                                <DialogFooter>
                                    <Button type="button" variant="outline" onClick={() => setIsOpen(false)}>
                                        Cancel
                                    </Button>
                                    <Button type="submit">
                                        {editingRole ? "Save Changes" : "Create Role"}
                                    </Button>
                                </DialogFooter>
                            </form>
                        </DialogContent>
                    </Dialog>
                </div>
            </div>

            <div className="flex items-center gap-2">
                <div className="relative flex-1 max-w-sm">
                    <Search className="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
                    <Input
                        type="search"
                        placeholder="Search roles..."
                        className="pl-8"
                        value={searchQuery}
                        onChange={(e) => setSearchQuery(e.target.value)}
                    />
                </div>
            </div>

            <div className="rounded-md border">
                <Table>
                    <TableHeader>
                        <TableRow>
                            <TableHead>Role Name</TableHead>
                            <TableHead>Level</TableHead>
                            <TableHead>Parent Role</TableHead>
                            <TableHead className="w-[100px]">Actions</TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {filteredRoles.length === 0 ? (
                            <TableRow>
                                <TableCell colSpan={4} className="h-24 text-center">
                                    No roles found.
                                </TableCell>
                            </TableRow>
                        ) : (
                            filteredRoles.map((role) => {
                                const parent = roles.find(r => r.id === role.parentId);
                                return (
                                    <TableRow key={role.id}>
                                        <TableCell className="font-medium">
                                            <div className="flex items-center gap-2">
                                                <ShieldCheck className="h-4 w-4 text-muted-foreground" />
                                                {role.name}
                                            </div>
                                        </TableCell>
                                        <TableCell>{role.level}</TableCell>
                                        <TableCell>
                                            {parent ? (
                                                <span className="inline-flex items-center rounded-md bg-zinc-100 px-2 py-1 text-xs font-medium text-zinc-600 ring-1 ring-inset ring-zinc-500/10 dark:bg-zinc-400/10 dark:text-zinc-400 dark:ring-zinc-400/20">
                                                    {parent.name}
                                                </span>
                                            ) : (
                                                <span className="text-muted-foreground">-</span>
                                            )}
                                        </TableCell>
                                        <TableCell>
                                            <div className="flex items-center gap-2">
                                                <Button
                                                    variant="ghost"
                                                    size="icon"
                                                    onClick={() => handleEdit(role)}
                                                >
                                                    <Pencil className="h-4 w-4" />
                                                </Button>
                                                <Button
                                                    variant="ghost"
                                                    size="icon"
                                                    className="text-red-500 hover:text-red-600"
                                                    onClick={() => handleDelete(role.id)}
                                                >
                                                    <Trash2 className="h-4 w-4" />
                                                </Button>
                                            </div>
                                        </TableCell>
                                    </TableRow>
                                );
                            })
                        )}
                    </TableBody>
                </Table>
            </div>
        </div>
    );
}
