package utils

import (
	"os"
	"regexp"
)

const projectDirName = "internet_service"

func getProjectName() *regexp.Regexp {
	return regexp.MustCompile(`^(.*` + projectDirName + `)`)
}

func getCurrentWorkDirectory() string {
	currentWorkDirectory, err := os.Getwd()
	if err != nil {
		// TODO log system.
	}
	return currentWorkDirectory
}

func getRootPath() []byte {
	projectName := getProjectName()
	currentWorkDirectory := getCurrentWorkDirectory()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	return rootPath
}

func GetPathAsBytes() []byte {
	return getRootPath()
}

// GetPathAsString pass current root project directory path,
// it has no `/` in end of path
func GetPathAsString() string {
	return string(getRootPath())
}
