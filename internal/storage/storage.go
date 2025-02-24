package storage

import ()

type Cat struct {
	Experience float32
	Salaty float64
	Name string
	Bread string
}

type Target struct {}

type Mision struct {
	Targets []Target
}

type Storage struct {
	
}

func (s *Storage) Conn() {

}

func (s *Storage) Close() {

}
 
func (s *Storage) AddCat(cat Cat) {

}

func (s *Storage) ListCat() []Cat {
	return nil
}

func (s *Storage) UpdateSalary() {

}

func (s *Storage) GetCat() {}

func (s *Storage) DeleteCat() {
	
}