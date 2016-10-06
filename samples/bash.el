// any unicode alphanumeric
unquotedString [-:_./\\\\\\p{L}\\p{Nd}]+

interpolatedString [\"].*[\"]
literalString '.*'

whitespace [ \t\r]+
comment #.*\n

%%

{comment}		{COMMENT}

"|"				{PIPE}
">>" 			{PIPE_APPEND}
">" 			{PIPE_OVERWRITE}
"<" 			{PIPE_INPUT}
"&>"			{PIPE_COMBINED_OVERWRITE}

"$"				{VARIABLE}
"${"			{VARIBLE_LONGFORM}
"$!"			{VARIABLE_LASTPID}
"$?"			{VARIABLE_LASTSTATUS}
"$@"			{VARIABLE_ALLARGS}
"export"		{EXPORT}
"local"			{VARIABLE_LOCAL}

"&"				{BACKGROUND}
"$("			{SUBSHELL}
"`"				{SUBSHELL_TICK}
"("				{PAREN_OPEN}
")"				{PAREN_CLOSE}
"{"				{BRACE_OPEN}
"}"				{BRACE_CLOSE}
"["				{BRACKET_OPEN}
"]"				{BRACKET_CLOSE}
"[["			{EXPRESSION_OPEN}
"]]"			{EXPRESSION_CLOSE}
"=" 			{ASSIGNMENT}
"!"				{LAST_COMMAND}

"if"			{IF}
"then"			{THEN}
"fi"			{FI}
"while"			{WHILE}
"do"			{DO}
"done"			{DONE}
"for"			{FOR}
"in"			{IN}
"function"		{FUNCTION}

"*"				{WILDCARD}
";"|"\n"|";\n"	{TERMINATOR}
","				{SEPARATOR}
"\\"			{ESCAPE}
"\\\n"			{ESCAPE_NEWLINE}

{unquotedString}		{STRING}
{interpolatedString}	{STRING_INTERPOLATED}
{literalString}			{STRING_LITERAL}

// ignore whitespace
{whitespace}
%%
