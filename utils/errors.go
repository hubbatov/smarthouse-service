package utils

import (
	"time"
)

//CustomError (error)
type CustomError struct {
	When time.Time
	What string
}

//GenerateCustomError - generates custom error from error string
func GenerateCustomError(err string) *CustomError {
	return &CustomError{time.Now(), err}
}

//ConvertCustomError - generates custom error from error
func ConvertCustomError(err error) *CustomError {
	if err != nil {
		return &CustomError{time.Now(), err.Error()}
	}
	return nil
}

func handleError(err *CustomError) {
	if err != nil {
		panic(err.What)
	}
}
