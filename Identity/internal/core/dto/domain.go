package dto

// CreateDomainRequest represents request to create a custom domain.
type CreateDomainRequest struct {
	Domain   string `json:"domain" validate:"required,min=4,max=100,hostname"`
	TenantID int    `json:"tenant_id" validate:"required"`
}

// UpdateDomainRequest represents request to update a domain.
type UpdateDomainRequest struct {
	ID         int     `json:"-" uri:"id"`
	Domain     *string `json:"domain" validate:"omitempty,min=4,max=100,hostname"`
	IsVerified *bool   `json:"is_verified"`
}

// GetDomainRequest represents request to get a domain by ID.
type GetDomainRequest struct {
	ID int `uri:"id" validate:"required"`
}

// DeleteDomainRequest represents request to delete a domain.
type DeleteDomainRequest struct {
	ID int `uri:"id" validate:"required"`
}

// DomainResponse represents domain data in API response.
type DomainResponse struct {
	ID         int    `json:"id"`
	Domain     string `json:"domain"`
	IsVerified bool   `json:"is_verified"`
	TenantID   int    `json:"tenant_id"`
}
