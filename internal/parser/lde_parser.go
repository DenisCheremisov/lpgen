// Code generated from LDE.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // LDE

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 27, 258,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 3, 2, 7, 2, 62, 10, 2, 12, 2, 14, 2, 65, 11, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 89, 10,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 109, 10, 5, 3, 6, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 122, 10, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 132, 10, 8, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 144, 10, 9, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 5, 22, 213, 10, 22, 3, 23, 3, 23, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 5, 24, 232, 10, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3,
	28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 2, 2,
	31, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
	38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 2, 4, 3, 2, 23, 24, 3, 2, 20,
	21, 2, 259, 2, 63, 3, 2, 2, 2, 4, 68, 3, 2, 2, 2, 6, 88, 3, 2, 2, 2, 8,
	108, 3, 2, 2, 2, 10, 110, 3, 2, 2, 2, 12, 121, 3, 2, 2, 2, 14, 131, 3,
	2, 2, 2, 16, 143, 3, 2, 2, 2, 18, 145, 3, 2, 2, 2, 20, 151, 3, 2, 2, 2,
	22, 154, 3, 2, 2, 2, 24, 158, 3, 2, 2, 2, 26, 161, 3, 2, 2, 2, 28, 165,
	3, 2, 2, 2, 30, 171, 3, 2, 2, 2, 32, 177, 3, 2, 2, 2, 34, 184, 3, 2, 2,
	2, 36, 191, 3, 2, 2, 2, 38, 196, 3, 2, 2, 2, 40, 202, 3, 2, 2, 2, 42, 212,
	3, 2, 2, 2, 44, 214, 3, 2, 2, 2, 46, 231, 3, 2, 2, 2, 48, 233, 3, 2, 2,
	2, 50, 235, 3, 2, 2, 2, 52, 241, 3, 2, 2, 2, 54, 246, 3, 2, 2, 2, 56, 251,
	3, 2, 2, 2, 58, 255, 3, 2, 2, 2, 60, 62, 5, 4, 3, 2, 61, 60, 3, 2, 2, 2,
	62, 65, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 66, 3,
	2, 2, 2, 65, 63, 3, 2, 2, 2, 66, 67, 7, 2, 2, 3, 67, 3, 3, 2, 2, 2, 68,
	69, 7, 20, 2, 2, 69, 70, 7, 3, 2, 2, 70, 71, 5, 6, 4, 2, 71, 72, 7, 4,
	2, 2, 72, 5, 3, 2, 2, 2, 73, 74, 7, 27, 2, 2, 74, 89, 5, 6, 4, 2, 75, 76,
	7, 5, 2, 2, 76, 77, 5, 6, 4, 2, 77, 78, 7, 6, 2, 2, 78, 79, 5, 6, 4, 2,
	79, 89, 3, 2, 2, 2, 80, 81, 7, 5, 2, 2, 81, 82, 5, 6, 4, 2, 82, 83, 7,
	6, 2, 2, 83, 89, 3, 2, 2, 2, 84, 85, 5, 8, 5, 2, 85, 86, 5, 6, 4, 2, 86,
	89, 3, 2, 2, 2, 87, 89, 5, 8, 5, 2, 88, 73, 3, 2, 2, 2, 88, 75, 3, 2, 2,
	2, 88, 80, 3, 2, 2, 2, 88, 84, 3, 2, 2, 2, 88, 87, 3, 2, 2, 2, 89, 7, 3,
	2, 2, 2, 90, 109, 5, 12, 7, 2, 91, 109, 5, 14, 8, 2, 92, 109, 5, 10, 6,
	2, 93, 109, 5, 16, 9, 2, 94, 109, 5, 18, 10, 2, 95, 109, 5, 24, 13, 2,
	96, 109, 5, 26, 14, 2, 97, 109, 5, 20, 11, 2, 98, 109, 5, 22, 12, 2, 99,
	109, 5, 28, 15, 2, 100, 109, 5, 30, 16, 2, 101, 109, 5, 32, 17, 2, 102,
	109, 5, 34, 18, 2, 103, 109, 5, 36, 19, 2, 104, 109, 5, 38, 20, 2, 105,
	109, 5, 40, 21, 2, 106, 109, 5, 42, 22, 2, 107, 109, 5, 44, 23, 2, 108,
	90, 3, 2, 2, 2, 108, 91, 3, 2, 2, 2, 108, 92, 3, 2, 2, 2, 108, 93, 3, 2,
	2, 2, 108, 94, 3, 2, 2, 2, 108, 95, 3, 2, 2, 2, 108, 96, 3, 2, 2, 2, 108,
	97, 3, 2, 2, 2, 108, 98, 3, 2, 2, 2, 108, 99, 3, 2, 2, 2, 108, 100, 3,
	2, 2, 2, 108, 101, 3, 2, 2, 2, 108, 102, 3, 2, 2, 2, 108, 103, 3, 2, 2,
	2, 108, 104, 3, 2, 2, 2, 108, 105, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 108,
	107, 3, 2, 2, 2, 109, 9, 3, 2, 2, 2, 110, 111, 7, 7, 2, 2, 111, 112, 7,
	24, 2, 2, 112, 11, 3, 2, 2, 2, 113, 114, 7, 8, 2, 2, 114, 115, 5, 48, 25,
	2, 115, 116, 7, 9, 2, 2, 116, 117, 7, 22, 2, 2, 117, 118, 7, 10, 2, 2,
	118, 122, 3, 2, 2, 2, 119, 120, 7, 8, 2, 2, 120, 122, 5, 48, 25, 2, 121,
	113, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 122, 13, 3, 2, 2, 2, 123, 124, 7,
	11, 2, 2, 124, 125, 5, 48, 25, 2, 125, 126, 7, 9, 2, 2, 126, 127, 7, 22,
	2, 2, 127, 128, 7, 10, 2, 2, 128, 132, 3, 2, 2, 2, 129, 130, 7, 11, 2,
	2, 130, 132, 5, 48, 25, 2, 131, 123, 3, 2, 2, 2, 131, 129, 3, 2, 2, 2,
	132, 15, 3, 2, 2, 2, 133, 134, 7, 12, 2, 2, 134, 135, 7, 8, 2, 2, 135,
	136, 5, 48, 25, 2, 136, 137, 7, 9, 2, 2, 137, 138, 7, 22, 2, 2, 138, 139,
	7, 10, 2, 2, 139, 144, 3, 2, 2, 2, 140, 141, 7, 12, 2, 2, 141, 142, 7,
	8, 2, 2, 142, 144, 5, 48, 25, 2, 143, 133, 3, 2, 2, 2, 143, 140, 3, 2,
	2, 2, 144, 17, 3, 2, 2, 2, 145, 146, 7, 13, 2, 2, 146, 147, 7, 9, 2, 2,
	147, 148, 7, 22, 2, 2, 148, 149, 7, 14, 2, 2, 149, 150, 7, 10, 2, 2, 150,
	19, 3, 2, 2, 2, 151, 152, 7, 15, 2, 2, 152, 153, 5, 46, 24, 2, 153, 21,
	3, 2, 2, 2, 154, 155, 7, 12, 2, 2, 155, 156, 7, 15, 2, 2, 156, 157, 5,
	46, 24, 2, 157, 23, 3, 2, 2, 2, 158, 159, 7, 13, 2, 2, 159, 160, 5, 46,
	24, 2, 160, 25, 3, 2, 2, 2, 161, 162, 7, 12, 2, 2, 162, 163, 7, 13, 2,
	2, 163, 164, 5, 46, 24, 2, 164, 27, 3, 2, 2, 2, 165, 166, 7, 20, 2, 2,
	166, 167, 7, 5, 2, 2, 167, 168, 5, 58, 30, 2, 168, 169, 7, 6, 2, 2, 169,
	170, 5, 46, 24, 2, 170, 29, 3, 2, 2, 2, 171, 172, 7, 20, 2, 2, 172, 173,
	7, 9, 2, 2, 173, 174, 5, 58, 30, 2, 174, 175, 7, 10, 2, 2, 175, 176, 5,
	46, 24, 2, 176, 31, 3, 2, 2, 2, 177, 178, 7, 20, 2, 2, 178, 179, 7, 5,
	2, 2, 179, 180, 5, 58, 30, 2, 180, 181, 7, 6, 2, 2, 181, 182, 7, 12, 2,
	2, 182, 183, 5, 46, 24, 2, 183, 33, 3, 2, 2, 2, 184, 185, 7, 20, 2, 2,
	185, 186, 7, 9, 2, 2, 186, 187, 5, 58, 30, 2, 187, 188, 7, 10, 2, 2, 188,
	189, 7, 12, 2, 2, 189, 190, 5, 46, 24, 2, 190, 35, 3, 2, 2, 2, 191, 192,
	7, 20, 2, 2, 192, 193, 7, 5, 2, 2, 193, 194, 5, 58, 30, 2, 194, 195, 7,
	6, 2, 2, 195, 37, 3, 2, 2, 2, 196, 197, 7, 12, 2, 2, 197, 198, 7, 20, 2,
	2, 198, 199, 7, 5, 2, 2, 199, 200, 5, 6, 4, 2, 200, 201, 7, 6, 2, 2, 201,
	39, 3, 2, 2, 2, 202, 203, 7, 12, 2, 2, 203, 204, 7, 5, 2, 2, 204, 205,
	5, 6, 4, 2, 205, 206, 7, 6, 2, 2, 206, 41, 3, 2, 2, 2, 207, 208, 7, 16,
	2, 2, 208, 213, 7, 22, 2, 2, 209, 210, 7, 16, 2, 2, 210, 211, 7, 19, 2,
	2, 211, 213, 7, 22, 2, 2, 212, 207, 3, 2, 2, 2, 212, 209, 3, 2, 2, 2, 213,
	43, 3, 2, 2, 2, 214, 215, 7, 17, 2, 2, 215, 45, 3, 2, 2, 2, 216, 217, 5,
	48, 25, 2, 217, 218, 5, 50, 26, 2, 218, 232, 3, 2, 2, 2, 219, 220, 5, 48,
	25, 2, 220, 221, 5, 52, 27, 2, 221, 232, 3, 2, 2, 2, 222, 223, 5, 48, 25,
	2, 223, 224, 5, 56, 29, 2, 224, 232, 3, 2, 2, 2, 225, 226, 5, 48, 25, 2,
	226, 227, 5, 54, 28, 2, 227, 232, 3, 2, 2, 2, 228, 232, 5, 48, 25, 2, 229,
	230, 7, 18, 2, 2, 230, 232, 5, 46, 24, 2, 231, 216, 3, 2, 2, 2, 231, 219,
	3, 2, 2, 2, 231, 222, 3, 2, 2, 2, 231, 225, 3, 2, 2, 2, 231, 228, 3, 2,
	2, 2, 231, 229, 3, 2, 2, 2, 232, 47, 3, 2, 2, 2, 233, 234, 9, 2, 2, 2,
	234, 49, 3, 2, 2, 2, 235, 236, 7, 9, 2, 2, 236, 237, 7, 22, 2, 2, 237,
	238, 7, 14, 2, 2, 238, 239, 7, 22, 2, 2, 239, 240, 7, 10, 2, 2, 240, 51,
	3, 2, 2, 2, 241, 242, 7, 9, 2, 2, 242, 243, 7, 14, 2, 2, 243, 244, 7, 22,
	2, 2, 244, 245, 7, 10, 2, 2, 245, 53, 3, 2, 2, 2, 246, 247, 7, 9, 2, 2,
	247, 248, 7, 22, 2, 2, 248, 249, 7, 14, 2, 2, 249, 250, 7, 10, 2, 2, 250,
	55, 3, 2, 2, 2, 251, 252, 7, 9, 2, 2, 252, 253, 7, 22, 2, 2, 253, 254,
	7, 10, 2, 2, 254, 57, 3, 2, 2, 2, 255, 256, 9, 3, 2, 2, 256, 59, 3, 2,
	2, 2, 10, 63, 88, 108, 121, 131, 143, 212, 231,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "';'", "'('", "')'", "'*'", "'^'", "'['", "']'", "'@'", "'?'",
	"'_'", "':'", "'..'", "'%'", "'$'", "'~'", "", "", "", "", "", "", "",
	"", "'!'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "ComparisonOperator",
	"Identifier", "IdentifierWithFraction", "IntLit", "StringLit", "CharLit",
	"WS", "LineComment", "Stress",
}

var ruleNames = []string{
	"rules", "atomicRule", "baseAction", "atomicAction", "passHeadingCharacters",
	"passTargetPrefix", "checkTargetPrefix", "mayBePassTargetPrefix", "passChars",
	"goUntil", "mayGoUntil", "passUntil", "mayPassUntil", "takeUntil", "takeUntilIncluding",
	"takeUntilOrRest", "takeUntilIncludingOrRest", "takeUntilRest", "optionalNamedArea",
	"optionalArea", "restCheck", "atEnd", "target", "targetLit", "bound", "limit",
	"jump", "exact", "fieldType",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type LDEParser struct {
	*antlr.BaseParser
}

func NewLDEParser(input antlr.TokenStream) *LDEParser {
	this := new(LDEParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "LDE.g4"

	return this
}

// LDEParser tokens.
const (
	LDEParserEOF                    = antlr.TokenEOF
	LDEParserT__0                   = 1
	LDEParserT__1                   = 2
	LDEParserT__2                   = 3
	LDEParserT__3                   = 4
	LDEParserT__4                   = 5
	LDEParserT__5                   = 6
	LDEParserT__6                   = 7
	LDEParserT__7                   = 8
	LDEParserT__8                   = 9
	LDEParserT__9                   = 10
	LDEParserT__10                  = 11
	LDEParserT__11                  = 12
	LDEParserT__12                  = 13
	LDEParserT__13                  = 14
	LDEParserT__14                  = 15
	LDEParserT__15                  = 16
	LDEParserComparisonOperator     = 17
	LDEParserIdentifier             = 18
	LDEParserIdentifierWithFraction = 19
	LDEParserIntLit                 = 20
	LDEParserStringLit              = 21
	LDEParserCharLit                = 22
	LDEParserWS                     = 23
	LDEParserLineComment            = 24
	LDEParserStress                 = 25
)

// LDEParser rules.
const (
	LDEParserRULE_rules                    = 0
	LDEParserRULE_atomicRule               = 1
	LDEParserRULE_baseAction               = 2
	LDEParserRULE_atomicAction             = 3
	LDEParserRULE_passHeadingCharacters    = 4
	LDEParserRULE_passTargetPrefix         = 5
	LDEParserRULE_checkTargetPrefix        = 6
	LDEParserRULE_mayBePassTargetPrefix    = 7
	LDEParserRULE_passChars                = 8
	LDEParserRULE_goUntil                  = 9
	LDEParserRULE_mayGoUntil               = 10
	LDEParserRULE_passUntil                = 11
	LDEParserRULE_mayPassUntil             = 12
	LDEParserRULE_takeUntil                = 13
	LDEParserRULE_takeUntilIncluding       = 14
	LDEParserRULE_takeUntilOrRest          = 15
	LDEParserRULE_takeUntilIncludingOrRest = 16
	LDEParserRULE_takeUntilRest            = 17
	LDEParserRULE_optionalNamedArea        = 18
	LDEParserRULE_optionalArea             = 19
	LDEParserRULE_restCheck                = 20
	LDEParserRULE_atEnd                    = 21
	LDEParserRULE_target                   = 22
	LDEParserRULE_targetLit                = 23
	LDEParserRULE_bound                    = 24
	LDEParserRULE_limit                    = 25
	LDEParserRULE_jump                     = 26
	LDEParserRULE_exact                    = 27
	LDEParserRULE_fieldType                = 28
)

// IRulesContext is an interface to support dynamic dispatch.
type IRulesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRulesContext differentiates from other interfaces.
	IsRulesContext()
}

type RulesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRulesContext() *RulesContext {
	var p = new(RulesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_rules
	return p
}

func (*RulesContext) IsRulesContext() {}

func NewRulesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RulesContext {
	var p = new(RulesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_rules

	return p
}

func (s *RulesContext) GetParser() antlr.Parser { return s.parser }

func (s *RulesContext) EOF() antlr.TerminalNode {
	return s.GetToken(LDEParserEOF, 0)
}

func (s *RulesContext) AllAtomicRule() []IAtomicRuleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtomicRuleContext)(nil)).Elem())
	var tst = make([]IAtomicRuleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtomicRuleContext)
		}
	}

	return tst
}

func (s *RulesContext) AtomicRule(i int) IAtomicRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomicRuleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtomicRuleContext)
}

func (s *RulesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RulesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RulesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterRules(s)
	}
}

func (s *RulesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitRules(s)
	}
}

func (p *LDEParser) Rules() (localctx IRulesContext) {
	localctx = NewRulesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LDEParserRULE_rules)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LDEParserIdentifier {
		{
			p.SetState(58)
			p.AtomicRule()
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(64)
		p.Match(LDEParserEOF)
	}

	return localctx
}

// IAtomicRuleContext is an interface to support dynamic dispatch.
type IAtomicRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomicRuleContext differentiates from other interfaces.
	IsAtomicRuleContext()
}

type AtomicRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomicRuleContext() *AtomicRuleContext {
	var p = new(AtomicRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_atomicRule
	return p
}

func (*AtomicRuleContext) IsAtomicRuleContext() {}

func NewAtomicRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomicRuleContext {
	var p = new(AtomicRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_atomicRule

	return p
}

func (s *AtomicRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomicRuleContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LDEParserIdentifier, 0)
}

func (s *AtomicRuleContext) BaseAction() IBaseActionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBaseActionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBaseActionContext)
}

func (s *AtomicRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomicRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomicRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterAtomicRule(s)
	}
}

func (s *AtomicRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitAtomicRule(s)
	}
}

func (p *LDEParser) AtomicRule() (localctx IAtomicRuleContext) {
	localctx = NewAtomicRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LDEParserRULE_atomicRule)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(LDEParserIdentifier)
	}
	{
		p.SetState(67)
		p.Match(LDEParserT__0)
	}
	{
		p.SetState(68)
		p.BaseAction()
	}
	{
		p.SetState(69)
		p.Match(LDEParserT__1)
	}

	return localctx
}

// IBaseActionContext is an interface to support dynamic dispatch.
type IBaseActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBaseActionContext differentiates from other interfaces.
	IsBaseActionContext()
}

type BaseActionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBaseActionContext() *BaseActionContext {
	var p = new(BaseActionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_baseAction
	return p
}

func (*BaseActionContext) IsBaseActionContext() {}

func NewBaseActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaseActionContext {
	var p = new(BaseActionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_baseAction

	return p
}

func (s *BaseActionContext) GetParser() antlr.Parser { return s.parser }

func (s *BaseActionContext) Stress() antlr.TerminalNode {
	return s.GetToken(LDEParserStress, 0)
}

func (s *BaseActionContext) AllBaseAction() []IBaseActionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBaseActionContext)(nil)).Elem())
	var tst = make([]IBaseActionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBaseActionContext)
		}
	}

	return tst
}

func (s *BaseActionContext) BaseAction(i int) IBaseActionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBaseActionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBaseActionContext)
}

func (s *BaseActionContext) AtomicAction() IAtomicActionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomicActionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomicActionContext)
}

func (s *BaseActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BaseActionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterBaseAction(s)
	}
}

func (s *BaseActionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitBaseAction(s)
	}
}

func (p *LDEParser) BaseAction() (localctx IBaseActionContext) {
	localctx = NewBaseActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LDEParserRULE_baseAction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)
			p.Match(LDEParserStress)
		}
		{
			p.SetState(72)
			p.BaseAction()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.Match(LDEParserT__2)
		}
		{
			p.SetState(74)
			p.BaseAction()
		}
		{
			p.SetState(75)
			p.Match(LDEParserT__3)
		}
		{
			p.SetState(76)
			p.BaseAction()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(78)
			p.Match(LDEParserT__2)
		}
		{
			p.SetState(79)
			p.BaseAction()
		}
		{
			p.SetState(80)
			p.Match(LDEParserT__3)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(82)
			p.AtomicAction()
		}
		{
			p.SetState(83)
			p.BaseAction()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(85)
			p.AtomicAction()
		}

	}

	return localctx
}

// IAtomicActionContext is an interface to support dynamic dispatch.
type IAtomicActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomicActionContext differentiates from other interfaces.
	IsAtomicActionContext()
}

type AtomicActionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomicActionContext() *AtomicActionContext {
	var p = new(AtomicActionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_atomicAction
	return p
}

func (*AtomicActionContext) IsAtomicActionContext() {}

func NewAtomicActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomicActionContext {
	var p = new(AtomicActionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_atomicAction

	return p
}

func (s *AtomicActionContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomicActionContext) PassTargetPrefix() IPassTargetPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPassTargetPrefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPassTargetPrefixContext)
}

func (s *AtomicActionContext) CheckTargetPrefix() ICheckTargetPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICheckTargetPrefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICheckTargetPrefixContext)
}

func (s *AtomicActionContext) PassHeadingCharacters() IPassHeadingCharactersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPassHeadingCharactersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPassHeadingCharactersContext)
}

func (s *AtomicActionContext) MayBePassTargetPrefix() IMayBePassTargetPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMayBePassTargetPrefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMayBePassTargetPrefixContext)
}

func (s *AtomicActionContext) PassChars() IPassCharsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPassCharsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPassCharsContext)
}

func (s *AtomicActionContext) PassUntil() IPassUntilContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPassUntilContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPassUntilContext)
}

func (s *AtomicActionContext) MayPassUntil() IMayPassUntilContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMayPassUntilContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMayPassUntilContext)
}

func (s *AtomicActionContext) GoUntil() IGoUntilContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGoUntilContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGoUntilContext)
}

func (s *AtomicActionContext) MayGoUntil() IMayGoUntilContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMayGoUntilContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMayGoUntilContext)
}

func (s *AtomicActionContext) TakeUntil() ITakeUntilContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITakeUntilContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITakeUntilContext)
}

func (s *AtomicActionContext) TakeUntilIncluding() ITakeUntilIncludingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITakeUntilIncludingContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITakeUntilIncludingContext)
}

func (s *AtomicActionContext) TakeUntilOrRest() ITakeUntilOrRestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITakeUntilOrRestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITakeUntilOrRestContext)
}

func (s *AtomicActionContext) TakeUntilIncludingOrRest() ITakeUntilIncludingOrRestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITakeUntilIncludingOrRestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITakeUntilIncludingOrRestContext)
}

func (s *AtomicActionContext) TakeUntilRest() ITakeUntilRestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITakeUntilRestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITakeUntilRestContext)
}

func (s *AtomicActionContext) OptionalNamedArea() IOptionalNamedAreaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalNamedAreaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalNamedAreaContext)
}

func (s *AtomicActionContext) OptionalArea() IOptionalAreaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalAreaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalAreaContext)
}

func (s *AtomicActionContext) RestCheck() IRestCheckContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRestCheckContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRestCheckContext)
}

func (s *AtomicActionContext) AtEnd() IAtEndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtEndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtEndContext)
}

func (s *AtomicActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomicActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomicActionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterAtomicAction(s)
	}
}

func (s *AtomicActionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitAtomicAction(s)
	}
}

func (p *LDEParser) AtomicAction() (localctx IAtomicActionContext) {
	localctx = NewAtomicActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LDEParserRULE_atomicAction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.PassTargetPrefix()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.CheckTargetPrefix()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(90)
			p.PassHeadingCharacters()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(91)
			p.MayBePassTargetPrefix()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(92)
			p.PassChars()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(93)
			p.PassUntil()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(94)
			p.MayPassUntil()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(95)
			p.GoUntil()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(96)
			p.MayGoUntil()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(97)
			p.TakeUntil()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(98)
			p.TakeUntilIncluding()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(99)
			p.TakeUntilOrRest()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(100)
			p.TakeUntilIncludingOrRest()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(101)
			p.TakeUntilRest()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(102)
			p.OptionalNamedArea()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(103)
			p.OptionalArea()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(104)
			p.RestCheck()
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(105)
			p.AtEnd()
		}

	}

	return localctx
}

// IPassHeadingCharactersContext is an interface to support dynamic dispatch.
type IPassHeadingCharactersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPassHeadingCharactersContext differentiates from other interfaces.
	IsPassHeadingCharactersContext()
}

type PassHeadingCharactersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPassHeadingCharactersContext() *PassHeadingCharactersContext {
	var p = new(PassHeadingCharactersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_passHeadingCharacters
	return p
}

func (*PassHeadingCharactersContext) IsPassHeadingCharactersContext() {}

func NewPassHeadingCharactersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PassHeadingCharactersContext {
	var p = new(PassHeadingCharactersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_passHeadingCharacters

	return p
}

func (s *PassHeadingCharactersContext) GetParser() antlr.Parser { return s.parser }

func (s *PassHeadingCharactersContext) CharLit() antlr.TerminalNode {
	return s.GetToken(LDEParserCharLit, 0)
}

func (s *PassHeadingCharactersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassHeadingCharactersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PassHeadingCharactersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterPassHeadingCharacters(s)
	}
}

func (s *PassHeadingCharactersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitPassHeadingCharacters(s)
	}
}

func (p *LDEParser) PassHeadingCharacters() (localctx IPassHeadingCharactersContext) {
	localctx = NewPassHeadingCharactersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LDEParserRULE_passHeadingCharacters)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(LDEParserT__4)
	}
	{
		p.SetState(109)
		p.Match(LDEParserCharLit)
	}

	return localctx
}

// IPassTargetPrefixContext is an interface to support dynamic dispatch.
type IPassTargetPrefixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPassTargetPrefixContext differentiates from other interfaces.
	IsPassTargetPrefixContext()
}

type PassTargetPrefixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPassTargetPrefixContext() *PassTargetPrefixContext {
	var p = new(PassTargetPrefixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_passTargetPrefix
	return p
}

func (*PassTargetPrefixContext) IsPassTargetPrefixContext() {}

func NewPassTargetPrefixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PassTargetPrefixContext {
	var p = new(PassTargetPrefixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_passTargetPrefix

	return p
}

func (s *PassTargetPrefixContext) GetParser() antlr.Parser { return s.parser }

func (s *PassTargetPrefixContext) TargetLit() ITargetLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetLitContext)
}

func (s *PassTargetPrefixContext) IntLit() antlr.TerminalNode {
	return s.GetToken(LDEParserIntLit, 0)
}

func (s *PassTargetPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassTargetPrefixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PassTargetPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterPassTargetPrefix(s)
	}
}

func (s *PassTargetPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitPassTargetPrefix(s)
	}
}

func (p *LDEParser) PassTargetPrefix() (localctx IPassTargetPrefixContext) {
	localctx = NewPassTargetPrefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LDEParserRULE_passTargetPrefix)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.Match(LDEParserT__5)
		}
		{
			p.SetState(112)
			p.TargetLit()
		}
		{
			p.SetState(113)
			p.Match(LDEParserT__6)
		}
		{
			p.SetState(114)
			p.Match(LDEParserIntLit)
		}
		{
			p.SetState(115)
			p.Match(LDEParserT__7)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.Match(LDEParserT__5)
		}
		{
			p.SetState(118)
			p.TargetLit()
		}

	}

	return localctx
}

// ICheckTargetPrefixContext is an interface to support dynamic dispatch.
type ICheckTargetPrefixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCheckTargetPrefixContext differentiates from other interfaces.
	IsCheckTargetPrefixContext()
}

type CheckTargetPrefixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCheckTargetPrefixContext() *CheckTargetPrefixContext {
	var p = new(CheckTargetPrefixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_checkTargetPrefix
	return p
}

func (*CheckTargetPrefixContext) IsCheckTargetPrefixContext() {}

func NewCheckTargetPrefixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CheckTargetPrefixContext {
	var p = new(CheckTargetPrefixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_checkTargetPrefix

	return p
}

func (s *CheckTargetPrefixContext) GetParser() antlr.Parser { return s.parser }

func (s *CheckTargetPrefixContext) TargetLit() ITargetLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetLitContext)
}

func (s *CheckTargetPrefixContext) IntLit() antlr.TerminalNode {
	return s.GetToken(LDEParserIntLit, 0)
}

func (s *CheckTargetPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CheckTargetPrefixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CheckTargetPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterCheckTargetPrefix(s)
	}
}

func (s *CheckTargetPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitCheckTargetPrefix(s)
	}
}

func (p *LDEParser) CheckTargetPrefix() (localctx ICheckTargetPrefixContext) {
	localctx = NewCheckTargetPrefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LDEParserRULE_checkTargetPrefix)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.Match(LDEParserT__8)
		}
		{
			p.SetState(122)
			p.TargetLit()
		}
		{
			p.SetState(123)
			p.Match(LDEParserT__6)
		}
		{
			p.SetState(124)
			p.Match(LDEParserIntLit)
		}
		{
			p.SetState(125)
			p.Match(LDEParserT__7)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(127)
			p.Match(LDEParserT__8)
		}
		{
			p.SetState(128)
			p.TargetLit()
		}

	}

	return localctx
}

// IMayBePassTargetPrefixContext is an interface to support dynamic dispatch.
type IMayBePassTargetPrefixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMayBePassTargetPrefixContext differentiates from other interfaces.
	IsMayBePassTargetPrefixContext()
}

type MayBePassTargetPrefixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMayBePassTargetPrefixContext() *MayBePassTargetPrefixContext {
	var p = new(MayBePassTargetPrefixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_mayBePassTargetPrefix
	return p
}

func (*MayBePassTargetPrefixContext) IsMayBePassTargetPrefixContext() {}

func NewMayBePassTargetPrefixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MayBePassTargetPrefixContext {
	var p = new(MayBePassTargetPrefixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_mayBePassTargetPrefix

	return p
}

func (s *MayBePassTargetPrefixContext) GetParser() antlr.Parser { return s.parser }

func (s *MayBePassTargetPrefixContext) TargetLit() ITargetLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetLitContext)
}

func (s *MayBePassTargetPrefixContext) IntLit() antlr.TerminalNode {
	return s.GetToken(LDEParserIntLit, 0)
}

func (s *MayBePassTargetPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MayBePassTargetPrefixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MayBePassTargetPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterMayBePassTargetPrefix(s)
	}
}

func (s *MayBePassTargetPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitMayBePassTargetPrefix(s)
	}
}

func (p *LDEParser) MayBePassTargetPrefix() (localctx IMayBePassTargetPrefixContext) {
	localctx = NewMayBePassTargetPrefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LDEParserRULE_mayBePassTargetPrefix)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)
			p.Match(LDEParserT__9)
		}
		{
			p.SetState(132)
			p.Match(LDEParserT__5)
		}
		{
			p.SetState(133)
			p.TargetLit()
		}
		{
			p.SetState(134)
			p.Match(LDEParserT__6)
		}
		{
			p.SetState(135)
			p.Match(LDEParserIntLit)
		}
		{
			p.SetState(136)
			p.Match(LDEParserT__7)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(138)
			p.Match(LDEParserT__9)
		}
		{
			p.SetState(139)
			p.Match(LDEParserT__5)
		}
		{
			p.SetState(140)
			p.TargetLit()
		}

	}

	return localctx
}

// IPassCharsContext is an interface to support dynamic dispatch.
type IPassCharsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPassCharsContext differentiates from other interfaces.
	IsPassCharsContext()
}

type PassCharsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPassCharsContext() *PassCharsContext {
	var p = new(PassCharsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_passChars
	return p
}

func (*PassCharsContext) IsPassCharsContext() {}

func NewPassCharsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PassCharsContext {
	var p = new(PassCharsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_passChars

	return p
}

func (s *PassCharsContext) GetParser() antlr.Parser { return s.parser }

func (s *PassCharsContext) IntLit() antlr.TerminalNode {
	return s.GetToken(LDEParserIntLit, 0)
}

func (s *PassCharsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassCharsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PassCharsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterPassChars(s)
	}
}

func (s *PassCharsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitPassChars(s)
	}
}

func (p *LDEParser) PassChars() (localctx IPassCharsContext) {
	localctx = NewPassCharsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LDEParserRULE_passChars)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(LDEParserT__10)
	}
	{
		p.SetState(144)
		p.Match(LDEParserT__6)
	}
	{
		p.SetState(145)
		p.Match(LDEParserIntLit)
	}
	{
		p.SetState(146)
		p.Match(LDEParserT__11)
	}
	{
		p.SetState(147)
		p.Match(LDEParserT__7)
	}

	return localctx
}

// IGoUntilContext is an interface to support dynamic dispatch.
type IGoUntilContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGoUntilContext differentiates from other interfaces.
	IsGoUntilContext()
}

type GoUntilContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGoUntilContext() *GoUntilContext {
	var p = new(GoUntilContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_goUntil
	return p
}

func (*GoUntilContext) IsGoUntilContext() {}

func NewGoUntilContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GoUntilContext {
	var p = new(GoUntilContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_goUntil

	return p
}

func (s *GoUntilContext) GetParser() antlr.Parser { return s.parser }

func (s *GoUntilContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *GoUntilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GoUntilContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GoUntilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterGoUntil(s)
	}
}

func (s *GoUntilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitGoUntil(s)
	}
}

func (p *LDEParser) GoUntil() (localctx IGoUntilContext) {
	localctx = NewGoUntilContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LDEParserRULE_goUntil)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(LDEParserT__12)
	}
	{
		p.SetState(150)
		p.Target()
	}

	return localctx
}

// IMayGoUntilContext is an interface to support dynamic dispatch.
type IMayGoUntilContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMayGoUntilContext differentiates from other interfaces.
	IsMayGoUntilContext()
}

type MayGoUntilContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMayGoUntilContext() *MayGoUntilContext {
	var p = new(MayGoUntilContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_mayGoUntil
	return p
}

func (*MayGoUntilContext) IsMayGoUntilContext() {}

func NewMayGoUntilContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MayGoUntilContext {
	var p = new(MayGoUntilContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_mayGoUntil

	return p
}

func (s *MayGoUntilContext) GetParser() antlr.Parser { return s.parser }

func (s *MayGoUntilContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *MayGoUntilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MayGoUntilContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MayGoUntilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterMayGoUntil(s)
	}
}

func (s *MayGoUntilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitMayGoUntil(s)
	}
}

func (p *LDEParser) MayGoUntil() (localctx IMayGoUntilContext) {
	localctx = NewMayGoUntilContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LDEParserRULE_mayGoUntil)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Match(LDEParserT__9)
	}
	{
		p.SetState(153)
		p.Match(LDEParserT__12)
	}
	{
		p.SetState(154)
		p.Target()
	}

	return localctx
}

// IPassUntilContext is an interface to support dynamic dispatch.
type IPassUntilContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPassUntilContext differentiates from other interfaces.
	IsPassUntilContext()
}

type PassUntilContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPassUntilContext() *PassUntilContext {
	var p = new(PassUntilContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_passUntil
	return p
}

func (*PassUntilContext) IsPassUntilContext() {}

func NewPassUntilContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PassUntilContext {
	var p = new(PassUntilContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_passUntil

	return p
}

func (s *PassUntilContext) GetParser() antlr.Parser { return s.parser }

func (s *PassUntilContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *PassUntilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassUntilContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PassUntilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterPassUntil(s)
	}
}

func (s *PassUntilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitPassUntil(s)
	}
}

func (p *LDEParser) PassUntil() (localctx IPassUntilContext) {
	localctx = NewPassUntilContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LDEParserRULE_passUntil)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(LDEParserT__10)
	}
	{
		p.SetState(157)
		p.Target()
	}

	return localctx
}

// IMayPassUntilContext is an interface to support dynamic dispatch.
type IMayPassUntilContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMayPassUntilContext differentiates from other interfaces.
	IsMayPassUntilContext()
}

type MayPassUntilContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMayPassUntilContext() *MayPassUntilContext {
	var p = new(MayPassUntilContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_mayPassUntil
	return p
}

func (*MayPassUntilContext) IsMayPassUntilContext() {}

func NewMayPassUntilContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MayPassUntilContext {
	var p = new(MayPassUntilContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_mayPassUntil

	return p
}

func (s *MayPassUntilContext) GetParser() antlr.Parser { return s.parser }

func (s *MayPassUntilContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *MayPassUntilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MayPassUntilContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MayPassUntilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterMayPassUntil(s)
	}
}

func (s *MayPassUntilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitMayPassUntil(s)
	}
}

func (p *LDEParser) MayPassUntil() (localctx IMayPassUntilContext) {
	localctx = NewMayPassUntilContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LDEParserRULE_mayPassUntil)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Match(LDEParserT__9)
	}
	{
		p.SetState(160)
		p.Match(LDEParserT__10)
	}
	{
		p.SetState(161)
		p.Target()
	}

	return localctx
}

// ITakeUntilContext is an interface to support dynamic dispatch.
type ITakeUntilContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTakeUntilContext differentiates from other interfaces.
	IsTakeUntilContext()
}

type TakeUntilContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTakeUntilContext() *TakeUntilContext {
	var p = new(TakeUntilContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_takeUntil
	return p
}

func (*TakeUntilContext) IsTakeUntilContext() {}

func NewTakeUntilContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TakeUntilContext {
	var p = new(TakeUntilContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_takeUntil

	return p
}

func (s *TakeUntilContext) GetParser() antlr.Parser { return s.parser }

func (s *TakeUntilContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LDEParserIdentifier, 0)
}

func (s *TakeUntilContext) FieldType() IFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *TakeUntilContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *TakeUntilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TakeUntilContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TakeUntilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterTakeUntil(s)
	}
}

func (s *TakeUntilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitTakeUntil(s)
	}
}

func (p *LDEParser) TakeUntil() (localctx ITakeUntilContext) {
	localctx = NewTakeUntilContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LDEParserRULE_takeUntil)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(LDEParserIdentifier)
	}
	{
		p.SetState(164)
		p.Match(LDEParserT__2)
	}
	{
		p.SetState(165)
		p.FieldType()
	}
	{
		p.SetState(166)
		p.Match(LDEParserT__3)
	}
	{
		p.SetState(167)
		p.Target()
	}

	return localctx
}

// ITakeUntilIncludingContext is an interface to support dynamic dispatch.
type ITakeUntilIncludingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTakeUntilIncludingContext differentiates from other interfaces.
	IsTakeUntilIncludingContext()
}

type TakeUntilIncludingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTakeUntilIncludingContext() *TakeUntilIncludingContext {
	var p = new(TakeUntilIncludingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_takeUntilIncluding
	return p
}

func (*TakeUntilIncludingContext) IsTakeUntilIncludingContext() {}

func NewTakeUntilIncludingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TakeUntilIncludingContext {
	var p = new(TakeUntilIncludingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_takeUntilIncluding

	return p
}

func (s *TakeUntilIncludingContext) GetParser() antlr.Parser { return s.parser }

func (s *TakeUntilIncludingContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LDEParserIdentifier, 0)
}

func (s *TakeUntilIncludingContext) FieldType() IFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *TakeUntilIncludingContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *TakeUntilIncludingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TakeUntilIncludingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TakeUntilIncludingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterTakeUntilIncluding(s)
	}
}

func (s *TakeUntilIncludingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitTakeUntilIncluding(s)
	}
}

func (p *LDEParser) TakeUntilIncluding() (localctx ITakeUntilIncludingContext) {
	localctx = NewTakeUntilIncludingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LDEParserRULE_takeUntilIncluding)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(LDEParserIdentifier)
	}
	{
		p.SetState(170)
		p.Match(LDEParserT__6)
	}
	{
		p.SetState(171)
		p.FieldType()
	}
	{
		p.SetState(172)
		p.Match(LDEParserT__7)
	}
	{
		p.SetState(173)
		p.Target()
	}

	return localctx
}

// ITakeUntilOrRestContext is an interface to support dynamic dispatch.
type ITakeUntilOrRestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTakeUntilOrRestContext differentiates from other interfaces.
	IsTakeUntilOrRestContext()
}

type TakeUntilOrRestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTakeUntilOrRestContext() *TakeUntilOrRestContext {
	var p = new(TakeUntilOrRestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_takeUntilOrRest
	return p
}

func (*TakeUntilOrRestContext) IsTakeUntilOrRestContext() {}

func NewTakeUntilOrRestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TakeUntilOrRestContext {
	var p = new(TakeUntilOrRestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_takeUntilOrRest

	return p
}

func (s *TakeUntilOrRestContext) GetParser() antlr.Parser { return s.parser }

func (s *TakeUntilOrRestContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LDEParserIdentifier, 0)
}

func (s *TakeUntilOrRestContext) FieldType() IFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *TakeUntilOrRestContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *TakeUntilOrRestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TakeUntilOrRestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TakeUntilOrRestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterTakeUntilOrRest(s)
	}
}

func (s *TakeUntilOrRestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitTakeUntilOrRest(s)
	}
}

func (p *LDEParser) TakeUntilOrRest() (localctx ITakeUntilOrRestContext) {
	localctx = NewTakeUntilOrRestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LDEParserRULE_takeUntilOrRest)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(LDEParserIdentifier)
	}
	{
		p.SetState(176)
		p.Match(LDEParserT__2)
	}
	{
		p.SetState(177)
		p.FieldType()
	}
	{
		p.SetState(178)
		p.Match(LDEParserT__3)
	}
	{
		p.SetState(179)
		p.Match(LDEParserT__9)
	}
	{
		p.SetState(180)
		p.Target()
	}

	return localctx
}

// ITakeUntilIncludingOrRestContext is an interface to support dynamic dispatch.
type ITakeUntilIncludingOrRestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTakeUntilIncludingOrRestContext differentiates from other interfaces.
	IsTakeUntilIncludingOrRestContext()
}

type TakeUntilIncludingOrRestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTakeUntilIncludingOrRestContext() *TakeUntilIncludingOrRestContext {
	var p = new(TakeUntilIncludingOrRestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_takeUntilIncludingOrRest
	return p
}

func (*TakeUntilIncludingOrRestContext) IsTakeUntilIncludingOrRestContext() {}

func NewTakeUntilIncludingOrRestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TakeUntilIncludingOrRestContext {
	var p = new(TakeUntilIncludingOrRestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_takeUntilIncludingOrRest

	return p
}

func (s *TakeUntilIncludingOrRestContext) GetParser() antlr.Parser { return s.parser }

func (s *TakeUntilIncludingOrRestContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LDEParserIdentifier, 0)
}

func (s *TakeUntilIncludingOrRestContext) FieldType() IFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *TakeUntilIncludingOrRestContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *TakeUntilIncludingOrRestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TakeUntilIncludingOrRestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TakeUntilIncludingOrRestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterTakeUntilIncludingOrRest(s)
	}
}

func (s *TakeUntilIncludingOrRestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitTakeUntilIncludingOrRest(s)
	}
}

func (p *LDEParser) TakeUntilIncludingOrRest() (localctx ITakeUntilIncludingOrRestContext) {
	localctx = NewTakeUntilIncludingOrRestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, LDEParserRULE_takeUntilIncludingOrRest)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(LDEParserIdentifier)
	}
	{
		p.SetState(183)
		p.Match(LDEParserT__6)
	}
	{
		p.SetState(184)
		p.FieldType()
	}
	{
		p.SetState(185)
		p.Match(LDEParserT__7)
	}
	{
		p.SetState(186)
		p.Match(LDEParserT__9)
	}
	{
		p.SetState(187)
		p.Target()
	}

	return localctx
}

// ITakeUntilRestContext is an interface to support dynamic dispatch.
type ITakeUntilRestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTakeUntilRestContext differentiates from other interfaces.
	IsTakeUntilRestContext()
}

type TakeUntilRestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTakeUntilRestContext() *TakeUntilRestContext {
	var p = new(TakeUntilRestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_takeUntilRest
	return p
}

func (*TakeUntilRestContext) IsTakeUntilRestContext() {}

func NewTakeUntilRestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TakeUntilRestContext {
	var p = new(TakeUntilRestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_takeUntilRest

	return p
}

func (s *TakeUntilRestContext) GetParser() antlr.Parser { return s.parser }

func (s *TakeUntilRestContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LDEParserIdentifier, 0)
}

func (s *TakeUntilRestContext) FieldType() IFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *TakeUntilRestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TakeUntilRestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TakeUntilRestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterTakeUntilRest(s)
	}
}

func (s *TakeUntilRestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitTakeUntilRest(s)
	}
}

func (p *LDEParser) TakeUntilRest() (localctx ITakeUntilRestContext) {
	localctx = NewTakeUntilRestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, LDEParserRULE_takeUntilRest)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.Match(LDEParserIdentifier)
	}
	{
		p.SetState(190)
		p.Match(LDEParserT__2)
	}
	{
		p.SetState(191)
		p.FieldType()
	}
	{
		p.SetState(192)
		p.Match(LDEParserT__3)
	}

	return localctx
}

// IOptionalNamedAreaContext is an interface to support dynamic dispatch.
type IOptionalNamedAreaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionalNamedAreaContext differentiates from other interfaces.
	IsOptionalNamedAreaContext()
}

type OptionalNamedAreaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionalNamedAreaContext() *OptionalNamedAreaContext {
	var p = new(OptionalNamedAreaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_optionalNamedArea
	return p
}

func (*OptionalNamedAreaContext) IsOptionalNamedAreaContext() {}

func NewOptionalNamedAreaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionalNamedAreaContext {
	var p = new(OptionalNamedAreaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_optionalNamedArea

	return p
}

func (s *OptionalNamedAreaContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionalNamedAreaContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LDEParserIdentifier, 0)
}

func (s *OptionalNamedAreaContext) BaseAction() IBaseActionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBaseActionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBaseActionContext)
}

func (s *OptionalNamedAreaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionalNamedAreaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionalNamedAreaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterOptionalNamedArea(s)
	}
}

func (s *OptionalNamedAreaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitOptionalNamedArea(s)
	}
}

func (p *LDEParser) OptionalNamedArea() (localctx IOptionalNamedAreaContext) {
	localctx = NewOptionalNamedAreaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, LDEParserRULE_optionalNamedArea)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
		p.Match(LDEParserT__9)
	}
	{
		p.SetState(195)
		p.Match(LDEParserIdentifier)
	}
	{
		p.SetState(196)
		p.Match(LDEParserT__2)
	}
	{
		p.SetState(197)
		p.BaseAction()
	}
	{
		p.SetState(198)
		p.Match(LDEParserT__3)
	}

	return localctx
}

// IOptionalAreaContext is an interface to support dynamic dispatch.
type IOptionalAreaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionalAreaContext differentiates from other interfaces.
	IsOptionalAreaContext()
}

type OptionalAreaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionalAreaContext() *OptionalAreaContext {
	var p = new(OptionalAreaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_optionalArea
	return p
}

func (*OptionalAreaContext) IsOptionalAreaContext() {}

func NewOptionalAreaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionalAreaContext {
	var p = new(OptionalAreaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_optionalArea

	return p
}

func (s *OptionalAreaContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionalAreaContext) BaseAction() IBaseActionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBaseActionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBaseActionContext)
}

func (s *OptionalAreaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionalAreaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionalAreaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterOptionalArea(s)
	}
}

func (s *OptionalAreaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitOptionalArea(s)
	}
}

func (p *LDEParser) OptionalArea() (localctx IOptionalAreaContext) {
	localctx = NewOptionalAreaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, LDEParserRULE_optionalArea)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(LDEParserT__9)
	}
	{
		p.SetState(201)
		p.Match(LDEParserT__2)
	}
	{
		p.SetState(202)
		p.BaseAction()
	}
	{
		p.SetState(203)
		p.Match(LDEParserT__3)
	}

	return localctx
}

// IRestCheckContext is an interface to support dynamic dispatch.
type IRestCheckContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRestCheckContext differentiates from other interfaces.
	IsRestCheckContext()
}

type RestCheckContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRestCheckContext() *RestCheckContext {
	var p = new(RestCheckContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_restCheck
	return p
}

func (*RestCheckContext) IsRestCheckContext() {}

func NewRestCheckContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RestCheckContext {
	var p = new(RestCheckContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_restCheck

	return p
}

func (s *RestCheckContext) GetParser() antlr.Parser { return s.parser }

func (s *RestCheckContext) IntLit() antlr.TerminalNode {
	return s.GetToken(LDEParserIntLit, 0)
}

func (s *RestCheckContext) ComparisonOperator() antlr.TerminalNode {
	return s.GetToken(LDEParserComparisonOperator, 0)
}

func (s *RestCheckContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RestCheckContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RestCheckContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterRestCheck(s)
	}
}

func (s *RestCheckContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitRestCheck(s)
	}
}

func (p *LDEParser) RestCheck() (localctx IRestCheckContext) {
	localctx = NewRestCheckContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, LDEParserRULE_restCheck)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(205)
			p.Match(LDEParserT__13)
		}
		{
			p.SetState(206)
			p.Match(LDEParserIntLit)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(207)
			p.Match(LDEParserT__13)
		}
		{
			p.SetState(208)
			p.Match(LDEParserComparisonOperator)
		}
		{
			p.SetState(209)
			p.Match(LDEParserIntLit)
		}

	}

	return localctx
}

// IAtEndContext is an interface to support dynamic dispatch.
type IAtEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtEndContext differentiates from other interfaces.
	IsAtEndContext()
}

type AtEndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtEndContext() *AtEndContext {
	var p = new(AtEndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_atEnd
	return p
}

func (*AtEndContext) IsAtEndContext() {}

func NewAtEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtEndContext {
	var p = new(AtEndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_atEnd

	return p
}

func (s *AtEndContext) GetParser() antlr.Parser { return s.parser }
func (s *AtEndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtEndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtEndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterAtEnd(s)
	}
}

func (s *AtEndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitAtEnd(s)
	}
}

func (p *LDEParser) AtEnd() (localctx IAtEndContext) {
	localctx = NewAtEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, LDEParserRULE_atEnd)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(LDEParserT__14)
	}

	return localctx
}

// ITargetContext is an interface to support dynamic dispatch.
type ITargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTargetContext differentiates from other interfaces.
	IsTargetContext()
}

type TargetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetContext() *TargetContext {
	var p = new(TargetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_target
	return p
}

func (*TargetContext) IsTargetContext() {}

func NewTargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetContext {
	var p = new(TargetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_target

	return p
}

func (s *TargetContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetContext) TargetLit() ITargetLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetLitContext)
}

func (s *TargetContext) Bound() IBoundContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoundContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoundContext)
}

func (s *TargetContext) Limit() ILimitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimitContext)
}

func (s *TargetContext) Exact() IExactContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExactContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExactContext)
}

func (s *TargetContext) Jump() IJumpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJumpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJumpContext)
}

func (s *TargetContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *TargetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterTarget(s)
	}
}

func (s *TargetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitTarget(s)
	}
}

func (p *LDEParser) Target() (localctx ITargetContext) {
	localctx = NewTargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, LDEParserRULE_target)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(214)
			p.TargetLit()
		}
		{
			p.SetState(215)
			p.Bound()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.TargetLit()
		}
		{
			p.SetState(218)
			p.Limit()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(220)
			p.TargetLit()
		}
		{
			p.SetState(221)
			p.Exact()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(223)
			p.TargetLit()
		}
		{
			p.SetState(224)
			p.Jump()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(226)
			p.TargetLit()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(227)
			p.Match(LDEParserT__15)
		}
		{
			p.SetState(228)
			p.Target()
		}

	}

	return localctx
}

// ITargetLitContext is an interface to support dynamic dispatch.
type ITargetLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTargetLitContext differentiates from other interfaces.
	IsTargetLitContext()
}

type TargetLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetLitContext() *TargetLitContext {
	var p = new(TargetLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_targetLit
	return p
}

func (*TargetLitContext) IsTargetLitContext() {}

func NewTargetLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetLitContext {
	var p = new(TargetLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_targetLit

	return p
}

func (s *TargetLitContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetLitContext) CharLit() antlr.TerminalNode {
	return s.GetToken(LDEParserCharLit, 0)
}

func (s *TargetLitContext) StringLit() antlr.TerminalNode {
	return s.GetToken(LDEParserStringLit, 0)
}

func (s *TargetLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterTargetLit(s)
	}
}

func (s *TargetLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitTargetLit(s)
	}
}

func (p *LDEParser) TargetLit() (localctx ITargetLitContext) {
	localctx = NewTargetLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, LDEParserRULE_targetLit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LDEParserStringLit || _la == LDEParserCharLit) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBoundContext is an interface to support dynamic dispatch.
type IBoundContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoundContext differentiates from other interfaces.
	IsBoundContext()
}

type BoundContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoundContext() *BoundContext {
	var p = new(BoundContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_bound
	return p
}

func (*BoundContext) IsBoundContext() {}

func NewBoundContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoundContext {
	var p = new(BoundContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_bound

	return p
}

func (s *BoundContext) GetParser() antlr.Parser { return s.parser }

func (s *BoundContext) AllIntLit() []antlr.TerminalNode {
	return s.GetTokens(LDEParserIntLit)
}

func (s *BoundContext) IntLit(i int) antlr.TerminalNode {
	return s.GetToken(LDEParserIntLit, i)
}

func (s *BoundContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoundContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoundContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterBound(s)
	}
}

func (s *BoundContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitBound(s)
	}
}

func (p *LDEParser) Bound() (localctx IBoundContext) {
	localctx = NewBoundContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, LDEParserRULE_bound)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		p.Match(LDEParserT__6)
	}
	{
		p.SetState(234)
		p.Match(LDEParserIntLit)
	}
	{
		p.SetState(235)
		p.Match(LDEParserT__11)
	}
	{
		p.SetState(236)
		p.Match(LDEParserIntLit)
	}
	{
		p.SetState(237)
		p.Match(LDEParserT__7)
	}

	return localctx
}

// ILimitContext is an interface to support dynamic dispatch.
type ILimitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLimitContext differentiates from other interfaces.
	IsLimitContext()
}

type LimitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitContext() *LimitContext {
	var p = new(LimitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_limit
	return p
}

func (*LimitContext) IsLimitContext() {}

func NewLimitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitContext {
	var p = new(LimitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_limit

	return p
}

func (s *LimitContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitContext) IntLit() antlr.TerminalNode {
	return s.GetToken(LDEParserIntLit, 0)
}

func (s *LimitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterLimit(s)
	}
}

func (s *LimitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitLimit(s)
	}
}

func (p *LDEParser) Limit() (localctx ILimitContext) {
	localctx = NewLimitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, LDEParserRULE_limit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.Match(LDEParserT__6)
	}
	{
		p.SetState(240)
		p.Match(LDEParserT__11)
	}
	{
		p.SetState(241)
		p.Match(LDEParserIntLit)
	}
	{
		p.SetState(242)
		p.Match(LDEParserT__7)
	}

	return localctx
}

// IJumpContext is an interface to support dynamic dispatch.
type IJumpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJumpContext differentiates from other interfaces.
	IsJumpContext()
}

type JumpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJumpContext() *JumpContext {
	var p = new(JumpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_jump
	return p
}

func (*JumpContext) IsJumpContext() {}

func NewJumpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JumpContext {
	var p = new(JumpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_jump

	return p
}

func (s *JumpContext) GetParser() antlr.Parser { return s.parser }

func (s *JumpContext) IntLit() antlr.TerminalNode {
	return s.GetToken(LDEParserIntLit, 0)
}

func (s *JumpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JumpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JumpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterJump(s)
	}
}

func (s *JumpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitJump(s)
	}
}

func (p *LDEParser) Jump() (localctx IJumpContext) {
	localctx = NewJumpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, LDEParserRULE_jump)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(244)
		p.Match(LDEParserT__6)
	}
	{
		p.SetState(245)
		p.Match(LDEParserIntLit)
	}
	{
		p.SetState(246)
		p.Match(LDEParserT__11)
	}
	{
		p.SetState(247)
		p.Match(LDEParserT__7)
	}

	return localctx
}

// IExactContext is an interface to support dynamic dispatch.
type IExactContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExactContext differentiates from other interfaces.
	IsExactContext()
}

type ExactContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExactContext() *ExactContext {
	var p = new(ExactContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_exact
	return p
}

func (*ExactContext) IsExactContext() {}

func NewExactContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExactContext {
	var p = new(ExactContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_exact

	return p
}

func (s *ExactContext) GetParser() antlr.Parser { return s.parser }

func (s *ExactContext) IntLit() antlr.TerminalNode {
	return s.GetToken(LDEParserIntLit, 0)
}

func (s *ExactContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExactContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExactContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterExact(s)
	}
}

func (s *ExactContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitExact(s)
	}
}

func (p *LDEParser) Exact() (localctx IExactContext) {
	localctx = NewExactContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, LDEParserRULE_exact)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.Match(LDEParserT__6)
	}
	{
		p.SetState(250)
		p.Match(LDEParserIntLit)
	}
	{
		p.SetState(251)
		p.Match(LDEParserT__7)
	}

	return localctx
}

// IFieldTypeContext is an interface to support dynamic dispatch.
type IFieldTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldTypeContext differentiates from other interfaces.
	IsFieldTypeContext()
}

type FieldTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldTypeContext() *FieldTypeContext {
	var p = new(FieldTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_fieldType
	return p
}

func (*FieldTypeContext) IsFieldTypeContext() {}

func NewFieldTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldTypeContext {
	var p = new(FieldTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_fieldType

	return p
}

func (s *FieldTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldTypeContext) IdentifierWithFraction() antlr.TerminalNode {
	return s.GetToken(LDEParserIdentifierWithFraction, 0)
}

func (s *FieldTypeContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LDEParserIdentifier, 0)
}

func (s *FieldTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterFieldType(s)
	}
}

func (s *FieldTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitFieldType(s)
	}
}

func (p *LDEParser) FieldType() (localctx IFieldTypeContext) {
	localctx = NewFieldTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, LDEParserRULE_fieldType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LDEParserIdentifier || _la == LDEParserIdentifierWithFraction) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
