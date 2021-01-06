package tool

import "fmt"

/*
时间常量
*/
const (
	//定义每分钟的秒数
	SecondsPerMinute = 60
	//定义每小时的秒数
	SecondsPerHour = SecondsPerMinute * 60
	//定义每天的秒数
	SecondsPerDay = SecondsPerHour * 24
)

// 将秒转换为 x分x秒 的格式
func SecondsFormat(duration int) string {
	minute := duration / SecondsPerMinute
	second := duration % SecondsPerMinute
	return fmt.Sprintf("%d分%d秒", minute, second)
}
