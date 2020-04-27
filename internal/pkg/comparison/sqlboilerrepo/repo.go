package repo

import (
	"context"
	"database/sql"
	"errors"
	"orm-comparison/internal/pkg/comparison"
	"orm-comparison/internal/pkg/comparison/sqlboilerrepo/models"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/types"
)

// SqlboilerRepo ...
type SqlboilerRepo struct {
	db *sql.DB
}

// NewSqlboilerRepo ...
func NewSqlboilerRepo(db *sql.DB) *SqlboilerRepo {
	return &SqlboilerRepo{db: db}
}

// ListUsers ...
func (r *SqlboilerRepo) ListUsers() ([]comparison.User, error) {
	return nil, nil
}

// GetUser ...
func (r *SqlboilerRepo) GetUser(userID uint64) (*comparison.User, error) {
	u, err := models.FindUser(context.Background(), r.db, userID)
	if err != nil {
		return nil, err
	}

	return &comparison.User{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Name:      u.Name,
		Gender:    u.Gender.String,
		Email:     u.Email.String,
	}, nil
}

// CreateUser ...
func (r *SqlboilerRepo) CreateUser(user comparison.User) (*comparison.User, error) {
	dbUser := models.User{
		Name:   user.Name,
		Gender: null.String{String: user.Gender, Valid: true},
		Email:  null.String{String: user.Email, Valid: true},
	}

	// DB에 User를 넣는게 아니라 User를 DB에 넣는다는 인터페이스 구조가 ORM 다움 (오브젝트가 주인공)
	if err := dbUser.Insert(context.Background(), r.db, boil.Infer()); err != nil {
		return nil, err
	}

	return &comparison.User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		Gender:    dbUser.Gender.String,
		Email:     dbUser.Email.String,
	}, nil
}

// UpdateUser ...
func (r *SqlboilerRepo) UpdateUser(user comparison.User) (*comparison.User, error) {
	u, _ := models.FindUser(context.Background(), r.db, user.ID)

	u.Name = user.Name
	u.Gender = null.String{String: user.Gender, Valid: true}
	u.Email = null.String{String: user.Email, Valid: true}

	_, err := u.Update(context.Background(), r.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &comparison.User{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Name:      u.Name,
		Gender:    u.Gender.String,
		Email:     u.Email.String,
	}, nil
}

// DeleteUser ...
func (r *SqlboilerRepo) DeleteUser(user comparison.User) error {
	u, _ := models.FindUser(context.Background(), r.db, user.ID)

	_, err := u.Delete(context.Background(), r.db)
	if err != nil {
		return err
	}

	return nil
}

// ListFoods ...
func (r *SqlboilerRepo) ListFoods() ([]comparison.Food, error) {
	return nil, nil
}

// GetFood ...
func (r *SqlboilerRepo) GetFood(foodID uint64) (*comparison.Food, error) {
	f, err := models.FindFood(context.Background(), r.db, foodID)
	if err != nil {
		return nil, err
	}

	price, ok := f.Price.Big.Float64()
	if !ok {
		return nil, errors.New("Error")
	}
	return &comparison.Food{
		ID:        f.ID,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
		Name:      f.Name,
		Price:     price,
	}, nil
}

// CreateFood ...
func (r *SqlboilerRepo) CreateFood(food comparison.Food) (*comparison.Food, error) {
	price := types.Decimal{}
	price.SetFloat64(food.Price)
	dbFood := models.Food{
		Name:  food.Name,
		Price: price,
	}

	if err := dbFood.Insert(context.Background(), r.db, boil.Infer()); err != nil {
		return nil, err
	}

	p, ok := dbFood.Price.Big.Float64()
	if !ok {
		return nil, errors.New("error")
	}
	return &comparison.Food{
		ID:        dbFood.ID,
		CreatedAt: dbFood.CreatedAt,
		UpdatedAt: dbFood.UpdatedAt,
		Name:      dbFood.Name,
		Price:     p,
	}, nil
}

// UpdateFood ...
func (r *SqlboilerRepo) UpdateFood(food comparison.Food) (*comparison.Food, error) {
	f, _ := models.FindFood(context.Background(), r.db, food.ID)

	f.Name = food.Name
	price := types.Decimal{}
	price.SetFloat64(food.Price)
	f.Price = price

	_, err := f.Update(context.Background(), r.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	p, ok := f.Price.Big.Float64()
	if !ok {
		return nil, errors.New("error")
	}
	return &comparison.Food{
		ID:        f.ID,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
		Name:      f.Name,
		Price:     p,
	}, nil
}

// DeleteFood ...
func (r *SqlboilerRepo) DeleteFood(food comparison.Food) error {
	f, _ := models.FindFood(context.Background(), r.db, food.ID)

	_, err := f.Delete(context.Background(), r.db)
	if err != nil {
		return err
	}

	return nil
}

// ListOrders ...
func (r *SqlboilerRepo) ListOrders() ([]comparison.Order, error) {
	return nil, nil
}

// GetOrder ...
func (r *SqlboilerRepo) GetOrder(orderID uint64) (*comparison.Order, error) {
	o, err := models.FindOrder(context.Background(), r.db, orderID)
	if err != nil {
		return nil, err
	}

	amount, ok := o.PurchaseAmount.Float64()
	if !ok {
		return nil, errors.New("error")
	}
	return &comparison.Order{
		ID:             o.ID,
		CreatedAt:      o.CreatedAt,
		UpdatedAt:      o.UpdatedAt,
		UserID:         o.UserID,
		FoodID:         o.FoodID,
		PurchaseCount:  o.PurchaseCount,
		PurchaseAmount: amount,
	}, nil
}

// CreateOrder ...
func (r *SqlboilerRepo) CreateOrder(comparison.User, comparison.Food) (*comparison.Order, error) {
	return nil, nil
}

// UpdateOrder ...
func (r *SqlboilerRepo) UpdateOrder(order comparison.Order) (*comparison.Order, error) {
	return nil, nil
}

// DeleteOrder ...
func (r *SqlboilerRepo) DeleteOrder(order comparison.Order) error {
	return nil
}
