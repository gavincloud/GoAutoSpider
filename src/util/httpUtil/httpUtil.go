package httpUtil

import (
	"net/http"
	"io/ioutil"
	"strings"
	"fmt"
	"net/url"
	"log"
	"net"
	"time"
	"crypto/tls"
	"logger"
)

var c *http.Client = &http.Client{

	Transport: &http.Transport{
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, time.Second * 5)
			if err != nil {
				fmt.Println("dail timeout", err)
				return nil, err
			}
			return c, nil

		},
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * 5,
	},
}
/**
Accept:
Accept-Encoding:gzip, deflate
Accept-Language:zh-CN,zh;q=0.8,en;q=0.6
Cache-Control:max-age=0
Connection:keep-alive
Cookie:_free_proxy_session=BAh7B0kiD3Nlc3Npb25faWQGOgZFVEkiJTcyMjM1YTU2ZjJmMzI3MDA3MDgxMDliNWY0NjdhY2NiBjsAVEkiEF9jc3JmX3Rva2VuBjsARkkiMWR4SGxIMzVHcGVOY25ncEkzbzVtQmNDbkZ4a0ZqZEdhY2RHRDBJZ21aMFU9BjsARg%3D%3D--93730f28c0a7458b9342d10d821a945f21a7c188; Hm_lvt_0cf76c77469e965d2957f0553e6ecf59=1500126412,1500126431,1500126497,1501655388; Hm_lpvt_0cf76c77469e965d2957f0553e6ecf59=1501655388
Host:www.xicidaili.com
If-None-Match:W/"efd09ce809ed7787a5c824d0f0098834"
Referer:http://www.baidu.com/link?url=4r-6aQt0lki-feczbb9WUwxQ3GYlooKuCBwN6rQLw6109flibWPw9sY3LdaO0GQt&wd=&eqid=b8822c6900001cbb0000000459817167
Upgrade-Insecure-Requests:1
User-Agent:Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36
 */
func HttpGet(uri string) string {
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Add("Accept", "*/*")
	//req.Header.Add("Accept-Encoding","gzip, deflate")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", "_free_proxy_session=BAh7B0kiD3Nlc3Npb25faWQGOgZFVEkiJTcyMjM1YTU2ZjJmMzI3MDA3MDgxMDliNWY0NjdhY2NiBjsAVEkiEF9jc3JmX3Rva2VuBjsARkkiMWR4SGxIMzVHcGVOY25ncEkzbzVtQmNDbkZ4a0ZqZEdhY2RHRDBJZ21aMFU9BjsARg%3D%3D--93730f28c0a7458b9342d10d821a945f21a7c188; Hm_lvt_0cf76c77469e965d2957f0553e6ecf59=1500126412,1500126431,1500126497,1501655388; Hm_lpvt_0cf76c77469e965d2957f0553e6ecf59=1501655388")
	req.Header.Add("User-Agent", "User-Agent:Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")
	resp, err := c.Do(req)
	if err != nil {
		logger.Info("http请求异常:", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)

}
/**
unc Post(url string, contentType string, body io.Reader) (resp *Response, err error) {
 */
func HttpPost(uri string, param string, mime string) string {
	resp, err := http.Post(uri, mime, strings.NewReader(param))
	if err != nil {
		// handle error
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return ""
	}
	return string(body)
}

func HttpDoGet(url string, method string, param string) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, strings.NewReader(param))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func HttpGetProxy(urls string, proxyIp string) string {
	proxy := func(_ *http.Request) (*url.URL, error) {
		logger.Info(proxyIp)
		return url.Parse(proxyIp)
	}
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport}

	resp, err := client.Get(urls) //请求并获取到对象,使用代理
	if err != nil {
		log.Fatal(err)
	}

	dataproxy, err := ioutil.ReadAll(resp.Body) //取出主体的内容
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s",data) //打印
	sproxy := string(dataproxy);
	resp.Body.Close()
	return sproxy

}