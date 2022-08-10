package config

import "gopkg.in/ini.v1"

var (
	MiyousheCookie string
	Act_id         string
	Region         string
	Uid            string
	Device_id      string
	App_version    string
	Client_type    string

	EmailHost     string
	EmailPort     int
	EmailUsername string
	EmailPass     string
	NotifyEmail   string
)

func InitConfig() error {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		return err
	}
	loadEmailConfig(file)
	loadCookieConfig(file)
	loadInfoConfig(file)
	return nil
}

func loadEmailConfig(file *ini.File) { //加载email服务配置
	EmailHost = file.Section("email").Key("Host").MustString("")
	EmailUsername = file.Section("email").Key("Username").MustString("")
	EmailPort = file.Section("email").Key("Port").MustInt(465)
	EmailPass = file.Section("email").Key("Pass").MustString("")
	NotifyEmail = file.Section("email").Key("NotifyEmail").MustString("")
}

func loadCookieConfig(file *ini.File) { //加载qq video cookie
	MiyousheCookie = file.Section("cookie").Key("MiyousheCookie").MustString("")
}

func loadInfoConfig(file *ini.File) { //加载米游社配置
	Act_id = file.Section("info").Key("Act_id").MustString("")
	Region = file.Section("info").Key("Region").MustString("")
	Uid = file.Section("info").Key("Uid").MustString("")
	Device_id = file.Section("info").Key("Device_id").MustString("")
	App_version = file.Section("info").Key("App_version").MustString("")
	Client_type = file.Section("info").Key("Client_type").MustString("")
}
