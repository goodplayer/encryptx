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
	connStr  string
)

func init() {
	flag.StringVar(&filename, "file", ".", "-file=abc.xyz")
	flag.StringVar(&keyfile, "key", ".", "-key=defg.key")
	flag.StringVar(&connStr, "conn", "127.0.0.1:15433", "-conn=127.0.0.1:15433")
}

func main() {
	flag.Parse()

	key, err := common.ReadKeyFile(keyfile)
	if err != nil {
		log.Fatalln(err)
	}

	var appfs afero.Fs = afero.NewOsFs()

	created, err := appfs.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := dealConn()
	if err != nil {
		log.Fatalln(err)
	}

	common.AesDecryptCopy(created, conn, key)

	err = conn.Close()
	if err != nil {
		log.Fatalln(err)
	}

	err = created.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

func dealConn() (*net.TCPConn, error) {
	addr, err := net.ResolveTCPAddr("tcp", connStr)
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatalln(err)
	}

	return conn, nil
}
