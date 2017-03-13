package user

import (
	"errors"
)

type user struct {
	id    string
	name  string
	age   int
	addr1 string
	addr2 string
}

func New(id, name string, age int, addr1, addr2 string) (*user, error) {
	if len(id) == 0 {
		return nil, errors.New("id is nil")
	}

	if len(name) == 0 {
		return nil, errors.New("name is nil")
	}

	if age <= 0 {
		return nil, errors.New("age is nil")
	}

	if len(addr1) == 0 {
		return nil, errors.New("addr1 is nil")
	}

	if len(addr2) == 0 {
		return nil, errors.New("addr2 is nil")
	}

	user := user{id, name, age, addr1, addr2}

	return &user, nil
}

func (u *user) Id() string {
	return u.id
}

func (u *user) Name() string {
	return u.name
}

func (u *user) Age() int {
	return u.age
}

func (u *user) Addr1() string {
	return u.addr1
}

func (u *user) Addr2() string {
	return u.addr2
}

func (u *user) SetName(n string) {
	u.name = n
}

func (u *user) SetAge(a int) {
	u.age = a
}
