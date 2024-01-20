package main

type DAO interface {
	CreateUser(name, email string) error
	UpdateUser(id, name, email string) error
}
