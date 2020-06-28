package file


var keyWordArr = []string{"using","rand()", "union", "select *", "*"}

//关键词的内容提示
func KeyWordParse(fileMap map[int]string, fileName string) []LogInfo {
	logList := []LogInfo{}
	if len(fileMap) == 0{
		return logList
	}

	for k, v := range fileMap{
		if c,keyword :=contains(keyWordArr, v); c {
			logList = append(logList, LogInfo{FileName:fileName, Location:string(k), Msg:"存在关键字"+ keyword})
		}
	}

	return logList
}

func contains(a []string, x string) (bool, string)  {
	for _, n := range a {
		if x == n {
			return true,x
		}
	}
	return false,x
}