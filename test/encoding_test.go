package test

import (
	encoding "help_me_unpack/src"
	"testing"
)

func TestEncoding(t *testing.T) {
	for _, tc := range encodingTestCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := runTest(t, tc.encodingType, tc.subject)
			if string(tc.expected) != string(actual) {
				t.Errorf("%s(\"%s\") = %#v, want: %#v", tc.encodingType.String(), tc.subject, actual, tc.expected)
			}
		})
	}
}

func runTest(t *testing.T, encodingType encodingType, subject []byte) []byte {
	var actual []byte
	var err error

	switch encodingType {
	case ENCODE:
		actual, err = encoding.Encode(subject)
	case DECODE:
		actual, err = encoding.Decode(subject)
	}

	if err != nil {
		t.Errorf("%s(\"%s\") produced: %s\n", encodingType.String(), subject, err)
	}

	return actual
}
