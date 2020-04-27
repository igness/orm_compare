package comparison

import "time"

// User ...
type User struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Gender    string
	Email     string
}

// Food ...
type Food struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Price     float64
}

// Order ...
type Order struct {
	ID             uint64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	UserID         uint64
	FoodID         uint64
	PurchaseCount  int
	PurchaseAmount float64
}
