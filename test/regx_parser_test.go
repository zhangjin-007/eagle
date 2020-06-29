package test

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestSubQuery(t *testing.T) {
	query :="select a from table1 where id in select b from table b)"
	matched, _ := regexp.MatchString(`\([\s\S]*select[\s\S]*\)`, query)
	assert.Equal(t, matched, true)
}