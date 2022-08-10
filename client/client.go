package client

import (
	"Miyoushe_Genshin_AutoSigner/config"
	"Miyoushe_Genshin_AutoSigner/email"
	"Miyoushe_Genshin_AutoSigner/service"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

//type HTTPClient struct {
//	*http.Client
//}

//func GetRandString() string {
//	r := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(100000)
//	r = r + 100000
//	return strconv.FormatInt(r, 10)
//}

var (
	wg           sync.WaitGroup
	cstSh, _     = time.LoadLocation("Asia/Shanghai")
	SignTimeRule = "0 01 00 * * *" //每天00:01:00执行 每日签到
)

func NewHttpClient() *http.Client {
	return &http.Client{}
}

func Task() {
	HTTPClient := NewHttpClient()
	reqSign := service.SignRequest()
	rspSign, err := HTTPClient.Do(reqSign) //发送签到请求
	if err != nil {
		log.Fatalln(err)
	}
	reqSignInfo := service.SignInfoRequest() //发送签到信息请求
	rspSignInfo, err := HTTPClient.Do(reqSignInfo)
	if err != nil {
		log.Fatalln(err)
	}
	respBodySign, err := ioutil.ReadAll(rspSign.Body)
	respBodySignInfo, err := ioutil.ReadAll(rspSignInfo.Body)
	email.SendEmail(config.NotifyEmail, "米游社签到成功通知", string(respBodySign)+"\n"+string(respBodySignInfo))
	log.Print(string(respBodySign))
}

func runTask(c *cron.Cron, timeRule string) {
	_, err := c.AddFunc(timeRule, func() { //开启定时任务
		Task() //任务处理
	})
	if err != nil {
		log.Fatalln("定时器启动失败,err:", err.Error())
	}
	c.Start()
}

func Run() {
	time.Local = cstSh
	c := cron.New(cron.WithSeconds(), cron.WithLocation(cstSh))
	go runTask(c, SignTimeRule)
	log.Println("服务已启动...")

	wg.Add(1)
	wg.Wait()
}
