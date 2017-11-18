package mason

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestCreateGoPath(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "gomason")
	if err != nil {
		log.Printf("Error creating temp dir\n")
		t.Fail()
	}

	defer os.RemoveAll(tmpDir)

	_, err = CreateGoPath(tmpDir)
	if err != nil {
		log.Printf("Error creating gopath in %q: %s", tmpDir, err)
		t.Fail()
	}

	dirs := []string{"go", "go/src", "go/pkg", "go/bin"}

	for _, dir := range dirs {
		if _, err := os.Stat(fmt.Sprintf("%s/%s", tmpDir, dir)); os.IsNotExist(err) {
			t.Fail()
		}

	}

}

func TestReadMetadata(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "gomason")
	if err != nil {
		log.Printf("Error creating temp dir\n")
		t.Fail()
	}

	defer os.RemoveAll(tmpDir)

	fileName := fmt.Sprintf("%s/%s", tmpDir, TestMetadataFileName())

	err = ioutil.WriteFile(fileName, []byte(TestMetaDataJson()), 0644)
	if err != nil {
		log.Printf("Error writing metadata file: %s", err)
		t.Fail()
	}

	expected := TestMetadataObj()

	actual, err := ReadMetadata(fileName)
	if err != nil {
		log.Printf("Error reading metadata from file: %s", err)
		t.Fail()
	}

	assert.Equal(t, expected, actual, "Generated metadata object meets expectations.")

}

func TestGitSSHUrlFromPackage(t *testing.T) {
	input := "github.com/nikogura/gomason"
	expected := "git@github.com:nikogura/gomason.git"

	assert.Equal(t, expected, GitSSHUrlFromPackage(input), "Git SSH URL from Package Name meets expectations.")
}