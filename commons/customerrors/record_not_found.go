package customerrors

<<<<<<< HEAD
import "fmt"

type RecordNotFound struct {
	id string
}

func (e RecordNotFound) Error() string {
	return fmt.Sprintf("database record not Found for given filter")
=======
type ResourceNotFound struct {
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (e ResourceNotFound) Error() string {
	return e.Message
}

func NewResourceNotFound(msg string) ResourceNotFound {
	if msg == "" {
		msg = "The requested resource was not found."
	}
	return ResourceNotFound{
		Message: msg,
	}
>>>>>>> cb617580f7e540b5109a595dbdc81d6aa6c40d39
}
