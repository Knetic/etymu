// any unicode alphanumeric
//string [\\p{L}\\p{Nd}_/\\.]+

string [a-zA-z0-9_/\\.]+
whitespace [\\s]+

%%

"|"			{PIPE}
">>" 		{PIPE_APPEND}
">" 		{PIPE_OVERWRITE}
"<" 		{PIPE_INPUT}
"&>"		{PIPE_COMBINED_OVERWRITE}

"$"			{VARIABLE}
"&"			{BACKGROUND}
"("			{SUBSHELL_START}
")"			{SUBSHELL_END}
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
";"			{ENDCAP}
","			{SEPARATOR}
"\""		{QUOTE_INTERPOLATED}
"'"			{QUOTE_LITERAL}
"\\"		{ESCAPE}
"\n"		{ESCAPE_NEWLINE}

{string}	{STRING}
{whitespace}
%%
