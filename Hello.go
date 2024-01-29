package Hello

import "fmt"

func Hello(name string) string {
	msg := fmt.Sprint("hii,%s  Welcome", name)
	return msg
}
