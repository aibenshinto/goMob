package models

type MobilePhone struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Image string  `json:"image"`
	Specs string  `json:"specs"`
	Price float64 `json:"price"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
