package go_micro_srv_user

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"github.com/satori/go.uuid"
)

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	uuid_, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("created uuid error: %v\n", err)
	}
	return scope.SetColumn("Id", uuid_.String())
}
