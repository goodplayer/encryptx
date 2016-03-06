package main

import (
	"crypto/rand"
	"log"
)

import (
	"github.com/spf13/afero"
)

func main() {
	key := make([]byte, 256)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatalln(err)
	}

	var appfs afero.Fs = afero.NewOsFs()

	created, err := appfs.Create("keygen.key")

	if err != nil {
		log.Fatalln(err)
	}

	_, err = created.Write(key)
	if err != nil {
		log.Fatalln(err)
	}
	created.Close()
}
