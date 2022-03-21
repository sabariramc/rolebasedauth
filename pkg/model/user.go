package model

import (
	"time"

	"sabariram.com/goserverbase/db/mongo"
	"sabariram.com/rolebasedauth/pkg/constants"
)

type ActiveLogin struct {
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	LoginId     string    `json:"loginId" bson:"loginId"`
	Application string    `json:"application" bson:"application"`
	Platform    string    `json:"platform" bson:"platform"`
}

type LoginLog struct {
	ActiveLogin
	IsActive    bool                   `json:"isActive" bson:"isActive"`
	LoggedOutAt time.Time              `json:"loggedOutAt" bson:"loggedOutAt"`
	LogoutType  constants.LogoutType   `json:"logoutType" bson:"logoutType"`
	Metadata    map[string]interface{} `json:"metadata" bson:"metadata"`
}

type User struct {
	mongo.BaseMongoModel
	TenantId             string          `json:"tenantId" bson:"tenantId"`
	TenantUserIdentifier string          `json:"tenantUserIdentifier" bson:"tenantUserIdentifier"`
	UserLoginId          []string        `json:"userLoginId" bson:"userLoginId"`
	UserId               string          `json:"userId" bson:"userId"`
	Roles                []string        `json:"roles" bson:"roles"`
	Claims               []string        `json:"claims" bson:"claims"`
	MaxParallelLogin     int             `json:"maxParallelLogin" bson:"maxParallelLogin"`
	IsActive             bool            `json:"isActive" bson:"isActive"`
	Auth                 *Authentication `json:"authentication" bson:"authentication"`
	LoginLog             []*LoginLog     `json:"loginLog" bson:"loginLog"`
	ActiveLogin          []*ActiveLogin  `json:"activeLogin" bson:"activeLogin"`
}
