// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Mutation struct {
}

type NewCategory struct {
	Name string `json:"name"`
}

type NewDelivery struct {
	Cep      string `json:"CEP"`
	Address  string `json:"address"`
	Number   string `json:"number"`
	Country  string `json:"country"`
	City     string `json:"city"`
	District string `json:"district"`
}

type NewItems struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type NewOrder struct {
	Items    []*NewItems  `json:"items"`
	Delivery *NewDelivery `json:"delivery"`
	UserID   string       `json:"userId"`
}

type NewProduct struct {
	Code        *string `json:"code,omitempty"`
	Name        string  `json:"name"`
	ImageURL    string  `json:"imageUrl"`
	Price       int     `json:"price"`
	Description *string `json:"description,omitempty"`
	Active      bool    `json:"active"`
	CategoryID  string  `json:"categoryId"`
}

type Order struct {
	PaymentURL string `json:"paymentUrl"`
}

type Product struct {
	ID          string    `json:"id"`
	Code        *string   `json:"code,omitempty"`
	Name        string    `json:"name"`
	ImageURL    string    `json:"imageUrl"`
	Price       int       `json:"price"`
	Description *string   `json:"description,omitempty"`
	Active      bool      `json:"active"`
	CategoryID  string    `json:"categoryId"`
	Category    *Category `json:"Category,omitempty"`
}

type Query struct {
}

type RetrieveByID struct {
	ID string `json:"ID"`
}
