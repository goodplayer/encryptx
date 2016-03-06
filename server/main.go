package main

import (
	"flag"
	"log"
	"net"
)

import (
	"github.com/spf13/afero"

	"github.com/goodplayer/encryptx/common"
)

var (
	filename string
	keyfile  string
)

func init() {
	flag.StringVar(&filename, "file", ".", "-file=abc.xyz")
	flag.StringVar(&keyfile, "key", ".", "-key=defg.key")
}

func main() {
	flag.Parse()

	key, err := common.ReadKeyFile(keyfile)
	if err != nil {
		log.Fatalln(err)
	}

	var (
		fileLength  int64
		dstFileName string
	)

	var appfs afero.Fs = afero.NewOsFs()

	fileinfo, err := appfs.Stat(filename)
	if err != nil {
		log.Fatalln(err)
	}
	fileLength = fileinfo.Size()
	dstFileName = fileinfo.Name()

	file, err := appfs.Open(filename)

	if err != nil {
		log.Fatalln(err)
	}

	conn, err := dealIncome()
	if err != nil {
		log.Fatalln(err)
	}

	common.AesEncryptCopy(file, conn, key)

	var _ = fileLength
	var _ = dstFileName

	err = conn.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

func dealIncome() (*net.TCPConn, error) {
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:15433")
	if err != nil {
		log.Fatalln(err)
	}
	li, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	conn, err := li.AcceptTCP()
	if err != nil {
		log.Fatalln(err)
	}
	return conn, nil
}
