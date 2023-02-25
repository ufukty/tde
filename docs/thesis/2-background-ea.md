# Background - Concepts and terms (EA-related topics)

## Evolutionary Algorithms and Genetic Programming

Developer doesn't need to know form of solution prior to development, being able to define the problem and test cases the solution should pass and fail is enough to start evolution.

## Koza's work on generating human competitive results with GP in the area of electric circuitry and symbolic regression

Used in symbolic regression like

## Program synyhesis with user-provided TDD

Test-Driven Synthesis (2014, Parelman)

### Benchmarks

-   String transformations
-   Table transformations
-   XML transformations
-   Pex4Fun

## Neuroevolution of Augmented Topologies (NEAT)

-   2002, Stanley

-   Main idea is evolving ANNs along with their network topologies instead of only the weights.

-   Primary sub-problems improved

    -   Cross-overs on completely different structured candidates, size and shape
    -   Premature elimination: speciation
    -   Repressing bloating tendency of evolution

-   Secondary sub-problems improved

    -   Competing conventions

-   Proved the advantage of main idea on pole balancing problem by outperforming predecessors

-   Influenced by:

    -   Fixed topology neuroevolution

        -   sGP (Dasgupta 1992)

            -   Binary encoding: bits represent connections
            -   Scalability issues

    -   Topology evolving:
        -   Chen, 1993 (supervised training)
        -   GNARL (Angeline, 1993)
        -   Branke, 1995
        -   Gruau, 1996
            -   "also evolving structure saves time otherwise spent by humans deciding parameters at start"
        -   Yao, 1999

-   Predecessor of HyperNEAT, same author

Regardless the product type is being different; NEAT and tree-based genetic programming have some parts in common.

NEAT has introduced some solutions to couple of the problems faced in implementation of TDE. Such as premature elimination found as one of the 3 important issues addressed in paper [Stanley, 2002, page 100] along with designing a cross-over operation that works on completely different topologies and repressing tendency of GP on increasing structural complexity at each generation. That one is called as "bloating" in GP context.

Although the technique can not be direclty entitled as Genetic Programming; the difference between ANN's and AST's are not breaking opportunity to apply same principles on GP too.
