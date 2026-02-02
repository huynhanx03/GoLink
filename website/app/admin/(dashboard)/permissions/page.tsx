"use client";

import { useState } from "react";
import { Role, Resource, Permission } from "@/lib/types";
import { fakeRoles, fakeResources, fakePermissions } from "@/lib/data/fake-data";
import { Button } from "@/components/ui/button";
import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table";
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/components/ui/card";
import { Checkbox } from "@/components/ui/checkbox";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { toast } from "sonner";
import { Loader2, Save } from "lucide-react";

// Scope constants
const SCOPE_READ = 1;
const SCOPE_CREATE = 2;
const SCOPE_UPDATE = 4;
const SCOPE_DELETE = 8;

const SCOPES = [
    { value: SCOPE_READ, label: "Read" },
    { value: SCOPE_CREATE, label: "Create" },
    { value: SCOPE_UPDATE, label: "Update" },
    { value: SCOPE_DELETE, label: "Delete" },
];

export default function PermissionsPage() {
    const [roles] = useState<Role[]>(fakeRoles);
    const [resources] = useState<Resource[]>(fakeResources);
    const [permissions, setPermissions] = useState<Permission[]>(fakePermissions);
    const [activeRoleId, setActiveRoleId] = useState<string>(roles[0]?.id || "");
    const [isSaving, setIsSaving] = useState(false);

    const hasScope = (roleId: string, resourceId: string, scope: number) => {
        const perm = permissions.find(p => p.roleId === roleId && p.resourceId === resourceId);
        if (!perm) return false;
        return (perm.scopes & scope) === scope;
    };

    const toggleScope = (roleId: string, resourceId: string, scope: number) => {
        const existingPerm = permissions.find(p => p.roleId === roleId && p.resourceId === resourceId);

        if (existingPerm) {
            const newScopes = (existingPerm.scopes & scope) === scope
                ? existingPerm.scopes & ~scope // Remove scope
                : existingPerm.scopes | scope; // Add scope

            setPermissions(permissions.map(p =>
                p.id === existingPerm.id ? { ...p, scopes: newScopes, updatedAt: new Date() } : p
            ));
        } else {
            // Create new permission entry if it doesn't exist
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

    const handleSave = async () => {
        setIsSaving(true);
        // Simulate API call
        await new Promise(resolve => setTimeout(resolve, 800));
        setIsSaving(false);
        toast.success("Permissions saved successfully");
    };

    return (
        <div className="space-y-4">
            <div className="flex items-center justify-between">
                <div>
                    <h2 className="text-3xl font-bold tracking-tight">Permissions</h2>
                    <p className="text-muted-foreground">
                        Manage role-based access control (RBAC) permissions.
                    </p>
                </div>
                <div className="flex items-center space-x-2">
                    <Button onClick={handleSave} disabled={isSaving}>
                        {isSaving && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                        {!isSaving && <Save className="mr-2 h-4 w-4" />}
                        Save Changes
                    </Button>
                </div>
            </div>

            <Tabs value={activeRoleId} onValueChange={setActiveRoleId} className="space-y-4">
                <TabsList className="justify-start w-full overflow-x-auto h-auto p-1">
                    {roles.map((role) => (
                        <TabsTrigger key={role.id} value={role.id} className="min-w-[100px]">
                            {role.name}
                        </TabsTrigger>
                    ))}
                </TabsList>
                {roles.map((role) => (
                    <TabsContent key={role.id} value={role.id} className="space-y-4">
                        <Card>
                            <CardHeader>
                                <CardTitle>{role.name} Permissions</CardTitle>
                                <CardDescription>
                                    Configure access levels for the {role.name} role across system resources.
                                </CardDescription>
                            </CardHeader>
                            <CardContent>
                                <Table>
                                    <TableHeader>
                                        <TableRow>
                                            <TableHead className="w-[200px]">Resource</TableHead>
                                            {SCOPES.map(scope => (
                                                <TableHead key={scope.value} className="text-center w-[100px]">
                                                    {scope.label}
                                                </TableHead>
                                            ))}
                                            <TableHead className="text-right">Total Scopes</TableHead>
                                        </TableRow>
                                    </TableHeader>
                                    <TableBody>
                                        {resources.map((resource) => {
                                            const perm = permissions.find(p => p.roleId === role.id && p.resourceId === resource.id);
                                            const currentScopes = perm?.scopes || 0;

                                            return (
                                                <TableRow key={resource.id}>
                                                    <TableCell className="font-medium">
                                                        <div>{resource.key}</div>
                                                        <div className="text-xs text-muted-foreground">{resource.description}</div>
                                                    </TableCell>
                                                    {SCOPES.map(scope => (
                                                        <TableCell key={scope.value} className="text-center">
                                                            <div className="flex justify-center">
                                                                <Checkbox
                                                                    checked={hasScope(role.id, resource.id, scope.value)}
                                                                    onCheckedChange={() => toggleScope(role.id, resource.id, scope.value)}
                                                                />
                                                            </div>
                                                        </TableCell>
                                                    ))}
                                                    <TableCell className="text-right font-mono text-xs text-muted-foreground">
                                                        {currentScopes.toString(2).padStart(4, '0')} ({currentScopes})
                                                    </TableCell>
                                                </TableRow>
                                            );
                                        })}
                                    </TableBody>
                                </Table>
                            </CardContent>
                        </Card>
                    </TabsContent>
                ))}
            </Tabs>
        </div>
    );
}
