# Quick Start

[TOC]

## Introduction

**Credits**

Creating, testing and selecting of randomly generated or modified variations of code requires significant amount of computational resources to be used. 

To be able to use DeepThinker, you need to purchase credits. Please read [Purchasing credits](#) 

**License**

License for DeepThinker Community Beta only allows you to use it on open sourced code. Criteria for that is to publishing exact module in a publicly accessible internet platform under one of the Apache, BSD, MIT licenses or their variations. Please read [Community Beta License](#) page for details. 

## Download tool(s)

Download the CLI tool from below.

- [Linux (ARM)](#)
- [Linux (AMD64)](#)
- [Mac (Intel)](#)
- [Mac (M)](#)

Move the binary into a  `$PATH` folder.

Additionally you can install DeepThinker for VisualStudio Code. That will allow user to start an evolution session by clicking the start button which is placed in gutter, lined with each test function. Also, results will be shown in editor as they arrive.

## Login

Run login command to enter username and password.

```sh
deepthinker login
```

## Upload your module

Both producing variations and testing them will happen in deepthinker.app servers. So, you need to upload your module to servers. Run below command from a directory in your go module. 

```sh
deepthinker upload
```

Note: By default `.git`, `build`, `docs`, `.vscode` directories are excluded and only files with  `go`, `mod`, `sum` extensions are included. You can exclude more directories and include more extensions. [Upload filters](#)

## Start an evolution session

You'll use the `evo` command to start evolution sessions. There are many optional flags you can use to optimize between cost and performance.

```sh
deepthinker evo start \
	-credits 10 \
  -population 1000 \
	-generations 10 \
	-archive 00000000-0000-0000-0000-000000000000 \
	-test examples/word-reverse/word-reverse_test.go:TDE_WordReverse \
	-target examples/word-reverse/word-reverse.go:WordReverse
```

This command will overwrite your target file's target function part as evolution cycles completes.
