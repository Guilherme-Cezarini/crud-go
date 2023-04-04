package response

type UserResponse struct {
<<<<<<< Updated upstream
	Id    int    `json:"id"`
=======
	ID    string `json:"id"`
>>>>>>> Stashed changes
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int8   `json:"age"`
}
