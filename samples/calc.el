digit 	[0-9]+
whitespace [\w]

%%

"+"	{ PLUS }
"-"	{ MINUS }
"/" { DIVIDE }
"*" { MULTIPLY}
{digit} { NUMERIC }

// skip whitespace, but treat as separator
{whitespace}

%%
