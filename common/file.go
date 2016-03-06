package common

import (
	"io"
	"log"
)

import (
	"github.com/spf13/afero"
)

func ReadFullFile(filename string) ([]byte, error) {

	var appfs afero.Fs = afero.NewOsFs()

	fileinfo, err := appfs.Stat(filename)
	if err != nil {
		log.Fatalln(err)
	}
	fileLength := fileinfo.Size()

	file, err := appfs.Open(filename)

	if err != nil {
		log.Fatalln(err)
	}

	data := make([]byte, fileLength)

	_, err = io.ReadFull(file, data)
	if err != nil {
		log.Fatalln(err)
	}
	return data, nil
}
