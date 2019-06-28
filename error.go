====================================
package builtin
type error interface {
	Error() string
}


====================================
package errors

// New returns an error that formats as the given text.
func New(text string) error {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}



============================================

var errDBClosed = errors.New("sql: database is closed")
func main() {
	
}
