package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/y2hO0ol23/weiver/env"
)

var (
	db *gorm.DB
)

type UserModel struct {
	gorm.Model
	ID string `gorm:"primaryKey;size:64;"`

	Wrote   []ReviewModel  `gorm:"foreignKey:AuthorID;"`
	Written []ReviewModel  `gorm:"foreignKey:SubjectID;"`
	Like    []*ReviewModel `gorm:"many2many:like;"`
	Hate    []*ReviewModel `gorm:"many2many:hate;"`

	Role []*RoleModel `gorm:"many2many:role;"`
}

type ReviewModel struct {
	gorm.Model
	ID        int    `gorm:"primaryKey;autoIncrement;"`
	MessageID string `gorm:"size:64;"`
	ChannelID string `gorm:"size:64;"`
	GuildID   string `gorm:"size:64;"`

	DMMessageID string `gorm:"size:64"`

	Score   int    `gorm:"not null;"`
	Title   string `gorm:"not null;size:20;"`
	Content string `gorm:"not null;size:300;"`

	AuthorID  string `gorm:"not null;size:64;"`
	SubjectID string `gorm:"not null;size:64;"`

	Like []*UserModel `gorm:"many2many:like;"`
	Hate []*UserModel `gorm:"many2many:hate;"`

	LikeTotal int64 `gorm:"default:0;not null;"`

	TimeStamp time.Time
}

type GuildModel struct {
	gorm.Model
	ID        string `gorm:"primaryKey;size:64;"`
	AllowRole bool   `gorm:"default:false"`

	InProgress bool

	Role []RoleModel `gorm:"foreignKey:GuildID;"`
}

type RoleModel struct {
	gorm.Model
	ID      string `gorm:"primaryKey;size:128;"`
	RoleID  string `gorm:"not null;"`
	GuildID string `gorm:"not null;"`

	Display string `gorm:"size:100"`

	User []*UserModel `gorm:"many2many:role;"`
}

func init() {
	var err error
	db, err = gorm.Open(mysql.Open(os.Getenv("MYSQL_URL")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Panicf("Error loading mysql\n%v", err)
	}

	db.AutoMigrate(&UserModel{}, &GuildModel{})
	db.AutoMigrate(&ReviewModel{}, &RoleModel{})
}
