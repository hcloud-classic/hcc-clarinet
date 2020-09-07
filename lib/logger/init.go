package logger

import "hcc/clarinet/lib/errors"

// Init : Prepare logger
func Init() error {
	if !Prepare() {
		errors.SetErrLogger(Logger)
		return errors.NewHccError(errors.ClarinetInternalInitFail, "logger").New()
	}
	errors.SetErrLogger(Logger)

	return nil
}

// End : Close logger
func End() {
	_ = FpLog.Close()
}
