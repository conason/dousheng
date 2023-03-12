package config

//Database connection's config

var (
	MySQLUser    = "root"
	MySQLPWD     = "root"
	DNS          = "localhost"
	DatabasePort = 3306
	DatabaseName = "dousheng_db"
)

//JWT secret Key

var JwtKey = []byte("tyhngebvfpliyergfgdf")

//virtual token

var Token = "shuxindouyin"

//Qi Niu cloud config file

var (
	AccessKey     = "PtGMD2Uxg-lUSBzFhdRcE6xHGZmPR1vxGGhR56e4"
	SecretKey     = ""
	VideoBucket   = "conason"
	PictureBucket = "conason-pic"
	DomainVideo   = "http://rrb05r4gx.hn-bkt.clouddn.com"
	DomainCover   = "http://rrb1dnse4.hn-bkt.clouddn.com"
)

//n:a CRUD process can get n video

const N = 30

const DICT = "./config/sensitive_words_lines.txt"

const TEMPTIME = "2006-01-02 15:04:05"

const REDISADDR = "127.0.0.1:6379"

const REDISDB = 0

const VIDEOSKEY = "videos"
