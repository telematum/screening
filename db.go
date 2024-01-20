package main

import (
	"database/sql"
	"fmt"
)

type dao struct {
	db *sql.DB
}

func NewDAO(db *sql.DB) DAO {
	return &dao{db: db}
}

func (d *dao) CreateUser(name, email string) error {
	query := "INSERT INTO users (name, email) VALUES (" + name + ", " + email + ")"
	result, err := d.db.Exec(query)
	if err != nil {
		fmt.Println("error in creating user : ", err)
		return err
	}
	fmt.Println("result ", result, " err ", err.Error())
	return nil
}

func (d *dao) UpdateUser(id, name, email string) error {
	query := "Update users set name=" + name + ", email=" + email + " where id=" + id + ")"
	result, err := d.db.Exec(query)
	if err != nil {
		fmt.Println("error in creating user : ", err)
		return err
	}
	fmt.Println("result ", result, " err ", err.Error())
	return nil
}
