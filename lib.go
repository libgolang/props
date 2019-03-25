package props

import (
	"os"
	"regexp"
)

type propsStruct struct {
	props map[string]string
	args  []string
}

var (
	globalProps   *propsStruct
	propEqRegex   = regexp.MustCompile(`--([\w\-]+)=(.*)$`)
	propRegex     = regexp.MustCompile(`--([\w\-]+)`)
	propFlagRegex = regexp.MustCompile(`-(\w+)`)
)

func init() {
	globalProps = parseArgs(os.Args[1:])
}

// GetProp returns the value of a property, if not set it returns
// an empty string
func GetProp(name string) string {
	if val, found := globalProps.props[name]; found {
		return val
	}
	return ""
}

// IsSet returns true if property is set, otherwise it returns false
func IsSet(name string) bool {
	_, found := globalProps.props[name]
	return found
}

// Regurns any argument that is not a property name, flag or property value
func GetArgs() []string {
	return globalProps.args
}

func parseArgs(args []string) *propsStruct {
	res := &propsStruct{
		props: make(map[string]string),
		args:  make([]string, 0),
	}

	numArgs := len(args)
	for i := 0; i < numArgs; i++ {
		arg := args[i]
		//for _, arg := range args {
		if arr := propEqRegex.FindStringSubmatch(arg); arr != nil {
			// prop: --some-prop=xyz
			res.props[arr[1]] = arr[2]
		} else if arr := propRegex.FindStringSubmatch(arg); arr != nil {
			// prop: --some-prop xyz
			if i < (numArgs - 1) {
				nextArg := args[i+1]
				if !propEqRegex.MatchString(nextArg) && !propRegex.MatchString(nextArg) {
					// next argument is not a property/flag --prop-name
					res.props[arr[1]] = args[i+1]
					i++
				} else {
					// if next argument is a property/flag, then set to empty
					res.props[arr[1]] = ""
				}
			} else {
				// property is the last parameter, set to empty
				res.props[arr[1]] = ""
			}
		} else if arr := propFlagRegex.FindStringSubmatch(arg); arr != nil {
			// flag -p, -a -x
			flagsArr := []rune(arr[1]) // -abc : -a -b -c
			for i := 0; i < len(flagsArr); i++ {
				letter := string(flagsArr[i])
				res.props[letter] = ""
			}
		} else {
			// arg
			res.args = append(res.args, arg)
		}
	}
	return res
}
