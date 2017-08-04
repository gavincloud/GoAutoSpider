package ipUtil

import (
	"regexp"
	"strings"
)

func GetIpFromText(src string) []string {

	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	reg := regexp.MustCompile(`((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?):\d{2,5}`)
	array := reg.FindAllString(src, -1)
	return array
}
func MathIpPort(src string) bool {

	reg := regexp.MustCompile(`[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}:[0-9]{2,5}`)
	return reg.MatchString(src)
}
func MathIp(src string) bool {

	reg := regexp.MustCompile(`((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)`)
	array := reg.MatchString(src)
	return array
}
func MathPort(src string) bool {

	reg := regexp.MustCompile(`^[0-9]*$`)
	array := reg.MatchString(src)
	return array
}