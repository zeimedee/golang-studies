package main

import (
	. "fmt"
	"os/exec"
)

func main() {
	// app := "go version"

	// arg0 := "./"
	// arg1 := "hi.go"
	// arg2 := "\n\tfrom"
	// arg3 := "golang"

	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	cmd := exec.Command("go", "run", "./readfile.go")
	stdout, err := cmd.Output()

	if err != nil {
		Println(err.Error())
		return
	}

	Print(string(stdout))
}
