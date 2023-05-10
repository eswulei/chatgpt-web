package model

import (
	"github.com/eswulei/chatgpt-web/pkg/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// BaseModel 主模型
type BaseModel struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement;not null"`
	CreatedAt time.Time `gorm:"column:created_at;index"`
	UpdatedAt time.Time `gorm:"column:updated_at;index"`
}

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(1024)"`
	Email    string `gorm:"type:varchar(1024)"`
	Password string `gorm:"type:varchar(1024)"`
	//Saltcode string `gorm:"type:varchar(1024)"`
	//Nickname string `gorm:"type:varchar(1024)"`
	ExistExPassword bool
	ExPassword      string `gorm:"type:varchar(1024)"`
}

type Message struct {
	gorm.Model
	UserName    string `gorm:"type:varchar(1024)"`
	MessageFrom string `gorm:"type:varchar(1024)"`
	Message     []byte `gorm:"type:blob"`
}

// GetStringID 获取主键字符串
func (model BaseModel) GetStringID() string {
	return types.UInt64ToString(model.ID)
}

// CreatedAtDate 获取创建时间
func (model BaseModel) CreatedAtDate() string {
	return model.CreatedAt.Format("2006-01-02")
}

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	//dsn := "chat.db"
	//var err error
	//DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
	//	Logger: gloger.Default.LogMode(gloger.Info),
	//})
	//if err != nil {
	//	logger.Danger("open sqlite error:", err)
	//}
	//return DB
	dsn := "root:renejia19890304@tcp(127.0.0.1:3306)/chat_gpt?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return DB
}
