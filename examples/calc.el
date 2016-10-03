digit 	[0-9]+
whitepace [\w+]

%%

"+"	{ return PLUS }
"-"	{ return MINUS }
"/" { return DIVIDE }
"*" { return MULTIPLY}
{digit} { return NUMERIC }

// skip whitespace, but treat as separator
{whitespace}

%%