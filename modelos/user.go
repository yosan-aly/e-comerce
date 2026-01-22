package modelos

import "time"

type User struct {
	UserId        int
	UserName      string
	CreatedAt time.Time
	Status    bool
}

func (u *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	u.UserId = id
	u.UserName = name
	u.CreatedAt = createdAt
	u.Status = status
}
