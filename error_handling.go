package main
import "C"

// Workaround for error handling.

// errMessage stores the message of the last encountered error.
// If no error was encountered it is "".
var errMessage string

// setErr stores the message of the given error in errMessage,
// so it can be checked from the wrapper C code.
func setErr(err error) {
	errMessage = err.Error()
}

// checkErr returns the message of the last encountered error.
// It is called from the C wrapper code.
//export checkErr
func checkErr() *C.char {
	return C.CString(errMessage)
}

// clearErr clears errMessage.
//export clearErr
func clearErr() {
	errMessage = ""
}

