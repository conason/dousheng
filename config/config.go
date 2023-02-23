package config

//Database connection's config

var (
	Name         = "root"
	Password     = "sx221410"
	DNS          = "localhost"
	DatabasePort = 701
	DatabaseName = "minidouyin"
)

//JWT secret Key

var JwtKey = []byte("tyhngebvfpliyergfgdf")

//virtual token

var Token = "shuxindouyin"

//Qi Niu cloud config file

var (
	AccessKey     = "XuigBGSCJ7vpAtRtpu04NqLGLXpEROCaqgOxTZ0W"
	SecretKey     = "mhV_z93CyJCcDTmSfU2cSfx_LiejWCjujCCRMuqg"
	VideoBucket   = "minidouyin-video"
	PictureBucket = "minidouyin-picture"
	DomainVideo   = "http://rq9lt9dry.bkt.clouddn.com"
	DomainCover   = "http://rq9lfs4ld.bkt.clouddn.com"
)

//n:a CRUD process can get n video

const N = 30

const DICT = "./config/sensitive_words_lines.txt"

const TEMPTIME = "2006-01-02 15:04:05"

const REDISADDR = "127.0.0.1:6379"

const REDISDB = 0

const VIDEOSKEY = "videos"
