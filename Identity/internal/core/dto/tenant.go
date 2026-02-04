package dto

// CreateTenantRequest represents request to create a tenant.
type CreateTenantRequest struct {
	Name string `json:"name" validate:"required,min=2,max=50"`
}

// UpdateTenantRequest represents request to update a tenant.
type UpdateTenantRequest struct {
	ID     int     `json:"-" uri:"id"`
	Name   *string `json:"name" validate:"omitempty,min=2,max=50"`
	PlanID *int    `json:"plan_id"`
}

// GetTenantRequest represents request to get a tenant by ID.
type GetTenantRequest struct {
	ID int `uri:"id" validate:"required"`
}

// DeleteTenantRequest represents request to delete a tenant.
type DeleteTenantRequest struct {
	ID int `uri:"id" validate:"required"`
}

// TenantResponse represents tenant data in API response.
type TenantResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	PlanID int    `json:"plan_id"`
}

// GetMyTenantsRequest represents request to get current user's tenants.
type GetMyTenantsRequest struct{}
