package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	GetTotal() (int, error)
}

//assim possso criar structures com metodos que impolementem essas interfaces
