package user

type User struct {
	// bson _id - _ид - ключ индекс с автогенерацией   omitempty - может быть пустым
	// поле с json: "-" не будет выдаваться
	ID           string `json:"id" bson:"_id,omitempty"`
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"-" bson:"password"`
	Email        string `json:"email" bson:"email"`
}

type CreateUserDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
