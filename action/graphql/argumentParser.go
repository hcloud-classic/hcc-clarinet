package argumentParser

import (
	"strconv"

	errors "github.com/hcloudclassic/hcc_errors"
)

func CheckArgsMin(args map[string]string, min int, mustchk ...string) bool {
	for key, value := range args {
		if value == "" || value == "0" {
			delete(args, key)
		}
	}

	for _, key := range mustchk {
		if _, exist := args[key]; !exist {
			return false
		}
	}

	return len(args) < min
}

func CheckArgsAll(args map[string]string, max int, exception ...string) (bool, string) {
	var emptyField string = ""

	for key, value := range args {
		for _, e := range exception {
			if e == key {
				goto CONTINUE
			}
		}
		if value == "" || value == "0" {
			emptyField += (key + " ")
			delete(args, key)
		}
	CONTINUE:
	}

	return len(args) < max, emptyField
}

func SplitIntArgs(strArgs map[string]string) map[string]int {
	intArgs := map[string]int{}
	for key, value := range strArgs {
		if i, err := strconv.Atoi(value); err == nil {
			intArgs[key] = i
			delete(strArgs, key)
		}
	}
	return intArgs
}

func GenIntArg(arguments *string, args map[string]int, fn func(int) (bool, *errors.HccError)) *errors.HccError {
	var err *errors.HccError = nil
	var checkF func(int) (bool, *errors.HccError)

	if fn == nil {
		checkF = func(arg int) (bool, *errors.HccError) {
			if arg > 0 {
				return true, nil
			} else if arg == 0 {
				return false, nil
			}
			return false, errors.NewHccError(errors.ClarinetGraphQLParsingError, "Integer argument ")
		}
	} else {
		checkF = fn
	}

	for key, arg := range args {
		if b, e := checkF(arg); b && e == nil {
			*arguments += key + ": " + strconv.Itoa(arg) + ", "
		} else if e != nil {
			err = errors.NewHccError(e.Code(), e.Text()+key)
			break
		}
	}

	return err
}

func GenStrArg(arguments *string, args map[string]string, fn func(string) (bool, *errors.HccError)) *errors.HccError {
	var err *errors.HccError = nil
	var checkF func(string) (bool, *errors.HccError)

	if fn == nil {
		checkF = func(arg string) (bool, *errors.HccError) {
			if arg == "" {
				return false, nil
			}
			return true, nil
		}
	} else {
		checkF = fn
	}

	for key, arg := range args {
		if b, _ := checkF(arg); b {
			*arguments += key + ": \"" + arg + "\", "
		}
	}
	return err
}

func GetArgumentStr(strArgs map[string]string, fn ...interface{}) (string, *errors.HccError) {

	intArgs := SplitIntArgs(strArgs)
	var arguments = ""
	var fnStr func(string) (bool, *errors.HccError)
	var fnInt func(int) (bool, *errors.HccError)

	if len(fn) == 2 {
		fnStr = fn[0].(func(string) (bool, *errors.HccError))
		fnInt = fn[1].(func(int) (bool, *errors.HccError))
	} else if len(fn) == 1 {
		fnStr = fn[0].(func(string) (bool, *errors.HccError))
		fnInt = nil
	} else {
		fnStr = nil
		fnInt = nil
	}

	err := GenStrArg(&arguments, strArgs, fnStr)
	if err != nil {
		return arguments, err
	}

	err = GenIntArg(&arguments, intArgs, fnInt)
	if err != nil {
		return arguments, err
	}

	if arguments != "" {
		arguments = "(" + arguments + ")"
	}
	return arguments, nil
}
