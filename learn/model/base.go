package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type BaseModel struct {
	Id              *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` //主键id需要加 omitempty
	CreateTimestamp *time.Time          `json:"createTimestamp,omitempty" bson:"createTimestamp,omitempty" `
	CreateUsername  string              `json:"createUsername,omitempty" bson:"createUsername,omitempty"`
	UpdateTimestamp *time.Time          `json:"updateTimestamp,omitempty" bson:"updateTimestamp,omitempty"`
	UpdateUsername  string              `json:"updateUsername,omitempty" bson:"updateUsername,omitempty"`
	CompanyId       string              `json:"companyId,omitempty" bson:"companyId,omitempty"`
	CompanyName     string              `json:"companyName,omitempty" bson:"companyName,omitempty"`
}
