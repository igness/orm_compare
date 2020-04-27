package repo

import "time"

// SqlxUser ...
type SqlxUser struct {
	ID        uint64
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Name      string
	Gender    string
	Email     string
}

// SqlxFood ...
type SqlxFood struct {
	ID        uint64
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Name      string
	Price     float64
}

// SqlxOrder ...
type SqlxOrder struct {
	ID             uint64
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
	UserID         uint64    `db:"user_id"`
	FoodID         uint64    `db:"food_id"`
	PurchaseCount  int       `db:"purchase_count"`
	PurchaseAmount float64   `db:"purchase_amount"`
}
