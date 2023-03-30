# 4. Discussion

## Improvement oppurtunities

### Reducing time spent on printing/compiling/parsing

Compilation has the most impact on resonse time amongst all steps, it is the current bottleneck. Measure of impact is avg. compilation time per package holding the candidate function times number of candidates in each generation (4000 used) times number of generations (10 used).

In a laptop laptop which has i9 processor, 16gb memory, SSD and subpar cooling system, compilation of 100 candidate is completed within 9 seconds. That means without rest of the calculations, compilation alone 10 generations and 4000 candidates per generation can make waiting time of user:

$$
9 \frac{\text{seconds}}{\text{candidate}} \times (4000/100) \frac{\text{candidate}}{\text{generation}} \times 10 \frac{\text{generation}}{\text{run}} \approx 1 \frac{\text{hours} }{\text{run}}
$$

Average wait time of 1 hour is unacceptable for average developer to integrate this tool in their development workflow; making this tool experimental for day. With that being said, there are ways to reduce waiting time.

**Making the whole process in-memory**

In current proof-of-concept, there are many steps that requires us to use various file operations. Lack of ability to compile AST-represented Go code -directly- to an executable in memory, creates further slow dependencies like:

-   Code printing from AST,
-   Managing folders (duplicates of original module to test candidates),
-   Creating a shell to run executable

Built-in Go compiler (and further a Go interpreter) in evolver program can remove all chores above.

**Re-compile only changed parts**

It is waste of time to compile a package in each generation from scratch when most of the time the only changed information is a single statement, expression, declaration or a specification in the implemented function.

Compilers divides compilation process into many steps involving parsing, type checking, IR construction, generating machine code. Not all artificats from previous compilation needs to redone, many unchanged can be reused to reduce time spent on compilation.

**Distributed compilation and testing**

$$
24 {\$ \over \text{pod} \times \text{month}} \times 10 \text{pod} \times 1\text{hour} = 0,13 \$
$$

Tasks of printing, compiling, testing all candidates can be distributed over a cluster of application servers that each is called runner. That way, the waiting time will be roughly divided by number of instances in the cluster. Using 10 pods (as the term of Kubernetes)

## Links

1. [Appendix: Performance comparison of compilation options](appendix/performance%20comparison%20of%20compilation%20options.md)
1. Go tools create a single object file in `pkg/<architecture>` per package https://groups.google.com/g/golang-nuts/c/RSd3B5_rIFE
