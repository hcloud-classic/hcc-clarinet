package argumentParser

import (
	"errors"
	"strconv"
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

func GenIntArg(arguments *string, args map[string]int, fn func(int) (bool, error)) error {
	var err error = nil
	var checkF func(int) (bool, error)

	if fn == nil {
		checkF = func(arg int) (bool, error) {
			if arg > 0 {
				return true, nil
			} else if arg == 0 {
				return false, nil
			}
			return false, errors.New("")
		}
	} else {
		checkF = fn
	}

	for key, arg := range args {
		if b, e := checkF(arg); b && e == nil {
			*arguments += key + ": " + strconv.Itoa(arg) + ", "
		} else if e != nil {
			err = errors.New("Check flag value of " + key)
			break
		}
	}

	return err
}

func GenStrArg(arguments *string, args map[string]string, fn func(string) (bool, error)) error {
	var err error = nil
	var checkF func(string) (bool, error)

	if fn == nil {
		checkF = func(arg string) (bool, error) {
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

func GetArgumentStr(strArgs map[string]string, fn ...interface{}) (string, error) {

	intArgs := SplitIntArgs(strArgs)
	var arguments = ""
	var fnStr func(string) (bool, error)
	var fnInt func(int) (bool, error)

	if len(fn) == 2 {
		fnStr = fn[0].(func(string) (bool, error))
		fnInt = fn[1].(func(int) (bool, error))
	} else if len(fn) == 1 {
		fnStr = fn[0].(func(string) (bool, error))
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
