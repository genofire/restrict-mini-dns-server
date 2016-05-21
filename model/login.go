package model

import "time"

type Login struct {
	Id         uint64    `xorm:"autoincr pk not null unique"  json:"id"`
	CreatedAt  time.Time `xorm:"created"`
	UpdatedAt  time.Time `xorm:"updated"`
	Username   string    `xorm:"not null unique 'username'" json:"username"`
	Password   string    `xorm:"'password'" json:"password"`
	Active     bool      `xorm:"'active'"`
	Superadmin bool      `xorm:"'superadmin'"`
	Code       string    `xorm:"'code'" json:"code"`
}

// TableName define table name for Xorm
func (e *Login) TableName() string {
	return "login"
}
