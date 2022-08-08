package file

import "os"

// WriteFile filename可以是绝对路径，也可以是相对路径，例如：./log.txt
func WriteFile(filename string, line string) error {
	var f *os.File
	var err error
	if isFileExist(filename) {
		f, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0666) // os.FileMode(0666).String() returns "-rw-rw-rw-"
	} else {
		f, err = os.Create(filename)
	}

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(line + "\n")
	if err != nil {
		return err
	}

	if err = f.Sync(); err != nil {
		return nil
	}

	return nil
}

func isFileExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
