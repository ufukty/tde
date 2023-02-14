# **Test Driven Evolution**: Combining Genetic Programming with Test Driven Development to put evolution to average programmer's toolboxg

## Background

## Suggestion

## Random code-fragment generation

### Alternative methods

Generation of Go codes that have valid syntaxes and compatible with compiler can be challenging. Three methods considered to be used in the research: Generating sequence of random characters, tokens and AST nodes.

Tokens consist by keywords and symbols (operators) that are allowed to be placed in a valid Go program by specification.

AST (Abstract-Syntax-Tree) is used to represent computer programs in a tree of nodes. Each node represents a section of the code and any node can consists by one or more nodes that each represents smaller portion of code that the parent node represents.

Ordering 3 methods from left to right, it can be seen that while the complexity of program that creates valid codes is increasing, rate of codes with invalid syntax created by that program decreases. With that; since the both token and AST node generation methods need details about the structure (both lexical tokens and syntax) of programming language; it can be forseeable that keeping search-space as all valid Go programs are reachable with initialization process at start and after performing genetic operations is tied with the success on implementing the token/node generating program in a way that is has no missing information on its due.

An example to the complexity of generating AST-nodes could be the amount of node types that has to be considered. In the time of that counting; (Version of Go compiler and libraries used is 1.19.3); there are more than 50 node types for representing some-nested structures that varies in the size and functionality of the area they represents. Examples can be given as `Package` and `File` for relatively bigger structures; and `Field`, `IfStmt` (If Statement), `SliceExpr` (Slice Expression) and `MapType` for smaller ones. Most of node types consists by more than one type and number subnodes. Thus, implementing an AST node based code generator is challanging and detailed task. More importantly, that kind of code generator is language-specific. That means a code generator for Go won't be able to used for generating codes in different languages.

### Reducing errors on genetic operations and streamlining the evolution process (search-space traversal)

Research has aimed to reduce mistakes on code fragment generation by making the process aware of context and types. Since any mistake leads to many of undetected lexical and syntax errors in code file creation and program generation (compile) phases, this reduction enabled evolution to waste less time trying to reach neighbour of a compilable but lacking candidate, combined with cruelness of selection algorithms against candidates who needs `n` generations before completing and surfacing underlying improvements in the phenotype (test performance) and fails on those phases throughout all of those `1st, 2nd, 3rd, ..., (n - 1)th` generations.

Reduction on failures made possible with making CFG **context and type aware**. That is restricting the CFG with only the symbols accessible from the insertion point. Meaning of symbols: variables, function and type declarations of both current package and imported packages. Then, CFG creates expressions which its resolving type comply with the insertion point. So, the resolving type of the expression given to `ReturnStmt` can only be the function's output parameter's type that is stated in by user:

```go
func FuncName() int {
    return ... // an expression that resolves to `int`
}
```

Any node placed before ReturnStmt will have a guide for

## Using dataset of web-scrapped code as initialization

Another method could be using set of example codes that are written for different problems, collected from web and presents each type of code block and design patterns for initial population. If the used genetic operations (when applied over and over) are capable to reach every point of search-space even in unlimited time (like eliminating blind spots theorically); it might mean that the program is able to turn any form of solution to any other; put it differently, is able to navigate from any point in search-space to another. Thus; it can be predicable that creating initial population with such dataset would not hurt the ability of program to create a solution to user-provided tests.

One issue is assuring the ability of program to `reduce` the complexity of codes carried by candidates in any given generation. This is challanging, especially because of the progress on getting success is usually made on the offspring that inherits parents' bigger parts (while the other offspring gets the smaller parts from each); and this is known to lead a phenomena called "bloating". One attempt to solve bloating problem is impacting the size of candidate into fitness evalution. Such change leads one candidate to get better score then another candidate that has same number of performance on tests but have bigger file size. FIXME: Another attempt to solve this problem is tracking genes and traits; so in a cut point decision phase of cross-over operation, one offspring getting will be avoided. That method was tried on NEAT (Neuro-Evolution of Augmented Topologies) and it seems successfully solves the problem.

## AST-node based code fragment generator

There are two approaches can be consiedered for implementing AST-node based code fragment generation (CAAST-RCFG), I seperate them on whether or not being context-aware. CAAST-RCFG tracks down each block statements from up in the package files down to the smallest block that wraps around the next line will be created on. Tracked information are declarated (and values assigned) variables, constants, previously defined functions in the file, accessible functions that defined in package and imports. Of course this method requires another functionality that, when new variable, constant, function is needed to be created, it handled gracefully, accordingly to the syntax of language. It might be align with the common sense that CAAST-RCFG accelarate evolution process since it avoids invalid syntax candidate code to reach to fitness evaluation phase and

One issue with context-aware (that is creating new lines that can only access previously declarated variables, constants and imported libraries) node generation is that the amount of computation needed to resolve

## Validity of the solution

### Turing completeness

A computational environment is Turing-complete only if it can reproduce itself. One thing to prove that the TDE can find solution when it is theorically possible can be trying to produce TDE with TDE.

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
func addition(a, b int) {
    return a // fault
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
```

The `Expr` that is given to `ReturnStmt` needs a rework. Existing `Ident{a}` should be wrapped with a `BinaryExpr{+}` along newly created `Ident(b)`.

```yaml
# Fixed version
*ast.ReturnStmt {
    Results: []ast.Expr{
        0: *ast.BinaryExpr {
            X: *ast.Ident {
                Name: "a"
            }
            Op: +
            Y: *ast.Ident {
                Name: "b"
            }
        }
    }
}
```

Important to note: The variable name that is choosen for `Ident(b)` is not arbitrary. In fact, it is one of the input parameters of the user provided `FuncDecl` along with the `Ident(a)`. If the "Code Fragment Generator" would've generated random variable names at each node creation, coming across to correct pair of variables withing correct expression would take much more try-error, and would increase its chance to get eliminated by selection algorithm.

## Dealing with inconsistencies: Evolving homologous structures at once

```go
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

## Mutations that work

Cross-over is there to move around in search space with bigger steps. They are good for quickly reaching to far regions of search space, and revealing what sits in there. Moving from the location of initial solution to a compilable candidate's, it takes a series of bigger and smaller steps in search space. Because it should not be expected that the first candidate solution in the landing point of one cross-over will just compile without errors. Put semantically incorrect candidates aside; finding a candidate that at least compiles without a "lexical error" (that is an error can be occured in tokenization step) or "syntax error" (belongs to compilation step) needs a combination of bigger and smaller steps performed in specific order.

As seen in literature, "explorative" and "exploitative" terms are used to describe (respectively) quickly scanning bigger areas and focusing to smaller area.

### Mutation's importance

In this work; many concerns involved in the design of mutation operations. Need is raised by that any haphazard edit on a compilable code's abstract syntax tree can easily cause syntax errors; which will make the candidate fall behind its anchestry in fitness order, which will also increase its chance to get discrarded by roulette in the next selection phase. Long before complimentary mutations get a change to take place in next generations and introduce an improvement in working code.

It should be expected from the final set of mutation operations to be able to iron out syntax and lexical errors of a candidate that is semantically close to transformed into the accepted solution. An example in context of string reverse problem, a candidate that creates the expected loop with expected statements inside it, but defines the temporary storage array's type as `[]int` instead of `string` or `[]byte`. Of course a solution that utilize a temporary array for such problem might not be the best alternative. Since just swapping elements from start and end until reach the center is better method. But, in the scope of this work, any solution that passes user-provided tests is perfect. Although:

```go
func StringReverse(in []rune) []rune {
    var1 := []int{} // wrong type
    for i := len(in) -1 ; i >= 0; i-- {
        var1 = append(var1, in[i])
    }
    return var1
}
```

### More examples on need for small-fixes

Many situations can be thought where inconsistencies are type-related and lead to syntax-errors on compilation.

### Variable name replacement after Cross-Over

> dsdsdsdsd

### Looping with condition

```go
for condition {
    list of statements
}
```

### Looping over a collection of items: map, slice.

```go
for i, v := range collection {
    list of statements
}
```

### Conditionalizing a list of statements with an expression

```go
if condition {
    list of statements
}
```

### Data structure transform

```go
keys := []string{...}
values := []int{...}
```

Above to below:

```go
type Item struct {
    key string
    value int
}
items := []Item{...}
```

Prerequisites of improving-mutation are determining of correct

-   set of statements to conditionalize
-   expression to use as condition (which resolves to a boolean)

### Performing trials of different mutations on same candidate simultaneously, for many candidates in a generation

For a set of _n_ candidates of the generation _i_, after selection with survival ratio of _s_ starts, the population size becomes `n * s` For the average number of candidates _c_, there should _m_ different mutation be carried over to next evalution & selection and the population size remain constant.

$$ n = n _ s _ c _ m$$
$$ 1 = s _ c \* m $$

## Attributes of TDE's search space

What does it look like? How many dimensions? What makes 2 candidates neighbours? Is it continuous or discrete? Why it is helpful to imagine candidates in a organized search space?

### Sorting all Go codes possible

Attributes that any Go code have:

-   Not helpful:
    -   Length
    -   Number of code blocks in some kind e.g. `GenDecl`, `FuncDecl`
-   Helpful:
    -   Least number of mutations required to transform candidate A into candidate B.

```go
f := f()
```
