package api

import (
	"util/httpUtil"
	"util/ipUtil"
)

const apiUrl string = "http://www.66ip.cn/nmtq.php?getnum=&isp=0&anonymoustype=0&start=&ports=&export=&ipaddress=&area=0&proxytype=2&api=66ip";

func GetIps(pipe chan string) {
	var ips string = httpUtil.HttpGet(apiUrl)
	array:=ipUtil.GetIpFromText(ips)
	for _, v := range array {
		pipe <- v
	}
	pipe <- "quit"

}

