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
		"", "';'", "'{'", "'}'", "'for'", "'in'", "'...'", "'while'", "'if'",
		"'else'", "'guard'", "'continue'", "'break'", "'return'", "'switch'",
		"'case'", "':'", "'default'", "'let'", "'='", "'var'", "'?'", "'String'",
		"'Int'", "'Float'", "'Bool'", "'Character'", "'['", "']'", "','", "'print'",
		"'('", "')'", "'+='", "'-='", "'!'", "'-'", "'/'", "'%'", "'*'", "'+'",
		"'<'", "'<='", "'>='", "'>'", "'=='", "'!='", "'&&'", "'||'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "WS", "Block_comment",
		"Line_comment", "Int", "Float", "Caracter", "String", "Bool", "Identificador",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
		"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40",
		"T__41", "T__42", "T__43", "T__44", "T__45", "T__46", "T__47", "WS",
		"Block_comment", "Line_comment", "Int", "Float", "Caracter", "String",
		"Quoted_text", "Quoted_text_item", "Bool", "Identificador", "Identifier_head",
		"Identifier_character", "Identifier_characters",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 57, 407, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
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
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1,
		32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36,
		1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1,
		41, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45,
		1, 45, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 48, 4, 48, 308, 8,
		48, 11, 48, 12, 48, 309, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49,
		5, 49, 319, 8, 49, 10, 49, 12, 49, 322, 9, 49, 1, 49, 1, 49, 1, 49, 1,
		49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 5, 50, 333, 8, 50, 10, 50, 12, 50,
		336, 9, 50, 1, 50, 3, 50, 339, 8, 50, 1, 50, 1, 50, 1, 51, 4, 51, 344,
		8, 51, 11, 51, 12, 51, 345, 1, 52, 4, 52, 349, 8, 52, 11, 52, 12, 52, 350,
		1, 52, 1, 52, 4, 52, 355, 8, 52, 11, 52, 12, 52, 356, 3, 52, 359, 8, 52,
		1, 53, 1, 53, 1, 53, 1, 53, 1, 54, 1, 54, 3, 54, 367, 8, 54, 1, 54, 1,
		54, 1, 55, 4, 55, 372, 8, 55, 11, 55, 12, 55, 373, 1, 56, 1, 56, 1, 56,
		3, 56, 379, 8, 56, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1,
		57, 1, 57, 3, 57, 390, 8, 57, 1, 58, 1, 58, 3, 58, 394, 8, 58, 1, 59, 3,
		59, 397, 8, 59, 1, 60, 1, 60, 3, 60, 401, 8, 60, 1, 61, 4, 61, 404, 8,
		61, 11, 61, 12, 61, 405, 2, 320, 334, 0, 62, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32,
		65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41,
		83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 50,
		101, 51, 103, 52, 105, 53, 107, 54, 109, 55, 111, 0, 113, 0, 115, 56, 117,
		57, 119, 0, 121, 0, 123, 0, 1, 0, 6, 3, 0, 0, 0, 9, 13, 32, 32, 1, 1, 10,
		10, 1, 0, 48, 57, 7, 0, 34, 34, 39, 39, 48, 48, 92, 92, 110, 110, 114,
		114, 116, 116, 4, 0, 10, 10, 13, 13, 34, 34, 92, 92, 3, 0, 65, 90, 95,
		95, 97, 122, 416, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0,
		0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0,
		0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0,
		0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1,
		0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37,
		1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0,
		45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0,
		0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0,
		0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0,
		0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1,
		0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83,
		1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0,
		91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0,
		0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0,
		0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 0, 117,
		1, 0, 0, 0, 1, 125, 1, 0, 0, 0, 3, 127, 1, 0, 0, 0, 5, 129, 1, 0, 0, 0,
		7, 131, 1, 0, 0, 0, 9, 135, 1, 0, 0, 0, 11, 138, 1, 0, 0, 0, 13, 142, 1,
		0, 0, 0, 15, 148, 1, 0, 0, 0, 17, 151, 1, 0, 0, 0, 19, 156, 1, 0, 0, 0,
		21, 162, 1, 0, 0, 0, 23, 171, 1, 0, 0, 0, 25, 177, 1, 0, 0, 0, 27, 184,
		1, 0, 0, 0, 29, 191, 1, 0, 0, 0, 31, 196, 1, 0, 0, 0, 33, 198, 1, 0, 0,
		0, 35, 206, 1, 0, 0, 0, 37, 210, 1, 0, 0, 0, 39, 212, 1, 0, 0, 0, 41, 216,
		1, 0, 0, 0, 43, 218, 1, 0, 0, 0, 45, 225, 1, 0, 0, 0, 47, 229, 1, 0, 0,
		0, 49, 235, 1, 0, 0, 0, 51, 240, 1, 0, 0, 0, 53, 250, 1, 0, 0, 0, 55, 252,
		1, 0, 0, 0, 57, 254, 1, 0, 0, 0, 59, 256, 1, 0, 0, 0, 61, 262, 1, 0, 0,
		0, 63, 264, 1, 0, 0, 0, 65, 266, 1, 0, 0, 0, 67, 269, 1, 0, 0, 0, 69, 272,
		1, 0, 0, 0, 71, 274, 1, 0, 0, 0, 73, 276, 1, 0, 0, 0, 75, 278, 1, 0, 0,
		0, 77, 280, 1, 0, 0, 0, 79, 282, 1, 0, 0, 0, 81, 284, 1, 0, 0, 0, 83, 286,
		1, 0, 0, 0, 85, 289, 1, 0, 0, 0, 87, 292, 1, 0, 0, 0, 89, 294, 1, 0, 0,
		0, 91, 297, 1, 0, 0, 0, 93, 300, 1, 0, 0, 0, 95, 303, 1, 0, 0, 0, 97, 307,
		1, 0, 0, 0, 99, 313, 1, 0, 0, 0, 101, 328, 1, 0, 0, 0, 103, 343, 1, 0,
		0, 0, 105, 348, 1, 0, 0, 0, 107, 360, 1, 0, 0, 0, 109, 364, 1, 0, 0, 0,
		111, 371, 1, 0, 0, 0, 113, 378, 1, 0, 0, 0, 115, 389, 1, 0, 0, 0, 117,
		391, 1, 0, 0, 0, 119, 396, 1, 0, 0, 0, 121, 400, 1, 0, 0, 0, 123, 403,
		1, 0, 0, 0, 125, 126, 5, 59, 0, 0, 126, 2, 1, 0, 0, 0, 127, 128, 5, 123,
		0, 0, 128, 4, 1, 0, 0, 0, 129, 130, 5, 125, 0, 0, 130, 6, 1, 0, 0, 0, 131,
		132, 5, 102, 0, 0, 132, 133, 5, 111, 0, 0, 133, 134, 5, 114, 0, 0, 134,
		8, 1, 0, 0, 0, 135, 136, 5, 105, 0, 0, 136, 137, 5, 110, 0, 0, 137, 10,
		1, 0, 0, 0, 138, 139, 5, 46, 0, 0, 139, 140, 5, 46, 0, 0, 140, 141, 5,
		46, 0, 0, 141, 12, 1, 0, 0, 0, 142, 143, 5, 119, 0, 0, 143, 144, 5, 104,
		0, 0, 144, 145, 5, 105, 0, 0, 145, 146, 5, 108, 0, 0, 146, 147, 5, 101,
		0, 0, 147, 14, 1, 0, 0, 0, 148, 149, 5, 105, 0, 0, 149, 150, 5, 102, 0,
		0, 150, 16, 1, 0, 0, 0, 151, 152, 5, 101, 0, 0, 152, 153, 5, 108, 0, 0,
		153, 154, 5, 115, 0, 0, 154, 155, 5, 101, 0, 0, 155, 18, 1, 0, 0, 0, 156,
		157, 5, 103, 0, 0, 157, 158, 5, 117, 0, 0, 158, 159, 5, 97, 0, 0, 159,
		160, 5, 114, 0, 0, 160, 161, 5, 100, 0, 0, 161, 20, 1, 0, 0, 0, 162, 163,
		5, 99, 0, 0, 163, 164, 5, 111, 0, 0, 164, 165, 5, 110, 0, 0, 165, 166,
		5, 116, 0, 0, 166, 167, 5, 105, 0, 0, 167, 168, 5, 110, 0, 0, 168, 169,
		5, 117, 0, 0, 169, 170, 5, 101, 0, 0, 170, 22, 1, 0, 0, 0, 171, 172, 5,
		98, 0, 0, 172, 173, 5, 114, 0, 0, 173, 174, 5, 101, 0, 0, 174, 175, 5,
		97, 0, 0, 175, 176, 5, 107, 0, 0, 176, 24, 1, 0, 0, 0, 177, 178, 5, 114,
		0, 0, 178, 179, 5, 101, 0, 0, 179, 180, 5, 116, 0, 0, 180, 181, 5, 117,
		0, 0, 181, 182, 5, 114, 0, 0, 182, 183, 5, 110, 0, 0, 183, 26, 1, 0, 0,
		0, 184, 185, 5, 115, 0, 0, 185, 186, 5, 119, 0, 0, 186, 187, 5, 105, 0,
		0, 187, 188, 5, 116, 0, 0, 188, 189, 5, 99, 0, 0, 189, 190, 5, 104, 0,
		0, 190, 28, 1, 0, 0, 0, 191, 192, 5, 99, 0, 0, 192, 193, 5, 97, 0, 0, 193,
		194, 5, 115, 0, 0, 194, 195, 5, 101, 0, 0, 195, 30, 1, 0, 0, 0, 196, 197,
		5, 58, 0, 0, 197, 32, 1, 0, 0, 0, 198, 199, 5, 100, 0, 0, 199, 200, 5,
		101, 0, 0, 200, 201, 5, 102, 0, 0, 201, 202, 5, 97, 0, 0, 202, 203, 5,
		117, 0, 0, 203, 204, 5, 108, 0, 0, 204, 205, 5, 116, 0, 0, 205, 34, 1,
		0, 0, 0, 206, 207, 5, 108, 0, 0, 207, 208, 5, 101, 0, 0, 208, 209, 5, 116,
		0, 0, 209, 36, 1, 0, 0, 0, 210, 211, 5, 61, 0, 0, 211, 38, 1, 0, 0, 0,
		212, 213, 5, 118, 0, 0, 213, 214, 5, 97, 0, 0, 214, 215, 5, 114, 0, 0,
		215, 40, 1, 0, 0, 0, 216, 217, 5, 63, 0, 0, 217, 42, 1, 0, 0, 0, 218, 219,
		5, 83, 0, 0, 219, 220, 5, 116, 0, 0, 220, 221, 5, 114, 0, 0, 221, 222,
		5, 105, 0, 0, 222, 223, 5, 110, 0, 0, 223, 224, 5, 103, 0, 0, 224, 44,
		1, 0, 0, 0, 225, 226, 5, 73, 0, 0, 226, 227, 5, 110, 0, 0, 227, 228, 5,
		116, 0, 0, 228, 46, 1, 0, 0, 0, 229, 230, 5, 70, 0, 0, 230, 231, 5, 108,
		0, 0, 231, 232, 5, 111, 0, 0, 232, 233, 5, 97, 0, 0, 233, 234, 5, 116,
		0, 0, 234, 48, 1, 0, 0, 0, 235, 236, 5, 66, 0, 0, 236, 237, 5, 111, 0,
		0, 237, 238, 5, 111, 0, 0, 238, 239, 5, 108, 0, 0, 239, 50, 1, 0, 0, 0,
		240, 241, 5, 67, 0, 0, 241, 242, 5, 104, 0, 0, 242, 243, 5, 97, 0, 0, 243,
		244, 5, 114, 0, 0, 244, 245, 5, 97, 0, 0, 245, 246, 5, 99, 0, 0, 246, 247,
		5, 116, 0, 0, 247, 248, 5, 101, 0, 0, 248, 249, 5, 114, 0, 0, 249, 52,
		1, 0, 0, 0, 250, 251, 5, 91, 0, 0, 251, 54, 1, 0, 0, 0, 252, 253, 5, 93,
		0, 0, 253, 56, 1, 0, 0, 0, 254, 255, 5, 44, 0, 0, 255, 58, 1, 0, 0, 0,
		256, 257, 5, 112, 0, 0, 257, 258, 5, 114, 0, 0, 258, 259, 5, 105, 0, 0,
		259, 260, 5, 110, 0, 0, 260, 261, 5, 116, 0, 0, 261, 60, 1, 0, 0, 0, 262,
		263, 5, 40, 0, 0, 263, 62, 1, 0, 0, 0, 264, 265, 5, 41, 0, 0, 265, 64,
		1, 0, 0, 0, 266, 267, 5, 43, 0, 0, 267, 268, 5, 61, 0, 0, 268, 66, 1, 0,
		0, 0, 269, 270, 5, 45, 0, 0, 270, 271, 5, 61, 0, 0, 271, 68, 1, 0, 0, 0,
		272, 273, 5, 33, 0, 0, 273, 70, 1, 0, 0, 0, 274, 275, 5, 45, 0, 0, 275,
		72, 1, 0, 0, 0, 276, 277, 5, 47, 0, 0, 277, 74, 1, 0, 0, 0, 278, 279, 5,
		37, 0, 0, 279, 76, 1, 0, 0, 0, 280, 281, 5, 42, 0, 0, 281, 78, 1, 0, 0,
		0, 282, 283, 5, 43, 0, 0, 283, 80, 1, 0, 0, 0, 284, 285, 5, 60, 0, 0, 285,
		82, 1, 0, 0, 0, 286, 287, 5, 60, 0, 0, 287, 288, 5, 61, 0, 0, 288, 84,
		1, 0, 0, 0, 289, 290, 5, 62, 0, 0, 290, 291, 5, 61, 0, 0, 291, 86, 1, 0,
		0, 0, 292, 293, 5, 62, 0, 0, 293, 88, 1, 0, 0, 0, 294, 295, 5, 61, 0, 0,
		295, 296, 5, 61, 0, 0, 296, 90, 1, 0, 0, 0, 297, 298, 5, 33, 0, 0, 298,
		299, 5, 61, 0, 0, 299, 92, 1, 0, 0, 0, 300, 301, 5, 38, 0, 0, 301, 302,
		5, 38, 0, 0, 302, 94, 1, 0, 0, 0, 303, 304, 5, 124, 0, 0, 304, 305, 5,
		124, 0, 0, 305, 96, 1, 0, 0, 0, 306, 308, 7, 0, 0, 0, 307, 306, 1, 0, 0,
		0, 308, 309, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 309, 310, 1, 0, 0, 0, 310,
		311, 1, 0, 0, 0, 311, 312, 6, 48, 0, 0, 312, 98, 1, 0, 0, 0, 313, 314,
		5, 47, 0, 0, 314, 315, 5, 42, 0, 0, 315, 320, 1, 0, 0, 0, 316, 319, 3,
		99, 49, 0, 317, 319, 9, 0, 0, 0, 318, 316, 1, 0, 0, 0, 318, 317, 1, 0,
		0, 0, 319, 322, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0, 320, 318, 1, 0, 0, 0,
		321, 323, 1, 0, 0, 0, 322, 320, 1, 0, 0, 0, 323, 324, 5, 42, 0, 0, 324,
		325, 5, 47, 0, 0, 325, 326, 1, 0, 0, 0, 326, 327, 6, 49, 0, 0, 327, 100,
		1, 0, 0, 0, 328, 329, 5, 47, 0, 0, 329, 330, 5, 47, 0, 0, 330, 334, 1,
		0, 0, 0, 331, 333, 9, 0, 0, 0, 332, 331, 1, 0, 0, 0, 333, 336, 1, 0, 0,
		0, 334, 335, 1, 0, 0, 0, 334, 332, 1, 0, 0, 0, 335, 338, 1, 0, 0, 0, 336,
		334, 1, 0, 0, 0, 337, 339, 7, 1, 0, 0, 338, 337, 1, 0, 0, 0, 339, 340,
		1, 0, 0, 0, 340, 341, 6, 50, 0, 0, 341, 102, 1, 0, 0, 0, 342, 344, 7, 2,
		0, 0, 343, 342, 1, 0, 0, 0, 344, 345, 1, 0, 0, 0, 345, 343, 1, 0, 0, 0,
		345, 346, 1, 0, 0, 0, 346, 104, 1, 0, 0, 0, 347, 349, 7, 2, 0, 0, 348,
		347, 1, 0, 0, 0, 349, 350, 1, 0, 0, 0, 350, 348, 1, 0, 0, 0, 350, 351,
		1, 0, 0, 0, 351, 358, 1, 0, 0, 0, 352, 354, 5, 46, 0, 0, 353, 355, 7, 2,
		0, 0, 354, 353, 1, 0, 0, 0, 355, 356, 1, 0, 0, 0, 356, 354, 1, 0, 0, 0,
		356, 357, 1, 0, 0, 0, 357, 359, 1, 0, 0, 0, 358, 352, 1, 0, 0, 0, 358,
		359, 1, 0, 0, 0, 359, 106, 1, 0, 0, 0, 360, 361, 5, 34, 0, 0, 361, 362,
		3, 113, 56, 0, 362, 363, 5, 34, 0, 0, 363, 108, 1, 0, 0, 0, 364, 366, 5,
		34, 0, 0, 365, 367, 3, 111, 55, 0, 366, 365, 1, 0, 0, 0, 366, 367, 1, 0,
		0, 0, 367, 368, 1, 0, 0, 0, 368, 369, 5, 34, 0, 0, 369, 110, 1, 0, 0, 0,
		370, 372, 3, 113, 56, 0, 371, 370, 1, 0, 0, 0, 372, 373, 1, 0, 0, 0, 373,
		371, 1, 0, 0, 0, 373, 374, 1, 0, 0, 0, 374, 112, 1, 0, 0, 0, 375, 376,
		5, 92, 0, 0, 376, 379, 7, 3, 0, 0, 377, 379, 8, 4, 0, 0, 378, 375, 1, 0,
		0, 0, 378, 377, 1, 0, 0, 0, 379, 114, 1, 0, 0, 0, 380, 381, 5, 116, 0,
		0, 381, 382, 5, 114, 0, 0, 382, 383, 5, 117, 0, 0, 383, 390, 5, 101, 0,
		0, 384, 385, 5, 102, 0, 0, 385, 386, 5, 97, 0, 0, 386, 387, 5, 108, 0,
		0, 387, 388, 5, 115, 0, 0, 388, 390, 5, 101, 0, 0, 389, 380, 1, 0, 0, 0,
		389, 384, 1, 0, 0, 0, 390, 116, 1, 0, 0, 0, 391, 393, 3, 119, 59, 0, 392,
		394, 3, 123, 61, 0, 393, 392, 1, 0, 0, 0, 393, 394, 1, 0, 0, 0, 394, 118,
		1, 0, 0, 0, 395, 397, 7, 5, 0, 0, 396, 395, 1, 0, 0, 0, 397, 120, 1, 0,
		0, 0, 398, 401, 7, 2, 0, 0, 399, 401, 3, 119, 59, 0, 400, 398, 1, 0, 0,
		0, 400, 399, 1, 0, 0, 0, 401, 122, 1, 0, 0, 0, 402, 404, 3, 121, 60, 0,
		403, 402, 1, 0, 0, 0, 404, 405, 1, 0, 0, 0, 405, 403, 1, 0, 0, 0, 405,
		406, 1, 0, 0, 0, 406, 124, 1, 0, 0, 0, 18, 0, 309, 318, 320, 334, 338,
		345, 350, 356, 358, 366, 373, 378, 389, 393, 396, 400, 405, 1, 0, 1, 0,
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
	T_swiftLexerWS            = 49
	T_swiftLexerBlock_comment = 50
	T_swiftLexerLine_comment  = 51
	T_swiftLexerInt           = 52
	T_swiftLexerFloat         = 53
	T_swiftLexerCaracter      = 54
	T_swiftLexerString_       = 55
	T_swiftLexerBool          = 56
	T_swiftLexerIdentificador = 57
)
