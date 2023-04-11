### Improvement Opportunites

#### Speed of execution:

Using `ast` and `print` package, to create always-syntactically correct Go programs to spend less time with eliminating invalid-syntax candidates. - Instead fresh-compilation at each generation, only compile the changed file and link it to previous ones.

#### Evolution speed

NEAT-like mechanisms such as tagged traits, breeds... Propagating the genes creates improvement to rest that has same genetic history.

#### Further reducing errors caused by mistakes on code fragment generation

todo

#### Improvement on applicability

Compilation of a computer program takes some time. To save time on compilation all candidates of one generation are embedded into file at once and tested out at only one compilation. When compared to per-candidate compilation, current version of per-generation compilation lacks the ability to use all tests in module/package for testing candidates through target function's evolution.

####

What programmers do when they assigned with a function's implementation is running through various design patterns and previously seen solutions to similar problems. They start from popular and simple ones and keep thinking until most complicated, hardest to implement, least effective patterns and solutions. They bring ideas first, and continue by eliminating. Just like evolutionary process.

Elimination reasons are many. In syntax level: type mismatch, wrong scoping, namespace collision etc. In higher levels: readability, maintainability, scalability (resource utilization should not grow faster than company income), stability, security (access control, encryption), feasibility (labor, budget) are some concerns.
