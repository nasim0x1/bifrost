package types

import "time"

type User struct {
	ID        int       `db:"id" json:"id"`
	FirstName string    `db:"firstName" json:"first_name"`
	LastName  string    `db:"lastName" json:"last_name"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	CreatedAt time.Time `db:"createdAt" json:"created_at"`
}

