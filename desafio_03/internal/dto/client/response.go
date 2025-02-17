package client

type ClientResponse struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	CPF       string  `json:"cpf"`
	Income    float64 `json:"income"`
	BirthDate string  `json:"birth_date"`
	Children  int     `json:"children"`
}

type PaginatedResponse struct {
	Content      []ClientResponse `json:"content"`
	TotalPages   int              `json:"total_pages"`
	TotalItems   int64            `json:"total_items"`
	CurrentPage  int              `json:"current_page"`
	ItemsPerPage int              `json:"items_per_page"`
}
