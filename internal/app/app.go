package app

type Storage interface {
	AddCat()
	DeleteCat()
	UpdateSalary()
	ListCat()
	GetCat()
}