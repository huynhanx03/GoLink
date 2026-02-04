"use client";

import { useState } from "react";
import { Resource } from "@/lib/types";
import { fakeResources } from "@/lib/data/fake-data";
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
import { Label } from "@/components/ui/label";
import { Plus, Trash2, Pencil, Search, Database } from "lucide-react";
import { toast } from "sonner";

export default function ResourcesPage() {
    const [resources, setResources] = useState<Resource[]>(fakeResources);
    const [searchQuery, setSearchQuery] = useState("");
    const [isOpen, setIsOpen] = useState(false);
    const [editingResource, setEditingResource] = useState<Resource | null>(null);

    const [formData, setFormData] = useState({
        key: "",
        description: "",
    });

    const filteredResources = resources.filter(
        (res) =>
            res.key.toLowerCase().includes(searchQuery.toLowerCase()) ||
            (res.description && res.description.toLowerCase().includes(searchQuery.toLowerCase()))
    );

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();

        if (editingResource) {
            // Update existing
            const updated = resources.map(res =>
                res.id === editingResource.id
                    ? { ...res, key: formData.key, description: formData.description, updatedAt: new Date() }
                    : res
            );
            setResources(updated);
            toast.success("Resource updated successfully");
        } else {
            // Create new
            const newResource: Resource = {
                id: `res-${Date.now()}`,
                key: formData.key,
                description: formData.description,
                createdAt: new Date(),
                updatedAt: new Date(),
            };
            setResources([...resources, newResource]);
            toast.success("Resource created successfully");
        }

        setIsOpen(false);
        resetForm();
    };

    const resetForm = () => {
        setFormData({ key: "", description: "" });
        setEditingResource(null);
    };

    const handleEdit = (resource: Resource) => {
        setEditingResource(resource);
        setFormData({
            key: resource.key,
            description: resource.description || "",
        });
        setIsOpen(true);
    };

    const handleDelete = (id: string) => {
        if (confirm("Are you sure you want to delete this resource?")) {
            setResources(resources.filter(res => res.id !== id));
            toast.success("Resource deleted successfully");
        }
    };

    return (
        <div className="flex-1 space-y-4 p-8 pt-6">
            <div className="flex items-center justify-between space-y-2">
                <div>
                    <h2 className="text-3xl font-bold tracking-tight">Resources</h2>
                    <p className="text-muted-foreground">
                        Manage system resources for RBAC.
                    </p>
                </div>
                <div className="flex items-center space-x-2">
                    <Dialog open={isOpen} onOpenChange={(open) => {
                        setIsOpen(open);
                        if (!open) resetForm();
                    }}>
                        <DialogTrigger asChild>
                            <Button>
                                <Plus className="mr-2 h-4 w-4" /> Add Resource
                            </Button>
                        </DialogTrigger>
                        <DialogContent>
                            <DialogHeader>
                                <DialogTitle>{editingResource ? "Edit Resource" : "Add Resource"}</DialogTitle>
                                <DialogDescription>
                                    Define a protected resource that can be assigned permissions.
                                </DialogDescription>
                            </DialogHeader>
                            <form onSubmit={handleSubmit} className="space-y-4">
                                <div className="space-y-2">
                                    <Label htmlFor="key">Key</Label>
                                    <Input
                                        id="key"
                                        placeholder="e.g. links"
                                        value={formData.key}
                                        onChange={(e) => setFormData({ ...formData, key: e.target.value })}
                                        required
                                    />
                                    <p className="text-[0.8rem] text-muted-foreground">
                                        Unique identifier for the resource.
                                    </p>
                                </div>
                                <div className="space-y-2">
                                    <Label htmlFor="description">Description (Optional)</Label>
                                    <Input
                                        id="description"
                                        placeholder="e.g. Manage short links"
                                        value={formData.description}
                                        onChange={(e) => setFormData({ ...formData, description: e.target.value })}
                                    />
                                </div>
                                <DialogFooter>
                                    <Button type="button" variant="outline" onClick={() => setIsOpen(false)}>
                                        Cancel
                                    </Button>
                                    <Button type="submit">
                                        {editingResource ? "Save Changes" : "Create Resource"}
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
                        placeholder="Search resources..."
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
                            <TableHead>Key</TableHead>
                            <TableHead>Description</TableHead>
                            <TableHead className="w-[100px]">Actions</TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {filteredResources.length === 0 ? (
                            <TableRow>
                                <TableCell colSpan={3} className="h-24 text-center">
                                    No resources found.
                                </TableCell>
                            </TableRow>
                        ) : (
                            filteredResources.map((resource) => (
                                <TableRow key={resource.id}>
                                    <TableCell className="font-medium">
                                        <div className="flex items-center gap-2">
                                            <Database className="h-4 w-4 text-muted-foreground" />
                                            {resource.key}
                                        </div>
                                    </TableCell>
                                    <TableCell>{resource.description}</TableCell>
                                    <TableCell>
                                        <div className="flex items-center gap-2">
                                            <Button
                                                variant="ghost"
                                                size="icon"
                                                onClick={() => handleEdit(resource)}
                                            >
                                                <Pencil className="h-4 w-4" />
                                            </Button>
                                            <Button
                                                variant="ghost"
                                                size="icon"
                                                className="text-red-500 hover:text-red-600"
                                                onClick={() => handleDelete(resource.id)}
                                            >
                                                <Trash2 className="h-4 w-4" />
                                            </Button>
                                        </div>
                                    </TableCell>
                                </TableRow>
                            ))
                        )}
                    </TableBody>
                </Table>
            </div>
        </div>
    );
}
