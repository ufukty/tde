# Prior Art

## GP Basics

**1992 Koza: Automatically Defining Functions**

-   Tree-based PG

**2001 Ferreira: Gene Expression Programming**

-   Operations on linear encoding, instead tree.
-   Multi-genes per candidate.
-   Rich set of genetic operations.
-   GeneXproTools based on this. [Which is used by many like GSK, GECC, University of Wales for symbolic regression, classification.](https://www.gepsoft.com/clients.htm)

## Using GP to generate test (opposite way)

Work that utilize GP to produce test-cases or test-data or test-data. Opposite of using GP to produce an implementation that passes a unit-test.

-   To produce **test cases**:

    -   2009 Fraser, Arcuri: Evolving Software Test Suites with Genetic Programming
    -   2008 Cuesta ea: Automatic Generation of Unit Tests for Java Classes
    -   2011 Fraser, Arcuri ea: EvoSuite: Automatic Test Suite Generation for Object-Oriented Software
    -   2012 Fraser, McMinn: Search-Based Unit Test Generation: Foundations, Trends, and Future Directions

-   To produce **test data**: The fitness function is based on the proportion of branches or paths that are covered by the evolved test cases (test coverage).

    -   2001 Harman, McMinn, Hierons: Automatic Test Data Generation Using Genetic Programming
    -   2004 Adorno ea: Test Data Generation Using Genetic Programming

-   Other

    -   2009 Weimer ea: Genetic Improvement of Software
    -   2009 Fraser ea: Evolutionary Testing of Classes
    -   2008 Harman ea: Co-evolving Fitness Functions for Regression Test Case Selection
    -   2004 Xie ea: Automated Test Data Generation in Evolutionary Testing

## Generalizing examples into programs (almost in triangle)

**2014 Perelman, Gulwani: Test-Driven Synthesis**

-   Influenced Excel Flash-fill (String completion)
-   Grammatical genetic programming through DSL
-   TDD.

## Triangle of GP-TDD-OOP

**2021 Illanes: Generating Object-Oriented Source Code Using Genetic Programming**

**2023 Fernandes: HOTGP Higher-Order Typed Genetic Programming**

**2023 Helmuth: Human-Driven Program Synthesis**

-   Using unit-tests instead fitness functions by enriching them artificially.

## Reviews

-   2008 Poli: A Field Guide to Genetic Programming
-   2021 Arı: A Review of Genetic Programming

## Benchmarks

-   2021 Helmuth: PSB2: The Second Program Synthesis Benchmark Suite
