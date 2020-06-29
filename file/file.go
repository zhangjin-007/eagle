package file

import (
	"fmt"
	"log"
	"github.com/spf13/afero"
	"os"
	"strings"
)

/**
过滤
 */
func filter(ss []string, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

var scanFileArr []string

var logInfoList []LogInfo

//执行xml文件的扫描 walk 方法使用RegexpFs
func Xmlscan(projectPath string) {
	appfs := afero.NewOsFs()
	var err = afero.Walk(appfs, projectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(path, ".xml") && !strings.Contains(path, "pom.xml") && !strings.Contains(path, "target") && !strings.Contains(path, "workspace") {
			scanFileArr = append(scanFileArr, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal("😂 scanner error, may be you should try the correct path")
		return
	}

	//1.找出所有mybatis文件
	scanFileArr = filter(scanFileArr, filterMybatisFiles)

	for _,path  := range scanFileArr{
		//KeyWordParse(, path)
		mybatisXmlFile := afero.NewOsFs()
		fileBytes,err := afero.ReadFile(mybatisXmlFile, path)
		if err != nil {
			log.Println("😂 xml file read error !")
			return
		}
		fileContent := string(fileBytes)
		//构建行号内容map
		fileLineMap := BuildFileLineMap(fileContent)
		//关键字收集日志
		logInfoList = append(logInfoList, KeyWordParse(fileLineMap, path)...)
	}
	fmt.Println(logInfoList)
}

/**
过滤操作mybatis的文件🚀
*/
func filterMybatisFiles(filePath string) bool {
	appfs := afero.NewOsFs()
	bytePath := []byte("mybatis")
	exist, er := afero.FileContainsBytes(appfs, filePath, bytePath)

	if er != nil {
		log.Fatal("afero filer scan error")
	}
	return exist
}

//构建文件的 行号跟内容的map
func BuildFileLineMap(fileContent string) map[int]string {
	if len(fileContent) <= 0 {
		log.Fatal("scan file cannot empty")
	}

	fileLineMap := map[int]string{}

	temp := strings.Split(fileContent, "\n")

	for index, item := range temp{
		fileLineMap[index+1] = item
	}
	return fileLineMap
}