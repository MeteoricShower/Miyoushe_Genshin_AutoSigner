# 米游社原神自动签到
## 使用方法
`1.`在[config.ini](https://github.com/MeteoricShower/Miyoushe_Genshin_AutoSigner/blob/master/config/config.ini) 中配置好你的米游社cookie和原神uid以及收发邮件的邮箱,其余不用修改

```
[cookie]
MiyousheCookie = `Your Cookie`

[info]
Uid = your uid

[email]
Username = 发送邮件的email@qq.com
Pass = 发送邮件的email的smtp授权码
NotifyEmail = 接收邮件的email@qq.com
```
`2.`在你配置好docker的服务器上输入以下代码
```
docker build -t mys_genshin_autosigner:v0 .
docker run -i -t mys_genshin_autosigner:v0
```
