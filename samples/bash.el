// any unicode alphanumeric
string [\\p{L}\\p{Nd}_/\\.]+

//string [a-zA-z0-9]+
whitespace [\\s]+

%%

"|"			{PIPE}
">>" 		{PIPE_APPEND}
">" 		{PIPE_OVERWRITE}
"<" 		{PIPE_INPUT}
"&>"		{PIPE_COMBINED_OVERWRITE}

"$"			{VARIABLE}
"&"			{BACKGROUND}
"("			{PAREN_OPEN}
")"			{PAREN_CLOSE}
"{"			{BRACE_OPEN}
"}"			{BRACE_CLOSE}
"["			{BRACKET_OPEN}
"]"			{BRACKET_CLOSE}
"[["		{EXPRESSION_OPEN}
"]]"		{EXPRESSION_CLOSE}

"if"		{IF}
"then"		{THEN}
"fi"		{FI}
"while"		{WHILE}
"do"		{DO}
"done"		{DONE}
"for"		{FOR}
"in"		{IN}
"function"	{FUNCTION}

"*"			{WILDCARD}
";"			{TERMINATOR}
","			{SEPARATOR}
"\""		{QUOTE_INTERPOLATED}
"'"			{QUOTE_LITERAL}
"\\"		{ESCAPE}
"\n"		{ESCAPE_NEWLINE}

{string}	{STRING}
{whitespace}
%%
