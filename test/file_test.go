package test

import (
	"eagle/file"
	"fmt"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var scanFileArr = []string{}



func TestExist(t *testing.T) {
	appfs := afero.NewOsFs()
	appfs.MkdirAll("src/a", 0755)
	// 创建 test 文件 和 目录
	//appFS.MkdirAll("src/a", 0755)
	afero.WriteFile(appfs, "src/a/b", []byte("file b"), 0644)
	afero.WriteFile(appfs, "src/c", []byte("file c"), 0644)

	exists, _ := afero.Exists(appfs, "src/a/b")
	fmt.Println("file :", exists)
	//afero.(appfs, "src")
	//name := "src/c"
	//_, err := appFS.Stat(name)
	//if os.IsNotExist(err) {
	//	t.Errorf("file \"%s\" does not exist.\n", name)
	//}
}

func TestWalk(t *testing.T)  {
	err := filepath.Walk("/Users/zhangjin/Project/user-center",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size())
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}


/**
找出所有的 xml文件
 */
func TestWalk2(t *testing.T)  {
	err := filepath.Walk("/Users/zhangjin/Project/user-center",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if(strings.Contains(path, ".xml") && !strings.Contains(path, "pom.xml")){
				fmt.Println(path, info.Size())
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

func TestAferoWalk(t *testing.T) {
	appfs := afero.NewOsFs()
	err := afero.Walk(appfs, "/Users/zhangjin/Project/user-center", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if(strings.Contains(path, ".xml") && !strings.Contains(path, "pom.xml")){
			fmt.Println(path, info.Size())
			scanFileArr = append(scanFileArr, path)
		}
		return nil
	})
	log.Println("file array ", scanFileArr)


	assert.Nil(t, err)
}

/**
* 文件包含内容
 */
func TestFileContains(t *testing.T)  {

	appfs := afero.NewOsFs()

	path := "/Users/zhangjin/Project/user-center/ucenter-common/src/main/resources/com/yzf/ucenter/common/mapper/UserExtraInfoDAO.xml"
	bytePath := []byte("mybatis")
	exist,er := afero.FileContainsBytes(appfs, path, bytePath)

	if er !=nil{

	}
	assert.Equal(t,exist, true)
}

func TestXmlscan(t *testing.T)  {
	//projectPath := "/Users/zhangjin/Project/user-center"
	projectPath := "/Users/zhangjin/Project/eagle"
	file.Xmlscan(projectPath)
}

func TestFileReadAt(t *testing.T)  {
	appfs := afero.NewOsFs()
	fileByte, _ := afero.ReadFile(appfs, "/Users/zhangjin/Project/user-center/ucenter-common/src/main/resources/com/yzf/ucenter/common/mapper/UserExtraInfoDAO.xml")
	fileString :=string(fileByte)
	log.Println(fileString)
}

func TestFileLine(t *testing.T)  {
	appfs := afero.NewOsFs()
	fileByte, _ := afero.ReadFile(appfs, "/Users/zhangjin/Project/user-center/ucenter-common/src/main/resources/com/yzf/ucenter/common/mapper/UserExtraInfoDAO.xml")
	fileString :=string(fileByte)
	fileLineMap := file.BuildFileLineMap(fileString)
	for line,content := range fileLineMap{
		fmt.Println("line: ", line, " content: ", content)
	}
}

//测试append
func TestAppend(t *testing.T)  {
	logs := []file.LogInfo{}
	logs = append(logs, file.LogInfo{"x123","123", "123"})
	logs = append(logs, file.LogInfo{"x123123123","123", "123123123"})
	var logsFinal = []file.LogInfo{}
	logsFinal = append(logsFinal, logs...)
	fmt.Println(logsFinal)
}

func TestFileReadDir(t *testing.T)  {
	// using the function
	mydir, _ := os.Getwd()
	fmt.Println(mydir)
	fmt.Println("regexp execute error")
}




