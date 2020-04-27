package repo

import (
	"orm-comparison/internal/pkg/comparison"

	"github.com/jinzhu/gorm"
)

// GormRepo ...
type GormRepo struct {
	db *gorm.DB
}

// NewGormRepo ...
func NewGormRepo(db *gorm.DB) *GormRepo {
	return &GormRepo{db: db}
}

// ListUsers ...
func (r *GormRepo) ListUsers() ([]comparison.User, error) {
	return nil, nil
}

// GetUser ...
func (r *GormRepo) GetUser(userID uint64) (*comparison.User, error) {
	var u GormUser
	if err := r.db.First(&u, userID).Error; err != nil {
		return nil, err
	}

	return &comparison.User{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Name:      u.Name,
		Gender:    u.Gender,
		Email:     u.Email,
	}, nil
}

// CreateUser ...
func (r *GormRepo) CreateUser(user comparison.User) (*comparison.User, error) {
	dbUser := GormUser{
		Name:   user.Name,
		Gender: user.Gender,
		Email:  user.Email,
	}

	// 여기서 "db"가 "User"를 생성한다는 느낌이 들어서 "db"가 주인공인 느낌 @Kyle
	if err := r.db.Create(&dbUser).Error; err != nil {
		return nil, err
	}

	return &comparison.User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		Gender:    dbUser.Gender,
		Email:     dbUser.Email,
	}, nil
}

// UpdateUser ...
func (r *GormRepo) UpdateUser(user comparison.User) (*comparison.User, error) {
	var u GormUser
	// 여기서 u가 데이터베이스에서 리턴된 값이라는 것을 직관적으로 못 느끼겠음 @Kyle
	// "db"의 "First" behavior를 사용해서 User 객체를 다룬다는 느낌이 없음 @Kyle
	if err := r.db.First(&u, user.ID).Error; err != nil {
		return nil, err
	}

	u.Name = user.Name
	u.Gender = user.Gender
	u.Email = user.Email

	// 여기서도 "User"를 다룬다는 느낌보다 "db"를 다룬다는 느낌이 먼저 듬 @Kyle
	if err := r.db.Save(&u).Error; err != nil {
		return nil, err
	}

	return &comparison.User{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Name:      u.Name,
		Gender:    u.Gender,
		Email:     u.Email,
	}, nil
}

// DeleteUser ...
func (r *GormRepo) DeleteUser(user comparison.User) error {
	u := GormUser{
		ID: user.ID,
	}
	if err := r.db.Delete(&u).Error; err != nil {
		return err
	}
	return nil
}

// ListFoods ...
func (r *GormRepo) ListFoods() ([]comparison.Food, error) {
	return nil, nil
}

// GetFood ...
func (r *GormRepo) GetFood(foodID uint64) (*comparison.Food, error) {
	var f GormFood
	if err := r.db.First(&f, foodID).Error; err != nil {
		return nil, err
	}

	return &comparison.Food{
		ID:        f.ID,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
		Name:      f.Name,
		Price:     f.Price,
	}, nil
}

// CreateFood ...
func (r *GormRepo) CreateFood(food comparison.Food) (*comparison.Food, error) {
	dbFood := GormFood{
		Name:  food.Name,
		Price: food.Price,
	}

	if err := r.db.Create(&dbFood).Error; err != nil {
		return nil, err
	}

	return &comparison.Food{
		ID:        dbFood.ID,
		CreatedAt: dbFood.CreatedAt,
		UpdatedAt: dbFood.UpdatedAt,
		Name:      dbFood.Name,
		Price:     dbFood.Price,
	}, nil
}

// UpdateFood ...
func (r *GormRepo) UpdateFood(food comparison.Food) (*comparison.Food, error) {
	var f GormFood

	if err := r.db.First(&f, food.ID).Error; err != nil {
		return nil, err
	}

	f.Name = food.Name
	f.Price = food.Price

	if err := r.db.Save(&f).Error; err != nil {
		return nil, err
	}

	return &comparison.Food{
		ID:        f.ID,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
		Name:      f.Name,
		Price:     f.Price,
	}, nil
}

// DeleteFood ...
func (r *GormRepo) DeleteFood(food comparison.Food) error {
	f := GormFood{
		ID: food.ID,
	}
	if err := r.db.Delete(&f).Error; err != nil {
		return err
	}
	return nil
}

// ListOrders ...
func (r *GormRepo) ListOrders() ([]comparison.Order, error) {
	return nil, nil
}

// GetOrder ...
func (r *GormRepo) GetOrder(orderID uint64) (*comparison.Order, error) {
	var o GormOrder
	if err := r.db.First(&o, orderID).Error; err != nil {
		return nil, err
	}

	return &comparison.Order{
		ID:             o.ID,
		CreatedAt:      o.CreatedAt,
		UpdatedAt:      o.UpdatedAt,
		UserID:         o.UserID,
		FoodID:         o.FoodID,
		PurchaseCount:  o.PurchaseCount,
		PurchaseAmount: o.PurchaseAmount,
	}, nil
}

// CreateOrder ...
func (r *GormRepo) CreateOrder(comparison.User, comparison.Food) (*comparison.Order, error) {
	return nil, nil
}

// UpdateOrder ...
func (r *GormRepo) UpdateOrder(order comparison.Order) (*comparison.Order, error) {
	return nil, nil
}

// DeleteOrder ...
func (r *GormRepo) DeleteOrder(order comparison.Order) error {
	return nil
}
