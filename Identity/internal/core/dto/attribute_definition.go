package dto

// CreateAttributeDefinitionRequest represents request to create an attribute definition.
type CreateAttributeDefinitionRequest struct {
	Key         string `json:"key" validate:"required"`
	DataType    string `json:"data_type"`
	Description string `json:"description"`
}

// UpdateAttributeDefinitionRequest represents request to update an attribute definition.
type UpdateAttributeDefinitionRequest struct {
	ID          int     `json:"-" uri:"id"`
	Key         *string `json:"key"`
	DataType    *string `json:"data_type"`
	Description *string `json:"description"`
}

// GetAttributeDefinitionRequest represents request to get an attribute definition by ID.
type GetAttributeDefinitionRequest struct {
	ID int `uri:"id" validate:"required"`
}

// DeleteAttributeDefinitionRequest represents request to delete an attribute definition.
type DeleteAttributeDefinitionRequest struct {
	ID int `uri:"id" validate:"required"`
}

// AttributeDefinitionResponse represents attribute definition data in API response.
type AttributeDefinitionResponse struct {
	ID          int    `json:"id"`
	Key         string `json:"key"`
	DataType    string `json:"data_type"`
	Description string `json:"description"`
}