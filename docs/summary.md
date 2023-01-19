# **Test Driven Evolution**: Combining Genetic Programming with Test Driven Development to put evolution to average programmer's toolboxg

## Liability and responsibility in Software Development

Software is frequently developed as a commercial product that solves a social (or derived) problem and creates responsibility and liability on the seller company. Responsibility comes from the need to maintain company reputation in target consumer demographic and the liability comes from promises that company explicitly gave to customers as compliences to certain standards or the requirement of complying with the law.

## Meaning of production-ready

Just like workers of companies' other product departments, software developers should be able to inspect the prior work as the team makes progress in development to spot unfitting progress to everchanging requirements.

The aim is to present the Genetic Programming within a tool to software developers for helping them to build their commercial purposed software. The tool eventually should have features like familiarity, easy-to-learn

Most of the companies develop CRUD programs. That is a software runs on a computer accessible to clients. Exposes an stable interface consists by serial of endpoints each creates/reads/updates/deletes a resource. Resources are abstract objects, meaningful for the domain: like users, posts, photos for a social media platform; or employees, customers, e-mails, calls for a customer relation management software. Developer teams from variety of companies operates on different sector, or industries; it is safe to say that most of the teams specialized on developing CRUD programs and their skill-sets are shaped around the requirements of developing CRUD programs.

CI/CD is trending topic for nearly 20 years. That is automating the cycle of integration and deployment which is expected to be a loop because of everchanging set of requirements and priorities that are faced by organizations produce software, in addition to a work force that consists by employees don't stay forever on same company and have different degrees of experiences on the field. Steps of CI/CD are mainly considered as code, build, test, release, deploy, operate, monitor and plan (for next round).

Familiarity: Most of the devs are experienced in building CRUD applications and their skill set don't cover the Genetic Programming or another EA variation. In order to enable those developers leveraging from fruits of GP, the method should be delivered to them in a tool that presents a

For the sake of our will to make Genetic Programming widely used method for generating computer programs it is necessary to make sure a regular dev team of a standard organization that is hired for developing CRUD programs with popular programming languages of domain (eg. Java, C#, Go) should be able to

## Changing requirements of projects

Software projects (service/product) mature with time and the increasing diffusion area on target group. Early adopters of innovations are the people who are expected to be the most educated/experienced ones amongst the whole target group. Through diffusion process, it is expected to feedbacks and requests raised by customers differantiate with the diffusion phase. Thus it is unavoidable that requirements for a particular product or a service will change with time; and software development process should

## Work force circulation

Companies can not rely their employees on staying in the company until a software project completes its lifecycle. Employees come and go, but products should keep get fixed for newly discovered requirements, and get updated for new features customers request without effecting by if employees who worked on first phases of project still in the company. Latest workforce should be able to work on the existing codebase to add new features without the fear of breaking previously working features that implemented by previous workers. Following by test driven approach, key requirements that the software expected to satisfy are added into the codebase as tests at the moment those requirements raised and their reasoning known by employees of that time. Codebase is supposed to get tested for all tests at any time a possibly breaking-change occurred, or before and after every deployment by the CI/CD pipeline.

## Test driven development

Test driven development are implemented by developers with writing so-called tests before actually implementing the program. Those tests can be divided into a couple of classes. Most popularly adopted and well known examples are unit tests and integration tests. Unit test is when a component tested by itself in isolation from the components it supposed to interact in active deployment. So, when a unit test failed, smallest component can be addressed immediately by tracking lately-turned-to-fail tests. Other class is when systems that consists by more than one components deployed into real or mock (a.k.a test, development) environment and tested for higher-level requirements than unit tests. To understand relation between test types, thinking about that system definition may help: A system consists by components and its behavior emerges by the interactions of components, also each component can be seen as sub-systems that consists by sub-components smaller than itself. Thus, unit test are for smallest components, and at each organizational-level from that to the higher, different type of test tools and techniques are considered.

## Evolutionary Algorithms and Genetic Programming

Developer doesn't need to know form of solution prior to development, being able to define the problem and test cases the solution should pass and fail is enough to start evolution.

## Improvement Opportunites

Speed of execution: - Using `ast` and `print` package, to create always-syntactically correct Go programs to spend less time with eliminating invalid-syntax candidates. - Instead fresh-compilation at each generation, only compile the changed file and link it to previous ones.

Evolution speed - NEAT-like mechanisms such as tagged traits, breeds... Propagating the genes creates improvement to rest that has same genetic history.

Improvement on applicability - Compilation of a computer program takes some time. To save time on compilation all candidates of one generation are embedded into file at once and tested out at only one compilation. When compared to per-candidate compilation, current version of per-generation compilation lacks the ability to use all tests in module/package for testing candidates through target function's evolution.

## Requirement that comes from commercialization

Genetic programming means one program is executing arbitrary computer programs in a computer in hope to run across "the one" before time runs. It can be considered amongst those arbitrarily (or carefully random) structured computer code, some of the candidates could potentially harm the operating system or hardware in execution. In practice, Genetic Programming known as a popular choise for generating computer malwares (viruses etc.). In scope of this research, environment is the computer of user; and avoiding harm the computer of user is one of the critical requirements of this project for making it come alive.

## Comparing fitness and test functions

Fitness functions "tests" each candidate and gives a fitness value represents how much a candidate comply with targeted solution, so the result can be any value in a range usually in [0, 1].

Test functions of TDD are used to test the only candidate function usually developed by a human and the result of testing process is either PASS or FAIL. Opposed to fitness functions, test functions return "poorly continous" results.

As the usage of TDD doesn't require knowing how close a candidate to perfect solution but only if the candidate passes all tests or not, we can't force user to write tests with more detail; cause such verbosity (meaning added lines to test code) can decrease the effectiveness of TDD on being the first source of knowledge what behaviours the function requires to have in order to comply with business requirements of software project it belongs to. TDD tests should stay "clean" and immediately "hint" the new developers to key points of the tested function, as the each members of a software development team in a company are mixed, not all of (current and future) team members are guaranteed to have equal and high experience in software development field.

## Limitations

Since panics are used to distinguish bad programs (or anomalies, as in the context of evolution), there should not be any expectation to produce a solution that uses panic/recover features for error handling (as some programmers may). There is no limitation for producing solutions that uses error-wrapping method. This should not be a problem since error-wrapping is more frequently suggested over panic/recover feature for most usual cases by programmers and [https://go.dev/blog/error-handling-and-go](official documentation of Go programming language).

The number of generations before the serial selections discards a candidate's breed from the population and blocks it from further development is limited with the case that number of candidates with at least better fitnesses are covering the contingent

Thus, local minumums may be expected to become pit-falls, impossible to get out from. One way to deal with that problem; is switching to exploratory behaviour rather than staying exploatitave. That means increasing the rate of crossovers by decreasing mutations or increasing genetic operations all together. Although, that appoach naively expects that selections won't discard all members of the breed (starts from the member, ends on the grandparent that has the progress in its genotype but not exposed in the fenotype because of a syntax-error raised from the statement/expression is not finizalized [or supported by prior declarations/assignments]) from the population before it gets chance to find supplementary modification.

## Random code-fragment generation

Generation of Go codes that have valid syntaxes and compatible with compiler can be challenging. Three methods used in research: Generating sequence of random characters, tokens and AST nodes.

Tokens consist by keywords and symbols (operators) that are allowed to be placed in a valid Go program by specification.

AST (Abstract-Syntax-Tree) is used to represent computer programs in a tree of nodes. Each node represents a section of the code and any node can consists by one or more nodes that each represents smaller portion of code that the parent node represents.

Ordering 3 methods from left to right, it can be seen that while the complexity of program that creates valid codes is increasing, rate of codes with invalid syntax created by that program decreases. With that; since the both token and AST node generation methods need details about the structure (both lexical tokens and syntax) of programming language; it can be forseeable that keeping search-space as all valid Go programs are reachable with initialization process at start and after performing genetic operations is tied with the success on implementing the token/node generating program in a way that is has no missing information on its due.

An example to the complexity of generating AST-nodes could be the amount of node types that has to be considered. In the time of that counting; (Version of Go compiler and libraries used is 1.19.3); there are more than 50 node types for representing some-nested structures that varies in the size and functionality of the area they represents. Examples can be given as "Package" and "File" for relatively bigger structures; and "Field", "IfStmt" ("If Statement"), "SliceExpr" ("Slice Expression") and "MapType" for smaller ones. Most of node types consists by more than one type and number subnodes. Thus, implementing an AST node based code generator is challanging and detailed task. More importantly, that kind of code generator is language-specific. That means a code generator for Go won't be able to used for generating codes in different languages.

## Using dataset of web-scrapped code as initialization

Another method could be using set of example codes that are written for different problems, collected from web and presents each type of code block and design patterns for initial population. If the used genetic operations (when applied over and over) are capable to reach every point of search-space even in unlimited time (like eliminating blind spots theorically); it might mean that the program is able to turn any form of solution to any other; put it differently, is able to navigate from any point in search-space to another. Thus; it can be predicable that creating initial population with such dataset would not hurt the ability of program to create a solution to user-provided tests.

One issue is assuring the ability of program to "reduce" the complexity of codes carried by candidates in any given generation. This is challanging, especially because of the progress on getting success is usually made on the offspring that inherits parents' bigger parts (while the other offspring gets the smaller parts from each); and this is known to lead a phenomena called "bloating". One attempt to solve bloating problem is impacting the size of candidate into fitness evalution. Such change leads one candidate to get better score then another candidate that has same number of performance on tests but have bigger file size. FIXME: Another attempt to solve this problem is tracking genes and traits; so in a cut point decision phase of cross-over operation, one offspring getting will be avoided. That method was tried on NEAT (Neuro-Evolution of Augmented Topologies) and it seems successfully solves the problem.

## AST-node based code fragment generator

There are two approaches can be consiedered for implementing AST-node based code fragment generation (CAAST-RCFG), I seperate them on whether or not being context-aware. CAAST-RCFG tracks down each block statements from up in the package files down to the smallest block that wraps around the next line will be created on. Tracked information are declarated (and values assigned) variables, constants, previously defined functions in the file, accessible functions that defined in package and imports. Of course this method requires another functionality that, when new variable, constant, function is needed to be created, it handled gracefully, accordingly to the syntax of language. It might be align with the common sense that CAAST-RCFG accelarate evolution process since it avoids invalid syntax candidate code to reach to fitness evaluation phase and

One issue with context-aware (that is creating new lines that can only access previously declarated variables, constants and imported libraries) node generation is that the amount of computation needed to resolve

## Turing completeness

A computational environment is Turing-complete only if it can reproduce itself. One thing to prove that the TDE can find solution when it is theorically possible can be trying to produce TDE with TDE.

## Comparing with alternative methods

Generating code from configuration files: Such, recently developed tools promise generating domain-specific computer programs in return of user-provided configuration files. Examples can be given as Swagger,

AI models trained on publicly accessible internet dataset: Github Copilot, ChatGPT. Those solutions work in a way that accept a text promt that explains the problem the user wants service to solve.

## Measure against the potential harm that might caused by running arbitrary code in a computer

Throughout n generation and m candidate at each; n\*m arbitrarily created computer codes are executed with full access on the computer with nothing between them with the hardware. Also filesystem and network access are also given to candidates for the scope of this research being creating production-ready code.

By its nature, the GP suits to the needs of generating harmful software, due to that, its capabilities are utilized by many on research of producing new computer viruses.

Alternative methods found for running arbitrary code without giving full access of the computer and file system to the program are sandboxing, virtualization and containerization. Popular choices are FireJail, AppArmor, SELinux for sandboxing; VirtualBox for virtualization; and Docker for containerization.

Similar use case is implemented for "Go Playground" (https://go.dev/play) which is a website that compiles user-provided Go code in a server, runs it and sends the output back to the user. Codebase shows that (https://go.googlesource.com/playground access date 26.12.2022) the method choosen for executing user-provided code is using a set of Docker containers with addition of a container isolation solution called gVisor. With that addition, communication between host device and the user code is intercepted by gVisor.

## Begging to the roluette for one more chance

The roulette wheel or many other popular selection algorithms don't count undergoing changes that have possibility to lead to a progress with addition of couple others. If the change in AST (let's say "the genotype") doesn't lead to a improvement in fitness score by the end of iteration (when the fenotype [compilable code] get evaluated and tested), has possibility to get eliminated on selection.

Evolving AST trees introduces another level of diffuculty on protecting n-step improvement against the roulette. Because, any change on a version of a code that has only semantic errors can lead to a syntax error; which has a penalty on the score (See example below).

One thing about second chances is that they are limited with the quota remains from candidates have better scores already, but not necessarily the best. It doesn’t matter if a candidate have a change in its genotype that doesn't show up as progress on the fenotype yet.When roulette starts to spin and seeks for which ones to survive, which ones to eliminate, it is all about what you have already to show; at the other hand your potential, doesn’t matter.

Dividing the fitness measurement into multiple objectives is a promising idea to reduce rate of losing progresses under surface into roulette’s careless business manner.

```go
// Faulty original
func addition(a, b int) {
    return a
}

// Fixed version
func addition(a, b int) {
    return a + b // changed line
}
```

AST representations (simplified):

```yaml
# Faulty original
*ast.ReturnStmt {
    Results: []ast.Expr {
        0: *ast.Ident {
            Name: "a"
        }
    }
}

# Fixed version
*ast.ReturnStmt {
    Results: []ast.Expr{
        0: *ast.BinaryExpr {   # 1 Newly created, choosen randomly amongst 10+ possible Expression variation
            X: *ast.Ident {    # 2 Previously existing structure; moved to inner level
                Name: "a"
            }
            Op: +              # 3 Choosen randomly amongst nearly 10 possible tokens
            Y: *ast.Ident {    # 4 Similar to 1st line
                Name: "b"      # 5 Chosen randomly amongst previously declarated variables and functions.
            }
        }
    }
}
```

One way to solve this changing the expression that is given to Return node

## Dealing with inconsistencies: Evolving homologous structures at once

```cpp
Fenotype = GenotypeToFenotypeGenerator(AST-based-data-structure)
```

One change in one property of one structure (instance of either Decl, Stmt, Spec, Expr) will reflect to every occurances of that instance in the actual code. Thus inconsistencies among occurances are elimintated by design.

As one approach; GenotypeToFenotypeGenerator function will be a kind of a permutation generator; since what all it does is generating a layout for the Go code by placing prebuilt building blocks in some order by choosing from given set. This approach satifies the above case.

## Getting the right one among the set of passing solutions: evolving from basic to complex

Producing a code piece that passes a test is easier challange than selecting the simplest one amongst set of passing solutions. Because of two reasons, the search space may contain more than one passing solution: alternative solutions and bloated solutions.

Alternative solutions can be thought as different implementations of same algorithm, or alternative algorithms to solve same problem each are valid (and more imporantly fitting) but varies between compute and memory utilization.

Bloated solutions are the ones that not only satisfy all requirements but also performs unnecessary actions for unrequested and sometimes harmful purposes. For example a simple string reversing function has no need to access a website in addition to a simple reversed for loop. An average software developer should not be expected to write tests that will exclude all bloated solutions. Deploying bloated solution into a server can be a problem; because (from cyber-security perspective) a system must only perform operations its meant to do and nothing else. Normally, when a developer fails to produce fitting solution, cyber-attackers explores those unfitting points as potential attack-surfaces and try to combine those with others to plan a cyber-attack.

It might be harder to filter those bloated solutions after they emerged rather than designing the evolution process in a way to reduce the possibility of producing them at first place.

(and possible more complicated than the ideal and has additional features that are not required or even wanted) from getting marked as approved. So, producing code with tests can only be viable if the tool presents to the coder the simplest one as solution from the set of passing solutions.
