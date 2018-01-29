package file

import (
	"os"
	"path"
	"strings"
	"fmt"
)

func WriteBytes(filePath string, b []byte) (int, error) {
	os.MkdirAll(path.Dir(filePath), os.ModePerm)
	fw, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer fw.Close()
	return fw.Write(b)
}

func WriteString(filePath string, s string) (int, error) {
	return WriteBytes(filePath, []byte(s))
}

func WriteLines(lines [] string,path string)(err error){
	var _file *os.File

	if _file,err = os.Create(path); err != nil{
		return
	}

	defer _file.Close()

	for _,elem := range lines {
		_,err := _file.WriteString(strings.TrimSpace(elem)+"\n")
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	return
}

