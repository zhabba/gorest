package main

type IStorage interface {
	Create(interface{}) (interface{}, error)
	Read(interface{}) (interface{}, error)
	Update(interface{}) (interface{}, error)
	Delete(interface{}) (interface{}, error)
}
