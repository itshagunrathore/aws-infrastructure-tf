package customerrors

import "fmt"

type JobAlreadyExistsError struct {
	JobName   string
	AccountId string
}

func (e JobAlreadyExistsError) Error() string {
	return fmt.Sprintf("job with Name: %s already exits for accountId: %s", e.JobName, e.AccountId)
}
