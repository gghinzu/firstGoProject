package entity

// User structure, ok for json usage
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Education string `json:"education"`
}
