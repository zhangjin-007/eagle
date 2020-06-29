package test

import (
	"eagle/file"
	"fmt"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"regexp"
	"strings"
	"testing"
)

func TestSubQuery(t *testing.T) {
	query :="select a from table1 where id in select b from table b)"
	matched, _ := regexp.MatchString(`\([\s\S]*select[\s\S]*\)`, query)
	assert.Equal(t, matched, true)
}

func TestAllSubSquery(t *testing.T) {
	appfs := afero.NewOsFs()
	fileName := "/Users/zhangjin/Project/eagle/test/src/RptInfoMapper.xml"
	fileByte, _ := afero.ReadFile(appfs, fileName)
	fileString :=string(fileByte)
	subQueryList := file.FileFindAllSubqueryString(fileString)
	fmt.Println(strings.Join(subQueryList[:], "\n"))
}