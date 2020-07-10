package domains

type Category struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Tag       string `json:"tag"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
