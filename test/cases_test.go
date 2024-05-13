package test

import (
	unpack "help_me_unpack/src"
)

type unpackTest struct {
	description string
	subject     string
	expected    unpack.UnpackedBytes
}

var testCases = []unpackTest{
	{
		description: "valid",
		subject:     "ln8YhXNgYcqVyAAAcmoHQx7oia9MF1VAQFUXTK+J6B4=",
		expected:    unpack.UnpackedBytes{},
	},
	{
		description: "incorrect length",
		subject:     "R4usg4PUzcVJAQFJxM+Vw5yI=",
		expected:    unpack.UnpackedBytes{},
	},
	{
		description: "empty",
		subject:     "",
		expected:    unpack.UnpackedBytes{},
	},
}
