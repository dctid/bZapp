package test

import (
	"fmt"
	"github.com/dctid/bZapp/format"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"
)

func ReadFile(t *testing.T, name string) string {
	path, err := filepath.Abs("./")
	assert.NoError(t, err)
	log.Printf("Path %s", path)
	dat, err := ioutil.ReadFile(fmt.Sprintf("%s/test/data/%s", path, name))
	assert.NoError(t, err)

	prettyJson := format.PrettyJson(t, string(dat))
	assert.NoError(t, err)
	return prettyJson
}
