// Generated from c:\Users\ruben\Desktop\Compi_2\Lab\Proyecto 1\backend\parser\T_swift.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class T_swiftParser extends Parser {
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
	public static final int
		RULE_inicio = 0, RULE_instrucciones = 1, RULE_instruccion = 2, RULE_loop_statement = 3, 
		RULE_code_block = 4, RULE_for_in_statement = 5, RULE_while_statement = 6, 
		RULE_branch_statement = 7, RULE_if_statement = 8, RULE_else_clause = 9, 
		RULE_guard_statement = 10, RULE_switch_statement = 11, RULE_switch_cases = 12, 
		RULE_switch_case = 13, RULE_case_label = 14, RULE_default_label = 15, 
		RULE_control_transfer_statement = 16, RULE_break_statement = 17, RULE_continue_statement = 18, 
		RULE_return_statement = 19, RULE_declaracion = 20, RULE_constant_declaracion = 21, 
		RULE_variable_declaracion = 22, RULE_anotacion_tipo = 23, RULE_tipos = 24, 
		RULE_array_declaracion = 25, RULE_definicion_vector = 26, RULE_lista_expresiones = 27, 
		RULE_funcion_print = 28, RULE_expresion = 29, RULE_primitivos = 30;
	private static String[] makeRuleNames() {
		return new String[] {
			"inicio", "instrucciones", "instruccion", "loop_statement", "code_block", 
			"for_in_statement", "while_statement", "branch_statement", "if_statement", 
			"else_clause", "guard_statement", "switch_statement", "switch_cases", 
			"switch_case", "case_label", "default_label", "control_transfer_statement", 
			"break_statement", "continue_statement", "return_statement", "declaracion", 
			"constant_declaracion", "variable_declaracion", "anotacion_tipo", "tipos", 
			"array_declaracion", "definicion_vector", "lista_expresiones", "funcion_print", 
			"expresion", "primitivos"
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

	@Override
	public String getGrammarFileName() { return "T_swift.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public T_swiftParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class InicioContext extends ParserRuleContext {
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode EOF() { return getToken(T_swiftParser.EOF, 0); }
		public InicioContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_inicio; }
	}

	public final InicioContext inicio() throws RecognitionException {
		InicioContext _localctx = new InicioContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_inicio);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(62);
			instrucciones();
			setState(63);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InstruccionesContext extends ParserRuleContext {
		public List<InstruccionContext> instruccion() {
			return getRuleContexts(InstruccionContext.class);
		}
		public InstruccionContext instruccion(int i) {
			return getRuleContext(InstruccionContext.class,i);
		}
		public InstruccionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instrucciones; }
	}

	public final InstruccionesContext instrucciones() throws RecognitionException {
		InstruccionesContext _localctx = new InstruccionesContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_instrucciones);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(68);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__3) | (1L << T__5) | (1L << T__6) | (1L << T__8) | (1L << T__9) | (1L << T__13) | (1L << T__14) | (1L << T__15) | (1L << T__16) | (1L << T__18) | (1L << T__27))) != 0)) {
				{
				{
				setState(65);
				instruccion();
				}
				}
				setState(70);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InstruccionContext extends ParserRuleContext {
		public DeclaracionContext declaracion() {
			return getRuleContext(DeclaracionContext.class,0);
		}
		public Loop_statementContext loop_statement() {
			return getRuleContext(Loop_statementContext.class,0);
		}
		public Branch_statementContext branch_statement() {
			return getRuleContext(Branch_statementContext.class,0);
		}
		public Control_transfer_statementContext control_transfer_statement() {
			return getRuleContext(Control_transfer_statementContext.class,0);
		}
		public Funcion_printContext funcion_print() {
			return getRuleContext(Funcion_printContext.class,0);
		}
		public InstruccionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion; }
	}

	public final InstruccionContext instruccion() throws RecognitionException {
		InstruccionContext _localctx = new InstruccionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruccion);
		int _la;
		try {
			setState(91);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__16:
			case T__18:
				enterOuterAlt(_localctx, 1);
				{
				setState(71);
				declaracion();
				setState(73);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__0) {
					{
					setState(72);
					match(T__0);
					}
				}

				}
				break;
			case T__3:
			case T__5:
				enterOuterAlt(_localctx, 2);
				{
				setState(75);
				loop_statement();
				setState(77);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__0) {
					{
					setState(76);
					match(T__0);
					}
				}

				}
				break;
			case T__6:
			case T__8:
			case T__9:
				enterOuterAlt(_localctx, 3);
				{
				setState(79);
				branch_statement();
				setState(81);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__0) {
					{
					setState(80);
					match(T__0);
					}
				}

				}
				break;
			case T__13:
			case T__14:
			case T__15:
				enterOuterAlt(_localctx, 4);
				{
				setState(83);
				control_transfer_statement();
				setState(85);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__0) {
					{
					setState(84);
					match(T__0);
					}
				}

				}
				break;
			case T__27:
				enterOuterAlt(_localctx, 5);
				{
				setState(87);
				funcion_print();
				setState(89);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__0) {
					{
					setState(88);
					match(T__0);
					}
				}

				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Loop_statementContext extends ParserRuleContext {
		public For_in_statementContext for_in_statement() {
			return getRuleContext(For_in_statementContext.class,0);
		}
		public While_statementContext while_statement() {
			return getRuleContext(While_statementContext.class,0);
		}
		public Loop_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_loop_statement; }
	}

	public final Loop_statementContext loop_statement() throws RecognitionException {
		Loop_statementContext _localctx = new Loop_statementContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_loop_statement);
		try {
			setState(95);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__3:
				enterOuterAlt(_localctx, 1);
				{
				setState(93);
				for_in_statement();
				}
				break;
			case T__5:
				enterOuterAlt(_localctx, 2);
				{
				setState(94);
				while_statement();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Code_blockContext extends ParserRuleContext {
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public Code_blockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_code_block; }
	}

	public final Code_blockContext code_block() throws RecognitionException {
		Code_blockContext _localctx = new Code_blockContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_code_block);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(97);
			match(T__1);
			setState(98);
			instrucciones();
			setState(99);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class For_in_statementContext extends ParserRuleContext {
		public TerminalNode Identificador() { return getToken(T_swiftParser.Identificador, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Code_blockContext code_block() {
			return getRuleContext(Code_blockContext.class,0);
		}
		public For_in_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_for_in_statement; }
	}

	public final For_in_statementContext for_in_statement() throws RecognitionException {
		For_in_statementContext _localctx = new For_in_statementContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_for_in_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(101);
			match(T__3);
			setState(102);
			match(Identificador);
			setState(103);
			match(T__4);
			setState(104);
			expresion(0);
			setState(105);
			code_block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class While_statementContext extends ParserRuleContext {
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Code_blockContext code_block() {
			return getRuleContext(Code_blockContext.class,0);
		}
		public While_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_while_statement; }
	}

	public final While_statementContext while_statement() throws RecognitionException {
		While_statementContext _localctx = new While_statementContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_while_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(107);
			match(T__5);
			setState(108);
			expresion(0);
			setState(109);
			code_block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Branch_statementContext extends ParserRuleContext {
		public If_statementContext if_statement() {
			return getRuleContext(If_statementContext.class,0);
		}
		public Guard_statementContext guard_statement() {
			return getRuleContext(Guard_statementContext.class,0);
		}
		public Switch_statementContext switch_statement() {
			return getRuleContext(Switch_statementContext.class,0);
		}
		public Branch_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_branch_statement; }
	}

	public final Branch_statementContext branch_statement() throws RecognitionException {
		Branch_statementContext _localctx = new Branch_statementContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_branch_statement);
		try {
			setState(114);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__6:
				enterOuterAlt(_localctx, 1);
				{
				setState(111);
				if_statement();
				}
				break;
			case T__8:
				enterOuterAlt(_localctx, 2);
				{
				setState(112);
				guard_statement();
				}
				break;
			case T__9:
				enterOuterAlt(_localctx, 3);
				{
				setState(113);
				switch_statement();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class If_statementContext extends ParserRuleContext {
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Code_blockContext code_block() {
			return getRuleContext(Code_blockContext.class,0);
		}
		public Else_clauseContext else_clause() {
			return getRuleContext(Else_clauseContext.class,0);
		}
		public If_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_statement; }
	}

	public final If_statementContext if_statement() throws RecognitionException {
		If_statementContext _localctx = new If_statementContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_if_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(116);
			match(T__6);
			setState(117);
			expresion(0);
			setState(118);
			code_block();
			setState(120);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__7) {
				{
				setState(119);
				else_clause();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Else_clauseContext extends ParserRuleContext {
		public Code_blockContext code_block() {
			return getRuleContext(Code_blockContext.class,0);
		}
		public If_statementContext if_statement() {
			return getRuleContext(If_statementContext.class,0);
		}
		public Else_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_else_clause; }
	}

	public final Else_clauseContext else_clause() throws RecognitionException {
		Else_clauseContext _localctx = new Else_clauseContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_else_clause);
		try {
			setState(126);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(122);
				match(T__7);
				setState(123);
				code_block();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(124);
				match(T__7);
				setState(125);
				if_statement();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Guard_statementContext extends ParserRuleContext {
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Code_blockContext code_block() {
			return getRuleContext(Code_blockContext.class,0);
		}
		public Guard_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_guard_statement; }
	}

	public final Guard_statementContext guard_statement() throws RecognitionException {
		Guard_statementContext _localctx = new Guard_statementContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_guard_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(128);
			match(T__8);
			setState(129);
			expresion(0);
			setState(130);
			match(T__7);
			setState(131);
			code_block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Switch_statementContext extends ParserRuleContext {
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Switch_casesContext switch_cases() {
			return getRuleContext(Switch_casesContext.class,0);
		}
		public Switch_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switch_statement; }
	}

	public final Switch_statementContext switch_statement() throws RecognitionException {
		Switch_statementContext _localctx = new Switch_statementContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_switch_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(133);
			match(T__9);
			setState(134);
			expresion(0);
			setState(135);
			match(T__1);
			setState(137);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__10 || _la==T__12) {
				{
				setState(136);
				switch_cases();
				}
			}

			setState(139);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Switch_casesContext extends ParserRuleContext {
		public Switch_caseContext switch_case() {
			return getRuleContext(Switch_caseContext.class,0);
		}
		public Switch_casesContext switch_cases() {
			return getRuleContext(Switch_casesContext.class,0);
		}
		public Switch_casesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switch_cases; }
	}

	public final Switch_casesContext switch_cases() throws RecognitionException {
		Switch_casesContext _localctx = new Switch_casesContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_switch_cases);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(141);
			switch_case();
			setState(143);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__10 || _la==T__12) {
				{
				setState(142);
				switch_cases();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Switch_caseContext extends ParserRuleContext {
		public Case_labelContext case_label() {
			return getRuleContext(Case_labelContext.class,0);
		}
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public Default_labelContext default_label() {
			return getRuleContext(Default_labelContext.class,0);
		}
		public Switch_caseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switch_case; }
	}

	public final Switch_caseContext switch_case() throws RecognitionException {
		Switch_caseContext _localctx = new Switch_caseContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_switch_case);
		try {
			setState(151);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__10:
				enterOuterAlt(_localctx, 1);
				{
				setState(145);
				case_label();
				setState(146);
				instrucciones();
				}
				break;
			case T__12:
				enterOuterAlt(_localctx, 2);
				{
				setState(148);
				default_label();
				setState(149);
				instrucciones();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Case_labelContext extends ParserRuleContext {
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Case_labelContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_case_label; }
	}

	public final Case_labelContext case_label() throws RecognitionException {
		Case_labelContext _localctx = new Case_labelContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_case_label);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(153);
			match(T__10);
			setState(154);
			expresion(0);
			setState(155);
			match(T__11);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Default_labelContext extends ParserRuleContext {
		public Default_labelContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_default_label; }
	}

	public final Default_labelContext default_label() throws RecognitionException {
		Default_labelContext _localctx = new Default_labelContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_default_label);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(157);
			match(T__12);
			setState(158);
			match(T__11);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Control_transfer_statementContext extends ParserRuleContext {
		public Break_statementContext break_statement() {
			return getRuleContext(Break_statementContext.class,0);
		}
		public Continue_statementContext continue_statement() {
			return getRuleContext(Continue_statementContext.class,0);
		}
		public Return_statementContext return_statement() {
			return getRuleContext(Return_statementContext.class,0);
		}
		public Control_transfer_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_control_transfer_statement; }
	}

	public final Control_transfer_statementContext control_transfer_statement() throws RecognitionException {
		Control_transfer_statementContext _localctx = new Control_transfer_statementContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_control_transfer_statement);
		try {
			setState(163);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__13:
				enterOuterAlt(_localctx, 1);
				{
				setState(160);
				break_statement();
				}
				break;
			case T__14:
				enterOuterAlt(_localctx, 2);
				{
				setState(161);
				continue_statement();
				}
				break;
			case T__15:
				enterOuterAlt(_localctx, 3);
				{
				setState(162);
				return_statement();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Break_statementContext extends ParserRuleContext {
		public Break_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_break_statement; }
	}

	public final Break_statementContext break_statement() throws RecognitionException {
		Break_statementContext _localctx = new Break_statementContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_break_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(165);
			match(T__13);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Continue_statementContext extends ParserRuleContext {
		public Continue_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_continue_statement; }
	}

	public final Continue_statementContext continue_statement() throws RecognitionException {
		Continue_statementContext _localctx = new Continue_statementContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_continue_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(167);
			match(T__14);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Return_statementContext extends ParserRuleContext {
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Return_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_return_statement; }
	}

	public final Return_statementContext return_statement() throws RecognitionException {
		Return_statementContext _localctx = new Return_statementContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_return_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(169);
			match(T__15);
			setState(171);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__28) | (1L << T__30) | (1L << T__41) | (1L << Int) | (1L << Float) | (1L << Caracter) | (1L << String) | (1L << Bool) | (1L << Identificador))) != 0)) {
				{
				setState(170);
				expresion(0);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DeclaracionContext extends ParserRuleContext {
		public Constant_declaracionContext constant_declaracion() {
			return getRuleContext(Constant_declaracionContext.class,0);
		}
		public Variable_declaracionContext variable_declaracion() {
			return getRuleContext(Variable_declaracionContext.class,0);
		}
		public Array_declaracionContext array_declaracion() {
			return getRuleContext(Array_declaracionContext.class,0);
		}
		public DeclaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion; }
	}

	public final DeclaracionContext declaracion() throws RecognitionException {
		DeclaracionContext _localctx = new DeclaracionContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_declaracion);
		try {
			setState(176);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(173);
				constant_declaracion();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(174);
				variable_declaracion();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(175);
				array_declaracion();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Constant_declaracionContext extends ParserRuleContext {
		public TerminalNode Identificador() { return getToken(T_swiftParser.Identificador, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Anotacion_tipoContext anotacion_tipo() {
			return getRuleContext(Anotacion_tipoContext.class,0);
		}
		public Constant_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_constant_declaracion; }
	}

	public final Constant_declaracionContext constant_declaracion() throws RecognitionException {
		Constant_declaracionContext _localctx = new Constant_declaracionContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_constant_declaracion);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(178);
			match(T__16);
			setState(179);
			match(Identificador);
			setState(181);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__11) {
				{
				setState(180);
				anotacion_tipo();
				}
			}

			setState(183);
			match(T__17);
			setState(184);
			expresion(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Variable_declaracionContext extends ParserRuleContext {
		public TerminalNode Identificador() { return getToken(T_swiftParser.Identificador, 0); }
		public Anotacion_tipoContext anotacion_tipo() {
			return getRuleContext(Anotacion_tipoContext.class,0);
		}
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Variable_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variable_declaracion; }
	}

	public final Variable_declaracionContext variable_declaracion() throws RecognitionException {
		Variable_declaracionContext _localctx = new Variable_declaracionContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_variable_declaracion);
		int _la;
		try {
			setState(198);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(186);
				match(T__18);
				setState(187);
				match(Identificador);
				setState(188);
				anotacion_tipo();
				setState(189);
				match(T__19);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(191);
				match(T__18);
				setState(192);
				match(Identificador);
				setState(194);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__11) {
					{
					setState(193);
					anotacion_tipo();
					}
				}

				setState(196);
				match(T__17);
				setState(197);
				expresion(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Anotacion_tipoContext extends ParserRuleContext {
		public TiposContext tipos() {
			return getRuleContext(TiposContext.class,0);
		}
		public Anotacion_tipoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_anotacion_tipo; }
	}

	public final Anotacion_tipoContext anotacion_tipo() throws RecognitionException {
		Anotacion_tipoContext _localctx = new Anotacion_tipoContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_anotacion_tipo);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(200);
			match(T__11);
			setState(201);
			tipos();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TiposContext extends ParserRuleContext {
		public TiposContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipos; }
	}

	public final TiposContext tipos() throws RecognitionException {
		TiposContext _localctx = new TiposContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_tipos);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(203);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Array_declaracionContext extends ParserRuleContext {
		public TerminalNode Identificador() { return getToken(T_swiftParser.Identificador, 0); }
		public TiposContext tipos() {
			return getRuleContext(TiposContext.class,0);
		}
		public Definicion_vectorContext definicion_vector() {
			return getRuleContext(Definicion_vectorContext.class,0);
		}
		public Array_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_array_declaracion; }
	}

	public final Array_declaracionContext array_declaracion() throws RecognitionException {
		Array_declaracionContext _localctx = new Array_declaracionContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_array_declaracion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(205);
			match(T__18);
			setState(206);
			match(Identificador);
			setState(207);
			match(T__11);
			setState(208);
			match(T__24);
			setState(209);
			tipos();
			setState(210);
			match(T__25);
			setState(211);
			definicion_vector();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Definicion_vectorContext extends ParserRuleContext {
		public Lista_expresionesContext lista_expresiones() {
			return getRuleContext(Lista_expresionesContext.class,0);
		}
		public TerminalNode Identificador() { return getToken(T_swiftParser.Identificador, 0); }
		public Definicion_vectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_definicion_vector; }
	}

	public final Definicion_vectorContext definicion_vector() throws RecognitionException {
		Definicion_vectorContext _localctx = new Definicion_vectorContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_definicion_vector);
		try {
			setState(222);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(213);
				match(T__17);
				setState(214);
				match(T__24);
				setState(215);
				lista_expresiones();
				setState(216);
				match(T__25);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(218);
				match(T__17);
				setState(219);
				match(T__24);
				setState(220);
				match(T__25);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(221);
				match(Identificador);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Lista_expresionesContext extends ParserRuleContext {
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public Lista_expresionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lista_expresiones; }
	}

	public final Lista_expresionesContext lista_expresiones() throws RecognitionException {
		Lista_expresionesContext _localctx = new Lista_expresionesContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_lista_expresiones);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(224);
			expresion(0);
			setState(229);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__26) {
				{
				{
				setState(225);
				match(T__26);
				setState(226);
				expresion(0);
				}
				}
				setState(231);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Funcion_printContext extends ParserRuleContext {
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Funcion_printContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcion_print; }
	}

	public final Funcion_printContext funcion_print() throws RecognitionException {
		Funcion_printContext _localctx = new Funcion_printContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_funcion_print);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(232);
			match(T__27);
			setState(233);
			match(T__28);
			setState(234);
			expresion(0);
			setState(235);
			match(T__29);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpresionContext extends ParserRuleContext {
		public ExpresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expresion; }
	 
		public ExpresionContext() { }
		public void copyFrom(ExpresionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Expresion_idContext extends ExpresionContext {
		public TerminalNode Identificador() { return getToken(T_swiftParser.Identificador, 0); }
		public Expresion_idContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	public static class Valor_primitivoContext extends ExpresionContext {
		public PrimitivosContext primitivos() {
			return getRuleContext(PrimitivosContext.class,0);
		}
		public Valor_primitivoContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	public static class Expresion_unarioContext extends ExpresionContext {
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Expresion_unarioContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	public static class Expresion_parenContext extends ExpresionContext {
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Expresion_parenContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	public static class Expresion_relaContext extends ExpresionContext {
		public ExpresionContext right;
		public Token op;
		public ExpresionContext left;
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public Expresion_relaContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	public static class Expresion_aritContext extends ExpresionContext {
		public ExpresionContext right;
		public Token op;
		public ExpresionContext left;
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public Expresion_aritContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	public static class Expresion_compaContext extends ExpresionContext {
		public ExpresionContext right;
		public Token op;
		public ExpresionContext left;
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public Expresion_compaContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	public static class Expresion_negaContext extends ExpresionContext {
		public Token op;
		public ExpresionContext left;
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Expresion_negaContext(ExpresionContext ctx) { copyFrom(ctx); }
	}

	public final ExpresionContext expresion() throws RecognitionException {
		return expresion(0);
	}

	private ExpresionContext expresion(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpresionContext _localctx = new ExpresionContext(_ctx, _parentState);
		ExpresionContext _prevctx = _localctx;
		int _startState = 58;
		enterRecursionRule(_localctx, 58, RULE_expresion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(248);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Int:
			case Float:
			case Caracter:
			case String:
			case Bool:
				{
				_localctx = new Valor_primitivoContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(238);
				primitivos();
				}
				break;
			case Identificador:
				{
				_localctx = new Expresion_idContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(239);
				match(Identificador);
				}
				break;
			case T__30:
				{
				_localctx = new Expresion_unarioContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(240);
				match(T__30);
				setState(241);
				expresion(7);
				}
				break;
			case T__28:
				{
				_localctx = new Expresion_parenContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(242);
				match(T__28);
				setState(243);
				expresion(0);
				setState(244);
				match(T__29);
				}
				break;
			case T__41:
				{
				_localctx = new Expresion_negaContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(246);
				((Expresion_negaContext)_localctx).op = match(T__41);
				setState(247);
				((Expresion_negaContext)_localctx).left = expresion(2);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(264);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,24,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(262);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
					case 1:
						{
						_localctx = new Expresion_aritContext(new ExpresionContext(_parentctx, _parentState));
						((Expresion_aritContext)_localctx).right = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(250);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(251);
						((Expresion_aritContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__31) | (1L << T__32) | (1L << T__33))) != 0)) ) {
							((Expresion_aritContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(252);
						((Expresion_aritContext)_localctx).left = expresion(6);
						}
						break;
					case 2:
						{
						_localctx = new Expresion_aritContext(new ExpresionContext(_parentctx, _parentState));
						((Expresion_aritContext)_localctx).right = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(253);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(254);
						((Expresion_aritContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__30 || _la==T__34) ) {
							((Expresion_aritContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(255);
						((Expresion_aritContext)_localctx).left = expresion(5);
						}
						break;
					case 3:
						{
						_localctx = new Expresion_compaContext(new ExpresionContext(_parentctx, _parentState));
						((Expresion_compaContext)_localctx).right = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(256);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(257);
						((Expresion_compaContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40))) != 0)) ) {
							((Expresion_compaContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(258);
						((Expresion_compaContext)_localctx).left = expresion(4);
						}
						break;
					case 4:
						{
						_localctx = new Expresion_relaContext(new ExpresionContext(_parentctx, _parentState));
						((Expresion_relaContext)_localctx).right = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(259);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(260);
						((Expresion_relaContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__42 || _la==T__43) ) {
							((Expresion_relaContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(261);
						((Expresion_relaContext)_localctx).left = expresion(2);
						}
						break;
					}
					} 
				}
				setState(266);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,24,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class PrimitivosContext extends ParserRuleContext {
		public PrimitivosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primitivos; }
	 
		public PrimitivosContext() { }
		public void copyFrom(PrimitivosContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Primitivo_boolContext extends PrimitivosContext {
		public TerminalNode Bool() { return getToken(T_swiftParser.Bool, 0); }
		public Primitivo_boolContext(PrimitivosContext ctx) { copyFrom(ctx); }
	}
	public static class Primitivo_stringContext extends PrimitivosContext {
		public TerminalNode String() { return getToken(T_swiftParser.String, 0); }
		public Primitivo_stringContext(PrimitivosContext ctx) { copyFrom(ctx); }
	}
	public static class Primitivo_floatContext extends PrimitivosContext {
		public TerminalNode Float() { return getToken(T_swiftParser.Float, 0); }
		public Primitivo_floatContext(PrimitivosContext ctx) { copyFrom(ctx); }
	}
	public static class Primitivo_caracterContext extends PrimitivosContext {
		public TerminalNode Caracter() { return getToken(T_swiftParser.Caracter, 0); }
		public Primitivo_caracterContext(PrimitivosContext ctx) { copyFrom(ctx); }
	}
	public static class Primitivo_intContext extends PrimitivosContext {
		public TerminalNode Int() { return getToken(T_swiftParser.Int, 0); }
		public Primitivo_intContext(PrimitivosContext ctx) { copyFrom(ctx); }
	}

	public final PrimitivosContext primitivos() throws RecognitionException {
		PrimitivosContext _localctx = new PrimitivosContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_primitivos);
		try {
			setState(272);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Caracter:
				_localctx = new Primitivo_caracterContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(267);
				match(Caracter);
				}
				break;
			case Int:
				_localctx = new Primitivo_intContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(268);
				match(Int);
				}
				break;
			case Float:
				_localctx = new Primitivo_floatContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(269);
				match(Float);
				}
				break;
			case String:
				_localctx = new Primitivo_stringContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(270);
				match(String);
				}
				break;
			case Bool:
				_localctx = new Primitivo_boolContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(271);
				match(Bool);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 29:
			return expresion_sempred((ExpresionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expresion_sempred(ExpresionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 5);
		case 1:
			return precpred(_ctx, 4);
		case 2:
			return precpred(_ctx, 3);
		case 3:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\67\u0115\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \3\2"+
		"\3\2\3\2\3\3\7\3E\n\3\f\3\16\3H\13\3\3\4\3\4\5\4L\n\4\3\4\3\4\5\4P\n\4"+
		"\3\4\3\4\5\4T\n\4\3\4\3\4\5\4X\n\4\3\4\3\4\5\4\\\n\4\5\4^\n\4\3\5\3\5"+
		"\5\5b\n\5\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\t"+
		"\3\t\3\t\5\tu\n\t\3\n\3\n\3\n\3\n\5\n{\n\n\3\13\3\13\3\13\3\13\5\13\u0081"+
		"\n\13\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\5\r\u008c\n\r\3\r\3\r\3\16\3"+
		"\16\5\16\u0092\n\16\3\17\3\17\3\17\3\17\3\17\3\17\5\17\u009a\n\17\3\20"+
		"\3\20\3\20\3\20\3\21\3\21\3\21\3\22\3\22\3\22\5\22\u00a6\n\22\3\23\3\23"+
		"\3\24\3\24\3\25\3\25\5\25\u00ae\n\25\3\26\3\26\3\26\5\26\u00b3\n\26\3"+
		"\27\3\27\3\27\5\27\u00b8\n\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30"+
		"\3\30\3\30\3\30\5\30\u00c5\n\30\3\30\3\30\5\30\u00c9\n\30\3\31\3\31\3"+
		"\31\3\32\3\32\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3"+
		"\34\3\34\3\34\3\34\3\34\3\34\5\34\u00e1\n\34\3\35\3\35\3\35\7\35\u00e6"+
		"\n\35\f\35\16\35\u00e9\13\35\3\36\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3"+
		"\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\5\37\u00fb\n\37\3\37\3\37\3\37"+
		"\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\7\37\u0109\n\37\f\37\16"+
		"\37\u010c\13\37\3 \3 \3 \3 \3 \5 \u0113\n \3 \2\3<!\2\4\6\b\n\f\16\20"+
		"\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\668:<>\2\7\3\2\27\32\3\2\"$\4"+
		"\2!!%%\3\2&+\3\2-.\2\u011e\2@\3\2\2\2\4F\3\2\2\2\6]\3\2\2\2\ba\3\2\2\2"+
		"\nc\3\2\2\2\fg\3\2\2\2\16m\3\2\2\2\20t\3\2\2\2\22v\3\2\2\2\24\u0080\3"+
		"\2\2\2\26\u0082\3\2\2\2\30\u0087\3\2\2\2\32\u008f\3\2\2\2\34\u0099\3\2"+
		"\2\2\36\u009b\3\2\2\2 \u009f\3\2\2\2\"\u00a5\3\2\2\2$\u00a7\3\2\2\2&\u00a9"+
		"\3\2\2\2(\u00ab\3\2\2\2*\u00b2\3\2\2\2,\u00b4\3\2\2\2.\u00c8\3\2\2\2\60"+
		"\u00ca\3\2\2\2\62\u00cd\3\2\2\2\64\u00cf\3\2\2\2\66\u00e0\3\2\2\28\u00e2"+
		"\3\2\2\2:\u00ea\3\2\2\2<\u00fa\3\2\2\2>\u0112\3\2\2\2@A\5\4\3\2AB\7\2"+
		"\2\3B\3\3\2\2\2CE\5\6\4\2DC\3\2\2\2EH\3\2\2\2FD\3\2\2\2FG\3\2\2\2G\5\3"+
		"\2\2\2HF\3\2\2\2IK\5*\26\2JL\7\3\2\2KJ\3\2\2\2KL\3\2\2\2L^\3\2\2\2MO\5"+
		"\b\5\2NP\7\3\2\2ON\3\2\2\2OP\3\2\2\2P^\3\2\2\2QS\5\20\t\2RT\7\3\2\2SR"+
		"\3\2\2\2ST\3\2\2\2T^\3\2\2\2UW\5\"\22\2VX\7\3\2\2WV\3\2\2\2WX\3\2\2\2"+
		"X^\3\2\2\2Y[\5:\36\2Z\\\7\3\2\2[Z\3\2\2\2[\\\3\2\2\2\\^\3\2\2\2]I\3\2"+
		"\2\2]M\3\2\2\2]Q\3\2\2\2]U\3\2\2\2]Y\3\2\2\2^\7\3\2\2\2_b\5\f\7\2`b\5"+
		"\16\b\2a_\3\2\2\2a`\3\2\2\2b\t\3\2\2\2cd\7\4\2\2de\5\4\3\2ef\7\5\2\2f"+
		"\13\3\2\2\2gh\7\6\2\2hi\7\67\2\2ij\7\7\2\2jk\5<\37\2kl\5\n\6\2l\r\3\2"+
		"\2\2mn\7\b\2\2no\5<\37\2op\5\n\6\2p\17\3\2\2\2qu\5\22\n\2ru\5\26\f\2s"+
		"u\5\30\r\2tq\3\2\2\2tr\3\2\2\2ts\3\2\2\2u\21\3\2\2\2vw\7\t\2\2wx\5<\37"+
		"\2xz\5\n\6\2y{\5\24\13\2zy\3\2\2\2z{\3\2\2\2{\23\3\2\2\2|}\7\n\2\2}\u0081"+
		"\5\n\6\2~\177\7\n\2\2\177\u0081\5\22\n\2\u0080|\3\2\2\2\u0080~\3\2\2\2"+
		"\u0081\25\3\2\2\2\u0082\u0083\7\13\2\2\u0083\u0084\5<\37\2\u0084\u0085"+
		"\7\n\2\2\u0085\u0086\5\n\6\2\u0086\27\3\2\2\2\u0087\u0088\7\f\2\2\u0088"+
		"\u0089\5<\37\2\u0089\u008b\7\4\2\2\u008a\u008c\5\32\16\2\u008b\u008a\3"+
		"\2\2\2\u008b\u008c\3\2\2\2\u008c\u008d\3\2\2\2\u008d\u008e\7\5\2\2\u008e"+
		"\31\3\2\2\2\u008f\u0091\5\34\17\2\u0090\u0092\5\32\16\2\u0091\u0090\3"+
		"\2\2\2\u0091\u0092\3\2\2\2\u0092\33\3\2\2\2\u0093\u0094\5\36\20\2\u0094"+
		"\u0095\5\4\3\2\u0095\u009a\3\2\2\2\u0096\u0097\5 \21\2\u0097\u0098\5\4"+
		"\3\2\u0098\u009a\3\2\2\2\u0099\u0093\3\2\2\2\u0099\u0096\3\2\2\2\u009a"+
		"\35\3\2\2\2\u009b\u009c\7\r\2\2\u009c\u009d\5<\37\2\u009d\u009e\7\16\2"+
		"\2\u009e\37\3\2\2\2\u009f\u00a0\7\17\2\2\u00a0\u00a1\7\16\2\2\u00a1!\3"+
		"\2\2\2\u00a2\u00a6\5$\23\2\u00a3\u00a6\5&\24\2\u00a4\u00a6\5(\25\2\u00a5"+
		"\u00a2\3\2\2\2\u00a5\u00a3\3\2\2\2\u00a5\u00a4\3\2\2\2\u00a6#\3\2\2\2"+
		"\u00a7\u00a8\7\20\2\2\u00a8%\3\2\2\2\u00a9\u00aa\7\21\2\2\u00aa\'\3\2"+
		"\2\2\u00ab\u00ad\7\22\2\2\u00ac\u00ae\5<\37\2\u00ad\u00ac\3\2\2\2\u00ad"+
		"\u00ae\3\2\2\2\u00ae)\3\2\2\2\u00af\u00b3\5,\27\2\u00b0\u00b3\5.\30\2"+
		"\u00b1\u00b3\5\64\33\2\u00b2\u00af\3\2\2\2\u00b2\u00b0\3\2\2\2\u00b2\u00b1"+
		"\3\2\2\2\u00b3+\3\2\2\2\u00b4\u00b5\7\23\2\2\u00b5\u00b7\7\67\2\2\u00b6"+
		"\u00b8\5\60\31\2\u00b7\u00b6\3\2\2\2\u00b7\u00b8\3\2\2\2\u00b8\u00b9\3"+
		"\2\2\2\u00b9\u00ba\7\24\2\2\u00ba\u00bb\5<\37\2\u00bb-\3\2\2\2\u00bc\u00bd"+
		"\7\25\2\2\u00bd\u00be\7\67\2\2\u00be\u00bf\5\60\31\2\u00bf\u00c0\7\26"+
		"\2\2\u00c0\u00c9\3\2\2\2\u00c1\u00c2\7\25\2\2\u00c2\u00c4\7\67\2\2\u00c3"+
		"\u00c5\5\60\31\2\u00c4\u00c3\3\2\2\2\u00c4\u00c5\3\2\2\2\u00c5\u00c6\3"+
		"\2\2\2\u00c6\u00c7\7\24\2\2\u00c7\u00c9\5<\37\2\u00c8\u00bc\3\2\2\2\u00c8"+
		"\u00c1\3\2\2\2\u00c9/\3\2\2\2\u00ca\u00cb\7\16\2\2\u00cb\u00cc\5\62\32"+
		"\2\u00cc\61\3\2\2\2\u00cd\u00ce\t\2\2\2\u00ce\63\3\2\2\2\u00cf\u00d0\7"+
		"\25\2\2\u00d0\u00d1\7\67\2\2\u00d1\u00d2\7\16\2\2\u00d2\u00d3\7\33\2\2"+
		"\u00d3\u00d4\5\62\32\2\u00d4\u00d5\7\34\2\2\u00d5\u00d6\5\66\34\2\u00d6"+
		"\65\3\2\2\2\u00d7\u00d8\7\24\2\2\u00d8\u00d9\7\33\2\2\u00d9\u00da\58\35"+
		"\2\u00da\u00db\7\34\2\2\u00db\u00e1\3\2\2\2\u00dc\u00dd\7\24\2\2\u00dd"+
		"\u00de\7\33\2\2\u00de\u00e1\7\34\2\2\u00df\u00e1\7\67\2\2\u00e0\u00d7"+
		"\3\2\2\2\u00e0\u00dc\3\2\2\2\u00e0\u00df\3\2\2\2\u00e1\67\3\2\2\2\u00e2"+
		"\u00e7\5<\37\2\u00e3\u00e4\7\35\2\2\u00e4\u00e6\5<\37\2\u00e5\u00e3\3"+
		"\2\2\2\u00e6\u00e9\3\2\2\2\u00e7\u00e5\3\2\2\2\u00e7\u00e8\3\2\2\2\u00e8"+
		"9\3\2\2\2\u00e9\u00e7\3\2\2\2\u00ea\u00eb\7\36\2\2\u00eb\u00ec\7\37\2"+
		"\2\u00ec\u00ed\5<\37\2\u00ed\u00ee\7 \2\2\u00ee;\3\2\2\2\u00ef\u00f0\b"+
		"\37\1\2\u00f0\u00fb\5> \2\u00f1\u00fb\7\67\2\2\u00f2\u00f3\7!\2\2\u00f3"+
		"\u00fb\5<\37\t\u00f4\u00f5\7\37\2\2\u00f5\u00f6\5<\37\2\u00f6\u00f7\7"+
		" \2\2\u00f7\u00fb\3\2\2\2\u00f8\u00f9\7,\2\2\u00f9\u00fb\5<\37\4\u00fa"+
		"\u00ef\3\2\2\2\u00fa\u00f1\3\2\2\2\u00fa\u00f2\3\2\2\2\u00fa\u00f4\3\2"+
		"\2\2\u00fa\u00f8\3\2\2\2\u00fb\u010a\3\2\2\2\u00fc\u00fd\f\7\2\2\u00fd"+
		"\u00fe\t\3\2\2\u00fe\u0109\5<\37\b\u00ff\u0100\f\6\2\2\u0100\u0101\t\4"+
		"\2\2\u0101\u0109\5<\37\7\u0102\u0103\f\5\2\2\u0103\u0104\t\5\2\2\u0104"+
		"\u0109\5<\37\6\u0105\u0106\f\3\2\2\u0106\u0107\t\6\2\2\u0107\u0109\5<"+
		"\37\4\u0108\u00fc\3\2\2\2\u0108\u00ff\3\2\2\2\u0108\u0102\3\2\2\2\u0108"+
		"\u0105\3\2\2\2\u0109\u010c\3\2\2\2\u010a\u0108\3\2\2\2\u010a\u010b\3\2"+
		"\2\2\u010b=\3\2\2\2\u010c\u010a\3\2\2\2\u010d\u0113\7\64\2\2\u010e\u0113"+
		"\7\62\2\2\u010f\u0113\7\63\2\2\u0110\u0113\7\65\2\2\u0111\u0113\7\66\2"+
		"\2\u0112\u010d\3\2\2\2\u0112\u010e\3\2\2\2\u0112\u010f\3\2\2\2\u0112\u0110"+
		"\3\2\2\2\u0112\u0111\3\2\2\2\u0113?\3\2\2\2\34FKOSW[]atz\u0080\u008b\u0091"+
		"\u0099\u00a5\u00ad\u00b2\u00b7\u00c4\u00c8\u00e0\u00e7\u00fa\u0108\u010a"+
		"\u0112";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}