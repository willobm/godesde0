package modelos

import "time"

type User struct {
	Id       int
	Name     string
	CreateAt time.Time
	Status   bool
}

func (this *User) AddUser(id int, name string, createAt time.Time, status bool) {
	this.Id = id
	this.Name = name
	this.CreateAt = createAt
	this.Status = status
}
