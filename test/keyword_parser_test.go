package test

import (
	"eagle/file"
	"fmt"
	"github.com/spf13/afero"
	"testing"
)

func TestKeyWordParse(t *testing.T){
	appfs := afero.NewOsFs()
	fileName := "/Users/zhangjin/Project/eagle/test/src/RptInfoMapper.xml"
	fileByte, _ := afero.ReadFile(appfs, fileName)
	fileString :=string(fileByte)
	fileLineMap := file.BuildFileLineMap(fileString)
	logInfoArray := file.KeyWordParse(fileLineMap, fileName)
	fmt.Println(logInfoArray[0].FormatLog())
}