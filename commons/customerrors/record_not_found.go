package customerrors

import "fmt"

type RecordNotFound struct {
	id string
}

func (e RecordNotFound) Error() string {
	return fmt.Sprintf("database record not Found for given filter")
}
