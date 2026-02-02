"use client";

import { createContext, useContext, useState, useEffect, ReactNode } from "react";

type Language = "en" | "vi";

interface Translations {
    [key: string]: {
        en: string;
        vi: string;
    };
}

const translations: Translations = {
    // Navbar
    "nav.home": { en: "Home", vi: "Trang chủ" },
    "nav.login": { en: "Login", vi: "Đăng nhập" },
    "nav.getStarted": { en: "Get Started", vi: "Bắt đầu" },

    // Hero
    "hero.badge": { en: "New: Team Workspaces now available", vi: "Mới: Không gian làm việc nhóm đã có" },
    "hero.title1": { en: "Short Links,", vi: "Link Ngắn," },
    "hero.title2": { en: "Powerful Results", vi: "Kết Quả Mạnh Mẽ" },
    "hero.subtitle": {
        en: "Create, manage, and track short links with powerful analytics. The modern URL shortener built for teams and businesses.",
        vi: "Tạo, quản lý và theo dõi link ngắn với phân tích mạnh mẽ. Công cụ rút gọn URL hiện đại dành cho đội nhóm và doanh nghiệp."
    },

    // URL Shortener
    "shortener.placeholder": { en: "Paste your long URL here...", vi: "Dán URL dài của bạn vào đây..." },
    "shortener.button": { en: "Shorten URL", vi: "Rút gọn URL" },
    "shortener.copy": { en: "Copy Link", vi: "Sao chép" },
    "shortener.copied": { en: "Copied!", vi: "Đã sao chép!" },
    "shortener.another": { en: "Shorten another URL", vi: "Rút gọn URL khác" },
    "shortener.freeLimit": { en: "Free users can create up to 5 links.", vi: "Người dùng miễn phí có thể tạo tối đa 5 link." },
    "shortener.signUp": { en: "Sign up", vi: "Đăng ký" },
    "shortener.forUnlimited": { en: "for unlimited access.", vi: "để truy cập không giới hạn." },

    // Stats
    "stats.linksCreated": { en: "Links Created", vi: "Link đã tạo" },
    "stats.clicksTracked": { en: "Clicks Tracked", vi: "Lượt click" },
    "stats.uptime": { en: "Uptime", vi: "Thời gian hoạt động" },
    "stats.happyTeams": { en: "Happy Teams", vi: "Đội nhóm hài lòng" },

    // Features
    "features.title": { en: "Everything you need to manage links", vi: "Mọi thứ bạn cần để quản lý link" },
    "features.subtitle": { en: "Powerful features designed for modern teams and businesses.", vi: "Tính năng mạnh mẽ được thiết kế cho đội nhóm và doanh nghiệp hiện đại." },
    "features.fast.title": { en: "Lightning Fast", vi: "Siêu nhanh" },
    "features.fast.desc": { en: "Create short links instantly with our optimized infrastructure. No delays, just speed.", vi: "Tạo link ngắn ngay lập tức với cơ sở hạ tầng tối ưu. Không chậm trễ, chỉ có tốc độ." },
    "features.analytics.title": { en: "Powerful Analytics", vi: "Phân tích mạnh mẽ" },
    "features.analytics.desc": { en: "Track clicks, geographic data, and referrers. Make data-driven decisions.", vi: "Theo dõi click, dữ liệu địa lý và nguồn giới thiệu. Đưa ra quyết định dựa trên dữ liệu." },
    "features.security.title": { en: "Enterprise Security", vi: "Bảo mật doanh nghiệp" },
    "features.security.desc": { en: "Your links are protected with enterprise-grade security and encryption.", vi: "Link của bạn được bảo vệ với bảo mật và mã hóa cấp doanh nghiệp." },
    "features.team.title": { en: "Team Collaboration", vi: "Cộng tác nhóm" },
    "features.team.desc": { en: "Work together with your team. Manage permissions and share links effortlessly.", vi: "Làm việc cùng nhóm. Quản lý quyền và chia sẻ link dễ dàng." },

    // How it works
    "howItWorks.title": { en: "How it works", vi: "Cách hoạt động" },
    "howItWorks.subtitle": { en: "Create your first short link in seconds.", vi: "Tạo link ngắn đầu tiên trong vài giây." },
    "howItWorks.step1.title": { en: "Paste Your URL", vi: "Dán URL của bạn" },
    "howItWorks.step1.desc": { en: "Drop in any long URL you want to shorten.", vi: "Dán bất kỳ URL dài nào bạn muốn rút gọn." },
    "howItWorks.step2.title": { en: "Click Shorten", vi: "Nhấn Rút gọn" },
    "howItWorks.step2.desc": { en: "We generate a short, memorable link instantly.", vi: "Chúng tôi tạo link ngắn, dễ nhớ ngay lập tức." },
    "howItWorks.step3.title": { en: "Share Anywhere", vi: "Chia sẻ mọi nơi" },
    "howItWorks.step3.desc": { en: "Use your short link across all platforms.", vi: "Sử dụng link ngắn trên tất cả nền tảng." },

    // Pricing
    "pricing.title": { en: "Simple, transparent pricing", vi: "Giá đơn giản, minh bạch" },
    "pricing.subtitle": { en: "Choose the perfect plan for your needs.", vi: "Chọn gói phù hợp với nhu cầu của bạn." },
    "pricing.monthly": { en: "Monthly", vi: "Hàng tháng" },
    "pricing.yearly": { en: "Yearly", vi: "Hàng năm" },
    "pricing.save": { en: "Save 20%", vi: "Tiết kiệm 20%" },
    "pricing.popular": { en: "Most Popular", vi: "Phổ biến nhất" },
    "pricing.whatsIncluded": { en: "What's included:", vi: "Bao gồm:" },
    "pricing.getStartedFree": { en: "Get Started Free", vi: "Bắt đầu miễn phí" },
    "pricing.startFreeTrial": { en: "Start Free Trial", vi: "Dùng thử miễn phí" },
    "pricing.contactSales": { en: "Contact Sales", vi: "Liên hệ" },

    // Pricing Features (Parameterized)
    "pricing.feature.links": { en: "{count} links", vi: "{count} links" },
    "pricing.feature.unlimitedLinks": { en: "Unlimited links", vi: "Không giới hạn link" },
    "pricing.feature.customDomains": { en: "Custom domain", vi: "Tên miền tùy chỉnh" },
    "pricing.feature.unlimitedDomains": { en: "Unlimited custom domains", vi: "Không giới hạn tên miền" },
    "pricing.feature.retention": { en: "{days} days data retention", vi: "Lưu dữ liệu {days} ngày" },
    "pricing.feature.unlimitedRetention": { en: "Unlimited data retention", vi: "Lưu trữ không giới hạn" },

    // Generic Plan Features (Matched from fake-data)
    "pricing.feature.basicAnalytics": { en: "Basic analytics", vi: "Phân tích cơ bản" },
    "pricing.feature.standardSupport": { en: "Standard support", vi: "Hỗ trợ tiêu chuẩn" },
    "pricing.feature.advAnalytics": { en: "Advanced analytics", vi: "Phân tích nâng cao" },
    "pricing.feature.prioritySupport": { en: "Priority support", vi: "Hỗ trợ ưu tiên" },
    "pricing.feature.apiAccess": { en: "API access", vi: "Truy cập API" },
    "pricing.feature.unlimitedClicks": { en: "Unlimited clicks", vi: "Không giới hạn click" },
    "pricing.feature.unlimitedTeam": { en: "Unlimited team members", vi: "Không giới hạn thành viên" },
    "pricing.feature.dedicatedSupport": { en: "24/7 dedicated support", vi: "Hỗ trợ 24/7" },
    "pricing.feature.customIntegrations": { en: "Custom integrations", vi: "Tích hợp tùy chỉnh" },
    "pricing.feature.sla": { en: "SLA guarantee", vi: "Cam kết SLA" },
    "pricing.feature.sso": { en: "SSO/SAML", vi: "Đăng nhập SSO/SAML" },
    "pricing.feature.teamMembers": { en: "{count} team members", vi: "{count} thành viên đội" },
    "pricing.feature.customDomainsCount": { en: "{count} custom domains", vi: "{count} tên miền riêng" },

    // CTA
    "cta.title": { en: "Ready to level up your links?", vi: "Sẵn sàng nâng cấp link của bạn?" },
    "cta.subtitle": { en: "Join thousands of teams using GoLink to create and manage short links. Get started for free today.", vi: "Tham gia cùng hàng nghìn đội nhóm sử dụng GoLink. Bắt đầu miễn phí ngay hôm nay." },
    "cta.viewPricing": { en: "View Pricing", vi: "Xem giá" },

    // Auth
    "auth.login": { en: "Login", vi: "Đăng nhập" },
    "auth.register": { en: "Register", vi: "Đăng ký" },
    "auth.welcomeBack": { en: "Welcome back", vi: "Chào mừng trở lại" },
    "auth.signInToContinue": { en: "Sign in to your account to continue", vi: "Đăng nhập vào tài khoản để tiếp tục" },
    "auth.createAccount": { en: "Create your account", vi: "Tạo tài khoản" },
    "auth.startCreating": { en: "Start creating short links for free", vi: "Bắt đầu tạo link ngắn miễn phí" },
    "auth.email": { en: "Email", vi: "Email" },
    "auth.password": { en: "Password", vi: "Mật khẩu" },
    "auth.confirmPassword": { en: "Confirm Password", vi: "Xác nhận mật khẩu" },
    "auth.fullName": { en: "Full Name", vi: "Họ và tên" },
    "auth.forgotPassword": { en: "Forgot password?", vi: "Quên mật khẩu?" },
    "auth.rememberMe": { en: "Remember me for 30 days", vi: "Ghi nhớ trong 30 ngày" },
    "auth.signIn": { en: "Sign In", vi: "Đăng nhập" },
    "auth.createAccountBtn": { en: "Create Account", vi: "Tạo tài khoản" },
    "auth.orContinueWith": { en: "or continue with", vi: "hoặc tiếp tục với" },
    "auth.agreeToTerms": { en: "I agree to the", vi: "Tôi đồng ý với" },
    "auth.termsOfService": { en: "Terms of Service", vi: "Điều khoản dịch vụ" },
    "auth.and": { en: "and", vi: "và" },
    "auth.privacyPolicy": { en: "Privacy Policy", vi: "Chính sách bảo mật" },
    "auth.minChars": { en: "Must be at least 8 characters", vi: "Tối thiểu 8 ký tự" },
    "auth.username": { en: "Username", vi: "Tên đăng nhập" },
    "auth.firstName": { en: "First Name", vi: "Tên" },
    "auth.lastName": { en: "Last Name", vi: "Họ" },
    "auth.gender": { en: "Gender", vi: "Giới tính" },
    "auth.genderMale": { en: "Male", vi: "Nam" },
    "auth.genderFemale": { en: "Female", vi: "Nữ" },
    "auth.genderOther": { en: "Other", vi: "Khác" },
    "auth.birthday": { en: "Birthday", vi: "Ngày sinh" },
    "auth.remember30days": { en: "Remember 30 days", vi: "Ghi nhớ 30 ngày" },
    "auth.selectGender": { en: "Select", vi: "Chọn" },
    "auth.enterUsername": { en: "Enter your username", vi: "Nhập tên đăng nhập" },
    "auth.chooseUsername": { en: "Choose a username", vi: "Chọn tên đăng nhập" },

    // Footer
    "footer.description": { en: "The modern URL shortener built for teams and businesses.", vi: "Công cụ rút gọn URL hiện đại cho đội nhóm và doanh nghiệp." },
    "footer.product": { en: "Product", vi: "Sản phẩm" },
    "footer.company": { en: "Company", vi: "Công ty" },
    "footer.legal": { en: "Legal", vi: "Pháp lý" },
    "footer.pricing": { en: "Pricing", vi: "Giá" },
    "footer.features": { en: "Features", vi: "Tính năng" },
    "footer.api": { en: "API", vi: "API" },
    "footer.about": { en: "About", vi: "Giới thiệu" },
    "footer.blog": { en: "Blog", vi: "Blog" },
    "footer.careers": { en: "Careers", vi: "Tuyển dụng" },
    "footer.privacy": { en: "Privacy", vi: "Riêng tư" },
    "footer.terms": { en: "Terms", vi: "Điều khoản" },
    "footer.rights": { en: "All rights reserved.", vi: "Bản quyền được bảo lưu." },

    // Theme
    "theme.light": { en: "Light", vi: "Sáng" },
    "theme.dark": { en: "Dark", vi: "Tối" },
    "theme.system": { en: "System", vi: "Hệ thống" },

    // Admin
    "admin.portal": { en: "Admin Portal", vi: "Cổng Admin" },
    "admin.access": { en: "Admin Access", vi: "Đăng nhập Admin" },
    "admin.signInCredentials": { en: "Sign in with your admin credentials", vi: "Đăng nhập với tài khoản admin của bạn" },
    "admin.accessPanel": { en: "Access Admin Panel", vi: "Vào trang Admin" },
    "admin.enterUsername": { en: "Enter your username", vi: "Nhập tên đăng nhập" },

    // Admin Dashboard
    "admin.dashboard": { en: "Admin Dashboard", vi: "Bảng điều khiển Admin" },
    "admin.overview": { en: "Overview of your platform's performance", vi: "Tổng quan hiệu suất nền tảng" },
    "admin.totalUsers": { en: "Total Users", vi: "Tổng người dùng" },
    "admin.totalTenants": { en: "Total Tenants", vi: "Tổng Tenant" },
    "admin.totalLinks": { en: "Total Links", vi: "Tổng Link" },
    "admin.totalClicks": { en: "Total Clicks", vi: "Tổng Click" },
    "admin.activeSubscriptions": { en: "Active Subscriptions", vi: "Đăng ký hoạt động" },
    "admin.monthlyRevenue": { en: "Monthly Revenue", vi: "Doanh thu tháng" },
    "admin.manageUsers": { en: "Manage Users", vi: "Quản lý người dùng" },
    "admin.viewEditUsers": { en: "View and edit users", vi: "Xem và chỉnh sửa người dùng" },
    "admin.manageTenants": { en: "Manage Tenants", vi: "Quản lý Tenant" },
    "admin.viewWorkspaces": { en: "View workspaces", vi: "Xem không gian làm việc" },
    "admin.editPlans": { en: "Edit Plans", vi: "Chỉnh sửa gói" },
    "admin.configurePricing": { en: "Configure pricing", vi: "Cấu hình giá" },
    "admin.viewAnalytics": { en: "View Analytics", vi: "Xem phân tích" },
    "admin.platformMetrics": { en: "Platform metrics", vi: "Chỉ số nền tảng" },

    // Admin Sidebar
    "admin.sidebar.dashboard": { en: "Dashboard", vi: "Bảng điều khiển" },
    "admin.sidebar.users": { en: "Users", vi: "Người dùng" },
    "admin.sidebar.plans": { en: "Plans", vi: "Gói dịch vụ" },
    "admin.sidebar.settings": { en: "Settings", vi: "Cài đặt" },
    "admin.sidebar.permissions": { en: "Permissions", vi: "Phân quyền" },
    "admin.sidebar.attributes": { en: "Attributes", vi: "Thuộc tính" },
    "admin.sidebar.profile": { en: "Profile", vi: "Hồ sơ" },
    "admin.sidebar.logout": { en: "Log out", vi: "Đăng xuất" },
    "admin.sidebar.account": { en: "Admin Account", vi: "Tài khoản Admin" },

    // Admin Permissions
    "admin.permissions.title": { en: "Permissions Management", vi: "Quản lý phân quyền" },
    "admin.permissions.subtitle": { en: "Manage roles and their permissions", vi: "Quản lý vai trò và quyền hạn" },
    "admin.permissions.roles": { en: "Roles", vi: "Vai trò" },
    "admin.permissions.addRole": { en: "Add Role", vi: "Thêm vai trò" },
    "admin.permissions.editRole": { en: "Edit Role", vi: "Sửa vai trò" },
    "admin.permissions.deleteRole": { en: "Delete Role", vi: "Xóa vai trò" },
    "admin.permissions.select": { en: "Select", vi: "Chọn" },
    "admin.permissions.edit": { en: "Edit", vi: "Sửa" },
    "admin.permissions.delete": { en: "Delete", vi: "Xóa" },
    "admin.permissions.permissionsFor": { en: "Permissions for:", vi: "Quyền cho:" },
    "admin.permissions.noRoleSelected": { en: "No role selected", vi: "Chưa chọn vai trò" },
    "admin.permissions.selectFromLeft": { en: "Please select a role from the list on the left", vi: "Vui lòng chọn vai trò từ danh sách bên trái" },
    "admin.permissions.view": { en: "View", vi: "Xem" },
    "admin.permissions.add": { en: "Add", vi: "Thêm" },
    "admin.permissions.save": { en: "Save Permissions", vi: "Lưu quyền" },

    // Admin Attributes
    "admin.attributes.title": { en: "Attributes", vi: "Thuộc tính" },
    "admin.attributes.subtitle": { en: "Manage user attribute definitions for custom fields.", vi: "Quản lý định nghĩa thuộc tính người dùng cho các trường tùy chỉnh." },
    "admin.attributes.addAttribute": { en: "Add Attribute", vi: "Thêm thuộc tính" },
    "admin.attributes.editAttribute": { en: "Edit Attribute", vi: "Sửa thuộc tính" },
    "admin.attributes.createAttribute": { en: "Create Attribute", vi: "Tạo thuộc tính" },
    "admin.attributes.key": { en: "Attribute Key", vi: "Khóa thuộc tính" },
    "admin.attributes.keyPlaceholder": { en: "e.g., department, employee_id", vi: "VD: department, employee_id" },
    "admin.attributes.keyHint": { en: "Use lowercase letters, numbers, and underscores only.", vi: "Chỉ sử dụng chữ thường, số và dấu gạch dưới." },
    "admin.attributes.dataType": { en: "Data Type", vi: "Kiểu dữ liệu" },
    "admin.attributes.description": { en: "Description", vi: "Mô tả" },
    "admin.attributes.descriptionPlaceholder": { en: "Describe the purpose of this attribute", vi: "Mô tả mục đích của thuộc tính này" },
    "admin.attributes.noAttributes": { en: "No attributes found", vi: "Không tìm thấy thuộc tính" },
    "admin.attributes.created": { en: "Created", vi: "Ngày tạo" },
    "admin.attributes.actions": { en: "Actions", vi: "Thao tác" },
    "admin.attributes.searchPlaceholder": { en: "Search attributes...", vi: "Tìm kiếm thuộc tính..." },
    "admin.attributes.cancel": { en: "Cancel", vi: "Hủy" },
    "admin.attributes.saveChanges": { en: "Save Changes", vi: "Lưu thay đổi" },
    "admin.attributes.defineNew": { en: "Define a new attribute for user profiles.", vi: "Định nghĩa thuộc tính mới cho hồ sơ người dùng." },

    // Profile
    "profile.title": { en: "Profile Information", vi: "Thông tin hồ sơ" },
    "profile.subtitle": { en: "Update your account profile details and settings.", vi: "Cập nhật thông tin hồ sơ và cài đặt tài khoản." },
    "profile.firstName": { en: "First Name", vi: "Tên" },
    "profile.lastName": { en: "Last Name", vi: "Họ" },
    "profile.gender": { en: "Gender", vi: "Giới tính" },
    "profile.birthday": { en: "Birthday", vi: "Ngày sinh" },
    "profile.username": { en: "Username", vi: "Tên đăng nhập" },
    "profile.usernameHint": { en: "Username cannot be changed.", vi: "Tên đăng nhập không thể thay đổi." },
    "profile.email": { en: "Email", vi: "Email" },
    "profile.picture": { en: "Profile Picture", vi: "Ảnh đại diện" },
    "profile.pictureHint": { en: "Click on the image to upload a new one.", vi: "Nhấn vào ảnh để tải lên ảnh mới." },
    "profile.randomize": { en: "Randomize Avatar", vi: "Tạo avatar ngẫu nhiên" },
    "profile.security": { en: "Security", vi: "Bảo mật" },
    "profile.securitySubtitle": { en: "Manage your password and security settings.", vi: "Quản lý mật khẩu và cài đặt bảo mật." },
    "profile.currentPassword": { en: "Current Password", vi: "Mật khẩu hiện tại" },
    "profile.currentPasswordPlaceholder": { en: "Enter current password to change it", vi: "Nhập mật khẩu hiện tại để thay đổi" },
    "profile.newPassword": { en: "New Password", vi: "Mật khẩu mới" },
    "profile.confirmPassword": { en: "Confirm Password", vi: "Xác nhận mật khẩu" },
    "profile.saveChanges": { en: "Save Changes", vi: "Lưu thay đổi" },
    "profile.notSpecified": { en: "Not Specified", vi: "Không xác định" },
    "profile.male": { en: "Male", vi: "Nam" },
    "profile.female": { en: "Female", vi: "Nữ" },
    "profile.other": { en: "Other", vi: "Khác" },

    // User Sidebar
    "user.sidebar.workspaces": { en: "Workspaces", vi: "Không gian làm việc" },
    "user.sidebar.overview": { en: "Overview", vi: "Tổng quan" },
    "user.sidebar.links": { en: "Links", vi: "Link" },
    "user.sidebar.settings": { en: "Settings", vi: "Cài đặt" },
    "user.sidebar.selectWorkspace": { en: "Select a workspace to see navigation", vi: "Chọn workspace để xem điều hướng" },
    "user.sidebar.account": { en: "My Account", vi: "Tài khoản" },
    "user.sidebar.allWorkspaces": { en: "All Workspaces", vi: "Tất cả Workspace" },
    "user.sidebar.profile": { en: "Profile", vi: "Hồ sơ" },
    "user.sidebar.accountSettings": { en: "Account Settings", vi: "Cài đặt tài khoản" },
    "user.sidebar.logout": { en: "Log out", vi: "Đăng xuất" },

    // User Profile Page
    "user.profile.title": { en: "Profile Settings", vi: "Cài đặt hồ sơ" },
    "user.profile.subtitle": { en: "Manage your personal information and security preferences.", vi: "Quản lý thông tin cá nhân và cài đặt bảo mật." },
    "user.profile.error": { en: "Error loading user profile. Please try logging in again.", vi: "Lỗi tải hồ sơ. Vui lòng đăng nhập lại." },

    // Workspace Page
    "workspace.manageLinks": { en: "Manage your short links and view analytics", vi: "Quản lý link ngắn và xem phân tích" },
    "workspace.createLink": { en: "Create Link", vi: "Tạo Link" },
    "workspace.createShortLink": { en: "Create Short Link", vi: "Tạo Link Ngắn" },
    "workspace.createShortLinkDesc": { en: "Enter a long URL to create a shortened version.", vi: "Nhập URL dài để tạo phiên bản rút gọn." },
    "workspace.destinationUrl": { en: "Destination URL", vi: "URL đích" },
    "workspace.customAlias": { en: "Custom Alias (optional)", vi: "Alias tùy chỉnh (tùy chọn)" },
    "workspace.totalLinks": { en: "Total Links", vi: "Tổng Link" },
    "workspace.totalClicks": { en: "Total Clicks", vi: "Tổng Click" },
    "workspace.activeLinks": { en: "Active Links", vi: "Link hoạt động" },
    "workspace.thisMonth": { en: "This Month", vi: "Tháng này" },
    "workspace.yourLinks": { en: "Your Links", vi: "Các Link của bạn" },
    "workspace.shortLink": { en: "Short Link", vi: "Link ngắn" },
    "workspace.destination": { en: "Destination", vi: "Đích" },
    "workspace.clicks": { en: "Clicks", vi: "Click" },
    "workspace.status": { en: "Status", vi: "Trạng thái" },
    "workspace.created": { en: "Created", vi: "Ngày tạo" },
    "workspace.active": { en: "Active", vi: "Hoạt động" },
    "workspace.inactive": { en: "Inactive", vi: "Không hoạt động" },
    "workspace.noLinks": { en: "No links yet", vi: "Chưa có link" },
    "workspace.noLinksDesc": { en: "Create your first short link to get started.", vi: "Tạo link ngắn đầu tiên để bắt đầu." },
    "workspace.edit": { en: "Edit", vi: "Sửa" },
    "workspace.activate": { en: "Activate", vi: "Kích hoạt" },
    "workspace.deactivate": { en: "Deactivate", vi: "Vô hiệu hóa" },
    "workspace.delete": { en: "Delete", vi: "Xóa" },
    "workspace.notFound": { en: "Workspace not found", vi: "Không tìm thấy workspace" },
    "workspace.notFoundDesc": { en: "The workspace you're looking for doesn't exist.", vi: "Workspace bạn đang tìm không tồn tại." },
    "workspace.domain": { en: "Domain", vi: "Miền" },
    "workspace.members": { en: "Members", vi: "Thành viên" },
    "workspace.addMember": { en: "Add Member", vi: "Thêm thành viên" },

    // Common
    "common.cancel": { en: "Cancel", vi: "Hủy" },
    "common.save": { en: "Save", vi: "Lưu" },
    "common.delete": { en: "Delete", vi: "Xóa" },
    "common.edit": { en: "Edit", vi: "Sửa" },
    "common.create": { en: "Create", vi: "Tạo" },
    "common.search": { en: "Search", vi: "Tìm kiếm" },
    "common.actions": { en: "Actions", vi: "Thao tác" },
};

interface LanguageContextType {
    language: Language;
    setLanguage: (lang: Language) => void;
    t: (key: string, params?: Record<string, string | number>) => string;
}

const LanguageContext = createContext<LanguageContextType | undefined>(undefined);

export function LanguageProvider({ children }: { children: ReactNode }) {
    const [language, setLanguage] = useState<Language>("en");

    useEffect(() => {
        const saved = localStorage.getItem("golink-language") as Language;
        if (saved && (saved === "en" || saved === "vi")) {
            setLanguage(saved);
        }
    }, []);

    const handleSetLanguage = (lang: Language) => {
        setLanguage(lang);
        localStorage.setItem("golink-language", lang);
    };

    const t = (key: string, params?: Record<string, string | number>): string => {
        const translation = translations[key];

        let text = key;
        if (translation) {
            text = translation[language] || translation.en || key;
        }

        // Interpolate params
        if (params) {
            Object.entries(params).forEach(([paramKey, paramValue]) => {
                text = text.replace(`{${paramKey}}`, String(paramValue));
            });
        }

        return text;
    };

    return (
        <LanguageContext.Provider value={{ language, setLanguage: handleSetLanguage, t }}>
            {children}
        </LanguageContext.Provider>
    );
}

export function useLanguage() {
    const context = useContext(LanguageContext);
    if (!context) {
        throw new Error("useLanguage must be used within a LanguageProvider");
    }
    return context;
}
