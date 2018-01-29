package file

import (
	"bufio"
	"io/ioutil"
	"strconv"
	"strings"
)

func ToBytes(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

func ToString(filePath string) (string, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ToTrimString(filePath string) (string, error) {
	str, err := ToString(filePath)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(str), nil
}

func ToUint64(filePath string) (uint64, error) {
	content, err := ToTrimString(filePath)
	if err != nil {
		return 0, err
	}

	var ret uint64
	if ret, err = strconv.ParseUint(content, 10, 64); err != nil {
		return 0, err
	}
	return ret, nil
}

func ToInt64(filePath string) (int64, error) {
	content, err := ToTrimString(filePath)
	if err != nil {
		return 0, err
	}

	var ret int64
	if ret, err = strconv.ParseInt(content, 10, 64); err != nil {
		return 0, err
	}
	return ret, nil
}

func ReadLine(r *bufio.Reader) ([]byte, error) {
	line, isPrefix, err := r.ReadLine()
	for isPrefix && err == nil {
		var bs []byte
		bs, isPrefix, err = r.ReadLine()
		line = append(line, bs...)
	}

	return line, err
}

func readLines(path string) (lines [] string, err error) {
	var (
		file   *os.File
		part   [] byte
		prefix bool
	)

	if file, err = os.Open(path); err != nil {
		return
	}

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 1024))

	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

