package entity
import "go.mongodb.org/mongo-driver/bson/primitive"
type UserEntity struct {
	ID         primitive.ObjectID  `json:"_id,omiempty"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Name     string `bson:"name"`
	Age      int8   `bson:"age"`
}
