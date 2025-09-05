package babbler

import (
	"bufio"
	"os"
)

func GetText(file string) string {
	f, err := os.Open(file)
	defer f.Close()

	// just panic if we can't open the file
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(f)
	return Collect(r)
}

func Collect(r *bufio.Reader) string {
	contents := ""

	line, rerr := r.ReadString('\n')
	for rerr != nil {
		contents += line
		line, rerr = r.ReadString('\n')
	}
	return contents
}
