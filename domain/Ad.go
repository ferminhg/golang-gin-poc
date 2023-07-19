package domain

type Ad struct {
	Title       string  `json:"title" binding:"min=2,max=10"`
	Description string  `json:"description" binding:"max=50"`
	Price       float32 `json:"price"`
	Author      Person  `json:"author"`
}

type Person struct {
	FirstName  string `json:"firstName" binding:"required" validate:"is-cool"`
	SecondName string `json:"secondName" binding:"required"`
	Email      string `json:"email" validate:"required,email"`
	Age        int8   `json:"age" binding:"gte=0,lte=130"`
}
