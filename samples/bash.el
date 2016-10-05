// any unicode alphanumeric
unquotedString [-:_./\\\\\\p{L}\\p{Nd}]+

interpolatedString [\"].*[\"]
literalString '.*'

//assignment [\n;][-:_./\\\\\\p{L}\\p{Nd}]+=

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
"("				{PAREN_OPEN}
")"				{PAREN_CLOSE}
"{"				{BRACE_OPEN}
"}"				{BRACE_CLOSE}
"["				{BRACKET_OPEN}
"]"				{BRACKET_CLOSE}
"[["			{EXPRESSION_OPEN}
"]]"			{EXPRESSION_CLOSE}
"=" 			{ASSIGNMENT}

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

// may not be useful,
//"\""		{QUOTE_INTERPOLATED}
//"'"		{QUOTE_LITERAL}
%%
