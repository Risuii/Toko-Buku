package services

import "Toko/modules"

type Kontrak interface {
	CreateToko(*modules.Toko) error
	GetToko(string) (*modules.Toko, error)
	GetAll() ([]*modules.Toko, error)
	UpdateToko(*modules.Toko, string) error
	DeleteToko(string) error
	CreateBook(string, *modules.Buku) error
	GetBook(string, string) (*modules.Toko, error)
	UpdateBook(string, string, *modules.Buku) error
	DeleteBook(string, string) error
}
