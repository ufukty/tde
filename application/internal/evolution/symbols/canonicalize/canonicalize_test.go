package canonicalize

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"tde/internal/evolution/evaluation/discovery"
	"tde/internal/evolution/evaluation/list"
	"tde/internal/utilities/osw"
	"testing"
)

type testcase struct {
	input, want []string
}

func prepareTestCases() ([]testcase, error) {
	testcases := []testcase{
		{
			input: []string{
				"fmt",
				"math",
				"tde/internal/evolution/evaluation/discovery",
				"tde/internal/evolution/evaluation/list",
				"tde/internal/utilities/osw",
			},
			want: []string{
				"{{goroot}}/src/fmt",
				"{{goroot}}/src/math",
				"{{cwd}}/../../../../../tde/internal/evolution/evaluation/discovery",
				"{{cwd}}/../../../../../tde/internal/evolution/evaluation/list",
				"{{cwd}}/../../../../../tde/internal/utilities/osw",
			},
		},
		{
			input: []string{
				"tde/internal/evolution/evaluation/discovery",
				"tde/internal/evolution/evaluation/list",
				"fmt",
				"math",
				"tde/internal/utilities/osw",
			},
			want: []string{
				"{{cwd}}/../../../../../tde/internal/evolution/evaluation/discovery",
				"{{cwd}}/../../../../../tde/internal/evolution/evaluation/list",
				"{{goroot}}/src/fmt",
				"{{goroot}}/src/math",
				"{{cwd}}/../../../../../tde/internal/utilities/osw",
			},
		},
	}

	goroot, err := GoRoot()
	if err != nil {
		return nil, fmt.Errorf("GoRoot: %w", err)
	}
	cwd, err := osw.WorkingDir()
	if err != nil {
		return nil, fmt.Errorf("osw.WorkingDir: %w", err)
	}

	for _, tc := range testcases {
		for i := range tc.want {
			tc.want[i] = strings.ReplaceAll(tc.want[i], "{{cwd}}", cwd)
			tc.want[i] = strings.ReplaceAll(tc.want[i], "{{goroot}}", goroot)
		}
	}

	return testcases, nil
}

func TestCanonicalize(t *testing.T) {

	tcs, err := prepareTestCases()
	if err != nil {
		t.Fatal(fmt.Errorf("preparing testcases: %w", err))
	}
	for i, tc := range tcs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {

			mod, err := discovery.ModuleRoot()
			if err != nil {
				t.Fatal(fmt.Errorf("prep, whereAmI: %w", err))
			}

			modpkgs, err := list.ListAllPackages(mod)
			if err != nil {
				t.Fatal(fmt.Errorf("prep, listing all module packages: %w", err))
			}

			goroot, err := GoRoot()
			if err != nil {
				t.Fatal(fmt.Errorf("prep, getting go root: %w", err))
			}

			got := CanonicalizePaths(goroot, mod, modpkgs, tc.input)

			if len(got) != len(tc.input) {
				t.Fatalf("assert 1: lengths mitmatch: got (%d) != tc.input (%d)\n", len(got), len(tc.input))
			}

			for _, path := range got {
				if !filepath.IsAbs(path) {
					t.Fatal(fmt.Errorf("assert 2: %q is not canonicalized", path))
				}
			}
		})
	}
}
