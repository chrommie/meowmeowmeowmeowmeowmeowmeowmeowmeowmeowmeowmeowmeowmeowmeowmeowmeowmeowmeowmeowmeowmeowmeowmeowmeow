package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const separator string = "​"

func main() {
	tmp := "print('Hello World!');"

	tmp = encode(tmp)
	writeToFile("encoded.txt", tmp)
	tmp = decode(tmp)

	fmt.Println(runPython(tmp))
}

func writeToFile(fileName string, data string) {
	f, err := os.Create(fileName)
	if err != nil {
		panic(any(err))
	}
	defer f.Close()
	f.WriteString(data)

}

func runPython(code string) string {
	out, err := exec.Command("python", "-c", fmt.Sprint(code)).Output()
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
}

func encode(s string) string {
	// for character in s, get its ascii value and print it
	var tmp = ""
	for _, c := range s {
		for i := 0; i < int(c); i++ {
			tmp += "meow"
		}
		tmp += separator

	}
	return tmp
}

func decode(s string) string {
	var tmp = ""
	// do the reverse of encode
	for _, c := range strings.Split(s, separator) {
		tmp += string(strings.Count(c, "meow"))
	}
	tmp = strings.Replace(tmp, separator, "", -1)
	tmp = string(bytes.Trim([]byte(tmp), "\x00"))
	return tmp
}
