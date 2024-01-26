package main
import (
	"os"
	"fmt"
	"bytes"
	"path"
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


/*
cli: 

simpgo file.simp # takes file.simp and produces file_simp.go
simpgo file1.simp file2.simp file3.simp # does the same as above but for many files

*/

func changefilename(s string) string {
	i := len(s) -1
	for ; i > 0; i-- {
		if s[i] == '.' {
			break
		}
	}
	return s[:i] + "_simp.go"
}

func main() {
	args := os.Args[1:]

	for _, p := range args {
		ext := path.Ext(p)
		if ext != ".simp" {
			fmt.Println("warning,", p, "does not have the right file extension '.simp'")
			continue
		}
		file, err := os.ReadFile(p)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		res := transform(file)
		nn := changefilename(p)
		out, err := os.Create(nn)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		out.Write(res)
		out.Sync()

	}
	return 
}

func transform(bs []byte) []byte {
	lines := []line{}
	lines = getLinesContent(bs)
	addLinesType(lines)
	return printLines(lines)
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

// TODO: need to escape double $ properly
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
