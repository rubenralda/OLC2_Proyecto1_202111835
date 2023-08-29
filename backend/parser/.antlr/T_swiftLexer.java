// Generated from c:\Users\ruben\Desktop\Compi_2\Lab\Proyecto 1\backend\parser\T_swift.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class T_swiftLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, T__27=28, T__28=29, T__29=30, T__30=31, 
		T__31=32, T__32=33, T__33=34, T__34=35, T__35=36, T__36=37, T__37=38, 
		T__38=39, T__39=40, T__40=41, T__41=42, T__42=43, T__43=44, WS=45, Block_comment=46, 
		Line_comment=47, Int=48, Float=49, Caracter=50, String=51, Bool=52, Identificador=53;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
			"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16", 
			"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24", 
			"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32", 
			"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40", 
			"T__41", "T__42", "T__43", "WS", "Block_comment", "Line_comment", "Int", 
			"Float", "Caracter", "String", "Quoted_text", "Quoted_text_item", "Bool", 
			"Identificador", "Identifier_head", "Identifier_character", "Identifier_characters"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "';'", "'{'", "'}'", "'for'", "'in'", "'while'", "'if'", "'else'", 
			"'guard'", "'switch'", "'case'", "':'", "'default'", "'break'", "'continue'", 
			"'return'", "'let '", "'='", "'var'", "'?'", "'String'", "'Int'", "'Float'", 
			"'Bool'", "'['", "']'", "','", "'print'", "'('", "')'", "'-'", "'*'", 
			"'/'", "'%'", "'+'", "'=='", "'!='", "'>'", "'<'", "'<='", "'>='", "'!'", 
			"'&&'", "'||'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, "WS", "Block_comment", 
			"Line_comment", "Int", "Float", "Caracter", "String", "Bool", "Identificador"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public T_swiftLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "T_swift.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\67\u017e\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64"+
		"\t\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\3\2\3\2\3"+
		"\3\3\3\3\4\3\4\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\b"+
		"\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3"+
		"\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22"+
		"\3\22\3\22\3\23\3\23\3\24\3\24\3\24\3\24\3\25\3\25\3\26\3\26\3\26\3\26"+
		"\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\30\3\31"+
		"\3\31\3\31\3\31\3\31\3\32\3\32\3\33\3\33\3\34\3\34\3\35\3\35\3\35\3\35"+
		"\3\35\3\35\3\36\3\36\3\37\3\37\3 \3 \3!\3!\3\"\3\"\3#\3#\3$\3$\3%\3%\3"+
		"%\3&\3&\3&\3\'\3\'\3(\3(\3)\3)\3)\3*\3*\3*\3+\3+\3,\3,\3,\3-\3-\3-\3."+
		"\6.\u011b\n.\r.\16.\u011c\3.\3.\3/\3/\3/\3/\3/\7/\u0126\n/\f/\16/\u0129"+
		"\13/\3/\3/\3/\3/\3/\3\60\3\60\3\60\3\60\7\60\u0134\n\60\f\60\16\60\u0137"+
		"\13\60\3\60\5\60\u013a\n\60\3\60\3\60\3\61\6\61\u013f\n\61\r\61\16\61"+
		"\u0140\3\62\6\62\u0144\n\62\r\62\16\62\u0145\3\62\3\62\6\62\u014a\n\62"+
		"\r\62\16\62\u014b\5\62\u014e\n\62\3\63\3\63\3\63\3\63\3\64\3\64\5\64\u0156"+
		"\n\64\3\64\3\64\3\65\6\65\u015b\n\65\r\65\16\65\u015c\3\66\3\66\3\66\5"+
		"\66\u0162\n\66\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67\5\67\u016d"+
		"\n\67\38\38\58\u0171\n8\39\59\u0174\n9\3:\3:\5:\u0178\n:\3;\6;\u017b\n"+
		";\r;\16;\u017c\4\u0127\u0135\2<\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23"+
		"\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31"+
		"\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60"+
		"_\61a\62c\63e\64g\65i\2k\2m\66o\67q\2s\2u\2\3\2\b\5\2\2\2\13\17\"\"\3"+
		"\3\f\f\3\2\62;\t\2$$))\62\62^^ppttvv\6\2\f\f\17\17$$^^\5\2C\\aac|\2\u0187"+
		"\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2"+
		"\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2"+
		"\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2"+
		"\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2"+
		"\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3"+
		"\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2"+
		"\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2"+
		"U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3"+
		"\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\3w\3\2\2"+
		"\2\5y\3\2\2\2\7{\3\2\2\2\t}\3\2\2\2\13\u0081\3\2\2\2\r\u0084\3\2\2\2\17"+
		"\u008a\3\2\2\2\21\u008d\3\2\2\2\23\u0092\3\2\2\2\25\u0098\3\2\2\2\27\u009f"+
		"\3\2\2\2\31\u00a4\3\2\2\2\33\u00a6\3\2\2\2\35\u00ae\3\2\2\2\37\u00b4\3"+
		"\2\2\2!\u00bd\3\2\2\2#\u00c4\3\2\2\2%\u00c9\3\2\2\2\'\u00cb\3\2\2\2)\u00cf"+
		"\3\2\2\2+\u00d1\3\2\2\2-\u00d8\3\2\2\2/\u00dc\3\2\2\2\61\u00e2\3\2\2\2"+
		"\63\u00e7\3\2\2\2\65\u00e9\3\2\2\2\67\u00eb\3\2\2\29\u00ed\3\2\2\2;\u00f3"+
		"\3\2\2\2=\u00f5\3\2\2\2?\u00f7\3\2\2\2A\u00f9\3\2\2\2C\u00fb\3\2\2\2E"+
		"\u00fd\3\2\2\2G\u00ff\3\2\2\2I\u0101\3\2\2\2K\u0104\3\2\2\2M\u0107\3\2"+
		"\2\2O\u0109\3\2\2\2Q\u010b\3\2\2\2S\u010e\3\2\2\2U\u0111\3\2\2\2W\u0113"+
		"\3\2\2\2Y\u0116\3\2\2\2[\u011a\3\2\2\2]\u0120\3\2\2\2_\u012f\3\2\2\2a"+
		"\u013e\3\2\2\2c\u0143\3\2\2\2e\u014f\3\2\2\2g\u0153\3\2\2\2i\u015a\3\2"+
		"\2\2k\u0161\3\2\2\2m\u016c\3\2\2\2o\u016e\3\2\2\2q\u0173\3\2\2\2s\u0177"+
		"\3\2\2\2u\u017a\3\2\2\2wx\7=\2\2x\4\3\2\2\2yz\7}\2\2z\6\3\2\2\2{|\7\177"+
		"\2\2|\b\3\2\2\2}~\7h\2\2~\177\7q\2\2\177\u0080\7t\2\2\u0080\n\3\2\2\2"+
		"\u0081\u0082\7k\2\2\u0082\u0083\7p\2\2\u0083\f\3\2\2\2\u0084\u0085\7y"+
		"\2\2\u0085\u0086\7j\2\2\u0086\u0087\7k\2\2\u0087\u0088\7n\2\2\u0088\u0089"+
		"\7g\2\2\u0089\16\3\2\2\2\u008a\u008b\7k\2\2\u008b\u008c\7h\2\2\u008c\20"+
		"\3\2\2\2\u008d\u008e\7g\2\2\u008e\u008f\7n\2\2\u008f\u0090\7u\2\2\u0090"+
		"\u0091\7g\2\2\u0091\22\3\2\2\2\u0092\u0093\7i\2\2\u0093\u0094\7w\2\2\u0094"+
		"\u0095\7c\2\2\u0095\u0096\7t\2\2\u0096\u0097\7f\2\2\u0097\24\3\2\2\2\u0098"+
		"\u0099\7u\2\2\u0099\u009a\7y\2\2\u009a\u009b\7k\2\2\u009b\u009c\7v\2\2"+
		"\u009c\u009d\7e\2\2\u009d\u009e\7j\2\2\u009e\26\3\2\2\2\u009f\u00a0\7"+
		"e\2\2\u00a0\u00a1\7c\2\2\u00a1\u00a2\7u\2\2\u00a2\u00a3\7g\2\2\u00a3\30"+
		"\3\2\2\2\u00a4\u00a5\7<\2\2\u00a5\32\3\2\2\2\u00a6\u00a7\7f\2\2\u00a7"+
		"\u00a8\7g\2\2\u00a8\u00a9\7h\2\2\u00a9\u00aa\7c\2\2\u00aa\u00ab\7w\2\2"+
		"\u00ab\u00ac\7n\2\2\u00ac\u00ad\7v\2\2\u00ad\34\3\2\2\2\u00ae\u00af\7"+
		"d\2\2\u00af\u00b0\7t\2\2\u00b0\u00b1\7g\2\2\u00b1\u00b2\7c\2\2\u00b2\u00b3"+
		"\7m\2\2\u00b3\36\3\2\2\2\u00b4\u00b5\7e\2\2\u00b5\u00b6\7q\2\2\u00b6\u00b7"+
		"\7p\2\2\u00b7\u00b8\7v\2\2\u00b8\u00b9\7k\2\2\u00b9\u00ba\7p\2\2\u00ba"+
		"\u00bb\7w\2\2\u00bb\u00bc\7g\2\2\u00bc \3\2\2\2\u00bd\u00be\7t\2\2\u00be"+
		"\u00bf\7g\2\2\u00bf\u00c0\7v\2\2\u00c0\u00c1\7w\2\2\u00c1\u00c2\7t\2\2"+
		"\u00c2\u00c3\7p\2\2\u00c3\"\3\2\2\2\u00c4\u00c5\7n\2\2\u00c5\u00c6\7g"+
		"\2\2\u00c6\u00c7\7v\2\2\u00c7\u00c8\7\"\2\2\u00c8$\3\2\2\2\u00c9\u00ca"+
		"\7?\2\2\u00ca&\3\2\2\2\u00cb\u00cc\7x\2\2\u00cc\u00cd\7c\2\2\u00cd\u00ce"+
		"\7t\2\2\u00ce(\3\2\2\2\u00cf\u00d0\7A\2\2\u00d0*\3\2\2\2\u00d1\u00d2\7"+
		"U\2\2\u00d2\u00d3\7v\2\2\u00d3\u00d4\7t\2\2\u00d4\u00d5\7k\2\2\u00d5\u00d6"+
		"\7p\2\2\u00d6\u00d7\7i\2\2\u00d7,\3\2\2\2\u00d8\u00d9\7K\2\2\u00d9\u00da"+
		"\7p\2\2\u00da\u00db\7v\2\2\u00db.\3\2\2\2\u00dc\u00dd\7H\2\2\u00dd\u00de"+
		"\7n\2\2\u00de\u00df\7q\2\2\u00df\u00e0\7c\2\2\u00e0\u00e1\7v\2\2\u00e1"+
		"\60\3\2\2\2\u00e2\u00e3\7D\2\2\u00e3\u00e4\7q\2\2\u00e4\u00e5\7q\2\2\u00e5"+
		"\u00e6\7n\2\2\u00e6\62\3\2\2\2\u00e7\u00e8\7]\2\2\u00e8\64\3\2\2\2\u00e9"+
		"\u00ea\7_\2\2\u00ea\66\3\2\2\2\u00eb\u00ec\7.\2\2\u00ec8\3\2\2\2\u00ed"+
		"\u00ee\7r\2\2\u00ee\u00ef\7t\2\2\u00ef\u00f0\7k\2\2\u00f0\u00f1\7p\2\2"+
		"\u00f1\u00f2\7v\2\2\u00f2:\3\2\2\2\u00f3\u00f4\7*\2\2\u00f4<\3\2\2\2\u00f5"+
		"\u00f6\7+\2\2\u00f6>\3\2\2\2\u00f7\u00f8\7/\2\2\u00f8@\3\2\2\2\u00f9\u00fa"+
		"\7,\2\2\u00faB\3\2\2\2\u00fb\u00fc\7\61\2\2\u00fcD\3\2\2\2\u00fd\u00fe"+
		"\7\'\2\2\u00feF\3\2\2\2\u00ff\u0100\7-\2\2\u0100H\3\2\2\2\u0101\u0102"+
		"\7?\2\2\u0102\u0103\7?\2\2\u0103J\3\2\2\2\u0104\u0105\7#\2\2\u0105\u0106"+
		"\7?\2\2\u0106L\3\2\2\2\u0107\u0108\7@\2\2\u0108N\3\2\2\2\u0109\u010a\7"+
		">\2\2\u010aP\3\2\2\2\u010b\u010c\7>\2\2\u010c\u010d\7?\2\2\u010dR\3\2"+
		"\2\2\u010e\u010f\7@\2\2\u010f\u0110\7?\2\2\u0110T\3\2\2\2\u0111\u0112"+
		"\7#\2\2\u0112V\3\2\2\2\u0113\u0114\7(\2\2\u0114\u0115\7(\2\2\u0115X\3"+
		"\2\2\2\u0116\u0117\7~\2\2\u0117\u0118\7~\2\2\u0118Z\3\2\2\2\u0119\u011b"+
		"\t\2\2\2\u011a\u0119\3\2\2\2\u011b\u011c\3\2\2\2\u011c\u011a\3\2\2\2\u011c"+
		"\u011d\3\2\2\2\u011d\u011e\3\2\2\2\u011e\u011f\b.\2\2\u011f\\\3\2\2\2"+
		"\u0120\u0121\7\61\2\2\u0121\u0122\7,\2\2\u0122\u0127\3\2\2\2\u0123\u0126"+
		"\5]/\2\u0124\u0126\13\2\2\2\u0125\u0123\3\2\2\2\u0125\u0124\3\2\2\2\u0126"+
		"\u0129\3\2\2\2\u0127\u0128\3\2\2\2\u0127\u0125\3\2\2\2\u0128\u012a\3\2"+
		"\2\2\u0129\u0127\3\2\2\2\u012a\u012b\7,\2\2\u012b\u012c\7\61\2\2\u012c"+
		"\u012d\3\2\2\2\u012d\u012e\b/\2\2\u012e^\3\2\2\2\u012f\u0130\7\61\2\2"+
		"\u0130\u0131\7\61\2\2\u0131\u0135\3\2\2\2\u0132\u0134\13\2\2\2\u0133\u0132"+
		"\3\2\2\2\u0134\u0137\3\2\2\2\u0135\u0136\3\2\2\2\u0135\u0133\3\2\2\2\u0136"+
		"\u0139\3\2\2\2\u0137\u0135\3\2\2\2\u0138\u013a\t\3\2\2\u0139\u0138\3\2"+
		"\2\2\u013a\u013b\3\2\2\2\u013b\u013c\b\60\2\2\u013c`\3\2\2\2\u013d\u013f"+
		"\t\4\2\2\u013e\u013d\3\2\2\2\u013f\u0140\3\2\2\2\u0140\u013e\3\2\2\2\u0140"+
		"\u0141\3\2\2\2\u0141b\3\2\2\2\u0142\u0144\t\4\2\2\u0143\u0142\3\2\2\2"+
		"\u0144\u0145\3\2\2\2\u0145\u0143\3\2\2\2\u0145\u0146\3\2\2\2\u0146\u014d"+
		"\3\2\2\2\u0147\u0149\7\60\2\2\u0148\u014a\t\4\2\2\u0149\u0148\3\2\2\2"+
		"\u014a\u014b\3\2\2\2\u014b\u0149\3\2\2\2\u014b\u014c\3\2\2\2\u014c\u014e"+
		"\3\2\2\2\u014d\u0147\3\2\2\2\u014d\u014e\3\2\2\2\u014ed\3\2\2\2\u014f"+
		"\u0150\7$\2\2\u0150\u0151\5k\66\2\u0151\u0152\7$\2\2\u0152f\3\2\2\2\u0153"+
		"\u0155\7$\2\2\u0154\u0156\5i\65\2\u0155\u0154\3\2\2\2\u0155\u0156\3\2"+
		"\2\2\u0156\u0157\3\2\2\2\u0157\u0158\7$\2\2\u0158h\3\2\2\2\u0159\u015b"+
		"\5k\66\2\u015a\u0159\3\2\2\2\u015b\u015c\3\2\2\2\u015c\u015a\3\2\2\2\u015c"+
		"\u015d\3\2\2\2\u015dj\3\2\2\2\u015e\u015f\7^\2\2\u015f\u0162\t\5\2\2\u0160"+
		"\u0162\n\6\2\2\u0161\u015e\3\2\2\2\u0161\u0160\3\2\2\2\u0162l\3\2\2\2"+
		"\u0163\u0164\7v\2\2\u0164\u0165\7t\2\2\u0165\u0166\7w\2\2\u0166\u016d"+
		"\7g\2\2\u0167\u0168\7h\2\2\u0168\u0169\7c\2\2\u0169\u016a\7n\2\2\u016a"+
		"\u016b\7u\2\2\u016b\u016d\7g\2\2\u016c\u0163\3\2\2\2\u016c\u0167\3\2\2"+
		"\2\u016dn\3\2\2\2\u016e\u0170\5q9\2\u016f\u0171\5u;\2\u0170\u016f\3\2"+
		"\2\2\u0170\u0171\3\2\2\2\u0171p\3\2\2\2\u0172\u0174\t\7\2\2\u0173\u0172"+
		"\3\2\2\2\u0174r\3\2\2\2\u0175\u0178\t\4\2\2\u0176\u0178\5q9\2\u0177\u0175"+
		"\3\2\2\2\u0177\u0176\3\2\2\2\u0178t\3\2\2\2\u0179\u017b\5s:\2\u017a\u0179"+
		"\3\2\2\2\u017b\u017c\3\2\2\2\u017c\u017a\3\2\2\2\u017c\u017d\3\2\2\2\u017d"+
		"v\3\2\2\2\24\2\u011c\u0125\u0127\u0135\u0139\u0140\u0145\u014b\u014d\u0155"+
		"\u015c\u0161\u016c\u0170\u0173\u0177\u017c\3\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}