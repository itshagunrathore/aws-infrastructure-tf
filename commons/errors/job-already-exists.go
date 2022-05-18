package errors

import "fmt"

type JobAlreadyExistsError struct {
	jobName   string
	accountId string
}

func (e JobAlreadyExistsError) Error() string {
	return fmt.Sprintf("job with Name: %s already exits for accountId: %s", e.jobName, e.accountId)
}
