package main
import (
	"os"
	"fmt"
	"bytes"
	"regexp"
)

var simpl = []byte("simpl")

type line struct {
	 // if true then we are dealing with the home language
	 // if false then the markup language
	home bool
	// if home is true then content should be exactly what will be printed
	// this means you have some responsibility to fix up content
	content []byte 
}

func main() {
	fmt.Println("hello!!!!")
	file, _ := os.ReadFile("small.simpl")
	lines := []line{}
	lines = getLinesContent(file)
	addLinesType(lines)
	// for _,line := range lines {
	// 	fmt.Printf("%t: %s",line.home,line.content)
	// }
	p := printLines(lines)
	fmt.Printf("%s",p)
}

func getLinesContent(bs []byte) []line {
	linestart := 0
	lines := make([]line, 0, 50)
	for i := 0; i < len(bs); i++ {
		if bs[i] != '\n' {
			continue
		}
		lines = append(lines, line{content: bs[linestart:i+1]})
		linestart = i+1
	}
	return lines
}

func addLinesType(lines []line) {
	home := true
	for i := range lines {
		l := &lines[i]
		cnt := l.content
		if home == false {
			if cnt[0] == '}' {
				home = true
				l.home = true
			}  else if cnt[0] == '@' {
				l.home = true
			}
			continue
		}
		l.home = home
		ln := len(cnt) > len(simpl)
		if ln && bytes.Equal(cnt[0:len(simpl)], simpl) {
			copy(cnt[0:len(simpl)], "func ")
			home = false
		}
	}
}

var varregex = regexp.MustCompile(`([^$]|^)\$([\w().]+)`)

func printLines(lines []line) []byte {
	size := 0
	for _,l := range lines {
		size += len(l.content)
	}
	buf := make([]byte, 0, size)
	for _,l := range lines {
		if l.home {
			res := l.content
			if l.content[0] == '@' {
				res = res[1:]
			}
			buf = append(buf, res...)
		} else {
			buf = append(buf, "b = append(b,(\""...)

			c := l.content[:len(l.content)-1]
			r := varregex.ReplaceAll(c, []byte(`$1" + $2.String() + "`))
			buf = append(buf, r...)

			buf = append(buf, "\\n\")...)\n"...)
		}
	}
	return buf
}
