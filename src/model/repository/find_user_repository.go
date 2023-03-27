package repository

func (ur *userRepository) FindUserByEmail(email string,) (model.UserDomainInterface, *rest_err.RestErr ){
	logger.Info("Init find user by email repository.")
	collection_name := os.Getenv(MONGO_DB_USER)
	collection := ur.databaseConnection.Collection(collection_name)

	UserEntity := &entity.UserEntity{}

	filter := bson.D({ Key: "email", Value: email})

	err := collection.FindOne( context.Background(), filter).Decode(UserEntity)
	if err != nil {
		if err == mongo.ErrNoDocumments {
			errorMessage := fmt.Sprintf(
				"User not found with this email: %s", email)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
	}

	return nil, nil
}