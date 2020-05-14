package errors

import (
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	// Reserved for command specific indications
	ErrorCommandSpecific = 1
	// Returned on connection failure to API endpoint
	ErrorConnectionFailure = 11
	// Returned on unexpected API response, i.e. authorization failure
	ErrorAPIResponse = 12
	// Generic error
	ErrorGeneric = 20
)

// CheckError is a convenience function to exit if an error is non-nil and exit if it was
func CheckErrorWithCode(err error, exitcode int) {
	if err != nil {
		Fatal(exitcode, err)
	}
}

// Wrapper for logrus.Fatal() to exit with custom code
func Fatal(exitcode int, args ...interface{}) {
	exitfunc := func() {
		os.Exit(exitcode)
	}
	log.RegisterExitHandler(exitfunc)
	log.Fatal(args...)
}

// Wrapper for logrus.Fatalf() to exit with custom code
func Fatalf(exitcode int, format string, args ...interface{}) {
	exitfunc := func() {
		os.Exit(exitcode)
	}
	log.RegisterExitHandler(exitfunc)
	log.Fatalf(format, args...)
}

// If err is not nil, logs a fatal message and exits with ErrorGeneric
func CheckError(err error) {
	Fatal(ErrorGeneric, err)
}

// panics if there is an error.
// This returns the first value so you can use it if you cast it:
// text := FailOrErr(Foo)).(string)
func FailOnErr(v interface{}, err error) interface{} {
	CheckError(err)
	return v
}
