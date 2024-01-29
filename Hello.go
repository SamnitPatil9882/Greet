package Hello

import "fmt"

func Hello(name string) string {
	msg := fmt.Sprintf("hii,%s  Welcome", name)
	return msg
}
