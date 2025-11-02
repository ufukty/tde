# TDE

![Logotype](assets/github-social-preview.png)

This is the implementation of my suggestions on solving the problem of applying Genetic Programming to produce test-passing Go functions that I've presented in [my Master's thesis](https://tez.yok.gov.tr/UlusalTezMerkezi/TezGoster?key=weFMBHaUra8rsS5wi2bmHDKlIvi-IwlFkdPWTMwNi0k9Pt1C4PzNAFzxcjzHPgAW) in 2024. For English readers an [LLM-provided pseudo-review](assets/llm-pseudo-review.pdf) might be helpful on understanding what this work currently is and is not. Note that, it is not shared to replace peer review.

The codebase provides many packages than can be helpful for future researchers to implement necessary tooling on performing genetic operations on Go AST. It also contains the partial implementations for the **Layered Search** and **Context Aware Subtree Generation** (CA-STG, the newest component of the Strongly-Typed ASTGP) algorithms that each are the main suggestions of my thesis. They still have time for completion.

If you are working on similar method or for goal, don't hesitate [contacting me](https://ufukty.com?utm_source=tde).

## Technic

Against the search-space explosion of applying GP I suggest two approaches: guiding genetic operations in favor of producing compiler compliant ASTs and spinning off sub-evolutions with isolated pools. Former gets beyond STGP in implementation because of compiler involvement and the care required by the disconnection between type checks performed by the AST printer and code compiler. Latter is expected to control search-space expansion and provide survival chance to those candidates require multiple generations before introducing progress to get rid of structural problems before getting compared against their mediocre yet perfectly running distant neighbors.

## License

-   All rights reserved (c) Ufuktan Yıldırım. Dependencies may belong to others. Open source software used under various licenses including MIT, BSD etc.
