package models

import (
	"time"

	"github.com/upper/db/v4"
)

type User struct {
	ID        int `db:"id,omitempty"`
	Name      string `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UserRequest struct {
	Name string `json:"name"`
}

type UsersModel struct {
	db db.Session
}

func (m UsersModel) Table() string {
	return "users"
}

func (m UsersModel) GetAll() (users []User, err error) {
	err = m.db.Collection(m.Table()).Find().All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (m UsersModel) Get() (user *User, err error) {
	err = m.db.Collection(m.Table()).Find().All(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m UsersModel) Create(u *User) (err error) {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	
	col := m.db.Collection(m.Table())
	_, err = col.Insert(u)
	if err != nil {
		return err
	}

	return nil
}

func (m UsersModel) Update(u *User) (*User, error) {
	u.UpdatedAt = time.Now()
	err := m.db.Collection(m.Table()).Find(db.Cond{"id": u.ID}).Update(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (m UsersModel) Delete(id int) (err error) {
	err = m.db.Collection(m.Table()).Find(db.Cond{"id": id}).Delete()
	if err != nil {
		return err
	}

	return nil
}