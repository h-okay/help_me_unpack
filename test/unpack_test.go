package test

import (
	unpack "help_me_unpack/src"
	"testing"
)

func TestUnpack(t *testing.T) {
	for _, tc := range unpackTestCases {
		t.Run(tc.description, func(t *testing.T) {
			actual, err := unpack.Unpack(tc.subject)
			if err != nil {
				t.Errorf("Unpack(%q) produced: %s\n", tc.subject, err)
			}
			if tc.expected != actual {
				t.Errorf("Unpack(%q) = %#v, want: %#v\n", tc.subject, actual, tc.expected)
			}
		})
	}
}
