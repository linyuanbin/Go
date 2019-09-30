package models

type User struct {
	id    int    `db:"code" type:"int"`
	name  string `db:"name" type:"string"`
	age   int    `db:"age" type:"int"`
	addr  string `db:"addr" type:"string"`
	email string `db:"email" type:"string"`
}
