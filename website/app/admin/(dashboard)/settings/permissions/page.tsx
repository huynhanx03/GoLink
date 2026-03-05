"use client";

import { useEffect, useState, useCallback } from "react";
import { identityService } from "@/lib/api";
import { RoleDto, ResourceDto, PermissionDto } from "@/lib/api/types";
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
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select";
import { Label } from "@/components/ui/label";
import { Checkbox } from "@/components/ui/checkbox";
import { ScrollArea } from "@/components/ui/scroll-area";
import { toast } from "sonner";
import { Plus, Loader2, Save, ChevronRight, Trash2 } from "lucide-react";

// Bitmask scope constants matching backend
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

// Track local changes before saving
interface PendingChange {
    type: "create" | "update" | "delete";
    roleId: number;
    resourceId: number;
    scopes: number;
    existingPermId?: number;
}

export default function PermissionsSettingsPage() {
    const [roles, setRoles] = useState<RoleDto[]>([]);
    const [allResources, setAllResources] = useState<ResourceDto[]>([]);
    const [permissions, setPermissions] = useState<PermissionDto[]>([]);
    const [selectedRoleId, setSelectedRoleId] = useState<number | null>(null);
    const [isLoading, setIsLoading] = useState(true);
    const [isSaving, setIsSaving] = useState(false);
    const [pendingChanges, setPendingChanges] = useState<Map<string, PendingChange>>(new Map());

    // Role dialog state
    const [isRoleDialogOpen, setIsRoleDialogOpen] = useState(false);
    const [roleFormData, setRoleFormData] = useState({ name: "", level: 0 });
    const [isCreatingRole, setIsCreatingRole] = useState(false);

    // Add resource dialog state
    const [isAddResourceDialogOpen, setIsAddResourceDialogOpen] = useState(false);
    const [selectedNewResourceId, setSelectedNewResourceId] = useState<string>("");

    const changeKey = (roleId: number, resourceId: number) => `${roleId}-${resourceId}`;

    const selectedRole = roles.find(r => r.id === selectedRoleId);

    // Load roles and resources once on mount
    const loadInitialData = useCallback(async () => {
        setIsLoading(true);
        const [rolesRes, resourcesRes] = await Promise.all([
            identityService.findRoles(),
            identityService.findResources(),
        ]);

        if (rolesRes.success && rolesRes.data?.records) {
            setRoles(rolesRes.data.records);
            if (rolesRes.data.records.length > 0 && !selectedRoleId) {
                setSelectedRoleId(rolesRes.data.records[0].id);
            }
        }
        if (resourcesRes.success && resourcesRes.data?.records) {
            setAllResources(resourcesRes.data.records);
        }

        setIsLoading(false);
    }, [selectedRoleId]);

    // Load permissions for selected role
    const loadRolePermissions = useCallback(async (roleId: number) => {
        const permsRes = await identityService.findPermissions({
            pagination: { page: 1, page_size: 1000 },
            filters: [{ key: "role_id", value: roleId, type: "exact" }],
        });

        if (permsRes.success && permsRes.data?.records) {
            setPermissions(permsRes.data.records);
        }
    }, []);

    useEffect(() => {
        loadInitialData();
    }, []); // eslint-disable-line react-hooks/exhaustive-deps

    // Reload permissions when selected role changes
    useEffect(() => {
        if (selectedRoleId) {
            setPendingChanges(new Map());
            loadRolePermissions(selectedRoleId);
        }
    }, [selectedRoleId, loadRolePermissions]);

    // Resources that the selected role has permissions for (original + pending creates)
    const getRoleResources = (): ResourceDto[] => {
        if (!selectedRoleId) return [];

        // Resource IDs from existing permissions
        const existingResourceIds = new Set(
            permissions
                .filter(p => p.role_id === selectedRoleId)
                .map(p => p.resource_id)
        );

        // Resource IDs from pending creates
        pendingChanges.forEach((change) => {
            if (change.roleId === selectedRoleId) {
                if (change.type === "create" || change.type === "update") {
                    existingResourceIds.add(change.resourceId);
                }
                if (change.type === "delete") {
                    existingResourceIds.delete(change.resourceId);
                }
            }
        });

        return allResources.filter(r => existingResourceIds.has(r.id));
    };

    // Resources not yet assigned to the selected role (for adding)
    const getAvailableResources = (): ResourceDto[] => {
        const assignedIds = new Set(getRoleResources().map(r => r.id));
        return allResources.filter(r => !assignedIds.has(r.id));
    };

    // Get combined scopes from all permission rows for a role-resource pair (backend stores 1 row per scope)
    const getOriginalScopes = (roleId: number, resourceId: number): number => {
        return permissions
            .filter(p => p.role_id === roleId && p.resource_id === resourceId)
            .reduce((acc, p) => acc | p.scopes, 0);
    };

    // Get permission IDs for a role-resource pair
    const getPermissionIds = (roleId: number, resourceId: number): number[] => {
        return permissions
            .filter(p => p.role_id === roleId && p.resource_id === resourceId)
            .map(p => p.id);
    };

    // Get effective scopes (pending change or original combined)
    const getScopes = (roleId: number, resourceId: number): number => {
        const key = changeKey(roleId, resourceId);
        const pending = pendingChanges.get(key);
        if (pending) {
            if (pending.type === "delete") return 0;
            return pending.scopes;
        }
        return getOriginalScopes(roleId, resourceId);
    };

    const hasScope = (roleId: number, resourceId: number, scope: number) => {
        return (getScopes(roleId, resourceId) & scope) === scope;
    };

    const toggleScope = (roleId: number, resourceId: number, scope: number) => {
        const key = changeKey(roleId, resourceId);
        const currentScopes = getScopes(roleId, resourceId);
        const newScopes = (currentScopes & scope) === scope
            ? currentScopes & ~scope
            : currentScopes | scope;

        const originalScopes = getOriginalScopes(roleId, resourceId);
        const hasExisting = permissions.some(p => p.role_id === roleId && p.resource_id === resourceId);

        setPendingChanges(prev => {
            const next = new Map(prev);
            // If back to original, remove pending change
            if (originalScopes === newScopes) {
                next.delete(key);
            } else if (!hasExisting && newScopes === 0) {
                next.delete(key);
            } else {
                next.set(key, {
                    type: hasExisting ? "update" : "create",
                    roleId,
                    resourceId,
                    scopes: newScopes,
                    existingPermId: undefined, // handled in save via getPermissionIds
                });
            }
            return next;
        });
    };

    const getPermissionCount = (roleId: number, resourceId: number) => {
        const scopes = getScopes(roleId, resourceId);
        let count = 0;
        SCOPES.forEach(s => {
            if ((scopes & s.value) === s.value) count++;
        });
        return count;
    };

    // Add a new resource to the selected role (pending create with 0 scopes initially)
    const handleAddResource = () => {
        if (!selectedRoleId || !selectedNewResourceId) return;
        const resourceId = Number(selectedNewResourceId);
        const key = changeKey(selectedRoleId, resourceId);

        setPendingChanges(prev => {
            const next = new Map(prev);
            next.set(key, {
                type: "create",
                roleId: selectedRoleId,
                resourceId,
                scopes: SCOPE_READ, // Default: read permission
            });
            return next;
        });

        setSelectedNewResourceId("");
        setIsAddResourceDialogOpen(false);
        toast.success("Đã thêm resource, nhớ bấm Lưu");
    };

    // Remove all permission rows for a resource from the selected role
    const handleRemoveResource = (resourceId: number) => {
        if (!selectedRoleId) return;
        const key = changeKey(selectedRoleId, resourceId);
        const hasExisting = permissions.some(p => p.role_id === selectedRoleId && p.resource_id === resourceId);

        setPendingChanges(prev => {
            const next = new Map(prev);
            if (hasExisting) {
                next.set(key, {
                    type: "delete",
                    roleId: selectedRoleId,
                    resourceId,
                    scopes: 0,
                });
            } else {
                // Was a pending create, just remove it
                next.delete(key);
            }
            return next;
        });
    };

    const handleSave = async () => {
        if (pendingChanges.size === 0) {
            toast.info("Không có thay đổi");
            return;
        }

        setIsSaving(true);
        let errorCount = 0;

        for (const change of pendingChanges.values()) {
            // Step 1: Delete all existing rows for this role-resource pair
            const existingIds = getPermissionIds(change.roleId, change.resourceId);
            for (const id of existingIds) {
                const result = await identityService.deletePermission(id);
                if (!result.success) errorCount++;
            }

            // Step 2: Create new rows for each scope bit (skip if delete or scopes=0)
            if (change.type !== "delete" && change.scopes > 0) {
                for (const scope of SCOPES) {
                    if ((change.scopes & scope.value) === scope.value) {
                        const result = await identityService.createPermission({
                            role_id: change.roleId,
                            resource_id: change.resourceId,
                            scopes: scope.value,
                        });
                        if (!result.success) errorCount++;
                    }
                }
            }
        }

        if (errorCount > 0) toast.error(`Có ${errorCount} lỗi khi lưu`);
        else toast.success("Đã lưu thành công");

        setPendingChanges(new Map());
        if (selectedRoleId) await loadRolePermissions(selectedRoleId);
        setIsSaving(false);
    };

    // Create new role via API
    const handleRoleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setIsCreatingRole(true);

        const result = await identityService.createRole({
            name: roleFormData.name,
            level: roleFormData.level,
        });

        if (result.success && result.data) {
            toast.success("Đã tạo vai trò mới");
            setIsRoleDialogOpen(false);
            setRoleFormData({ name: "", level: 0 });
            await loadInitialData();
            setSelectedRoleId(result.data.id);
        } else {
            toast.error(result.error || "Tạo vai trò thất bại");
        }

        setIsCreatingRole(false);
    };

    if (isLoading) {
        return (
            <div className="flex items-center justify-center py-20">
                <Loader2 className="h-8 w-8 animate-spin text-muted-foreground" />
            </div>
        );
    }

    const roleResources = getRoleResources();
    const availableResources = getAvailableResources();

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
                        if (!open) setRoleFormData({ name: "", level: 0 });
                    }}>
                        <DialogTrigger asChild>
                            <Button variant="outline">
                                <Plus className="mr-2 h-4 w-4" /> Thêm vai trò
                            </Button>
                        </DialogTrigger>
                        <DialogContent>
                            <DialogHeader>
                                <DialogTitle>Thêm vai trò mới</DialogTitle>
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
                                    <Button type="submit" disabled={isCreatingRole}>
                                        {isCreatingRole && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                                        Tạo vai trò
                                    </Button>
                                </DialogFooter>
                            </form>
                        </DialogContent>
                    </Dialog>
                    <Button onClick={handleSave} disabled={isSaving || pendingChanges.size === 0}>
                        {isSaving && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                        {!isSaving && <Save className="mr-2 h-4 w-4" />}
                        Lưu thay đổi {pendingChanges.size > 0 && `(${pendingChanges.size})`}
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
                                    {roles.length === 0 ? (
                                        <p className="text-center text-muted-foreground py-4">
                                            Chưa có vai trò nào
                                        </p>
                                    ) : (
                                        roles.map((role) => (
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
                                                            Level: {role.level} · {permissions.filter(p => p.role_id === role.id).length} resources
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                        ))
                                    )}
                                </div>
                            </ScrollArea>
                        </CardContent>
                    </Card>
                </div>

                {/* Permissions - Right Side */}
                <div className="col-span-8">
                    <Card className="h-full">
                        <CardHeader>
                            <div className="flex items-center justify-between">
                                <div>
                                    <CardTitle>
                                        Quyền cho: {selectedRole?.name || "Chưa chọn vai trò"}
                                    </CardTitle>
                                    <CardDescription>
                                        Chỉ hiện các resource đã được gán cho vai trò này
                                    </CardDescription>
                                </div>
                                {selectedRoleId && availableResources.length > 0 && (
                                    <Dialog open={isAddResourceDialogOpen} onOpenChange={setIsAddResourceDialogOpen}>
                                        <DialogTrigger asChild>
                                            <Button variant="outline" size="sm">
                                                <Plus className="mr-2 h-4 w-4" /> Thêm resource
                                            </Button>
                                        </DialogTrigger>
                                        <DialogContent>
                                            <DialogHeader>
                                                <DialogTitle>Thêm resource cho {selectedRole?.name}</DialogTitle>
                                                <DialogDescription>
                                                    Chọn resource muốn gán quyền cho vai trò này.
                                                </DialogDescription>
                                            </DialogHeader>
                                            <div className="space-y-4">
                                                <div className="space-y-2">
                                                    <Label>Resource</Label>
                                                    <Select value={selectedNewResourceId} onValueChange={setSelectedNewResourceId}>
                                                        <SelectTrigger>
                                                            <SelectValue placeholder="Chọn resource..." />
                                                        </SelectTrigger>
                                                        <SelectContent>
                                                            {availableResources.map(r => (
                                                                <SelectItem key={r.id} value={String(r.id)}>
                                                                    {r.key} {r.description ? `- ${r.description}` : ""}
                                                                </SelectItem>
                                                            ))}
                                                        </SelectContent>
                                                    </Select>
                                                </div>
                                            </div>
                                            <DialogFooter>
                                                <Button variant="outline" onClick={() => setIsAddResourceDialogOpen(false)}>
                                                    Hủy
                                                </Button>
                                                <Button onClick={handleAddResource} disabled={!selectedNewResourceId}>
                                                    Thêm
                                                </Button>
                                            </DialogFooter>
                                        </DialogContent>
                                    </Dialog>
                                )}
                            </div>
                        </CardHeader>
                        <CardContent>
                            {selectedRoleId ? (
                                roleResources.length === 0 ? (
                                    <div className="flex flex-col items-center justify-center h-[450px] text-muted-foreground gap-2">
                                        <p>Vai trò này chưa được gán resource nào</p>
                                        {availableResources.length > 0 && (
                                            <Button variant="outline" size="sm" onClick={() => setIsAddResourceDialogOpen(true)}>
                                                <Plus className="mr-2 h-4 w-4" /> Thêm resource
                                            </Button>
                                        )}
                                    </div>
                                ) : (
                                    <ScrollArea className="h-[450px] pr-4">
                                        <div className="space-y-4">
                                            {roleResources.map((resource) => {
                                                const key = changeKey(selectedRoleId, resource.id);
                                                const isPendingDelete = pendingChanges.get(key)?.type === "delete";
                                                if (isPendingDelete) return null;

                                                return (
                                                    <div key={resource.id} className={`border rounded-lg p-4 ${pendingChanges.has(key) ? "border-[var(--gold)]/50 bg-[var(--gold)]/5" : ""}`}>
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
                                                            <div className="flex items-center gap-2">
                                                                <span className="text-xs text-muted-foreground">
                                                                    {getPermissionCount(selectedRoleId, resource.id)}/{SCOPES.length}
                                                                </span>
                                                                <Button
                                                                    variant="ghost"
                                                                    size="sm"
                                                                    className="h-7 w-7 p-0 text-red-500 hover:text-red-600"
                                                                    onClick={() => handleRemoveResource(resource.id)}
                                                                >
                                                                    <Trash2 className="h-3.5 w-3.5" />
                                                                </Button>
                                                            </div>
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
                                                );
                                            })}
                                        </div>
                                    </ScrollArea>
                                )
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
