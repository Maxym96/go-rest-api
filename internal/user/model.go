package user

type User struct {
	ID           string `json:"id" bson:"_id,omitempty"`
	Email        string `json:"email" bson:"email"`
	UserName     string `json:"username" bson:"username"`
	PasswordHash string `json:"-" bson:"password"`
}

// Pattern DTO - Data Transfer Object

type CreateUserDTO struct {
	Email    string `json:"email"`
	UserName string `json:"username"`
	Password string `json:"-"`
}
