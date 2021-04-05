package handlers

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func CreateDb() {
	pwd, err := os.Getwd()
	fpCurrentDir := filepath.Join(pwd)
	files, err := ioutil.ReadDir(fpCurrentDir)
	var dbExists = false
	for _, f := range files {
		if f.Name() == "db" {
			dbExists = true
		}
	}

	fpDbDir := filepath.Join(pwd, "db")
	if !dbExists {
		os.Mkdir(fpDbDir, os.FileMode(0777))
	}

	filesFromDb, err := ioutil.ReadDir(fpDbDir)
	Check(err)
	if len(filesFromDb) == 0 {
		fpDbFile := filepath.Join(pwd, "db", "coffeeorders.json")
		emptyArray := []byte("[]")
		err := ioutil.WriteFile(fpDbFile, emptyArray, 0777)
		Check(err)
	}
}

func DeleteDb() {
	// first lets make sure the db folder exists, so that we don't crash the application if it got deleted.
	pwd, err := os.Getwd()
	fpCurrentDir := filepath.Join(pwd)
	files, err := ioutil.ReadDir(fpCurrentDir)
	var dbExists = false
	for _, f := range files {
		if f.Name() == "db" {
			dbExists = true
		}
	}

	fpDbDir := filepath.Join(pwd, "db")
	if !dbExists {
		os.Mkdir(fpDbDir, os.FileMode(0777))
	}

	fpDbFile := filepath.Join(pwd, "db", "coffeeorders.json")
	emptyArray := []byte("[]")
	err = ioutil.WriteFile(fpDbFile, emptyArray, 0777)
	Check(err)
}
