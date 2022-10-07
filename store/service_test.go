package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreservice = &StorageService{}

func init() {
	testStoreservice = IntializeConnection()
}

func TestStore(t *testing.T) {
	assert.True(t, testStoreservice.Client != nil)
}

func TestSetGetUrl(t *testing.T) {
	originalurl := "https://www.goodstudent.com/college/department/rollno/name/age"
	generatedurl := "Ghjrywuxn"

	SaveUrl(generatedurl, originalurl)

	geturl := GetUrl(generatedurl)

	assert.Equal(t, originalurl, geturl)

}
		