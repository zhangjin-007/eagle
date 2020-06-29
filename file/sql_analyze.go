package file

/**
 * *解析出来的异常信息
 */
var errorMsgList = []LogInfo{}

//日志内容
type LogInfo struct {
	FileName string
	Location string
	Msg string
}

//日志格式，是不是格式都放在这里维护
func (log *LogInfo) FormatLog() string {
	if(len(log.Location) >0 ) {
		return "文件: " + log.FileName + " " + "第" +log.Location + "行 " + log.Msg
	}else {
		return "文件: " + log.FileName + " " + log.Msg
	}
}