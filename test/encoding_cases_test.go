package test

type encodingType int

const (
	ENCODE encodingType = iota
	DECODE
)

type encodingTest struct {
	encodingType encodingType
	description  string
	subject      string
	expected     string
}

var encodingTestCases = []encodingTest{
	{
		encodingType: ENCODE,
		description:  "encode",
		subject:      "hello,world!",
		expected:     "aGVsbG8sd29ybGQh",
	},
	{
		encodingType: DECODE,
		description:  "decode",
		subject:      "aGVsbG8sd29ybGQh",
		expected:     "hello,world!",
	},
}
