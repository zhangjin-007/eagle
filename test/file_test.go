package test

import (
	"fmt"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

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
		}
		return nil
	})
	assert.Nil(t, err)
}