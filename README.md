# Lox Language Interpreter

This repository contains an implementation of the Lox programming language interpreter as part of the "Build Your Own Interpreter" challenge from codecrafters.
This Challenge follows the book - [Crafting Interpreter by Robert Nystrom](https://craftinginterpreters.com)

## Project Description

The Lox interpreter is built using Go and follows the principles of Test-Driven Development (TDD). It uses a recursive descent parser to convert a stream of lexical tokens into an Abstract Syntax Tree (AST).

## Getting Started

### Prerequisites

Please make sure you have Go 1.22 installed on your machine.

### Installation

Clone the repository:
```sh
git clone https://github.com/harish876/lox-lang.git
cd lox-lang

```sh
git add.
git commit -m "pass 1st stage" # any msg
git push origin master
```

It's time to move on to the next stage!

# Usage
1. Add additional grammar rules and code to parser/parser.go
2. Add test cases to parser/parser_test.go
3. cd cmd/myinterpreter/parser and run ```go test``` to run all test cases or ```go test -run [test_name]``` to run a particular test.


# Todo
1. Adding functions, statements, variable binding, classes ( Completing AST - full language feature set ) 
2. Converting to Bytecode and running on a VM (future goal) 


