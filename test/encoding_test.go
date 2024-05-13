package test

import (
	encoding "help_me_unpack/src"
	"testing"
)

func TestEncoding(t *testing.T) {
	for _, tc := range encodingTestCases {
		t.Run(tc.description, func(t *testing.T) {
			var actual string
			switch tc.encodingType {
			case ENCODE:
				actual = encoding.Encode(tc.subject)
			case DECODE:
				actual = encoding.Decode(tc.subject)
			}
			if tc.expected != actual {
				t.Errorf("Encode(%q) = %#v, want: %#v", tc.subject, actual, tc.expected)
			}
		})
	}
}
