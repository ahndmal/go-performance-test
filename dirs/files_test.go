package dirs

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func CountWordsByText(text string) int {
	var count int
	for i, _ := range strings.Split(text, " ") {
		count += i
	}
	return count
}

func countWordsByPath(path string) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		println(err)
	}
	count := 0
	for _, w := range strings.Split(string(fileBytes), " ") {
		println(w)
		count += 1
	}
	fmt.Println(count)
}

func TestFiles(t *testing.T) {
	dirName := "/home/andrii/Documents"
	//lorem1 := "/home/andrii/Documents/lorem1.txt"

	dirs, err := os.ReadDir(dirName)
	if err != nil {
		log.Panicln(err)
	}
	for _, dir := range dirs {
		if !dir.IsDir() {
			readFile, err := os.ReadFile(dir.Name())
			if err != nil {
				return
			}
			fmt.Println(readFile)
		}
	}

	//dirs, err := ioutil.ReadDir("/home/andrii/docs")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, file := range dirs {
	//	fmt.Println(file)
	//}
}

func TestGetWholeText(t *testing.T) {
	wholeText := GetWholeText()
	words := strings.Split(wholeText, " ")
	println(len(words))
}

func GetWholeText() string {
	var initStr string
	for i := 1; i < 500; i++ {
		file, err := os.ReadFile(fmt.Sprintf("/home/andrii/Documents/lorem%d.txt", i))
		initStr += string(file)
		if err != nil {
			fmt.Println(err)
		}
	}
	return initStr
}
