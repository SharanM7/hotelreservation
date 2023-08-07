package types

import "context"

type UserStore interface {
	GetUser(string) (*User, error)
	AddUser(context.Context) error
}
