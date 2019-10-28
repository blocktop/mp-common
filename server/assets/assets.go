package assets

func GetString(path string) (string, error) {
	return FSString(false, path)
}
