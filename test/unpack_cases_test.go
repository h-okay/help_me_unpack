package test

import (
	unpack "help_me_unpack/src"
)

type unpackTest struct {
	description string
	subject     string
	expected    unpack.UnpackedBytes
}

var unpackTestCases = []unpackTest{
	{
		description: "one",
		subject:     "ln8YhXNgYcqVyAAAcmoHQx7oia9MF1VAQFUXTK+J6B4=",
		expected: unpack.UnpackedBytes{
			Integer:         -2061992042,
			UnsignedInteger: 0xca616073,
			Short:           -14187,
			Float:           135.41580200195312,
			Double:          84.36405552356197,
			BigEndianDouble: 84.36405552356197,
		},
	},
	{
		description: "two",
		subject:     "oItAgvpVGcNLZQAAdP3RQrXSojSEwH1AQH3AhDSi0rU=",
		expected: unpack.UnpackedBytes{
			Integer:         -2109699168,
			UnsignedInteger: 0xc31955fa,
			Short:           25931,
			Float:           104.99502563476562,
			Double:          476.0322767601277,
			BigEndianDouble: 476.0322767601277,
		},
	},
	{
		description: "three",
		subject:     "61YAgbqP9vC5dQAAbqJpQyZbEfpHGntAQHsaR/oRWyY=",
		expected: unpack.UnpackedBytes{
			Integer:         -2130684181,
			UnsignedInteger: 0xf0f68fba,
			Short:           30137,
			Float:           233.63449096679688,
			Double:          433.6425724676104,
			BigEndianDouble: 433.6425724676104,
		},
	},
}
