package test

import (
	unpack "help_me_unpack/src"
	"testing"
)

func TestUnpack(t *testing.T) {
	for _, tc := range unpackTestCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := unpack.Unpack(tc.subject)
			if tc.expected != actual {
				t.Errorf("Unpack(%q) = %#v, want: %#v", tc.subject, actual, tc.expected)
			}
		})
	}
}
