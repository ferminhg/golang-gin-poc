package domain

type Ad struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}
