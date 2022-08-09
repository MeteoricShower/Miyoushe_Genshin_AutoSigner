package service

import (
	"Miyoushe_Genshin_AutoSigner/config"
	"Miyoushe_Genshin_AutoSigner/util"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

func SignRequest() *http.Request {
	Cookie := config.MiyousheCookie
	SignUrl, _ := url.Parse("https://api-takumi.mihoyo.com/event/bbs_sign_reward/sign")
	values := SignUrl.Query()
	values.Set("act_id", config.Act_id)
	values.Set("region", config.Region)
	values.Set("uid", config.Uid)
	SignUrl.RawQuery = values.Encode()

	//req, err := http.NewRequest(http.MethodGet, SignUrl.String(), nil)

	data := map[string]string{
		"act_id": config.Act_id,
		"region": config.Region,
		"uid":    config.Uid,
	}
	jsonByte, err := json.Marshal(data)

	req, err := http.NewRequest(http.MethodPost, SignUrl.String(), bytes.NewBuffer(jsonByte))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Cookie", Cookie)
	req.Header.Set("x-rpc-device_id", config.Device_id)
	req.Header.Set("x-rpc-app_version", config.App_version)
	req.Header.Set("x-rpc-client_type", config.Client_type)
	req.Header.Set("DS", util.GetDs())

	return req
}
