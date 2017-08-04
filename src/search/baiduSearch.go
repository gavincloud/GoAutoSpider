package search

import (
	"github.com/PuerkitoBio/goquery"
	"logger"
	"util/ipUtil"
	"strings"
	"strconv"
	"time"
	"sync"
)

const pageSize int = 10

var containerMap = make(map[string]string)
var LockMap sync.Mutex
/**
 * 百度搜索关键字
 */
func BaiduSearch(keyWord string) {

	//定义翻页数量
	var index int = 1

	var str string = "http://www.baidu.com/s?pn="
	str += strconv.Itoa((index - 1) * pageSize)
	str += "&wd=" + keyWord
	for {
		if (index > 5) {
			break
		}
		//根据百度搜索获取url
		doc, err := goquery.NewDocument(str)
		if (err != nil) {
			logger.Error("error occur")
			return
		}
		time.Sleep(1000)
		//处理url
		go dealSearch(doc)
		index ++;

	}

}
// 获取url
func dealSearch(doc *goquery.Document) {
	doc.Find("h3>a").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		//logger.Info(url)
		time.Sleep(1000)
		go dealPage(url)
	})

}
//解析页面
func dealPage(url string) {
	//如果url抓取过了 就不再抓取
	//放入map 有并发问题
	isOk := get(url)
	if len(isOk) != 0 {
		return
	}
	set(url, url)
	time.Sleep(1000)
	//content := httpUtil.HttpGet(url)
	//reader := strings.NewReader(content)

	//doc, err := goquery.NewDocumentFromReader(reader)
	//根据url获取页面
	doc, err := goquery.NewDocument(url)
	if (err != nil) {
		return
	}
	//继续判断这个页面是否含有有效的链接
	go judgeA(doc, doc.Url.Host)
	//判断文字是否含有ip代理信息
	go judgeText(doc.Text())
	//list := list.New()
	//解析table 定位ip 和port的位置
	x, y := getPos(doc)

	//判断是否有合理的table标签
	tableText := doc.Find("table").Text()

	if (!(strings.Contains(strings.ToLower(tableText), "时间") || strings.Contains(strings.ToLower(tableText), "ip") || strings.Contains(strings.ToLower(tableText), "port")&&strings.Contains(strings.ToLower(tableText), "端口")&&strings.Contains(strings.ToLower(tableText), "last"))) {
		return
	}

	//解析table 解析代理ip
	doc.Find("table").Each(func(i int, s *goquery.Selection) {
		s.Find("tr").Each(func(i int, s *goquery.Selection) {

			var ip string;
			var port string;
			s.Find("td").Each(func(i int, s *goquery.Selection) {
				if ( i == x) {
					ip = s.Text()
				}
				if (i == y ) {
					port = s.Text()
				}
			})
			//portInt, _ := strconv.Atoi(port)
			if len(ip) != 0 && ipUtil.MathIp(ip)&& len(port) != 0&&ipUtil.MathPort(port) {
				//ips := ip + port
				//result := httpUtil.HttpGetProxy("https://kyfw.12306.cn/otn/leftTicket/query?leftTicketDTO.train_date=2017-08-04&leftTicketDTO.from_station=SHH&leftTicketDTO.to_station=JJS&purpose_codes=ADULT", ips)
				//logger.Info(result)
				logger.Info(ip, ":", port)
			}
		})

	})
	//logger.Info, "抓取结束")
	//logger.Info(list)

}
//定位位置
func getPos(doc *goquery.Document) (x, y int) {
	var ipPos int = -1
	var portPos int = -1

	doc.Find("table").Each(func(i int, s *goquery.Selection) {
		s.Find("tr").Each(func(i int, s *goquery.Selection) {
			s.Find("td").Each(func(i int, s *goquery.Selection) {

				var temp string = s.Text()
				//判断ip和端口是否放在一起的
				if (!strings.Contains(temp, ":")) {
					if ipUtil.MathIp(temp) {
						ipPos = i
					}
					if (ipUtil.MathPort(temp)) {
						portPos = i;
					}
					if ipPos != -1 && portPos != -1 {
						return
					}

				} else {
					if ipUtil.MathIpPort(temp) {
						ipPos = i
						portPos = i
						return
					}
				}
			})
		})

	})
	return ipPos, portPos
}
//判断A标签
func judgeA(doc *goquery.Document, host string) {
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		var text string = s.Text()
		var url, _ = s.Attr("href");
		if strings.ContainsAny(text, "代理") || strings.ContainsAny(text, "ip") {
			if (strings.HasSuffix(url, "html")) {
				if (strings.HasPrefix(url, "http")) {
					go dealPage(url)
				} else if strings.HasPrefix(url, "/") {
					//logger.Info("http://" + host + url)
					go dealPage("http://" + host + url)
				} else {
					return
				}
			}
		}

		d, error := strconv.Atoi(text)
		if error == nil {
			if (strings.Index(url, text) >= 0 && d <= 30) {
				go dealPage("http://" + host + url)
			}
		}

		if (strings.Index(text, "月") >= 0&&strings.Index(text, "日") >= 0) {
			t1 := time.Now()
			_, month, day := t1.Date()
			dayStr := strconv.Itoa(day)
			monthStr := month.String()
			if (strings.Contains(text, dayStr)&&strings.Contains(text, monthStr)) {
				logger.Info(text)
				go dealPage(url)
			}
		}

	})
}
//判断是否有文字IP信息
func judgeText(src string) {
	var array = ipUtil.GetIpFromText(src)
	for index := range array {
		logger.Info(array[index])
	}
}


//加锁map get
func get(key string) string {
	LockMap.Lock()
	defer LockMap.Unlock()
	return containerMap[key]
}
//加锁map set
func set(k, v string) {
	LockMap.Lock()
	defer LockMap.Unlock()
	containerMap[k] = v
}