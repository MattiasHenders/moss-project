package users

type User struct {
	ID             *string `json:"id" db:"id"`
	FirstName      *string `json:"firstName" db:"first_name"`
	LastName       *string `json:"lastName" db:"last_name"`
	Email          *string `json:"email" db:"email"`
	Country        *string `json:"country" db:"country"`
	Sex            *string `json:"sex" db:"sex"`
	HashedPassword *string `json:"hashedPassword" db:"hashed_password"`
	UserType       *string `json:"userType" db:"user_type"`
}
