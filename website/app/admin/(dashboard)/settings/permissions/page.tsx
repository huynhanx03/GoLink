"use client";

import { useState } from "react";
import { Role, Resource, Permission } from "@/lib/types";
import { fakeRoles, fakeResources, fakePermissions } from "@/lib/data/fake-data";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/components/ui/card";
import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog";
import { Label } from "@/components/ui/label";
import { Checkbox } from "@/components/ui/checkbox";
import { ScrollArea } from "@/components/ui/scroll-area";
import { toast } from "sonner";
import { Plus, Pencil, Trash2, Loader2, Save, ChevronRight } from "lucide-react";

// Scope constants
const SCOPE_READ = 1;
const SCOPE_CREATE = 2;
const SCOPE_UPDATE = 4;
const SCOPE_DELETE = 8;

const SCOPES = [
    { value: SCOPE_READ, label: "Xem", key: "view" },
    { value: SCOPE_CREATE, label: "Thêm", key: "add" },
    { value: SCOPE_UPDATE, label: "Sửa", key: "edit" },
    { value: SCOPE_DELETE, label: "Xóa", key: "delete" },
];

export default function PermissionsPage() {
    const [roles, setRoles] = useState<Role[]>(fakeRoles);
    const [resources] = useState<Resource[]>(fakeResources);
    const [permissions, setPermissions] = useState<Permission[]>(fakePermissions);
    const [selectedRoleId, setSelectedRoleId] = useState<string | null>(roles[0]?.id || null);
    const [isSaving, setIsSaving] = useState(false);

    // Role dialog state
    const [isRoleDialogOpen, setIsRoleDialogOpen] = useState(false);
    const [editingRole, setEditingRole] = useState<Role | null>(null);
    const [roleFormData, setRoleFormData] = useState({ name: "", level: 0 });

    const selectedRole = roles.find(r => r.id === selectedRoleId);

    const hasScope = (roleId: string, resourceId: string, scope: number) => {
        const perm = permissions.find(p => p.roleId === roleId && p.resourceId === resourceId);
        if (!perm) return false;
        return (perm.scopes & scope) === scope;
    };

    const toggleScope = (roleId: string, resourceId: string, scope: number) => {
        const existingPerm = permissions.find(p => p.roleId === roleId && p.resourceId === resourceId);

        if (existingPerm) {
            const newScopes = (existingPerm.scopes & scope) === scope
                ? existingPerm.scopes & ~scope
                : existingPerm.scopes | scope;

            setPermissions(permissions.map(p =>
                p.id === existingPerm.id ? { ...p, scopes: newScopes, updatedAt: new Date() } : p
            ));
        } else {
            const newPerm: Permission = {
                id: `perm-${Date.now()}`,
                roleId,
                resourceId,
                scopes: scope,
                createdAt: new Date(),
                updatedAt: new Date(),
            };
            setPermissions([...permissions, newPerm]);
        }
    };

    const getPermissionCount = (roleId: string, resourceId: string) => {
        const perm = permissions.find(p => p.roleId === roleId && p.resourceId === resourceId);
        if (!perm) return 0;
        let count = 0;
        SCOPES.forEach(s => {
            if ((perm.scopes & s.value) === s.value) count++;
        });
        return count;
    };

    const handleSave = async () => {
        setIsSaving(true);
        await new Promise(resolve => setTimeout(resolve, 800));
        setIsSaving(false);
        toast.success("Đã lưu thay đổi thành công");
    };

    // Role CRUD
    const handleRoleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        if (editingRole) {
            setRoles(roles.map(r =>
                r.id === editingRole.id
                    ? { ...r, name: roleFormData.name, level: roleFormData.level, updatedAt: new Date() }
                    : r
            ));
            toast.success("Đã cập nhật vai trò");
        } else {
            const newRole: Role = {
                id: `role-${Date.now()}`,
                name: roleFormData.name,
                level: roleFormData.level,
                lft: 0,
                rgt: 0,
                createdAt: new Date(),
                updatedAt: new Date(),
            };
            setRoles([...roles, newRole]);
            toast.success("Đã tạo vai trò mới");
        }
        setIsRoleDialogOpen(false);
        resetRoleForm();
    };

    const resetRoleForm = () => {
        setRoleFormData({ name: "", level: 0 });
        setEditingRole(null);
    };

    const handleEditRole = (role: Role) => {
        setEditingRole(role);
        setRoleFormData({ name: role.name, level: role.level });
        setIsRoleDialogOpen(true);
    };

    const handleDeleteRole = (roleId: string) => {
        if (confirm("Bạn có chắc muốn xóa vai trò này?")) {
            setRoles(roles.filter(r => r.id !== roleId));
            setPermissions(permissions.filter(p => p.roleId !== roleId));
            if (selectedRoleId === roleId) {
                setSelectedRoleId(roles[0]?.id || null);
            }
            toast.success("Đã xóa vai trò");
        }
    };

    return (
        <div className="flex-1 space-y-4 p-8 pt-6">
            <div className="flex items-center justify-between">
                <div>
                    <h2 className="text-3xl font-bold tracking-tight">Quản lý phân quyền</h2>
                    <p className="text-muted-foreground">
                        Thiết lập quyền truy cập cho từng vai trò trong hệ thống.
                    </p>
                </div>
                <div className="flex items-center gap-2">
                    <Dialog open={isRoleDialogOpen} onOpenChange={(open) => {
                        setIsRoleDialogOpen(open);
                        if (!open) resetRoleForm();
                    }}>
                        <DialogTrigger asChild>
                            <Button variant="outline">
                                <Plus className="mr-2 h-4 w-4" /> Thêm vai trò
                            </Button>
                        </DialogTrigger>
                        <DialogContent>
                            <DialogHeader>
                                <DialogTitle>{editingRole ? "Sửa vai trò" : "Thêm vai trò mới"}</DialogTitle>
                                <DialogDescription>
                                    Nhập thông tin cho vai trò.
                                </DialogDescription>
                            </DialogHeader>
                            <form onSubmit={handleRoleSubmit} className="space-y-4">
                                <div className="space-y-2">
                                    <Label htmlFor="name">Tên vai trò</Label>
                                    <Input
                                        id="name"
                                        placeholder="VD: Quản lý"
                                        value={roleFormData.name}
                                        onChange={(e) => setRoleFormData({ ...roleFormData, name: e.target.value })}
                                        required
                                    />
                                </div>
                                <div className="space-y-2">
                                    <Label htmlFor="level">Mức độ ưu tiên</Label>
                                    <Input
                                        id="level"
                                        type="number"
                                        placeholder="VD: 50"
                                        value={roleFormData.level}
                                        onChange={(e) => setRoleFormData({ ...roleFormData, level: Number(e.target.value) })}
                                        required
                                    />
                                    <p className="text-[0.8rem] text-muted-foreground">
                                        Số càng cao càng có quyền lớn hơn.
                                    </p>
                                </div>
                                <DialogFooter>
                                    <Button type="button" variant="outline" onClick={() => setIsRoleDialogOpen(false)}>
                                        Hủy
                                    </Button>
                                    <Button type="submit">
                                        {editingRole ? "Lưu" : "Tạo vai trò"}
                                    </Button>
                                </DialogFooter>
                            </form>
                        </DialogContent>
                    </Dialog>
                    <Button onClick={handleSave} disabled={isSaving}>
                        {isSaving && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                        {!isSaving && <Save className="mr-2 h-4 w-4" />}
                        Lưu thay đổi
                    </Button>
                </div>
            </div>

            <div className="grid grid-cols-12 gap-6">
                {/* Roles List - Left Side */}
                <div className="col-span-4">
                    <Card className="h-full">
                        <CardHeader>
                            <CardTitle>Vai trò</CardTitle>
                            <CardDescription>Chọn vai trò để thiết lập quyền</CardDescription>
                        </CardHeader>
                        <CardContent className="p-0">
                            <ScrollArea className="h-[500px]">
                                <div className="space-y-1 p-4">
                                    {roles.map((role) => (
                                        <div
                                            key={role.id}
                                            className={`flex items-center justify-between p-3 rounded-lg cursor-pointer transition-colors ${selectedRoleId === role.id
                                                    ? "bg-[var(--gold)] text-[var(--gold-foreground)]"
                                                    : "hover:bg-muted"
                                                }`}
                                            onClick={() => setSelectedRoleId(role.id)}
                                        >
                                            <div className="flex items-center gap-2">
                                                <ChevronRight className={`h-4 w-4 transition-transform ${selectedRoleId === role.id ? "rotate-90" : ""}`} />
                                                <div>
                                                    <div className="font-medium">{role.name}</div>
                                                    <div className={`text-xs ${selectedRoleId === role.id ? "opacity-80" : "text-muted-foreground"}`}>
                                                        Level: {role.level}
                                                    </div>
                                                </div>
                                            </div>
                                            <div className="flex items-center gap-1" onClick={(e) => e.stopPropagation()}>
                                                <Button
                                                    variant="ghost"
                                                    size="sm"
                                                    className={selectedRoleId === role.id ? "text-[var(--gold-foreground)] hover:bg-[var(--gold-foreground)]/20" : ""}
                                                    onClick={() => handleEditRole(role)}
                                                >
                                                    Sửa
                                                </Button>
                                                <Button
                                                    variant="ghost"
                                                    size="sm"
                                                    className={`${selectedRoleId === role.id ? "text-[var(--gold-foreground)] hover:bg-[var(--gold-foreground)]/20" : "text-red-500 hover:text-red-600"}`}
                                                    onClick={() => handleDeleteRole(role.id)}
                                                >
                                                    Xóa
                                                </Button>
                                            </div>
                                        </div>
                                    ))}
                                </div>
                            </ScrollArea>
                        </CardContent>
                    </Card>
                </div>

                {/* Permissions - Right Side */}
                <div className="col-span-8">
                    <Card className="h-full">
                        <CardHeader>
                            <CardTitle>
                                Quyền cho: {selectedRole?.name || "Chưa chọn vai trò"}
                            </CardTitle>
                            <CardDescription>
                                Đánh dấu vào các quyền thao tác tương ứng với vai trò
                            </CardDescription>
                        </CardHeader>
                        <CardContent>
                            {selectedRoleId ? (
                                <ScrollArea className="h-[450px] pr-4">
                                    <div className="space-y-4">
                                        {resources.map((resource) => (
                                            <div key={resource.id} className="border rounded-lg p-4">
                                                <div className="flex items-center justify-between mb-3">
                                                    <div className="flex items-center gap-2">
                                                        <div className="w-2 h-2 rounded-full bg-[var(--gold)]" />
                                                        <span className="font-medium">{resource.key}</span>
                                                        {resource.description && (
                                                            <span className="text-sm text-muted-foreground">
                                                                - {resource.description}
                                                            </span>
                                                        )}
                                                    </div>
                                                    <span className="text-xs text-muted-foreground">
                                                        {getPermissionCount(selectedRoleId, resource.id)}/{SCOPES.length}
                                                    </span>
                                                </div>
                                                <div className="grid grid-cols-2 md:grid-cols-4 gap-3">
                                                    {SCOPES.map((scope) => (
                                                        <label
                                                            key={scope.value}
                                                            className="flex items-center gap-2 p-2 rounded-md hover:bg-muted cursor-pointer transition-colors"
                                                        >
                                                            <Checkbox
                                                                checked={hasScope(selectedRoleId, resource.id, scope.value)}
                                                                onCheckedChange={() => toggleScope(selectedRoleId, resource.id, scope.value)}
                                                            />
                                                            <div>
                                                                <div className="text-sm font-medium">{scope.label}</div>
                                                                <div className="text-xs text-muted-foreground">
                                                                    {resource.key}.{scope.key}
                                                                </div>
                                                            </div>
                                                        </label>
                                                    ))}
                                                </div>
                                            </div>
                                        ))}
                                    </div>
                                </ScrollArea>
                            ) : (
                                <div className="flex items-center justify-center h-[450px] text-muted-foreground">
                                    Vui lòng chọn vai trò từ danh sách bên trái
                                </div>
                            )}
                        </CardContent>
                    </Card>
                </div>
            </div>
        </div>
    );
}
