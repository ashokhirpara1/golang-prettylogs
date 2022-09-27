// Package prettylogs uses concept of structured logging and the methods to default add structure to log entry payload fields.
//
// The prettylogs package can be used to log entries that use the jsonPayload field to add structure to their payloads.
// By default it logs "level", "method", "msg", "time" json fields.
package prettylogs

import (
	"fmt"
	"regexp"
	"runtime"
	"time"

	"github.com/ashokhirpara1/golang-prettylogs/logger"
	"github.com/sirupsen/logrus"
)

var logs logger.Storage

// init initialises logger at a global level
func init() {
	logs = logger.New()
}

// Info accepts the string message
// It logs the info level message with method name and timestamp
func Info(str string) {

	// Skip this function, and fetch the PC and file for its parent
	pc, _, _, _ := runtime.Caller(1)

	// Retrieve a Function object this functions parent
	functionObject := runtime.FuncForPC(pc)

	// Regex to extract just the function name (and not the module path)
	extractFnName := regexp.MustCompile(`^.*\.(.*)$`)
	method := extractFnName.ReplaceAllString(functionObject.Name(), "$1")

	fields := logrus.Fields{
		"method": method,
	}

	logs.Info(fields, str)
}

// Error accepts the string message and error
// It logs the error level message with method name and timestamp
func Error(msg string, errr error) {
	// Skip this function, and fetch the PC and file for its parent
	pc, _, _, _ := runtime.Caller(1)

	// Retrieve a Function object this functions parent
	functionObject := runtime.FuncForPC(pc)

	// Regex to extract just the function name (and not the module path)
	extractFnName := regexp.MustCompile(`^.*\.(.*)$`)
	method := extractFnName.ReplaceAllString(functionObject.Name(), "$1")

	fields := logrus.Fields{
		"method": method,
	}

	msg = msg + ": " + errr.Error()
	logs.Error(fields, msg)
}

// Fatal accepts the method name, string message and error
// It logs the message and exits the program
func Fatal(method string, msg string, errr error) {
	fields := logrus.Fields{
		"method": method,
	}

	msg = msg + ": " + errr.Error()
	logs.Fatal(fields, msg)
}

// Enter logs the start time of the method
func Enter() (time.Time, string) {
	start := time.Now()

	// Skip this function, and fetch the PC and file for its parent
	pc, _, _, _ := runtime.Caller(1)

	// Retrieve a Function object this functions parent
	functionObject := runtime.FuncForPC(pc)

	// Regex to extract just the function name (and not the module path)
	extractFnName := regexp.MustCompile(`^.*\.(.*)$`)
	method := extractFnName.ReplaceAllString(functionObject.Name(), "$1")

	fields := logrus.Fields{
		"method": method,
	}

	msg := "Started " + method
	logs.Info(fields, msg)

	return start, method
}

// Exit mesures the execution time of the method
func Exit(start time.Time, name string) {
	elapsed := time.Since(start)
	fields := logrus.Fields{
		"method": name,
	}
	msg := fmt.Sprintf("%s took %s", name, elapsed)
	logs.Info(fields, msg)
}
