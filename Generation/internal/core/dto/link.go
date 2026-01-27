package dto

type CreateLinkRequest struct {
	OriginalURL string `json:"original_url"`
}

type LinkResponse struct {
	ShortLink string `json:"short_link"`
}

type DeleteLinkRequest struct {
	ID string `json:"id"`
}
