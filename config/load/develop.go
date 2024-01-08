package load

func Develop(path string, dst any) error {
	content, err := loadFile(path)
	if err != nil {
		return err
	}

	return unmarshal(content, dst)
}
