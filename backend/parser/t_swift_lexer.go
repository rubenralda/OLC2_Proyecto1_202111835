// Code generated from .\parser\T_swift.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type T_swiftLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var T_swiftLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func t_swiftlexerLexerInit() {
	staticData := &T_swiftLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "';'", "'func'", "'('", "')'", "'->'", "','", "':'", "'inout'",
		"'{'", "'}'", "'struct'", "'let'", "'var'", "'='", "'mutating'", "'&'",
		"'.'", "'IsEmpty'", "'count'", "'+='", "'-='", "'for'", "'in'", "'...'",
		"'while'", "'if'", "'else'", "'guard'", "'continue'", "'break'", "'return'",
		"'switch'", "'case'", "'default'", "'?'", "'String'", "'Int'", "'Float'",
		"'Bool'", "'Character'", "'['", "']'", "'print'", "'append'", "'removeLast'",
		"'remove'", "'at'", "'!'", "'-'", "'/'", "'%'", "'*'", "'+'", "'<'",
		"'<='", "'>='", "'>'", "'=='", "'!='", "'&&'", "'||'", "'nil'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "WS", "Block_comment",
		"Line_comment", "Int", "Float", "Caracter", "String", "Bool", "Identificador",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
		"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40",
		"T__41", "T__42", "T__43", "T__44", "T__45", "T__46", "T__47", "T__48",
		"T__49", "T__50", "T__51", "T__52", "T__53", "T__54", "T__55", "T__56",
		"T__57", "T__58", "T__59", "T__60", "T__61", "WS", "Block_comment",
		"Line_comment", "Int", "Float", "Caracter", "String", "Quoted_text",
		"Quoted_text_item", "Bool", "Identificador",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 71, 500, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38,
		1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42,
		1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1,
		44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45,
		1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 47, 1,
		47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 52, 1, 52,
		1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 55, 1, 55, 1, 55, 1, 56, 1, 56, 1,
		57, 1, 57, 1, 57, 1, 58, 1, 58, 1, 58, 1, 59, 1, 59, 1, 59, 1, 60, 1, 60,
		1, 60, 1, 61, 1, 61, 1, 61, 1, 61, 1, 62, 4, 62, 410, 8, 62, 11, 62, 12,
		62, 411, 1, 62, 1, 62, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 5, 63, 421, 8,
		63, 10, 63, 12, 63, 424, 9, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 64,
		1, 64, 1, 64, 1, 64, 5, 64, 435, 8, 64, 10, 64, 12, 64, 438, 9, 64, 1,
		64, 3, 64, 441, 8, 64, 1, 64, 1, 64, 1, 65, 4, 65, 446, 8, 65, 11, 65,
		12, 65, 447, 1, 66, 4, 66, 451, 8, 66, 11, 66, 12, 66, 452, 1, 66, 1, 66,
		4, 66, 457, 8, 66, 11, 66, 12, 66, 458, 3, 66, 461, 8, 66, 1, 67, 1, 67,
		1, 67, 1, 67, 1, 68, 1, 68, 3, 68, 469, 8, 68, 1, 68, 1, 68, 1, 69, 4,
		69, 474, 8, 69, 11, 69, 12, 69, 475, 1, 70, 1, 70, 1, 70, 3, 70, 481, 8,
		70, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 3, 71,
		492, 8, 71, 1, 72, 1, 72, 5, 72, 496, 8, 72, 10, 72, 12, 72, 499, 9, 72,
		2, 422, 436, 0, 73, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8,
		17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35,
		71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44,
		89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105,
		53, 107, 54, 109, 55, 111, 56, 113, 57, 115, 58, 117, 59, 119, 60, 121,
		61, 123, 62, 125, 63, 127, 64, 129, 65, 131, 66, 133, 67, 135, 68, 137,
		69, 139, 0, 141, 0, 143, 70, 145, 71, 1, 0, 7, 3, 0, 0, 0, 9, 13, 32, 32,
		1, 1, 10, 10, 1, 0, 48, 57, 7, 0, 34, 34, 39, 39, 48, 48, 92, 92, 110,
		110, 114, 114, 116, 116, 4, 0, 10, 10, 13, 13, 34, 34, 92, 92, 3, 0, 65,
		90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 510, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0,
		0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0,
		0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0,
		0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1,
		0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55,
		1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0,
		63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0,
		0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0,
		0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0,
		0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1,
		0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101,
		1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0,
		0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1,
		0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121, 1, 0, 0, 0, 0,
		123, 1, 0, 0, 0, 0, 125, 1, 0, 0, 0, 0, 127, 1, 0, 0, 0, 0, 129, 1, 0,
		0, 0, 0, 131, 1, 0, 0, 0, 0, 133, 1, 0, 0, 0, 0, 135, 1, 0, 0, 0, 0, 137,
		1, 0, 0, 0, 0, 143, 1, 0, 0, 0, 0, 145, 1, 0, 0, 0, 1, 147, 1, 0, 0, 0,
		3, 149, 1, 0, 0, 0, 5, 154, 1, 0, 0, 0, 7, 156, 1, 0, 0, 0, 9, 158, 1,
		0, 0, 0, 11, 161, 1, 0, 0, 0, 13, 163, 1, 0, 0, 0, 15, 165, 1, 0, 0, 0,
		17, 171, 1, 0, 0, 0, 19, 173, 1, 0, 0, 0, 21, 175, 1, 0, 0, 0, 23, 182,
		1, 0, 0, 0, 25, 186, 1, 0, 0, 0, 27, 190, 1, 0, 0, 0, 29, 192, 1, 0, 0,
		0, 31, 201, 1, 0, 0, 0, 33, 203, 1, 0, 0, 0, 35, 205, 1, 0, 0, 0, 37, 213,
		1, 0, 0, 0, 39, 219, 1, 0, 0, 0, 41, 222, 1, 0, 0, 0, 43, 225, 1, 0, 0,
		0, 45, 229, 1, 0, 0, 0, 47, 232, 1, 0, 0, 0, 49, 236, 1, 0, 0, 0, 51, 242,
		1, 0, 0, 0, 53, 245, 1, 0, 0, 0, 55, 250, 1, 0, 0, 0, 57, 256, 1, 0, 0,
		0, 59, 265, 1, 0, 0, 0, 61, 271, 1, 0, 0, 0, 63, 278, 1, 0, 0, 0, 65, 285,
		1, 0, 0, 0, 67, 290, 1, 0, 0, 0, 69, 298, 1, 0, 0, 0, 71, 300, 1, 0, 0,
		0, 73, 307, 1, 0, 0, 0, 75, 311, 1, 0, 0, 0, 77, 317, 1, 0, 0, 0, 79, 322,
		1, 0, 0, 0, 81, 332, 1, 0, 0, 0, 83, 334, 1, 0, 0, 0, 85, 336, 1, 0, 0,
		0, 87, 342, 1, 0, 0, 0, 89, 349, 1, 0, 0, 0, 91, 360, 1, 0, 0, 0, 93, 367,
		1, 0, 0, 0, 95, 370, 1, 0, 0, 0, 97, 372, 1, 0, 0, 0, 99, 374, 1, 0, 0,
		0, 101, 376, 1, 0, 0, 0, 103, 378, 1, 0, 0, 0, 105, 380, 1, 0, 0, 0, 107,
		382, 1, 0, 0, 0, 109, 384, 1, 0, 0, 0, 111, 387, 1, 0, 0, 0, 113, 390,
		1, 0, 0, 0, 115, 392, 1, 0, 0, 0, 117, 395, 1, 0, 0, 0, 119, 398, 1, 0,
		0, 0, 121, 401, 1, 0, 0, 0, 123, 404, 1, 0, 0, 0, 125, 409, 1, 0, 0, 0,
		127, 415, 1, 0, 0, 0, 129, 430, 1, 0, 0, 0, 131, 445, 1, 0, 0, 0, 133,
		450, 1, 0, 0, 0, 135, 462, 1, 0, 0, 0, 137, 466, 1, 0, 0, 0, 139, 473,
		1, 0, 0, 0, 141, 480, 1, 0, 0, 0, 143, 491, 1, 0, 0, 0, 145, 493, 1, 0,
		0, 0, 147, 148, 5, 59, 0, 0, 148, 2, 1, 0, 0, 0, 149, 150, 5, 102, 0, 0,
		150, 151, 5, 117, 0, 0, 151, 152, 5, 110, 0, 0, 152, 153, 5, 99, 0, 0,
		153, 4, 1, 0, 0, 0, 154, 155, 5, 40, 0, 0, 155, 6, 1, 0, 0, 0, 156, 157,
		5, 41, 0, 0, 157, 8, 1, 0, 0, 0, 158, 159, 5, 45, 0, 0, 159, 160, 5, 62,
		0, 0, 160, 10, 1, 0, 0, 0, 161, 162, 5, 44, 0, 0, 162, 12, 1, 0, 0, 0,
		163, 164, 5, 58, 0, 0, 164, 14, 1, 0, 0, 0, 165, 166, 5, 105, 0, 0, 166,
		167, 5, 110, 0, 0, 167, 168, 5, 111, 0, 0, 168, 169, 5, 117, 0, 0, 169,
		170, 5, 116, 0, 0, 170, 16, 1, 0, 0, 0, 171, 172, 5, 123, 0, 0, 172, 18,
		1, 0, 0, 0, 173, 174, 5, 125, 0, 0, 174, 20, 1, 0, 0, 0, 175, 176, 5, 115,
		0, 0, 176, 177, 5, 116, 0, 0, 177, 178, 5, 114, 0, 0, 178, 179, 5, 117,
		0, 0, 179, 180, 5, 99, 0, 0, 180, 181, 5, 116, 0, 0, 181, 22, 1, 0, 0,
		0, 182, 183, 5, 108, 0, 0, 183, 184, 5, 101, 0, 0, 184, 185, 5, 116, 0,
		0, 185, 24, 1, 0, 0, 0, 186, 187, 5, 118, 0, 0, 187, 188, 5, 97, 0, 0,
		188, 189, 5, 114, 0, 0, 189, 26, 1, 0, 0, 0, 190, 191, 5, 61, 0, 0, 191,
		28, 1, 0, 0, 0, 192, 193, 5, 109, 0, 0, 193, 194, 5, 117, 0, 0, 194, 195,
		5, 116, 0, 0, 195, 196, 5, 97, 0, 0, 196, 197, 5, 116, 0, 0, 197, 198,
		5, 105, 0, 0, 198, 199, 5, 110, 0, 0, 199, 200, 5, 103, 0, 0, 200, 30,
		1, 0, 0, 0, 201, 202, 5, 38, 0, 0, 202, 32, 1, 0, 0, 0, 203, 204, 5, 46,
		0, 0, 204, 34, 1, 0, 0, 0, 205, 206, 5, 73, 0, 0, 206, 207, 5, 115, 0,
		0, 207, 208, 5, 69, 0, 0, 208, 209, 5, 109, 0, 0, 209, 210, 5, 112, 0,
		0, 210, 211, 5, 116, 0, 0, 211, 212, 5, 121, 0, 0, 212, 36, 1, 0, 0, 0,
		213, 214, 5, 99, 0, 0, 214, 215, 5, 111, 0, 0, 215, 216, 5, 117, 0, 0,
		216, 217, 5, 110, 0, 0, 217, 218, 5, 116, 0, 0, 218, 38, 1, 0, 0, 0, 219,
		220, 5, 43, 0, 0, 220, 221, 5, 61, 0, 0, 221, 40, 1, 0, 0, 0, 222, 223,
		5, 45, 0, 0, 223, 224, 5, 61, 0, 0, 224, 42, 1, 0, 0, 0, 225, 226, 5, 102,
		0, 0, 226, 227, 5, 111, 0, 0, 227, 228, 5, 114, 0, 0, 228, 44, 1, 0, 0,
		0, 229, 230, 5, 105, 0, 0, 230, 231, 5, 110, 0, 0, 231, 46, 1, 0, 0, 0,
		232, 233, 5, 46, 0, 0, 233, 234, 5, 46, 0, 0, 234, 235, 5, 46, 0, 0, 235,
		48, 1, 0, 0, 0, 236, 237, 5, 119, 0, 0, 237, 238, 5, 104, 0, 0, 238, 239,
		5, 105, 0, 0, 239, 240, 5, 108, 0, 0, 240, 241, 5, 101, 0, 0, 241, 50,
		1, 0, 0, 0, 242, 243, 5, 105, 0, 0, 243, 244, 5, 102, 0, 0, 244, 52, 1,
		0, 0, 0, 245, 246, 5, 101, 0, 0, 246, 247, 5, 108, 0, 0, 247, 248, 5, 115,
		0, 0, 248, 249, 5, 101, 0, 0, 249, 54, 1, 0, 0, 0, 250, 251, 5, 103, 0,
		0, 251, 252, 5, 117, 0, 0, 252, 253, 5, 97, 0, 0, 253, 254, 5, 114, 0,
		0, 254, 255, 5, 100, 0, 0, 255, 56, 1, 0, 0, 0, 256, 257, 5, 99, 0, 0,
		257, 258, 5, 111, 0, 0, 258, 259, 5, 110, 0, 0, 259, 260, 5, 116, 0, 0,
		260, 261, 5, 105, 0, 0, 261, 262, 5, 110, 0, 0, 262, 263, 5, 117, 0, 0,
		263, 264, 5, 101, 0, 0, 264, 58, 1, 0, 0, 0, 265, 266, 5, 98, 0, 0, 266,
		267, 5, 114, 0, 0, 267, 268, 5, 101, 0, 0, 268, 269, 5, 97, 0, 0, 269,
		270, 5, 107, 0, 0, 270, 60, 1, 0, 0, 0, 271, 272, 5, 114, 0, 0, 272, 273,
		5, 101, 0, 0, 273, 274, 5, 116, 0, 0, 274, 275, 5, 117, 0, 0, 275, 276,
		5, 114, 0, 0, 276, 277, 5, 110, 0, 0, 277, 62, 1, 0, 0, 0, 278, 279, 5,
		115, 0, 0, 279, 280, 5, 119, 0, 0, 280, 281, 5, 105, 0, 0, 281, 282, 5,
		116, 0, 0, 282, 283, 5, 99, 0, 0, 283, 284, 5, 104, 0, 0, 284, 64, 1, 0,
		0, 0, 285, 286, 5, 99, 0, 0, 286, 287, 5, 97, 0, 0, 287, 288, 5, 115, 0,
		0, 288, 289, 5, 101, 0, 0, 289, 66, 1, 0, 0, 0, 290, 291, 5, 100, 0, 0,
		291, 292, 5, 101, 0, 0, 292, 293, 5, 102, 0, 0, 293, 294, 5, 97, 0, 0,
		294, 295, 5, 117, 0, 0, 295, 296, 5, 108, 0, 0, 296, 297, 5, 116, 0, 0,
		297, 68, 1, 0, 0, 0, 298, 299, 5, 63, 0, 0, 299, 70, 1, 0, 0, 0, 300, 301,
		5, 83, 0, 0, 301, 302, 5, 116, 0, 0, 302, 303, 5, 114, 0, 0, 303, 304,
		5, 105, 0, 0, 304, 305, 5, 110, 0, 0, 305, 306, 5, 103, 0, 0, 306, 72,
		1, 0, 0, 0, 307, 308, 5, 73, 0, 0, 308, 309, 5, 110, 0, 0, 309, 310, 5,
		116, 0, 0, 310, 74, 1, 0, 0, 0, 311, 312, 5, 70, 0, 0, 312, 313, 5, 108,
		0, 0, 313, 314, 5, 111, 0, 0, 314, 315, 5, 97, 0, 0, 315, 316, 5, 116,
		0, 0, 316, 76, 1, 0, 0, 0, 317, 318, 5, 66, 0, 0, 318, 319, 5, 111, 0,
		0, 319, 320, 5, 111, 0, 0, 320, 321, 5, 108, 0, 0, 321, 78, 1, 0, 0, 0,
		322, 323, 5, 67, 0, 0, 323, 324, 5, 104, 0, 0, 324, 325, 5, 97, 0, 0, 325,
		326, 5, 114, 0, 0, 326, 327, 5, 97, 0, 0, 327, 328, 5, 99, 0, 0, 328, 329,
		5, 116, 0, 0, 329, 330, 5, 101, 0, 0, 330, 331, 5, 114, 0, 0, 331, 80,
		1, 0, 0, 0, 332, 333, 5, 91, 0, 0, 333, 82, 1, 0, 0, 0, 334, 335, 5, 93,
		0, 0, 335, 84, 1, 0, 0, 0, 336, 337, 5, 112, 0, 0, 337, 338, 5, 114, 0,
		0, 338, 339, 5, 105, 0, 0, 339, 340, 5, 110, 0, 0, 340, 341, 5, 116, 0,
		0, 341, 86, 1, 0, 0, 0, 342, 343, 5, 97, 0, 0, 343, 344, 5, 112, 0, 0,
		344, 345, 5, 112, 0, 0, 345, 346, 5, 101, 0, 0, 346, 347, 5, 110, 0, 0,
		347, 348, 5, 100, 0, 0, 348, 88, 1, 0, 0, 0, 349, 350, 5, 114, 0, 0, 350,
		351, 5, 101, 0, 0, 351, 352, 5, 109, 0, 0, 352, 353, 5, 111, 0, 0, 353,
		354, 5, 118, 0, 0, 354, 355, 5, 101, 0, 0, 355, 356, 5, 76, 0, 0, 356,
		357, 5, 97, 0, 0, 357, 358, 5, 115, 0, 0, 358, 359, 5, 116, 0, 0, 359,
		90, 1, 0, 0, 0, 360, 361, 5, 114, 0, 0, 361, 362, 5, 101, 0, 0, 362, 363,
		5, 109, 0, 0, 363, 364, 5, 111, 0, 0, 364, 365, 5, 118, 0, 0, 365, 366,
		5, 101, 0, 0, 366, 92, 1, 0, 0, 0, 367, 368, 5, 97, 0, 0, 368, 369, 5,
		116, 0, 0, 369, 94, 1, 0, 0, 0, 370, 371, 5, 33, 0, 0, 371, 96, 1, 0, 0,
		0, 372, 373, 5, 45, 0, 0, 373, 98, 1, 0, 0, 0, 374, 375, 5, 47, 0, 0, 375,
		100, 1, 0, 0, 0, 376, 377, 5, 37, 0, 0, 377, 102, 1, 0, 0, 0, 378, 379,
		5, 42, 0, 0, 379, 104, 1, 0, 0, 0, 380, 381, 5, 43, 0, 0, 381, 106, 1,
		0, 0, 0, 382, 383, 5, 60, 0, 0, 383, 108, 1, 0, 0, 0, 384, 385, 5, 60,
		0, 0, 385, 386, 5, 61, 0, 0, 386, 110, 1, 0, 0, 0, 387, 388, 5, 62, 0,
		0, 388, 389, 5, 61, 0, 0, 389, 112, 1, 0, 0, 0, 390, 391, 5, 62, 0, 0,
		391, 114, 1, 0, 0, 0, 392, 393, 5, 61, 0, 0, 393, 394, 5, 61, 0, 0, 394,
		116, 1, 0, 0, 0, 395, 396, 5, 33, 0, 0, 396, 397, 5, 61, 0, 0, 397, 118,
		1, 0, 0, 0, 398, 399, 5, 38, 0, 0, 399, 400, 5, 38, 0, 0, 400, 120, 1,
		0, 0, 0, 401, 402, 5, 124, 0, 0, 402, 403, 5, 124, 0, 0, 403, 122, 1, 0,
		0, 0, 404, 405, 5, 110, 0, 0, 405, 406, 5, 105, 0, 0, 406, 407, 5, 108,
		0, 0, 407, 124, 1, 0, 0, 0, 408, 410, 7, 0, 0, 0, 409, 408, 1, 0, 0, 0,
		410, 411, 1, 0, 0, 0, 411, 409, 1, 0, 0, 0, 411, 412, 1, 0, 0, 0, 412,
		413, 1, 0, 0, 0, 413, 414, 6, 62, 0, 0, 414, 126, 1, 0, 0, 0, 415, 416,
		5, 47, 0, 0, 416, 417, 5, 42, 0, 0, 417, 422, 1, 0, 0, 0, 418, 421, 3,
		127, 63, 0, 419, 421, 9, 0, 0, 0, 420, 418, 1, 0, 0, 0, 420, 419, 1, 0,
		0, 0, 421, 424, 1, 0, 0, 0, 422, 423, 1, 0, 0, 0, 422, 420, 1, 0, 0, 0,
		423, 425, 1, 0, 0, 0, 424, 422, 1, 0, 0, 0, 425, 426, 5, 42, 0, 0, 426,
		427, 5, 47, 0, 0, 427, 428, 1, 0, 0, 0, 428, 429, 6, 63, 0, 0, 429, 128,
		1, 0, 0, 0, 430, 431, 5, 47, 0, 0, 431, 432, 5, 47, 0, 0, 432, 436, 1,
		0, 0, 0, 433, 435, 9, 0, 0, 0, 434, 433, 1, 0, 0, 0, 435, 438, 1, 0, 0,
		0, 436, 437, 1, 0, 0, 0, 436, 434, 1, 0, 0, 0, 437, 440, 1, 0, 0, 0, 438,
		436, 1, 0, 0, 0, 439, 441, 7, 1, 0, 0, 440, 439, 1, 0, 0, 0, 441, 442,
		1, 0, 0, 0, 442, 443, 6, 64, 0, 0, 443, 130, 1, 0, 0, 0, 444, 446, 7, 2,
		0, 0, 445, 444, 1, 0, 0, 0, 446, 447, 1, 0, 0, 0, 447, 445, 1, 0, 0, 0,
		447, 448, 1, 0, 0, 0, 448, 132, 1, 0, 0, 0, 449, 451, 7, 2, 0, 0, 450,
		449, 1, 0, 0, 0, 451, 452, 1, 0, 0, 0, 452, 450, 1, 0, 0, 0, 452, 453,
		1, 0, 0, 0, 453, 460, 1, 0, 0, 0, 454, 456, 5, 46, 0, 0, 455, 457, 7, 2,
		0, 0, 456, 455, 1, 0, 0, 0, 457, 458, 1, 0, 0, 0, 458, 456, 1, 0, 0, 0,
		458, 459, 1, 0, 0, 0, 459, 461, 1, 0, 0, 0, 460, 454, 1, 0, 0, 0, 460,
		461, 1, 0, 0, 0, 461, 134, 1, 0, 0, 0, 462, 463, 5, 39, 0, 0, 463, 464,
		3, 141, 70, 0, 464, 465, 5, 39, 0, 0, 465, 136, 1, 0, 0, 0, 466, 468, 5,
		34, 0, 0, 467, 469, 3, 139, 69, 0, 468, 467, 1, 0, 0, 0, 468, 469, 1, 0,
		0, 0, 469, 470, 1, 0, 0, 0, 470, 471, 5, 34, 0, 0, 471, 138, 1, 0, 0, 0,
		472, 474, 3, 141, 70, 0, 473, 472, 1, 0, 0, 0, 474, 475, 1, 0, 0, 0, 475,
		473, 1, 0, 0, 0, 475, 476, 1, 0, 0, 0, 476, 140, 1, 0, 0, 0, 477, 478,
		5, 92, 0, 0, 478, 481, 7, 3, 0, 0, 479, 481, 8, 4, 0, 0, 480, 477, 1, 0,
		0, 0, 480, 479, 1, 0, 0, 0, 481, 142, 1, 0, 0, 0, 482, 483, 5, 116, 0,
		0, 483, 484, 5, 114, 0, 0, 484, 485, 5, 117, 0, 0, 485, 492, 5, 101, 0,
		0, 486, 487, 5, 102, 0, 0, 487, 488, 5, 97, 0, 0, 488, 489, 5, 108, 0,
		0, 489, 490, 5, 115, 0, 0, 490, 492, 5, 101, 0, 0, 491, 482, 1, 0, 0, 0,
		491, 486, 1, 0, 0, 0, 492, 144, 1, 0, 0, 0, 493, 497, 7, 5, 0, 0, 494,
		496, 7, 6, 0, 0, 495, 494, 1, 0, 0, 0, 496, 499, 1, 0, 0, 0, 497, 495,
		1, 0, 0, 0, 497, 498, 1, 0, 0, 0, 498, 146, 1, 0, 0, 0, 499, 497, 1, 0,
		0, 0, 15, 0, 411, 420, 422, 436, 440, 447, 452, 458, 460, 468, 475, 480,
		491, 497, 1, 0, 1, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// T_swiftLexerInit initializes any static state used to implement T_swiftLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewT_swiftLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func T_swiftLexerInit() {
	staticData := &T_swiftLexerLexerStaticData
	staticData.once.Do(t_swiftlexerLexerInit)
}

// NewT_swiftLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewT_swiftLexer(input antlr.CharStream) *T_swiftLexer {
	T_swiftLexerInit()
	l := new(T_swiftLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &T_swiftLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "T_swift.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// T_swiftLexer tokens.
const (
	T_swiftLexerT__0          = 1
	T_swiftLexerT__1          = 2
	T_swiftLexerT__2          = 3
	T_swiftLexerT__3          = 4
	T_swiftLexerT__4          = 5
	T_swiftLexerT__5          = 6
	T_swiftLexerT__6          = 7
	T_swiftLexerT__7          = 8
	T_swiftLexerT__8          = 9
	T_swiftLexerT__9          = 10
	T_swiftLexerT__10         = 11
	T_swiftLexerT__11         = 12
	T_swiftLexerT__12         = 13
	T_swiftLexerT__13         = 14
	T_swiftLexerT__14         = 15
	T_swiftLexerT__15         = 16
	T_swiftLexerT__16         = 17
	T_swiftLexerT__17         = 18
	T_swiftLexerT__18         = 19
	T_swiftLexerT__19         = 20
	T_swiftLexerT__20         = 21
	T_swiftLexerT__21         = 22
	T_swiftLexerT__22         = 23
	T_swiftLexerT__23         = 24
	T_swiftLexerT__24         = 25
	T_swiftLexerT__25         = 26
	T_swiftLexerT__26         = 27
	T_swiftLexerT__27         = 28
	T_swiftLexerT__28         = 29
	T_swiftLexerT__29         = 30
	T_swiftLexerT__30         = 31
	T_swiftLexerT__31         = 32
	T_swiftLexerT__32         = 33
	T_swiftLexerT__33         = 34
	T_swiftLexerT__34         = 35
	T_swiftLexerT__35         = 36
	T_swiftLexerT__36         = 37
	T_swiftLexerT__37         = 38
	T_swiftLexerT__38         = 39
	T_swiftLexerT__39         = 40
	T_swiftLexerT__40         = 41
	T_swiftLexerT__41         = 42
	T_swiftLexerT__42         = 43
	T_swiftLexerT__43         = 44
	T_swiftLexerT__44         = 45
	T_swiftLexerT__45         = 46
	T_swiftLexerT__46         = 47
	T_swiftLexerT__47         = 48
	T_swiftLexerT__48         = 49
	T_swiftLexerT__49         = 50
	T_swiftLexerT__50         = 51
	T_swiftLexerT__51         = 52
	T_swiftLexerT__52         = 53
	T_swiftLexerT__53         = 54
	T_swiftLexerT__54         = 55
	T_swiftLexerT__55         = 56
	T_swiftLexerT__56         = 57
	T_swiftLexerT__57         = 58
	T_swiftLexerT__58         = 59
	T_swiftLexerT__59         = 60
	T_swiftLexerT__60         = 61
	T_swiftLexerT__61         = 62
	T_swiftLexerWS            = 63
	T_swiftLexerBlock_comment = 64
	T_swiftLexerLine_comment  = 65
	T_swiftLexerInt           = 66
	T_swiftLexerFloat         = 67
	T_swiftLexerCaracter      = 68
	T_swiftLexerString_       = 69
	T_swiftLexerBool          = 70
	T_swiftLexerIdentificador = 71
)
