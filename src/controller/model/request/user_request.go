package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Passowrd string `json:"password" binding:"required,min=6,containsany=!@#$%*&"`
	Name     string `json:"name" binding:"required,min=4,max=50"`
	Age      int8   `json:"age" binding:"required,numeric,min=1,max=140"`
}

type UserUpdateRequest struct {

	Email    string `json:"email" binding:"omitempty,email"`
	Passowrd string `json:"password" binding:"omitempty,min=6,containsany=!@#$%*&"`
	Name     string `json:"name" binding:"omitempty,min=4,max=50"`
	Age      int8   `json:"age" binding:"omitempty,numeric,min=1,max=140"`
}


type UserLogin struct {

	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*&"`
}
