package comparison

// Repo ...
type Repo interface {
	ListUsers() ([]User, error)
	GetUser(userID uint64) (*User, error)
	CreateUser(User) (*User, error)
	UpdateUser(User) (*User, error)
	DeleteUser(User) error

	ListFoods() ([]Food, error)
	GetFood(foodID uint64) (*Food, error)
	CreateFood(Food) (*Food, error)
	UpdateFood(Food) (*Food, error)
	DeleteFood(Food) error

	ListOrders() ([]Order, error)
	GetOrder(orderID uint64) (*Order, error)
	CreateOrder(User, Food) (*Order, error)
	UpdateOrder(Order) (*Order, error)
	DeleteOrder(Order) error
}
