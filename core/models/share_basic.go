package models

import "time"

type ShareBasic struct {
	Id                     int
	Identity               string
	UserIdentity           string
	RepositoryIdentity     string
	UserRepositoryIdentity string
	ExpiredTime            int
	ClickNum               int
	CreatedAt              time.Time `xorm:"created"`
	UpdatedAt              time.Time `xorm:"created"`
	DeletedAt              time.Time `xorm:"created"`
}

func (table ShareBasic) TableName() string {
	return "share_basic"
}
