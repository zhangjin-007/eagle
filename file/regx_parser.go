package file

import (
	"fmt"
	"regexp"
)

//子查询正则
const subQueryPattern string = `\([\s]*select[\s\S]*?\)`

//正则处理每一行文件
func RegxParse(fileMap map[int]string, fileName string) []LogInfo {
	logList := []LogInfo{}
	if len(fileMap) == 0{
		return logList
	}
	for lineNum, lineString := range fileMap{
		if FileLineContainsSubQuery(lineString){
			logList = append(logList, LogInfo{FileName:fileName, Location:fmt.Sprint(lineNum), Msg:"可能存在子查询 "})
		}
	}

	return logList
}


//查询某一行的sql存在子查询
func FileLineContainsSubQuery(line string)  bool{
	matched, err := regexp.MatchString(subQueryPattern, line)
	if err != nil {
		fmt.Println("😢 regexp execute error")
	}
	return matched
}


//当每行检查没有logs,这个有就要加
//这个功能需要增强
func FileFindAllSubqueryString(fileContent string)  []string{
	re := regexp.MustCompile(subQueryPattern)
	subStringArr := re.FindAllString(fileContent, -1)
	return subStringArr
}