package svc

import (
	"backend/internal/config"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/prometheus"
)

type ServiceContext struct {
	Config config.Config
	Mysql  *gorm.DB
	Redis  *redis.Client
}
type Sbjjson struct {
	Lession1  string
	Lession2  string
	Lession3  string
	Lession4  string
	Lession5  string
	Lession6  string
	Lession7  string
	Lession8  string
	Lession9  string
	Lession10 string
	Lession11 string
	Lession12 string
	Lession13 string
	Lession14 string
}
type Teachjson struct {
	Teacher1  string
	Teacher2  string
	Teacher3  string
	Teacher4  string
	Teacher5  string
	Teacher6  string
	Teacher7  string
	Teacher8  string
	Teacher9  string
	Teacher10 string
	Teacher11 string
	Teacher12 string
	Teacher13 string
	Teacher14 string
}
type Srcjson struct {
	Src1  bool
	Src2  bool
	Src3  bool
	Src4  bool
	Src5  bool
	Src6  bool
	Src7  bool
	Src8  bool
	Src9  bool
	Src10 bool
	Src11 bool
	Src12 bool
	Src13 bool
	Src14 bool
}
type Userinf struct { //用户大表
	Uuid      int64  `gorm:"primaryKey"`
	ClassCode string `gorm:"size:20;index:idx_user"`
	Passwd    string `gorm:"size:20;index:idx_user"`
	Info      string `gorm:"size:20"`
}
type ClassSubject struct { //课表json
	Uuid      int64     `json:"uuid"`
	ClassCode string    `json:"classcode"`
	Weekly    int8      `json:"weekly"`
	Sbjname   Sbjjson   `json:"sbjname"`
	Teacher   Teachjson `json:"teacher"`
	Src       Srcjson   `json:"src"`
}
type ClassSubject2 struct { //课表
	Uuid      int64  `gorm:"primaryKey"`
	ClassCode string `gorm:"size:20;index:idx_subj"`
	Weekly    int8   `gorm:"index:idx_subj"`
	Sbjname   string
	Teacher   string
	Src       string
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
	Mysql.Use(prometheus.New(prometheus.Config{
		DBName:          "digicalssDB",               // 使用 `DBName` 作为指标 label
		RefreshInterval: 15,                          // 指标刷新频率（默认为 15 秒）
		PushAddr:        "prometheus pusher address", // 如果配置了 `PushAddr`，则推送指标
		StartServer:     true,                        // 启用一个 http 服务来暴露指标
		HTTPServerPort:  10002,                       // 配置 http 服务监听端口，默认端口为 8080 （如果您配置了多个，只有第一个 `HTTPServerPort` 会被使用）
		MetricsCollector: []prometheus.MetricsCollector{
			&prometheus.MySQL{
				VariableNames: []string{"Threads_running"},
			},
		}, // 用户自定义指标
	}))
	sqlDB, err2 := Mysql.DB()
	if err != nil {
		println(err.Error())
	}
	if err2 != nil {
		println(err2)
	}
	sqlDB.SetMaxOpenConns(95)
	Mysql.AutoMigrate(&Userinf{})
	Mysql.AutoMigrate(&ClassSubject2{})
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
