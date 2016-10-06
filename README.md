etymu
====

[![Build Status](https://travis-ci.org/Knetic/etymu.svg?branch=master)](https://travis-ci.org/Knetic/etymu)
[![Godoc](https://godoc.org/github.com/Knetic/etymu?status.png)](https://godoc.org/github.com/Knetic/etymu)

Generates lexers from language-agnostic lexer files. All the good stuff from flex/lex, none of the bad decisions.

Why not flex/lex?
====

`flex` and `lex` implement the ".l" file format, which suffers from the notable problem of combining data with code - that is, C code is written _inline_ with the definition of the language's tokenization. It's impossible to use an ".l" file to generate code for any language other than C. Therefore, if you have a lex-able format and want to distribute parsers for more than one languages, you will have to write one lex file for each language - and find a program which will translate that file into the desired language correctly (hint: it's not easy).

Instead of all that, this program takes _one_ format and can create _multiple_ languages worth of lexer implementations for that format.

This came out of the author's frustration trying to even tokenize a bash file - because bash's tokenization is so closely coupled with its implementation.

How does the format and usage differ from lex?
====

* Lex files have three parts; definition, rules, and code. "el" files have no code section, and no ability for inline code.
* Regex are Go, not the original C regex used by lex.
* If any syntax is not specified, it is an error. Lex will usually just print unused sections of the file.
* Lex uses logic that involves matching a token to the longest rule. `etymu` prefers rule precedence - top to bottom, the first rule that matches the token string will be used.

What languages does this support?
====

Right now, in this earliest phase, only Go. However, the entire system _is_ set up for anyone to write a generator for any language. The author simply hasn't needed anything but Go yet, but isn't under the impression that only Go users want to write tokenizers.
