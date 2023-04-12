package models

type Client struct {
	SureName string
	Name     string
	LastName string
	Age      int
	Address  string
	Phone    string
	Card     []Card
}
