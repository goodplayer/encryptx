package common

func ReadKeyFile(keyfile string) ([]byte, error) {
	return ReadFullFile(keyfile)
}
