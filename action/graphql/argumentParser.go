package argumentParser

import (
	"errors"
	"strconv"
)

func CheckArgsMin(args map[string]string, min int) bool {
	for key, value := range args {
		if value == "" || value == "0" {
			delete(args, key)
		}
	}

	return len(args) < min
}

func CheckArgsAll(args map[string]string, max int) (bool, string) {
	var emptyField string = ""

	for key, value := range args {
		if value == "" || value == "0" {
			emptyField += (key + " ")
			delete(args, key)
		}
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

func GetArgumentStr(strArgs map[string]string) (string, error) {

	intArgs := SplitIntArgs(strArgs)
	var arguments = ""

	err := GenStrArg(&arguments, strArgs, nil)
	if err != nil {
		return arguments, err
	}

	err = GenIntArg(&arguments, intArgs, nil)
	if err != nil {
		return arguments, err
	}

	if arguments != "" {
		arguments = "(" + arguments + ")"
	}
	return arguments, nil
}
