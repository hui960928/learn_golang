package yk_base

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"learn_golang/learn/model"
	"learn_golang/learn/resource"
	"time"
)

var (
	userCollection = resource.GetMongoClient().Database("yk_base").Collection("User")
)

type User struct {
	Accessible          bool      `json:"accessible" bson:"accessible"`
	ValidateDate        time.Time `json:"validateDate,omitempty" bson:"validateDate,omitempty" `
	PermissionTimestamp int64     `json:"permissionTimestamp" bson:"permissionTimestamp"`
	model.BaseModel     `bson:",inline"`
}

func FindUserById(obj *User, id string) (User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err == nil {
		b := bson.D{{"_id", objID}}
		err = userCollection.Find(context.TODO(), b).One(&obj)
	}
	return *obj, err
}
