package svc

import (
	"backend/internal/config"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Mysql  *gorm.DB
	Redis  *redis.Client
}
type Sbjjson struct {
	Lession1  string `json:"lession1"`
	Lession2  string `json:"lession2"`
	Lession3  string `json:"lession3"`
	Lession4  string `json:"lession4"`
	Lession5  string `json:"lession5"`
	Lession6  string `json:"lession6"`
	Lession7  string `json:"lession7"`
	Lession8  string `json:"lession8"`
	Lession9  string `json:"lession9"`
	Lession10 string `json:"lession10"`
	Lession11 string `json:"lession11"`
	Lession12 string `json:"lession12"`
	Lession13 string `json:"lession13"`
}
type Teachjson struct {
	Teacher1  string `json:"teacher1"`
	Teacher2  string `json:"teacher2"`
	Teacher3  string `json:"teacher3"`
	Teacher4  string `json:"teacher4"`
	Teacher5  string `json:"teacher5"`
	Teacher6  string `json:"teacher6"`
	Teacher7  string `json:"teacher7"`
	Teacher8  string `json:"teacher8"`
	Teacher9  string `json:"teacher9"`
	Teacher10 string `json:"teacher10"`
	Teacher11 string `json:"teacher11"`
	Teacher12 string `json:"teacher12"`
	Teacher13 string `json:"teacher13"`
}
type Srcjson struct {
	Src1  string `json:"src1"`
	Src2  string `json:"src2"`
	Src3  string `json:"src3"`
	Src4  string `json:"src4"`
	Src5  string `json:"src5"`
	Src6  string `json:"src6"`
	Src7  string `json:"src7"`
	Src8  string `json:"src8"`
	Src9  string `json:"src9"`
	Src10 string `json:"src10"`
	Src11 string `json:"src11"`
	Src12 string `json:"src12"`
	Src13 string `json:"src13"`
}
type Userinf struct {
	Uuid      int64  `gorm:"primaryKey"`
	ClassCode string `gorm:"size:20;index:idx_user"`
	Passwd    string `gorm:"size:20;index:idx_user"`
	Info      string `gorm:"size:20"`
}
type Subject struct {
	Uuid      int64     `gorm:"primaryKey"`
	ClassCode string    `gorm:"size:20;index:idx_subj"`
	Weekly    int       `gorm:"index:idx_subj"`
	Sbjname   Sbjjson   `gorm:"serializer:json"`
	Teacher   Teachjson `gorm:"serializer:json"`
	Src       Srcjson   `gorm:"serializer:json"`
}
type Message struct {
	Uuid      int64  `gorm:"primaryKey;index:idx_msgid"`
	ClassCode string `gorm:"size:20"`
	Msg       string `gorm:"size:200"`
	MsgTitle  string `gorm:"size:50"`
	Time      time.Time
}
type Log struct {
	Uuid int64  `gorm:"primaryKey"`
	Type string `gorm:"size:20"`
	Logs string `gorm:"size:1000"`
	time.Time
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := "digiclass:nhZhfep3pK3txs2c@tcp(119.23.69.180:3306)/digiclass?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	Mysql, err := gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true})
	sqlDB, err2 := Mysql.DB()
	if err != nil {
		println(err.Error())
	}
	if err2 != nil {
		println(err2)
	}
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(95)
	Mysql.AutoMigrate(&Userinf{})
	Mysql.AutoMigrate(&Subject{})
	Mysql.AutoMigrate(&Message{})
	Mysql.AutoMigrate(&Log{})
	Redis := redis.NewClient(&redis.Options{
		Addr:     "119.23.69.180:6378",
		Password: "Qxr7g9hh386.",
		DB:       0,
	})
	_, rerr := Redis.Ping().Result()
	if rerr != nil {
		print(rerr.Error())
	}
	return &ServiceContext{
		Config: c,
		Mysql:  Mysql,
		Redis:  Redis,
	}
}
