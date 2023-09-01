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
		T__45=46, T__46=47, WS=48, Block_comment=49, Line_comment=50, Int=51, 
		Float=52, Caracter=53, String=54, Bool=55, Identificador=56;
	public static final int
		RULE_inicio = 0, RULE_instrucciones = 1, RULE_instruccion = 2, RULE_loop_statement = 3, 
		RULE_code_block = 4, RULE_for_in_statement = 5, RULE_while_statement = 6, 
		RULE_branch_statement = 7, RULE_if_statement = 8, RULE_guard_statement = 9, 
		RULE_switch_statement = 10, RULE_switch_case = 11, RULE_default_case = 12, 
		RULE_control_transfer_statement = 13, RULE_break_statement = 14, RULE_continue_statement = 15, 
		RULE_return_statement = 16, RULE_declaracion = 17, RULE_constant_declaracion = 18, 
		RULE_variable_declaracion = 19, RULE_anotacion_tipo = 20, RULE_tipos = 21, 
		RULE_array_declaracion = 22, RULE_definicion_vector = 23, RULE_lista_expresiones = 24, 
		RULE_funcion_print = 25, RULE_asignacion = 26, RULE_expresion = 27, RULE_primitivos = 28;
	private static String[] makeRuleNames() {
		return new String[] {
			"inicio", "instrucciones", "instruccion", "loop_statement", "code_block", 
			"for_in_statement", "while_statement", "branch_statement", "if_statement", 
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
			null, "';'", "'{'", "'}'", "'for'", "'in'", "'while'", "'if'", "'else'", 
			"'guard'", "'switch'", "'case'", "':'", "'break'", "'default'", "'continue'", 
			"'return'", "'let'", "'='", "'var'", "'?'", "'String'", "'Int'", "'Float'", 
			"'Bool'", "'Character'", "'['", "']'", "','", "'print'", "'('", "')'", 
			"'+='", "'-='", "'-'", "'*'", "'/'", "'%'", "'+'", "'=='", "'!='", "'>'", 
			"'<'", "'<='", "'>='", "'!'", "'&&'", "'||'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			"WS", "Block_comment", "Line_comment", "Int", "Float", "Caracter", "String", 
			"Bool", "Identificador"
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
			setState(58);
			instrucciones();
			setState(59);
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
			setState(64);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(61);
					instruccion();
					}
					} 
				}
				setState(66);
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
			setState(91);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__16:
			case T__18:
				enterOuterAlt(_localctx, 1);
				{
				setState(67);
				declaracion();
				setState(69);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__0) {
					{
					setState(68);
					match(T__0);
					}
				}

				}
				break;
			case T__3:
			case T__5:
				enterOuterAlt(_localctx, 2);
				{
				setState(71);
				loop_statement();
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
			case T__6:
			case T__8:
			case T__9:
				enterOuterAlt(_localctx, 3);
				{
				setState(75);
				branch_statement();
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
			case T__12:
			case T__14:
			case T__15:
				enterOuterAlt(_localctx, 4);
				{
				setState(79);
				control_transfer_statement();
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
			case T__28:
				enterOuterAlt(_localctx, 5);
				{
				setState(83);
				funcion_print();
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
			case Identificador:
				enterOuterAlt(_localctx, 6);
				{
				setState(87);
				asignacion();
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
		enterRule(_localctx, 14, RULE_branch_statement);
		try {
			setState(114);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__6:
				_localctx = new Declarar_ifContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(111);
				if_statement();
				}
				break;
			case T__8:
				_localctx = new Declarar_guardContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(112);
				guard_statement();
				}
				break;
			case T__9:
				_localctx = new Declarar_switchContext(_localctx);
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
		enterRule(_localctx, 16, RULE_if_statement);
		try {
			setState(132);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				_localctx = new If_simpleContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(116);
				match(T__6);
				setState(117);
				expresion(0);
				setState(118);
				code_block();
				}
				break;
			case 2:
				_localctx = new Else_finalContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(120);
				match(T__6);
				setState(121);
				expresion(0);
				setState(122);
				code_block();
				setState(123);
				match(T__7);
				setState(124);
				code_block();
				}
				break;
			case 3:
				_localctx = new Siguiente_ifContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(126);
				match(T__6);
				setState(127);
				expresion(0);
				setState(128);
				code_block();
				setState(129);
				match(T__7);
				setState(130);
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
		enterRule(_localctx, 18, RULE_guard_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(134);
			match(T__8);
			setState(135);
			expresion(0);
			setState(136);
			match(T__7);
			setState(137);
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
		enterRule(_localctx, 20, RULE_switch_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(139);
			match(T__9);
			setState(140);
			expresion(0);
			setState(141);
			match(T__1);
			setState(145);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__10) {
				{
				{
				setState(142);
				switch_case();
				}
				}
				setState(147);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(149);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__13) {
				{
				setState(148);
				default_case();
				}
			}

			setState(151);
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
		enterRule(_localctx, 22, RULE_switch_case);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(153);
			match(T__10);
			setState(154);
			expresion(0);
			setState(155);
			match(T__11);
			setState(156);
			instrucciones();
			setState(158);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__12) {
				{
				setState(157);
				match(T__12);
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
		enterRule(_localctx, 24, RULE_default_case);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(160);
			match(T__13);
			setState(161);
			match(T__11);
			setState(162);
			instrucciones();
			setState(164);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__12) {
				{
				setState(163);
				match(T__12);
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
		enterRule(_localctx, 26, RULE_control_transfer_statement);
		try {
			setState(169);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__12:
				enterOuterAlt(_localctx, 1);
				{
				setState(166);
				break_statement();
				}
				break;
			case T__14:
				enterOuterAlt(_localctx, 2);
				{
				setState(167);
				continue_statement();
				}
				break;
			case T__15:
				enterOuterAlt(_localctx, 3);
				{
				setState(168);
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
		enterRule(_localctx, 28, RULE_break_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(171);
			match(T__12);
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
		enterRule(_localctx, 30, RULE_continue_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(173);
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
		enterRule(_localctx, 32, RULE_return_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(175);
			match(T__15);
			setState(177);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				{
				setState(176);
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
		enterRule(_localctx, 34, RULE_declaracion);
		try {
			setState(182);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(179);
				constant_declaracion();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(180);
				variable_declaracion();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(181);
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
		enterRule(_localctx, 36, RULE_constant_declaracion);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(184);
			match(T__16);
			setState(185);
			match(Identificador);
			setState(187);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__11) {
				{
				setState(186);
				anotacion_tipo();
				}
			}

			setState(189);
			match(T__17);
			setState(190);
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
		enterRule(_localctx, 38, RULE_variable_declaracion);
		int _la;
		try {
			setState(204);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(192);
				match(T__18);
				setState(193);
				match(Identificador);
				setState(194);
				anotacion_tipo();
				setState(195);
				match(T__19);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(197);
				match(T__18);
				setState(198);
				match(Identificador);
				setState(200);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__11) {
					{
					setState(199);
					anotacion_tipo();
					}
				}

				setState(202);
				match(T__17);
				setState(203);
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
		enterRule(_localctx, 40, RULE_anotacion_tipo);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(206);
			match(T__11);
			setState(207);
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
		enterRule(_localctx, 42, RULE_tipos);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(209);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24))) != 0)) ) {
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
		enterRule(_localctx, 44, RULE_array_declaracion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(211);
			match(T__18);
			setState(212);
			match(Identificador);
			setState(213);
			match(T__11);
			setState(214);
			match(T__25);
			setState(215);
			tipos();
			setState(216);
			match(T__26);
			setState(217);
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
		enterRule(_localctx, 46, RULE_definicion_vector);
		try {
			setState(228);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(219);
				match(T__17);
				setState(220);
				match(T__25);
				setState(221);
				lista_expresiones();
				setState(222);
				match(T__26);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(224);
				match(T__17);
				setState(225);
				match(T__25);
				setState(226);
				match(T__26);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(227);
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
		enterRule(_localctx, 48, RULE_lista_expresiones);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(230);
			expresion(0);
			setState(235);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__27) {
				{
				{
				setState(231);
				match(T__27);
				setState(232);
				expresion(0);
				}
				}
				setState(237);
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
		enterRule(_localctx, 50, RULE_funcion_print);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(238);
			match(T__28);
			setState(239);
			match(T__29);
			setState(240);
			expresion(0);
			setState(241);
			match(T__30);
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
	public static class Asignacion_incrementoContext extends AsignacionContext {
		public TerminalNode Identificador() { return getToken(T_swiftParser.Identificador, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Asignacion_incrementoContext(AsignacionContext ctx) { copyFrom(ctx); }
	}

	public final AsignacionContext asignacion() throws RecognitionException {
		AsignacionContext _localctx = new AsignacionContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_asignacion);
		try {
			setState(252);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
			case 1:
				_localctx = new Asignacion_normalContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(243);
				match(Identificador);
				setState(244);
				match(T__17);
				setState(245);
				expresion(0);
				}
				break;
			case 2:
				_localctx = new Asignacion_incrementoContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(246);
				match(Identificador);
				setState(247);
				match(T__31);
				setState(248);
				expresion(0);
				}
				break;
			case 3:
				_localctx = new Asignacion_decrementoContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(249);
				match(Identificador);
				setState(250);
				match(T__32);
				setState(251);
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
		int _startState = 54;
		enterRecursionRule(_localctx, 54, RULE_expresion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(265);
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

				setState(255);
				primitivos();
				}
				break;
			case Identificador:
				{
				_localctx = new Expresion_idContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(256);
				match(Identificador);
				}
				break;
			case T__33:
				{
				_localctx = new Expresion_unarioContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(257);
				match(T__33);
				setState(258);
				expresion(7);
				}
				break;
			case T__29:
				{
				_localctx = new Expresion_parenContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(259);
				match(T__29);
				setState(260);
				expresion(0);
				setState(261);
				match(T__30);
				}
				break;
			case T__44:
				{
				_localctx = new Expresion_negaContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(263);
				((Expresion_negaContext)_localctx).op = match(T__44);
				setState(264);
				expresion(2);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(281);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,26,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(279);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
					case 1:
						{
						_localctx = new Expresion_aritContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(267);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(268);
						((Expresion_aritContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__34) | (1L << T__35) | (1L << T__36))) != 0)) ) {
							((Expresion_aritContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(269);
						expresion(6);
						}
						break;
					case 2:
						{
						_localctx = new Expresion_aritContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(270);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(271);
						((Expresion_aritContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__33 || _la==T__37) ) {
							((Expresion_aritContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(272);
						expresion(5);
						}
						break;
					case 3:
						{
						_localctx = new Expresion_compaContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(273);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(274);
						((Expresion_compaContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << T__43))) != 0)) ) {
							((Expresion_compaContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(275);
						expresion(4);
						}
						break;
					case 4:
						{
						_localctx = new Expresion_relaContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(276);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(277);
						((Expresion_relaContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__45 || _la==T__46) ) {
							((Expresion_relaContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(278);
						expresion(2);
						}
						break;
					}
					} 
				}
				setState(283);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,26,_ctx);
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
		enterRule(_localctx, 56, RULE_primitivos);
		try {
			setState(289);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Caracter:
				_localctx = new Primitivo_caracterContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(284);
				match(Caracter);
				}
				break;
			case Int:
				_localctx = new Primitivo_intContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(285);
				match(Int);
				}
				break;
			case Float:
				_localctx = new Primitivo_floatContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(286);
				match(Float);
				}
				break;
			case String:
				_localctx = new Primitivo_stringContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(287);
				match(String);
				}
				break;
			case Bool:
				_localctx = new Primitivo_boolContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(288);
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
		case 27:
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3:\u0126\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\3\2\3\2\3\2\3\3\7\3"+
		"A\n\3\f\3\16\3D\13\3\3\4\3\4\5\4H\n\4\3\4\3\4\5\4L\n\4\3\4\3\4\5\4P\n"+
		"\4\3\4\3\4\5\4T\n\4\3\4\3\4\5\4X\n\4\3\4\3\4\5\4\\\n\4\5\4^\n\4\3\5\3"+
		"\5\5\5b\n\5\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3"+
		"\t\3\t\3\t\5\tu\n\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\5\n\u0087\n\n\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f"+
		"\7\f\u0092\n\f\f\f\16\f\u0095\13\f\3\f\5\f\u0098\n\f\3\f\3\f\3\r\3\r\3"+
		"\r\3\r\3\r\5\r\u00a1\n\r\3\16\3\16\3\16\3\16\5\16\u00a7\n\16\3\17\3\17"+
		"\3\17\5\17\u00ac\n\17\3\20\3\20\3\21\3\21\3\22\3\22\5\22\u00b4\n\22\3"+
		"\23\3\23\3\23\5\23\u00b9\n\23\3\24\3\24\3\24\5\24\u00be\n\24\3\24\3\24"+
		"\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\5\25\u00cb\n\25\3\25\3\25"+
		"\5\25\u00cf\n\25\3\26\3\26\3\26\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\30"+
		"\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\5\31\u00e7\n\31"+
		"\3\32\3\32\3\32\7\32\u00ec\n\32\f\32\16\32\u00ef\13\32\3\33\3\33\3\33"+
		"\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\5\34\u00ff\n\34"+
		"\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\5\35\u010c\n\35"+
		"\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\7\35\u011a"+
		"\n\35\f\35\16\35\u011d\13\35\3\36\3\36\3\36\3\36\3\36\5\36\u0124\n\36"+
		"\3\36\2\38\37\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\62\64"+
		"\668:\2\7\3\2\27\33\3\2%\'\4\2$$((\3\2).\3\2\60\61\2\u0136\2<\3\2\2\2"+
		"\4B\3\2\2\2\6]\3\2\2\2\ba\3\2\2\2\nc\3\2\2\2\fg\3\2\2\2\16m\3\2\2\2\20"+
		"t\3\2\2\2\22\u0086\3\2\2\2\24\u0088\3\2\2\2\26\u008d\3\2\2\2\30\u009b"+
		"\3\2\2\2\32\u00a2\3\2\2\2\34\u00ab\3\2\2\2\36\u00ad\3\2\2\2 \u00af\3\2"+
		"\2\2\"\u00b1\3\2\2\2$\u00b8\3\2\2\2&\u00ba\3\2\2\2(\u00ce\3\2\2\2*\u00d0"+
		"\3\2\2\2,\u00d3\3\2\2\2.\u00d5\3\2\2\2\60\u00e6\3\2\2\2\62\u00e8\3\2\2"+
		"\2\64\u00f0\3\2\2\2\66\u00fe\3\2\2\28\u010b\3\2\2\2:\u0123\3\2\2\2<=\5"+
		"\4\3\2=>\7\2\2\3>\3\3\2\2\2?A\5\6\4\2@?\3\2\2\2AD\3\2\2\2B@\3\2\2\2BC"+
		"\3\2\2\2C\5\3\2\2\2DB\3\2\2\2EG\5$\23\2FH\7\3\2\2GF\3\2\2\2GH\3\2\2\2"+
		"H^\3\2\2\2IK\5\b\5\2JL\7\3\2\2KJ\3\2\2\2KL\3\2\2\2L^\3\2\2\2MO\5\20\t"+
		"\2NP\7\3\2\2ON\3\2\2\2OP\3\2\2\2P^\3\2\2\2QS\5\34\17\2RT\7\3\2\2SR\3\2"+
		"\2\2ST\3\2\2\2T^\3\2\2\2UW\5\64\33\2VX\7\3\2\2WV\3\2\2\2WX\3\2\2\2X^\3"+
		"\2\2\2Y[\5\66\34\2Z\\\7\3\2\2[Z\3\2\2\2[\\\3\2\2\2\\^\3\2\2\2]E\3\2\2"+
		"\2]I\3\2\2\2]M\3\2\2\2]Q\3\2\2\2]U\3\2\2\2]Y\3\2\2\2^\7\3\2\2\2_b\5\f"+
		"\7\2`b\5\16\b\2a_\3\2\2\2a`\3\2\2\2b\t\3\2\2\2cd\7\4\2\2de\5\4\3\2ef\7"+
		"\5\2\2f\13\3\2\2\2gh\7\6\2\2hi\7:\2\2ij\7\7\2\2jk\58\35\2kl\5\n\6\2l\r"+
		"\3\2\2\2mn\7\b\2\2no\58\35\2op\5\n\6\2p\17\3\2\2\2qu\5\22\n\2ru\5\24\13"+
		"\2su\5\26\f\2tq\3\2\2\2tr\3\2\2\2ts\3\2\2\2u\21\3\2\2\2vw\7\t\2\2wx\5"+
		"8\35\2xy\5\n\6\2y\u0087\3\2\2\2z{\7\t\2\2{|\58\35\2|}\5\n\6\2}~\7\n\2"+
		"\2~\177\5\n\6\2\177\u0087\3\2\2\2\u0080\u0081\7\t\2\2\u0081\u0082\58\35"+
		"\2\u0082\u0083\5\n\6\2\u0083\u0084\7\n\2\2\u0084\u0085\5\22\n\2\u0085"+
		"\u0087\3\2\2\2\u0086v\3\2\2\2\u0086z\3\2\2\2\u0086\u0080\3\2\2\2\u0087"+
		"\23\3\2\2\2\u0088\u0089\7\13\2\2\u0089\u008a\58\35\2\u008a\u008b\7\n\2"+
		"\2\u008b\u008c\5\n\6\2\u008c\25\3\2\2\2\u008d\u008e\7\f\2\2\u008e\u008f"+
		"\58\35\2\u008f\u0093\7\4\2\2\u0090\u0092\5\30\r\2\u0091\u0090\3\2\2\2"+
		"\u0092\u0095\3\2\2\2\u0093\u0091\3\2\2\2\u0093\u0094\3\2\2\2\u0094\u0097"+
		"\3\2\2\2\u0095\u0093\3\2\2\2\u0096\u0098\5\32\16\2\u0097\u0096\3\2\2\2"+
		"\u0097\u0098\3\2\2\2\u0098\u0099\3\2\2\2\u0099\u009a\7\5\2\2\u009a\27"+
		"\3\2\2\2\u009b\u009c\7\r\2\2\u009c\u009d\58\35\2\u009d\u009e\7\16\2\2"+
		"\u009e\u00a0\5\4\3\2\u009f\u00a1\7\17\2\2\u00a0\u009f\3\2\2\2\u00a0\u00a1"+
		"\3\2\2\2\u00a1\31\3\2\2\2\u00a2\u00a3\7\20\2\2\u00a3\u00a4\7\16\2\2\u00a4"+
		"\u00a6\5\4\3\2\u00a5\u00a7\7\17\2\2\u00a6\u00a5\3\2\2\2\u00a6\u00a7\3"+
		"\2\2\2\u00a7\33\3\2\2\2\u00a8\u00ac\5\36\20\2\u00a9\u00ac\5 \21\2\u00aa"+
		"\u00ac\5\"\22\2\u00ab\u00a8\3\2\2\2\u00ab\u00a9\3\2\2\2\u00ab\u00aa\3"+
		"\2\2\2\u00ac\35\3\2\2\2\u00ad\u00ae\7\17\2\2\u00ae\37\3\2\2\2\u00af\u00b0"+
		"\7\21\2\2\u00b0!\3\2\2\2\u00b1\u00b3\7\22\2\2\u00b2\u00b4\58\35\2\u00b3"+
		"\u00b2\3\2\2\2\u00b3\u00b4\3\2\2\2\u00b4#\3\2\2\2\u00b5\u00b9\5&\24\2"+
		"\u00b6\u00b9\5(\25\2\u00b7\u00b9\5.\30\2\u00b8\u00b5\3\2\2\2\u00b8\u00b6"+
		"\3\2\2\2\u00b8\u00b7\3\2\2\2\u00b9%\3\2\2\2\u00ba\u00bb\7\23\2\2\u00bb"+
		"\u00bd\7:\2\2\u00bc\u00be\5*\26\2\u00bd\u00bc\3\2\2\2\u00bd\u00be\3\2"+
		"\2\2\u00be\u00bf\3\2\2\2\u00bf\u00c0\7\24\2\2\u00c0\u00c1\58\35\2\u00c1"+
		"\'\3\2\2\2\u00c2\u00c3\7\25\2\2\u00c3\u00c4\7:\2\2\u00c4\u00c5\5*\26\2"+
		"\u00c5\u00c6\7\26\2\2\u00c6\u00cf\3\2\2\2\u00c7\u00c8\7\25\2\2\u00c8\u00ca"+
		"\7:\2\2\u00c9\u00cb\5*\26\2\u00ca\u00c9\3\2\2\2\u00ca\u00cb\3\2\2\2\u00cb"+
		"\u00cc\3\2\2\2\u00cc\u00cd\7\24\2\2\u00cd\u00cf\58\35\2\u00ce\u00c2\3"+
		"\2\2\2\u00ce\u00c7\3\2\2\2\u00cf)\3\2\2\2\u00d0\u00d1\7\16\2\2\u00d1\u00d2"+
		"\5,\27\2\u00d2+\3\2\2\2\u00d3\u00d4\t\2\2\2\u00d4-\3\2\2\2\u00d5\u00d6"+
		"\7\25\2\2\u00d6\u00d7\7:\2\2\u00d7\u00d8\7\16\2\2\u00d8\u00d9\7\34\2\2"+
		"\u00d9\u00da\5,\27\2\u00da\u00db\7\35\2\2\u00db\u00dc\5\60\31\2\u00dc"+
		"/\3\2\2\2\u00dd\u00de\7\24\2\2\u00de\u00df\7\34\2\2\u00df\u00e0\5\62\32"+
		"\2\u00e0\u00e1\7\35\2\2\u00e1\u00e7\3\2\2\2\u00e2\u00e3\7\24\2\2\u00e3"+
		"\u00e4\7\34\2\2\u00e4\u00e7\7\35\2\2\u00e5\u00e7\7:\2\2\u00e6\u00dd\3"+
		"\2\2\2\u00e6\u00e2\3\2\2\2\u00e6\u00e5\3\2\2\2\u00e7\61\3\2\2\2\u00e8"+
		"\u00ed\58\35\2\u00e9\u00ea\7\36\2\2\u00ea\u00ec\58\35\2\u00eb\u00e9\3"+
		"\2\2\2\u00ec\u00ef\3\2\2\2\u00ed\u00eb\3\2\2\2\u00ed\u00ee\3\2\2\2\u00ee"+
		"\63\3\2\2\2\u00ef\u00ed\3\2\2\2\u00f0\u00f1\7\37\2\2\u00f1\u00f2\7 \2"+
		"\2\u00f2\u00f3\58\35\2\u00f3\u00f4\7!\2\2\u00f4\65\3\2\2\2\u00f5\u00f6"+
		"\7:\2\2\u00f6\u00f7\7\24\2\2\u00f7\u00ff\58\35\2\u00f8\u00f9\7:\2\2\u00f9"+
		"\u00fa\7\"\2\2\u00fa\u00ff\58\35\2\u00fb\u00fc\7:\2\2\u00fc\u00fd\7#\2"+
		"\2\u00fd\u00ff\58\35\2\u00fe\u00f5\3\2\2\2\u00fe\u00f8\3\2\2\2\u00fe\u00fb"+
		"\3\2\2\2\u00ff\67\3\2\2\2\u0100\u0101\b\35\1\2\u0101\u010c\5:\36\2\u0102"+
		"\u010c\7:\2\2\u0103\u0104\7$\2\2\u0104\u010c\58\35\t\u0105\u0106\7 \2"+
		"\2\u0106\u0107\58\35\2\u0107\u0108\7!\2\2\u0108\u010c\3\2\2\2\u0109\u010a"+
		"\7/\2\2\u010a\u010c\58\35\4\u010b\u0100\3\2\2\2\u010b\u0102\3\2\2\2\u010b"+
		"\u0103\3\2\2\2\u010b\u0105\3\2\2\2\u010b\u0109\3\2\2\2\u010c\u011b\3\2"+
		"\2\2\u010d\u010e\f\7\2\2\u010e\u010f\t\3\2\2\u010f\u011a\58\35\b\u0110"+
		"\u0111\f\6\2\2\u0111\u0112\t\4\2\2\u0112\u011a\58\35\7\u0113\u0114\f\5"+
		"\2\2\u0114\u0115\t\5\2\2\u0115\u011a\58\35\6\u0116\u0117\f\3\2\2\u0117"+
		"\u0118\t\6\2\2\u0118\u011a\58\35\4\u0119\u010d\3\2\2\2\u0119\u0110\3\2"+
		"\2\2\u0119\u0113\3\2\2\2\u0119\u0116\3\2\2\2\u011a\u011d\3\2\2\2\u011b"+
		"\u0119\3\2\2\2\u011b\u011c\3\2\2\2\u011c9\3\2\2\2\u011d\u011b\3\2\2\2"+
		"\u011e\u0124\7\67\2\2\u011f\u0124\7\65\2\2\u0120\u0124\7\66\2\2\u0121"+
		"\u0124\78\2\2\u0122\u0124\79\2\2\u0123\u011e\3\2\2\2\u0123\u011f\3\2\2"+
		"\2\u0123\u0120\3\2\2\2\u0123\u0121\3\2\2\2\u0123\u0122\3\2\2\2\u0124;"+
		"\3\2\2\2\36BGKOSW[]at\u0086\u0093\u0097\u00a0\u00a6\u00ab\u00b3\u00b8"+
		"\u00bd\u00ca\u00ce\u00e6\u00ed\u00fe\u010b\u0119\u011b\u0123";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}