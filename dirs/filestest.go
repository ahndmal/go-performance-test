package dirs

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func CountWords(text string) int {
	var count int
	for i, _ := range strings.Split(text, " ") {
		count += i
	}
	return count
}

func TestFiles() {
	//dirName := "/home/andrii/docs"
	//dirs, err := os.ReadDir(dirName)
	//if err != nil {
	//	println(err)
	//}

	//dirs, err := os.ReadDir(dirName)
	//if err != nil {
	//	log.Panicln(err)
	//}
	//for i, file := range dirs {
	//	ioutil.ReadAll(file)
	//}

	//dirs, err := ioutil.ReadDir("/home/andrii/docs")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, file := range dirs {
	//	fmt.Println(file)
	//}
}

func GetWholeText() string {
	var initStr string
	for i := 1; i < 500; i++ {
		file, err := ioutil.ReadFile(fmt.Sprintf("/home/andrii/docs/file%d.txt", i))
		initStr += string(file)
		if err != nil {
			fmt.Println(err)
		}
	}
	return initStr
}
