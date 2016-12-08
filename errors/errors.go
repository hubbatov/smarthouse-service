package errors

import (
	"fmt"
	"time"
)

//CustomError is an error with time
type CustomError struct {
	When time.Time
	What string
}

//GenerateCustomError generates custom error from error string
func GenerateCustomError(err string) *CustomError {
	return &CustomError{time.Now(), err}
}

//ConvertCustomError generates custom error from error
func ConvertCustomError(err error) *CustomError {
	if err != nil {
		return &CustomError{time.Now(), err.Error()}
	}
	return nil
}

//HandleError prints error in console
func HandleError(err *CustomError) {
	if err != nil {
		fmt.Println("Error: ", err.What, err.When)
	}
}
