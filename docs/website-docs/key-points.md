# Short examples

## Comparing TDE and DeepThinker

**TDE** is for Test Driven Evolution. Meaning that using unit tests to produce computer code with **little-to-none modification on them**. It is important because **in past**, many people found that unit tests was non-functional for genetic programming. 

**DeepThinker** is a coding tool which produces Go code from unit tests. DeepThinker utilizes many technique as well as TDE. 

- Tooling is designed to **abstract away** all the implementation details of genetic programming, enabling **coders with any experience** to use without need to understand those techniques works. 
- What the coder has to do is writing an Arrange-Act-Assert style **unit tests** (which is always good practice) and press the **green triangle** on the editor to start a session.
- DeepThinker has its own cluster of runners to let you produce code as quick as possible. It is enough to testify 1000s of candidates **within seconds** and return only the best performing one.

## Secondary advantages of adopting DeepThinker

- Greater test coverage within same labor cost
- 

## When to use DeepThinker instead of design or asking to AI?

DeepThinker is good for novel coding problems:

- If you can define the problem and know the answer: design. 
- If you can define the problem but don't know the answer: ask AI. 
- If you **can't define the problem** but you are able to classify a suggestion good or bad fit: use DeepThinker.

So it is like playing **hot & cold** for coding.

Consider that with slight turnover, rewrites and nightly patches; each repository moves away from common-practices and data structures, unwillingly becomes unique. Any standard problem's implementation can be classified as novel.

Also:

- DeepThinker only returns a code piece as the **solution** if it **passes all of the test cases** you provide. So it is not like checking if language model's code works after each.

- DeepThinker can find an implementation that is in **harmony** with rest of your codebase. DeepThinker have the ability to utilize any symbol you previously defined in the module. Including **constants, variables, your data structures, functions** and more.

## Purchasing Credits

