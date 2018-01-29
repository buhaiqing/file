package file

import (
	"os"
	"path"
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

func writeLines(lines [] string,path string)(err error){
	var file *os.File

	if file,err = os.Create(path); err != nil{
		return
	}

	defer file.Close()

	for _,elem := range lines {
		_,err := file.WriteString(strings.TrimSpace(elem)+"\n")
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	return
}

