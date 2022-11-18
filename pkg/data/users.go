package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID            string    `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName     string    `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName      string    `json:"last_name,omitempty" bson:"last_name,omitempty"`
	DisplayName   string    `json:"display_name,omitempty" bson:"display_name,omitempty"`
	Phone         string    `json:"phone,omitempty" bson:"phone,omitempty"`
	Email         string    `json:"email,omitempty" bson:"email,omitempty"`
	EmailVerified string    `json:"email_verified,omitempty" bson:"email_verified,omitempty"`
	Password      string    `json:"password,omitempty" bson:"password,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type UserModel struct {
	client *mongo.Client
}

func NewUserModel(client *mongo.Client) *UserModel {
	return &UserModel{client: client}
}

func (u UserModel) Create(ctx context.Context, user *User) error {

	coll := u.client.Database("authbird").Collection("users")

	result, err := coll.InsertOne(ctx, &user)
	if err != nil {
		return err
	}

	if result.InsertedID.(string) != user.ID {
		return ErrCreateDocument
	}

	return nil
}

func (u UserModel) GetByID(ctx context.Context, id string) (*User, error) {

	coll := u.client.Database("authbird").Collection("users")

	var user User

	if err := coll.FindOne(ctx, bson.M{"_id": id}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u UserModel) Get(ctx context.Context, skip, limit int64) (*[]User, error) {

	opts := options.Find().SetSkip(skip).SetLimit(limit)

	coll := u.client.Database("authbird").Collection("users")

	filterCursor, err := coll.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}

	var users []User

	if err := filterCursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return &users, nil
}

func (u UserModel) GetByEmail(ctx context.Context, email string) (*User, error) {

	coll := u.client.Database("authbird").Collection("users")

	var user User

	if err := coll.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u UserModel) GetByPhone(ctx context.Context, phone string) (*User, error) {

	coll := u.client.Database("authbird").Collection("users")

	var user User

	if err := coll.FindOne(ctx, bson.M{"phone": phone}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u UserModel) Update(ctx context.Context, id string, user *User) error {

	coll := u.client.Database("authbird").Collection("users")

	result, err := coll.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.D{{Key: "$set", Value: &user}})
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return ErrNoDocument
	}

	return nil
}

func (u UserModel) DeleteByID(ctx context.Context, id string) error {

	coll := u.client.Database("authbird").Collection("users")

	result, err := coll.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return ErrNoDocument
	}

	return nil
}
