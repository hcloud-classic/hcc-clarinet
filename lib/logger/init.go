package logger

import errors "github.com/hcloudclassic/hcc_errors"

// Init : Prepare logger
func Init() error {
	if !Prepare() {
		errors.SetErrLogger(Logger)
		return errors.NewHccError(errors.ClarinetInternalInitFail, "logger").ToError()
	}
	errors.SetErrLogger(Logger)

	return nil
}

// End : Close logger
func End() {
	_ = FpLog.Close()
}
