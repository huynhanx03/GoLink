// Core TypeScript interfaces for GoLink SaaS

// ============================================
// User Types
// ============================================

export interface User {
  id: string;
  username: string;
  email?: string;
  name: string;
  avatar?: string;
  role: "user" | "admin";
  createdAt: Date;
  updatedAt: Date;
}

export interface UserWithTenants extends User {
  tenants: TenantMembership[];
}

// ============================================
// Tenant/Workspace Types
// ============================================

export interface Tenant {
  id: string;
  name: string;
  slug: string;
  planId: string;
  ownerId: string;
  domain?: string;
  logo?: string;
  createdAt: Date;
  updatedAt: Date;
}

export interface TenantMembership {
  tenantId: string;
  tenantName: string;
  tenantSlug: string;
  role: "owner" | "admin" | "member";
  joinedAt: Date;
}

export interface TenantWithPlan extends Tenant {
  plan: Plan;
  memberCount: number;
  linkCount: number;
}

// ============================================
// Plan/Subscription Types
// ============================================

export type BillingInterval = "monthly" | "yearly";

export interface PlanFeature {
  key: string;
  params?: Record<string, string | number>;
  fallback?: string;
}

export type Feature = string | PlanFeature;

export interface Plan {
  id: string;
  name: string;
  description: string;
  price: number;
  interval: BillingInterval;
  features: Feature[];
  isActive: boolean;
  isPopular?: boolean;
  createdAt: Date;
  updatedAt: Date;
}

// ============================================
// Short Link Types
// ============================================

export interface ShortLink {
  id: string;
  tenantId: string;
  originalUrl: string;
  shortCode: string;
  customAlias?: string;
  clicks: number;
  isActive: boolean;
  expiresAt?: Date;
  password?: string;
  createdAt: Date;
  updatedAt: Date;
  createdBy: string;
}

export interface ShortLinkWithCreator extends ShortLink {
  creator: Pick<User, "id" | "username" | "name" | "email" | "avatar">;
}

export interface CreateShortLinkInput {
  originalUrl: string;
  customAlias?: string;
  expiresAt?: Date;
  password?: string;
}

// ============================================
// Domain Types
// ============================================

export interface CustomDomain {
  id: string;
  tenantId: string;
  domain: string;
  isVerified: boolean;
  isActive: boolean;
  verificationToken: string;
  createdAt: Date;
  updatedAt: Date;
}

export interface Role {
  id: string;
  name: string;
  level: number;
  parentId?: string;
  lft: number;
  rgt: number;
  createdAt: Date;
  updatedAt: Date;
}

export interface Resource {
  id: string;
  key: string;
  description?: string;
  createdAt: Date;
  updatedAt: Date;
}

export interface Permission {
  id: string;
  roleId: string;
  resourceId: string;
  description?: string;
  scopes: number; // Bitmask of allowed operations
  createdAt: Date;
  updatedAt: Date;
}

// ============================================
// Attribute Types
// ============================================

export interface AttributeDefinition {
  id: string;
  key: string;
  dataType: string; // string, number, boolean, date, etc.
  description?: string;
  createdAt: Date;
  updatedAt: Date;
}

// ============================================
// Auth Types
// ============================================

export interface AuthState {
  user: UserWithTenants | null;
  isAuthenticated: boolean;
  isAdmin: boolean;
  isLoading: boolean;
}

export interface LoginCredentials {
  username: string; // Changed from email
  password: string;
  rememberMe?: boolean;
}

export interface RegisterData {
  name: string;
  username: string;
  password: string;
  confirmPassword: string;
}

// ============================================
// API Response Types
// ============================================

export interface ApiResponse<T> {
  success: boolean;
  data?: T;
  error?: string;
  message?: string;
}

export interface PaginatedResponse<T> {
  items: T[];
  total: number;
  page: number;
  pageSize: number;
  totalPages: number;
}

// ============================================
// Stats Types
// ============================================

export interface DashboardStats {
  totalLinks: number;
  totalClicks: number;
  activeLinks: number;
  linksThisMonth: number;
}

export interface AdminStats {
  totalUsers: number;
  totalTenants: number;
  totalLinks: number;
  totalClicks: number;
  activeSubscriptions: number;
  monthlyRevenue: number;
}
