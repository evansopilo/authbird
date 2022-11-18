package data

import "context"

type Models struct {
	User interface {
		Create(ctx context.Context, user *User) error

		GetByID(ctx context.Context, id string) (*User, error)

		Get(ctx context.Context, skip, limit int64) (*[]User, error)

		GetByEmail(ctx context.Context, email string) (*User, error)

		GetByPhone(ctx context.Context, phone string) (*User, error)

		Update(ctx context.Context, id string, user *User) error

		DeleteByID(ctx context.Context, id string) error
	}
}

type TokenStr struct {
	Value string `json:"value,omitempty"`
}

type AuthCred struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
