"use client";

import { useState } from "react";
import { useLanguage } from "@/lib/i18n";
import { AttributeDefinition } from "@/lib/types";
import { fakeAttributeDefinitions } from "@/lib/data/fake-data";
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
import { Textarea } from "@/components/ui/textarea";
import { toast } from "sonner";
import { Plus, Pencil, Trash2, Search, Tags } from "lucide-react";

const DATA_TYPES = [
    { value: "string", label: "String" },
    { value: "number", label: "Number" },
    { value: "boolean", label: "Boolean" },
    { value: "date", label: "Date" },
    { value: "json", label: "JSON" },
];

export default function AttributesPage() {
    const { t } = useLanguage();
    const [attributes, setAttributes] = useState<AttributeDefinition[]>(fakeAttributeDefinitions);
    const [searchQuery, setSearchQuery] = useState("");
    const [isDialogOpen, setIsDialogOpen] = useState(false);
    const [editingAttribute, setEditingAttribute] = useState<AttributeDefinition | null>(null);
    const [formData, setFormData] = useState({
        key: "",
        dataType: "string",
        description: "",
    });

    const filteredAttributes = attributes.filter(
        (attr) =>
            attr.key.toLowerCase().includes(searchQuery.toLowerCase()) ||
            attr.description?.toLowerCase().includes(searchQuery.toLowerCase())
    );

    const resetForm = () => {
        setFormData({ key: "", dataType: "string", description: "" });
        setEditingAttribute(null);
    };

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();

        if (editingAttribute) {
            setAttributes(attributes.map(attr =>
                attr.id === editingAttribute.id
                    ? { ...attr, ...formData, updatedAt: new Date() }
                    : attr
            ));
            toast.success(t("admin.attributes.editAttribute") + " ✓");
        } else {
            const newAttr: AttributeDefinition = {
                id: `attr-${Date.now()}`,
                key: formData.key,
                dataType: formData.dataType,
                description: formData.description,
                createdAt: new Date(),
                updatedAt: new Date(),
            };
            setAttributes([...attributes, newAttr]);
            toast.success(t("admin.attributes.createAttribute") + " ✓");
        }

        setIsDialogOpen(false);
        resetForm();
    };

    const handleEdit = (attr: AttributeDefinition) => {
        setEditingAttribute(attr);
        setFormData({
            key: attr.key,
            dataType: attr.dataType,
            description: attr.description || "",
        });
        setIsDialogOpen(true);
    };

    const handleDelete = (id: string) => {
        if (confirm(t("common.delete") + "?")) {
            setAttributes(attributes.filter(attr => attr.id !== id));
            toast.success(t("common.delete") + " ✓");
        }
    };

    return (
        <div className="space-y-4">
            <div className="flex items-center justify-between">
                <div>
                    <h2 className="text-3xl font-bold tracking-tight">{t("admin.attributes.title")}</h2>
                    <p className="text-muted-foreground">
                        {t("admin.attributes.subtitle")}
                    </p>
                </div>
                <Dialog open={isDialogOpen} onOpenChange={(open) => {
                    setIsDialogOpen(open);
                    if (!open) resetForm();
                }}>
                    <DialogTrigger asChild>
                        <Button>
                            <Plus className="mr-2 h-4 w-4" />
                            {t("admin.attributes.addAttribute")}
                        </Button>
                    </DialogTrigger>
                    <DialogContent>
                        <DialogHeader>
                            <DialogTitle>
                                {editingAttribute ? t("admin.attributes.editAttribute") : t("admin.attributes.createAttribute")}
                            </DialogTitle>
                            <DialogDescription>
                                {t("admin.attributes.defineNew")}
                            </DialogDescription>
                        </DialogHeader>
                        <form onSubmit={handleSubmit} className="space-y-4">
                            <div className="space-y-2">
                                <Label htmlFor="key">{t("admin.attributes.key")}</Label>
                                <Input
                                    id="key"
                                    placeholder={t("admin.attributes.keyPlaceholder")}
                                    value={formData.key}
                                    onChange={(e) => setFormData({ ...formData, key: e.target.value })}
                                    required
                                />
                                <p className="text-[0.8rem] text-muted-foreground">
                                    {t("admin.attributes.keyHint")}
                                </p>
                            </div>
                            <div className="space-y-2">
                                <Label htmlFor="dataType">{t("admin.attributes.dataType")}</Label>
                                <Select
                                    value={formData.dataType}
                                    onValueChange={(value) => setFormData({ ...formData, dataType: value })}
                                >
                                    <SelectTrigger className="w-full">
                                        <SelectValue placeholder={t("admin.attributes.dataType")} />
                                    </SelectTrigger>
                                    <SelectContent>
                                        {DATA_TYPES.map((type) => (
                                            <SelectItem key={type.value} value={type.value}>
                                                {type.label}
                                            </SelectItem>
                                        ))}
                                    </SelectContent>
                                </Select>
                            </div>
                            <div className="space-y-2">
                                <Label htmlFor="description">{t("admin.attributes.description")}</Label>
                                <Textarea
                                    id="description"
                                    placeholder={t("admin.attributes.descriptionPlaceholder")}
                                    value={formData.description}
                                    onChange={(e) => setFormData({ ...formData, description: e.target.value })}
                                    className="min-h-[120px] max-h-[200px] resize-y"
                                />
                            </div>
                            <DialogFooter>
                                <Button type="button" variant="outline" onClick={() => setIsDialogOpen(false)}>
                                    {t("admin.attributes.cancel")}
                                </Button>
                                <Button type="submit">
                                    {editingAttribute ? t("admin.attributes.saveChanges") : t("admin.attributes.createAttribute")}
                                </Button>
                            </DialogFooter>
                        </form>
                    </DialogContent>
                </Dialog>
            </div>

            <div className="flex items-center gap-4">
                <div className="relative flex-1 max-w-sm">
                    <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
                    <Input
                        placeholder={t("admin.attributes.searchPlaceholder")}
                        className="pl-10"
                        value={searchQuery}
                        onChange={(e) => setSearchQuery(e.target.value)}
                    />
                </div>
            </div>

            <div className="rounded-md border">
                <Table>
                    <TableHeader>
                        <TableRow>
                            <TableHead>{t("admin.attributes.key")}</TableHead>
                            <TableHead>{t("admin.attributes.dataType")}</TableHead>
                            <TableHead className="hidden md:table-cell">{t("admin.attributes.description")}</TableHead>
                            <TableHead className="hidden md:table-cell">{t("admin.attributes.created")}</TableHead>
                            <TableHead className="text-right">{t("admin.attributes.actions")}</TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {filteredAttributes.length === 0 ? (
                            <TableRow>
                                <TableCell colSpan={5} className="text-center py-8">
                                    <Tags className="h-8 w-8 mx-auto mb-2 text-muted-foreground" />
                                    <p className="text-muted-foreground">{t("admin.attributes.noAttributes")}</p>
                                </TableCell>
                            </TableRow>
                        ) : (
                            filteredAttributes.map((attr) => (
                                <TableRow key={attr.id}>
                                    <TableCell className="font-medium font-mono text-sm">
                                        {attr.key}
                                    </TableCell>
                                    <TableCell>
                                        <span className="inline-flex items-center rounded-md bg-muted px-2 py-1 text-xs font-medium">
                                            {attr.dataType}
                                        </span>
                                    </TableCell>
                                    <TableCell className="hidden md:table-cell text-muted-foreground">
                                        {attr.description || "-"}
                                    </TableCell>
                                    <TableCell className="hidden md:table-cell text-muted-foreground">
                                        {attr.createdAt.toLocaleDateString()}
                                    </TableCell>
                                    <TableCell className="text-right">
                                        <div className="flex justify-end gap-2">
                                            <Button
                                                variant="ghost"
                                                size="sm"
                                                onClick={() => handleEdit(attr)}
                                            >
                                                <Pencil className="h-4 w-4" />
                                            </Button>
                                            <Button
                                                variant="ghost"
                                                size="sm"
                                                className="text-red-500 hover:text-red-600"
                                                onClick={() => handleDelete(attr.id)}
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
