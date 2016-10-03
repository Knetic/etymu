etymu
====

Generates lexers from language-agnostic lexer files.

Why not flex/lex?
====

`flex` and `lex` implement the ".l" file format, which suffers from the notable problem of combining data with code. It's impossible to use an ".l" file to generate code for more than one language. Therefore, if you have a lex-able format and want to distribute parsers for more than one languages, you will have to write one lex file for each language - and find a program which will translate that file into the desired language correctly (hint: it's not easy).

Instead of all that, this program takes _one_ format and can create _multiple_ languages worth of lexer implementations for that format.

This came out of the author's frustration trying to even tokenize a bash file - because bash's tokenization is so closely coupled with its implementation.

How does the format and usage differ from lex?
====

* Lex files have three parts; definition, rules, and code. "el" files have no code section, and no ability for inline code.
* Regex are Go, not the original C regex used by lex.
* If any syntax is not specified, it is an error. Lex will usually just print unused sections of the file.
