package database

import (
	"log"
	"os"
	"time"

	_ "github.com/y2hO0ol23/weiver/utils/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

type UserModel struct {
	gorm.Model
	ID string `gorm:"primaryKey;size:64;"`

	Wrote   []ReviewModel  `gorm:"foreignKey:FromID;"`
	Written []ReviewModel  `gorm:"foreignKey:ToID;"`
	Like    []*ReviewModel `gorm:"many2many:like;"`
	Hate    []*ReviewModel `gorm:"many2many:hate;"`
}

type ReviewModel struct {
	gorm.Model
	ID        int    `gorm:"primaryKey;autoIncrement;"`
	MessageID string `gorm:"size:64;"`
	ChannelID string `gorm:"size:64;"`
	GuildID   string `gorm:"size:64;"`

	Score   int    `gorm:"not null;"`
	Title   string `gorm:"not null;size:20;"`
	Content string `gorm:"not null;size:300;"`

	FromID string `gorm:"not null;size:64;"`
	ToID   string `gorm:"not null;size:64;"`

	Like []*UserModel `gorm:"many2many:like;"`
	Hate []*UserModel `gorm:"many2many:hate;"`

	LikeTotal int64 `gorm:"default:0;"`

	TimeStamp time.Time
}

func init() {
	db, err = gorm.Open(mysql.Open(os.Getenv("MYSQL_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal("Error loading mysql, ", err)
	}

	db.AutoMigrate(&UserModel{}, &ReviewModel{})
}
