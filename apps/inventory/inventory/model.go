package inventory

type Inventory struct {
	ID          int    `json:"id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
}
