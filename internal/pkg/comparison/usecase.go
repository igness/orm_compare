package comparison

// UseCase ...
type UseCase struct {
	repo Repo
}

// NewUseCase ...
func NewUseCase(repo Repo) *UseCase {
	return &UseCase{repo: repo}
}

// ListUsers ...
func (u *UseCase) ListUsers() ([]User, error) {
	return nil, nil
}

// GetUser ...
func (u *UseCase) GetUser(userID uint64) (*User, error) {
	user, err := u.repo.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return user, err
}

// CreateUser ...
func (u *UseCase) CreateUser(user User) (*User, error) {
	createdUser, err := u.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return createdUser, err
}

// UpdateUser ...
func (u *UseCase) UpdateUser(user User) (*User, error) {
	updatedUser, err := u.repo.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

// DeleteUser ...
func (u *UseCase) DeleteUser(user User) error {
	if err := u.repo.DeleteUser(user); err != nil {
		return err
	}
	return nil
}

// ListFoods ...
func (u *UseCase) ListFoods() ([]Food, error) {
	return nil, nil
}

// GetFood ...
func (u *UseCase) GetFood(foodID uint64) (*Food, error) {
	food, err := u.repo.GetFood(foodID)
	if err != nil {
		return nil, err
	}
	return food, nil
}

// CreateFood ...
func (u *UseCase) CreateFood(food Food) (*Food, error) {
	createdFood, err := u.repo.CreateFood(food)
	if err != nil {
		return nil, err
	}
	return createdFood, nil
}

// UpdateFood ...
func (u *UseCase) UpdateFood(food Food) (*Food, error) {
	updatedFood, err := u.repo.UpdateFood(food)
	if err != nil {
		return nil, err
	}
	return updatedFood, nil
}

// DeleteFood ...
func (u *UseCase) DeleteFood(food Food) error {
	if err := u.repo.DeleteFood(food); err != nil {
		return err
	}
	return nil
}

// ListOrders ...
func (u *UseCase) ListOrders() ([]Order, error) {
	return nil, nil
}

// GetOrder ...
func (u *UseCase) GetOrder(orderID uint64) (*Order, error) {
	order, err := u.repo.GetOrder(orderID)
	if err != nil {
		return nil, err
	}
	return order, nil
}

// CreateOrder ...
func (u *UseCase) CreateOrder(user User, food Food) (*Order, error) {
	order, err := u.repo.CreateOrder(user, food)
	if err != nil {
		return nil, err
	}
	return order, nil
}

// UpdateOrder ...
func (u *UseCase) UpdateOrder(order Order) (*Order, error) {
	updatedOrder, err := u.repo.UpdateOrder(order)
	if err != nil {
		return nil, err
	}
	return updatedOrder, nil
}

// DeleteOrder ...
func (u *UseCase) DeleteOrder(order Order) error {
	if err := u.repo.DeleteOrder(order); err != nil {
		return err
	}
	return nil
}
