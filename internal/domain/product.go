package domain

type Product struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
