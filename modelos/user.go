package modelos

import (
	"time"
)

type User struct {
	Id     int
	Nome   string
	Email  string
	Status bool
	DataCreated time.Time
}

func (usuario *User) AddUser(id int, nome string, email string, status bool, dateCreated time.Time) {
	usuario.Id = id
	usuario.Nome = nome
	usuario.Email = email
	usuario.Status = status
	usuario.DataCreated = dateCreated
}