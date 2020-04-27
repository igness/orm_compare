package repo

import "time"

// GormUser ...
type GormUser struct {
	ID        uint64    `gorm:"type:BIGINT;AUTO_INCREMENT"`
	CreatedAt time.Time `gorm:"type:DATETIME;NOT NULL"`
	UpdatedAt time.Time `gorm:"type:DATETIME;NOT NULL"`
	Name      string    `gorm:"type:VARCHAR(255);NOT NULL"`
	Gender    string    `gorm:"type:VARCHAR(255);NULL"`
	Email     string    `gorm:"type:VARCHAR(255);NULL"`
}

// GormFood ...
type GormFood struct {
	ID        uint64    `gorm:"type:BIGINT;AUTO_INCREMENT"`
	CreatedAt time.Time `gorm:"type:DATETIME;NOT NULL"`
	UpdatedAt time.Time `gorm:"type:DATETIME;NOT NULL"`
	Name      string    `gorm:"type:VARCHAR(255);NOT NULL"`
	Price     float64   `gorm:"type:NUMERIC(27,9);NOT NULL"`
}

// GormOrder ...
type GormOrder struct {
	ID             uint64    `gorm:"type:BIGINT;AUTO_INCREMENT"`
	CreatedAt      time.Time `gorm:"type:DATETIME;NOT NULL"`
	UpdatedAt      time.Time `gorm:"type:DATETIME;NOT NULL"`
	UserID         uint64    `gorm:"type:INTEGER UNSIGNED;NOT NULL"`
	FoodID         uint64    `gorm:"type:INTEGER UNSIGNED;NOT NULL"`
	PurchaseCount  int       `gorm:"type:INTEGER;NOT NULL"`
	PurchaseAmount float64   `gorm:"type:NUMERIC(27,9);NOT NULL"`
}
