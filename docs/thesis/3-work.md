# Work

## Setting up and securing the test environment

Genetic programming means one program is executing arbitrary computer programs in a computer in hope to run across "the one" before time runs. It can be considered amongst those arbitrarily (or carefully random) structured computer code, some of the candidates could potentially harm the operating system or hardware in execution. In practice, Genetic Programming known as a popular choise for generating computer malwares (viruses etc.). In scope of this research, environment is the computer of user; and avoiding harm the computer of user is one of the critical requirements of this project for making it come alive.

## `tde/testing` an interface designed for TDD users

Aim of the tool is reducing the work of developers and save labor for writing better tests rather allocating it to implement the actual code.

In order to reach tool's aim; the `tde/testing` package is designed. Source code is shared in popular public platforms such as golang.com and github.com which Go developers are familiar to download and install packages from.

Package provides functions and types necessary to setup test scenarios that will be used as TDE's testing functions. As reducing the work; `tde/testing` is also perform mostly used TDD unit test features, can be integrated into CI/CD systems. Thus, developers can comply with TDD and TDE at once with just using `tde/testing` package to write their unit tests.

### Writing a test from user's perspective

A test file can be placed in the user's package's directory and should be names as `xyz_tde.go`, xyz can be changed to whatever developer wants. First lines should contain build directives to exlude this file to be included in compiled binary and only counted when `tde` compiles the package for running tests.

```go
//go:build tde
// +build tde

package word_reverse

import (
	"tde/pkg/tde"
)

func TDE_WordReverse(e *tde.E) {
	testParameters := map[string]string{
		"Hello world":         "dlrow olleH",
		"dlrow olleH":         "Hello world",
		"The quick brown fox": "xof nworb kciuq ehT",
	}

	s := St{}
	for input, want := range testParameters {
		output := s.WordReverse(input)
		e.AssertEqual(output, want)
	}
}
```

An example TDE test file; note that package name is same with other files in same directory, tde package has to be imported, test functions declared with `TDE_` prefix in their name and their input parameter is only the `*tde.E` type.

Types and functions in the `tde/testing` are designed to make it familiar and intuitive to developers, minimizing needs for long duration learning. A for loop ranges over a collection of input and output parameters and checks if the expected result is what the function returned for given input parameters.

That interface doesn't expose details like which candidate of `s.WordReverse` is currently being testing. This is because of the rule of abstraction, a tool should present controls only the user needs. Simplicity is prioritized for reducing the cost of learning and making `tde/testing` intuitive for long-term users of `go/testing`.

In defiance of standard testing package; pair of assert functions are provided to developers built-in. Those functions returns a boolean to user when the tests are used for unit testing purposes. But when invoked by `tde` runner program; they calculate the distance between the desired and returned values for int, string and float types. The distance contains more information than simple false/true value and much more helpful when two candidates are compared to see which one is closer to what the developer wants.

### Comparing test functions of TDD and GP

Fitness functions "tests" each candidate and gives a fitness value represents how much a candidate comply with targeted solution, so the result can be any value in a range usually in [0, 1].

Test functions of TDD are used to test the only candidate function usually developed by a human and the result of testing process is either PASS or FAIL. Opposed to fitness functions, test functions return "poorly continous" results.

As the usage of TDD doesn't require knowing how close a candidate to perfect solution but only if the candidate passes all tests or not, we can't force user to write tests with more detail; cause such verbosity (meaning added lines to test code) can decrease the effectiveness of TDD on being the first source of knowledge what behaviours the function requires to have in order to comply with business requirements of software project it belongs to. TDD tests should stay "clean" and immediately "hint" the new developers to key points of the tested function, as the each members of a software development team in a company are mixed, not all of (current and future) team members are guaranteed to have equal and high experience in software development field.

### Limitations

Since panics are used to distinguish bad programs (or anomalies, as in the context of evolution), there should not be any expectation to produce a solution that uses panic/recover features for error handling (as some programmers may). There is no limitation for producing solutions that uses error-wrapping method. This should not be a problem since error-wrapping is more frequently suggested over panic/recover feature for most usual cases by programmers and [https://go.dev/blog/error-handling-and-go](official documentation of Go programming language).

The number of generations before the serial selections discards a candidate's breed from the population and blocks it from further development is limited with the case that number of candidates with at least better fitnesses are covering the contingent

Thus, local minumums may be expected to become pit-falls, impossible to get out from. One way to deal with that problem; is switching to exploratory behaviour rather than staying exploatitave. That means increasing the rate of crossovers by decreasing mutations or increasing genetic operations all together. Although, that appoach naively expects that selections won't discard all members of the breed (starts from the member, ends on the grandparent that has the progress in its genotype but not exposed in the fenotype because of a syntax-error raised from the statement/expression is not finizalized [or supported by prior declarations/assignments]) from the population before it gets chance to find supplementary modification.
