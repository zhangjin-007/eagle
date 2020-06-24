package file

/**
 * *解析出来的异常信息
 */
var errorMsgList = []LogInfo{}

//日志内容
type LogInfo struct {
	FileName string
	location string
	msg string
}

//日志格式，是不是格式都放在这里维护
func (log *LogInfo) FormatLog() string {
	return "文件: " + log.FileName + log.msg
}