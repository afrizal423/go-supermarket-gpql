// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Customer struct {
	ID string `json:"id"`
	// Nama customer dengan tipe data string
	CustomerName string      `json:"customer_name"`
	Segment      string      `json:"segment"`
	Age          string      `json:"age"`
	Country      string      `json:"country"`
	City         string      `json:"city"`
	State        string      `json:"state"`
	PostalCode   string      `json:"postal_code"`
	Region       string      `json:"region"`
	Related      []*Customer `json:"related"`
}

type NewCustomer struct {
	ID           string `json:"id"`
	CustomerName string `json:"customer_name"`
	Segment      string `json:"segment"`
	Age          string `json:"age"`
	Country      string `json:"country"`
	City         string `json:"city"`
	State        string `json:"state"`
	PostalCode   string `json:"postal_code"`
	Region       string `json:"region"`
}
