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
	Lession14 string `json:"lession14"`
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
	Teacher14 string `json:"teacher14"`
}
type Srcjson struct {
	Src1  bool `json:"src1"`
	Src2  bool `json:"src2"`
	Src3  bool `json:"src3"`
	Src4  bool `json:"src4"`
	Src5  bool `json:"src5"`
	Src6  bool `json:"src6"`
	Src7  bool `json:"src7"`
	Src8  bool `json:"src8"`
	Src9  bool `json:"src9"`
	Src10 bool `json:"src10"`
	Src11 bool `json:"src11"`
	Src12 bool `json:"src12"`
	Src13 bool `json:"src13"`
	Src14 bool `json:"src14"`
}
type Userinf struct { //用户大表
	Uuid      int64  `gorm:"primaryKey"`
	ClassCode string `gorm:"size:20;index:idx_user"`
	Passwd    string `gorm:"size:20;index:idx_user"`
	Info      string `gorm:"size:20"`
}
type ClassSubject struct { //课表
	Uuid      int64     `gorm:"primaryKey"`
	ClassCode string    `gorm:"size:20;index:idx_subj"`
	Weekly    int8      `gorm:"index:idx_subj"`
	Sbjname   Sbjjson   `gorm:"serializer:json"`
	Teacher   Teachjson `gorm:"serializer:json"`
	Src       Srcjson   `gorm:"serializer:json"`
}
type Message struct { //消息发布
	Uuid      int64  `gorm:"primaryKey;index:idx_msgid"`
	ClassCode string `gorm:"size:20"`
	Msg       string `gorm:"size:200"`
	MsgTitle  string `gorm:"size:50"`
	Sender    string `gorm:"size:20"`
	Time      time.Time
}
type Log struct { //系统日志
	Uuid int64  `gorm:"primaryKey"`
	Type string `gorm:"size:20"`
	Logs string `gorm:"size:1000"`
	time.Time
}
type Student struct { //学生大表
	Uuid    int64  `gorm:"primaryKey"`
	JkUid   string `gorm:"size:10"`               //极课号
	Name    string `gorm:"size:10"`               //姓名
	Class   string `gorm:"size:20;index:idx_stu"` //所在班级
	Subject string `gorm:"size:10;index:idx_stu"` //选课
	Score   int16  //当前积分
}
type StudentScore struct { //学生积分日志
	Uuid       int64  `gorm:"primaryKey"`
	Name       string `gorm:"size:10;index:idx_stuscore"` //学生姓名
	Type       string `gorm:"size:10;index:idx_stuscore"` //记录类型
	Class      string `gorm:"size:20;index:idx_stuscore"` //所在班级
	Score      int16  //加减积分
	SubInfo    string `gorm:"size:100"` //备注
	UpdateTime time.Time
}
type TeacherLession struct { //老师课时统计
	Uuid             int64  `gorm:"primaryKey"`
	Teacher          string `gorm:"size:10;index:idx_tec"`
	Subject          string `gorm:"size:10;index:idx_tec"`
	LastMonthLession int16
	Lession          int16
	UpdateTime       time.Time
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
	sqlDB.SetMaxOpenConns(95)
	Mysql.AutoMigrate(&Userinf{})
	Mysql.AutoMigrate(&ClassSubject{})
	Mysql.AutoMigrate(&Message{})
	Mysql.AutoMigrate(&Log{})
	Mysql.AutoMigrate(&Student{})
	Mysql.AutoMigrate(&StudentScore{})
	Mysql.AutoMigrate(&TeacherLession{})
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
