package services

import (
	"context"
	"errors"
	"fmt"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"streamChatServer/db"
	"streamChatServer/types"
	"streamChatServer/utils"
)

func GetAllUsers() ([]*types.User, error) {

	var users []*types.User

	client, err := db.GetMongoClient()

	if err != nil {
		return nil, err
	}

	coll := client.Database(db.Database).Collection(string(db.UsersCollection))

	curr, err := coll.Find(context.TODO(), bson.D{
		primitive.E{},
	})

	if err != nil {
		return nil, err
	}

	for curr.Next(context.TODO()) {
		var u types.User
		err := curr.Decode(&u)
		if err != nil {
			return nil, err
		}

		users = append(users, &u)
	}

	return users, nil
}

func CreateNewUser(user types.User) (*mongo.InsertOneResult, error) {
	hashedPass, err := utils.HashPass(user.Password)
	if err != nil {
		return nil, err
	}
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}
	newUser := &types.User{
		Id:       id,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Password: hashedPass,
	}
	client, err := db.GetMongoClient()
	if err != nil {
		return nil, err
	}
	coll := client.Database(db.Database).Collection(string(db.UsersCollection))
	result := coll.FindOne(context.TODO(), bson.D{primitive.E{
		Key:   "email",
		Value: &user.Email,
	}})
	searchResult := result.Decode(&newUser)
	if searchResult == nil {
		return nil, errors.New("username or email has already been taken")
	}
	if searchResult.Error() != "mongo: no documents in result" {
		return nil, errors.New("username or email has already been taken")
	}
	insertResult, err := coll.InsertOne(context.TODO(), newUser)
	if err != nil {
		return nil, err
	}
	return insertResult, nil
}

func LoginUser(user types.User) (*types.User, error) {
	fmt.Println("input user --> ", user)
	var existingUser types.User

	client, err := db.GetMongoClient()

	if err != nil {
		return nil, err
	}

	coll := client.Database(db.Database).Collection(string(db.UsersCollection))

	result := coll.FindOne(context.TODO(), bson.D{
		{"email", user.Email},
	})

	if result == nil {
		return nil, errors.New("no user with that email or username found")
	}
	fmt.Println("search results ---> ", result.Decode(&existingUser))
	err = result.Decode(&existingUser)
	if err != nil {
		return nil, err
	}

	matchedPass, err := utils.VerifyPass(user.Password, existingUser.Password)

	if err != nil {
		return nil, err
	}

	if !matchedPass {
		return nil, errors.New("incorrect email or password")
	}

	return &existingUser, nil
}
