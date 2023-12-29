package reader

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func GetFileInput(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = file.Close() }()
	r := bufio.NewReader(file)
	var bytes []byte
	var lines []string
	for {
		line, isPrefix, err := r.ReadLine()
		if err != nil {
			break
		}
		bytes = append(bytes, line...)
		if !isPrefix {
			str := strings.TrimSpace(string(bytes))
			if len(str) > 0 {
				lines = append(lines, str)
				bytes = []byte{}
			}
		}
	}
	if len(bytes) > 0 {
		lines = append(lines, string(bytes))
	}
	return lines
}
