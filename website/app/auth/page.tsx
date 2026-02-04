"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import Image from "next/image";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Checkbox } from "@/components/ui/checkbox";
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select";
import { useLanguage } from "@/lib/i18n";
import { authService } from "@/lib/api";
import { toast } from "sonner";
import { Loader2, Lock, User, Calendar, Users } from "lucide-react";
import { Navbar } from "@/components/shared";

export default function AuthPage() {
    const router = useRouter();
    const { t } = useLanguage();
    const [activeTab, setActiveTab] = useState<"login" | "register">("login");
    const [isLoading, setIsLoading] = useState(false);

    // Login form state
    const [loginUsername, setLoginUsername] = useState("");
    const [loginPassword, setLoginPassword] = useState("");
    const [rememberMe, setRememberMe] = useState(false);

    // Register form state
    const [registerUsername, setRegisterUsername] = useState("");
    const [registerPassword, setRegisterPassword] = useState("");
    const [registerFirstName, setRegisterFirstName] = useState("");
    const [registerLastName, setRegisterLastName] = useState("");
    const [registerGender, setRegisterGender] = useState<string>("");
    const [registerBirthday, setRegisterBirthday] = useState("");

    const handleLogin = async (e: React.FormEvent) => {
        e.preventDefault();
        setIsLoading(true);

        try {
            const result = await authService.login({
                username: loginUsername,
                password: loginPassword,
            });

            if (result.success) {
                toast.success("Welcome back!");
                router.push("/dashboard");
            } else {
                toast.error(result.error || "Login failed");
            }
        } catch {
            toast.error("Something went wrong");
        } finally {
            setIsLoading(false);
        }
    };

    const handleRegister = async (e: React.FormEvent) => {
        e.preventDefault();

        if (registerPassword.length < 8) {
            toast.error("Password must be at least 8 characters");
            return;
        }

        if (!registerBirthday) {
            toast.error("Please enter your birthday");
            return;
        }

        if (!registerGender) {
            toast.error("Please select your gender");
            return;
        }

        setIsLoading(true);

        try {
            const result = await authService.register({
                username: registerUsername,
                password: registerPassword,
                first_name: registerFirstName,
                last_name: registerLastName,
                gender: parseInt(registerGender, 10),
                birthday: registerBirthday, // Format: "YYYY-MM-DD"
            });

            if (result.success) {
                toast.success("Account created successfully! Please login.");
                // Switch to login tab
                setActiveTab("login");
                // Clear register form
                setRegisterUsername("");
                setRegisterPassword("");
                setRegisterFirstName("");
                setRegisterLastName("");
                setRegisterGender("");
                setRegisterBirthday("");
            } else {
                toast.error(result.error || "Registration failed");
            }
        } catch {
            toast.error("Something went wrong");
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <div className="min-h-screen relative">
            {/* Background Image */}
            <div className="fixed inset-0 -z-10">
                <Image
                    src="/auth-bg.png"
                    alt="Background"
                    fill
                    className="object-cover"
                    priority
                />
                <div className="absolute inset-0 bg-black/50 backdrop-blur-sm" />
            </div>

            {/* Navbar */}
            <Navbar />

            {/* Centered Auth Card */}
            <div className="min-h-screen flex items-start justify-center px-4 pt-32 pb-8">
                <div className="w-full max-w-md bg-background/95 backdrop-blur-xl rounded-3xl border border-border/50 shadow-2xl p-8">
                    {/* Toggle Switch Tab Selector */}
                    <div className="relative flex bg-muted rounded-full p-1 mb-8">
                        {/* Sliding Background */}
                        <div
                            className={`absolute top-1 bottom-1 w-[calc(50%-4px)] bg-[var(--gold)] rounded-full transition-all duration-300 ease-in-out ${activeTab === "login" ? "left-1" : "left-[calc(50%+2px)]"
                                }`}
                        />
                        {/* Login Button */}
                        <button
                            onClick={() => setActiveTab("login")}
                            className={`relative z-10 flex-1 py-2.5 text-center text-sm font-medium transition-colors duration-300 rounded-full cursor-pointer ${activeTab === "login"
                                ? "text-[var(--gold-foreground)]"
                                : "text-muted-foreground hover:text-foreground"
                                }`}
                        >
                            {t("auth.login")}
                        </button>
                        {/* Register Button */}
                        <button
                            onClick={() => setActiveTab("register")}
                            className={`relative z-10 flex-1 py-2.5 text-center text-sm font-medium transition-colors duration-300 rounded-full cursor-pointer ${activeTab === "register"
                                ? "text-[var(--gold-foreground)]"
                                : "text-muted-foreground hover:text-foreground"
                                }`}
                        >
                            {t("auth.register")}
                        </button>
                    </div>

                    {/* Form Container */}
                    <div>
                        {activeTab === "login" ? (
                            /* Login Form */
                            <form onSubmit={handleLogin} className="space-y-4">
                                <div className="space-y-2">
                                    <Label htmlFor="login-username">{t("auth.username")}</Label>
                                    <div className="relative">
                                        <User className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
                                        <Input
                                            id="login-username"
                                            type="text"
                                            placeholder={t("auth.enterUsername")}
                                            value={loginUsername}
                                            onChange={(e) => setLoginUsername(e.target.value)}
                                            className="pl-10"
                                            required
                                        />
                                    </div>
                                </div>

                                <div className="space-y-2">
                                    <Label htmlFor="login-password">{t("auth.password")}</Label>
                                    <div className="relative">
                                        <Lock className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
                                        <Input
                                            id="login-password"
                                            type="password"
                                            placeholder="••••••••"
                                            value={loginPassword}
                                            onChange={(e) => setLoginPassword(e.target.value)}
                                            className="pl-10"
                                            required
                                        />
                                    </div>
                                </div>

                                {/* Remember Me & Forgot Password on same line */}
                                <div className="flex items-center justify-between">
                                    <div className="flex items-center gap-2">
                                        <Checkbox
                                            id="remember"
                                            checked={rememberMe}
                                            onCheckedChange={(c) => setRememberMe(c === true)}
                                        />
                                        <Label htmlFor="remember" className="text-sm font-normal cursor-pointer">
                                            {t("auth.remember30days")}
                                        </Label>
                                    </div>
                                    <Link
                                        href="/forgot-password"
                                        className="text-sm text-[var(--gold)] hover:underline"
                                    >
                                        {t("auth.forgotPassword")}
                                    </Link>
                                </div>

                                <Button
                                    type="submit"
                                    disabled={isLoading}
                                    className="w-full bg-[var(--gold)] hover:bg-[var(--gold)]/90 text-[var(--gold-foreground)] cursor-pointer mt-6"
                                >
                                    {isLoading ? (
                                        <Loader2 className="h-4 w-4 animate-spin" />
                                    ) : (
                                        t("auth.signIn")
                                    )}
                                </Button>
                            </form>
                        ) : (
                            /* Register Form */
                            <form onSubmit={handleRegister} className="space-y-4">
                                <div className="space-y-2">
                                    <Label htmlFor="register-username">{t("auth.username")}</Label>
                                    <div className="relative">
                                        <User className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
                                        <Input
                                            id="register-username"
                                            type="text"
                                            placeholder={t("auth.chooseUsername")}
                                            value={registerUsername}
                                            onChange={(e) => setRegisterUsername(e.target.value)}
                                            className="pl-10"
                                            required
                                            minLength={3}
                                            maxLength={50}
                                        />
                                    </div>
                                </div>

                                <div className="space-y-2">
                                    <Label htmlFor="register-password">{t("auth.password")}</Label>
                                    <div className="relative">
                                        <Lock className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
                                        <Input
                                            id="register-password"
                                            type="password"
                                            placeholder="••••••••"
                                            value={registerPassword}
                                            onChange={(e) => setRegisterPassword(e.target.value)}
                                            className="pl-10"
                                            required
                                            minLength={8}
                                        />
                                    </div>
                                    <p className="text-xs text-muted-foreground">{t("auth.minChars")}</p>
                                </div>

                                {/* First Name & Last Name side by side */}
                                <div className="grid grid-cols-2 gap-4">
                                    <div className="space-y-2">
                                        <Label htmlFor="register-firstname">{t("auth.firstName")}</Label>
                                        <Input
                                            id="register-firstname"
                                            type="text"
                                            placeholder={t("auth.firstName")}
                                            value={registerFirstName}
                                            onChange={(e) => setRegisterFirstName(e.target.value)}
                                            required
                                            maxLength={100}
                                        />
                                    </div>
                                    <div className="space-y-2">
                                        <Label htmlFor="register-lastname">{t("auth.lastName")}</Label>
                                        <Input
                                            id="register-lastname"
                                            type="text"
                                            placeholder={t("auth.lastName")}
                                            value={registerLastName}
                                            onChange={(e) => setRegisterLastName(e.target.value)}
                                            required
                                            maxLength={100}
                                        />
                                    </div>
                                </div>

                                {/* Gender & Birthday side by side */}
                                <div className="grid grid-cols-2 gap-4">
                                    <div className="space-y-2">
                                        <Label htmlFor="register-gender">{t("auth.gender")}</Label>
                                        <Select
                                            value={registerGender}
                                            onValueChange={setRegisterGender}
                                        >
                                            <SelectTrigger id="register-gender" className="w-full">
                                                <Users className="h-4 w-4 text-muted-foreground mr-2" />
                                                <SelectValue placeholder={t("auth.selectGender")} />
                                            </SelectTrigger>
                                            <SelectContent>
                                                <SelectItem value="0">{t("auth.genderMale")}</SelectItem>
                                                <SelectItem value="1">{t("auth.genderFemale")}</SelectItem>
                                                <SelectItem value="2">{t("auth.genderOther")}</SelectItem>
                                            </SelectContent>
                                        </Select>
                                    </div>
                                    <div className="space-y-2">
                                        <Label htmlFor="register-birthday">{t("auth.birthday")}</Label>
                                        <div className="relative">
                                            <Calendar className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
                                            <Input
                                                id="register-birthday"
                                                type="date"
                                                value={registerBirthday}
                                                onChange={(e) => setRegisterBirthday(e.target.value)}
                                                className="pl-10"
                                                required
                                            />
                                        </div>
                                    </div>
                                </div>

                                <Button
                                    type="submit"
                                    disabled={isLoading}
                                    className="w-full bg-[var(--gold)] hover:bg-[var(--gold)]/90 text-[var(--gold-foreground)] cursor-pointer mt-6"
                                >
                                    {isLoading ? (
                                        <Loader2 className="h-4 w-4 animate-spin" />
                                    ) : (
                                        t("auth.createAccountBtn")
                                    )}
                                </Button>
                            </form>
                        )}
                    </div>
                </div>
            </div>
        </div>
    );
}
