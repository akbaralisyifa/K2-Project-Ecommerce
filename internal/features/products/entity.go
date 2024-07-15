package products

type Product struct {
	ID          uint
	Name        string
	Category    string
	Description string
	Price       int
	Stock       int
	ImageUrl    string
}

type Handler interface {
}

type Query interface {
	AddProduct(newProducts Product) error
	GetProduct(ID uint) (Product, error)
	UpdateProduct(ID uint, updateProduct Product) error
	DeleteProduct(ID uint) error
}

type Service interface {
	AddProduct(newProduct Product) error
}

type AddProductValidation struct {
	Name  string `validate:"required"`
	Price int    `validate:"required, number"`
	Stock int    `validate:"required, number"`
}

type LoginValidation struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}
