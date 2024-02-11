package main

import (
	"fmt"
	"os"
)

func main() {
	total := 10000

	for i := 0; i < total; i++ {
		file, err := os.Create(fmt.Sprintf("./tmp/file%d.txt", i))

		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()

		/* Conteúdo do arquivo */
		file.WriteString("ok")
	}
}
