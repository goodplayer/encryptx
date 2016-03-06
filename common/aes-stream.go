package common

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
	"log"
)

func AesEncryptCopy(r io.Reader, w io.Writer, key []byte) error {
	block, err := aes.NewCipher(key[:32])
	if err != nil {
		log.Fatalln(err)
	}

	s := cipher.NewOFB(block, key[:aes.BlockSize])

	writer := &cipher.StreamWriter{S: s, W: w}

	if _, err := io.Copy(writer, r); err != nil {
		log.Fatalln(err)
	}
	return nil
}

func AesDecryptCopy(w io.Writer, r io.Reader, key []byte) error {
	block, err := aes.NewCipher(key[:32])
	if err != nil {
		log.Fatalln(err)
	}

	stream := cipher.NewOFB(block, key[:aes.BlockSize])

	reader := &cipher.StreamReader{S: stream, R: r}

	if _, err := io.Copy(w, reader); err != nil {
		log.Fatalln(err)
	}

	return nil
}
