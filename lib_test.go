package props

import (
	"testing"
)

func TestFlagRegex(t *testing.T) {
	// given
	given := "-abc"

	// when
	arr := propFlagRegex.FindStringSubmatch(given)

	// then
	if len(arr) < 2 {
		t.Error("propFlagRegex match should contain two elements")
	}
}

func TestGetProp(t *testing.T) {
	// given
	args := []string{
		"-abc",
		"-x",
		"-y",
		"--bool",
		"--prop",
		"prop-val",
		"arg1",
		"--prop-eq1=prop-eq-val1",
		"arg2",
		"--prop-eq2=prop-eq-val2",
		"--prop-last",
	}

	// when
	globalProps = parseArgs(args)

	// then
	for _, v := range []string{"a", "b", "c", "x", "y", "bool", "prop", "prop-eq1", "prop-eq2", "prop-last"} {
		if !IsSet(v) {
			t.Errorf("flag `%s` must be set", v)
		}
	}

	for k, v := range map[string]string{"prop": "prop-val"} {
		if GetProp(k) != v {
			t.Errorf("expected `--%[1]s=%[2]s`, but got `--%[1]s=%[3]s`", k, v, GetProp(k))
		}
	}
}
