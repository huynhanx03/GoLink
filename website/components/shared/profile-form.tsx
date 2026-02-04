"use client";

import { useState } from "react";
import { useLanguage } from "@/lib/i18n";
import { User } from "@/lib/types";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select";
import { toast } from "sonner";
import { Loader2, Camera, Users, Calendar } from "lucide-react";

interface ProfileFormProps {
    user: User;
    onSave?: (data: Partial<User>) => Promise<void>;
}

export function ProfileForm({ user, onSave }: ProfileFormProps) {
    const { t } = useLanguage();
    const [isLoading, setIsLoading] = useState(false);

    // Split name into first and last name
    const nameParts = user.name.split(" ");
    const initialFirstName = nameParts[0] || "";
    const initialLastName = nameParts.slice(1).join(" ") || "";

    const [firstName, setFirstName] = useState(initialFirstName);
    const [lastName, setLastName] = useState(initialLastName);
    const [email, setEmail] = useState(user.email || "");
    const [gender, setGender] = useState<string>("not_specified");
    const [birthday, setBirthday] = useState<string>("");
    const [avatar, setAvatar] = useState(user.avatar || "");

    // Password fields
    const [currentPassword, setCurrentPassword] = useState("");
    const [newPassword, setNewPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");

    const handleSave = async (e: React.FormEvent) => {
        e.preventDefault();
        setIsLoading(true);

        try {
            if (newPassword && newPassword !== confirmPassword) {
                toast.error("New passwords do not match");
                setIsLoading(false);
                return;
            }

            const fullName = `${firstName} ${lastName}`.trim();
            const updates: Partial<User> & { password?: string; currentPassword?: string } = {
                name: fullName,
                email,
                avatar,
            };

            if (newPassword) {
                updates.password = newPassword;
                updates.currentPassword = currentPassword;
            }

            if (onSave) {
                await onSave(updates);
            } else {
                await new Promise(resolve => setTimeout(resolve, 1000));
            }

            toast.success("Profile updated successfully");

            // Reset password fields
            setCurrentPassword("");
            setNewPassword("");
            setConfirmPassword("");

        } catch (error) {
            console.error(error);
            toast.error("Failed to update profile");
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <form onSubmit={handleSave} className="space-y-6">
            <Card>
                <CardHeader>
                    <CardTitle>{t("profile.title")}</CardTitle>
                    <CardDescription>
                        {t("profile.subtitle")}
                    </CardDescription>
                </CardHeader>
                <CardContent className="space-y-6">
                    {/* Avatar Section */}
                    <div className="flex flex-col items-center sm:flex-row gap-6">
                        <div className="relative group cursor-pointer">
                            <Avatar className="h-24 w-24">
                                <AvatarImage src={avatar} alt={firstName} />
                                <AvatarFallback className="text-xl">
                                    {firstName.charAt(0).toUpperCase()}
                                </AvatarFallback>
                            </Avatar>
                            <div className="absolute inset-0 flex items-center justify-center bg-black/60 rounded-full opacity-0 group-hover:opacity-100 transition-opacity">
                                <Camera className="h-6 w-6 text-white" />
                            </div>
                        </div>
                        <div className="space-y-1 text-center sm:text-left">
                            <h3 className="font-medium">{t("profile.picture")}</h3>
                            <p className="text-sm text-muted-foreground">
                                {t("profile.pictureHint")}
                            </p>
                            <Button variant="outline" size="sm" type="button" onClick={() => {
                                setAvatar(`https://api.dicebear.com/7.x/avataaars/svg?seed=${Date.now()}`);
                            }}>
                                {t("profile.randomize")}
                            </Button>
                        </div>
                    </div>

                    <div className="grid gap-4 sm:grid-cols-2">
                        <div className="space-y-2">
                            <Label htmlFor="firstName">{t("profile.firstName")}</Label>
                            <Input
                                id="firstName"
                                value={firstName}
                                onChange={(e) => setFirstName(e.target.value)}
                                placeholder={t("profile.firstName")}
                            />
                        </div>
                        <div className="space-y-2">
                            <Label htmlFor="lastName">{t("profile.lastName")}</Label>
                            <Input
                                id="lastName"
                                value={lastName}
                                onChange={(e) => setLastName(e.target.value)}
                                placeholder={t("profile.lastName")}
                            />
                        </div>
                    </div>

                    <div className="grid gap-4 sm:grid-cols-2">
                        <div className="space-y-2">
                            <Label htmlFor="gender">{t("profile.gender")}</Label>
                            <Select value={gender} onValueChange={setGender}>
                                <SelectTrigger className="w-full">
                                    <Users className="mr-2 h-4 w-4 text-muted-foreground" />
                                    <SelectValue placeholder={t("auth.selectGender")} />
                                </SelectTrigger>
                                <SelectContent className="min-w-[var(--radix-select-trigger-width)]">
                                    <SelectItem value="not_specified">{t("profile.notSpecified")}</SelectItem>
                                    <SelectItem value="male">{t("profile.male")}</SelectItem>
                                    <SelectItem value="female">{t("profile.female")}</SelectItem>
                                    <SelectItem value="other">{t("profile.other")}</SelectItem>
                                </SelectContent>
                            </Select>
                        </div>
                        <div className="space-y-2">
                            <Label htmlFor="birthday">{t("profile.birthday")}</Label>
                            <div className="relative">
                                <Calendar className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
                                <Input
                                    id="birthday"
                                    type="date"
                                    value={birthday}
                                    onChange={(e) => setBirthday(e.target.value)}
                                    className="pl-10"
                                />
                            </div>
                        </div>
                    </div>

                    <div className="grid gap-4 sm:grid-cols-2">
                        <div className="space-y-2">
                            <Label htmlFor="username">{t("profile.username")}</Label>
                            <Input
                                id="username"
                                value={user.username}
                                disabled
                                className="bg-muted"
                            />
                            <p className="text-[0.8rem] text-muted-foreground">
                                {t("profile.usernameHint")}
                            </p>
                        </div>
                        <div className="space-y-2">
                            <Label htmlFor="email">{t("profile.email")}</Label>
                            <Input
                                id="email"
                                type="email"
                                value={email}
                                onChange={(e) => setEmail(e.target.value)}
                                placeholder="your@email.com"
                            />
                        </div>
                    </div>
                </CardContent>
            </Card>

            <Card>
                <CardHeader>
                    <CardTitle>{t("profile.security")}</CardTitle>
                    <CardDescription>
                        {t("profile.securitySubtitle")}
                    </CardDescription>
                </CardHeader>
                <CardContent className="space-y-4">
                    <div className="space-y-2">
                        <Label htmlFor="current-password">{t("profile.currentPassword")}</Label>
                        <Input
                            id="current-password"
                            type="password"
                            value={currentPassword}
                            onChange={(e) => setCurrentPassword(e.target.value)}
                            placeholder={t("profile.currentPasswordPlaceholder")}
                        />
                    </div>
                    <div className="grid gap-4 sm:grid-cols-2">
                        <div className="space-y-2">
                            <Label htmlFor="new-password">{t("profile.newPassword")}</Label>
                            <Input
                                id="new-password"
                                type="password"
                                value={newPassword}
                                onChange={(e) => setNewPassword(e.target.value)}
                            />
                        </div>
                        <div className="space-y-2">
                            <Label htmlFor="confirm-password">{t("profile.confirmPassword")}</Label>
                            <Input
                                id="confirm-password"
                                type="password"
                                value={confirmPassword}
                                onChange={(e) => setConfirmPassword(e.target.value)}
                            />
                        </div>
                    </div>
                </CardContent>
                <CardFooter className="flex justify-end border-t pt-6">
                    <Button type="submit" disabled={isLoading}>
                        {isLoading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                        {t("profile.saveChanges")}
                    </Button>
                </CardFooter>
            </Card>
        </form>
    );
}
