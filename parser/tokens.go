package parser

var tokenNames = map[int]string{
	ADDASSIGN:     "ADDASSIGN",
	ALIAS:         "ALIAS",
	AND:           "AND",
	ASSIGN:        "ASSIGN",
	ASTERISK:      "ASTERISK",
	BANG:          "BANG",
	BEGIN:         "BEGIN",
	CVAR:          "CVAR",
	CARET:         "CARET",
	CASE:          "CASE",
	CLASS:         "CLASS",
	COLON:         "COLON",
	COMMENT:       "COMMENT",
	CONSTANT:      "CONSTANT",
	COMMA:         "COMMA",
	DEF:           "DEF",
	DIVASSIGN:     "DIVASSIGN",
	DO:            "DO",
	DOT:           "DOT",
	DOT2:          "DOT2",
	DOT3:          "DOT3",
	ELSE:          "ELSE",
	ELSIF:         "ELSIF",
	END:           "END",
	ENSURE:        "ENSURE",
	EQ:            "EQ",
	FALSE:         "FALSE",
	FLOAT:         "FLOAT",
	GVAR:          "GVAR",
	GT:            "GT",
	GTE:           "GTE",
	INTERPBEG:     "INTERPBEG",
	INTERPEND:     "INTERPEND",
	IVAR:          "IVAR",
	IDENT:         "IDENT",
	IF:            "IF",
	INT:           "INT",
	LABEL:         "LABEL",
	LBRACE:        "LBRACE",
	LBRACKET:      "LBRACKET",
	LBRACKETSTART: "LBRACKETSTART",
	LSHIFT:        "LSHIFT",
	LSHIFTASSIGN:  "LSHIFTASSIGN",
	LOGICALAND:    "LOGICALAND",
	LOGICALOR:     "LOGICALOR",
	LPAREN:        "LPAREN",
	LPARENSTART:   "LPARENSTART",
	LT:            "LT",
	LTE:           "LTE",
	MATCH:         "MATCH",
	METHODIDENT:   "METHODIDENT",
	MINUS:         "MINUS",
	MODASSIGN:     "MODASSIGN",
	MODULE:        "MODULE",
	MODULO:        "MODULO",
	MULASSIGN:     "MULASSIGN",
	NEQ:           "NEQ",
	NEWLINE:       "NEWLINE",
	NIL:           "NIL",
	NOTMATCH:      "NOTMATCH",
	PIPE:          "PIPE",
	PLUS:          "PLUS",
	POW:           "POW",
	PRIVATE:       "PRIVATE",
	PROTECTED:     "PROTECTED",
	QMARK:         "QMARK",
	RAWWORDSBEG:   "RAWWORDSBEG",
	RBRACE:        "RBRACE",
	REGEXBEG:      "REGEXBEG",
	REGEXEND:      "REGEXEND",
	RSHIFT:        "RSHIFT",
	RSHIFTASSIGN:  "RSHIFTASSIGN",
	RESCUE:        "RESCUE",
	SCOPE:         "SCOPE",
	SLASH:         "SLASH",
	SPACESHIP:     "SPACESHIP",
	STRINGBEG:     "STRINGBEG",
	STRINGBODY:    "STRINGBODY",
	STRINGEND:     "STRINGEND",
	SUBASSIGN:     "SUBASSIGN",
	SUPER:         "SUPER",
	SYMBOL:        "SYMBOL",
	THEN:          "THEN",
	TRUE:          "TRUE",
	UNARY_NUM:     "UNARY_NUM",
	UNLESS:        "UNLESS",
	WORDSBEG:      "WORDSBEG",
}

var exprStartTokens = map[int]int{
	LPAREN:   LPARENSTART,
	LBRACKET: LBRACKETSTART,
}

var keywordModifierTokens = map[int]int{
	IF:     IF_MOD,
	UNLESS: UNLESS_MOD,
	WHILE:  WHILE_MOD,
	UNTIL:  UNTIL_MOD,
	RESCUE: RESCUE_MOD,
}
