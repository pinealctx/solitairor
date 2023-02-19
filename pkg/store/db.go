package store

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Dsn struct {
	User     string
	Password string
	Schema   string
	Host     string
	Proto    string

	// options
	MaxOpenConn int
	MaxIdle     int
	MaxLifeTime time.Duration
	DebugMode   bool

	Options map[string]string
}

func (d *Dsn) url() string {
	var url = fmt.Sprintf("%s:%s@%s(%s)/%s", d.User, d.Password, d.Proto, d.Host, d.Schema)
	if len(d.Options) == 0 {
		return url
	}
	url += "?"
	for k, v := range d.Options {
		url += fmt.Sprintf("%s=%s&", k, v)
	}
	url = strings.TrimSuffix(url, "&")
	return url
}

func NewDsn(opts ...Option) *Dsn {
	var d = genDefaultDsn()
	for _, opt := range opts {
		opt(d)
	}
	return d
}

func NewDB(opts ...Option) *gorm.DB {
	var dsn = NewDsn(opts...)
	var db, err = gorm.Open(mysql.Open(dsn.url()))
	if err != nil {
		panic(err)
	}
	if dsn.DebugMode {
		db = db.Debug()
	}
	var ddb *sql.DB
	ddb, err = db.DB()
	if err != nil {
		panic(err)
	}
	ddb.SetMaxOpenConns(dsn.MaxOpenConn)
	ddb.SetMaxIdleConns(dsn.MaxIdle)
	ddb.SetConnMaxLifetime(dsn.MaxLifeTime)

	return db
}

type Option func(*Dsn)

func User(user string) Option {
	return func(d *Dsn) {
		d.User = user
	}
}

func Password(password string) Option {
	return func(d *Dsn) {
		d.Password = password
	}
}

func Schema(schema string) Option {
	return func(d *Dsn) {
		d.Schema = schema
	}
}

func Host(host string) Option {
	return func(d *Dsn) {
		d.Host = host
	}
}

func Proto(proto string) Option {
	return func(d *Dsn) {
		d.Proto = proto
	}
}

func MaxOpenConn(maxOpenConn int) Option {
	return func(d *Dsn) {
		d.MaxOpenConn = maxOpenConn
	}
}

func MaxIdle(maxIdle int) Option {
	return func(d *Dsn) {
		d.MaxIdle = maxIdle
	}
}

func MaxLifeTime(maxLifeTime time.Duration) Option {
	return func(d *Dsn) {
		d.MaxLifeTime = maxLifeTime
	}
}

func DebugMode(debugMode bool) Option {
	return func(d *Dsn) {
		d.DebugMode = debugMode
	}
}

func Charset(charset string) Option {
	return func(dsn *Dsn) {
		dsn.Options["charset"] = charset
	}
}

func ParseTime(parseTime string) Option {
	return func(dsn *Dsn) {
		dsn.Options["parseTime"] = parseTime
	}
}

func Loc(loc string) Option {
	return func(dsn *Dsn) {
		dsn.Options["loc"] = loc
	}
}

func genDefaultDsn() *Dsn {
	var d = &Dsn{
		User:     "root",
		Password: "12345678",
		Schema:   "testDB",
		Host:     "127.0.0.1",
		Proto:    "tcp",

		MaxOpenConn: 40,
		MaxIdle:     20,
		MaxLifeTime: time.Hour,
		DebugMode:   false,

		Options: make(map[string]string),
	}
	d.Options["charset"] = "utf8mb4"
	d.Options["parseTime"] = "True"
	d.Options["loc"] = "Local"
	return d
}

var _ = User("root")
var _ = Password("12345678")
var _ = Schema("testDB")
var _ = Host("127.0.0.1")
var _ = Proto("tcp")
var _ = MaxOpenConn(40)
var _ = MaxIdle(20)
var _ = MaxLifeTime(time.Hour)
var _ = DebugMode(true)
var _ = Charset("utf8mb4")
var _ = ParseTime("True")
var _ = Loc("Local")
