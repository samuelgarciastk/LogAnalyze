package common

import (
	"bytes"
)

//CheckErr is a general method to handle error
//'messages' is an optional argument that stores the messages you want to show
func CheckErr(err error, messages ...string) {
	if err != nil {
		buf := bytes.NewBufferString(err.Error())
		for _, v := range messages {
			buf.WriteString("\n")
			buf.WriteString(v)
		}
		panic(buf.String())
	}
}
