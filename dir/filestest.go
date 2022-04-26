package dir

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
	//dir, err := os.ReadDir(dirName)
	//if err != nil {
	//	println(err)
	//}

	//dir, err := os.ReadDir(dirName)
	//if err != nil {
	//	log.Panicln(err)
	//}
	//for i, file := range dir {
	//	ioutil.ReadAll(file)
	//}

	//dir, err := ioutil.ReadDir("/home/andrii/docs")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, file := range dir {
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
