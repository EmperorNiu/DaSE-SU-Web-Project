package conf

import(
	"github.com/go-ini/ini"
	"log"
)

var (
	RunMode string
	Port int

	User string
	Password string
	Host string
	Database string
	DBType string
	DbUrl string
)

func init() {
	Cfg,err := ini.Load("conf/app.ini")
	if err!=nil{
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	RunMode = Cfg.Section("").Key("RUN_MODE").String()
	Port, _ = Cfg.Section("").Key("HTTP_PORT").Int()

	loadDatabase(Cfg)
}

func loadDatabase(Cfg *ini.File) {
	sec,err := Cfg.GetSection("database")
	if err!=nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}
	User = sec.Key("USER").String()
	Password = sec.Key("PASSWORD").String()
	Host = sec.Key("PORT").String()
	Database = sec.Key("DATABASE").String()
	DBType = sec.Key("TYPE").String()
	DbUrl = User + ":" + Password + "@tcp(" + Host + ")/" +
		Database + "?parseTime=true"
}
