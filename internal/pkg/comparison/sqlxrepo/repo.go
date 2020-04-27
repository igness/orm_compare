package repo

import (
	"orm-comparison/internal/pkg/comparison"
	"strconv"

	"github.com/jmoiron/sqlx"
)

// SqlxRepo ...
type SqlxRepo struct {
	db *sqlx.DB
}

// NewSqlxRepo ...
func NewSqlxRepo(db *sqlx.DB) *SqlxRepo {
	return &SqlxRepo{db: db}
}

// ListUsers ...
func (r *SqlxRepo) ListUsers() ([]comparison.User, error) {
	return nil, nil
}

// GetUser ...
func (r *SqlxRepo) GetUser(userID uint64) (*comparison.User, error) {
	u := SqlxUser{}
	if err := r.db.Get(&u, "SELECT * FROM user WHERE id=$1", strconv.FormatUint(userID, 10)); err != nil {
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
func (r *SqlxRepo) CreateUser(user comparison.User) (*comparison.User, error) {
	_, err := r.db.NamedExec(`INSERT INTO user (name,gender,email) VALUES (:name,:gender,:email)`,
		map[string]interface{}{
			"name":   user.Name,
			"gender": user.Gender,
			"email":  user.Email,
		})
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser ...
func (r *SqlxRepo) UpdateUser(user comparison.User) (*comparison.User, error) {
	var u SqlxUser
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
func (r *SqlxRepo) DeleteUser(user comparison.User) error {
	u := SqlxUser{
		ID: user.ID,
	}
	if err := r.db.Delete(&u).Error; err != nil {
		return err
	}
	return nil
}

// ListFoods ...
func (r *SqlxRepo) ListFoods() ([]comparison.Food, error) {
	return nil, nil
}

// GetFood ...
func (r *SqlxRepo) GetFood(foodID uint64) (*comparison.Food, error) {
	var f SqlxFood
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
func (r *SqlxRepo) CreateFood(food comparison.Food) (*comparison.Food, error) {
	dbFood := SqlxFood{
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
func (r *SqlxRepo) UpdateFood(food comparison.Food) (*comparison.Food, error) {
	var f SqlxFood

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
func (r *SqlxRepo) DeleteFood(food comparison.Food) error {
	f := SqlxFood{
		ID: food.ID,
	}
	if err := r.db.Delete(&f).Error; err != nil {
		return err
	}
	return nil
}

// ListOrders ...
func (r *SqlxRepo) ListOrders() ([]comparison.Order, error) {
	return nil, nil
}

// GetOrder ...
func (r *SqlxRepo) GetOrder(orderID uint64) (*comparison.Order, error) {
	var o SqlxOrder
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
func (r *SqlxRepo) CreateOrder(comparison.User, comparison.Food) (*comparison.Order, error) {
	return nil, nil
}

// UpdateOrder ...
func (r *SqlxRepo) UpdateOrder(order comparison.Order) (*comparison.Order, error) {
	return nil, nil
}

// DeleteOrder ...
func (r *SqlxRepo) DeleteOrder(order comparison.Order) error {
	return nil
}
