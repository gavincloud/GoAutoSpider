package main

import (
	"strconv"
	"github.com/donnie4w/go-logger/logger"
	"search"
)

func _log(i int) {
	logger.Debug("Debug>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", strconv.Itoa(i))
	//	logger.Info("Info>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", strconv.Itoa(i))
	//	logger.Warn("Warn>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
	//	logger.Error("Error>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
	//	logger.Fatal("Fatal>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
}
func init() {
	logger.SetRollingFile(`/Users/gavinding/software/logs/go`, "logs.txt", 10, 1, logger.MB)

}

func main() {

	pipe := make(chan string, 1)
	//str := httpUtil.HttpGet("https://kyfw.12306.cn/otn/leftTicket/query?leftTicketDTO.train_date=2017-08-04&leftTicketDTO.from_station=SHH&leftTicketDTO.to_station=JJS&purpose_codes=ADUL")
	//logger.Info(str)
	// 、go api.GetIps(pipe)
	//var result string
	//w:for {
	//	select {
	//	case result = <-pipe:
	//		if result == "quit" {
	//			fmt.Println("End.")
	//			break w
	//		}
	//		fmt.Println(result)
	//	}
	//
	//}
	search.BaiduSearch("proxy+ip")
	<-pipe
}

func Test() {
	//runtime.GOMAXPROCS(runtime.NumCPU())

	//指定是否控制台打印，默认为true
	//	logger.SetConsole(true)
	//	logger.SetFormat("=====>%s##%s")
	//指定日志文件备份方式为文件大小的方式
	//第一个参数为日志文件存放目录
	//第二个参数为日志文件命名
	//第三个参数为备份文件最大数量
	//第四个参数为备份文件大小
	//第五个参数为文件大小的单位

	//指定日志文件备份方式为日期的方式
	//第一个参数为日志文件存放目录
	//第二个参数为日志文件命名
	//	logger.SetRollingDaily(`C:\Users\Thinkpad\Desktop\logtest`, "test.log")

	//指定日志级别  ALL，DEBUG，INFO，WARN，ERROR，FATAL，OFF 级别由低到高
	//一般习惯是测试阶段为debug，		 生成环境为info以上
	//logger.SetLevel(logger.DEBUG)
	//
	//for i := 100; i > 0; i-- {
	//	go _log(i)
	//}
	//time.Sleep(2 * time.Second)
	//var lg = logger.GetLogger()
	//
	////重新指定log文件
	//lg.SetRollingFile(`/Users/gavinding/software/logs/go`, "test.log", 10, 1, logger.MB)
	//lg.SetLevelFile(logger.INFO, `/Users/gavinding/software/logs/go`, "info.log")
	//lg.SetLevelFile(logger.WARN, `/Users/gavinding/software/logs/go`, "warn.log")
	//lg.Debug("debug hello world")
	//for i := 100; i > 0; i-- {
	//	go lg.Info("info hello world >>>>>>>>>>>>>>>>>> ", i)
	//}
	//lg.Warn("warn hello world")
	//
	//time.Sleep(2 * time.Second)

}