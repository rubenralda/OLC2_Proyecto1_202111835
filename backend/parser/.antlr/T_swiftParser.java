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
		T__38=39, T__39=40, T__40=41, T__41=42, T__42=43, T__43=44, T__44=45, 
		T__45=46, T__46=47, T__47=48, WS=49, Block_comment=50, Line_comment=51, 
		Int=52, Float=53, Caracter=54, String=55, Bool=56, Identificador=57;
	public static final int
		RULE_inicio = 0, RULE_instrucciones = 1, RULE_instruccion = 2, RULE_loop_statement = 3, 
		RULE_code_block = 4, RULE_for_in_statement = 5, RULE_rango = 6, RULE_while_statement = 7, 
		RULE_branch_statement = 8, RULE_if_statement = 9, RULE_guard_statement = 10, 
		RULE_switch_statement = 11, RULE_switch_case = 12, RULE_default_case = 13, 
		RULE_control_transfer_statement = 14, RULE_break_statement = 15, RULE_continue_statement = 16, 
		RULE_return_statement = 17, RULE_declaracion = 18, RULE_constant_declaracion = 19, 
		RULE_variable_declaracion = 20, RULE_anotacion_tipo = 21, RULE_tipos = 22, 
		RULE_array_declaracion = 23, RULE_definicion_vector = 24, RULE_lista_expresiones = 25, 
		RULE_funcion_print = 26, RULE_asignacion = 27, RULE_expresion = 28, RULE_primitivos = 29;
	private static String[] makeRuleNames() {
		return new String[] {
			"inicio", "instrucciones", "instruccion", "loop_statement", "code_block", 
			"for_in_statement", "rango", "while_statement", "branch_statement", "if_statement", 
			"guard_statement", "switch_statement", "switch_case", "default_case", 
			"control_transfer_statement", "break_statement", "continue_statement", 
			"return_statement", "declaracion", "constant_declaracion", "variable_declaracion", 
			"anotacion_tipo", "tipos", "array_declaracion", "definicion_vector", 
			"lista_expresiones", "funcion_print", "asignacion", "expresion", "primitivos"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "';'", "'{'", "'}'", "'for'", "'in'", "'...'", "'while'", "'if'", 
			"'else'", "'guard'", "'continue'", "'break'", "'return'", "'switch'", 
			"'case'", "':'", "'default'", "'let'", "'='", "'var'", "'?'", "'String'", 
			"'Int'", "'Float'", "'Bool'", "'Character'", "'['", "']'", "','", "'print'", 
			"'('", "')'", "'+='", "'-='", "'!'", "'-'", "'/'", "'%'", "'*'", "'+'", 
			"'<'", "'<='", "'>='", "'>'", "'=='", "'!='", "'&&'", "'||'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, "WS", "Block_comment", "Line_comment", "Int", "Float", "Caracter", 
			"String", "Bool", "Identificador"
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
			setState(60);
			instrucciones();
			setState(61);
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
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(66);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(63);
					instruccion();
					}
					} 
				}
				setState(68);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
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
		public AsignacionContext asignacion() {
			return getRuleContext(AsignacionContext.class,0);
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
			setState(93);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__17:
			case T__19:
				enterOuterAlt(_localctx, 1);
				{
				setState(69);
				declaracion();
				setState(71);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__0) {
					{
					setState(70);
					match(T__0);
					}
				}

				}
				break;
			case T__3:
			case T__6:
				enterOuterAlt(_localctx, 2);
				{
				setState(73);
				loop_statement();
				setState(75);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__0) {
					{
					setState(74);
					match(T__0);
					}
				}

				}
				break;
			case T__7:
			case T__9:
			case T__13:
				enterOuterAlt(_localctx, 3);
				{
				setState(77);
				branch_statement();
				setState(79);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__0) {
					{
					setState(78);
					match(T__0);
					}
				}

				}
				break;
			case T__10:
			case T__11:
			case T__12:
				enterOuterAlt(_localctx, 4);
				{
				setState(81);
				control_transfer_statement();
				setState(83);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__0) {
					{
					setState(82);
					match(T__0);
					}
				}

				}
				break;
			case T__29:
				enterOuterAlt(_localctx, 5);
				{
				setState(85);
				funcion_print();
				setState(87);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__0) {
					{
					setState(86);
					match(T__0);
					}
				}

				}
				break;
			case Identificador:
				enterOuterAlt(_localctx, 6);
				{
				setState(89);
				asignacion();
				setState(91);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__0) {
					{
					setState(90);
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
			setState(97);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__3:
				enterOuterAlt(_localctx, 1);
				{
				setState(95);
				for_in_statement();
				}
				break;
			case T__6:
				enterOuterAlt(_localctx, 2);
				{
				setState(96);
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
			setState(99);
			match(T__1);
			setState(100);
			instrucciones();
			setState(101);
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
		public Code_blockContext code_block() {
			return getRuleContext(Code_blockContext.class,0);
		}
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public RangoContext rango() {
			return getRuleContext(RangoContext.class,0);
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
			setState(103);
			match(T__3);
			setState(104);
			match(Identificador);
			setState(105);
			match(T__4);
			setState(108);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				{
				setState(106);
				expresion(0);
				}
				break;
			case 2:
				{
				setState(107);
				rango();
				}
				break;
			}
			setState(110);
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

	public static class RangoContext extends ParserRuleContext {
		public List<TerminalNode> Int() { return getTokens(T_swiftParser.Int); }
		public TerminalNode Int(int i) {
			return getToken(T_swiftParser.Int, i);
		}
		public RangoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rango; }
	}

	public final RangoContext rango() throws RecognitionException {
		RangoContext _localctx = new RangoContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_rango);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(112);
			match(Int);
			setState(113);
			match(T__5);
			setState(114);
			match(Int);
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
		enterRule(_localctx, 14, RULE_while_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(116);
			match(T__6);
			setState(117);
			expresion(0);
			setState(118);
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
		public Branch_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_branch_statement; }
	 
		public Branch_statementContext() { }
		public void copyFrom(Branch_statementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Declarar_guardContext extends Branch_statementContext {
		public Guard_statementContext guard_statement() {
			return getRuleContext(Guard_statementContext.class,0);
		}
		public Declarar_guardContext(Branch_statementContext ctx) { copyFrom(ctx); }
	}
	public static class Declarar_ifContext extends Branch_statementContext {
		public If_statementContext if_statement() {
			return getRuleContext(If_statementContext.class,0);
		}
		public Declarar_ifContext(Branch_statementContext ctx) { copyFrom(ctx); }
	}
	public static class Declarar_switchContext extends Branch_statementContext {
		public Switch_statementContext switch_statement() {
			return getRuleContext(Switch_statementContext.class,0);
		}
		public Declarar_switchContext(Branch_statementContext ctx) { copyFrom(ctx); }
	}

	public final Branch_statementContext branch_statement() throws RecognitionException {
		Branch_statementContext _localctx = new Branch_statementContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_branch_statement);
		try {
			setState(123);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__7:
				_localctx = new Declarar_ifContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(120);
				if_statement();
				}
				break;
			case T__9:
				_localctx = new Declarar_guardContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(121);
				guard_statement();
				}
				break;
			case T__13:
				_localctx = new Declarar_switchContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(122);
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
		public If_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_statement; }
	 
		public If_statementContext() { }
		public void copyFrom(If_statementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Else_finalContext extends If_statementContext {
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public List<Code_blockContext> code_block() {
			return getRuleContexts(Code_blockContext.class);
		}
		public Code_blockContext code_block(int i) {
			return getRuleContext(Code_blockContext.class,i);
		}
		public Else_finalContext(If_statementContext ctx) { copyFrom(ctx); }
	}
	public static class If_simpleContext extends If_statementContext {
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Code_blockContext code_block() {
			return getRuleContext(Code_blockContext.class,0);
		}
		public If_simpleContext(If_statementContext ctx) { copyFrom(ctx); }
	}
	public static class Siguiente_ifContext extends If_statementContext {
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Code_blockContext code_block() {
			return getRuleContext(Code_blockContext.class,0);
		}
		public If_statementContext if_statement() {
			return getRuleContext(If_statementContext.class,0);
		}
		public Siguiente_ifContext(If_statementContext ctx) { copyFrom(ctx); }
	}

	public final If_statementContext if_statement() throws RecognitionException {
		If_statementContext _localctx = new If_statementContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_if_statement);
		try {
			setState(141);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				_localctx = new If_simpleContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(125);
				match(T__7);
				setState(126);
				expresion(0);
				setState(127);
				code_block();
				}
				break;
			case 2:
				_localctx = new Else_finalContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(129);
				match(T__7);
				setState(130);
				expresion(0);
				setState(131);
				code_block();
				setState(132);
				match(T__8);
				setState(133);
				code_block();
				}
				break;
			case 3:
				_localctx = new Siguiente_ifContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(135);
				match(T__7);
				setState(136);
				expresion(0);
				setState(137);
				code_block();
				setState(138);
				match(T__8);
				setState(139);
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
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public Guard_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_guard_statement; }
	}

	public final Guard_statementContext guard_statement() throws RecognitionException {
		Guard_statementContext _localctx = new Guard_statementContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_guard_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(143);
			match(T__9);
			setState(144);
			expresion(0);
			setState(145);
			match(T__8);
			setState(146);
			match(T__1);
			setState(147);
			instrucciones();
			setState(148);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__10) | (1L << T__11) | (1L << T__12))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(149);
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

	public static class Switch_statementContext extends ParserRuleContext {
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public List<Switch_caseContext> switch_case() {
			return getRuleContexts(Switch_caseContext.class);
		}
		public Switch_caseContext switch_case(int i) {
			return getRuleContext(Switch_caseContext.class,i);
		}
		public Default_caseContext default_case() {
			return getRuleContext(Default_caseContext.class,0);
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
			setState(151);
			match(T__13);
			setState(152);
			expresion(0);
			setState(153);
			match(T__1);
			setState(157);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__14) {
				{
				{
				setState(154);
				switch_case();
				}
				}
				setState(159);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(161);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__16) {
				{
				setState(160);
				default_case();
				}
			}

			setState(163);
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

	public static class Switch_caseContext extends ParserRuleContext {
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public Switch_caseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switch_case; }
	}

	public final Switch_caseContext switch_case() throws RecognitionException {
		Switch_caseContext _localctx = new Switch_caseContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_switch_case);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(165);
			match(T__14);
			setState(166);
			expresion(0);
			setState(167);
			match(T__15);
			setState(168);
			instrucciones();
			setState(170);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__11) {
				{
				setState(169);
				match(T__11);
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

	public static class Default_caseContext extends ParserRuleContext {
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public Default_caseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_default_case; }
	}

	public final Default_caseContext default_case() throws RecognitionException {
		Default_caseContext _localctx = new Default_caseContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_default_case);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(172);
			match(T__16);
			setState(173);
			match(T__15);
			setState(174);
			instrucciones();
			setState(176);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__11) {
				{
				setState(175);
				match(T__11);
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
		enterRule(_localctx, 28, RULE_control_transfer_statement);
		try {
			setState(181);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__11:
				enterOuterAlt(_localctx, 1);
				{
				setState(178);
				break_statement();
				}
				break;
			case T__10:
				enterOuterAlt(_localctx, 2);
				{
				setState(179);
				continue_statement();
				}
				break;
			case T__12:
				enterOuterAlt(_localctx, 3);
				{
				setState(180);
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
		enterRule(_localctx, 30, RULE_break_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(183);
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

	public static class Continue_statementContext extends ParserRuleContext {
		public Continue_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_continue_statement; }
	}

	public final Continue_statementContext continue_statement() throws RecognitionException {
		Continue_statementContext _localctx = new Continue_statementContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_continue_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(185);
			match(T__10);
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
		enterRule(_localctx, 34, RULE_return_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(187);
			match(T__12);
			setState(189);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				{
				setState(188);
				expresion(0);
				}
				break;
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
		enterRule(_localctx, 36, RULE_declaracion);
		try {
			setState(194);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(191);
				constant_declaracion();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(192);
				variable_declaracion();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(193);
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
		enterRule(_localctx, 38, RULE_constant_declaracion);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(196);
			match(T__17);
			setState(197);
			match(Identificador);
			setState(199);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__15) {
				{
				setState(198);
				anotacion_tipo();
				}
			}

			setState(201);
			match(T__18);
			setState(202);
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
		enterRule(_localctx, 40, RULE_variable_declaracion);
		int _la;
		try {
			setState(216);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(204);
				match(T__19);
				setState(205);
				match(Identificador);
				setState(206);
				anotacion_tipo();
				setState(207);
				match(T__20);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(209);
				match(T__19);
				setState(210);
				match(Identificador);
				setState(212);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__15) {
					{
					setState(211);
					anotacion_tipo();
					}
				}

				setState(214);
				match(T__18);
				setState(215);
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
		enterRule(_localctx, 42, RULE_anotacion_tipo);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(218);
			match(T__15);
			setState(219);
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
		enterRule(_localctx, 44, RULE_tipos);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(221);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25))) != 0)) ) {
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
		enterRule(_localctx, 46, RULE_array_declaracion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(223);
			match(T__19);
			setState(224);
			match(Identificador);
			setState(225);
			match(T__15);
			setState(226);
			match(T__26);
			setState(227);
			tipos();
			setState(228);
			match(T__27);
			setState(229);
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
		enterRule(_localctx, 48, RULE_definicion_vector);
		try {
			setState(241);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(231);
				match(T__18);
				setState(232);
				match(T__26);
				setState(233);
				lista_expresiones();
				setState(234);
				match(T__27);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(236);
				match(T__18);
				setState(237);
				match(T__26);
				setState(238);
				match(T__27);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(239);
				match(T__18);
				setState(240);
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
		enterRule(_localctx, 50, RULE_lista_expresiones);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(243);
			expresion(0);
			setState(248);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__28) {
				{
				{
				setState(244);
				match(T__28);
				setState(245);
				expresion(0);
				}
				}
				setState(250);
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
		enterRule(_localctx, 52, RULE_funcion_print);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(251);
			match(T__29);
			setState(252);
			match(T__30);
			setState(253);
			expresion(0);
			setState(254);
			match(T__31);
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

	public static class AsignacionContext extends ParserRuleContext {
		public AsignacionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignacion; }
	 
		public AsignacionContext() { }
		public void copyFrom(AsignacionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Asignacion_normalContext extends AsignacionContext {
		public TerminalNode Identificador() { return getToken(T_swiftParser.Identificador, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Asignacion_normalContext(AsignacionContext ctx) { copyFrom(ctx); }
	}
	public static class Asignacion_decrementoContext extends AsignacionContext {
		public TerminalNode Identificador() { return getToken(T_swiftParser.Identificador, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Asignacion_decrementoContext(AsignacionContext ctx) { copyFrom(ctx); }
	}
	public static class Asignacion_vectorContext extends AsignacionContext {
		public TerminalNode Identificador() { return getToken(T_swiftParser.Identificador, 0); }
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public Asignacion_vectorContext(AsignacionContext ctx) { copyFrom(ctx); }
	}
	public static class Asignacion_incremento_vectorContext extends AsignacionContext {
		public TerminalNode Identificador() { return getToken(T_swiftParser.Identificador, 0); }
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public Asignacion_incremento_vectorContext(AsignacionContext ctx) { copyFrom(ctx); }
	}
	public static class Asignacion_incrementoContext extends AsignacionContext {
		public TerminalNode Identificador() { return getToken(T_swiftParser.Identificador, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Asignacion_incrementoContext(AsignacionContext ctx) { copyFrom(ctx); }
	}
	public static class Asignacion_decremento_vectorContext extends AsignacionContext {
		public TerminalNode Identificador() { return getToken(T_swiftParser.Identificador, 0); }
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public Asignacion_decremento_vectorContext(AsignacionContext ctx) { copyFrom(ctx); }
	}

	public final AsignacionContext asignacion() throws RecognitionException {
		AsignacionContext _localctx = new AsignacionContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_asignacion);
		try {
			setState(286);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
			case 1:
				_localctx = new Asignacion_normalContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(256);
				match(Identificador);
				setState(257);
				match(T__18);
				setState(258);
				expresion(0);
				}
				break;
			case 2:
				_localctx = new Asignacion_incrementoContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(259);
				match(Identificador);
				setState(260);
				match(T__32);
				setState(261);
				expresion(0);
				}
				break;
			case 3:
				_localctx = new Asignacion_decrementoContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(262);
				match(Identificador);
				setState(263);
				match(T__33);
				setState(264);
				expresion(0);
				}
				break;
			case 4:
				_localctx = new Asignacion_vectorContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(265);
				match(Identificador);
				setState(266);
				match(T__26);
				setState(267);
				expresion(0);
				setState(268);
				match(T__27);
				setState(269);
				match(T__18);
				setState(270);
				expresion(0);
				}
				break;
			case 5:
				_localctx = new Asignacion_incremento_vectorContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(272);
				match(Identificador);
				setState(273);
				match(T__26);
				setState(274);
				expresion(0);
				setState(275);
				match(T__27);
				setState(276);
				match(T__32);
				setState(277);
				expresion(0);
				}
				break;
			case 6:
				_localctx = new Asignacion_decremento_vectorContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(279);
				match(Identificador);
				setState(280);
				match(T__26);
				setState(281);
				expresion(0);
				setState(282);
				match(T__27);
				setState(283);
				match(T__33);
				setState(284);
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
		public Token op;
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public Expresion_relaContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	public static class Expresion_aritContext extends ExpresionContext {
		public Token op;
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public Expresion_aritContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	public static class Expresion_compaContext extends ExpresionContext {
		public Token op;
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public Expresion_compaContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	public static class Expresion_vectorContext extends ExpresionContext {
		public TerminalNode Identificador() { return getToken(T_swiftParser.Identificador, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Expresion_vectorContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	public static class Expresion_negaContext extends ExpresionContext {
		public Token op;
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
		int _startState = 56;
		enterRecursionRule(_localctx, 56, RULE_expresion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(304);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				{
				_localctx = new Valor_primitivoContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(289);
				primitivos();
				}
				break;
			case 2:
				{
				_localctx = new Expresion_vectorContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(290);
				match(Identificador);
				setState(291);
				match(T__26);
				setState(292);
				expresion(0);
				setState(293);
				match(T__27);
				}
				break;
			case 3:
				{
				_localctx = new Expresion_idContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(295);
				match(Identificador);
				}
				break;
			case 4:
				{
				_localctx = new Expresion_parenContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(296);
				match(T__30);
				setState(297);
				expresion(0);
				setState(298);
				match(T__31);
				}
				break;
			case 5:
				{
				_localctx = new Expresion_negaContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(300);
				((Expresion_negaContext)_localctx).op = match(T__34);
				setState(301);
				expresion(6);
				}
				break;
			case 6:
				{
				_localctx = new Expresion_unarioContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(302);
				match(T__35);
				setState(303);
				expresion(5);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(320);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,27,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(318);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
					case 1:
						{
						_localctx = new Expresion_aritContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(306);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(307);
						((Expresion_aritContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__36) | (1L << T__37) | (1L << T__38))) != 0)) ) {
							((Expresion_aritContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(308);
						expresion(5);
						}
						break;
					case 2:
						{
						_localctx = new Expresion_aritContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(309);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(310);
						((Expresion_aritContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__35 || _la==T__39) ) {
							((Expresion_aritContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(311);
						expresion(4);
						}
						break;
					case 3:
						{
						_localctx = new Expresion_compaContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(312);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(313);
						((Expresion_compaContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << T__43) | (1L << T__44) | (1L << T__45))) != 0)) ) {
							((Expresion_compaContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(314);
						expresion(3);
						}
						break;
					case 4:
						{
						_localctx = new Expresion_relaContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(315);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(316);
						((Expresion_relaContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__46 || _la==T__47) ) {
							((Expresion_relaContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(317);
						expresion(2);
						}
						break;
					}
					} 
				}
				setState(322);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,27,_ctx);
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
		enterRule(_localctx, 58, RULE_primitivos);
		try {
			setState(328);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Caracter:
				_localctx = new Primitivo_caracterContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(323);
				match(Caracter);
				}
				break;
			case Int:
				_localctx = new Primitivo_intContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(324);
				match(Int);
				}
				break;
			case Float:
				_localctx = new Primitivo_floatContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(325);
				match(Float);
				}
				break;
			case String:
				_localctx = new Primitivo_stringContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(326);
				match(String);
				}
				break;
			case Bool:
				_localctx = new Primitivo_boolContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(327);
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
		case 28:
			return expresion_sempred((ExpresionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expresion_sempred(ExpresionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 4);
		case 1:
			return precpred(_ctx, 3);
		case 2:
			return precpred(_ctx, 2);
		case 3:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3;\u014d\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\3\2\3\2\3"+
		"\2\3\3\7\3C\n\3\f\3\16\3F\13\3\3\4\3\4\5\4J\n\4\3\4\3\4\5\4N\n\4\3\4\3"+
		"\4\5\4R\n\4\3\4\3\4\5\4V\n\4\3\4\3\4\5\4Z\n\4\3\4\3\4\5\4^\n\4\5\4`\n"+
		"\4\3\5\3\5\5\5d\n\5\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\5\7o\n\7\3\7\3"+
		"\7\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\n\3\n\3\n\5\n~\n\n\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\5\13"+
		"\u0090\n\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\7\r\u009e"+
		"\n\r\f\r\16\r\u00a1\13\r\3\r\5\r\u00a4\n\r\3\r\3\r\3\16\3\16\3\16\3\16"+
		"\3\16\5\16\u00ad\n\16\3\17\3\17\3\17\3\17\5\17\u00b3\n\17\3\20\3\20\3"+
		"\20\5\20\u00b8\n\20\3\21\3\21\3\22\3\22\3\23\3\23\5\23\u00c0\n\23\3\24"+
		"\3\24\3\24\5\24\u00c5\n\24\3\25\3\25\3\25\5\25\u00ca\n\25\3\25\3\25\3"+
		"\25\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\5\26\u00d7\n\26\3\26\3\26"+
		"\5\26\u00db\n\26\3\27\3\27\3\27\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31"+
		"\3\31\3\31\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\5\32\u00f4"+
		"\n\32\3\33\3\33\3\33\7\33\u00f9\n\33\f\33\16\33\u00fc\13\33\3\34\3\34"+
		"\3\34\3\34\3\34\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35"+
		"\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35"+
		"\3\35\3\35\3\35\3\35\3\35\5\35\u0121\n\35\3\36\3\36\3\36\3\36\3\36\3\36"+
		"\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\36\5\36\u0133\n\36\3\36"+
		"\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\36\7\36\u0141\n\36"+
		"\f\36\16\36\u0144\13\36\3\37\3\37\3\37\3\37\3\37\5\37\u014b\n\37\3\37"+
		"\2\3: \2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\668:<"+
		"\2\b\3\2\r\17\3\2\30\34\3\2\')\4\2&&**\3\2+\60\3\2\61\62\2\u0161\2>\3"+
		"\2\2\2\4D\3\2\2\2\6_\3\2\2\2\bc\3\2\2\2\ne\3\2\2\2\fi\3\2\2\2\16r\3\2"+
		"\2\2\20v\3\2\2\2\22}\3\2\2\2\24\u008f\3\2\2\2\26\u0091\3\2\2\2\30\u0099"+
		"\3\2\2\2\32\u00a7\3\2\2\2\34\u00ae\3\2\2\2\36\u00b7\3\2\2\2 \u00b9\3\2"+
		"\2\2\"\u00bb\3\2\2\2$\u00bd\3\2\2\2&\u00c4\3\2\2\2(\u00c6\3\2\2\2*\u00da"+
		"\3\2\2\2,\u00dc\3\2\2\2.\u00df\3\2\2\2\60\u00e1\3\2\2\2\62\u00f3\3\2\2"+
		"\2\64\u00f5\3\2\2\2\66\u00fd\3\2\2\28\u0120\3\2\2\2:\u0132\3\2\2\2<\u014a"+
		"\3\2\2\2>?\5\4\3\2?@\7\2\2\3@\3\3\2\2\2AC\5\6\4\2BA\3\2\2\2CF\3\2\2\2"+
		"DB\3\2\2\2DE\3\2\2\2E\5\3\2\2\2FD\3\2\2\2GI\5&\24\2HJ\7\3\2\2IH\3\2\2"+
		"\2IJ\3\2\2\2J`\3\2\2\2KM\5\b\5\2LN\7\3\2\2ML\3\2\2\2MN\3\2\2\2N`\3\2\2"+
		"\2OQ\5\22\n\2PR\7\3\2\2QP\3\2\2\2QR\3\2\2\2R`\3\2\2\2SU\5\36\20\2TV\7"+
		"\3\2\2UT\3\2\2\2UV\3\2\2\2V`\3\2\2\2WY\5\66\34\2XZ\7\3\2\2YX\3\2\2\2Y"+
		"Z\3\2\2\2Z`\3\2\2\2[]\58\35\2\\^\7\3\2\2]\\\3\2\2\2]^\3\2\2\2^`\3\2\2"+
		"\2_G\3\2\2\2_K\3\2\2\2_O\3\2\2\2_S\3\2\2\2_W\3\2\2\2_[\3\2\2\2`\7\3\2"+
		"\2\2ad\5\f\7\2bd\5\20\t\2ca\3\2\2\2cb\3\2\2\2d\t\3\2\2\2ef\7\4\2\2fg\5"+
		"\4\3\2gh\7\5\2\2h\13\3\2\2\2ij\7\6\2\2jk\7;\2\2kn\7\7\2\2lo\5:\36\2mo"+
		"\5\16\b\2nl\3\2\2\2nm\3\2\2\2op\3\2\2\2pq\5\n\6\2q\r\3\2\2\2rs\7\66\2"+
		"\2st\7\b\2\2tu\7\66\2\2u\17\3\2\2\2vw\7\t\2\2wx\5:\36\2xy\5\n\6\2y\21"+
		"\3\2\2\2z~\5\24\13\2{~\5\26\f\2|~\5\30\r\2}z\3\2\2\2}{\3\2\2\2}|\3\2\2"+
		"\2~\23\3\2\2\2\177\u0080\7\n\2\2\u0080\u0081\5:\36\2\u0081\u0082\5\n\6"+
		"\2\u0082\u0090\3\2\2\2\u0083\u0084\7\n\2\2\u0084\u0085\5:\36\2\u0085\u0086"+
		"\5\n\6\2\u0086\u0087\7\13\2\2\u0087\u0088\5\n\6\2\u0088\u0090\3\2\2\2"+
		"\u0089\u008a\7\n\2\2\u008a\u008b\5:\36\2\u008b\u008c\5\n\6\2\u008c\u008d"+
		"\7\13\2\2\u008d\u008e\5\24\13\2\u008e\u0090\3\2\2\2\u008f\177\3\2\2\2"+
		"\u008f\u0083\3\2\2\2\u008f\u0089\3\2\2\2\u0090\25\3\2\2\2\u0091\u0092"+
		"\7\f\2\2\u0092\u0093\5:\36\2\u0093\u0094\7\13\2\2\u0094\u0095\7\4\2\2"+
		"\u0095\u0096\5\4\3\2\u0096\u0097\t\2\2\2\u0097\u0098\7\5\2\2\u0098\27"+
		"\3\2\2\2\u0099\u009a\7\20\2\2\u009a\u009b\5:\36\2\u009b\u009f\7\4\2\2"+
		"\u009c\u009e\5\32\16\2\u009d\u009c\3\2\2\2\u009e\u00a1\3\2\2\2\u009f\u009d"+
		"\3\2\2\2\u009f\u00a0\3\2\2\2\u00a0\u00a3\3\2\2\2\u00a1\u009f\3\2\2\2\u00a2"+
		"\u00a4\5\34\17\2\u00a3\u00a2\3\2\2\2\u00a3\u00a4\3\2\2\2\u00a4\u00a5\3"+
		"\2\2\2\u00a5\u00a6\7\5\2\2\u00a6\31\3\2\2\2\u00a7\u00a8\7\21\2\2\u00a8"+
		"\u00a9\5:\36\2\u00a9\u00aa\7\22\2\2\u00aa\u00ac\5\4\3\2\u00ab\u00ad\7"+
		"\16\2\2\u00ac\u00ab\3\2\2\2\u00ac\u00ad\3\2\2\2\u00ad\33\3\2\2\2\u00ae"+
		"\u00af\7\23\2\2\u00af\u00b0\7\22\2\2\u00b0\u00b2\5\4\3\2\u00b1\u00b3\7"+
		"\16\2\2\u00b2\u00b1\3\2\2\2\u00b2\u00b3\3\2\2\2\u00b3\35\3\2\2\2\u00b4"+
		"\u00b8\5 \21\2\u00b5\u00b8\5\"\22\2\u00b6\u00b8\5$\23\2\u00b7\u00b4\3"+
		"\2\2\2\u00b7\u00b5\3\2\2\2\u00b7\u00b6\3\2\2\2\u00b8\37\3\2\2\2\u00b9"+
		"\u00ba\7\16\2\2\u00ba!\3\2\2\2\u00bb\u00bc\7\r\2\2\u00bc#\3\2\2\2\u00bd"+
		"\u00bf\7\17\2\2\u00be\u00c0\5:\36\2\u00bf\u00be\3\2\2\2\u00bf\u00c0\3"+
		"\2\2\2\u00c0%\3\2\2\2\u00c1\u00c5\5(\25\2\u00c2\u00c5\5*\26\2\u00c3\u00c5"+
		"\5\60\31\2\u00c4\u00c1\3\2\2\2\u00c4\u00c2\3\2\2\2\u00c4\u00c3\3\2\2\2"+
		"\u00c5\'\3\2\2\2\u00c6\u00c7\7\24\2\2\u00c7\u00c9\7;\2\2\u00c8\u00ca\5"+
		",\27\2\u00c9\u00c8\3\2\2\2\u00c9\u00ca\3\2\2\2\u00ca\u00cb\3\2\2\2\u00cb"+
		"\u00cc\7\25\2\2\u00cc\u00cd\5:\36\2\u00cd)\3\2\2\2\u00ce\u00cf\7\26\2"+
		"\2\u00cf\u00d0\7;\2\2\u00d0\u00d1\5,\27\2\u00d1\u00d2\7\27\2\2\u00d2\u00db"+
		"\3\2\2\2\u00d3\u00d4\7\26\2\2\u00d4\u00d6\7;\2\2\u00d5\u00d7\5,\27\2\u00d6"+
		"\u00d5\3\2\2\2\u00d6\u00d7\3\2\2\2\u00d7\u00d8\3\2\2\2\u00d8\u00d9\7\25"+
		"\2\2\u00d9\u00db\5:\36\2\u00da\u00ce\3\2\2\2\u00da\u00d3\3\2\2\2\u00db"+
		"+\3\2\2\2\u00dc\u00dd\7\22\2\2\u00dd\u00de\5.\30\2\u00de-\3\2\2\2\u00df"+
		"\u00e0\t\3\2\2\u00e0/\3\2\2\2\u00e1\u00e2\7\26\2\2\u00e2\u00e3\7;\2\2"+
		"\u00e3\u00e4\7\22\2\2\u00e4\u00e5\7\35\2\2\u00e5\u00e6\5.\30\2\u00e6\u00e7"+
		"\7\36\2\2\u00e7\u00e8\5\62\32\2\u00e8\61\3\2\2\2\u00e9\u00ea\7\25\2\2"+
		"\u00ea\u00eb\7\35\2\2\u00eb\u00ec\5\64\33\2\u00ec\u00ed\7\36\2\2\u00ed"+
		"\u00f4\3\2\2\2\u00ee\u00ef\7\25\2\2\u00ef\u00f0\7\35\2\2\u00f0\u00f4\7"+
		"\36\2\2\u00f1\u00f2\7\25\2\2\u00f2\u00f4\7;\2\2\u00f3\u00e9\3\2\2\2\u00f3"+
		"\u00ee\3\2\2\2\u00f3\u00f1\3\2\2\2\u00f4\63\3\2\2\2\u00f5\u00fa\5:\36"+
		"\2\u00f6\u00f7\7\37\2\2\u00f7\u00f9\5:\36\2\u00f8\u00f6\3\2\2\2\u00f9"+
		"\u00fc\3\2\2\2\u00fa\u00f8\3\2\2\2\u00fa\u00fb\3\2\2\2\u00fb\65\3\2\2"+
		"\2\u00fc\u00fa\3\2\2\2\u00fd\u00fe\7 \2\2\u00fe\u00ff\7!\2\2\u00ff\u0100"+
		"\5:\36\2\u0100\u0101\7\"\2\2\u0101\67\3\2\2\2\u0102\u0103\7;\2\2\u0103"+
		"\u0104\7\25\2\2\u0104\u0121\5:\36\2\u0105\u0106\7;\2\2\u0106\u0107\7#"+
		"\2\2\u0107\u0121\5:\36\2\u0108\u0109\7;\2\2\u0109\u010a\7$\2\2\u010a\u0121"+
		"\5:\36\2\u010b\u010c\7;\2\2\u010c\u010d\7\35\2\2\u010d\u010e\5:\36\2\u010e"+
		"\u010f\7\36\2\2\u010f\u0110\7\25\2\2\u0110\u0111\5:\36\2\u0111\u0121\3"+
		"\2\2\2\u0112\u0113\7;\2\2\u0113\u0114\7\35\2\2\u0114\u0115\5:\36\2\u0115"+
		"\u0116\7\36\2\2\u0116\u0117\7#\2\2\u0117\u0118\5:\36\2\u0118\u0121\3\2"+
		"\2\2\u0119\u011a\7;\2\2\u011a\u011b\7\35\2\2\u011b\u011c\5:\36\2\u011c"+
		"\u011d\7\36\2\2\u011d\u011e\7$\2\2\u011e\u011f\5:\36\2\u011f\u0121\3\2"+
		"\2\2\u0120\u0102\3\2\2\2\u0120\u0105\3\2\2\2\u0120\u0108\3\2\2\2\u0120"+
		"\u010b\3\2\2\2\u0120\u0112\3\2\2\2\u0120\u0119\3\2\2\2\u01219\3\2\2\2"+
		"\u0122\u0123\b\36\1\2\u0123\u0133\5<\37\2\u0124\u0125\7;\2\2\u0125\u0126"+
		"\7\35\2\2\u0126\u0127\5:\36\2\u0127\u0128\7\36\2\2\u0128\u0133\3\2\2\2"+
		"\u0129\u0133\7;\2\2\u012a\u012b\7!\2\2\u012b\u012c\5:\36\2\u012c\u012d"+
		"\7\"\2\2\u012d\u0133\3\2\2\2\u012e\u012f\7%\2\2\u012f\u0133\5:\36\b\u0130"+
		"\u0131\7&\2\2\u0131\u0133\5:\36\7\u0132\u0122\3\2\2\2\u0132\u0124\3\2"+
		"\2\2\u0132\u0129\3\2\2\2\u0132\u012a\3\2\2\2\u0132\u012e\3\2\2\2\u0132"+
		"\u0130\3\2\2\2\u0133\u0142\3\2\2\2\u0134\u0135\f\6\2\2\u0135\u0136\t\4"+
		"\2\2\u0136\u0141\5:\36\7\u0137\u0138\f\5\2\2\u0138\u0139\t\5\2\2\u0139"+
		"\u0141\5:\36\6\u013a\u013b\f\4\2\2\u013b\u013c\t\6\2\2\u013c\u0141\5:"+
		"\36\5\u013d\u013e\f\3\2\2\u013e\u013f\t\7\2\2\u013f\u0141\5:\36\4\u0140"+
		"\u0134\3\2\2\2\u0140\u0137\3\2\2\2\u0140\u013a\3\2\2\2\u0140\u013d\3\2"+
		"\2\2\u0141\u0144\3\2\2\2\u0142\u0140\3\2\2\2\u0142\u0143\3\2\2\2\u0143"+
		";\3\2\2\2\u0144\u0142\3\2\2\2\u0145\u014b\78\2\2\u0146\u014b\7\66\2\2"+
		"\u0147\u014b\7\67\2\2\u0148\u014b\79\2\2\u0149\u014b\7:\2\2\u014a\u0145"+
		"\3\2\2\2\u014a\u0146\3\2\2\2\u014a\u0147\3\2\2\2\u014a\u0148\3\2\2\2\u014a"+
		"\u0149\3\2\2\2\u014b=\3\2\2\2\37DIMQUY]_cn}\u008f\u009f\u00a3\u00ac\u00b2"+
		"\u00b7\u00bf\u00c4\u00c9\u00d6\u00da\u00f3\u00fa\u0120\u0132\u0140\u0142"+
		"\u014a";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}