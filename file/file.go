package file

import (
	"github.com/prometheus/common/log"
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

var scanFileArr = []string{}

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
		log.Error("scanner error")
	}
	scanFileArr = filter(scanFileArr, filterMybatisFiles)


	for _,path  := range scanFileArr{
		log.Info(path)
	}

}

/**
过滤操作mybatis的文件
*/
func filterMybatisFiles(filePath string) bool {
	appfs := afero.NewOsFs()
	bytePath := []byte("mybatis")
	exist, er := afero.FileContainsBytes(appfs, filePath, bytePath)

	if er != nil {
		log.Error("afero filer scan error")
	}
	return exist
}

//构建文件的 行号跟内容的map
func BuildFileLineMap(fileContent string) map[int]string {
	if len(fileContent) <= 0 {
		log.Error("scan file cannot empty")
	}

	fileLineMap := map[int]string{}

	temp := strings.Split(fileContent, "\n")

	for index, item := range temp{
		fileLineMap[index+1] = item
	}
	return fileLineMap
}