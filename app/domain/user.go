package domain

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type UserRepository interface {
	Save(user User) error
	GetByID(id int) (*User, error)
	GetByEmail(email string) (*User, error)
	Delete(id int) error
}
