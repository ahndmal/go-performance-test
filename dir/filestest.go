package dir

import (
	"fmt"
	"io/ioutil"
)

func testFiles() {
	for i := 1; i < 200; i++ {
		file, err := ioutil.ReadFile(fmt.Sprintf("/home/andrii/docs/file%d.txt", i))
		fmt.Println(string(file))
		if err != nil {
			fmt.Println(err)
		}
	}

	//dir, err := ioutil.ReadDir("/home/andrii/docs")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, file := range dir {
	//	fmt.Println(file)
	//}
}
