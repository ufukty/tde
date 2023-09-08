package archive

import "testing"

func Test_RegexAgainstZipSlip(t *testing.T) {
	testCasesContaining := []string{
		`/`,
		`.`,
		`./`,
		`..`,
		`../`,
		`../../`,
		`./foo`,
		`l//`,
		`../foo/bar/zoo`,
		`foo/bar/../zoo`,
		`/foo/../../bar`,
		`foo/bar/..`,
		`/foo/bar/..`,
		`.\foo`,
		`..\foo\bar\zoo`,
		`foo\bar\.\zoo`,
		`foo\bar\..\zoo`,
		`\foo\..\..\bar`,
		`foo/bar/.`,
		`foo\bar\..`,
		`\foo\bar`,
		`C:\foo\bar`,
	}
	testCasesSafe := []string{
		`foo`,
		`foo/bar`,
		`foo/bar/baz`,
		`lorem/ipsum/d/o/l/o/r`,
		`foo/`,
		`foo/bar/baz.zip`,
		`foo/bar/.baz.zip`,
		`foobar`,
	}

	for i, testCase := range testCasesContaining {
		if !unsafePathFragmentMatcher.Match([]byte(testCase)) {
			t.Errorf("False negative. i=%d input=%s", i, testCase)
		}
	}

	for i, testCase := range testCasesSafe {
		if unsafePathFragmentMatcher.Match([]byte(testCase)) {
			t.Errorf("False positive. i=%d input=%s", i, testCase)
		}
	}

}
