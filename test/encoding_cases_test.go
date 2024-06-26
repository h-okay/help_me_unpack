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
	expected     []byte
}

func (et encodingType) String() string {
	switch et {
	case ENCODE:
		return "Encode"
	case DECODE:
		return "Decode"
	default:
		return "Unknown"
	}
}

var encodingTestCases = []encodingTest{
	{
		encodingType: ENCODE,
		description:  "encode",
		subject:      "hello,world!",
		expected:     []byte("aGVsbG8sd29ybGQh"),
	},
	{
		encodingType: DECODE,
		description:  "decode",
		subject:      "aGVsbG8sd29ybGQh",
		expected:     []byte("hello,world!"),
	},
}
