package models

type Product struct {
    ID      int    `json:"id"`
    Model   string `json:"model"`
    Company string `json:"company"`
    Price   int    `json:"price"`
}
