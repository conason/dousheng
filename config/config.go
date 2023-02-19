package config

//Database connection's config

var (
	Name = "root"
	Password = "sx221410"
	DNS = "localhost"
	DatabasePort = 701
	DatabaseName = "minidouyin"
)

//JWT secret Key

var JwtKey = []byte("tyhngebvfpliyergfgdf")

//virtual token

var Token = "shuxindouyin"

//Qi Niu cloud config file

var(
	AccessKey = "XuigBGSCJ7vpAtRtpu04NqLGLXpEROCaqgOxTZ0W"
	SecretKey = "mhV_z93CyJCcDTmSfU2cSfx_LiejWCjujCCRMuqg"
	VideoBucket = "minidouyin-video"
	PictureBucket = "minidouyin-picture"
	Domain = "http://rq9lt9dry.bkt.clouddn.com"
)

//n:a CRUD process can get n video

const N =5
