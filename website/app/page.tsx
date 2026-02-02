"use client";

import { Navbar, Footer, UrlShortenerWidget, PricingSection } from "@/components/shared";
import { useLanguage } from "@/lib/i18n";
import { Zap, BarChart3, Shield, Users, Sparkles } from "lucide-react";

export default function LandingPage() {
  const { t } = useLanguage();

  return (
    <div className="min-h-screen bg-background">
      <Navbar />

      {/* Hero Section */}
      <section className="relative pt-32 pb-20 md:pt-40 md:pb-28 overflow-hidden">
        {/* Background decorations */}
        <div className="absolute inset-0 -z-10">
          <div className="absolute top-1/4 left-1/4 w-96 h-96 bg-[var(--gold)]/10 rounded-full blur-3xl" />
          <div className="absolute bottom-1/4 right-1/4 w-96 h-96 bg-[var(--gold)]/5 rounded-full blur-3xl" />
        </div>

        <div className="container mx-auto px-4">
          <div className="max-w-4xl mx-auto text-center">
            {/* Badge */}
            <div className="inline-flex items-center gap-2 rounded-full border border-border/50 bg-muted/30 backdrop-blur-sm px-4 py-2 mb-8">
              <Sparkles className="h-4 w-4 text-[var(--gold)]" />
              <span className="text-sm">{t("hero.badge")}</span>
            </div>

            {/* Title */}
            <h1 className="text-4xl md:text-6xl lg:text-7xl font-bold tracking-tight mb-6">
              <span className="text-foreground">{t("hero.title1")}</span>
              <br />
              <span className="bg-gradient-to-r from-[var(--gold)] to-amber-500 bg-clip-text text-transparent">
                {t("hero.title2")}
              </span>
            </h1>

            {/* Subtitle */}
            <p className="text-lg md:text-xl text-muted-foreground max-w-2xl mx-auto mb-10">
              {t("hero.subtitle")}
            </p>

            {/* URL Shortener Widget */}
            <UrlShortenerWidget variant="hero" className="max-w-2xl mx-auto" />
          </div>
        </div>
      </section>

      {/* Stats Section */}
      <section className="py-16 border-y border-border/50 bg-muted/20">
        <div className="container mx-auto px-4">
          <div className="grid grid-cols-2 md:grid-cols-4 gap-8">
            {[
              { value: "10M+", label: t("stats.linksCreated") },
              { value: "500M+", label: t("stats.clicksTracked") },
              { value: "99.9%", label: t("stats.uptime") },
              { value: "50K+", label: t("stats.happyTeams") },
            ].map((stat) => (
              <div key={stat.label} className="text-center">
                <div className="text-3xl md:text-4xl font-bold text-[var(--gold)] mb-2">
                  {stat.value}
                </div>
                <div className="text-sm text-muted-foreground">{stat.label}</div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-20 md:py-28">
        <div className="container mx-auto px-4">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-4xl font-bold mb-4">
              {t("features.title")}
            </h2>
            <p className="text-lg text-muted-foreground max-w-2xl mx-auto">
              {t("features.subtitle")}
            </p>
          </div>

          <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-6">
            {[
              {
                icon: Zap,
                title: t("features.fast.title"),
                description: t("features.fast.desc"),
              },
              {
                icon: BarChart3,
                title: t("features.analytics.title"),
                description: t("features.analytics.desc"),
              },
              {
                icon: Shield,
                title: t("features.security.title"),
                description: t("features.security.desc"),
              },
              {
                icon: Users,
                title: t("features.team.title"),
                description: t("features.team.desc"),
              },
            ].map((feature) => (
              <div
                key={feature.title}
                className="group p-6 rounded-2xl border border-border/50 bg-background hover:border-[var(--gold)]/50 hover:bg-muted/30 transition-all duration-300"
              >
                <div className="flex h-12 w-12 items-center justify-center rounded-xl bg-[var(--gold)]/10 text-[var(--gold)] mb-4 group-hover:scale-110 transition-transform">
                  <feature.icon className="h-6 w-6" />
                </div>
                <h3 className="text-lg font-semibold mb-2">{feature.title}</h3>
                <p className="text-sm text-muted-foreground">{feature.description}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Pricing Section */}
      <section className="py-20 md:py-28" id="pricing">
        <div className="container mx-auto px-4">
          <PricingSection />
        </div>
      </section>

      <Footer />
    </div>
  );
}
