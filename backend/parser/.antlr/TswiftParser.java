// Generated from c:\Users\ruben\Desktop\Compi_2\Lab\Proyecto 1\backend\parser\Tswift.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class TswiftParser extends Parser {
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
		T__45=46, T__46=47, T__47=48, T__48=49, T__49=50, T__50=51, T__51=52, 
		T__52=53, T__53=54, T__54=55, T__55=56, T__56=57, T__57=58, T__58=59, 
		T__59=60, T__60=61, T__61=62, T__62=63, T__63=64, T__64=65, T__65=66, 
		T__66=67, T__67=68, T__68=69, T__69=70, T__70=71, T__71=72, T__72=73, 
		T__73=74, T__74=75, T__75=76, T__76=77, T__77=78, T__78=79, T__79=80, 
		T__80=81, T__81=82, T__82=83, T__83=84, T__84=85, T__85=86, T__86=87, 
		T__87=88, T__88=89, T__89=90, T__90=91, T__91=92, T__92=93, T__93=94, 
		T__94=95, T__95=96, T__96=97, T__97=98, T__98=99, T__99=100, T__100=101, 
		T__101=102, T__102=103, T__103=104, T__104=105, T__105=106, T__106=107, 
		T__107=108, T__108=109, T__109=110, T__110=111, T__111=112, T__112=113, 
		T__113=114, T__114=115, T__115=116, T__116=117, T__117=118, T__118=119, 
		T__119=120, T__120=121, T__121=122, T__122=123, COMENTARIO_COMPUESTO=124, 
		COMENTARIO_SIMPLE=125, Identifier=126, DOT=127, LCURLY=128, LPAREN=129, 
		LBRACK=130, RCURLY=131, RPAREN=132, RBRACK=133, COMMA=134, COLON=135, 
		SEMI=136, LT=137, GT=138, UNDERSCORE=139, BANG=140, QUESTION=141, AT=142, 
		AND=143, SUB=144, EQUAL=145, OR=146, DIV=147, ADD=148, MUL=149, MOD=150, 
		CARET=151, TILDE=152, Operator_head_other=153, Operator_following_character=154, 
		Implicit_parameter_name=155, Binary_literal=156, Octal_literal=157, Decimal_literal=158, 
		Pure_decimal_digits=159, Hexadecimal_literal=160, Floating_point_literal=161, 
		Static_string_literal=162, Interpolated_string_literal=163, WS=164, Block_comment=165, 
		Line_comment=166;
	public static final int
		RULE_inicio = 0, RULE_instrucciones = 1, RULE_instruccion = 2, RULE_loop_statement = 3, 
		RULE_for_in_statement = 4, RULE_while_statement = 5, RULE_branch_statement = 6, 
		RULE_if_statement = 7, RULE_else_clause = 8, RULE_guard_statement = 9, 
		RULE_switch_statement = 10, RULE_switch_cases = 11, RULE_switch_case = 12, 
		RULE_case_label = 13, RULE_default_label = 14, RULE_control_transfer_statement = 15, 
		RULE_break_statement = 16, RULE_continue_statement = 17, RULE_return_statement = 18, 
		RULE_generic_parameter_clause = 19, RULE_generic_parameter_list = 20, 
		RULE_generic_parameter = 21, RULE_generic_where_clause = 22, RULE_requirement_list = 23, 
		RULE_requirement = 24, RULE_conformance_requirement = 25, RULE_same_type_requirement = 26, 
		RULE_generic_argument_clause = 27, RULE_generic_argument_list = 28, RULE_generic_argument = 29, 
		RULE_declaracion = 30, RULE_declaracions = 31, RULE_code_block = 32, RULE_import_kind = 33, 
		RULE_import_path = 34, RULE_import_path_identifier = 35, RULE_constant_declaracion = 36, 
		RULE_pattern_initializer_list = 37, RULE_pattern_initializer = 38, RULE_initializer = 39, 
		RULE_variable_declaracion = 40, RULE_variable_declaracion_head = 41, RULE_variable_name = 42, 
		RULE_getter_setter_block = 43, RULE_getter_clause = 44, RULE_setter_clause = 45, 
		RULE_setter_name = 46, RULE_getter_setter_keyword_block = 47, RULE_getter_keyword_clause = 48, 
		RULE_setter_keyword_clause = 49, RULE_willSet_didSet_block = 50, RULE_willSet_clause = 51, 
		RULE_didSet_clause = 52, RULE_function_declaracion = 53, RULE_function_head = 54, 
		RULE_function_name = 55, RULE_function_signature = 56, RULE_function_result = 57, 
		RULE_function_body = 58, RULE_parameter_clause = 59, RULE_parameter_list = 60, 
		RULE_parameter = 61, RULE_external_parameter_name = 62, RULE_local_parameter_name = 63, 
		RULE_default_argument_clause = 64, RULE_enum_declaracion = 65, RULE_union_style_enum = 66, 
		RULE_union_style_enum_members = 67, RULE_union_style_enum_member = 68, 
		RULE_union_style_enum_case_clause = 69, RULE_union_style_enum_case_list = 70, 
		RULE_union_style_enum_case = 71, RULE_enum_name = 72, RULE_enum_case_name = 73, 
		RULE_raw_value_style_enum = 74, RULE_raw_value_style_enum_members = 75, 
		RULE_raw_value_style_enum_member = 76, RULE_raw_value_style_enum_case_clause = 77, 
		RULE_raw_value_style_enum_case_list = 78, RULE_raw_value_style_enum_case = 79, 
		RULE_raw_value_assignment = 80, RULE_raw_value_literal = 81, RULE_struct_declaracion = 82, 
		RULE_struct_name = 83, RULE_struct_body = 84, RULE_struct_member = 85, 
		RULE_class_declaracion = 86, RULE_class_name = 87, RULE_class_body = 88, 
		RULE_class_member = 89, RULE_protocol_declaracion = 90, RULE_protocol_name = 91, 
		RULE_protocol_body = 92, RULE_protocol_member = 93, RULE_protocol_member_declaracion = 94, 
		RULE_protocol_property_declaracion = 95, RULE_protocol_method_declaracion = 96, 
		RULE_protocol_initializer_declaracion = 97, RULE_protocol_subscript_declaracion = 98, 
		RULE_protocol_associated_type_declaracion = 99, RULE_initializer_declaracion = 100, 
		RULE_initializer_head = 101, RULE_initializer_body = 102, RULE_deinitializer_declaracion = 103, 
		RULE_extension_declaracion = 104, RULE_extension_body = 105, RULE_extension_member = 106, 
		RULE_subscript_declaracion = 107, RULE_subscript_head = 108, RULE_subscript_result = 109, 
		RULE_operator_declaracion = 110, RULE_prefix_operator_declaracion = 111, 
		RULE_postfix_operator_declaracion = 112, RULE_infix_operator_declaracion = 113, 
		RULE_infix_operator_group = 114, RULE_precedence_group_declaracion = 115, 
		RULE_precedence_group_attribute = 116, RULE_precedence_group_relation = 117, 
		RULE_precedence_group_assignment = 118, RULE_precedence_group_associativity = 119, 
		RULE_associativity_ = 120, RULE_precedence_group_names = 121, RULE_precedence_group_name = 122, 
		RULE_declaracion_modifier = 123, RULE_declaracion_modifiers = 124, RULE_access_level_modifier = 125, 
		RULE_mutation_modifier = 126, RULE_pattern = 127, RULE_wildcard_pattern = 128, 
		RULE_identifier_pattern = 129, RULE_value_binding_pattern = 130, RULE_tuple_pattern = 131, 
		RULE_tuple_pattern_element_list = 132, RULE_tuple_pattern_element = 133, 
		RULE_enum_case_pattern = 134, RULE_optional_pattern = 135, RULE_expresion_pattern = 136, 
		RULE_attribute = 137, RULE_attribute_name = 138, RULE_attribute_argument_clause = 139, 
		RULE_attributes = 140, RULE_balanced_tokens = 141, RULE_balanced_token = 142, 
		RULE_any_punctuation_for_balanced_token = 143, RULE_expresion = 144, RULE_expresion_list = 145, 
		RULE_prefix_expresion = 146, RULE_in_out_expresion = 147, RULE_try_operator = 148, 
		RULE_binary_expresion = 149, RULE_binary_expresions = 150, RULE_conditional_operator = 151, 
		RULE_type_casting_operator = 152, RULE_primary_expresion = 153, RULE_literal_expresion = 154, 
		RULE_array_literal = 155, RULE_array_literal_items = 156, RULE_array_literal_item = 157, 
		RULE_dictionary_literal = 158, RULE_dictionary_literal_items = 159, RULE_dictionary_literal_item = 160, 
		RULE_playground_literal = 161, RULE_self_expresion = 162, RULE_superclass_expresion = 163, 
		RULE_superclass_method_expresion = 164, RULE_superclass_subscript_expresion = 165, 
		RULE_superclass_initializer_expresion = 166, RULE_closure_expresion = 167, 
		RULE_closure_signature = 168, RULE_closure_parameter_clause = 169, RULE_closure_parameter_clause_identifier_list = 170, 
		RULE_closure_parameter_list = 171, RULE_closure_parameter = 172, RULE_closure_parameter_name = 173, 
		RULE_capture_list = 174, RULE_capture_list_items = 175, RULE_capture_list_item = 176, 
		RULE_capture_specifier = 177, RULE_implicit_member_expresion = 178, RULE_parenthesized_expresion = 179, 
		RULE_tuple_expresion = 180, RULE_tuple_element = 181, RULE_wildcard_expresion = 182, 
		RULE_selector_expresion = 183, RULE_key_path_expresion = 184, RULE_postfix_expresion = 185, 
		RULE_function_call_argument_clause = 186, RULE_function_call_argument_list = 187, 
		RULE_function_call_argument = 188, RULE_trailing_closure = 189, RULE_argument_names = 190, 
		RULE_argument_name = 191, RULE_dynamic_type_expresion = 192, RULE_type_ = 193, 
		RULE_type_annotation = 194, RULE_type_identifier = 195, RULE_type_name = 196, 
		RULE_tuple_type = 197, RULE_tuple_type_element_list = 198, RULE_tuple_type_element = 199, 
		RULE_element_name = 200, RULE_function_type = 201, RULE_function_type_argument_clause = 202, 
		RULE_function_type_argument_list = 203, RULE_function_type_argument = 204, 
		RULE_argument_label = 205, RULE_array_type = 206, RULE_dictionary_type = 207, 
		RULE_protocol_composition_type = 208, RULE_protocol_identifier = 209, 
		RULE_type_inheritance_clause = 210, RULE_type_inheritance_list = 211, 
		RULE_class_requirement = 212, RULE_declaracion_identifier = 213, RULE_label_identifier = 214, 
		RULE_keyword_as_identifier_in_declaracions = 215, RULE_keyword_as_identifier_in_labels = 216, 
		RULE_assignment_operator = 217, RULE_negate_prefix_operator = 218, RULE_compilation_condition_AND = 219, 
		RULE_compilation_condition_OR = 220, RULE_compilation_condition_GE = 221, 
		RULE_arrow_operator = 222, RULE_range_operator = 223, RULE_same_type_equals = 224, 
		RULE_binary_operator = 225, RULE_prefix_operator = 226, RULE_postfix_operator = 227, 
		RULE_operator_ = 228, RULE_operator_character = 229, RULE_operator_head = 230, 
		RULE_dot_operator_head = 231, RULE_dot_operator_character = 232, RULE_literal = 233, 
		RULE_numeric_literal = 234, RULE_boolean_literal = 235, RULE_nil_literal = 236, 
		RULE_integer_literal = 237, RULE_string_literal = 238;
	private static String[] makeRuleNames() {
		return new String[] {
			"inicio", "instrucciones", "instruccion", "loop_statement", "for_in_statement", 
			"while_statement", "branch_statement", "if_statement", "else_clause", 
			"guard_statement", "switch_statement", "switch_cases", "switch_case", 
			"case_label", "default_label", "control_transfer_statement", "break_statement", 
			"continue_statement", "return_statement", "generic_parameter_clause", 
			"generic_parameter_list", "generic_parameter", "generic_where_clause", 
			"requirement_list", "requirement", "conformance_requirement", "same_type_requirement", 
			"generic_argument_clause", "generic_argument_list", "generic_argument", 
			"declaracion", "declaracions", "code_block", "import_kind", "import_path", 
			"import_path_identifier", "constant_declaracion", "pattern_initializer_list", 
			"pattern_initializer", "initializer", "variable_declaracion", "variable_declaracion_head", 
			"variable_name", "getter_setter_block", "getter_clause", "setter_clause", 
			"setter_name", "getter_setter_keyword_block", "getter_keyword_clause", 
			"setter_keyword_clause", "willSet_didSet_block", "willSet_clause", "didSet_clause", 
			"function_declaracion", "function_head", "function_name", "function_signature", 
			"function_result", "function_body", "parameter_clause", "parameter_list", 
			"parameter", "external_parameter_name", "local_parameter_name", "default_argument_clause", 
			"enum_declaracion", "union_style_enum", "union_style_enum_members", "union_style_enum_member", 
			"union_style_enum_case_clause", "union_style_enum_case_list", "union_style_enum_case", 
			"enum_name", "enum_case_name", "raw_value_style_enum", "raw_value_style_enum_members", 
			"raw_value_style_enum_member", "raw_value_style_enum_case_clause", "raw_value_style_enum_case_list", 
			"raw_value_style_enum_case", "raw_value_assignment", "raw_value_literal", 
			"struct_declaracion", "struct_name", "struct_body", "struct_member", 
			"class_declaracion", "class_name", "class_body", "class_member", "protocol_declaracion", 
			"protocol_name", "protocol_body", "protocol_member", "protocol_member_declaracion", 
			"protocol_property_declaracion", "protocol_method_declaracion", "protocol_initializer_declaracion", 
			"protocol_subscript_declaracion", "protocol_associated_type_declaracion", 
			"initializer_declaracion", "initializer_head", "initializer_body", "deinitializer_declaracion", 
			"extension_declaracion", "extension_body", "extension_member", "subscript_declaracion", 
			"subscript_head", "subscript_result", "operator_declaracion", "prefix_operator_declaracion", 
			"postfix_operator_declaracion", "infix_operator_declaracion", "infix_operator_group", 
			"precedence_group_declaracion", "precedence_group_attribute", "precedence_group_relation", 
			"precedence_group_assignment", "precedence_group_associativity", "associativity_", 
			"precedence_group_names", "precedence_group_name", "declaracion_modifier", 
			"declaracion_modifiers", "access_level_modifier", "mutation_modifier", 
			"pattern", "wildcard_pattern", "identifier_pattern", "value_binding_pattern", 
			"tuple_pattern", "tuple_pattern_element_list", "tuple_pattern_element", 
			"enum_case_pattern", "optional_pattern", "expresion_pattern", "attribute", 
			"attribute_name", "attribute_argument_clause", "attributes", "balanced_tokens", 
			"balanced_token", "any_punctuation_for_balanced_token", "expresion", 
			"expresion_list", "prefix_expresion", "in_out_expresion", "try_operator", 
			"binary_expresion", "binary_expresions", "conditional_operator", "type_casting_operator", 
			"primary_expresion", "literal_expresion", "array_literal", "array_literal_items", 
			"array_literal_item", "dictionary_literal", "dictionary_literal_items", 
			"dictionary_literal_item", "playground_literal", "self_expresion", "superclass_expresion", 
			"superclass_method_expresion", "superclass_subscript_expresion", "superclass_initializer_expresion", 
			"closure_expresion", "closure_signature", "closure_parameter_clause", 
			"closure_parameter_clause_identifier_list", "closure_parameter_list", 
			"closure_parameter", "closure_parameter_name", "capture_list", "capture_list_items", 
			"capture_list_item", "capture_specifier", "implicit_member_expresion", 
			"parenthesized_expresion", "tuple_expresion", "tuple_element", "wildcard_expresion", 
			"selector_expresion", "key_path_expresion", "postfix_expresion", "function_call_argument_clause", 
			"function_call_argument_list", "function_call_argument", "trailing_closure", 
			"argument_names", "argument_name", "dynamic_type_expresion", "type_", 
			"type_annotation", "type_identifier", "type_name", "tuple_type", "tuple_type_element_list", 
			"tuple_type_element", "element_name", "function_type", "function_type_argument_clause", 
			"function_type_argument_list", "function_type_argument", "argument_label", 
			"array_type", "dictionary_type", "protocol_composition_type", "protocol_identifier", 
			"type_inheritance_clause", "type_inheritance_list", "class_requirement", 
			"declaracion_identifier", "label_identifier", "keyword_as_identifier_in_declaracions", 
			"keyword_as_identifier_in_labels", "assignment_operator", "negate_prefix_operator", 
			"compilation_condition_AND", "compilation_condition_OR", "compilation_condition_GE", 
			"arrow_operator", "range_operator", "same_type_equals", "binary_operator", 
			"prefix_operator", "postfix_operator", "operator_", "operator_character", 
			"operator_head", "dot_operator_head", "dot_operator_character", "literal", 
			"numeric_literal", "boolean_literal", "nil_literal", "integer_literal", 
			"string_literal"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'for'", "'in'", "'while'", "'if'", "'else'", "'guard'", "'switch'", 
			"'case'", "'default'", "'break'", "'continue'", "'return'", "'where'", 
			"'typealias'", "'struct'", "'class'", "'enum'", "'protocol'", "'var'", 
			"'func'", "'let'", "'get'", "'set'", "'willSet'", "'didSet'", "'throws'", 
			"'rethrows'", "'indirect'", "'final'", "'associatedtype'", "'init'", 
			"'deinit'", "'extension'", "'subscript'", "'prefix'", "'operator'", "'postfix'", 
			"'infix'", "'precedencegroup'", "'higherThan'", "'lowerThan'", "'assignment'", 
			"'associativity'", "'left'", "'right'", "'none'", "'convenience'", "'dynamic'", 
			"'lazy'", "'optional'", "'override'", "'required'", "'static'", "'unowned'", 
			"'safe'", "'unsafe'", "'weak'", "'private'", "'fileprivate'", "'internal'", 
			"'public'", "'open'", "'mutating'", "'nonmutating'", "'is'", "'as'", 
			"'#'", "'`'", "'try'", "'#file'", "'#line'", "'#column'", "'#function'", 
			"'#dsohandle'", "'#colorLiteral'", "'red'", "'green'", "'blue'", "'alpha'", 
			"'#fileLiteral'", "'resourceName'", "'#imageLiteral'", "'self'", "'Self'", 
			"'super'", "'unowned(safe)'", "'unowned(unsafe)'", "'#selector'", "'getter:'", 
			"'setter:'", "'#keyPath'", "'type'", "'of'", "'Type'", "'Protocol'", 
			"'Any'", "'inout'", "'arch'", "'arm'", "'arm64'", "'file'", "'i386'", 
			"'iOS'", "'iOSApplicationExtension'", "'line'", "'macOS'", "'macOSApplicationExtension'", 
			"'os'", "'precedence'", "'swift'", "'tvOS'", "'watchOS'", "'x86_64'", 
			"'catch'", "'defer'", "'do'", "'fallthrough'", "'false'", "'import'", 
			"'nil'", "'repeat'", "'throw'", "'true'", null, null, null, "'.'", "'{'", 
			"'('", "'['", "'}'", "')'", "']'", "','", "':'", "';'", "'<'", "'>'", 
			"'_'", "'!'", "'?'", "'@'", "'&'", "'-'", "'='", "'|'", "'/'", "'+'", 
			"'*'", "'%'", "'^'", "'~'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, "COMENTARIO_COMPUESTO", "COMENTARIO_SIMPLE", 
			"Identifier", "DOT", "LCURLY", "LPAREN", "LBRACK", "RCURLY", "RPAREN", 
			"RBRACK", "COMMA", "COLON", "SEMI", "LT", "GT", "UNDERSCORE", "BANG", 
			"QUESTION", "AT", "AND", "SUB", "EQUAL", "OR", "DIV", "ADD", "MUL", "MOD", 
			"CARET", "TILDE", "Operator_head_other", "Operator_following_character", 
			"Implicit_parameter_name", "Binary_literal", "Octal_literal", "Decimal_literal", 
			"Pure_decimal_digits", "Hexadecimal_literal", "Floating_point_literal", 
			"Static_string_literal", "Interpolated_string_literal", "WS", "Block_comment", 
			"Line_comment"
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
	public String getGrammarFileName() { return "Tswift.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public TswiftParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class InicioContext extends ParserRuleContext {
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode EOF() { return getToken(TswiftParser.EOF, 0); }
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
			setState(478);
			instrucciones();
			setState(479);
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
			setState(484);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(481);
					instruccion();
					}
					} 
				}
				setState(486);
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
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
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
		public InstruccionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion; }
	}

	public final InstruccionContext instruccion() throws RecognitionException {
		InstruccionContext _localctx = new InstruccionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruccion);
		try {
			setState(492);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(487);
				expresion();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(488);
				declaracion();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(489);
				loop_statement();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(490);
				branch_statement();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(491);
				control_transfer_statement();
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
			setState(496);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				enterOuterAlt(_localctx, 1);
				{
				setState(494);
				for_in_statement();
				}
				break;
			case T__2:
				enterOuterAlt(_localctx, 2);
				{
				setState(495);
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

	public static class For_in_statementContext extends ParserRuleContext {
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
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
		enterRule(_localctx, 8, RULE_for_in_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(498);
			match(T__0);
			setState(499);
			pattern(0);
			setState(500);
			match(T__1);
			setState(501);
			expresion();
			setState(502);
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
		enterRule(_localctx, 10, RULE_while_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(504);
			match(T__2);
			setState(505);
			expresion();
			setState(506);
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
		enterRule(_localctx, 12, RULE_branch_statement);
		try {
			setState(511);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__3:
				enterOuterAlt(_localctx, 1);
				{
				setState(508);
				if_statement();
				}
				break;
			case T__5:
				enterOuterAlt(_localctx, 2);
				{
				setState(509);
				guard_statement();
				}
				break;
			case T__6:
				enterOuterAlt(_localctx, 3);
				{
				setState(510);
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
		enterRule(_localctx, 14, RULE_if_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(513);
			match(T__3);
			setState(514);
			expresion();
			setState(515);
			code_block();
			setState(517);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				{
				setState(516);
				else_clause();
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
		enterRule(_localctx, 16, RULE_else_clause);
		try {
			setState(523);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(519);
				match(T__4);
				setState(520);
				code_block();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(521);
				match(T__4);
				setState(522);
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
			setState(525);
			match(T__5);
			setState(526);
			expresion();
			setState(527);
			match(T__4);
			setState(528);
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
		public TerminalNode LCURLY() { return getToken(TswiftParser.LCURLY, 0); }
		public TerminalNode RCURLY() { return getToken(TswiftParser.RCURLY, 0); }
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
		enterRule(_localctx, 20, RULE_switch_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(530);
			match(T__6);
			setState(531);
			expresion();
			setState(532);
			match(LCURLY);
			setState(534);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__7 || _la==T__8) {
				{
				setState(533);
				switch_cases();
				}
			}

			setState(536);
			match(RCURLY);
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
		enterRule(_localctx, 22, RULE_switch_cases);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(538);
			switch_case();
			setState(540);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__7 || _la==T__8) {
				{
				setState(539);
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
		enterRule(_localctx, 24, RULE_switch_case);
		try {
			setState(548);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__7:
				enterOuterAlt(_localctx, 1);
				{
				setState(542);
				case_label();
				setState(543);
				instrucciones();
				}
				break;
			case T__8:
				enterOuterAlt(_localctx, 2);
				{
				setState(545);
				default_label();
				setState(546);
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
		public TerminalNode COLON() { return getToken(TswiftParser.COLON, 0); }
		public Case_labelContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_case_label; }
	}

	public final Case_labelContext case_label() throws RecognitionException {
		Case_labelContext _localctx = new Case_labelContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_case_label);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(550);
			match(T__7);
			setState(551);
			expresion();
			setState(552);
			match(COLON);
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
		public TerminalNode COLON() { return getToken(TswiftParser.COLON, 0); }
		public Default_labelContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_default_label; }
	}

	public final Default_labelContext default_label() throws RecognitionException {
		Default_labelContext _localctx = new Default_labelContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_default_label);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(554);
			match(T__8);
			setState(555);
			match(COLON);
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
		enterRule(_localctx, 30, RULE_control_transfer_statement);
		try {
			setState(560);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__9:
				enterOuterAlt(_localctx, 1);
				{
				setState(557);
				break_statement();
				}
				break;
			case T__10:
				enterOuterAlt(_localctx, 2);
				{
				setState(558);
				continue_statement();
				}
				break;
			case T__11:
				enterOuterAlt(_localctx, 3);
				{
				setState(559);
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
		enterRule(_localctx, 32, RULE_break_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(562);
			match(T__9);
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
		enterRule(_localctx, 34, RULE_continue_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(564);
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
		enterRule(_localctx, 36, RULE_return_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(566);
			match(T__11);
			setState(568);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				{
				setState(567);
				expresion();
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

	public static class Generic_parameter_clauseContext extends ParserRuleContext {
		public TerminalNode LT() { return getToken(TswiftParser.LT, 0); }
		public Generic_parameter_listContext generic_parameter_list() {
			return getRuleContext(Generic_parameter_listContext.class,0);
		}
		public TerminalNode GT() { return getToken(TswiftParser.GT, 0); }
		public Generic_parameter_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_generic_parameter_clause; }
	}

	public final Generic_parameter_clauseContext generic_parameter_clause() throws RecognitionException {
		Generic_parameter_clauseContext _localctx = new Generic_parameter_clauseContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_generic_parameter_clause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(570);
			match(LT);
			setState(571);
			generic_parameter_list();
			setState(572);
			match(GT);
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

	public static class Generic_parameter_listContext extends ParserRuleContext {
		public List<Generic_parameterContext> generic_parameter() {
			return getRuleContexts(Generic_parameterContext.class);
		}
		public Generic_parameterContext generic_parameter(int i) {
			return getRuleContext(Generic_parameterContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(TswiftParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(TswiftParser.COMMA, i);
		}
		public Generic_parameter_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_generic_parameter_list; }
	}

	public final Generic_parameter_listContext generic_parameter_list() throws RecognitionException {
		Generic_parameter_listContext _localctx = new Generic_parameter_listContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_generic_parameter_list);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(574);
			generic_parameter();
			setState(579);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(575);
				match(COMMA);
				setState(576);
				generic_parameter();
				}
				}
				setState(581);
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

	public static class Generic_parameterContext extends ParserRuleContext {
		public Type_nameContext type_name() {
			return getRuleContext(Type_nameContext.class,0);
		}
		public TerminalNode COLON() { return getToken(TswiftParser.COLON, 0); }
		public Type_identifierContext type_identifier() {
			return getRuleContext(Type_identifierContext.class,0);
		}
		public Protocol_composition_typeContext protocol_composition_type() {
			return getRuleContext(Protocol_composition_typeContext.class,0);
		}
		public Generic_parameterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_generic_parameter; }
	}

	public final Generic_parameterContext generic_parameter() throws RecognitionException {
		Generic_parameterContext _localctx = new Generic_parameterContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_generic_parameter);
		try {
			setState(591);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(582);
				type_name();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(583);
				type_name();
				setState(584);
				match(COLON);
				setState(585);
				type_identifier();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(587);
				type_name();
				setState(588);
				match(COLON);
				setState(589);
				protocol_composition_type();
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

	public static class Generic_where_clauseContext extends ParserRuleContext {
		public Requirement_listContext requirement_list() {
			return getRuleContext(Requirement_listContext.class,0);
		}
		public Generic_where_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_generic_where_clause; }
	}

	public final Generic_where_clauseContext generic_where_clause() throws RecognitionException {
		Generic_where_clauseContext _localctx = new Generic_where_clauseContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_generic_where_clause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(593);
			match(T__12);
			setState(594);
			requirement_list();
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

	public static class Requirement_listContext extends ParserRuleContext {
		public List<RequirementContext> requirement() {
			return getRuleContexts(RequirementContext.class);
		}
		public RequirementContext requirement(int i) {
			return getRuleContext(RequirementContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(TswiftParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(TswiftParser.COMMA, i);
		}
		public Requirement_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_requirement_list; }
	}

	public final Requirement_listContext requirement_list() throws RecognitionException {
		Requirement_listContext _localctx = new Requirement_listContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_requirement_list);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(596);
			requirement();
			setState(601);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(597);
					match(COMMA);
					setState(598);
					requirement();
					}
					} 
				}
				setState(603);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
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

	public static class RequirementContext extends ParserRuleContext {
		public Conformance_requirementContext conformance_requirement() {
			return getRuleContext(Conformance_requirementContext.class,0);
		}
		public Same_type_requirementContext same_type_requirement() {
			return getRuleContext(Same_type_requirementContext.class,0);
		}
		public RequirementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_requirement; }
	}

	public final RequirementContext requirement() throws RecognitionException {
		RequirementContext _localctx = new RequirementContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_requirement);
		try {
			setState(606);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(604);
				conformance_requirement();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(605);
				same_type_requirement();
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

	public static class Conformance_requirementContext extends ParserRuleContext {
		public List<Type_identifierContext> type_identifier() {
			return getRuleContexts(Type_identifierContext.class);
		}
		public Type_identifierContext type_identifier(int i) {
			return getRuleContext(Type_identifierContext.class,i);
		}
		public TerminalNode COLON() { return getToken(TswiftParser.COLON, 0); }
		public Protocol_composition_typeContext protocol_composition_type() {
			return getRuleContext(Protocol_composition_typeContext.class,0);
		}
		public Conformance_requirementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conformance_requirement; }
	}

	public final Conformance_requirementContext conformance_requirement() throws RecognitionException {
		Conformance_requirementContext _localctx = new Conformance_requirementContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_conformance_requirement);
		try {
			setState(616);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(608);
				type_identifier();
				setState(609);
				match(COLON);
				setState(610);
				type_identifier();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(612);
				type_identifier();
				setState(613);
				match(COLON);
				setState(614);
				protocol_composition_type();
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

	public static class Same_type_requirementContext extends ParserRuleContext {
		public Type_identifierContext type_identifier() {
			return getRuleContext(Type_identifierContext.class,0);
		}
		public Same_type_equalsContext same_type_equals() {
			return getRuleContext(Same_type_equalsContext.class,0);
		}
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public Same_type_requirementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_same_type_requirement; }
	}

	public final Same_type_requirementContext same_type_requirement() throws RecognitionException {
		Same_type_requirementContext _localctx = new Same_type_requirementContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_same_type_requirement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(618);
			type_identifier();
			setState(619);
			same_type_equals();
			setState(620);
			type_(0);
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

	public static class Generic_argument_clauseContext extends ParserRuleContext {
		public TerminalNode LT() { return getToken(TswiftParser.LT, 0); }
		public Generic_argument_listContext generic_argument_list() {
			return getRuleContext(Generic_argument_listContext.class,0);
		}
		public TerminalNode GT() { return getToken(TswiftParser.GT, 0); }
		public Generic_argument_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_generic_argument_clause; }
	}

	public final Generic_argument_clauseContext generic_argument_clause() throws RecognitionException {
		Generic_argument_clauseContext _localctx = new Generic_argument_clauseContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_generic_argument_clause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(622);
			match(LT);
			setState(623);
			generic_argument_list();
			setState(624);
			match(GT);
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

	public static class Generic_argument_listContext extends ParserRuleContext {
		public List<Generic_argumentContext> generic_argument() {
			return getRuleContexts(Generic_argumentContext.class);
		}
		public Generic_argumentContext generic_argument(int i) {
			return getRuleContext(Generic_argumentContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(TswiftParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(TswiftParser.COMMA, i);
		}
		public Generic_argument_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_generic_argument_list; }
	}

	public final Generic_argument_listContext generic_argument_list() throws RecognitionException {
		Generic_argument_listContext _localctx = new Generic_argument_listContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_generic_argument_list);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(626);
			generic_argument();
			setState(631);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(627);
				match(COMMA);
				setState(628);
				generic_argument();
				}
				}
				setState(633);
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

	public static class Generic_argumentContext extends ParserRuleContext {
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public Generic_argumentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_generic_argument; }
	}

	public final Generic_argumentContext generic_argument() throws RecognitionException {
		Generic_argumentContext _localctx = new Generic_argumentContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_generic_argument);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(634);
			type_(0);
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
		public Function_declaracionContext function_declaracion() {
			return getRuleContext(Function_declaracionContext.class,0);
		}
		public Enum_declaracionContext enum_declaracion() {
			return getRuleContext(Enum_declaracionContext.class,0);
		}
		public Struct_declaracionContext struct_declaracion() {
			return getRuleContext(Struct_declaracionContext.class,0);
		}
		public Class_declaracionContext class_declaracion() {
			return getRuleContext(Class_declaracionContext.class,0);
		}
		public Protocol_declaracionContext protocol_declaracion() {
			return getRuleContext(Protocol_declaracionContext.class,0);
		}
		public Initializer_declaracionContext initializer_declaracion() {
			return getRuleContext(Initializer_declaracionContext.class,0);
		}
		public Deinitializer_declaracionContext deinitializer_declaracion() {
			return getRuleContext(Deinitializer_declaracionContext.class,0);
		}
		public Extension_declaracionContext extension_declaracion() {
			return getRuleContext(Extension_declaracionContext.class,0);
		}
		public Subscript_declaracionContext subscript_declaracion() {
			return getRuleContext(Subscript_declaracionContext.class,0);
		}
		public Operator_declaracionContext operator_declaracion() {
			return getRuleContext(Operator_declaracionContext.class,0);
		}
		public Precedence_group_declaracionContext precedence_group_declaracion() {
			return getRuleContext(Precedence_group_declaracionContext.class,0);
		}
		public DeclaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion; }
	}

	public final DeclaracionContext declaracion() throws RecognitionException {
		DeclaracionContext _localctx = new DeclaracionContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_declaracion);
		try {
			setState(650);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(636);
				constant_declaracion();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(637);
				variable_declaracion();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(638);
				function_declaracion();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(639);
				enum_declaracion();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(640);
				struct_declaracion();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(641);
				class_declaracion();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(642);
				protocol_declaracion();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(643);
				initializer_declaracion();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(644);
				deinitializer_declaracion();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(645);
				extension_declaracion();
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(646);
				subscript_declaracion();
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(647);
				operator_declaracion();
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(648);
				operator_declaracion();
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(649);
				precedence_group_declaracion();
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

	public static class DeclaracionsContext extends ParserRuleContext {
		public List<DeclaracionContext> declaracion() {
			return getRuleContexts(DeclaracionContext.class);
		}
		public DeclaracionContext declaracion(int i) {
			return getRuleContext(DeclaracionContext.class,i);
		}
		public DeclaracionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracions; }
	}

	public final DeclaracionsContext declaracions() throws RecognitionException {
		DeclaracionsContext _localctx = new DeclaracionsContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_declaracions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(653); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(652);
				declaracion();
				}
				}
				setState(655); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( ((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & ((1L << (T__14 - 15)) | (1L << (T__15 - 15)) | (1L << (T__16 - 15)) | (1L << (T__17 - 15)) | (1L << (T__18 - 15)) | (1L << (T__19 - 15)) | (1L << (T__20 - 15)) | (1L << (T__27 - 15)) | (1L << (T__28 - 15)) | (1L << (T__30 - 15)) | (1L << (T__31 - 15)) | (1L << (T__32 - 15)) | (1L << (T__33 - 15)) | (1L << (T__34 - 15)) | (1L << (T__36 - 15)) | (1L << (T__37 - 15)) | (1L << (T__38 - 15)) | (1L << (T__46 - 15)) | (1L << (T__47 - 15)) | (1L << (T__48 - 15)) | (1L << (T__49 - 15)) | (1L << (T__50 - 15)) | (1L << (T__51 - 15)) | (1L << (T__52 - 15)) | (1L << (T__53 - 15)) | (1L << (T__56 - 15)) | (1L << (T__57 - 15)) | (1L << (T__58 - 15)) | (1L << (T__59 - 15)) | (1L << (T__60 - 15)) | (1L << (T__61 - 15)) | (1L << (T__62 - 15)) | (1L << (T__63 - 15)))) != 0) || _la==AT );
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
		public TerminalNode LCURLY() { return getToken(TswiftParser.LCURLY, 0); }
		public TerminalNode RCURLY() { return getToken(TswiftParser.RCURLY, 0); }
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
		enterRule(_localctx, 64, RULE_code_block);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(657);
			match(LCURLY);
			setState(659);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				{
				setState(658);
				instrucciones();
				}
				break;
			}
			setState(661);
			match(RCURLY);
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

	public static class Import_kindContext extends ParserRuleContext {
		public Import_kindContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_import_kind; }
	}

	public final Import_kindContext import_kind() throws RecognitionException {
		Import_kindContext _localctx = new Import_kindContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_import_kind);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(663);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__13) | (1L << T__14) | (1L << T__15) | (1L << T__16) | (1L << T__17) | (1L << T__18) | (1L << T__19))) != 0)) ) {
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

	public static class Import_pathContext extends ParserRuleContext {
		public List<Import_path_identifierContext> import_path_identifier() {
			return getRuleContexts(Import_path_identifierContext.class);
		}
		public Import_path_identifierContext import_path_identifier(int i) {
			return getRuleContext(Import_path_identifierContext.class,i);
		}
		public List<TerminalNode> DOT() { return getTokens(TswiftParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(TswiftParser.DOT, i);
		}
		public Import_pathContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_import_path; }
	}

	public final Import_pathContext import_path() throws RecognitionException {
		Import_pathContext _localctx = new Import_pathContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_import_path);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(665);
			import_path_identifier();
			setState(670);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOT) {
				{
				{
				setState(666);
				match(DOT);
				setState(667);
				import_path_identifier();
				}
				}
				setState(672);
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

	public static class Import_path_identifierContext extends ParserRuleContext {
		public Declaracion_identifierContext declaracion_identifier() {
			return getRuleContext(Declaracion_identifierContext.class,0);
		}
		public Operator_Context operator_() {
			return getRuleContext(Operator_Context.class,0);
		}
		public Import_path_identifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_import_path_identifier; }
	}

	public final Import_path_identifierContext import_path_identifier() throws RecognitionException {
		Import_path_identifierContext _localctx = new Import_path_identifierContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_import_path_identifier);
		try {
			setState(675);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__21:
			case T__22:
			case T__23:
			case T__24:
			case T__27:
			case T__28:
			case T__34:
			case T__36:
			case T__37:
			case T__39:
			case T__40:
			case T__41:
			case T__42:
			case T__43:
			case T__44:
			case T__45:
			case T__46:
			case T__47:
			case T__48:
			case T__49:
			case T__50:
			case T__51:
			case T__53:
			case T__54:
			case T__55:
			case T__56:
			case T__61:
			case T__62:
			case T__63:
			case T__75:
			case T__76:
			case T__77:
			case T__78:
			case T__80:
			case T__91:
			case T__92:
			case T__93:
			case T__94:
			case T__97:
			case T__98:
			case T__99:
			case T__100:
			case T__101:
			case T__102:
			case T__103:
			case T__104:
			case T__105:
			case T__106:
			case T__107:
			case T__108:
			case T__109:
			case T__110:
			case T__111:
			case T__112:
			case Identifier:
				enterOuterAlt(_localctx, 1);
				{
				setState(673);
				declaracion_identifier();
				}
				break;
			case DOT:
			case LT:
			case GT:
			case BANG:
			case QUESTION:
			case AND:
			case SUB:
			case EQUAL:
			case OR:
			case DIV:
			case ADD:
			case MUL:
			case MOD:
			case CARET:
			case TILDE:
			case Operator_head_other:
				enterOuterAlt(_localctx, 2);
				{
				setState(674);
				operator_();
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

	public static class Constant_declaracionContext extends ParserRuleContext {
		public Pattern_initializer_listContext pattern_initializer_list() {
			return getRuleContext(Pattern_initializer_listContext.class,0);
		}
		public Constant_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_constant_declaracion; }
	}

	public final Constant_declaracionContext constant_declaracion() throws RecognitionException {
		Constant_declaracionContext _localctx = new Constant_declaracionContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_constant_declaracion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(677);
			match(T__20);
			setState(678);
			pattern_initializer_list();
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

	public static class Pattern_initializer_listContext extends ParserRuleContext {
		public List<Pattern_initializerContext> pattern_initializer() {
			return getRuleContexts(Pattern_initializerContext.class);
		}
		public Pattern_initializerContext pattern_initializer(int i) {
			return getRuleContext(Pattern_initializerContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(TswiftParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(TswiftParser.COMMA, i);
		}
		public Pattern_initializer_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pattern_initializer_list; }
	}

	public final Pattern_initializer_listContext pattern_initializer_list() throws RecognitionException {
		Pattern_initializer_listContext _localctx = new Pattern_initializer_listContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_pattern_initializer_list);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(680);
			pattern_initializer();
			setState(685);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,22,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(681);
					match(COMMA);
					setState(682);
					pattern_initializer();
					}
					} 
				}
				setState(687);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,22,_ctx);
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

	public static class Pattern_initializerContext extends ParserRuleContext {
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public InitializerContext initializer() {
			return getRuleContext(InitializerContext.class,0);
		}
		public Pattern_initializerContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pattern_initializer; }
	}

	public final Pattern_initializerContext pattern_initializer() throws RecognitionException {
		Pattern_initializerContext _localctx = new Pattern_initializerContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_pattern_initializer);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(688);
			pattern(0);
			setState(690);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
			case 1:
				{
				setState(689);
				initializer();
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

	public static class InitializerContext extends ParserRuleContext {
		public Assignment_operatorContext assignment_operator() {
			return getRuleContext(Assignment_operatorContext.class,0);
		}
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public InitializerContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_initializer; }
	}

	public final InitializerContext initializer() throws RecognitionException {
		InitializerContext _localctx = new InitializerContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_initializer);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(692);
			assignment_operator();
			setState(693);
			expresion();
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
		public Variable_declaracion_headContext variable_declaracion_head() {
			return getRuleContext(Variable_declaracion_headContext.class,0);
		}
		public Variable_nameContext variable_name() {
			return getRuleContext(Variable_nameContext.class,0);
		}
		public List<Type_annotationContext> type_annotation() {
			return getRuleContexts(Type_annotationContext.class);
		}
		public Type_annotationContext type_annotation(int i) {
			return getRuleContext(Type_annotationContext.class,i);
		}
		public Code_blockContext code_block() {
			return getRuleContext(Code_blockContext.class,0);
		}
		public Getter_setter_blockContext getter_setter_block() {
			return getRuleContext(Getter_setter_blockContext.class,0);
		}
		public Getter_setter_keyword_blockContext getter_setter_keyword_block() {
			return getRuleContext(Getter_setter_keyword_blockContext.class,0);
		}
		public WillSet_didSet_blockContext willSet_didSet_block() {
			return getRuleContext(WillSet_didSet_blockContext.class,0);
		}
		public InitializerContext initializer() {
			return getRuleContext(InitializerContext.class,0);
		}
		public Pattern_initializer_listContext pattern_initializer_list() {
			return getRuleContext(Pattern_initializer_listContext.class,0);
		}
		public Variable_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variable_declaracion; }
	}

	public final Variable_declaracionContext variable_declaracion() throws RecognitionException {
		Variable_declaracionContext _localctx = new Variable_declaracionContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_variable_declaracion);
		try {
			setState(730);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(695);
				variable_declaracion_head();
				setState(696);
				variable_name();
				setState(697);
				type_annotation();
				setState(698);
				code_block();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(700);
				variable_declaracion_head();
				setState(701);
				variable_name();
				setState(702);
				type_annotation();
				setState(703);
				getter_setter_block();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(705);
				variable_declaracion_head();
				setState(706);
				variable_name();
				setState(707);
				type_annotation();
				setState(708);
				getter_setter_keyword_block();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(710);
				variable_declaracion_head();
				setState(711);
				variable_name();
				setState(712);
				type_annotation();
				setState(714);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
				case 1:
					{
					setState(713);
					initializer();
					}
					break;
				}
				setState(716);
				willSet_didSet_block();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(718);
				variable_declaracion_head();
				setState(719);
				variable_name();
				setState(720);
				type_annotation();
				setState(721);
				type_annotation();
				setState(723);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
				case 1:
					{
					setState(722);
					initializer();
					}
					break;
				}
				setState(725);
				willSet_didSet_block();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(727);
				variable_declaracion_head();
				setState(728);
				pattern_initializer_list();
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

	public static class Variable_declaracion_headContext extends ParserRuleContext {
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Declaracion_modifiersContext declaracion_modifiers() {
			return getRuleContext(Declaracion_modifiersContext.class,0);
		}
		public Variable_declaracion_headContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variable_declaracion_head; }
	}

	public final Variable_declaracion_headContext variable_declaracion_head() throws RecognitionException {
		Variable_declaracion_headContext _localctx = new Variable_declaracion_headContext(_ctx, getState());
		enterRule(_localctx, 82, RULE_variable_declaracion_head);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(733);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(732);
				attributes();
				}
			}

			setState(736);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 16)) & ~0x3f) == 0 && ((1L << (_la - 16)) & ((1L << (T__15 - 16)) | (1L << (T__28 - 16)) | (1L << (T__34 - 16)) | (1L << (T__36 - 16)) | (1L << (T__37 - 16)) | (1L << (T__46 - 16)) | (1L << (T__47 - 16)) | (1L << (T__48 - 16)) | (1L << (T__49 - 16)) | (1L << (T__50 - 16)) | (1L << (T__51 - 16)) | (1L << (T__52 - 16)) | (1L << (T__53 - 16)) | (1L << (T__56 - 16)) | (1L << (T__57 - 16)) | (1L << (T__58 - 16)) | (1L << (T__59 - 16)) | (1L << (T__60 - 16)) | (1L << (T__61 - 16)) | (1L << (T__62 - 16)) | (1L << (T__63 - 16)))) != 0)) {
				{
				setState(735);
				declaracion_modifiers();
				}
			}

			setState(738);
			match(T__18);
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

	public static class Variable_nameContext extends ParserRuleContext {
		public Declaracion_identifierContext declaracion_identifier() {
			return getRuleContext(Declaracion_identifierContext.class,0);
		}
		public Variable_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variable_name; }
	}

	public final Variable_nameContext variable_name() throws RecognitionException {
		Variable_nameContext _localctx = new Variable_nameContext(_ctx, getState());
		enterRule(_localctx, 84, RULE_variable_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(740);
			declaracion_identifier();
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

	public static class Getter_setter_blockContext extends ParserRuleContext {
		public TerminalNode LCURLY() { return getToken(TswiftParser.LCURLY, 0); }
		public Getter_clauseContext getter_clause() {
			return getRuleContext(Getter_clauseContext.class,0);
		}
		public TerminalNode RCURLY() { return getToken(TswiftParser.RCURLY, 0); }
		public Setter_clauseContext setter_clause() {
			return getRuleContext(Setter_clauseContext.class,0);
		}
		public Getter_setter_blockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_getter_setter_block; }
	}

	public final Getter_setter_blockContext getter_setter_block() throws RecognitionException {
		Getter_setter_blockContext _localctx = new Getter_setter_blockContext(_ctx, getState());
		enterRule(_localctx, 86, RULE_getter_setter_block);
		int _la;
		try {
			setState(754);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,30,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(742);
				match(LCURLY);
				setState(743);
				getter_clause();
				setState(745);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (((((_la - 23)) & ~0x3f) == 0 && ((1L << (_la - 23)) & ((1L << (T__22 - 23)) | (1L << (T__62 - 23)) | (1L << (T__63 - 23)))) != 0) || _la==AT) {
					{
					setState(744);
					setter_clause();
					}
				}

				setState(747);
				match(RCURLY);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(749);
				match(LCURLY);
				setState(750);
				setter_clause();
				setState(751);
				getter_clause();
				setState(752);
				match(RCURLY);
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

	public static class Getter_clauseContext extends ParserRuleContext {
		public Code_blockContext code_block() {
			return getRuleContext(Code_blockContext.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Mutation_modifierContext mutation_modifier() {
			return getRuleContext(Mutation_modifierContext.class,0);
		}
		public Getter_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_getter_clause; }
	}

	public final Getter_clauseContext getter_clause() throws RecognitionException {
		Getter_clauseContext _localctx = new Getter_clauseContext(_ctx, getState());
		enterRule(_localctx, 88, RULE_getter_clause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(757);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(756);
				attributes();
				}
			}

			setState(760);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__62 || _la==T__63) {
				{
				setState(759);
				mutation_modifier();
				}
			}

			setState(762);
			match(T__21);
			setState(763);
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

	public static class Setter_clauseContext extends ParserRuleContext {
		public Code_blockContext code_block() {
			return getRuleContext(Code_blockContext.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Mutation_modifierContext mutation_modifier() {
			return getRuleContext(Mutation_modifierContext.class,0);
		}
		public Setter_nameContext setter_name() {
			return getRuleContext(Setter_nameContext.class,0);
		}
		public Setter_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_setter_clause; }
	}

	public final Setter_clauseContext setter_clause() throws RecognitionException {
		Setter_clauseContext _localctx = new Setter_clauseContext(_ctx, getState());
		enterRule(_localctx, 90, RULE_setter_clause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(766);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(765);
				attributes();
				}
			}

			setState(769);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__62 || _la==T__63) {
				{
				setState(768);
				mutation_modifier();
				}
			}

			setState(771);
			match(T__22);
			setState(773);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LPAREN) {
				{
				setState(772);
				setter_name();
				}
			}

			setState(775);
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

	public static class Setter_nameContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(TswiftParser.LPAREN, 0); }
		public Declaracion_identifierContext declaracion_identifier() {
			return getRuleContext(Declaracion_identifierContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(TswiftParser.RPAREN, 0); }
		public Setter_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_setter_name; }
	}

	public final Setter_nameContext setter_name() throws RecognitionException {
		Setter_nameContext _localctx = new Setter_nameContext(_ctx, getState());
		enterRule(_localctx, 92, RULE_setter_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(777);
			match(LPAREN);
			setState(778);
			declaracion_identifier();
			setState(779);
			match(RPAREN);
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

	public static class Getter_setter_keyword_blockContext extends ParserRuleContext {
		public TerminalNode LCURLY() { return getToken(TswiftParser.LCURLY, 0); }
		public Getter_keyword_clauseContext getter_keyword_clause() {
			return getRuleContext(Getter_keyword_clauseContext.class,0);
		}
		public TerminalNode RCURLY() { return getToken(TswiftParser.RCURLY, 0); }
		public Setter_keyword_clauseContext setter_keyword_clause() {
			return getRuleContext(Setter_keyword_clauseContext.class,0);
		}
		public Getter_setter_keyword_blockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_getter_setter_keyword_block; }
	}

	public final Getter_setter_keyword_blockContext getter_setter_keyword_block() throws RecognitionException {
		Getter_setter_keyword_blockContext _localctx = new Getter_setter_keyword_blockContext(_ctx, getState());
		enterRule(_localctx, 94, RULE_getter_setter_keyword_block);
		int _la;
		try {
			setState(793);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,37,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(781);
				match(LCURLY);
				setState(782);
				getter_keyword_clause();
				setState(784);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (((((_la - 23)) & ~0x3f) == 0 && ((1L << (_la - 23)) & ((1L << (T__22 - 23)) | (1L << (T__62 - 23)) | (1L << (T__63 - 23)))) != 0) || _la==AT) {
					{
					setState(783);
					setter_keyword_clause();
					}
				}

				setState(786);
				match(RCURLY);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(788);
				match(LCURLY);
				setState(789);
				setter_keyword_clause();
				setState(790);
				getter_keyword_clause();
				setState(791);
				match(RCURLY);
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

	public static class Getter_keyword_clauseContext extends ParserRuleContext {
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Mutation_modifierContext mutation_modifier() {
			return getRuleContext(Mutation_modifierContext.class,0);
		}
		public Getter_keyword_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_getter_keyword_clause; }
	}

	public final Getter_keyword_clauseContext getter_keyword_clause() throws RecognitionException {
		Getter_keyword_clauseContext _localctx = new Getter_keyword_clauseContext(_ctx, getState());
		enterRule(_localctx, 96, RULE_getter_keyword_clause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(796);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(795);
				attributes();
				}
			}

			setState(799);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__62 || _la==T__63) {
				{
				setState(798);
				mutation_modifier();
				}
			}

			setState(801);
			match(T__21);
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

	public static class Setter_keyword_clauseContext extends ParserRuleContext {
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Mutation_modifierContext mutation_modifier() {
			return getRuleContext(Mutation_modifierContext.class,0);
		}
		public Setter_keyword_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_setter_keyword_clause; }
	}

	public final Setter_keyword_clauseContext setter_keyword_clause() throws RecognitionException {
		Setter_keyword_clauseContext _localctx = new Setter_keyword_clauseContext(_ctx, getState());
		enterRule(_localctx, 98, RULE_setter_keyword_clause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(804);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(803);
				attributes();
				}
			}

			setState(807);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__62 || _la==T__63) {
				{
				setState(806);
				mutation_modifier();
				}
			}

			setState(809);
			match(T__22);
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

	public static class WillSet_didSet_blockContext extends ParserRuleContext {
		public TerminalNode LCURLY() { return getToken(TswiftParser.LCURLY, 0); }
		public WillSet_clauseContext willSet_clause() {
			return getRuleContext(WillSet_clauseContext.class,0);
		}
		public TerminalNode RCURLY() { return getToken(TswiftParser.RCURLY, 0); }
		public DidSet_clauseContext didSet_clause() {
			return getRuleContext(DidSet_clauseContext.class,0);
		}
		public WillSet_didSet_blockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_willSet_didSet_block; }
	}

	public final WillSet_didSet_blockContext willSet_didSet_block() throws RecognitionException {
		WillSet_didSet_blockContext _localctx = new WillSet_didSet_blockContext(_ctx, getState());
		enterRule(_localctx, 100, RULE_willSet_didSet_block);
		int _la;
		try {
			setState(823);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,43,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(811);
				match(LCURLY);
				setState(812);
				willSet_clause();
				setState(814);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__24 || _la==AT) {
					{
					setState(813);
					didSet_clause();
					}
				}

				setState(816);
				match(RCURLY);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(818);
				match(LCURLY);
				setState(819);
				didSet_clause();
				setState(820);
				willSet_clause();
				setState(821);
				match(RCURLY);
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

	public static class WillSet_clauseContext extends ParserRuleContext {
		public Code_blockContext code_block() {
			return getRuleContext(Code_blockContext.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Setter_nameContext setter_name() {
			return getRuleContext(Setter_nameContext.class,0);
		}
		public WillSet_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_willSet_clause; }
	}

	public final WillSet_clauseContext willSet_clause() throws RecognitionException {
		WillSet_clauseContext _localctx = new WillSet_clauseContext(_ctx, getState());
		enterRule(_localctx, 102, RULE_willSet_clause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(826);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(825);
				attributes();
				}
			}

			setState(828);
			match(T__23);
			setState(830);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LPAREN) {
				{
				setState(829);
				setter_name();
				}
			}

			setState(832);
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

	public static class DidSet_clauseContext extends ParserRuleContext {
		public Code_blockContext code_block() {
			return getRuleContext(Code_blockContext.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Setter_nameContext setter_name() {
			return getRuleContext(Setter_nameContext.class,0);
		}
		public DidSet_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_didSet_clause; }
	}

	public final DidSet_clauseContext didSet_clause() throws RecognitionException {
		DidSet_clauseContext _localctx = new DidSet_clauseContext(_ctx, getState());
		enterRule(_localctx, 104, RULE_didSet_clause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(835);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(834);
				attributes();
				}
			}

			setState(837);
			match(T__24);
			setState(839);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LPAREN) {
				{
				setState(838);
				setter_name();
				}
			}

			setState(841);
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

	public static class Function_declaracionContext extends ParserRuleContext {
		public Function_headContext function_head() {
			return getRuleContext(Function_headContext.class,0);
		}
		public Function_nameContext function_name() {
			return getRuleContext(Function_nameContext.class,0);
		}
		public Function_signatureContext function_signature() {
			return getRuleContext(Function_signatureContext.class,0);
		}
		public Generic_parameter_clauseContext generic_parameter_clause() {
			return getRuleContext(Generic_parameter_clauseContext.class,0);
		}
		public Generic_where_clauseContext generic_where_clause() {
			return getRuleContext(Generic_where_clauseContext.class,0);
		}
		public Function_bodyContext function_body() {
			return getRuleContext(Function_bodyContext.class,0);
		}
		public Function_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_declaracion; }
	}

	public final Function_declaracionContext function_declaracion() throws RecognitionException {
		Function_declaracionContext _localctx = new Function_declaracionContext(_ctx, getState());
		enterRule(_localctx, 106, RULE_function_declaracion);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(843);
			function_head();
			setState(844);
			function_name();
			setState(846);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(845);
				generic_parameter_clause();
				}
			}

			setState(848);
			function_signature();
			setState(850);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,49,_ctx) ) {
			case 1:
				{
				setState(849);
				generic_where_clause();
				}
				break;
			}
			setState(853);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,50,_ctx) ) {
			case 1:
				{
				setState(852);
				function_body();
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

	public static class Function_headContext extends ParserRuleContext {
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Declaracion_modifiersContext declaracion_modifiers() {
			return getRuleContext(Declaracion_modifiersContext.class,0);
		}
		public Function_headContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_head; }
	}

	public final Function_headContext function_head() throws RecognitionException {
		Function_headContext _localctx = new Function_headContext(_ctx, getState());
		enterRule(_localctx, 108, RULE_function_head);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(856);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(855);
				attributes();
				}
			}

			setState(859);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 16)) & ~0x3f) == 0 && ((1L << (_la - 16)) & ((1L << (T__15 - 16)) | (1L << (T__28 - 16)) | (1L << (T__34 - 16)) | (1L << (T__36 - 16)) | (1L << (T__37 - 16)) | (1L << (T__46 - 16)) | (1L << (T__47 - 16)) | (1L << (T__48 - 16)) | (1L << (T__49 - 16)) | (1L << (T__50 - 16)) | (1L << (T__51 - 16)) | (1L << (T__52 - 16)) | (1L << (T__53 - 16)) | (1L << (T__56 - 16)) | (1L << (T__57 - 16)) | (1L << (T__58 - 16)) | (1L << (T__59 - 16)) | (1L << (T__60 - 16)) | (1L << (T__61 - 16)) | (1L << (T__62 - 16)) | (1L << (T__63 - 16)))) != 0)) {
				{
				setState(858);
				declaracion_modifiers();
				}
			}

			setState(861);
			match(T__19);
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

	public static class Function_nameContext extends ParserRuleContext {
		public Declaracion_identifierContext declaracion_identifier() {
			return getRuleContext(Declaracion_identifierContext.class,0);
		}
		public Operator_Context operator_() {
			return getRuleContext(Operator_Context.class,0);
		}
		public Function_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_name; }
	}

	public final Function_nameContext function_name() throws RecognitionException {
		Function_nameContext _localctx = new Function_nameContext(_ctx, getState());
		enterRule(_localctx, 110, RULE_function_name);
		try {
			setState(865);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__21:
			case T__22:
			case T__23:
			case T__24:
			case T__27:
			case T__28:
			case T__34:
			case T__36:
			case T__37:
			case T__39:
			case T__40:
			case T__41:
			case T__42:
			case T__43:
			case T__44:
			case T__45:
			case T__46:
			case T__47:
			case T__48:
			case T__49:
			case T__50:
			case T__51:
			case T__53:
			case T__54:
			case T__55:
			case T__56:
			case T__61:
			case T__62:
			case T__63:
			case T__75:
			case T__76:
			case T__77:
			case T__78:
			case T__80:
			case T__91:
			case T__92:
			case T__93:
			case T__94:
			case T__97:
			case T__98:
			case T__99:
			case T__100:
			case T__101:
			case T__102:
			case T__103:
			case T__104:
			case T__105:
			case T__106:
			case T__107:
			case T__108:
			case T__109:
			case T__110:
			case T__111:
			case T__112:
			case Identifier:
				enterOuterAlt(_localctx, 1);
				{
				setState(863);
				declaracion_identifier();
				}
				break;
			case DOT:
			case LT:
			case GT:
			case BANG:
			case QUESTION:
			case AND:
			case SUB:
			case EQUAL:
			case OR:
			case DIV:
			case ADD:
			case MUL:
			case MOD:
			case CARET:
			case TILDE:
			case Operator_head_other:
				enterOuterAlt(_localctx, 2);
				{
				setState(864);
				operator_();
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

	public static class Function_signatureContext extends ParserRuleContext {
		public Parameter_clauseContext parameter_clause() {
			return getRuleContext(Parameter_clauseContext.class,0);
		}
		public Function_resultContext function_result() {
			return getRuleContext(Function_resultContext.class,0);
		}
		public Function_signatureContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_signature; }
	}

	public final Function_signatureContext function_signature() throws RecognitionException {
		Function_signatureContext _localctx = new Function_signatureContext(_ctx, getState());
		enterRule(_localctx, 112, RULE_function_signature);
		try {
			setState(879);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,57,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(867);
				parameter_clause();
				setState(869);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,54,_ctx) ) {
				case 1:
					{
					setState(868);
					match(T__25);
					}
					break;
				}
				setState(872);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,55,_ctx) ) {
				case 1:
					{
					setState(871);
					function_result();
					}
					break;
				}
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(874);
				parameter_clause();
				setState(875);
				match(T__26);
				setState(877);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,56,_ctx) ) {
				case 1:
					{
					setState(876);
					function_result();
					}
					break;
				}
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

	public static class Function_resultContext extends ParserRuleContext {
		public Arrow_operatorContext arrow_operator() {
			return getRuleContext(Arrow_operatorContext.class,0);
		}
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Function_resultContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_result; }
	}

	public final Function_resultContext function_result() throws RecognitionException {
		Function_resultContext _localctx = new Function_resultContext(_ctx, getState());
		enterRule(_localctx, 114, RULE_function_result);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(881);
			arrow_operator();
			setState(883);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,58,_ctx) ) {
			case 1:
				{
				setState(882);
				attributes();
				}
				break;
			}
			setState(885);
			type_(0);
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

	public static class Function_bodyContext extends ParserRuleContext {
		public Code_blockContext code_block() {
			return getRuleContext(Code_blockContext.class,0);
		}
		public Function_bodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_body; }
	}

	public final Function_bodyContext function_body() throws RecognitionException {
		Function_bodyContext _localctx = new Function_bodyContext(_ctx, getState());
		enterRule(_localctx, 116, RULE_function_body);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(887);
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

	public static class Parameter_clauseContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(TswiftParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(TswiftParser.RPAREN, 0); }
		public Parameter_listContext parameter_list() {
			return getRuleContext(Parameter_listContext.class,0);
		}
		public Parameter_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameter_clause; }
	}

	public final Parameter_clauseContext parameter_clause() throws RecognitionException {
		Parameter_clauseContext _localctx = new Parameter_clauseContext(_ctx, getState());
		enterRule(_localctx, 118, RULE_parameter_clause);
		try {
			setState(895);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,59,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(889);
				match(LPAREN);
				setState(890);
				match(RPAREN);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(891);
				match(LPAREN);
				setState(892);
				parameter_list();
				setState(893);
				match(RPAREN);
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

	public static class Parameter_listContext extends ParserRuleContext {
		public List<ParameterContext> parameter() {
			return getRuleContexts(ParameterContext.class);
		}
		public ParameterContext parameter(int i) {
			return getRuleContext(ParameterContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(TswiftParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(TswiftParser.COMMA, i);
		}
		public Parameter_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameter_list; }
	}

	public final Parameter_listContext parameter_list() throws RecognitionException {
		Parameter_listContext _localctx = new Parameter_listContext(_ctx, getState());
		enterRule(_localctx, 120, RULE_parameter_list);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(897);
			parameter();
			setState(902);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(898);
				match(COMMA);
				setState(899);
				parameter();
				}
				}
				setState(904);
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

	public static class ParameterContext extends ParserRuleContext {
		public Local_parameter_nameContext local_parameter_name() {
			return getRuleContext(Local_parameter_nameContext.class,0);
		}
		public Type_annotationContext type_annotation() {
			return getRuleContext(Type_annotationContext.class,0);
		}
		public External_parameter_nameContext external_parameter_name() {
			return getRuleContext(External_parameter_nameContext.class,0);
		}
		public Default_argument_clauseContext default_argument_clause() {
			return getRuleContext(Default_argument_clauseContext.class,0);
		}
		public Range_operatorContext range_operator() {
			return getRuleContext(Range_operatorContext.class,0);
		}
		public ParameterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameter; }
	}

	public final ParameterContext parameter() throws RecognitionException {
		ParameterContext _localctx = new ParameterContext(_ctx, getState());
		enterRule(_localctx, 122, RULE_parameter);
		try {
			setState(926);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,65,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(906);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,61,_ctx) ) {
				case 1:
					{
					setState(905);
					external_parameter_name();
					}
					break;
				}
				setState(908);
				local_parameter_name();
				setState(909);
				type_annotation();
				setState(911);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,62,_ctx) ) {
				case 1:
					{
					setState(910);
					default_argument_clause();
					}
					break;
				}
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(914);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,63,_ctx) ) {
				case 1:
					{
					setState(913);
					external_parameter_name();
					}
					break;
				}
				setState(916);
				local_parameter_name();
				setState(917);
				type_annotation();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(920);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,64,_ctx) ) {
				case 1:
					{
					setState(919);
					external_parameter_name();
					}
					break;
				}
				setState(922);
				local_parameter_name();
				setState(923);
				type_annotation();
				setState(924);
				range_operator();
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

	public static class External_parameter_nameContext extends ParserRuleContext {
		public Label_identifierContext label_identifier() {
			return getRuleContext(Label_identifierContext.class,0);
		}
		public External_parameter_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_external_parameter_name; }
	}

	public final External_parameter_nameContext external_parameter_name() throws RecognitionException {
		External_parameter_nameContext _localctx = new External_parameter_nameContext(_ctx, getState());
		enterRule(_localctx, 124, RULE_external_parameter_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(928);
			label_identifier();
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

	public static class Local_parameter_nameContext extends ParserRuleContext {
		public Label_identifierContext label_identifier() {
			return getRuleContext(Label_identifierContext.class,0);
		}
		public Local_parameter_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_local_parameter_name; }
	}

	public final Local_parameter_nameContext local_parameter_name() throws RecognitionException {
		Local_parameter_nameContext _localctx = new Local_parameter_nameContext(_ctx, getState());
		enterRule(_localctx, 126, RULE_local_parameter_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(930);
			label_identifier();
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

	public static class Default_argument_clauseContext extends ParserRuleContext {
		public Assignment_operatorContext assignment_operator() {
			return getRuleContext(Assignment_operatorContext.class,0);
		}
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Default_argument_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_default_argument_clause; }
	}

	public final Default_argument_clauseContext default_argument_clause() throws RecognitionException {
		Default_argument_clauseContext _localctx = new Default_argument_clauseContext(_ctx, getState());
		enterRule(_localctx, 128, RULE_default_argument_clause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(932);
			assignment_operator();
			setState(933);
			expresion();
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

	public static class Enum_declaracionContext extends ParserRuleContext {
		public Union_style_enumContext union_style_enum() {
			return getRuleContext(Union_style_enumContext.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Access_level_modifierContext access_level_modifier() {
			return getRuleContext(Access_level_modifierContext.class,0);
		}
		public Raw_value_style_enumContext raw_value_style_enum() {
			return getRuleContext(Raw_value_style_enumContext.class,0);
		}
		public Enum_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enum_declaracion; }
	}

	public final Enum_declaracionContext enum_declaracion() throws RecognitionException {
		Enum_declaracionContext _localctx = new Enum_declaracionContext(_ctx, getState());
		enterRule(_localctx, 130, RULE_enum_declaracion);
		int _la;
		try {
			setState(949);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,70,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(936);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AT) {
					{
					setState(935);
					attributes();
					}
				}

				setState(939);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__57) | (1L << T__58) | (1L << T__59) | (1L << T__60) | (1L << T__61))) != 0)) {
					{
					setState(938);
					access_level_modifier();
					}
				}

				setState(941);
				union_style_enum();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(943);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AT) {
					{
					setState(942);
					attributes();
					}
				}

				setState(946);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__57) | (1L << T__58) | (1L << T__59) | (1L << T__60) | (1L << T__61))) != 0)) {
					{
					setState(945);
					access_level_modifier();
					}
				}

				setState(948);
				raw_value_style_enum();
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

	public static class Union_style_enumContext extends ParserRuleContext {
		public Enum_nameContext enum_name() {
			return getRuleContext(Enum_nameContext.class,0);
		}
		public TerminalNode LCURLY() { return getToken(TswiftParser.LCURLY, 0); }
		public TerminalNode RCURLY() { return getToken(TswiftParser.RCURLY, 0); }
		public Generic_parameter_clauseContext generic_parameter_clause() {
			return getRuleContext(Generic_parameter_clauseContext.class,0);
		}
		public Type_inheritance_clauseContext type_inheritance_clause() {
			return getRuleContext(Type_inheritance_clauseContext.class,0);
		}
		public Generic_where_clauseContext generic_where_clause() {
			return getRuleContext(Generic_where_clauseContext.class,0);
		}
		public Union_style_enum_membersContext union_style_enum_members() {
			return getRuleContext(Union_style_enum_membersContext.class,0);
		}
		public Union_style_enumContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_union_style_enum; }
	}

	public final Union_style_enumContext union_style_enum() throws RecognitionException {
		Union_style_enumContext _localctx = new Union_style_enumContext(_ctx, getState());
		enterRule(_localctx, 132, RULE_union_style_enum);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(952);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__27) {
				{
				setState(951);
				match(T__27);
				}
			}

			setState(954);
			match(T__16);
			setState(955);
			enum_name();
			setState(957);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(956);
				generic_parameter_clause();
				}
			}

			setState(960);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COLON) {
				{
				setState(959);
				type_inheritance_clause();
				}
			}

			setState(963);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__12) {
				{
				setState(962);
				generic_where_clause();
				}
			}

			setState(965);
			match(LCURLY);
			setState(967);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 8)) & ~0x3f) == 0 && ((1L << (_la - 8)) & ((1L << (T__7 - 8)) | (1L << (T__14 - 8)) | (1L << (T__15 - 8)) | (1L << (T__16 - 8)) | (1L << (T__17 - 8)) | (1L << (T__18 - 8)) | (1L << (T__19 - 8)) | (1L << (T__20 - 8)) | (1L << (T__27 - 8)) | (1L << (T__28 - 8)) | (1L << (T__30 - 8)) | (1L << (T__31 - 8)) | (1L << (T__32 - 8)) | (1L << (T__33 - 8)) | (1L << (T__34 - 8)) | (1L << (T__36 - 8)) | (1L << (T__37 - 8)) | (1L << (T__38 - 8)) | (1L << (T__46 - 8)) | (1L << (T__47 - 8)) | (1L << (T__48 - 8)) | (1L << (T__49 - 8)) | (1L << (T__50 - 8)) | (1L << (T__51 - 8)) | (1L << (T__52 - 8)) | (1L << (T__53 - 8)) | (1L << (T__56 - 8)) | (1L << (T__57 - 8)) | (1L << (T__58 - 8)) | (1L << (T__59 - 8)) | (1L << (T__60 - 8)) | (1L << (T__61 - 8)) | (1L << (T__62 - 8)) | (1L << (T__63 - 8)))) != 0) || _la==AT) {
				{
				setState(966);
				union_style_enum_members();
				}
			}

			setState(969);
			match(RCURLY);
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

	public static class Union_style_enum_membersContext extends ParserRuleContext {
		public Union_style_enum_memberContext union_style_enum_member() {
			return getRuleContext(Union_style_enum_memberContext.class,0);
		}
		public Union_style_enum_membersContext union_style_enum_members() {
			return getRuleContext(Union_style_enum_membersContext.class,0);
		}
		public Union_style_enum_membersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_union_style_enum_members; }
	}

	public final Union_style_enum_membersContext union_style_enum_members() throws RecognitionException {
		Union_style_enum_membersContext _localctx = new Union_style_enum_membersContext(_ctx, getState());
		enterRule(_localctx, 134, RULE_union_style_enum_members);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(971);
			union_style_enum_member();
			setState(973);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 8)) & ~0x3f) == 0 && ((1L << (_la - 8)) & ((1L << (T__7 - 8)) | (1L << (T__14 - 8)) | (1L << (T__15 - 8)) | (1L << (T__16 - 8)) | (1L << (T__17 - 8)) | (1L << (T__18 - 8)) | (1L << (T__19 - 8)) | (1L << (T__20 - 8)) | (1L << (T__27 - 8)) | (1L << (T__28 - 8)) | (1L << (T__30 - 8)) | (1L << (T__31 - 8)) | (1L << (T__32 - 8)) | (1L << (T__33 - 8)) | (1L << (T__34 - 8)) | (1L << (T__36 - 8)) | (1L << (T__37 - 8)) | (1L << (T__38 - 8)) | (1L << (T__46 - 8)) | (1L << (T__47 - 8)) | (1L << (T__48 - 8)) | (1L << (T__49 - 8)) | (1L << (T__50 - 8)) | (1L << (T__51 - 8)) | (1L << (T__52 - 8)) | (1L << (T__53 - 8)) | (1L << (T__56 - 8)) | (1L << (T__57 - 8)) | (1L << (T__58 - 8)) | (1L << (T__59 - 8)) | (1L << (T__60 - 8)) | (1L << (T__61 - 8)) | (1L << (T__62 - 8)) | (1L << (T__63 - 8)))) != 0) || _la==AT) {
				{
				setState(972);
				union_style_enum_members();
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

	public static class Union_style_enum_memberContext extends ParserRuleContext {
		public DeclaracionContext declaracion() {
			return getRuleContext(DeclaracionContext.class,0);
		}
		public Union_style_enum_case_clauseContext union_style_enum_case_clause() {
			return getRuleContext(Union_style_enum_case_clauseContext.class,0);
		}
		public Union_style_enum_memberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_union_style_enum_member; }
	}

	public final Union_style_enum_memberContext union_style_enum_member() throws RecognitionException {
		Union_style_enum_memberContext _localctx = new Union_style_enum_memberContext(_ctx, getState());
		enterRule(_localctx, 136, RULE_union_style_enum_member);
		try {
			setState(977);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,77,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(975);
				declaracion();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(976);
				union_style_enum_case_clause();
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

	public static class Union_style_enum_case_clauseContext extends ParserRuleContext {
		public Union_style_enum_case_listContext union_style_enum_case_list() {
			return getRuleContext(Union_style_enum_case_listContext.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Union_style_enum_case_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_union_style_enum_case_clause; }
	}

	public final Union_style_enum_case_clauseContext union_style_enum_case_clause() throws RecognitionException {
		Union_style_enum_case_clauseContext _localctx = new Union_style_enum_case_clauseContext(_ctx, getState());
		enterRule(_localctx, 138, RULE_union_style_enum_case_clause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(980);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(979);
				attributes();
				}
			}

			setState(983);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__27) {
				{
				setState(982);
				match(T__27);
				}
			}

			setState(985);
			match(T__7);
			setState(986);
			union_style_enum_case_list();
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

	public static class Union_style_enum_case_listContext extends ParserRuleContext {
		public Union_style_enum_caseContext union_style_enum_case() {
			return getRuleContext(Union_style_enum_caseContext.class,0);
		}
		public TerminalNode COMMA() { return getToken(TswiftParser.COMMA, 0); }
		public Union_style_enum_case_listContext union_style_enum_case_list() {
			return getRuleContext(Union_style_enum_case_listContext.class,0);
		}
		public Union_style_enum_case_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_union_style_enum_case_list; }
	}

	public final Union_style_enum_case_listContext union_style_enum_case_list() throws RecognitionException {
		Union_style_enum_case_listContext _localctx = new Union_style_enum_case_listContext(_ctx, getState());
		enterRule(_localctx, 140, RULE_union_style_enum_case_list);
		try {
			setState(993);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,80,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(988);
				union_style_enum_case();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(989);
				union_style_enum_case();
				setState(990);
				match(COMMA);
				setState(991);
				union_style_enum_case_list();
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

	public static class Union_style_enum_caseContext extends ParserRuleContext {
		public Enum_case_nameContext enum_case_name() {
			return getRuleContext(Enum_case_nameContext.class,0);
		}
		public Tuple_typeContext tuple_type() {
			return getRuleContext(Tuple_typeContext.class,0);
		}
		public Union_style_enum_caseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_union_style_enum_case; }
	}

	public final Union_style_enum_caseContext union_style_enum_case() throws RecognitionException {
		Union_style_enum_caseContext _localctx = new Union_style_enum_caseContext(_ctx, getState());
		enterRule(_localctx, 142, RULE_union_style_enum_case);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(995);
			enum_case_name();
			setState(997);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LPAREN) {
				{
				setState(996);
				tuple_type();
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

	public static class Enum_nameContext extends ParserRuleContext {
		public Declaracion_identifierContext declaracion_identifier() {
			return getRuleContext(Declaracion_identifierContext.class,0);
		}
		public Enum_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enum_name; }
	}

	public final Enum_nameContext enum_name() throws RecognitionException {
		Enum_nameContext _localctx = new Enum_nameContext(_ctx, getState());
		enterRule(_localctx, 144, RULE_enum_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(999);
			declaracion_identifier();
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

	public static class Enum_case_nameContext extends ParserRuleContext {
		public Declaracion_identifierContext declaracion_identifier() {
			return getRuleContext(Declaracion_identifierContext.class,0);
		}
		public Enum_case_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enum_case_name; }
	}

	public final Enum_case_nameContext enum_case_name() throws RecognitionException {
		Enum_case_nameContext _localctx = new Enum_case_nameContext(_ctx, getState());
		enterRule(_localctx, 146, RULE_enum_case_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1001);
			declaracion_identifier();
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

	public static class Raw_value_style_enumContext extends ParserRuleContext {
		public Enum_nameContext enum_name() {
			return getRuleContext(Enum_nameContext.class,0);
		}
		public Type_inheritance_clauseContext type_inheritance_clause() {
			return getRuleContext(Type_inheritance_clauseContext.class,0);
		}
		public TerminalNode LCURLY() { return getToken(TswiftParser.LCURLY, 0); }
		public Raw_value_style_enum_membersContext raw_value_style_enum_members() {
			return getRuleContext(Raw_value_style_enum_membersContext.class,0);
		}
		public TerminalNode RCURLY() { return getToken(TswiftParser.RCURLY, 0); }
		public Generic_parameter_clauseContext generic_parameter_clause() {
			return getRuleContext(Generic_parameter_clauseContext.class,0);
		}
		public Generic_where_clauseContext generic_where_clause() {
			return getRuleContext(Generic_where_clauseContext.class,0);
		}
		public Raw_value_style_enumContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_raw_value_style_enum; }
	}

	public final Raw_value_style_enumContext raw_value_style_enum() throws RecognitionException {
		Raw_value_style_enumContext _localctx = new Raw_value_style_enumContext(_ctx, getState());
		enterRule(_localctx, 148, RULE_raw_value_style_enum);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1003);
			match(T__16);
			setState(1004);
			enum_name();
			setState(1006);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(1005);
				generic_parameter_clause();
				}
			}

			setState(1008);
			type_inheritance_clause();
			setState(1010);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__12) {
				{
				setState(1009);
				generic_where_clause();
				}
			}

			setState(1012);
			match(LCURLY);
			setState(1013);
			raw_value_style_enum_members();
			setState(1014);
			match(RCURLY);
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

	public static class Raw_value_style_enum_membersContext extends ParserRuleContext {
		public Raw_value_style_enum_memberContext raw_value_style_enum_member() {
			return getRuleContext(Raw_value_style_enum_memberContext.class,0);
		}
		public Raw_value_style_enum_membersContext raw_value_style_enum_members() {
			return getRuleContext(Raw_value_style_enum_membersContext.class,0);
		}
		public Raw_value_style_enum_membersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_raw_value_style_enum_members; }
	}

	public final Raw_value_style_enum_membersContext raw_value_style_enum_members() throws RecognitionException {
		Raw_value_style_enum_membersContext _localctx = new Raw_value_style_enum_membersContext(_ctx, getState());
		enterRule(_localctx, 150, RULE_raw_value_style_enum_members);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1016);
			raw_value_style_enum_member();
			setState(1018);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 8)) & ~0x3f) == 0 && ((1L << (_la - 8)) & ((1L << (T__7 - 8)) | (1L << (T__14 - 8)) | (1L << (T__15 - 8)) | (1L << (T__16 - 8)) | (1L << (T__17 - 8)) | (1L << (T__18 - 8)) | (1L << (T__19 - 8)) | (1L << (T__20 - 8)) | (1L << (T__27 - 8)) | (1L << (T__28 - 8)) | (1L << (T__30 - 8)) | (1L << (T__31 - 8)) | (1L << (T__32 - 8)) | (1L << (T__33 - 8)) | (1L << (T__34 - 8)) | (1L << (T__36 - 8)) | (1L << (T__37 - 8)) | (1L << (T__38 - 8)) | (1L << (T__46 - 8)) | (1L << (T__47 - 8)) | (1L << (T__48 - 8)) | (1L << (T__49 - 8)) | (1L << (T__50 - 8)) | (1L << (T__51 - 8)) | (1L << (T__52 - 8)) | (1L << (T__53 - 8)) | (1L << (T__56 - 8)) | (1L << (T__57 - 8)) | (1L << (T__58 - 8)) | (1L << (T__59 - 8)) | (1L << (T__60 - 8)) | (1L << (T__61 - 8)) | (1L << (T__62 - 8)) | (1L << (T__63 - 8)))) != 0) || _la==AT) {
				{
				setState(1017);
				raw_value_style_enum_members();
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

	public static class Raw_value_style_enum_memberContext extends ParserRuleContext {
		public DeclaracionContext declaracion() {
			return getRuleContext(DeclaracionContext.class,0);
		}
		public Raw_value_style_enum_case_clauseContext raw_value_style_enum_case_clause() {
			return getRuleContext(Raw_value_style_enum_case_clauseContext.class,0);
		}
		public Raw_value_style_enum_memberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_raw_value_style_enum_member; }
	}

	public final Raw_value_style_enum_memberContext raw_value_style_enum_member() throws RecognitionException {
		Raw_value_style_enum_memberContext _localctx = new Raw_value_style_enum_memberContext(_ctx, getState());
		enterRule(_localctx, 152, RULE_raw_value_style_enum_member);
		try {
			setState(1022);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,85,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1020);
				declaracion();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1021);
				raw_value_style_enum_case_clause();
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

	public static class Raw_value_style_enum_case_clauseContext extends ParserRuleContext {
		public Raw_value_style_enum_case_listContext raw_value_style_enum_case_list() {
			return getRuleContext(Raw_value_style_enum_case_listContext.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Raw_value_style_enum_case_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_raw_value_style_enum_case_clause; }
	}

	public final Raw_value_style_enum_case_clauseContext raw_value_style_enum_case_clause() throws RecognitionException {
		Raw_value_style_enum_case_clauseContext _localctx = new Raw_value_style_enum_case_clauseContext(_ctx, getState());
		enterRule(_localctx, 154, RULE_raw_value_style_enum_case_clause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1025);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(1024);
				attributes();
				}
			}

			setState(1027);
			match(T__7);
			setState(1028);
			raw_value_style_enum_case_list();
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

	public static class Raw_value_style_enum_case_listContext extends ParserRuleContext {
		public Raw_value_style_enum_caseContext raw_value_style_enum_case() {
			return getRuleContext(Raw_value_style_enum_caseContext.class,0);
		}
		public TerminalNode COMMA() { return getToken(TswiftParser.COMMA, 0); }
		public Raw_value_style_enum_case_listContext raw_value_style_enum_case_list() {
			return getRuleContext(Raw_value_style_enum_case_listContext.class,0);
		}
		public Raw_value_style_enum_case_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_raw_value_style_enum_case_list; }
	}

	public final Raw_value_style_enum_case_listContext raw_value_style_enum_case_list() throws RecognitionException {
		Raw_value_style_enum_case_listContext _localctx = new Raw_value_style_enum_case_listContext(_ctx, getState());
		enterRule(_localctx, 156, RULE_raw_value_style_enum_case_list);
		try {
			setState(1035);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,87,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1030);
				raw_value_style_enum_case();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1031);
				raw_value_style_enum_case();
				setState(1032);
				match(COMMA);
				setState(1033);
				raw_value_style_enum_case_list();
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

	public static class Raw_value_style_enum_caseContext extends ParserRuleContext {
		public Enum_case_nameContext enum_case_name() {
			return getRuleContext(Enum_case_nameContext.class,0);
		}
		public Raw_value_assignmentContext raw_value_assignment() {
			return getRuleContext(Raw_value_assignmentContext.class,0);
		}
		public Raw_value_style_enum_caseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_raw_value_style_enum_case; }
	}

	public final Raw_value_style_enum_caseContext raw_value_style_enum_case() throws RecognitionException {
		Raw_value_style_enum_caseContext _localctx = new Raw_value_style_enum_caseContext(_ctx, getState());
		enterRule(_localctx, 158, RULE_raw_value_style_enum_case);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1037);
			enum_case_name();
			setState(1039);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,88,_ctx) ) {
			case 1:
				{
				setState(1038);
				raw_value_assignment();
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

	public static class Raw_value_assignmentContext extends ParserRuleContext {
		public Assignment_operatorContext assignment_operator() {
			return getRuleContext(Assignment_operatorContext.class,0);
		}
		public Raw_value_literalContext raw_value_literal() {
			return getRuleContext(Raw_value_literalContext.class,0);
		}
		public Raw_value_assignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_raw_value_assignment; }
	}

	public final Raw_value_assignmentContext raw_value_assignment() throws RecognitionException {
		Raw_value_assignmentContext _localctx = new Raw_value_assignmentContext(_ctx, getState());
		enterRule(_localctx, 160, RULE_raw_value_assignment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1041);
			assignment_operator();
			setState(1042);
			raw_value_literal();
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

	public static class Raw_value_literalContext extends ParserRuleContext {
		public Numeric_literalContext numeric_literal() {
			return getRuleContext(Numeric_literalContext.class,0);
		}
		public TerminalNode Static_string_literal() { return getToken(TswiftParser.Static_string_literal, 0); }
		public Boolean_literalContext boolean_literal() {
			return getRuleContext(Boolean_literalContext.class,0);
		}
		public Raw_value_literalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_raw_value_literal; }
	}

	public final Raw_value_literalContext raw_value_literal() throws RecognitionException {
		Raw_value_literalContext _localctx = new Raw_value_literalContext(_ctx, getState());
		enterRule(_localctx, 162, RULE_raw_value_literal);
		try {
			setState(1047);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,89,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1044);
				numeric_literal();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1045);
				match(Static_string_literal);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1046);
				boolean_literal();
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

	public static class Struct_declaracionContext extends ParserRuleContext {
		public Struct_nameContext struct_name() {
			return getRuleContext(Struct_nameContext.class,0);
		}
		public Struct_bodyContext struct_body() {
			return getRuleContext(Struct_bodyContext.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Access_level_modifierContext access_level_modifier() {
			return getRuleContext(Access_level_modifierContext.class,0);
		}
		public Generic_parameter_clauseContext generic_parameter_clause() {
			return getRuleContext(Generic_parameter_clauseContext.class,0);
		}
		public Type_inheritance_clauseContext type_inheritance_clause() {
			return getRuleContext(Type_inheritance_clauseContext.class,0);
		}
		public Generic_where_clauseContext generic_where_clause() {
			return getRuleContext(Generic_where_clauseContext.class,0);
		}
		public Struct_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_struct_declaracion; }
	}

	public final Struct_declaracionContext struct_declaracion() throws RecognitionException {
		Struct_declaracionContext _localctx = new Struct_declaracionContext(_ctx, getState());
		enterRule(_localctx, 164, RULE_struct_declaracion);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1050);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(1049);
				attributes();
				}
			}

			setState(1053);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__57) | (1L << T__58) | (1L << T__59) | (1L << T__60) | (1L << T__61))) != 0)) {
				{
				setState(1052);
				access_level_modifier();
				}
			}

			setState(1055);
			match(T__14);
			setState(1056);
			struct_name();
			setState(1058);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(1057);
				generic_parameter_clause();
				}
			}

			setState(1061);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COLON) {
				{
				setState(1060);
				type_inheritance_clause();
				}
			}

			setState(1064);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__12) {
				{
				setState(1063);
				generic_where_clause();
				}
			}

			setState(1066);
			struct_body();
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

	public static class Struct_nameContext extends ParserRuleContext {
		public Declaracion_identifierContext declaracion_identifier() {
			return getRuleContext(Declaracion_identifierContext.class,0);
		}
		public Struct_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_struct_name; }
	}

	public final Struct_nameContext struct_name() throws RecognitionException {
		Struct_nameContext _localctx = new Struct_nameContext(_ctx, getState());
		enterRule(_localctx, 166, RULE_struct_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1068);
			declaracion_identifier();
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

	public static class Struct_bodyContext extends ParserRuleContext {
		public TerminalNode LCURLY() { return getToken(TswiftParser.LCURLY, 0); }
		public TerminalNode RCURLY() { return getToken(TswiftParser.RCURLY, 0); }
		public List<Struct_memberContext> struct_member() {
			return getRuleContexts(Struct_memberContext.class);
		}
		public Struct_memberContext struct_member(int i) {
			return getRuleContext(Struct_memberContext.class,i);
		}
		public Struct_bodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_struct_body; }
	}

	public final Struct_bodyContext struct_body() throws RecognitionException {
		Struct_bodyContext _localctx = new Struct_bodyContext(_ctx, getState());
		enterRule(_localctx, 168, RULE_struct_body);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1070);
			match(LCURLY);
			setState(1074);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & ((1L << (T__14 - 15)) | (1L << (T__15 - 15)) | (1L << (T__16 - 15)) | (1L << (T__17 - 15)) | (1L << (T__18 - 15)) | (1L << (T__19 - 15)) | (1L << (T__20 - 15)) | (1L << (T__27 - 15)) | (1L << (T__28 - 15)) | (1L << (T__30 - 15)) | (1L << (T__31 - 15)) | (1L << (T__32 - 15)) | (1L << (T__33 - 15)) | (1L << (T__34 - 15)) | (1L << (T__36 - 15)) | (1L << (T__37 - 15)) | (1L << (T__38 - 15)) | (1L << (T__46 - 15)) | (1L << (T__47 - 15)) | (1L << (T__48 - 15)) | (1L << (T__49 - 15)) | (1L << (T__50 - 15)) | (1L << (T__51 - 15)) | (1L << (T__52 - 15)) | (1L << (T__53 - 15)) | (1L << (T__56 - 15)) | (1L << (T__57 - 15)) | (1L << (T__58 - 15)) | (1L << (T__59 - 15)) | (1L << (T__60 - 15)) | (1L << (T__61 - 15)) | (1L << (T__62 - 15)) | (1L << (T__63 - 15)))) != 0) || _la==AT) {
				{
				{
				setState(1071);
				struct_member();
				}
				}
				setState(1076);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1077);
			match(RCURLY);
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

	public static class Struct_memberContext extends ParserRuleContext {
		public DeclaracionContext declaracion() {
			return getRuleContext(DeclaracionContext.class,0);
		}
		public Struct_memberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_struct_member; }
	}

	public final Struct_memberContext struct_member() throws RecognitionException {
		Struct_memberContext _localctx = new Struct_memberContext(_ctx, getState());
		enterRule(_localctx, 170, RULE_struct_member);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1079);
			declaracion();
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

	public static class Class_declaracionContext extends ParserRuleContext {
		public Class_nameContext class_name() {
			return getRuleContext(Class_nameContext.class,0);
		}
		public Class_bodyContext class_body() {
			return getRuleContext(Class_bodyContext.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public List<Access_level_modifierContext> access_level_modifier() {
			return getRuleContexts(Access_level_modifierContext.class);
		}
		public Access_level_modifierContext access_level_modifier(int i) {
			return getRuleContext(Access_level_modifierContext.class,i);
		}
		public Generic_parameter_clauseContext generic_parameter_clause() {
			return getRuleContext(Generic_parameter_clauseContext.class,0);
		}
		public Type_inheritance_clauseContext type_inheritance_clause() {
			return getRuleContext(Type_inheritance_clauseContext.class,0);
		}
		public Generic_where_clauseContext generic_where_clause() {
			return getRuleContext(Generic_where_clauseContext.class,0);
		}
		public Class_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_class_declaracion; }
	}

	public final Class_declaracionContext class_declaracion() throws RecognitionException {
		Class_declaracionContext _localctx = new Class_declaracionContext(_ctx, getState());
		enterRule(_localctx, 172, RULE_class_declaracion);
		int _la;
		try {
			setState(1126);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,108,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1082);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AT) {
					{
					setState(1081);
					attributes();
					}
				}

				setState(1085);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__57) | (1L << T__58) | (1L << T__59) | (1L << T__60) | (1L << T__61))) != 0)) {
					{
					setState(1084);
					access_level_modifier();
					}
				}

				setState(1088);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__28) {
					{
					setState(1087);
					match(T__28);
					}
				}

				setState(1090);
				match(T__15);
				setState(1091);
				class_name();
				setState(1093);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LT) {
					{
					setState(1092);
					generic_parameter_clause();
					}
				}

				setState(1096);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COLON) {
					{
					setState(1095);
					type_inheritance_clause();
					}
				}

				setState(1099);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__12) {
					{
					setState(1098);
					generic_where_clause();
					}
				}

				setState(1101);
				class_body();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1104);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AT) {
					{
					setState(1103);
					attributes();
					}
				}

				setState(1107);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__57) | (1L << T__58) | (1L << T__59) | (1L << T__60) | (1L << T__61))) != 0)) {
					{
					setState(1106);
					access_level_modifier();
					}
				}

				setState(1109);
				match(T__28);
				setState(1111);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__57) | (1L << T__58) | (1L << T__59) | (1L << T__60) | (1L << T__61))) != 0)) {
					{
					setState(1110);
					access_level_modifier();
					}
				}

				setState(1113);
				match(T__15);
				setState(1114);
				class_name();
				setState(1116);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LT) {
					{
					setState(1115);
					generic_parameter_clause();
					}
				}

				setState(1119);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COLON) {
					{
					setState(1118);
					type_inheritance_clause();
					}
				}

				setState(1122);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__12) {
					{
					setState(1121);
					generic_where_clause();
					}
				}

				setState(1124);
				class_body();
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

	public static class Class_nameContext extends ParserRuleContext {
		public Declaracion_identifierContext declaracion_identifier() {
			return getRuleContext(Declaracion_identifierContext.class,0);
		}
		public Class_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_class_name; }
	}

	public final Class_nameContext class_name() throws RecognitionException {
		Class_nameContext _localctx = new Class_nameContext(_ctx, getState());
		enterRule(_localctx, 174, RULE_class_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1128);
			declaracion_identifier();
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

	public static class Class_bodyContext extends ParserRuleContext {
		public TerminalNode LCURLY() { return getToken(TswiftParser.LCURLY, 0); }
		public TerminalNode RCURLY() { return getToken(TswiftParser.RCURLY, 0); }
		public List<Class_memberContext> class_member() {
			return getRuleContexts(Class_memberContext.class);
		}
		public Class_memberContext class_member(int i) {
			return getRuleContext(Class_memberContext.class,i);
		}
		public Class_bodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_class_body; }
	}

	public final Class_bodyContext class_body() throws RecognitionException {
		Class_bodyContext _localctx = new Class_bodyContext(_ctx, getState());
		enterRule(_localctx, 176, RULE_class_body);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1130);
			match(LCURLY);
			setState(1134);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & ((1L << (T__14 - 15)) | (1L << (T__15 - 15)) | (1L << (T__16 - 15)) | (1L << (T__17 - 15)) | (1L << (T__18 - 15)) | (1L << (T__19 - 15)) | (1L << (T__20 - 15)) | (1L << (T__27 - 15)) | (1L << (T__28 - 15)) | (1L << (T__30 - 15)) | (1L << (T__31 - 15)) | (1L << (T__32 - 15)) | (1L << (T__33 - 15)) | (1L << (T__34 - 15)) | (1L << (T__36 - 15)) | (1L << (T__37 - 15)) | (1L << (T__38 - 15)) | (1L << (T__46 - 15)) | (1L << (T__47 - 15)) | (1L << (T__48 - 15)) | (1L << (T__49 - 15)) | (1L << (T__50 - 15)) | (1L << (T__51 - 15)) | (1L << (T__52 - 15)) | (1L << (T__53 - 15)) | (1L << (T__56 - 15)) | (1L << (T__57 - 15)) | (1L << (T__58 - 15)) | (1L << (T__59 - 15)) | (1L << (T__60 - 15)) | (1L << (T__61 - 15)) | (1L << (T__62 - 15)) | (1L << (T__63 - 15)))) != 0) || _la==AT) {
				{
				{
				setState(1131);
				class_member();
				}
				}
				setState(1136);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1137);
			match(RCURLY);
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

	public static class Class_memberContext extends ParserRuleContext {
		public DeclaracionContext declaracion() {
			return getRuleContext(DeclaracionContext.class,0);
		}
		public Class_memberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_class_member; }
	}

	public final Class_memberContext class_member() throws RecognitionException {
		Class_memberContext _localctx = new Class_memberContext(_ctx, getState());
		enterRule(_localctx, 178, RULE_class_member);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1139);
			declaracion();
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

	public static class Protocol_declaracionContext extends ParserRuleContext {
		public Protocol_nameContext protocol_name() {
			return getRuleContext(Protocol_nameContext.class,0);
		}
		public Protocol_bodyContext protocol_body() {
			return getRuleContext(Protocol_bodyContext.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Access_level_modifierContext access_level_modifier() {
			return getRuleContext(Access_level_modifierContext.class,0);
		}
		public Type_inheritance_clauseContext type_inheritance_clause() {
			return getRuleContext(Type_inheritance_clauseContext.class,0);
		}
		public Protocol_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_protocol_declaracion; }
	}

	public final Protocol_declaracionContext protocol_declaracion() throws RecognitionException {
		Protocol_declaracionContext _localctx = new Protocol_declaracionContext(_ctx, getState());
		enterRule(_localctx, 180, RULE_protocol_declaracion);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1142);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(1141);
				attributes();
				}
			}

			setState(1145);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__57) | (1L << T__58) | (1L << T__59) | (1L << T__60) | (1L << T__61))) != 0)) {
				{
				setState(1144);
				access_level_modifier();
				}
			}

			setState(1147);
			match(T__17);
			setState(1148);
			protocol_name();
			setState(1150);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COLON) {
				{
				setState(1149);
				type_inheritance_clause();
				}
			}

			setState(1152);
			protocol_body();
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

	public static class Protocol_nameContext extends ParserRuleContext {
		public Declaracion_identifierContext declaracion_identifier() {
			return getRuleContext(Declaracion_identifierContext.class,0);
		}
		public Protocol_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_protocol_name; }
	}

	public final Protocol_nameContext protocol_name() throws RecognitionException {
		Protocol_nameContext _localctx = new Protocol_nameContext(_ctx, getState());
		enterRule(_localctx, 182, RULE_protocol_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1154);
			declaracion_identifier();
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

	public static class Protocol_bodyContext extends ParserRuleContext {
		public TerminalNode LCURLY() { return getToken(TswiftParser.LCURLY, 0); }
		public TerminalNode RCURLY() { return getToken(TswiftParser.RCURLY, 0); }
		public List<Protocol_memberContext> protocol_member() {
			return getRuleContexts(Protocol_memberContext.class);
		}
		public Protocol_memberContext protocol_member(int i) {
			return getRuleContext(Protocol_memberContext.class,i);
		}
		public Protocol_bodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_protocol_body; }
	}

	public final Protocol_bodyContext protocol_body() throws RecognitionException {
		Protocol_bodyContext _localctx = new Protocol_bodyContext(_ctx, getState());
		enterRule(_localctx, 184, RULE_protocol_body);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1156);
			match(LCURLY);
			setState(1160);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 16)) & ~0x3f) == 0 && ((1L << (_la - 16)) & ((1L << (T__15 - 16)) | (1L << (T__18 - 16)) | (1L << (T__19 - 16)) | (1L << (T__28 - 16)) | (1L << (T__29 - 16)) | (1L << (T__30 - 16)) | (1L << (T__33 - 16)) | (1L << (T__34 - 16)) | (1L << (T__36 - 16)) | (1L << (T__37 - 16)) | (1L << (T__46 - 16)) | (1L << (T__47 - 16)) | (1L << (T__48 - 16)) | (1L << (T__49 - 16)) | (1L << (T__50 - 16)) | (1L << (T__51 - 16)) | (1L << (T__52 - 16)) | (1L << (T__53 - 16)) | (1L << (T__56 - 16)) | (1L << (T__57 - 16)) | (1L << (T__58 - 16)) | (1L << (T__59 - 16)) | (1L << (T__60 - 16)) | (1L << (T__61 - 16)) | (1L << (T__62 - 16)) | (1L << (T__63 - 16)))) != 0) || _la==AT) {
				{
				{
				setState(1157);
				protocol_member();
				}
				}
				setState(1162);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1163);
			match(RCURLY);
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

	public static class Protocol_memberContext extends ParserRuleContext {
		public Protocol_member_declaracionContext protocol_member_declaracion() {
			return getRuleContext(Protocol_member_declaracionContext.class,0);
		}
		public Protocol_memberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_protocol_member; }
	}

	public final Protocol_memberContext protocol_member() throws RecognitionException {
		Protocol_memberContext _localctx = new Protocol_memberContext(_ctx, getState());
		enterRule(_localctx, 186, RULE_protocol_member);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1165);
			protocol_member_declaracion();
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

	public static class Protocol_member_declaracionContext extends ParserRuleContext {
		public Protocol_property_declaracionContext protocol_property_declaracion() {
			return getRuleContext(Protocol_property_declaracionContext.class,0);
		}
		public Protocol_method_declaracionContext protocol_method_declaracion() {
			return getRuleContext(Protocol_method_declaracionContext.class,0);
		}
		public Protocol_initializer_declaracionContext protocol_initializer_declaracion() {
			return getRuleContext(Protocol_initializer_declaracionContext.class,0);
		}
		public Protocol_subscript_declaracionContext protocol_subscript_declaracion() {
			return getRuleContext(Protocol_subscript_declaracionContext.class,0);
		}
		public Protocol_associated_type_declaracionContext protocol_associated_type_declaracion() {
			return getRuleContext(Protocol_associated_type_declaracionContext.class,0);
		}
		public Protocol_member_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_protocol_member_declaracion; }
	}

	public final Protocol_member_declaracionContext protocol_member_declaracion() throws RecognitionException {
		Protocol_member_declaracionContext _localctx = new Protocol_member_declaracionContext(_ctx, getState());
		enterRule(_localctx, 188, RULE_protocol_member_declaracion);
		try {
			setState(1172);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,114,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1167);
				protocol_property_declaracion();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1168);
				protocol_method_declaracion();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1169);
				protocol_initializer_declaracion();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(1170);
				protocol_subscript_declaracion();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(1171);
				protocol_associated_type_declaracion();
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

	public static class Protocol_property_declaracionContext extends ParserRuleContext {
		public Variable_declaracion_headContext variable_declaracion_head() {
			return getRuleContext(Variable_declaracion_headContext.class,0);
		}
		public Variable_nameContext variable_name() {
			return getRuleContext(Variable_nameContext.class,0);
		}
		public Type_annotationContext type_annotation() {
			return getRuleContext(Type_annotationContext.class,0);
		}
		public Getter_setter_keyword_blockContext getter_setter_keyword_block() {
			return getRuleContext(Getter_setter_keyword_blockContext.class,0);
		}
		public Protocol_property_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_protocol_property_declaracion; }
	}

	public final Protocol_property_declaracionContext protocol_property_declaracion() throws RecognitionException {
		Protocol_property_declaracionContext _localctx = new Protocol_property_declaracionContext(_ctx, getState());
		enterRule(_localctx, 190, RULE_protocol_property_declaracion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1174);
			variable_declaracion_head();
			setState(1175);
			variable_name();
			setState(1176);
			type_annotation();
			setState(1177);
			getter_setter_keyword_block();
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

	public static class Protocol_method_declaracionContext extends ParserRuleContext {
		public Function_headContext function_head() {
			return getRuleContext(Function_headContext.class,0);
		}
		public Function_nameContext function_name() {
			return getRuleContext(Function_nameContext.class,0);
		}
		public Function_signatureContext function_signature() {
			return getRuleContext(Function_signatureContext.class,0);
		}
		public Generic_parameter_clauseContext generic_parameter_clause() {
			return getRuleContext(Generic_parameter_clauseContext.class,0);
		}
		public Generic_where_clauseContext generic_where_clause() {
			return getRuleContext(Generic_where_clauseContext.class,0);
		}
		public Protocol_method_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_protocol_method_declaracion; }
	}

	public final Protocol_method_declaracionContext protocol_method_declaracion() throws RecognitionException {
		Protocol_method_declaracionContext _localctx = new Protocol_method_declaracionContext(_ctx, getState());
		enterRule(_localctx, 192, RULE_protocol_method_declaracion);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1179);
			function_head();
			setState(1180);
			function_name();
			setState(1182);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(1181);
				generic_parameter_clause();
				}
			}

			setState(1184);
			function_signature();
			setState(1186);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__12) {
				{
				setState(1185);
				generic_where_clause();
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

	public static class Protocol_initializer_declaracionContext extends ParserRuleContext {
		public Initializer_headContext initializer_head() {
			return getRuleContext(Initializer_headContext.class,0);
		}
		public Parameter_clauseContext parameter_clause() {
			return getRuleContext(Parameter_clauseContext.class,0);
		}
		public Generic_parameter_clauseContext generic_parameter_clause() {
			return getRuleContext(Generic_parameter_clauseContext.class,0);
		}
		public Generic_where_clauseContext generic_where_clause() {
			return getRuleContext(Generic_where_clauseContext.class,0);
		}
		public Protocol_initializer_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_protocol_initializer_declaracion; }
	}

	public final Protocol_initializer_declaracionContext protocol_initializer_declaracion() throws RecognitionException {
		Protocol_initializer_declaracionContext _localctx = new Protocol_initializer_declaracionContext(_ctx, getState());
		enterRule(_localctx, 194, RULE_protocol_initializer_declaracion);
		int _la;
		try {
			setState(1208);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,122,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1188);
				initializer_head();
				setState(1190);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LT) {
					{
					setState(1189);
					generic_parameter_clause();
					}
				}

				setState(1192);
				parameter_clause();
				setState(1194);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__25) {
					{
					setState(1193);
					match(T__25);
					}
				}

				setState(1197);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__12) {
					{
					setState(1196);
					generic_where_clause();
					}
				}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1199);
				initializer_head();
				setState(1201);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LT) {
					{
					setState(1200);
					generic_parameter_clause();
					}
				}

				setState(1203);
				parameter_clause();
				setState(1204);
				match(T__26);
				setState(1206);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__12) {
					{
					setState(1205);
					generic_where_clause();
					}
				}

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

	public static class Protocol_subscript_declaracionContext extends ParserRuleContext {
		public Subscript_headContext subscript_head() {
			return getRuleContext(Subscript_headContext.class,0);
		}
		public Subscript_resultContext subscript_result() {
			return getRuleContext(Subscript_resultContext.class,0);
		}
		public Getter_setter_keyword_blockContext getter_setter_keyword_block() {
			return getRuleContext(Getter_setter_keyword_blockContext.class,0);
		}
		public Protocol_subscript_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_protocol_subscript_declaracion; }
	}

	public final Protocol_subscript_declaracionContext protocol_subscript_declaracion() throws RecognitionException {
		Protocol_subscript_declaracionContext _localctx = new Protocol_subscript_declaracionContext(_ctx, getState());
		enterRule(_localctx, 196, RULE_protocol_subscript_declaracion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1210);
			subscript_head();
			setState(1211);
			subscript_result();
			setState(1212);
			getter_setter_keyword_block();
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

	public static class Protocol_associated_type_declaracionContext extends ParserRuleContext {
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Access_level_modifierContext access_level_modifier() {
			return getRuleContext(Access_level_modifierContext.class,0);
		}
		public Type_inheritance_clauseContext type_inheritance_clause() {
			return getRuleContext(Type_inheritance_clauseContext.class,0);
		}
		public Protocol_associated_type_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_protocol_associated_type_declaracion; }
	}

	public final Protocol_associated_type_declaracionContext protocol_associated_type_declaracion() throws RecognitionException {
		Protocol_associated_type_declaracionContext _localctx = new Protocol_associated_type_declaracionContext(_ctx, getState());
		enterRule(_localctx, 198, RULE_protocol_associated_type_declaracion);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1215);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(1214);
				attributes();
				}
			}

			setState(1218);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__57) | (1L << T__58) | (1L << T__59) | (1L << T__60) | (1L << T__61))) != 0)) {
				{
				setState(1217);
				access_level_modifier();
				}
			}

			setState(1220);
			match(T__29);
			setState(1222);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COLON) {
				{
				setState(1221);
				type_inheritance_clause();
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

	public static class Initializer_declaracionContext extends ParserRuleContext {
		public Initializer_headContext initializer_head() {
			return getRuleContext(Initializer_headContext.class,0);
		}
		public Parameter_clauseContext parameter_clause() {
			return getRuleContext(Parameter_clauseContext.class,0);
		}
		public Initializer_bodyContext initializer_body() {
			return getRuleContext(Initializer_bodyContext.class,0);
		}
		public Generic_parameter_clauseContext generic_parameter_clause() {
			return getRuleContext(Generic_parameter_clauseContext.class,0);
		}
		public Generic_where_clauseContext generic_where_clause() {
			return getRuleContext(Generic_where_clauseContext.class,0);
		}
		public Initializer_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_initializer_declaracion; }
	}

	public final Initializer_declaracionContext initializer_declaracion() throws RecognitionException {
		Initializer_declaracionContext _localctx = new Initializer_declaracionContext(_ctx, getState());
		enterRule(_localctx, 200, RULE_initializer_declaracion);
		int _la;
		try {
			setState(1248);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,131,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1224);
				initializer_head();
				setState(1226);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LT) {
					{
					setState(1225);
					generic_parameter_clause();
					}
				}

				setState(1228);
				parameter_clause();
				setState(1230);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__25) {
					{
					setState(1229);
					match(T__25);
					}
				}

				setState(1233);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__12) {
					{
					setState(1232);
					generic_where_clause();
					}
				}

				setState(1235);
				initializer_body();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1237);
				initializer_head();
				setState(1239);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LT) {
					{
					setState(1238);
					generic_parameter_clause();
					}
				}

				setState(1241);
				parameter_clause();
				setState(1242);
				match(T__26);
				setState(1244);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__12) {
					{
					setState(1243);
					generic_where_clause();
					}
				}

				setState(1246);
				initializer_body();
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

	public static class Initializer_headContext extends ParserRuleContext {
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Declaracion_modifiersContext declaracion_modifiers() {
			return getRuleContext(Declaracion_modifiersContext.class,0);
		}
		public TerminalNode QUESTION() { return getToken(TswiftParser.QUESTION, 0); }
		public TerminalNode BANG() { return getToken(TswiftParser.BANG, 0); }
		public Initializer_headContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_initializer_head; }
	}

	public final Initializer_headContext initializer_head() throws RecognitionException {
		Initializer_headContext _localctx = new Initializer_headContext(_ctx, getState());
		enterRule(_localctx, 202, RULE_initializer_head);
		int _la;
		try {
			setState(1273);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,138,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1251);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AT) {
					{
					setState(1250);
					attributes();
					}
				}

				setState(1254);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (((((_la - 16)) & ~0x3f) == 0 && ((1L << (_la - 16)) & ((1L << (T__15 - 16)) | (1L << (T__28 - 16)) | (1L << (T__34 - 16)) | (1L << (T__36 - 16)) | (1L << (T__37 - 16)) | (1L << (T__46 - 16)) | (1L << (T__47 - 16)) | (1L << (T__48 - 16)) | (1L << (T__49 - 16)) | (1L << (T__50 - 16)) | (1L << (T__51 - 16)) | (1L << (T__52 - 16)) | (1L << (T__53 - 16)) | (1L << (T__56 - 16)) | (1L << (T__57 - 16)) | (1L << (T__58 - 16)) | (1L << (T__59 - 16)) | (1L << (T__60 - 16)) | (1L << (T__61 - 16)) | (1L << (T__62 - 16)) | (1L << (T__63 - 16)))) != 0)) {
					{
					setState(1253);
					declaracion_modifiers();
					}
				}

				setState(1256);
				match(T__30);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1258);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AT) {
					{
					setState(1257);
					attributes();
					}
				}

				setState(1261);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (((((_la - 16)) & ~0x3f) == 0 && ((1L << (_la - 16)) & ((1L << (T__15 - 16)) | (1L << (T__28 - 16)) | (1L << (T__34 - 16)) | (1L << (T__36 - 16)) | (1L << (T__37 - 16)) | (1L << (T__46 - 16)) | (1L << (T__47 - 16)) | (1L << (T__48 - 16)) | (1L << (T__49 - 16)) | (1L << (T__50 - 16)) | (1L << (T__51 - 16)) | (1L << (T__52 - 16)) | (1L << (T__53 - 16)) | (1L << (T__56 - 16)) | (1L << (T__57 - 16)) | (1L << (T__58 - 16)) | (1L << (T__59 - 16)) | (1L << (T__60 - 16)) | (1L << (T__61 - 16)) | (1L << (T__62 - 16)) | (1L << (T__63 - 16)))) != 0)) {
					{
					setState(1260);
					declaracion_modifiers();
					}
				}

				setState(1263);
				match(T__30);
				setState(1264);
				match(QUESTION);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1266);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AT) {
					{
					setState(1265);
					attributes();
					}
				}

				setState(1269);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (((((_la - 16)) & ~0x3f) == 0 && ((1L << (_la - 16)) & ((1L << (T__15 - 16)) | (1L << (T__28 - 16)) | (1L << (T__34 - 16)) | (1L << (T__36 - 16)) | (1L << (T__37 - 16)) | (1L << (T__46 - 16)) | (1L << (T__47 - 16)) | (1L << (T__48 - 16)) | (1L << (T__49 - 16)) | (1L << (T__50 - 16)) | (1L << (T__51 - 16)) | (1L << (T__52 - 16)) | (1L << (T__53 - 16)) | (1L << (T__56 - 16)) | (1L << (T__57 - 16)) | (1L << (T__58 - 16)) | (1L << (T__59 - 16)) | (1L << (T__60 - 16)) | (1L << (T__61 - 16)) | (1L << (T__62 - 16)) | (1L << (T__63 - 16)))) != 0)) {
					{
					setState(1268);
					declaracion_modifiers();
					}
				}

				setState(1271);
				match(T__30);
				setState(1272);
				match(BANG);
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

	public static class Initializer_bodyContext extends ParserRuleContext {
		public Code_blockContext code_block() {
			return getRuleContext(Code_blockContext.class,0);
		}
		public Initializer_bodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_initializer_body; }
	}

	public final Initializer_bodyContext initializer_body() throws RecognitionException {
		Initializer_bodyContext _localctx = new Initializer_bodyContext(_ctx, getState());
		enterRule(_localctx, 204, RULE_initializer_body);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1275);
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

	public static class Deinitializer_declaracionContext extends ParserRuleContext {
		public Code_blockContext code_block() {
			return getRuleContext(Code_blockContext.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Deinitializer_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_deinitializer_declaracion; }
	}

	public final Deinitializer_declaracionContext deinitializer_declaracion() throws RecognitionException {
		Deinitializer_declaracionContext _localctx = new Deinitializer_declaracionContext(_ctx, getState());
		enterRule(_localctx, 206, RULE_deinitializer_declaracion);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1278);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(1277);
				attributes();
				}
			}

			setState(1280);
			match(T__31);
			setState(1281);
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

	public static class Extension_declaracionContext extends ParserRuleContext {
		public Type_identifierContext type_identifier() {
			return getRuleContext(Type_identifierContext.class,0);
		}
		public Extension_bodyContext extension_body() {
			return getRuleContext(Extension_bodyContext.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Access_level_modifierContext access_level_modifier() {
			return getRuleContext(Access_level_modifierContext.class,0);
		}
		public Type_inheritance_clauseContext type_inheritance_clause() {
			return getRuleContext(Type_inheritance_clauseContext.class,0);
		}
		public Generic_where_clauseContext generic_where_clause() {
			return getRuleContext(Generic_where_clauseContext.class,0);
		}
		public Extension_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_extension_declaracion; }
	}

	public final Extension_declaracionContext extension_declaracion() throws RecognitionException {
		Extension_declaracionContext _localctx = new Extension_declaracionContext(_ctx, getState());
		enterRule(_localctx, 208, RULE_extension_declaracion);
		int _la;
		try {
			setState(1307);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,145,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1284);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AT) {
					{
					setState(1283);
					attributes();
					}
				}

				setState(1287);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__57) | (1L << T__58) | (1L << T__59) | (1L << T__60) | (1L << T__61))) != 0)) {
					{
					setState(1286);
					access_level_modifier();
					}
				}

				setState(1289);
				match(T__32);
				setState(1290);
				type_identifier();
				setState(1292);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COLON) {
					{
					setState(1291);
					type_inheritance_clause();
					}
				}

				setState(1294);
				extension_body();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1297);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AT) {
					{
					setState(1296);
					attributes();
					}
				}

				setState(1300);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__57) | (1L << T__58) | (1L << T__59) | (1L << T__60) | (1L << T__61))) != 0)) {
					{
					setState(1299);
					access_level_modifier();
					}
				}

				setState(1302);
				match(T__32);
				setState(1303);
				type_identifier();
				setState(1304);
				generic_where_clause();
				setState(1305);
				extension_body();
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

	public static class Extension_bodyContext extends ParserRuleContext {
		public TerminalNode LCURLY() { return getToken(TswiftParser.LCURLY, 0); }
		public TerminalNode RCURLY() { return getToken(TswiftParser.RCURLY, 0); }
		public List<Extension_memberContext> extension_member() {
			return getRuleContexts(Extension_memberContext.class);
		}
		public Extension_memberContext extension_member(int i) {
			return getRuleContext(Extension_memberContext.class,i);
		}
		public Extension_bodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_extension_body; }
	}

	public final Extension_bodyContext extension_body() throws RecognitionException {
		Extension_bodyContext _localctx = new Extension_bodyContext(_ctx, getState());
		enterRule(_localctx, 210, RULE_extension_body);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1309);
			match(LCURLY);
			setState(1313);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & ((1L << (T__14 - 15)) | (1L << (T__15 - 15)) | (1L << (T__16 - 15)) | (1L << (T__17 - 15)) | (1L << (T__18 - 15)) | (1L << (T__19 - 15)) | (1L << (T__20 - 15)) | (1L << (T__27 - 15)) | (1L << (T__28 - 15)) | (1L << (T__30 - 15)) | (1L << (T__31 - 15)) | (1L << (T__32 - 15)) | (1L << (T__33 - 15)) | (1L << (T__34 - 15)) | (1L << (T__36 - 15)) | (1L << (T__37 - 15)) | (1L << (T__38 - 15)) | (1L << (T__46 - 15)) | (1L << (T__47 - 15)) | (1L << (T__48 - 15)) | (1L << (T__49 - 15)) | (1L << (T__50 - 15)) | (1L << (T__51 - 15)) | (1L << (T__52 - 15)) | (1L << (T__53 - 15)) | (1L << (T__56 - 15)) | (1L << (T__57 - 15)) | (1L << (T__58 - 15)) | (1L << (T__59 - 15)) | (1L << (T__60 - 15)) | (1L << (T__61 - 15)) | (1L << (T__62 - 15)) | (1L << (T__63 - 15)))) != 0) || _la==AT) {
				{
				{
				setState(1310);
				extension_member();
				}
				}
				setState(1315);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1316);
			match(RCURLY);
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

	public static class Extension_memberContext extends ParserRuleContext {
		public DeclaracionContext declaracion() {
			return getRuleContext(DeclaracionContext.class,0);
		}
		public Extension_memberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_extension_member; }
	}

	public final Extension_memberContext extension_member() throws RecognitionException {
		Extension_memberContext _localctx = new Extension_memberContext(_ctx, getState());
		enterRule(_localctx, 212, RULE_extension_member);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1318);
			declaracion();
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

	public static class Subscript_declaracionContext extends ParserRuleContext {
		public Subscript_headContext subscript_head() {
			return getRuleContext(Subscript_headContext.class,0);
		}
		public Subscript_resultContext subscript_result() {
			return getRuleContext(Subscript_resultContext.class,0);
		}
		public Code_blockContext code_block() {
			return getRuleContext(Code_blockContext.class,0);
		}
		public Getter_setter_blockContext getter_setter_block() {
			return getRuleContext(Getter_setter_blockContext.class,0);
		}
		public Getter_setter_keyword_blockContext getter_setter_keyword_block() {
			return getRuleContext(Getter_setter_keyword_blockContext.class,0);
		}
		public Subscript_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_subscript_declaracion; }
	}

	public final Subscript_declaracionContext subscript_declaracion() throws RecognitionException {
		Subscript_declaracionContext _localctx = new Subscript_declaracionContext(_ctx, getState());
		enterRule(_localctx, 214, RULE_subscript_declaracion);
		try {
			setState(1332);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,147,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1320);
				subscript_head();
				setState(1321);
				subscript_result();
				setState(1322);
				code_block();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1324);
				subscript_head();
				setState(1325);
				subscript_result();
				setState(1326);
				getter_setter_block();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1328);
				subscript_head();
				setState(1329);
				subscript_result();
				setState(1330);
				getter_setter_keyword_block();
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

	public static class Subscript_headContext extends ParserRuleContext {
		public Parameter_clauseContext parameter_clause() {
			return getRuleContext(Parameter_clauseContext.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Declaracion_modifiersContext declaracion_modifiers() {
			return getRuleContext(Declaracion_modifiersContext.class,0);
		}
		public Subscript_headContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_subscript_head; }
	}

	public final Subscript_headContext subscript_head() throws RecognitionException {
		Subscript_headContext _localctx = new Subscript_headContext(_ctx, getState());
		enterRule(_localctx, 216, RULE_subscript_head);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1335);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(1334);
				attributes();
				}
			}

			setState(1338);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 16)) & ~0x3f) == 0 && ((1L << (_la - 16)) & ((1L << (T__15 - 16)) | (1L << (T__28 - 16)) | (1L << (T__34 - 16)) | (1L << (T__36 - 16)) | (1L << (T__37 - 16)) | (1L << (T__46 - 16)) | (1L << (T__47 - 16)) | (1L << (T__48 - 16)) | (1L << (T__49 - 16)) | (1L << (T__50 - 16)) | (1L << (T__51 - 16)) | (1L << (T__52 - 16)) | (1L << (T__53 - 16)) | (1L << (T__56 - 16)) | (1L << (T__57 - 16)) | (1L << (T__58 - 16)) | (1L << (T__59 - 16)) | (1L << (T__60 - 16)) | (1L << (T__61 - 16)) | (1L << (T__62 - 16)) | (1L << (T__63 - 16)))) != 0)) {
				{
				setState(1337);
				declaracion_modifiers();
				}
			}

			setState(1340);
			match(T__33);
			setState(1341);
			parameter_clause();
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

	public static class Subscript_resultContext extends ParserRuleContext {
		public Arrow_operatorContext arrow_operator() {
			return getRuleContext(Arrow_operatorContext.class,0);
		}
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Subscript_resultContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_subscript_result; }
	}

	public final Subscript_resultContext subscript_result() throws RecognitionException {
		Subscript_resultContext _localctx = new Subscript_resultContext(_ctx, getState());
		enterRule(_localctx, 218, RULE_subscript_result);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1343);
			arrow_operator();
			setState(1345);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,150,_ctx) ) {
			case 1:
				{
				setState(1344);
				attributes();
				}
				break;
			}
			setState(1347);
			type_(0);
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

	public static class Operator_declaracionContext extends ParserRuleContext {
		public Prefix_operator_declaracionContext prefix_operator_declaracion() {
			return getRuleContext(Prefix_operator_declaracionContext.class,0);
		}
		public Postfix_operator_declaracionContext postfix_operator_declaracion() {
			return getRuleContext(Postfix_operator_declaracionContext.class,0);
		}
		public Infix_operator_declaracionContext infix_operator_declaracion() {
			return getRuleContext(Infix_operator_declaracionContext.class,0);
		}
		public Operator_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operator_declaracion; }
	}

	public final Operator_declaracionContext operator_declaracion() throws RecognitionException {
		Operator_declaracionContext _localctx = new Operator_declaracionContext(_ctx, getState());
		enterRule(_localctx, 220, RULE_operator_declaracion);
		try {
			setState(1352);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__34:
				enterOuterAlt(_localctx, 1);
				{
				setState(1349);
				prefix_operator_declaracion();
				}
				break;
			case T__36:
				enterOuterAlt(_localctx, 2);
				{
				setState(1350);
				postfix_operator_declaracion();
				}
				break;
			case T__37:
				enterOuterAlt(_localctx, 3);
				{
				setState(1351);
				infix_operator_declaracion();
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

	public static class Prefix_operator_declaracionContext extends ParserRuleContext {
		public Operator_Context operator_() {
			return getRuleContext(Operator_Context.class,0);
		}
		public Prefix_operator_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prefix_operator_declaracion; }
	}

	public final Prefix_operator_declaracionContext prefix_operator_declaracion() throws RecognitionException {
		Prefix_operator_declaracionContext _localctx = new Prefix_operator_declaracionContext(_ctx, getState());
		enterRule(_localctx, 222, RULE_prefix_operator_declaracion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1354);
			match(T__34);
			setState(1355);
			match(T__35);
			setState(1356);
			operator_();
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

	public static class Postfix_operator_declaracionContext extends ParserRuleContext {
		public Operator_Context operator_() {
			return getRuleContext(Operator_Context.class,0);
		}
		public Postfix_operator_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_postfix_operator_declaracion; }
	}

	public final Postfix_operator_declaracionContext postfix_operator_declaracion() throws RecognitionException {
		Postfix_operator_declaracionContext _localctx = new Postfix_operator_declaracionContext(_ctx, getState());
		enterRule(_localctx, 224, RULE_postfix_operator_declaracion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1358);
			match(T__36);
			setState(1359);
			match(T__35);
			setState(1360);
			operator_();
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

	public static class Infix_operator_declaracionContext extends ParserRuleContext {
		public Operator_Context operator_() {
			return getRuleContext(Operator_Context.class,0);
		}
		public Infix_operator_groupContext infix_operator_group() {
			return getRuleContext(Infix_operator_groupContext.class,0);
		}
		public Infix_operator_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_infix_operator_declaracion; }
	}

	public final Infix_operator_declaracionContext infix_operator_declaracion() throws RecognitionException {
		Infix_operator_declaracionContext _localctx = new Infix_operator_declaracionContext(_ctx, getState());
		enterRule(_localctx, 226, RULE_infix_operator_declaracion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1362);
			match(T__37);
			setState(1363);
			match(T__35);
			setState(1364);
			operator_();
			setState(1366);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,152,_ctx) ) {
			case 1:
				{
				setState(1365);
				infix_operator_group();
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

	public static class Infix_operator_groupContext extends ParserRuleContext {
		public TerminalNode COLON() { return getToken(TswiftParser.COLON, 0); }
		public Precedence_group_nameContext precedence_group_name() {
			return getRuleContext(Precedence_group_nameContext.class,0);
		}
		public Infix_operator_groupContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_infix_operator_group; }
	}

	public final Infix_operator_groupContext infix_operator_group() throws RecognitionException {
		Infix_operator_groupContext _localctx = new Infix_operator_groupContext(_ctx, getState());
		enterRule(_localctx, 228, RULE_infix_operator_group);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1368);
			match(COLON);
			setState(1369);
			precedence_group_name();
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

	public static class Precedence_group_declaracionContext extends ParserRuleContext {
		public Precedence_group_nameContext precedence_group_name() {
			return getRuleContext(Precedence_group_nameContext.class,0);
		}
		public TerminalNode LCURLY() { return getToken(TswiftParser.LCURLY, 0); }
		public TerminalNode RCURLY() { return getToken(TswiftParser.RCURLY, 0); }
		public List<Precedence_group_attributeContext> precedence_group_attribute() {
			return getRuleContexts(Precedence_group_attributeContext.class);
		}
		public Precedence_group_attributeContext precedence_group_attribute(int i) {
			return getRuleContext(Precedence_group_attributeContext.class,i);
		}
		public Precedence_group_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_precedence_group_declaracion; }
	}

	public final Precedence_group_declaracionContext precedence_group_declaracion() throws RecognitionException {
		Precedence_group_declaracionContext _localctx = new Precedence_group_declaracionContext(_ctx, getState());
		enterRule(_localctx, 230, RULE_precedence_group_declaracion);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1371);
			match(T__38);
			setState(1372);
			precedence_group_name();
			setState(1373);
			match(LCURLY);
			setState(1377);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42))) != 0)) {
				{
				{
				setState(1374);
				precedence_group_attribute();
				}
				}
				setState(1379);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1380);
			match(RCURLY);
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

	public static class Precedence_group_attributeContext extends ParserRuleContext {
		public Precedence_group_relationContext precedence_group_relation() {
			return getRuleContext(Precedence_group_relationContext.class,0);
		}
		public Precedence_group_assignmentContext precedence_group_assignment() {
			return getRuleContext(Precedence_group_assignmentContext.class,0);
		}
		public Precedence_group_associativityContext precedence_group_associativity() {
			return getRuleContext(Precedence_group_associativityContext.class,0);
		}
		public Precedence_group_attributeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_precedence_group_attribute; }
	}

	public final Precedence_group_attributeContext precedence_group_attribute() throws RecognitionException {
		Precedence_group_attributeContext _localctx = new Precedence_group_attributeContext(_ctx, getState());
		enterRule(_localctx, 232, RULE_precedence_group_attribute);
		try {
			setState(1385);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__39:
			case T__40:
				enterOuterAlt(_localctx, 1);
				{
				setState(1382);
				precedence_group_relation();
				}
				break;
			case T__41:
				enterOuterAlt(_localctx, 2);
				{
				setState(1383);
				precedence_group_assignment();
				}
				break;
			case T__42:
				enterOuterAlt(_localctx, 3);
				{
				setState(1384);
				precedence_group_associativity();
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

	public static class Precedence_group_relationContext extends ParserRuleContext {
		public TerminalNode COLON() { return getToken(TswiftParser.COLON, 0); }
		public Precedence_group_namesContext precedence_group_names() {
			return getRuleContext(Precedence_group_namesContext.class,0);
		}
		public Precedence_group_relationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_precedence_group_relation; }
	}

	public final Precedence_group_relationContext precedence_group_relation() throws RecognitionException {
		Precedence_group_relationContext _localctx = new Precedence_group_relationContext(_ctx, getState());
		enterRule(_localctx, 234, RULE_precedence_group_relation);
		try {
			setState(1393);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__39:
				enterOuterAlt(_localctx, 1);
				{
				setState(1387);
				match(T__39);
				setState(1388);
				match(COLON);
				setState(1389);
				precedence_group_names();
				}
				break;
			case T__40:
				enterOuterAlt(_localctx, 2);
				{
				setState(1390);
				match(T__40);
				setState(1391);
				match(COLON);
				setState(1392);
				precedence_group_names();
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

	public static class Precedence_group_assignmentContext extends ParserRuleContext {
		public TerminalNode COLON() { return getToken(TswiftParser.COLON, 0); }
		public Boolean_literalContext boolean_literal() {
			return getRuleContext(Boolean_literalContext.class,0);
		}
		public Precedence_group_assignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_precedence_group_assignment; }
	}

	public final Precedence_group_assignmentContext precedence_group_assignment() throws RecognitionException {
		Precedence_group_assignmentContext _localctx = new Precedence_group_assignmentContext(_ctx, getState());
		enterRule(_localctx, 236, RULE_precedence_group_assignment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1395);
			match(T__41);
			setState(1396);
			match(COLON);
			setState(1397);
			boolean_literal();
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

	public static class Precedence_group_associativityContext extends ParserRuleContext {
		public TerminalNode COLON() { return getToken(TswiftParser.COLON, 0); }
		public Associativity_Context associativity_() {
			return getRuleContext(Associativity_Context.class,0);
		}
		public Precedence_group_associativityContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_precedence_group_associativity; }
	}

	public final Precedence_group_associativityContext precedence_group_associativity() throws RecognitionException {
		Precedence_group_associativityContext _localctx = new Precedence_group_associativityContext(_ctx, getState());
		enterRule(_localctx, 238, RULE_precedence_group_associativity);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1399);
			match(T__42);
			setState(1400);
			match(COLON);
			setState(1401);
			associativity_();
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

	public static class Associativity_Context extends ParserRuleContext {
		public Associativity_Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_associativity_; }
	}

	public final Associativity_Context associativity_() throws RecognitionException {
		Associativity_Context _localctx = new Associativity_Context(_ctx, getState());
		enterRule(_localctx, 240, RULE_associativity_);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1403);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__43) | (1L << T__44) | (1L << T__45))) != 0)) ) {
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

	public static class Precedence_group_namesContext extends ParserRuleContext {
		public List<Precedence_group_nameContext> precedence_group_name() {
			return getRuleContexts(Precedence_group_nameContext.class);
		}
		public Precedence_group_nameContext precedence_group_name(int i) {
			return getRuleContext(Precedence_group_nameContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(TswiftParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(TswiftParser.COMMA, i);
		}
		public Precedence_group_namesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_precedence_group_names; }
	}

	public final Precedence_group_namesContext precedence_group_names() throws RecognitionException {
		Precedence_group_namesContext _localctx = new Precedence_group_namesContext(_ctx, getState());
		enterRule(_localctx, 242, RULE_precedence_group_names);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1405);
			precedence_group_name();
			setState(1410);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(1406);
				match(COMMA);
				setState(1407);
				precedence_group_name();
				}
				}
				setState(1412);
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

	public static class Precedence_group_nameContext extends ParserRuleContext {
		public Declaracion_identifierContext declaracion_identifier() {
			return getRuleContext(Declaracion_identifierContext.class,0);
		}
		public Precedence_group_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_precedence_group_name; }
	}

	public final Precedence_group_nameContext precedence_group_name() throws RecognitionException {
		Precedence_group_nameContext _localctx = new Precedence_group_nameContext(_ctx, getState());
		enterRule(_localctx, 244, RULE_precedence_group_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1413);
			declaracion_identifier();
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

	public static class Declaracion_modifierContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(TswiftParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(TswiftParser.RPAREN, 0); }
		public Access_level_modifierContext access_level_modifier() {
			return getRuleContext(Access_level_modifierContext.class,0);
		}
		public Mutation_modifierContext mutation_modifier() {
			return getRuleContext(Mutation_modifierContext.class,0);
		}
		public Declaracion_modifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion_modifier; }
	}

	public final Declaracion_modifierContext declaracion_modifier() throws RecognitionException {
		Declaracion_modifierContext _localctx = new Declaracion_modifierContext(_ctx, getState());
		enterRule(_localctx, 246, RULE_declaracion_modifier);
		try {
			setState(1439);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,157,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1415);
				match(T__15);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1416);
				match(T__46);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1417);
				match(T__47);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(1418);
				match(T__28);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(1419);
				match(T__37);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(1420);
				match(T__48);
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(1421);
				match(T__49);
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(1422);
				match(T__50);
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(1423);
				match(T__36);
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(1424);
				match(T__34);
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(1425);
				match(T__51);
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(1426);
				match(T__52);
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(1427);
				match(T__53);
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(1428);
				match(T__53);
				setState(1429);
				match(LPAREN);
				setState(1430);
				match(T__54);
				setState(1431);
				match(RPAREN);
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(1432);
				match(T__53);
				setState(1433);
				match(LPAREN);
				setState(1434);
				match(T__55);
				setState(1435);
				match(RPAREN);
				}
				break;
			case 16:
				enterOuterAlt(_localctx, 16);
				{
				setState(1436);
				match(T__56);
				}
				break;
			case 17:
				enterOuterAlt(_localctx, 17);
				{
				setState(1437);
				access_level_modifier();
				}
				break;
			case 18:
				enterOuterAlt(_localctx, 18);
				{
				setState(1438);
				mutation_modifier();
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

	public static class Declaracion_modifiersContext extends ParserRuleContext {
		public List<Declaracion_modifierContext> declaracion_modifier() {
			return getRuleContexts(Declaracion_modifierContext.class);
		}
		public Declaracion_modifierContext declaracion_modifier(int i) {
			return getRuleContext(Declaracion_modifierContext.class,i);
		}
		public Declaracion_modifiersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion_modifiers; }
	}

	public final Declaracion_modifiersContext declaracion_modifiers() throws RecognitionException {
		Declaracion_modifiersContext _localctx = new Declaracion_modifiersContext(_ctx, getState());
		enterRule(_localctx, 248, RULE_declaracion_modifiers);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1442); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(1441);
				declaracion_modifier();
				}
				}
				setState(1444); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( ((((_la - 16)) & ~0x3f) == 0 && ((1L << (_la - 16)) & ((1L << (T__15 - 16)) | (1L << (T__28 - 16)) | (1L << (T__34 - 16)) | (1L << (T__36 - 16)) | (1L << (T__37 - 16)) | (1L << (T__46 - 16)) | (1L << (T__47 - 16)) | (1L << (T__48 - 16)) | (1L << (T__49 - 16)) | (1L << (T__50 - 16)) | (1L << (T__51 - 16)) | (1L << (T__52 - 16)) | (1L << (T__53 - 16)) | (1L << (T__56 - 16)) | (1L << (T__57 - 16)) | (1L << (T__58 - 16)) | (1L << (T__59 - 16)) | (1L << (T__60 - 16)) | (1L << (T__61 - 16)) | (1L << (T__62 - 16)) | (1L << (T__63 - 16)))) != 0) );
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

	public static class Access_level_modifierContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(TswiftParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(TswiftParser.RPAREN, 0); }
		public Access_level_modifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_access_level_modifier; }
	}

	public final Access_level_modifierContext access_level_modifier() throws RecognitionException {
		Access_level_modifierContext _localctx = new Access_level_modifierContext(_ctx, getState());
		enterRule(_localctx, 250, RULE_access_level_modifier);
		try {
			setState(1471);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,159,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1446);
				match(T__57);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1447);
				match(T__57);
				setState(1448);
				match(LPAREN);
				setState(1449);
				match(T__22);
				setState(1450);
				match(RPAREN);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1451);
				match(T__58);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(1452);
				match(T__58);
				setState(1453);
				match(LPAREN);
				setState(1454);
				match(T__22);
				setState(1455);
				match(RPAREN);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(1456);
				match(T__59);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(1457);
				match(T__59);
				setState(1458);
				match(LPAREN);
				setState(1459);
				match(T__22);
				setState(1460);
				match(RPAREN);
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(1461);
				match(T__60);
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(1462);
				match(T__60);
				setState(1463);
				match(LPAREN);
				setState(1464);
				match(T__22);
				setState(1465);
				match(RPAREN);
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(1466);
				match(T__61);
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(1467);
				match(T__61);
				setState(1468);
				match(LPAREN);
				setState(1469);
				match(T__22);
				setState(1470);
				match(RPAREN);
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

	public static class Mutation_modifierContext extends ParserRuleContext {
		public Mutation_modifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mutation_modifier; }
	}

	public final Mutation_modifierContext mutation_modifier() throws RecognitionException {
		Mutation_modifierContext _localctx = new Mutation_modifierContext(_ctx, getState());
		enterRule(_localctx, 252, RULE_mutation_modifier);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1473);
			_la = _input.LA(1);
			if ( !(_la==T__62 || _la==T__63) ) {
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

	public static class PatternContext extends ParserRuleContext {
		public Wildcard_patternContext wildcard_pattern() {
			return getRuleContext(Wildcard_patternContext.class,0);
		}
		public Type_annotationContext type_annotation() {
			return getRuleContext(Type_annotationContext.class,0);
		}
		public Identifier_patternContext identifier_pattern() {
			return getRuleContext(Identifier_patternContext.class,0);
		}
		public Value_binding_patternContext value_binding_pattern() {
			return getRuleContext(Value_binding_patternContext.class,0);
		}
		public Tuple_patternContext tuple_pattern() {
			return getRuleContext(Tuple_patternContext.class,0);
		}
		public Enum_case_patternContext enum_case_pattern() {
			return getRuleContext(Enum_case_patternContext.class,0);
		}
		public Optional_patternContext optional_pattern() {
			return getRuleContext(Optional_patternContext.class,0);
		}
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public Expresion_patternContext expresion_pattern() {
			return getRuleContext(Expresion_patternContext.class,0);
		}
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public PatternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pattern; }
	}

	public final PatternContext pattern() throws RecognitionException {
		return pattern(0);
	}

	private PatternContext pattern(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		PatternContext _localctx = new PatternContext(_ctx, _parentState);
		PatternContext _prevctx = _localctx;
		int _startState = 254;
		enterRecursionRule(_localctx, 254, RULE_pattern, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1494);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,163,_ctx) ) {
			case 1:
				{
				setState(1476);
				wildcard_pattern();
				setState(1478);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,160,_ctx) ) {
				case 1:
					{
					setState(1477);
					type_annotation();
					}
					break;
				}
				}
				break;
			case 2:
				{
				setState(1480);
				identifier_pattern();
				setState(1482);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,161,_ctx) ) {
				case 1:
					{
					setState(1481);
					type_annotation();
					}
					break;
				}
				}
				break;
			case 3:
				{
				setState(1484);
				value_binding_pattern();
				}
				break;
			case 4:
				{
				setState(1485);
				tuple_pattern();
				setState(1487);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,162,_ctx) ) {
				case 1:
					{
					setState(1486);
					type_annotation();
					}
					break;
				}
				}
				break;
			case 5:
				{
				setState(1489);
				enum_case_pattern();
				}
				break;
			case 6:
				{
				setState(1490);
				optional_pattern();
				}
				break;
			case 7:
				{
				setState(1491);
				match(T__64);
				setState(1492);
				type_(0);
				}
				break;
			case 8:
				{
				setState(1493);
				expresion_pattern();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(1501);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,164,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new PatternContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_pattern);
					setState(1496);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(1497);
					match(T__65);
					setState(1498);
					type_(0);
					}
					} 
				}
				setState(1503);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,164,_ctx);
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

	public static class Wildcard_patternContext extends ParserRuleContext {
		public TerminalNode UNDERSCORE() { return getToken(TswiftParser.UNDERSCORE, 0); }
		public Wildcard_patternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_wildcard_pattern; }
	}

	public final Wildcard_patternContext wildcard_pattern() throws RecognitionException {
		Wildcard_patternContext _localctx = new Wildcard_patternContext(_ctx, getState());
		enterRule(_localctx, 256, RULE_wildcard_pattern);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1504);
			match(UNDERSCORE);
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

	public static class Identifier_patternContext extends ParserRuleContext {
		public Declaracion_identifierContext declaracion_identifier() {
			return getRuleContext(Declaracion_identifierContext.class,0);
		}
		public Identifier_patternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifier_pattern; }
	}

	public final Identifier_patternContext identifier_pattern() throws RecognitionException {
		Identifier_patternContext _localctx = new Identifier_patternContext(_ctx, getState());
		enterRule(_localctx, 258, RULE_identifier_pattern);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1506);
			declaracion_identifier();
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

	public static class Value_binding_patternContext extends ParserRuleContext {
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public Value_binding_patternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_value_binding_pattern; }
	}

	public final Value_binding_patternContext value_binding_pattern() throws RecognitionException {
		Value_binding_patternContext _localctx = new Value_binding_patternContext(_ctx, getState());
		enterRule(_localctx, 260, RULE_value_binding_pattern);
		try {
			setState(1512);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__18:
				enterOuterAlt(_localctx, 1);
				{
				setState(1508);
				match(T__18);
				setState(1509);
				pattern(0);
				}
				break;
			case T__20:
				enterOuterAlt(_localctx, 2);
				{
				setState(1510);
				match(T__20);
				setState(1511);
				pattern(0);
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

	public static class Tuple_patternContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(TswiftParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(TswiftParser.RPAREN, 0); }
		public Tuple_pattern_element_listContext tuple_pattern_element_list() {
			return getRuleContext(Tuple_pattern_element_listContext.class,0);
		}
		public Tuple_patternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuple_pattern; }
	}

	public final Tuple_patternContext tuple_pattern() throws RecognitionException {
		Tuple_patternContext _localctx = new Tuple_patternContext(_ctx, getState());
		enterRule(_localctx, 262, RULE_tuple_pattern);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1514);
			match(LPAREN);
			setState(1516);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,166,_ctx) ) {
			case 1:
				{
				setState(1515);
				tuple_pattern_element_list();
				}
				break;
			}
			setState(1518);
			match(RPAREN);
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

	public static class Tuple_pattern_element_listContext extends ParserRuleContext {
		public List<Tuple_pattern_elementContext> tuple_pattern_element() {
			return getRuleContexts(Tuple_pattern_elementContext.class);
		}
		public Tuple_pattern_elementContext tuple_pattern_element(int i) {
			return getRuleContext(Tuple_pattern_elementContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(TswiftParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(TswiftParser.COMMA, i);
		}
		public Tuple_pattern_element_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuple_pattern_element_list; }
	}

	public final Tuple_pattern_element_listContext tuple_pattern_element_list() throws RecognitionException {
		Tuple_pattern_element_listContext _localctx = new Tuple_pattern_element_listContext(_ctx, getState());
		enterRule(_localctx, 264, RULE_tuple_pattern_element_list);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1520);
			tuple_pattern_element();
			setState(1525);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(1521);
				match(COMMA);
				setState(1522);
				tuple_pattern_element();
				}
				}
				setState(1527);
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

	public static class Tuple_pattern_elementContext extends ParserRuleContext {
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public Tuple_pattern_elementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuple_pattern_element; }
	}

	public final Tuple_pattern_elementContext tuple_pattern_element() throws RecognitionException {
		Tuple_pattern_elementContext _localctx = new Tuple_pattern_elementContext(_ctx, getState());
		enterRule(_localctx, 266, RULE_tuple_pattern_element);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1528);
			pattern(0);
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

	public static class Enum_case_patternContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(TswiftParser.DOT, 0); }
		public Enum_case_nameContext enum_case_name() {
			return getRuleContext(Enum_case_nameContext.class,0);
		}
		public Type_identifierContext type_identifier() {
			return getRuleContext(Type_identifierContext.class,0);
		}
		public Tuple_patternContext tuple_pattern() {
			return getRuleContext(Tuple_patternContext.class,0);
		}
		public Enum_case_patternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enum_case_pattern; }
	}

	public final Enum_case_patternContext enum_case_pattern() throws RecognitionException {
		Enum_case_patternContext _localctx = new Enum_case_patternContext(_ctx, getState());
		enterRule(_localctx, 268, RULE_enum_case_pattern);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1531);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__27) | (1L << T__28) | (1L << T__34) | (1L << T__36) | (1L << T__37) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << T__43) | (1L << T__44) | (1L << T__45) | (1L << T__46) | (1L << T__47) | (1L << T__48) | (1L << T__49) | (1L << T__50) | (1L << T__51) | (1L << T__53) | (1L << T__54) | (1L << T__55) | (1L << T__56) | (1L << T__61) | (1L << T__62))) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & ((1L << (T__63 - 64)) | (1L << (T__75 - 64)) | (1L << (T__76 - 64)) | (1L << (T__77 - 64)) | (1L << (T__78 - 64)) | (1L << (T__80 - 64)) | (1L << (T__91 - 64)) | (1L << (T__92 - 64)) | (1L << (T__93 - 64)) | (1L << (T__94 - 64)) | (1L << (T__97 - 64)) | (1L << (T__98 - 64)) | (1L << (T__99 - 64)) | (1L << (T__100 - 64)) | (1L << (T__101 - 64)) | (1L << (T__102 - 64)) | (1L << (T__103 - 64)) | (1L << (T__104 - 64)) | (1L << (T__105 - 64)) | (1L << (T__106 - 64)) | (1L << (T__107 - 64)) | (1L << (T__108 - 64)) | (1L << (T__109 - 64)) | (1L << (T__110 - 64)) | (1L << (T__111 - 64)) | (1L << (T__112 - 64)) | (1L << (Identifier - 64)))) != 0)) {
				{
				setState(1530);
				type_identifier();
				}
			}

			setState(1533);
			match(DOT);
			setState(1534);
			enum_case_name();
			setState(1536);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,169,_ctx) ) {
			case 1:
				{
				setState(1535);
				tuple_pattern();
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

	public static class Optional_patternContext extends ParserRuleContext {
		public Identifier_patternContext identifier_pattern() {
			return getRuleContext(Identifier_patternContext.class,0);
		}
		public TerminalNode QUESTION() { return getToken(TswiftParser.QUESTION, 0); }
		public Optional_patternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_optional_pattern; }
	}

	public final Optional_patternContext optional_pattern() throws RecognitionException {
		Optional_patternContext _localctx = new Optional_patternContext(_ctx, getState());
		enterRule(_localctx, 270, RULE_optional_pattern);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1538);
			identifier_pattern();
			setState(1539);
			match(QUESTION);
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

	public static class Expresion_patternContext extends ParserRuleContext {
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Expresion_patternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expresion_pattern; }
	}

	public final Expresion_patternContext expresion_pattern() throws RecognitionException {
		Expresion_patternContext _localctx = new Expresion_patternContext(_ctx, getState());
		enterRule(_localctx, 272, RULE_expresion_pattern);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1541);
			expresion();
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

	public static class AttributeContext extends ParserRuleContext {
		public TerminalNode AT() { return getToken(TswiftParser.AT, 0); }
		public Attribute_nameContext attribute_name() {
			return getRuleContext(Attribute_nameContext.class,0);
		}
		public Attribute_argument_clauseContext attribute_argument_clause() {
			return getRuleContext(Attribute_argument_clauseContext.class,0);
		}
		public AttributeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attribute; }
	}

	public final AttributeContext attribute() throws RecognitionException {
		AttributeContext _localctx = new AttributeContext(_ctx, getState());
		enterRule(_localctx, 274, RULE_attribute);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1543);
			match(AT);
			setState(1544);
			attribute_name();
			setState(1546);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,170,_ctx) ) {
			case 1:
				{
				setState(1545);
				attribute_argument_clause();
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

	public static class Attribute_nameContext extends ParserRuleContext {
		public Declaracion_identifierContext declaracion_identifier() {
			return getRuleContext(Declaracion_identifierContext.class,0);
		}
		public Attribute_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attribute_name; }
	}

	public final Attribute_nameContext attribute_name() throws RecognitionException {
		Attribute_nameContext _localctx = new Attribute_nameContext(_ctx, getState());
		enterRule(_localctx, 276, RULE_attribute_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1548);
			declaracion_identifier();
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

	public static class Attribute_argument_clauseContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(TswiftParser.LPAREN, 0); }
		public Balanced_tokensContext balanced_tokens() {
			return getRuleContext(Balanced_tokensContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(TswiftParser.RPAREN, 0); }
		public Attribute_argument_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attribute_argument_clause; }
	}

	public final Attribute_argument_clauseContext attribute_argument_clause() throws RecognitionException {
		Attribute_argument_clauseContext _localctx = new Attribute_argument_clauseContext(_ctx, getState());
		enterRule(_localctx, 278, RULE_attribute_argument_clause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1550);
			match(LPAREN);
			setState(1551);
			balanced_tokens();
			setState(1552);
			match(RPAREN);
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

	public static class AttributesContext extends ParserRuleContext {
		public List<AttributeContext> attribute() {
			return getRuleContexts(AttributeContext.class);
		}
		public AttributeContext attribute(int i) {
			return getRuleContext(AttributeContext.class,i);
		}
		public AttributesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attributes; }
	}

	public final AttributesContext attributes() throws RecognitionException {
		AttributesContext _localctx = new AttributesContext(_ctx, getState());
		enterRule(_localctx, 280, RULE_attributes);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1555); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(1554);
					attribute();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(1557); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,171,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
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

	public static class Balanced_tokensContext extends ParserRuleContext {
		public List<Balanced_tokenContext> balanced_token() {
			return getRuleContexts(Balanced_tokenContext.class);
		}
		public Balanced_tokenContext balanced_token(int i) {
			return getRuleContext(Balanced_tokenContext.class,i);
		}
		public Balanced_tokensContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_balanced_tokens; }
	}

	public final Balanced_tokensContext balanced_tokens() throws RecognitionException {
		Balanced_tokensContext _localctx = new Balanced_tokensContext(_ctx, getState());
		enterRule(_localctx, 282, RULE_balanced_tokens);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1562);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,172,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1559);
					balanced_token();
					}
					} 
				}
				setState(1564);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,172,_ctx);
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

	public static class Balanced_tokenContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(TswiftParser.LPAREN, 0); }
		public Balanced_tokensContext balanced_tokens() {
			return getRuleContext(Balanced_tokensContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(TswiftParser.RPAREN, 0); }
		public TerminalNode LBRACK() { return getToken(TswiftParser.LBRACK, 0); }
		public TerminalNode RBRACK() { return getToken(TswiftParser.RBRACK, 0); }
		public TerminalNode LCURLY() { return getToken(TswiftParser.LCURLY, 0); }
		public TerminalNode RCURLY() { return getToken(TswiftParser.RCURLY, 0); }
		public Label_identifierContext label_identifier() {
			return getRuleContext(Label_identifierContext.class,0);
		}
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public Operator_Context operator_() {
			return getRuleContext(Operator_Context.class,0);
		}
		public Any_punctuation_for_balanced_tokenContext any_punctuation_for_balanced_token() {
			return getRuleContext(Any_punctuation_for_balanced_tokenContext.class,0);
		}
		public Balanced_tokenContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_balanced_token; }
	}

	public final Balanced_tokenContext balanced_token() throws RecognitionException {
		Balanced_tokenContext _localctx = new Balanced_tokenContext(_ctx, getState());
		enterRule(_localctx, 284, RULE_balanced_token);
		try {
			setState(1581);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,173,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1565);
				match(LPAREN);
				setState(1566);
				balanced_tokens();
				setState(1567);
				match(RPAREN);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1569);
				match(LBRACK);
				setState(1570);
				balanced_tokens();
				setState(1571);
				match(RBRACK);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1573);
				match(LCURLY);
				setState(1574);
				balanced_tokens();
				setState(1575);
				match(RCURLY);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(1577);
				label_identifier();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(1578);
				literal();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(1579);
				operator_();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(1580);
				any_punctuation_for_balanced_token();
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

	public static class Any_punctuation_for_balanced_tokenContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(TswiftParser.DOT, 0); }
		public TerminalNode COMMA() { return getToken(TswiftParser.COMMA, 0); }
		public TerminalNode COLON() { return getToken(TswiftParser.COLON, 0); }
		public TerminalNode SEMI() { return getToken(TswiftParser.SEMI, 0); }
		public TerminalNode EQUAL() { return getToken(TswiftParser.EQUAL, 0); }
		public TerminalNode AT() { return getToken(TswiftParser.AT, 0); }
		public TerminalNode QUESTION() { return getToken(TswiftParser.QUESTION, 0); }
		public Arrow_operatorContext arrow_operator() {
			return getRuleContext(Arrow_operatorContext.class,0);
		}
		public TerminalNode AND() { return getToken(TswiftParser.AND, 0); }
		public TerminalNode BANG() { return getToken(TswiftParser.BANG, 0); }
		public Any_punctuation_for_balanced_tokenContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_any_punctuation_for_balanced_token; }
	}

	public final Any_punctuation_for_balanced_tokenContext any_punctuation_for_balanced_token() throws RecognitionException {
		Any_punctuation_for_balanced_tokenContext _localctx = new Any_punctuation_for_balanced_tokenContext(_ctx, getState());
		enterRule(_localctx, 286, RULE_any_punctuation_for_balanced_token);
		int _la;
		try {
			setState(1589);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,174,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1583);
				_la = _input.LA(1);
				if ( !(((((_la - 67)) & ~0x3f) == 0 && ((1L << (_la - 67)) & ((1L << (T__66 - 67)) | (1L << (T__67 - 67)) | (1L << (DOT - 67)))) != 0) || ((((_la - 134)) & ~0x3f) == 0 && ((1L << (_la - 134)) & ((1L << (COMMA - 134)) | (1L << (COLON - 134)) | (1L << (SEMI - 134)) | (1L << (QUESTION - 134)) | (1L << (AT - 134)) | (1L << (EQUAL - 134)))) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1584);
				arrow_operator();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1585);
				if (!(SwiftSupport.isPrefixOp(_input))) throw new FailedPredicateException(this, "SwiftSupport.isPrefixOp(_input)");
				setState(1586);
				match(AND);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(1587);
				if (!(SwiftSupport.isPostfixOp(_input))) throw new FailedPredicateException(this, "SwiftSupport.isPostfixOp(_input)");
				setState(1588);
				match(BANG);
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
		public Prefix_expresionContext prefix_expresion() {
			return getRuleContext(Prefix_expresionContext.class,0);
		}
		public Try_operatorContext try_operator() {
			return getRuleContext(Try_operatorContext.class,0);
		}
		public Binary_expresionsContext binary_expresions() {
			return getRuleContext(Binary_expresionsContext.class,0);
		}
		public ExpresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expresion; }
	}

	public final ExpresionContext expresion() throws RecognitionException {
		ExpresionContext _localctx = new ExpresionContext(_ctx, getState());
		enterRule(_localctx, 288, RULE_expresion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1592);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,175,_ctx) ) {
			case 1:
				{
				setState(1591);
				try_operator();
				}
				break;
			}
			setState(1594);
			prefix_expresion();
			setState(1596);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,176,_ctx) ) {
			case 1:
				{
				setState(1595);
				binary_expresions();
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

	public static class Expresion_listContext extends ParserRuleContext {
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(TswiftParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(TswiftParser.COMMA, i);
		}
		public Expresion_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expresion_list; }
	}

	public final Expresion_listContext expresion_list() throws RecognitionException {
		Expresion_listContext _localctx = new Expresion_listContext(_ctx, getState());
		enterRule(_localctx, 290, RULE_expresion_list);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1598);
			expresion();
			setState(1603);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(1599);
				match(COMMA);
				setState(1600);
				expresion();
				}
				}
				setState(1605);
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

	public static class Prefix_expresionContext extends ParserRuleContext {
		public Prefix_operatorContext prefix_operator() {
			return getRuleContext(Prefix_operatorContext.class,0);
		}
		public Postfix_expresionContext postfix_expresion() {
			return getRuleContext(Postfix_expresionContext.class,0);
		}
		public In_out_expresionContext in_out_expresion() {
			return getRuleContext(In_out_expresionContext.class,0);
		}
		public Prefix_expresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prefix_expresion; }
	}

	public final Prefix_expresionContext prefix_expresion() throws RecognitionException {
		Prefix_expresionContext _localctx = new Prefix_expresionContext(_ctx, getState());
		enterRule(_localctx, 292, RULE_prefix_expresion);
		try {
			setState(1611);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,178,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1606);
				prefix_operator();
				setState(1607);
				postfix_expresion(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1609);
				postfix_expresion(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1610);
				in_out_expresion();
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

	public static class In_out_expresionContext extends ParserRuleContext {
		public TerminalNode AND() { return getToken(TswiftParser.AND, 0); }
		public Declaracion_identifierContext declaracion_identifier() {
			return getRuleContext(Declaracion_identifierContext.class,0);
		}
		public In_out_expresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_in_out_expresion; }
	}

	public final In_out_expresionContext in_out_expresion() throws RecognitionException {
		In_out_expresionContext _localctx = new In_out_expresionContext(_ctx, getState());
		enterRule(_localctx, 294, RULE_in_out_expresion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1613);
			match(AND);
			setState(1614);
			declaracion_identifier();
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

	public static class Try_operatorContext extends ParserRuleContext {
		public TerminalNode QUESTION() { return getToken(TswiftParser.QUESTION, 0); }
		public TerminalNode BANG() { return getToken(TswiftParser.BANG, 0); }
		public Try_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_try_operator; }
	}

	public final Try_operatorContext try_operator() throws RecognitionException {
		Try_operatorContext _localctx = new Try_operatorContext(_ctx, getState());
		enterRule(_localctx, 296, RULE_try_operator);
		try {
			setState(1621);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,179,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1616);
				match(T__68);
				setState(1617);
				match(QUESTION);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1618);
				match(T__68);
				setState(1619);
				match(BANG);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1620);
				match(T__68);
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

	public static class Binary_expresionContext extends ParserRuleContext {
		public Binary_operatorContext binary_operator() {
			return getRuleContext(Binary_operatorContext.class,0);
		}
		public Prefix_expresionContext prefix_expresion() {
			return getRuleContext(Prefix_expresionContext.class,0);
		}
		public Assignment_operatorContext assignment_operator() {
			return getRuleContext(Assignment_operatorContext.class,0);
		}
		public Try_operatorContext try_operator() {
			return getRuleContext(Try_operatorContext.class,0);
		}
		public Conditional_operatorContext conditional_operator() {
			return getRuleContext(Conditional_operatorContext.class,0);
		}
		public Type_casting_operatorContext type_casting_operator() {
			return getRuleContext(Type_casting_operatorContext.class,0);
		}
		public Binary_expresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_binary_expresion; }
	}

	public final Binary_expresionContext binary_expresion() throws RecognitionException {
		Binary_expresionContext _localctx = new Binary_expresionContext(_ctx, getState());
		enterRule(_localctx, 298, RULE_binary_expresion);
		try {
			setState(1639);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,182,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1623);
				binary_operator();
				setState(1624);
				prefix_expresion();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1626);
				assignment_operator();
				setState(1628);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,180,_ctx) ) {
				case 1:
					{
					setState(1627);
					try_operator();
					}
					break;
				}
				setState(1630);
				prefix_expresion();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1632);
				conditional_operator();
				setState(1634);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,181,_ctx) ) {
				case 1:
					{
					setState(1633);
					try_operator();
					}
					break;
				}
				setState(1636);
				prefix_expresion();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(1638);
				type_casting_operator();
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

	public static class Binary_expresionsContext extends ParserRuleContext {
		public List<Binary_expresionContext> binary_expresion() {
			return getRuleContexts(Binary_expresionContext.class);
		}
		public Binary_expresionContext binary_expresion(int i) {
			return getRuleContext(Binary_expresionContext.class,i);
		}
		public Binary_expresionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_binary_expresions; }
	}

	public final Binary_expresionsContext binary_expresions() throws RecognitionException {
		Binary_expresionsContext _localctx = new Binary_expresionsContext(_ctx, getState());
		enterRule(_localctx, 300, RULE_binary_expresions);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1642); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(1641);
					binary_expresion();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(1644); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,183,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
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

	public static class Conditional_operatorContext extends ParserRuleContext {
		public TerminalNode QUESTION() { return getToken(TswiftParser.QUESTION, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode COLON() { return getToken(TswiftParser.COLON, 0); }
		public Try_operatorContext try_operator() {
			return getRuleContext(Try_operatorContext.class,0);
		}
		public Conditional_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditional_operator; }
	}

	public final Conditional_operatorContext conditional_operator() throws RecognitionException {
		Conditional_operatorContext _localctx = new Conditional_operatorContext(_ctx, getState());
		enterRule(_localctx, 302, RULE_conditional_operator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1646);
			match(QUESTION);
			setState(1648);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,184,_ctx) ) {
			case 1:
				{
				setState(1647);
				try_operator();
				}
				break;
			}
			setState(1650);
			expresion();
			setState(1651);
			match(COLON);
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

	public static class Type_casting_operatorContext extends ParserRuleContext {
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public TerminalNode QUESTION() { return getToken(TswiftParser.QUESTION, 0); }
		public TerminalNode BANG() { return getToken(TswiftParser.BANG, 0); }
		public Type_casting_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_casting_operator; }
	}

	public final Type_casting_operatorContext type_casting_operator() throws RecognitionException {
		Type_casting_operatorContext _localctx = new Type_casting_operatorContext(_ctx, getState());
		enterRule(_localctx, 304, RULE_type_casting_operator);
		try {
			setState(1663);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,185,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1653);
				match(T__64);
				setState(1654);
				type_(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1655);
				match(T__65);
				setState(1656);
				type_(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1657);
				match(T__65);
				setState(1658);
				match(QUESTION);
				setState(1659);
				type_(0);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(1660);
				match(T__65);
				setState(1661);
				match(BANG);
				setState(1662);
				type_(0);
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

	public static class Primary_expresionContext extends ParserRuleContext {
		public Declaracion_identifierContext declaracion_identifier() {
			return getRuleContext(Declaracion_identifierContext.class,0);
		}
		public Generic_argument_clauseContext generic_argument_clause() {
			return getRuleContext(Generic_argument_clauseContext.class,0);
		}
		public Literal_expresionContext literal_expresion() {
			return getRuleContext(Literal_expresionContext.class,0);
		}
		public Self_expresionContext self_expresion() {
			return getRuleContext(Self_expresionContext.class,0);
		}
		public Superclass_expresionContext superclass_expresion() {
			return getRuleContext(Superclass_expresionContext.class,0);
		}
		public Closure_expresionContext closure_expresion() {
			return getRuleContext(Closure_expresionContext.class,0);
		}
		public Parenthesized_expresionContext parenthesized_expresion() {
			return getRuleContext(Parenthesized_expresionContext.class,0);
		}
		public Tuple_expresionContext tuple_expresion() {
			return getRuleContext(Tuple_expresionContext.class,0);
		}
		public Implicit_member_expresionContext implicit_member_expresion() {
			return getRuleContext(Implicit_member_expresionContext.class,0);
		}
		public Wildcard_expresionContext wildcard_expresion() {
			return getRuleContext(Wildcard_expresionContext.class,0);
		}
		public Selector_expresionContext selector_expresion() {
			return getRuleContext(Selector_expresionContext.class,0);
		}
		public Key_path_expresionContext key_path_expresion() {
			return getRuleContext(Key_path_expresionContext.class,0);
		}
		public Primary_expresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primary_expresion; }
	}

	public final Primary_expresionContext primary_expresion() throws RecognitionException {
		Primary_expresionContext _localctx = new Primary_expresionContext(_ctx, getState());
		enterRule(_localctx, 306, RULE_primary_expresion);
		try {
			setState(1679);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,187,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1665);
				declaracion_identifier();
				setState(1667);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,186,_ctx) ) {
				case 1:
					{
					setState(1666);
					generic_argument_clause();
					}
					break;
				}
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1669);
				literal_expresion();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1670);
				self_expresion();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(1671);
				superclass_expresion();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(1672);
				closure_expresion();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(1673);
				parenthesized_expresion();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(1674);
				tuple_expresion();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(1675);
				implicit_member_expresion();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(1676);
				wildcard_expresion();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(1677);
				selector_expresion();
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(1678);
				key_path_expresion();
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

	public static class Literal_expresionContext extends ParserRuleContext {
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public Array_literalContext array_literal() {
			return getRuleContext(Array_literalContext.class,0);
		}
		public Dictionary_literalContext dictionary_literal() {
			return getRuleContext(Dictionary_literalContext.class,0);
		}
		public Literal_expresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal_expresion; }
	}

	public final Literal_expresionContext literal_expresion() throws RecognitionException {
		Literal_expresionContext _localctx = new Literal_expresionContext(_ctx, getState());
		enterRule(_localctx, 308, RULE_literal_expresion);
		try {
			setState(1689);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,188,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1681);
				literal();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1682);
				array_literal();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1683);
				dictionary_literal();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(1684);
				match(T__69);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(1685);
				match(T__70);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(1686);
				match(T__71);
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(1687);
				match(T__72);
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(1688);
				match(T__73);
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

	public static class Array_literalContext extends ParserRuleContext {
		public TerminalNode LBRACK() { return getToken(TswiftParser.LBRACK, 0); }
		public TerminalNode RBRACK() { return getToken(TswiftParser.RBRACK, 0); }
		public Array_literal_itemsContext array_literal_items() {
			return getRuleContext(Array_literal_itemsContext.class,0);
		}
		public Array_literalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_array_literal; }
	}

	public final Array_literalContext array_literal() throws RecognitionException {
		Array_literalContext _localctx = new Array_literalContext(_ctx, getState());
		enterRule(_localctx, 310, RULE_array_literal);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1691);
			match(LBRACK);
			setState(1693);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,189,_ctx) ) {
			case 1:
				{
				setState(1692);
				array_literal_items();
				}
				break;
			}
			setState(1695);
			match(RBRACK);
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

	public static class Array_literal_itemsContext extends ParserRuleContext {
		public Array_literal_itemContext array_literal_item() {
			return getRuleContext(Array_literal_itemContext.class,0);
		}
		public TerminalNode COMMA() { return getToken(TswiftParser.COMMA, 0); }
		public Array_literal_itemsContext array_literal_items() {
			return getRuleContext(Array_literal_itemsContext.class,0);
		}
		public Array_literal_itemsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_array_literal_items; }
	}

	public final Array_literal_itemsContext array_literal_items() throws RecognitionException {
		Array_literal_itemsContext _localctx = new Array_literal_itemsContext(_ctx, getState());
		enterRule(_localctx, 312, RULE_array_literal_items);
		int _la;
		try {
			setState(1705);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,191,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1697);
				array_literal_item();
				setState(1699);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COMMA) {
					{
					setState(1698);
					match(COMMA);
					}
				}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1701);
				array_literal_item();
				setState(1702);
				match(COMMA);
				setState(1703);
				array_literal_items();
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

	public static class Array_literal_itemContext extends ParserRuleContext {
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Array_literal_itemContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_array_literal_item; }
	}

	public final Array_literal_itemContext array_literal_item() throws RecognitionException {
		Array_literal_itemContext _localctx = new Array_literal_itemContext(_ctx, getState());
		enterRule(_localctx, 314, RULE_array_literal_item);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1707);
			expresion();
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

	public static class Dictionary_literalContext extends ParserRuleContext {
		public TerminalNode LBRACK() { return getToken(TswiftParser.LBRACK, 0); }
		public Dictionary_literal_itemsContext dictionary_literal_items() {
			return getRuleContext(Dictionary_literal_itemsContext.class,0);
		}
		public TerminalNode RBRACK() { return getToken(TswiftParser.RBRACK, 0); }
		public TerminalNode COLON() { return getToken(TswiftParser.COLON, 0); }
		public Dictionary_literalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dictionary_literal; }
	}

	public final Dictionary_literalContext dictionary_literal() throws RecognitionException {
		Dictionary_literalContext _localctx = new Dictionary_literalContext(_ctx, getState());
		enterRule(_localctx, 316, RULE_dictionary_literal);
		try {
			setState(1716);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,192,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1709);
				match(LBRACK);
				setState(1710);
				dictionary_literal_items();
				setState(1711);
				match(RBRACK);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1713);
				match(LBRACK);
				setState(1714);
				match(COLON);
				setState(1715);
				match(RBRACK);
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

	public static class Dictionary_literal_itemsContext extends ParserRuleContext {
		public Dictionary_literal_itemContext dictionary_literal_item() {
			return getRuleContext(Dictionary_literal_itemContext.class,0);
		}
		public TerminalNode COMMA() { return getToken(TswiftParser.COMMA, 0); }
		public Dictionary_literal_itemsContext dictionary_literal_items() {
			return getRuleContext(Dictionary_literal_itemsContext.class,0);
		}
		public Dictionary_literal_itemsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dictionary_literal_items; }
	}

	public final Dictionary_literal_itemsContext dictionary_literal_items() throws RecognitionException {
		Dictionary_literal_itemsContext _localctx = new Dictionary_literal_itemsContext(_ctx, getState());
		enterRule(_localctx, 318, RULE_dictionary_literal_items);
		int _la;
		try {
			setState(1726);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,194,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1718);
				dictionary_literal_item();
				setState(1720);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COMMA) {
					{
					setState(1719);
					match(COMMA);
					}
				}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1722);
				dictionary_literal_item();
				setState(1723);
				match(COMMA);
				setState(1724);
				dictionary_literal_items();
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

	public static class Dictionary_literal_itemContext extends ParserRuleContext {
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode COLON() { return getToken(TswiftParser.COLON, 0); }
		public Dictionary_literal_itemContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dictionary_literal_item; }
	}

	public final Dictionary_literal_itemContext dictionary_literal_item() throws RecognitionException {
		Dictionary_literal_itemContext _localctx = new Dictionary_literal_itemContext(_ctx, getState());
		enterRule(_localctx, 320, RULE_dictionary_literal_item);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1728);
			expresion();
			setState(1729);
			match(COLON);
			setState(1730);
			expresion();
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

	public static class Playground_literalContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(TswiftParser.LPAREN, 0); }
		public List<TerminalNode> COLON() { return getTokens(TswiftParser.COLON); }
		public TerminalNode COLON(int i) {
			return getToken(TswiftParser.COLON, i);
		}
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(TswiftParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(TswiftParser.COMMA, i);
		}
		public TerminalNode RPAREN() { return getToken(TswiftParser.RPAREN, 0); }
		public Playground_literalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_playground_literal; }
	}

	public final Playground_literalContext playground_literal() throws RecognitionException {
		Playground_literalContext _localctx = new Playground_literalContext(_ctx, getState());
		enterRule(_localctx, 322, RULE_playground_literal);
		try {
			setState(1765);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__74:
				enterOuterAlt(_localctx, 1);
				{
				setState(1732);
				match(T__74);
				setState(1733);
				match(LPAREN);
				setState(1734);
				match(T__75);
				setState(1735);
				match(COLON);
				setState(1736);
				expresion();
				setState(1737);
				match(COMMA);
				setState(1738);
				match(T__76);
				setState(1739);
				match(COLON);
				setState(1740);
				expresion();
				setState(1741);
				match(COMMA);
				setState(1742);
				match(T__77);
				setState(1743);
				match(COLON);
				setState(1744);
				expresion();
				setState(1745);
				match(COMMA);
				setState(1746);
				match(T__78);
				setState(1747);
				match(COLON);
				setState(1748);
				expresion();
				setState(1749);
				match(RPAREN);
				}
				break;
			case T__79:
				enterOuterAlt(_localctx, 2);
				{
				setState(1751);
				match(T__79);
				setState(1752);
				match(LPAREN);
				setState(1753);
				match(T__80);
				setState(1754);
				match(COLON);
				setState(1755);
				expresion();
				setState(1756);
				match(RPAREN);
				}
				break;
			case T__81:
				enterOuterAlt(_localctx, 3);
				{
				setState(1758);
				match(T__81);
				setState(1759);
				match(LPAREN);
				setState(1760);
				match(T__80);
				setState(1761);
				match(COLON);
				setState(1762);
				expresion();
				setState(1763);
				match(RPAREN);
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

	public static class Self_expresionContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(TswiftParser.DOT, 0); }
		public Declaracion_identifierContext declaracion_identifier() {
			return getRuleContext(Declaracion_identifierContext.class,0);
		}
		public TerminalNode LBRACK() { return getToken(TswiftParser.LBRACK, 0); }
		public Expresion_listContext expresion_list() {
			return getRuleContext(Expresion_listContext.class,0);
		}
		public TerminalNode RBRACK() { return getToken(TswiftParser.RBRACK, 0); }
		public Self_expresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_self_expresion; }
	}

	public final Self_expresionContext self_expresion() throws RecognitionException {
		Self_expresionContext _localctx = new Self_expresionContext(_ctx, getState());
		enterRule(_localctx, 324, RULE_self_expresion);
		try {
			setState(1786);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,196,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1767);
				match(T__82);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1768);
				match(T__82);
				setState(1769);
				match(DOT);
				setState(1770);
				declaracion_identifier();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1771);
				match(T__82);
				setState(1772);
				match(LBRACK);
				setState(1773);
				expresion_list();
				setState(1774);
				match(RBRACK);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(1776);
				match(T__82);
				setState(1777);
				match(DOT);
				setState(1778);
				match(T__30);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(1779);
				match(T__83);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(1780);
				match(T__83);
				setState(1781);
				match(DOT);
				setState(1782);
				declaracion_identifier();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(1783);
				match(T__83);
				setState(1784);
				match(DOT);
				setState(1785);
				match(T__30);
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

	public static class Superclass_expresionContext extends ParserRuleContext {
		public Superclass_method_expresionContext superclass_method_expresion() {
			return getRuleContext(Superclass_method_expresionContext.class,0);
		}
		public Superclass_subscript_expresionContext superclass_subscript_expresion() {
			return getRuleContext(Superclass_subscript_expresionContext.class,0);
		}
		public Superclass_initializer_expresionContext superclass_initializer_expresion() {
			return getRuleContext(Superclass_initializer_expresionContext.class,0);
		}
		public Superclass_expresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_superclass_expresion; }
	}

	public final Superclass_expresionContext superclass_expresion() throws RecognitionException {
		Superclass_expresionContext _localctx = new Superclass_expresionContext(_ctx, getState());
		enterRule(_localctx, 326, RULE_superclass_expresion);
		try {
			setState(1791);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,197,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1788);
				superclass_method_expresion();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1789);
				superclass_subscript_expresion();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1790);
				superclass_initializer_expresion();
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

	public static class Superclass_method_expresionContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(TswiftParser.DOT, 0); }
		public Declaracion_identifierContext declaracion_identifier() {
			return getRuleContext(Declaracion_identifierContext.class,0);
		}
		public Superclass_method_expresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_superclass_method_expresion; }
	}

	public final Superclass_method_expresionContext superclass_method_expresion() throws RecognitionException {
		Superclass_method_expresionContext _localctx = new Superclass_method_expresionContext(_ctx, getState());
		enterRule(_localctx, 328, RULE_superclass_method_expresion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1793);
			match(T__84);
			setState(1794);
			match(DOT);
			setState(1795);
			declaracion_identifier();
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

	public static class Superclass_subscript_expresionContext extends ParserRuleContext {
		public TerminalNode LBRACK() { return getToken(TswiftParser.LBRACK, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RBRACK() { return getToken(TswiftParser.RBRACK, 0); }
		public Superclass_subscript_expresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_superclass_subscript_expresion; }
	}

	public final Superclass_subscript_expresionContext superclass_subscript_expresion() throws RecognitionException {
		Superclass_subscript_expresionContext _localctx = new Superclass_subscript_expresionContext(_ctx, getState());
		enterRule(_localctx, 330, RULE_superclass_subscript_expresion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1797);
			match(T__84);
			setState(1798);
			match(LBRACK);
			setState(1799);
			expresion();
			setState(1800);
			match(RBRACK);
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

	public static class Superclass_initializer_expresionContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(TswiftParser.DOT, 0); }
		public Superclass_initializer_expresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_superclass_initializer_expresion; }
	}

	public final Superclass_initializer_expresionContext superclass_initializer_expresion() throws RecognitionException {
		Superclass_initializer_expresionContext _localctx = new Superclass_initializer_expresionContext(_ctx, getState());
		enterRule(_localctx, 332, RULE_superclass_initializer_expresion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1802);
			match(T__84);
			setState(1803);
			match(DOT);
			setState(1804);
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

	public static class Closure_expresionContext extends ParserRuleContext {
		public TerminalNode LCURLY() { return getToken(TswiftParser.LCURLY, 0); }
		public TerminalNode RCURLY() { return getToken(TswiftParser.RCURLY, 0); }
		public Closure_signatureContext closure_signature() {
			return getRuleContext(Closure_signatureContext.class,0);
		}
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public Closure_expresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_closure_expresion; }
	}

	public final Closure_expresionContext closure_expresion() throws RecognitionException {
		Closure_expresionContext _localctx = new Closure_expresionContext(_ctx, getState());
		enterRule(_localctx, 334, RULE_closure_expresion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1806);
			match(LCURLY);
			setState(1808);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,198,_ctx) ) {
			case 1:
				{
				setState(1807);
				closure_signature();
				}
				break;
			}
			setState(1811);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,199,_ctx) ) {
			case 1:
				{
				setState(1810);
				instrucciones();
				}
				break;
			}
			setState(1813);
			match(RCURLY);
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

	public static class Closure_signatureContext extends ParserRuleContext {
		public Closure_parameter_clauseContext closure_parameter_clause() {
			return getRuleContext(Closure_parameter_clauseContext.class,0);
		}
		public Capture_listContext capture_list() {
			return getRuleContext(Capture_listContext.class,0);
		}
		public Function_resultContext function_result() {
			return getRuleContext(Function_resultContext.class,0);
		}
		public Closure_signatureContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_closure_signature; }
	}

	public final Closure_signatureContext closure_signature() throws RecognitionException {
		Closure_signatureContext _localctx = new Closure_signatureContext(_ctx, getState());
		enterRule(_localctx, 336, RULE_closure_signature);
		int _la;
		try {
			setState(1830);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,203,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1816);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LBRACK) {
					{
					setState(1815);
					capture_list();
					}
				}

				setState(1818);
				closure_parameter_clause();
				setState(1820);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,201,_ctx) ) {
				case 1:
					{
					setState(1819);
					match(T__25);
					}
					break;
				}
				setState(1823);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,202,_ctx) ) {
				case 1:
					{
					setState(1822);
					function_result();
					}
					break;
				}
				setState(1825);
				match(T__1);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1827);
				capture_list();
				setState(1828);
				match(T__1);
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

	public static class Closure_parameter_clauseContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(TswiftParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(TswiftParser.RPAREN, 0); }
		public Closure_parameter_listContext closure_parameter_list() {
			return getRuleContext(Closure_parameter_listContext.class,0);
		}
		public Closure_parameter_clause_identifier_listContext closure_parameter_clause_identifier_list() {
			return getRuleContext(Closure_parameter_clause_identifier_listContext.class,0);
		}
		public Closure_parameter_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_closure_parameter_clause; }
	}

	public final Closure_parameter_clauseContext closure_parameter_clause() throws RecognitionException {
		Closure_parameter_clauseContext _localctx = new Closure_parameter_clauseContext(_ctx, getState());
		enterRule(_localctx, 338, RULE_closure_parameter_clause);
		try {
			setState(1839);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,204,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1832);
				match(LPAREN);
				setState(1833);
				match(RPAREN);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1834);
				match(LPAREN);
				setState(1835);
				closure_parameter_list();
				setState(1836);
				match(RPAREN);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1838);
				closure_parameter_clause_identifier_list();
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

	public static class Closure_parameter_clause_identifier_listContext extends ParserRuleContext {
		public List<Declaracion_identifierContext> declaracion_identifier() {
			return getRuleContexts(Declaracion_identifierContext.class);
		}
		public Declaracion_identifierContext declaracion_identifier(int i) {
			return getRuleContext(Declaracion_identifierContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(TswiftParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(TswiftParser.COMMA, i);
		}
		public Closure_parameter_clause_identifier_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_closure_parameter_clause_identifier_list; }
	}

	public final Closure_parameter_clause_identifier_listContext closure_parameter_clause_identifier_list() throws RecognitionException {
		Closure_parameter_clause_identifier_listContext _localctx = new Closure_parameter_clause_identifier_listContext(_ctx, getState());
		enterRule(_localctx, 340, RULE_closure_parameter_clause_identifier_list);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1841);
			declaracion_identifier();
			setState(1846);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,205,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1842);
					match(COMMA);
					setState(1843);
					declaracion_identifier();
					}
					} 
				}
				setState(1848);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,205,_ctx);
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

	public static class Closure_parameter_listContext extends ParserRuleContext {
		public List<Closure_parameterContext> closure_parameter() {
			return getRuleContexts(Closure_parameterContext.class);
		}
		public Closure_parameterContext closure_parameter(int i) {
			return getRuleContext(Closure_parameterContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(TswiftParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(TswiftParser.COMMA, i);
		}
		public Closure_parameter_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_closure_parameter_list; }
	}

	public final Closure_parameter_listContext closure_parameter_list() throws RecognitionException {
		Closure_parameter_listContext _localctx = new Closure_parameter_listContext(_ctx, getState());
		enterRule(_localctx, 342, RULE_closure_parameter_list);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1849);
			closure_parameter();
			setState(1854);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(1850);
				match(COMMA);
				setState(1851);
				closure_parameter();
				}
				}
				setState(1856);
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

	public static class Closure_parameterContext extends ParserRuleContext {
		public Closure_parameter_nameContext closure_parameter_name() {
			return getRuleContext(Closure_parameter_nameContext.class,0);
		}
		public Type_annotationContext type_annotation() {
			return getRuleContext(Type_annotationContext.class,0);
		}
		public Range_operatorContext range_operator() {
			return getRuleContext(Range_operatorContext.class,0);
		}
		public Closure_parameterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_closure_parameter; }
	}

	public final Closure_parameterContext closure_parameter() throws RecognitionException {
		Closure_parameterContext _localctx = new Closure_parameterContext(_ctx, getState());
		enterRule(_localctx, 344, RULE_closure_parameter);
		int _la;
		try {
			setState(1865);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,208,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1857);
				closure_parameter_name();
				setState(1859);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COLON) {
					{
					setState(1858);
					type_annotation();
					}
				}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1861);
				closure_parameter_name();
				setState(1862);
				type_annotation();
				setState(1863);
				range_operator();
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

	public static class Closure_parameter_nameContext extends ParserRuleContext {
		public Label_identifierContext label_identifier() {
			return getRuleContext(Label_identifierContext.class,0);
		}
		public Closure_parameter_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_closure_parameter_name; }
	}

	public final Closure_parameter_nameContext closure_parameter_name() throws RecognitionException {
		Closure_parameter_nameContext _localctx = new Closure_parameter_nameContext(_ctx, getState());
		enterRule(_localctx, 346, RULE_closure_parameter_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1867);
			label_identifier();
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

	public static class Capture_listContext extends ParserRuleContext {
		public TerminalNode LBRACK() { return getToken(TswiftParser.LBRACK, 0); }
		public Capture_list_itemsContext capture_list_items() {
			return getRuleContext(Capture_list_itemsContext.class,0);
		}
		public TerminalNode RBRACK() { return getToken(TswiftParser.RBRACK, 0); }
		public Capture_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_capture_list; }
	}

	public final Capture_listContext capture_list() throws RecognitionException {
		Capture_listContext _localctx = new Capture_listContext(_ctx, getState());
		enterRule(_localctx, 348, RULE_capture_list);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1869);
			match(LBRACK);
			setState(1870);
			capture_list_items();
			setState(1871);
			match(RBRACK);
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

	public static class Capture_list_itemsContext extends ParserRuleContext {
		public List<Capture_list_itemContext> capture_list_item() {
			return getRuleContexts(Capture_list_itemContext.class);
		}
		public Capture_list_itemContext capture_list_item(int i) {
			return getRuleContext(Capture_list_itemContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(TswiftParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(TswiftParser.COMMA, i);
		}
		public Capture_list_itemsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_capture_list_items; }
	}

	public final Capture_list_itemsContext capture_list_items() throws RecognitionException {
		Capture_list_itemsContext _localctx = new Capture_list_itemsContext(_ctx, getState());
		enterRule(_localctx, 350, RULE_capture_list_items);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1873);
			capture_list_item();
			setState(1878);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(1874);
				match(COMMA);
				setState(1875);
				capture_list_item();
				}
				}
				setState(1880);
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

	public static class Capture_list_itemContext extends ParserRuleContext {
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Capture_specifierContext capture_specifier() {
			return getRuleContext(Capture_specifierContext.class,0);
		}
		public Capture_list_itemContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_capture_list_item; }
	}

	public final Capture_list_itemContext capture_list_item() throws RecognitionException {
		Capture_list_itemContext _localctx = new Capture_list_itemContext(_ctx, getState());
		enterRule(_localctx, 352, RULE_capture_list_item);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1882);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,210,_ctx) ) {
			case 1:
				{
				setState(1881);
				capture_specifier();
				}
				break;
			}
			setState(1884);
			expresion();
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

	public static class Capture_specifierContext extends ParserRuleContext {
		public Capture_specifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_capture_specifier; }
	}

	public final Capture_specifierContext capture_specifier() throws RecognitionException {
		Capture_specifierContext _localctx = new Capture_specifierContext(_ctx, getState());
		enterRule(_localctx, 354, RULE_capture_specifier);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1886);
			_la = _input.LA(1);
			if ( !(((((_la - 54)) & ~0x3f) == 0 && ((1L << (_la - 54)) & ((1L << (T__53 - 54)) | (1L << (T__56 - 54)) | (1L << (T__85 - 54)) | (1L << (T__86 - 54)))) != 0)) ) {
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

	public static class Implicit_member_expresionContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(TswiftParser.DOT, 0); }
		public Label_identifierContext label_identifier() {
			return getRuleContext(Label_identifierContext.class,0);
		}
		public Implicit_member_expresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_implicit_member_expresion; }
	}

	public final Implicit_member_expresionContext implicit_member_expresion() throws RecognitionException {
		Implicit_member_expresionContext _localctx = new Implicit_member_expresionContext(_ctx, getState());
		enterRule(_localctx, 356, RULE_implicit_member_expresion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1888);
			match(DOT);
			setState(1889);
			label_identifier();
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

	public static class Parenthesized_expresionContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(TswiftParser.LPAREN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(TswiftParser.RPAREN, 0); }
		public Parenthesized_expresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parenthesized_expresion; }
	}

	public final Parenthesized_expresionContext parenthesized_expresion() throws RecognitionException {
		Parenthesized_expresionContext _localctx = new Parenthesized_expresionContext(_ctx, getState());
		enterRule(_localctx, 358, RULE_parenthesized_expresion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1891);
			match(LPAREN);
			setState(1892);
			expresion();
			setState(1893);
			match(RPAREN);
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

	public static class Tuple_expresionContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(TswiftParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(TswiftParser.RPAREN, 0); }
		public List<Tuple_elementContext> tuple_element() {
			return getRuleContexts(Tuple_elementContext.class);
		}
		public Tuple_elementContext tuple_element(int i) {
			return getRuleContext(Tuple_elementContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(TswiftParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(TswiftParser.COMMA, i);
		}
		public Tuple_expresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuple_expresion; }
	}

	public final Tuple_expresionContext tuple_expresion() throws RecognitionException {
		Tuple_expresionContext _localctx = new Tuple_expresionContext(_ctx, getState());
		enterRule(_localctx, 360, RULE_tuple_expresion);
		int _la;
		try {
			setState(1907);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,212,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1895);
				match(LPAREN);
				setState(1896);
				match(RPAREN);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1897);
				match(LPAREN);
				setState(1898);
				tuple_element();
				setState(1901); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(1899);
					match(COMMA);
					setState(1900);
					tuple_element();
					}
					}
					setState(1903); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==COMMA );
				setState(1905);
				match(RPAREN);
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

	public static class Tuple_elementContext extends ParserRuleContext {
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Label_identifierContext label_identifier() {
			return getRuleContext(Label_identifierContext.class,0);
		}
		public TerminalNode COLON() { return getToken(TswiftParser.COLON, 0); }
		public Tuple_elementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuple_element; }
	}

	public final Tuple_elementContext tuple_element() throws RecognitionException {
		Tuple_elementContext _localctx = new Tuple_elementContext(_ctx, getState());
		enterRule(_localctx, 362, RULE_tuple_element);
		try {
			setState(1914);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,213,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1909);
				expresion();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1910);
				label_identifier();
				setState(1911);
				match(COLON);
				setState(1912);
				expresion();
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

	public static class Wildcard_expresionContext extends ParserRuleContext {
		public TerminalNode UNDERSCORE() { return getToken(TswiftParser.UNDERSCORE, 0); }
		public Wildcard_expresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_wildcard_expresion; }
	}

	public final Wildcard_expresionContext wildcard_expresion() throws RecognitionException {
		Wildcard_expresionContext _localctx = new Wildcard_expresionContext(_ctx, getState());
		enterRule(_localctx, 364, RULE_wildcard_expresion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1916);
			match(UNDERSCORE);
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

	public static class Selector_expresionContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(TswiftParser.LPAREN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(TswiftParser.RPAREN, 0); }
		public Selector_expresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_selector_expresion; }
	}

	public final Selector_expresionContext selector_expresion() throws RecognitionException {
		Selector_expresionContext _localctx = new Selector_expresionContext(_ctx, getState());
		enterRule(_localctx, 366, RULE_selector_expresion);
		try {
			setState(1935);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,214,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1918);
				match(T__87);
				setState(1919);
				match(LPAREN);
				setState(1920);
				expresion();
				setState(1921);
				match(RPAREN);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1923);
				match(T__87);
				setState(1924);
				match(LPAREN);
				setState(1925);
				match(T__88);
				setState(1926);
				expresion();
				setState(1927);
				match(RPAREN);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1929);
				match(T__87);
				setState(1930);
				match(LPAREN);
				setState(1931);
				match(T__89);
				setState(1932);
				expresion();
				setState(1933);
				match(RPAREN);
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

	public static class Key_path_expresionContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(TswiftParser.LPAREN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(TswiftParser.RPAREN, 0); }
		public Key_path_expresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_key_path_expresion; }
	}

	public final Key_path_expresionContext key_path_expresion() throws RecognitionException {
		Key_path_expresionContext _localctx = new Key_path_expresionContext(_ctx, getState());
		enterRule(_localctx, 368, RULE_key_path_expresion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1937);
			match(T__90);
			setState(1938);
			match(LPAREN);
			setState(1939);
			expresion();
			setState(1940);
			match(RPAREN);
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

	public static class Postfix_expresionContext extends ParserRuleContext {
		public Postfix_expresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_postfix_expresion; }
	 
		public Postfix_expresionContext() { }
		public void copyFrom(Postfix_expresionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Explicit_member_expresion4Context extends Postfix_expresionContext {
		public Postfix_expresionContext postfix_expresion() {
			return getRuleContext(Postfix_expresionContext.class,0);
		}
		public TerminalNode LPAREN() { return getToken(TswiftParser.LPAREN, 0); }
		public Argument_namesContext argument_names() {
			return getRuleContext(Argument_namesContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(TswiftParser.RPAREN, 0); }
		public Explicit_member_expresion4Context(Postfix_expresionContext ctx) { copyFrom(ctx); }
	}
	public static class Function_call_expresionContext extends Postfix_expresionContext {
		public Postfix_expresionContext postfix_expresion() {
			return getRuleContext(Postfix_expresionContext.class,0);
		}
		public Function_call_argument_clauseContext function_call_argument_clause() {
			return getRuleContext(Function_call_argument_clauseContext.class,0);
		}
		public Function_call_expresionContext(Postfix_expresionContext ctx) { copyFrom(ctx); }
	}
	public static class Explicit_member_expresion3Context extends Postfix_expresionContext {
		public Postfix_expresionContext postfix_expresion() {
			return getRuleContext(Postfix_expresionContext.class,0);
		}
		public TerminalNode DOT() { return getToken(TswiftParser.DOT, 0); }
		public Declaracion_identifierContext declaracion_identifier() {
			return getRuleContext(Declaracion_identifierContext.class,0);
		}
		public TerminalNode LPAREN() { return getToken(TswiftParser.LPAREN, 0); }
		public Argument_namesContext argument_names() {
			return getRuleContext(Argument_namesContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(TswiftParser.RPAREN, 0); }
		public Explicit_member_expresion3Context(Postfix_expresionContext ctx) { copyFrom(ctx); }
	}
	public static class Explicit_member_expresion2Context extends Postfix_expresionContext {
		public Postfix_expresionContext postfix_expresion() {
			return getRuleContext(Postfix_expresionContext.class,0);
		}
		public TerminalNode DOT() { return getToken(TswiftParser.DOT, 0); }
		public Declaracion_identifierContext declaracion_identifier() {
			return getRuleContext(Declaracion_identifierContext.class,0);
		}
		public Generic_argument_clauseContext generic_argument_clause() {
			return getRuleContext(Generic_argument_clauseContext.class,0);
		}
		public Explicit_member_expresion2Context(Postfix_expresionContext ctx) { copyFrom(ctx); }
	}
	public static class Explicit_member_expresion1Context extends Postfix_expresionContext {
		public Postfix_expresionContext postfix_expresion() {
			return getRuleContext(Postfix_expresionContext.class,0);
		}
		public TerminalNode DOT() { return getToken(TswiftParser.DOT, 0); }
		public TerminalNode Pure_decimal_digits() { return getToken(TswiftParser.Pure_decimal_digits, 0); }
		public Explicit_member_expresion1Context(Postfix_expresionContext ctx) { copyFrom(ctx); }
	}
	public static class Initializer_expresion_with_argsContext extends Postfix_expresionContext {
		public Postfix_expresionContext postfix_expresion() {
			return getRuleContext(Postfix_expresionContext.class,0);
		}
		public TerminalNode DOT() { return getToken(TswiftParser.DOT, 0); }
		public TerminalNode LPAREN() { return getToken(TswiftParser.LPAREN, 0); }
		public Argument_namesContext argument_names() {
			return getRuleContext(Argument_namesContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(TswiftParser.RPAREN, 0); }
		public Initializer_expresion_with_argsContext(Postfix_expresionContext ctx) { copyFrom(ctx); }
	}
	public static class Postfix_self_expresionContext extends Postfix_expresionContext {
		public Postfix_expresionContext postfix_expresion() {
			return getRuleContext(Postfix_expresionContext.class,0);
		}
		public TerminalNode DOT() { return getToken(TswiftParser.DOT, 0); }
		public Postfix_self_expresionContext(Postfix_expresionContext ctx) { copyFrom(ctx); }
	}
	public static class Dynamic_typeContext extends Postfix_expresionContext {
		public Dynamic_type_expresionContext dynamic_type_expresion() {
			return getRuleContext(Dynamic_type_expresionContext.class,0);
		}
		public Dynamic_typeContext(Postfix_expresionContext ctx) { copyFrom(ctx); }
	}
	public static class Function_call_expresion_with_closureContext extends Postfix_expresionContext {
		public Postfix_expresionContext postfix_expresion() {
			return getRuleContext(Postfix_expresionContext.class,0);
		}
		public Trailing_closureContext trailing_closure() {
			return getRuleContext(Trailing_closureContext.class,0);
		}
		public Function_call_argument_clauseContext function_call_argument_clause() {
			return getRuleContext(Function_call_argument_clauseContext.class,0);
		}
		public Function_call_expresion_with_closureContext(Postfix_expresionContext ctx) { copyFrom(ctx); }
	}
	public static class Initializer_expresionContext extends Postfix_expresionContext {
		public Postfix_expresionContext postfix_expresion() {
			return getRuleContext(Postfix_expresionContext.class,0);
		}
		public TerminalNode DOT() { return getToken(TswiftParser.DOT, 0); }
		public Initializer_expresionContext(Postfix_expresionContext ctx) { copyFrom(ctx); }
	}
	public static class Postfix_operationContext extends Postfix_expresionContext {
		public Postfix_expresionContext postfix_expresion() {
			return getRuleContext(Postfix_expresionContext.class,0);
		}
		public Postfix_operatorContext postfix_operator() {
			return getRuleContext(Postfix_operatorContext.class,0);
		}
		public Postfix_operationContext(Postfix_expresionContext ctx) { copyFrom(ctx); }
	}
	public static class PrimaryContext extends Postfix_expresionContext {
		public Primary_expresionContext primary_expresion() {
			return getRuleContext(Primary_expresionContext.class,0);
		}
		public PrimaryContext(Postfix_expresionContext ctx) { copyFrom(ctx); }
	}
	public static class Subscript_expresionContext extends Postfix_expresionContext {
		public Postfix_expresionContext postfix_expresion() {
			return getRuleContext(Postfix_expresionContext.class,0);
		}
		public TerminalNode LBRACK() { return getToken(TswiftParser.LBRACK, 0); }
		public Expresion_listContext expresion_list() {
			return getRuleContext(Expresion_listContext.class,0);
		}
		public TerminalNode RBRACK() { return getToken(TswiftParser.RBRACK, 0); }
		public Subscript_expresionContext(Postfix_expresionContext ctx) { copyFrom(ctx); }
	}

	public final Postfix_expresionContext postfix_expresion() throws RecognitionException {
		return postfix_expresion(0);
	}

	private Postfix_expresionContext postfix_expresion(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Postfix_expresionContext _localctx = new Postfix_expresionContext(_ctx, _parentState);
		Postfix_expresionContext _prevctx = _localctx;
		int _startState = 370;
		enterRecursionRule(_localctx, 370, RULE_postfix_expresion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1945);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,215,_ctx) ) {
			case 1:
				{
				_localctx = new PrimaryContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(1943);
				primary_expresion();
				}
				break;
			case 2:
				{
				_localctx = new Dynamic_typeContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1944);
				dynamic_type_expresion();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(1997);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,219,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(1995);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,218,_ctx) ) {
					case 1:
						{
						_localctx = new Postfix_operationContext(new Postfix_expresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_postfix_expresion);
						setState(1947);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(1948);
						postfix_operator();
						}
						break;
					case 2:
						{
						_localctx = new Function_call_expresionContext(new Postfix_expresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_postfix_expresion);
						setState(1949);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(1950);
						function_call_argument_clause();
						}
						break;
					case 3:
						{
						_localctx = new Function_call_expresion_with_closureContext(new Postfix_expresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_postfix_expresion);
						setState(1951);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(1953);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if (_la==LPAREN) {
							{
							setState(1952);
							function_call_argument_clause();
							}
						}

						setState(1955);
						trailing_closure();
						}
						break;
					case 4:
						{
						_localctx = new Initializer_expresionContext(new Postfix_expresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_postfix_expresion);
						setState(1956);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(1957);
						match(DOT);
						setState(1958);
						match(T__30);
						}
						break;
					case 5:
						{
						_localctx = new Initializer_expresion_with_argsContext(new Postfix_expresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_postfix_expresion);
						setState(1959);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(1960);
						match(DOT);
						setState(1961);
						match(T__30);
						setState(1962);
						match(LPAREN);
						setState(1963);
						argument_names();
						setState(1964);
						match(RPAREN);
						}
						break;
					case 6:
						{
						_localctx = new Explicit_member_expresion1Context(new Postfix_expresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_postfix_expresion);
						setState(1966);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(1967);
						match(DOT);
						setState(1968);
						match(Pure_decimal_digits);
						}
						break;
					case 7:
						{
						_localctx = new Explicit_member_expresion2Context(new Postfix_expresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_postfix_expresion);
						setState(1969);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(1970);
						match(DOT);
						setState(1971);
						declaracion_identifier();
						setState(1973);
						_errHandler.sync(this);
						switch ( getInterpreter().adaptivePredict(_input,217,_ctx) ) {
						case 1:
							{
							setState(1972);
							generic_argument_clause();
							}
							break;
						}
						}
						break;
					case 8:
						{
						_localctx = new Explicit_member_expresion3Context(new Postfix_expresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_postfix_expresion);
						setState(1975);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(1976);
						match(DOT);
						setState(1977);
						declaracion_identifier();
						setState(1978);
						match(LPAREN);
						setState(1979);
						argument_names();
						setState(1980);
						match(RPAREN);
						}
						break;
					case 9:
						{
						_localctx = new Explicit_member_expresion4Context(new Postfix_expresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_postfix_expresion);
						setState(1982);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(1983);
						match(LPAREN);
						setState(1984);
						argument_names();
						setState(1985);
						match(RPAREN);
						}
						break;
					case 10:
						{
						_localctx = new Postfix_self_expresionContext(new Postfix_expresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_postfix_expresion);
						setState(1987);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(1988);
						match(DOT);
						setState(1989);
						match(T__82);
						}
						break;
					case 11:
						{
						_localctx = new Subscript_expresionContext(new Postfix_expresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_postfix_expresion);
						setState(1990);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(1991);
						match(LBRACK);
						setState(1992);
						expresion_list();
						setState(1993);
						match(RBRACK);
						}
						break;
					}
					} 
				}
				setState(1999);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,219,_ctx);
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

	public static class Function_call_argument_clauseContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(TswiftParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(TswiftParser.RPAREN, 0); }
		public Function_call_argument_listContext function_call_argument_list() {
			return getRuleContext(Function_call_argument_listContext.class,0);
		}
		public Function_call_argument_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_call_argument_clause; }
	}

	public final Function_call_argument_clauseContext function_call_argument_clause() throws RecognitionException {
		Function_call_argument_clauseContext _localctx = new Function_call_argument_clauseContext(_ctx, getState());
		enterRule(_localctx, 372, RULE_function_call_argument_clause);
		try {
			setState(2006);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,220,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(2000);
				match(LPAREN);
				setState(2001);
				match(RPAREN);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(2002);
				match(LPAREN);
				setState(2003);
				function_call_argument_list();
				setState(2004);
				match(RPAREN);
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

	public static class Function_call_argument_listContext extends ParserRuleContext {
		public List<Function_call_argumentContext> function_call_argument() {
			return getRuleContexts(Function_call_argumentContext.class);
		}
		public Function_call_argumentContext function_call_argument(int i) {
			return getRuleContext(Function_call_argumentContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(TswiftParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(TswiftParser.COMMA, i);
		}
		public Function_call_argument_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_call_argument_list; }
	}

	public final Function_call_argument_listContext function_call_argument_list() throws RecognitionException {
		Function_call_argument_listContext _localctx = new Function_call_argument_listContext(_ctx, getState());
		enterRule(_localctx, 374, RULE_function_call_argument_list);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2008);
			function_call_argument();
			setState(2013);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(2009);
				match(COMMA);
				setState(2010);
				function_call_argument();
				}
				}
				setState(2015);
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

	public static class Function_call_argumentContext extends ParserRuleContext {
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Label_identifierContext label_identifier() {
			return getRuleContext(Label_identifierContext.class,0);
		}
		public TerminalNode COLON() { return getToken(TswiftParser.COLON, 0); }
		public Operator_Context operator_() {
			return getRuleContext(Operator_Context.class,0);
		}
		public Function_call_argumentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_call_argument; }
	}

	public final Function_call_argumentContext function_call_argument() throws RecognitionException {
		Function_call_argumentContext _localctx = new Function_call_argumentContext(_ctx, getState());
		enterRule(_localctx, 376, RULE_function_call_argument);
		try {
			setState(2026);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,222,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(2016);
				expresion();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(2017);
				label_identifier();
				setState(2018);
				match(COLON);
				setState(2019);
				expresion();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(2021);
				operator_();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(2022);
				label_identifier();
				setState(2023);
				match(COLON);
				setState(2024);
				operator_();
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

	public static class Trailing_closureContext extends ParserRuleContext {
		public Closure_expresionContext closure_expresion() {
			return getRuleContext(Closure_expresionContext.class,0);
		}
		public Trailing_closureContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_trailing_closure; }
	}

	public final Trailing_closureContext trailing_closure() throws RecognitionException {
		Trailing_closureContext _localctx = new Trailing_closureContext(_ctx, getState());
		enterRule(_localctx, 378, RULE_trailing_closure);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2028);
			closure_expresion();
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

	public static class Argument_namesContext extends ParserRuleContext {
		public List<Argument_nameContext> argument_name() {
			return getRuleContexts(Argument_nameContext.class);
		}
		public Argument_nameContext argument_name(int i) {
			return getRuleContext(Argument_nameContext.class,i);
		}
		public Argument_namesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_argument_names; }
	}

	public final Argument_namesContext argument_names() throws RecognitionException {
		Argument_namesContext _localctx = new Argument_namesContext(_ctx, getState());
		enterRule(_localctx, 380, RULE_argument_names);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2030);
			argument_name();
			setState(2034);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << T__2) | (1L << T__3) | (1L << T__4) | (1L << T__5) | (1L << T__6) | (1L << T__7) | (1L << T__8) | (1L << T__9) | (1L << T__10) | (1L << T__11) | (1L << T__12) | (1L << T__13) | (1L << T__14) | (1L << T__15) | (1L << T__16) | (1L << T__17) | (1L << T__19) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << T__43) | (1L << T__44) | (1L << T__45) | (1L << T__46) | (1L << T__47) | (1L << T__48) | (1L << T__49) | (1L << T__50) | (1L << T__51) | (1L << T__52) | (1L << T__53) | (1L << T__54) | (1L << T__55) | (1L << T__56) | (1L << T__57) | (1L << T__58) | (1L << T__59) | (1L << T__60) | (1L << T__61) | (1L << T__62))) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & ((1L << (T__63 - 64)) | (1L << (T__64 - 64)) | (1L << (T__65 - 64)) | (1L << (T__68 - 64)) | (1L << (T__75 - 64)) | (1L << (T__76 - 64)) | (1L << (T__77 - 64)) | (1L << (T__78 - 64)) | (1L << (T__80 - 64)) | (1L << (T__82 - 64)) | (1L << (T__83 - 64)) | (1L << (T__84 - 64)) | (1L << (T__91 - 64)) | (1L << (T__92 - 64)) | (1L << (T__93 - 64)) | (1L << (T__94 - 64)) | (1L << (T__95 - 64)) | (1L << (T__97 - 64)) | (1L << (T__98 - 64)) | (1L << (T__99 - 64)) | (1L << (T__100 - 64)) | (1L << (T__101 - 64)) | (1L << (T__102 - 64)) | (1L << (T__103 - 64)) | (1L << (T__104 - 64)) | (1L << (T__105 - 64)) | (1L << (T__106 - 64)) | (1L << (T__107 - 64)) | (1L << (T__108 - 64)) | (1L << (T__109 - 64)) | (1L << (T__110 - 64)) | (1L << (T__111 - 64)) | (1L << (T__112 - 64)) | (1L << (T__113 - 64)) | (1L << (T__114 - 64)) | (1L << (T__115 - 64)) | (1L << (T__116 - 64)) | (1L << (T__117 - 64)) | (1L << (T__118 - 64)) | (1L << (T__119 - 64)) | (1L << (T__120 - 64)) | (1L << (T__121 - 64)) | (1L << (T__122 - 64)) | (1L << (Identifier - 64)))) != 0)) {
				{
				{
				setState(2031);
				argument_name();
				}
				}
				setState(2036);
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

	public static class Argument_nameContext extends ParserRuleContext {
		public Label_identifierContext label_identifier() {
			return getRuleContext(Label_identifierContext.class,0);
		}
		public TerminalNode COLON() { return getToken(TswiftParser.COLON, 0); }
		public Argument_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_argument_name; }
	}

	public final Argument_nameContext argument_name() throws RecognitionException {
		Argument_nameContext _localctx = new Argument_nameContext(_ctx, getState());
		enterRule(_localctx, 382, RULE_argument_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2037);
			label_identifier();
			setState(2038);
			match(COLON);
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

	public static class Dynamic_type_expresionContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(TswiftParser.LPAREN, 0); }
		public TerminalNode COLON() { return getToken(TswiftParser.COLON, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(TswiftParser.RPAREN, 0); }
		public Dynamic_type_expresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dynamic_type_expresion; }
	}

	public final Dynamic_type_expresionContext dynamic_type_expresion() throws RecognitionException {
		Dynamic_type_expresionContext _localctx = new Dynamic_type_expresionContext(_ctx, getState());
		enterRule(_localctx, 384, RULE_dynamic_type_expresion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2040);
			match(T__91);
			setState(2041);
			match(LPAREN);
			setState(2042);
			match(T__92);
			setState(2043);
			match(COLON);
			setState(2044);
			expresion();
			setState(2045);
			match(RPAREN);
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

	public static class Type_Context extends ParserRuleContext {
		public Type_Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_; }
	 
		public Type_Context() { }
		public void copyFrom(Type_Context ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class The_metatype_protocol_typeContext extends Type_Context {
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public TerminalNode DOT() { return getToken(TswiftParser.DOT, 0); }
		public The_metatype_protocol_typeContext(Type_Context ctx) { copyFrom(ctx); }
	}
	public static class The_function_typeContext extends Type_Context {
		public Function_typeContext function_type() {
			return getRuleContext(Function_typeContext.class,0);
		}
		public The_function_typeContext(Type_Context ctx) { copyFrom(ctx); }
	}
	public static class The_implicitly_unwrapped_optional_typeContext extends Type_Context {
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public TerminalNode BANG() { return getToken(TswiftParser.BANG, 0); }
		public The_implicitly_unwrapped_optional_typeContext(Type_Context ctx) { copyFrom(ctx); }
	}
	public static class The_dictionary_typeContext extends Type_Context {
		public Dictionary_typeContext dictionary_type() {
			return getRuleContext(Dictionary_typeContext.class,0);
		}
		public The_dictionary_typeContext(Type_Context ctx) { copyFrom(ctx); }
	}
	public static class The_optional_typeContext extends Type_Context {
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public TerminalNode QUESTION() { return getToken(TswiftParser.QUESTION, 0); }
		public The_optional_typeContext(Type_Context ctx) { copyFrom(ctx); }
	}
	public static class The_tuple_typeContext extends Type_Context {
		public Tuple_typeContext tuple_type() {
			return getRuleContext(Tuple_typeContext.class,0);
		}
		public The_tuple_typeContext(Type_Context ctx) { copyFrom(ctx); }
	}
	public static class The_self_typeContext extends Type_Context {
		public The_self_typeContext(Type_Context ctx) { copyFrom(ctx); }
	}
	public static class The_array_typeContext extends Type_Context {
		public Array_typeContext array_type() {
			return getRuleContext(Array_typeContext.class,0);
		}
		public The_array_typeContext(Type_Context ctx) { copyFrom(ctx); }
	}
	public static class The_metatype_type_typeContext extends Type_Context {
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public TerminalNode DOT() { return getToken(TswiftParser.DOT, 0); }
		public The_metatype_type_typeContext(Type_Context ctx) { copyFrom(ctx); }
	}
	public static class The_protocol_composition_typeContext extends Type_Context {
		public Protocol_composition_typeContext protocol_composition_type() {
			return getRuleContext(Protocol_composition_typeContext.class,0);
		}
		public The_protocol_composition_typeContext(Type_Context ctx) { copyFrom(ctx); }
	}
	public static class The_any_typeContext extends Type_Context {
		public The_any_typeContext(Type_Context ctx) { copyFrom(ctx); }
	}
	public static class The_type_identifierContext extends Type_Context {
		public Type_identifierContext type_identifier() {
			return getRuleContext(Type_identifierContext.class,0);
		}
		public The_type_identifierContext(Type_Context ctx) { copyFrom(ctx); }
	}

	public final Type_Context type_() throws RecognitionException {
		return type_(0);
	}

	private Type_Context type_(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Type_Context _localctx = new Type_Context(_ctx, _parentState);
		Type_Context _prevctx = _localctx;
		int _startState = 386;
		enterRecursionRule(_localctx, 386, RULE_type_, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(2056);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,224,_ctx) ) {
			case 1:
				{
				_localctx = new The_array_typeContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(2048);
				array_type();
				}
				break;
			case 2:
				{
				_localctx = new The_dictionary_typeContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(2049);
				dictionary_type();
				}
				break;
			case 3:
				{
				_localctx = new The_function_typeContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(2050);
				function_type();
				}
				break;
			case 4:
				{
				_localctx = new The_type_identifierContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(2051);
				type_identifier();
				}
				break;
			case 5:
				{
				_localctx = new The_tuple_typeContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(2052);
				tuple_type();
				}
				break;
			case 6:
				{
				_localctx = new The_protocol_composition_typeContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(2053);
				protocol_composition_type();
				}
				break;
			case 7:
				{
				_localctx = new The_any_typeContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(2054);
				match(T__95);
				}
				break;
			case 8:
				{
				_localctx = new The_self_typeContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(2055);
				match(T__83);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(2070);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,226,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(2068);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,225,_ctx) ) {
					case 1:
						{
						_localctx = new The_optional_typeContext(new Type_Context(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_type_);
						setState(2058);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(2059);
						match(QUESTION);
						}
						break;
					case 2:
						{
						_localctx = new The_implicitly_unwrapped_optional_typeContext(new Type_Context(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_type_);
						setState(2060);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(2061);
						match(BANG);
						}
						break;
					case 3:
						{
						_localctx = new The_metatype_type_typeContext(new Type_Context(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_type_);
						setState(2062);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(2063);
						match(DOT);
						setState(2064);
						match(T__93);
						}
						break;
					case 4:
						{
						_localctx = new The_metatype_protocol_typeContext(new Type_Context(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_type_);
						setState(2065);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(2066);
						match(DOT);
						setState(2067);
						match(T__94);
						}
						break;
					}
					} 
				}
				setState(2072);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,226,_ctx);
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

	public static class Type_annotationContext extends ParserRuleContext {
		public TerminalNode COLON() { return getToken(TswiftParser.COLON, 0); }
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Type_annotationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_annotation; }
	}

	public final Type_annotationContext type_annotation() throws RecognitionException {
		Type_annotationContext _localctx = new Type_annotationContext(_ctx, getState());
		enterRule(_localctx, 388, RULE_type_annotation);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2073);
			match(COLON);
			setState(2075);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,227,_ctx) ) {
			case 1:
				{
				setState(2074);
				attributes();
				}
				break;
			}
			setState(2078);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__96) {
				{
				setState(2077);
				match(T__96);
				}
			}

			setState(2080);
			type_(0);
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

	public static class Type_identifierContext extends ParserRuleContext {
		public Type_nameContext type_name() {
			return getRuleContext(Type_nameContext.class,0);
		}
		public Generic_argument_clauseContext generic_argument_clause() {
			return getRuleContext(Generic_argument_clauseContext.class,0);
		}
		public TerminalNode DOT() { return getToken(TswiftParser.DOT, 0); }
		public Type_identifierContext type_identifier() {
			return getRuleContext(Type_identifierContext.class,0);
		}
		public Type_identifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_identifier; }
	}

	public final Type_identifierContext type_identifier() throws RecognitionException {
		Type_identifierContext _localctx = new Type_identifierContext(_ctx, getState());
		enterRule(_localctx, 390, RULE_type_identifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2082);
			type_name();
			setState(2084);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,229,_ctx) ) {
			case 1:
				{
				setState(2083);
				generic_argument_clause();
				}
				break;
			}
			setState(2088);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,230,_ctx) ) {
			case 1:
				{
				setState(2086);
				match(DOT);
				setState(2087);
				type_identifier();
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

	public static class Type_nameContext extends ParserRuleContext {
		public Declaracion_identifierContext declaracion_identifier() {
			return getRuleContext(Declaracion_identifierContext.class,0);
		}
		public Type_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_name; }
	}

	public final Type_nameContext type_name() throws RecognitionException {
		Type_nameContext _localctx = new Type_nameContext(_ctx, getState());
		enterRule(_localctx, 392, RULE_type_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2090);
			declaracion_identifier();
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

	public static class Tuple_typeContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(TswiftParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(TswiftParser.RPAREN, 0); }
		public Tuple_type_element_listContext tuple_type_element_list() {
			return getRuleContext(Tuple_type_element_listContext.class,0);
		}
		public Tuple_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuple_type; }
	}

	public final Tuple_typeContext tuple_type() throws RecognitionException {
		Tuple_typeContext _localctx = new Tuple_typeContext(_ctx, getState());
		enterRule(_localctx, 394, RULE_tuple_type);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2092);
			match(LPAREN);
			setState(2094);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << T__2) | (1L << T__3) | (1L << T__4) | (1L << T__5) | (1L << T__6) | (1L << T__7) | (1L << T__8) | (1L << T__9) | (1L << T__10) | (1L << T__11) | (1L << T__12) | (1L << T__13) | (1L << T__14) | (1L << T__15) | (1L << T__16) | (1L << T__17) | (1L << T__19) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << T__43) | (1L << T__44) | (1L << T__45) | (1L << T__46) | (1L << T__47) | (1L << T__48) | (1L << T__49) | (1L << T__50) | (1L << T__51) | (1L << T__52) | (1L << T__53) | (1L << T__54) | (1L << T__55) | (1L << T__56) | (1L << T__57) | (1L << T__58) | (1L << T__59) | (1L << T__60) | (1L << T__61) | (1L << T__62))) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & ((1L << (T__63 - 64)) | (1L << (T__64 - 64)) | (1L << (T__65 - 64)) | (1L << (T__68 - 64)) | (1L << (T__75 - 64)) | (1L << (T__76 - 64)) | (1L << (T__77 - 64)) | (1L << (T__78 - 64)) | (1L << (T__80 - 64)) | (1L << (T__82 - 64)) | (1L << (T__83 - 64)) | (1L << (T__84 - 64)) | (1L << (T__91 - 64)) | (1L << (T__92 - 64)) | (1L << (T__93 - 64)) | (1L << (T__94 - 64)) | (1L << (T__95 - 64)) | (1L << (T__97 - 64)) | (1L << (T__98 - 64)) | (1L << (T__99 - 64)) | (1L << (T__100 - 64)) | (1L << (T__101 - 64)) | (1L << (T__102 - 64)) | (1L << (T__103 - 64)) | (1L << (T__104 - 64)) | (1L << (T__105 - 64)) | (1L << (T__106 - 64)) | (1L << (T__107 - 64)) | (1L << (T__108 - 64)) | (1L << (T__109 - 64)) | (1L << (T__110 - 64)) | (1L << (T__111 - 64)) | (1L << (T__112 - 64)) | (1L << (T__113 - 64)) | (1L << (T__114 - 64)) | (1L << (T__115 - 64)) | (1L << (T__116 - 64)) | (1L << (T__117 - 64)) | (1L << (T__118 - 64)) | (1L << (T__119 - 64)) | (1L << (T__120 - 64)) | (1L << (T__121 - 64)) | (1L << (T__122 - 64)) | (1L << (Identifier - 64)))) != 0) || ((((_la - 129)) & ~0x3f) == 0 && ((1L << (_la - 129)) & ((1L << (LPAREN - 129)) | (1L << (LBRACK - 129)) | (1L << (AT - 129)))) != 0)) {
				{
				setState(2093);
				tuple_type_element_list();
				}
			}

			setState(2096);
			match(RPAREN);
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

	public static class Tuple_type_element_listContext extends ParserRuleContext {
		public Tuple_type_elementContext tuple_type_element() {
			return getRuleContext(Tuple_type_elementContext.class,0);
		}
		public TerminalNode COMMA() { return getToken(TswiftParser.COMMA, 0); }
		public Tuple_type_element_listContext tuple_type_element_list() {
			return getRuleContext(Tuple_type_element_listContext.class,0);
		}
		public Tuple_type_element_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuple_type_element_list; }
	}

	public final Tuple_type_element_listContext tuple_type_element_list() throws RecognitionException {
		Tuple_type_element_listContext _localctx = new Tuple_type_element_listContext(_ctx, getState());
		enterRule(_localctx, 396, RULE_tuple_type_element_list);
		try {
			setState(2103);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,232,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(2098);
				tuple_type_element();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(2099);
				tuple_type_element();
				setState(2100);
				match(COMMA);
				setState(2101);
				tuple_type_element_list();
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

	public static class Tuple_type_elementContext extends ParserRuleContext {
		public Element_nameContext element_name() {
			return getRuleContext(Element_nameContext.class,0);
		}
		public Type_annotationContext type_annotation() {
			return getRuleContext(Type_annotationContext.class,0);
		}
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public Tuple_type_elementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuple_type_element; }
	}

	public final Tuple_type_elementContext tuple_type_element() throws RecognitionException {
		Tuple_type_elementContext _localctx = new Tuple_type_elementContext(_ctx, getState());
		enterRule(_localctx, 398, RULE_tuple_type_element);
		try {
			setState(2109);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,233,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(2105);
				element_name();
				setState(2106);
				type_annotation();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(2108);
				type_(0);
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

	public static class Element_nameContext extends ParserRuleContext {
		public Label_identifierContext label_identifier() {
			return getRuleContext(Label_identifierContext.class,0);
		}
		public Element_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_element_name; }
	}

	public final Element_nameContext element_name() throws RecognitionException {
		Element_nameContext _localctx = new Element_nameContext(_ctx, getState());
		enterRule(_localctx, 400, RULE_element_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2111);
			label_identifier();
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

	public static class Function_typeContext extends ParserRuleContext {
		public Function_type_argument_clauseContext function_type_argument_clause() {
			return getRuleContext(Function_type_argument_clauseContext.class,0);
		}
		public Arrow_operatorContext arrow_operator() {
			return getRuleContext(Arrow_operatorContext.class,0);
		}
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Function_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_type; }
	}

	public final Function_typeContext function_type() throws RecognitionException {
		Function_typeContext _localctx = new Function_typeContext(_ctx, getState());
		enterRule(_localctx, 402, RULE_function_type);
		int _la;
		try {
			setState(2131);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,237,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(2114);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AT) {
					{
					setState(2113);
					attributes();
					}
				}

				setState(2116);
				function_type_argument_clause();
				setState(2118);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,235,_ctx) ) {
				case 1:
					{
					setState(2117);
					match(T__25);
					}
					break;
				}
				setState(2120);
				arrow_operator();
				setState(2121);
				type_(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(2124);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AT) {
					{
					setState(2123);
					attributes();
					}
				}

				setState(2126);
				function_type_argument_clause();
				setState(2127);
				match(T__26);
				setState(2128);
				arrow_operator();
				setState(2129);
				type_(0);
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

	public static class Function_type_argument_clauseContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(TswiftParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(TswiftParser.RPAREN, 0); }
		public Function_type_argument_listContext function_type_argument_list() {
			return getRuleContext(Function_type_argument_listContext.class,0);
		}
		public Range_operatorContext range_operator() {
			return getRuleContext(Range_operatorContext.class,0);
		}
		public Function_type_argument_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_type_argument_clause; }
	}

	public final Function_type_argument_clauseContext function_type_argument_clause() throws RecognitionException {
		Function_type_argument_clauseContext _localctx = new Function_type_argument_clauseContext(_ctx, getState());
		enterRule(_localctx, 404, RULE_function_type_argument_clause);
		try {
			setState(2142);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,239,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(2133);
				match(LPAREN);
				setState(2134);
				match(RPAREN);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(2135);
				match(LPAREN);
				setState(2136);
				function_type_argument_list();
				setState(2138);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,238,_ctx) ) {
				case 1:
					{
					setState(2137);
					range_operator();
					}
					break;
				}
				setState(2140);
				match(RPAREN);
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

	public static class Function_type_argument_listContext extends ParserRuleContext {
		public Function_type_argumentContext function_type_argument() {
			return getRuleContext(Function_type_argumentContext.class,0);
		}
		public TerminalNode COMMA() { return getToken(TswiftParser.COMMA, 0); }
		public Function_type_argument_listContext function_type_argument_list() {
			return getRuleContext(Function_type_argument_listContext.class,0);
		}
		public Function_type_argument_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_type_argument_list; }
	}

	public final Function_type_argument_listContext function_type_argument_list() throws RecognitionException {
		Function_type_argument_listContext _localctx = new Function_type_argument_listContext(_ctx, getState());
		enterRule(_localctx, 406, RULE_function_type_argument_list);
		try {
			setState(2149);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,240,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(2144);
				function_type_argument();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(2145);
				function_type_argument();
				setState(2146);
				match(COMMA);
				setState(2147);
				function_type_argument_list();
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

	public static class Function_type_argumentContext extends ParserRuleContext {
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Argument_labelContext argument_label() {
			return getRuleContext(Argument_labelContext.class,0);
		}
		public Type_annotationContext type_annotation() {
			return getRuleContext(Type_annotationContext.class,0);
		}
		public Function_type_argumentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_type_argument; }
	}

	public final Function_type_argumentContext function_type_argument() throws RecognitionException {
		Function_type_argumentContext _localctx = new Function_type_argumentContext(_ctx, getState());
		enterRule(_localctx, 408, RULE_function_type_argument);
		int _la;
		try {
			setState(2161);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,243,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(2152);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,241,_ctx) ) {
				case 1:
					{
					setState(2151);
					attributes();
					}
					break;
				}
				setState(2155);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__96) {
					{
					setState(2154);
					match(T__96);
					}
				}

				setState(2157);
				type_(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(2158);
				argument_label();
				setState(2159);
				type_annotation();
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

	public static class Argument_labelContext extends ParserRuleContext {
		public Label_identifierContext label_identifier() {
			return getRuleContext(Label_identifierContext.class,0);
		}
		public Argument_labelContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_argument_label; }
	}

	public final Argument_labelContext argument_label() throws RecognitionException {
		Argument_labelContext _localctx = new Argument_labelContext(_ctx, getState());
		enterRule(_localctx, 410, RULE_argument_label);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2163);
			label_identifier();
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

	public static class Array_typeContext extends ParserRuleContext {
		public TerminalNode LBRACK() { return getToken(TswiftParser.LBRACK, 0); }
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public TerminalNode RBRACK() { return getToken(TswiftParser.RBRACK, 0); }
		public Array_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_array_type; }
	}

	public final Array_typeContext array_type() throws RecognitionException {
		Array_typeContext _localctx = new Array_typeContext(_ctx, getState());
		enterRule(_localctx, 412, RULE_array_type);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2165);
			match(LBRACK);
			setState(2166);
			type_(0);
			setState(2167);
			match(RBRACK);
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

	public static class Dictionary_typeContext extends ParserRuleContext {
		public TerminalNode LBRACK() { return getToken(TswiftParser.LBRACK, 0); }
		public List<Type_Context> type_() {
			return getRuleContexts(Type_Context.class);
		}
		public Type_Context type_(int i) {
			return getRuleContext(Type_Context.class,i);
		}
		public TerminalNode COLON() { return getToken(TswiftParser.COLON, 0); }
		public TerminalNode RBRACK() { return getToken(TswiftParser.RBRACK, 0); }
		public Dictionary_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dictionary_type; }
	}

	public final Dictionary_typeContext dictionary_type() throws RecognitionException {
		Dictionary_typeContext _localctx = new Dictionary_typeContext(_ctx, getState());
		enterRule(_localctx, 414, RULE_dictionary_type);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2169);
			match(LBRACK);
			setState(2170);
			type_(0);
			setState(2171);
			match(COLON);
			setState(2172);
			type_(0);
			setState(2173);
			match(RBRACK);
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

	public static class Protocol_composition_typeContext extends ParserRuleContext {
		public List<Protocol_identifierContext> protocol_identifier() {
			return getRuleContexts(Protocol_identifierContext.class);
		}
		public Protocol_identifierContext protocol_identifier(int i) {
			return getRuleContext(Protocol_identifierContext.class,i);
		}
		public List<TerminalNode> AND() { return getTokens(TswiftParser.AND); }
		public TerminalNode AND(int i) {
			return getToken(TswiftParser.AND, i);
		}
		public Protocol_composition_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_protocol_composition_type; }
	}

	public final Protocol_composition_typeContext protocol_composition_type() throws RecognitionException {
		Protocol_composition_typeContext _localctx = new Protocol_composition_typeContext(_ctx, getState());
		enterRule(_localctx, 416, RULE_protocol_composition_type);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(2175);
			protocol_identifier();
			setState(2178); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(2176);
					match(AND);
					setState(2177);
					protocol_identifier();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(2180); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,244,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
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

	public static class Protocol_identifierContext extends ParserRuleContext {
		public Type_identifierContext type_identifier() {
			return getRuleContext(Type_identifierContext.class,0);
		}
		public Protocol_identifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_protocol_identifier; }
	}

	public final Protocol_identifierContext protocol_identifier() throws RecognitionException {
		Protocol_identifierContext _localctx = new Protocol_identifierContext(_ctx, getState());
		enterRule(_localctx, 418, RULE_protocol_identifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2182);
			type_identifier();
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

	public static class Type_inheritance_clauseContext extends ParserRuleContext {
		public TerminalNode COLON() { return getToken(TswiftParser.COLON, 0); }
		public Class_requirementContext class_requirement() {
			return getRuleContext(Class_requirementContext.class,0);
		}
		public TerminalNode COMMA() { return getToken(TswiftParser.COMMA, 0); }
		public Type_inheritance_listContext type_inheritance_list() {
			return getRuleContext(Type_inheritance_listContext.class,0);
		}
		public Type_inheritance_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_inheritance_clause; }
	}

	public final Type_inheritance_clauseContext type_inheritance_clause() throws RecognitionException {
		Type_inheritance_clauseContext _localctx = new Type_inheritance_clauseContext(_ctx, getState());
		enterRule(_localctx, 420, RULE_type_inheritance_clause);
		try {
			setState(2193);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,245,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(2184);
				match(COLON);
				setState(2185);
				class_requirement();
				setState(2186);
				match(COMMA);
				setState(2187);
				type_inheritance_list();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(2189);
				match(COLON);
				setState(2190);
				class_requirement();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(2191);
				match(COLON);
				setState(2192);
				type_inheritance_list();
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

	public static class Type_inheritance_listContext extends ParserRuleContext {
		public Type_identifierContext type_identifier() {
			return getRuleContext(Type_identifierContext.class,0);
		}
		public TerminalNode COMMA() { return getToken(TswiftParser.COMMA, 0); }
		public Type_inheritance_listContext type_inheritance_list() {
			return getRuleContext(Type_inheritance_listContext.class,0);
		}
		public Type_inheritance_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_inheritance_list; }
	}

	public final Type_inheritance_listContext type_inheritance_list() throws RecognitionException {
		Type_inheritance_listContext _localctx = new Type_inheritance_listContext(_ctx, getState());
		enterRule(_localctx, 422, RULE_type_inheritance_list);
		try {
			setState(2200);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,246,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(2195);
				type_identifier();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(2196);
				type_identifier();
				setState(2197);
				match(COMMA);
				setState(2198);
				type_inheritance_list();
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

	public static class Class_requirementContext extends ParserRuleContext {
		public Class_requirementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_class_requirement; }
	}

	public final Class_requirementContext class_requirement() throws RecognitionException {
		Class_requirementContext _localctx = new Class_requirementContext(_ctx, getState());
		enterRule(_localctx, 424, RULE_class_requirement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2202);
			match(T__15);
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

	public static class Declaracion_identifierContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(TswiftParser.Identifier, 0); }
		public Keyword_as_identifier_in_declaracionsContext keyword_as_identifier_in_declaracions() {
			return getRuleContext(Keyword_as_identifier_in_declaracionsContext.class,0);
		}
		public Declaracion_identifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion_identifier; }
	}

	public final Declaracion_identifierContext declaracion_identifier() throws RecognitionException {
		Declaracion_identifierContext _localctx = new Declaracion_identifierContext(_ctx, getState());
		enterRule(_localctx, 426, RULE_declaracion_identifier);
		try {
			setState(2206);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Identifier:
				enterOuterAlt(_localctx, 1);
				{
				setState(2204);
				match(Identifier);
				}
				break;
			case T__21:
			case T__22:
			case T__23:
			case T__24:
			case T__27:
			case T__28:
			case T__34:
			case T__36:
			case T__37:
			case T__39:
			case T__40:
			case T__41:
			case T__42:
			case T__43:
			case T__44:
			case T__45:
			case T__46:
			case T__47:
			case T__48:
			case T__49:
			case T__50:
			case T__51:
			case T__53:
			case T__54:
			case T__55:
			case T__56:
			case T__61:
			case T__62:
			case T__63:
			case T__75:
			case T__76:
			case T__77:
			case T__78:
			case T__80:
			case T__91:
			case T__92:
			case T__93:
			case T__94:
			case T__97:
			case T__98:
			case T__99:
			case T__100:
			case T__101:
			case T__102:
			case T__103:
			case T__104:
			case T__105:
			case T__106:
			case T__107:
			case T__108:
			case T__109:
			case T__110:
			case T__111:
			case T__112:
				enterOuterAlt(_localctx, 2);
				{
				setState(2205);
				keyword_as_identifier_in_declaracions();
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

	public static class Label_identifierContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(TswiftParser.Identifier, 0); }
		public Keyword_as_identifier_in_labelsContext keyword_as_identifier_in_labels() {
			return getRuleContext(Keyword_as_identifier_in_labelsContext.class,0);
		}
		public Label_identifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_label_identifier; }
	}

	public final Label_identifierContext label_identifier() throws RecognitionException {
		Label_identifierContext _localctx = new Label_identifierContext(_ctx, getState());
		enterRule(_localctx, 428, RULE_label_identifier);
		try {
			setState(2210);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Identifier:
				enterOuterAlt(_localctx, 1);
				{
				setState(2208);
				match(Identifier);
				}
				break;
			case T__0:
			case T__1:
			case T__2:
			case T__3:
			case T__4:
			case T__5:
			case T__6:
			case T__7:
			case T__8:
			case T__9:
			case T__10:
			case T__11:
			case T__12:
			case T__13:
			case T__14:
			case T__15:
			case T__16:
			case T__17:
			case T__19:
			case T__21:
			case T__22:
			case T__23:
			case T__24:
			case T__25:
			case T__26:
			case T__27:
			case T__28:
			case T__29:
			case T__30:
			case T__31:
			case T__32:
			case T__33:
			case T__34:
			case T__35:
			case T__36:
			case T__37:
			case T__38:
			case T__39:
			case T__40:
			case T__41:
			case T__42:
			case T__43:
			case T__44:
			case T__45:
			case T__46:
			case T__47:
			case T__48:
			case T__49:
			case T__50:
			case T__51:
			case T__52:
			case T__53:
			case T__54:
			case T__55:
			case T__56:
			case T__57:
			case T__58:
			case T__59:
			case T__60:
			case T__61:
			case T__62:
			case T__63:
			case T__64:
			case T__65:
			case T__68:
			case T__75:
			case T__76:
			case T__77:
			case T__78:
			case T__80:
			case T__82:
			case T__83:
			case T__84:
			case T__91:
			case T__92:
			case T__93:
			case T__94:
			case T__95:
			case T__97:
			case T__98:
			case T__99:
			case T__100:
			case T__101:
			case T__102:
			case T__103:
			case T__104:
			case T__105:
			case T__106:
			case T__107:
			case T__108:
			case T__109:
			case T__110:
			case T__111:
			case T__112:
			case T__113:
			case T__114:
			case T__115:
			case T__116:
			case T__117:
			case T__118:
			case T__119:
			case T__120:
			case T__121:
			case T__122:
				enterOuterAlt(_localctx, 2);
				{
				setState(2209);
				keyword_as_identifier_in_labels();
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

	public static class Keyword_as_identifier_in_declaracionsContext extends ParserRuleContext {
		public Keyword_as_identifier_in_declaracionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_keyword_as_identifier_in_declaracions; }
	}

	public final Keyword_as_identifier_in_declaracionsContext keyword_as_identifier_in_declaracions() throws RecognitionException {
		Keyword_as_identifier_in_declaracionsContext _localctx = new Keyword_as_identifier_in_declaracionsContext(_ctx, getState());
		enterRule(_localctx, 430, RULE_keyword_as_identifier_in_declaracions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2212);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__27) | (1L << T__28) | (1L << T__34) | (1L << T__36) | (1L << T__37) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << T__43) | (1L << T__44) | (1L << T__45) | (1L << T__46) | (1L << T__47) | (1L << T__48) | (1L << T__49) | (1L << T__50) | (1L << T__51) | (1L << T__53) | (1L << T__54) | (1L << T__55) | (1L << T__56) | (1L << T__61) | (1L << T__62))) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & ((1L << (T__63 - 64)) | (1L << (T__75 - 64)) | (1L << (T__76 - 64)) | (1L << (T__77 - 64)) | (1L << (T__78 - 64)) | (1L << (T__80 - 64)) | (1L << (T__91 - 64)) | (1L << (T__92 - 64)) | (1L << (T__93 - 64)) | (1L << (T__94 - 64)) | (1L << (T__97 - 64)) | (1L << (T__98 - 64)) | (1L << (T__99 - 64)) | (1L << (T__100 - 64)) | (1L << (T__101 - 64)) | (1L << (T__102 - 64)) | (1L << (T__103 - 64)) | (1L << (T__104 - 64)) | (1L << (T__105 - 64)) | (1L << (T__106 - 64)) | (1L << (T__107 - 64)) | (1L << (T__108 - 64)) | (1L << (T__109 - 64)) | (1L << (T__110 - 64)) | (1L << (T__111 - 64)) | (1L << (T__112 - 64)))) != 0)) ) {
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

	public static class Keyword_as_identifier_in_labelsContext extends ParserRuleContext {
		public Keyword_as_identifier_in_labelsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_keyword_as_identifier_in_labels; }
	}

	public final Keyword_as_identifier_in_labelsContext keyword_as_identifier_in_labels() throws RecognitionException {
		Keyword_as_identifier_in_labelsContext _localctx = new Keyword_as_identifier_in_labelsContext(_ctx, getState());
		enterRule(_localctx, 432, RULE_keyword_as_identifier_in_labels);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2214);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << T__2) | (1L << T__3) | (1L << T__4) | (1L << T__5) | (1L << T__6) | (1L << T__7) | (1L << T__8) | (1L << T__9) | (1L << T__10) | (1L << T__11) | (1L << T__12) | (1L << T__13) | (1L << T__14) | (1L << T__15) | (1L << T__16) | (1L << T__17) | (1L << T__19) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << T__43) | (1L << T__44) | (1L << T__45) | (1L << T__46) | (1L << T__47) | (1L << T__48) | (1L << T__49) | (1L << T__50) | (1L << T__51) | (1L << T__52) | (1L << T__53) | (1L << T__54) | (1L << T__55) | (1L << T__56) | (1L << T__57) | (1L << T__58) | (1L << T__59) | (1L << T__60) | (1L << T__61) | (1L << T__62))) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & ((1L << (T__63 - 64)) | (1L << (T__64 - 64)) | (1L << (T__65 - 64)) | (1L << (T__68 - 64)) | (1L << (T__75 - 64)) | (1L << (T__76 - 64)) | (1L << (T__77 - 64)) | (1L << (T__78 - 64)) | (1L << (T__80 - 64)) | (1L << (T__82 - 64)) | (1L << (T__83 - 64)) | (1L << (T__84 - 64)) | (1L << (T__91 - 64)) | (1L << (T__92 - 64)) | (1L << (T__93 - 64)) | (1L << (T__94 - 64)) | (1L << (T__95 - 64)) | (1L << (T__97 - 64)) | (1L << (T__98 - 64)) | (1L << (T__99 - 64)) | (1L << (T__100 - 64)) | (1L << (T__101 - 64)) | (1L << (T__102 - 64)) | (1L << (T__103 - 64)) | (1L << (T__104 - 64)) | (1L << (T__105 - 64)) | (1L << (T__106 - 64)) | (1L << (T__107 - 64)) | (1L << (T__108 - 64)) | (1L << (T__109 - 64)) | (1L << (T__110 - 64)) | (1L << (T__111 - 64)) | (1L << (T__112 - 64)) | (1L << (T__113 - 64)) | (1L << (T__114 - 64)) | (1L << (T__115 - 64)) | (1L << (T__116 - 64)) | (1L << (T__117 - 64)) | (1L << (T__118 - 64)) | (1L << (T__119 - 64)) | (1L << (T__120 - 64)) | (1L << (T__121 - 64)) | (1L << (T__122 - 64)))) != 0)) ) {
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

	public static class Assignment_operatorContext extends ParserRuleContext {
		public TerminalNode EQUAL() { return getToken(TswiftParser.EQUAL, 0); }
		public Assignment_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignment_operator; }
	}

	public final Assignment_operatorContext assignment_operator() throws RecognitionException {
		Assignment_operatorContext _localctx = new Assignment_operatorContext(_ctx, getState());
		enterRule(_localctx, 434, RULE_assignment_operator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2216);
			if (!(SwiftSupport.isBinaryOp(_input))) throw new FailedPredicateException(this, "SwiftSupport.isBinaryOp(_input)");
			setState(2217);
			match(EQUAL);
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

	public static class Negate_prefix_operatorContext extends ParserRuleContext {
		public TerminalNode SUB() { return getToken(TswiftParser.SUB, 0); }
		public Negate_prefix_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_negate_prefix_operator; }
	}

	public final Negate_prefix_operatorContext negate_prefix_operator() throws RecognitionException {
		Negate_prefix_operatorContext _localctx = new Negate_prefix_operatorContext(_ctx, getState());
		enterRule(_localctx, 436, RULE_negate_prefix_operator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2219);
			if (!(SwiftSupport.isPrefixOp(_input))) throw new FailedPredicateException(this, "SwiftSupport.isPrefixOp(_input)");
			setState(2220);
			match(SUB);
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

	public static class Compilation_condition_ANDContext extends ParserRuleContext {
		public List<TerminalNode> AND() { return getTokens(TswiftParser.AND); }
		public TerminalNode AND(int i) {
			return getToken(TswiftParser.AND, i);
		}
		public Compilation_condition_ANDContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_compilation_condition_AND; }
	}

	public final Compilation_condition_ANDContext compilation_condition_AND() throws RecognitionException {
		Compilation_condition_ANDContext _localctx = new Compilation_condition_ANDContext(_ctx, getState());
		enterRule(_localctx, 438, RULE_compilation_condition_AND);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2222);
			if (!(SwiftSupport.isOperator(_input,"&&"))) throw new FailedPredicateException(this, "SwiftSupport.isOperator(_input,\"&&\")");
			setState(2223);
			match(AND);
			setState(2224);
			match(AND);
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

	public static class Compilation_condition_ORContext extends ParserRuleContext {
		public List<TerminalNode> OR() { return getTokens(TswiftParser.OR); }
		public TerminalNode OR(int i) {
			return getToken(TswiftParser.OR, i);
		}
		public Compilation_condition_ORContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_compilation_condition_OR; }
	}

	public final Compilation_condition_ORContext compilation_condition_OR() throws RecognitionException {
		Compilation_condition_ORContext _localctx = new Compilation_condition_ORContext(_ctx, getState());
		enterRule(_localctx, 440, RULE_compilation_condition_OR);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2226);
			if (!(SwiftSupport.isOperator(_input,"||"))) throw new FailedPredicateException(this, "SwiftSupport.isOperator(_input,\"||\")");
			setState(2227);
			match(OR);
			setState(2228);
			match(OR);
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

	public static class Compilation_condition_GEContext extends ParserRuleContext {
		public TerminalNode GT() { return getToken(TswiftParser.GT, 0); }
		public TerminalNode EQUAL() { return getToken(TswiftParser.EQUAL, 0); }
		public Compilation_condition_GEContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_compilation_condition_GE; }
	}

	public final Compilation_condition_GEContext compilation_condition_GE() throws RecognitionException {
		Compilation_condition_GEContext _localctx = new Compilation_condition_GEContext(_ctx, getState());
		enterRule(_localctx, 442, RULE_compilation_condition_GE);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2230);
			if (!(SwiftSupport.isOperator(_input,">="))) throw new FailedPredicateException(this, "SwiftSupport.isOperator(_input,\">=\")");
			setState(2231);
			match(GT);
			setState(2232);
			match(EQUAL);
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

	public static class Arrow_operatorContext extends ParserRuleContext {
		public TerminalNode SUB() { return getToken(TswiftParser.SUB, 0); }
		public TerminalNode GT() { return getToken(TswiftParser.GT, 0); }
		public Arrow_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrow_operator; }
	}

	public final Arrow_operatorContext arrow_operator() throws RecognitionException {
		Arrow_operatorContext _localctx = new Arrow_operatorContext(_ctx, getState());
		enterRule(_localctx, 444, RULE_arrow_operator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2234);
			if (!(SwiftSupport.isOperator(_input,"->"))) throw new FailedPredicateException(this, "SwiftSupport.isOperator(_input,\"->\")");
			setState(2235);
			match(SUB);
			setState(2236);
			match(GT);
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

	public static class Range_operatorContext extends ParserRuleContext {
		public List<TerminalNode> DOT() { return getTokens(TswiftParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(TswiftParser.DOT, i);
		}
		public Range_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_range_operator; }
	}

	public final Range_operatorContext range_operator() throws RecognitionException {
		Range_operatorContext _localctx = new Range_operatorContext(_ctx, getState());
		enterRule(_localctx, 446, RULE_range_operator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2238);
			if (!(SwiftSupport.isOperator(_input,"..."))) throw new FailedPredicateException(this, "SwiftSupport.isOperator(_input,\"...\")");
			setState(2239);
			match(DOT);
			setState(2240);
			match(DOT);
			setState(2241);
			match(DOT);
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

	public static class Same_type_equalsContext extends ParserRuleContext {
		public List<TerminalNode> EQUAL() { return getTokens(TswiftParser.EQUAL); }
		public TerminalNode EQUAL(int i) {
			return getToken(TswiftParser.EQUAL, i);
		}
		public Same_type_equalsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_same_type_equals; }
	}

	public final Same_type_equalsContext same_type_equals() throws RecognitionException {
		Same_type_equalsContext _localctx = new Same_type_equalsContext(_ctx, getState());
		enterRule(_localctx, 448, RULE_same_type_equals);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2243);
			if (!(SwiftSupport.isOperator(_input,"=="))) throw new FailedPredicateException(this, "SwiftSupport.isOperator(_input,\"==\")");
			setState(2244);
			match(EQUAL);
			setState(2245);
			match(EQUAL);
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

	public static class Binary_operatorContext extends ParserRuleContext {
		public Operator_Context operator_() {
			return getRuleContext(Operator_Context.class,0);
		}
		public Binary_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_binary_operator; }
	}

	public final Binary_operatorContext binary_operator() throws RecognitionException {
		Binary_operatorContext _localctx = new Binary_operatorContext(_ctx, getState());
		enterRule(_localctx, 450, RULE_binary_operator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2247);
			if (!(SwiftSupport.isBinaryOp(_input))) throw new FailedPredicateException(this, "SwiftSupport.isBinaryOp(_input)");
			setState(2248);
			operator_();
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

	public static class Prefix_operatorContext extends ParserRuleContext {
		public Operator_Context operator_() {
			return getRuleContext(Operator_Context.class,0);
		}
		public Prefix_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prefix_operator; }
	}

	public final Prefix_operatorContext prefix_operator() throws RecognitionException {
		Prefix_operatorContext _localctx = new Prefix_operatorContext(_ctx, getState());
		enterRule(_localctx, 452, RULE_prefix_operator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2250);
			if (!(SwiftSupport.isPrefixOp(_input))) throw new FailedPredicateException(this, "SwiftSupport.isPrefixOp(_input)");
			setState(2251);
			operator_();
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

	public static class Postfix_operatorContext extends ParserRuleContext {
		public Operator_Context operator_() {
			return getRuleContext(Operator_Context.class,0);
		}
		public Postfix_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_postfix_operator; }
	}

	public final Postfix_operatorContext postfix_operator() throws RecognitionException {
		Postfix_operatorContext _localctx = new Postfix_operatorContext(_ctx, getState());
		enterRule(_localctx, 454, RULE_postfix_operator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2253);
			if (!(SwiftSupport.isPostfixOp(_input))) throw new FailedPredicateException(this, "SwiftSupport.isPostfixOp(_input)");
			setState(2254);
			operator_();
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

	public static class Operator_Context extends ParserRuleContext {
		public Operator_headContext operator_head() {
			return getRuleContext(Operator_headContext.class,0);
		}
		public List<Operator_characterContext> operator_character() {
			return getRuleContexts(Operator_characterContext.class);
		}
		public Operator_characterContext operator_character(int i) {
			return getRuleContext(Operator_characterContext.class,i);
		}
		public Dot_operator_headContext dot_operator_head() {
			return getRuleContext(Dot_operator_headContext.class,0);
		}
		public List<Dot_operator_characterContext> dot_operator_character() {
			return getRuleContexts(Dot_operator_characterContext.class);
		}
		public Dot_operator_characterContext dot_operator_character(int i) {
			return getRuleContext(Dot_operator_characterContext.class,i);
		}
		public Operator_Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operator_; }
	}

	public final Operator_Context operator_() throws RecognitionException {
		Operator_Context _localctx = new Operator_Context(_ctx, getState());
		enterRule(_localctx, 456, RULE_operator_);
		try {
			int _alt;
			setState(2272);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LT:
			case GT:
			case BANG:
			case QUESTION:
			case AND:
			case SUB:
			case EQUAL:
			case OR:
			case DIV:
			case ADD:
			case MUL:
			case MOD:
			case CARET:
			case TILDE:
			case Operator_head_other:
				enterOuterAlt(_localctx, 1);
				{
				setState(2256);
				operator_head();
				setState(2261);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,249,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(2257);
						if (!(_input.get(_input.index()-1).getType()!=WS)) throw new FailedPredicateException(this, "_input.get(_input.index()-1).getType()!=WS");
						setState(2258);
						operator_character();
						}
						} 
					}
					setState(2263);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,249,_ctx);
				}
				}
				break;
			case DOT:
				enterOuterAlt(_localctx, 2);
				{
				setState(2264);
				dot_operator_head();
				setState(2269);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,250,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(2265);
						if (!(_input.get(_input.index()-1).getType()!=WS)) throw new FailedPredicateException(this, "_input.get(_input.index()-1).getType()!=WS");
						setState(2266);
						dot_operator_character();
						}
						} 
					}
					setState(2271);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,250,_ctx);
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

	public static class Operator_characterContext extends ParserRuleContext {
		public Operator_headContext operator_head() {
			return getRuleContext(Operator_headContext.class,0);
		}
		public TerminalNode Operator_following_character() { return getToken(TswiftParser.Operator_following_character, 0); }
		public Operator_characterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operator_character; }
	}

	public final Operator_characterContext operator_character() throws RecognitionException {
		Operator_characterContext _localctx = new Operator_characterContext(_ctx, getState());
		enterRule(_localctx, 458, RULE_operator_character);
		try {
			setState(2276);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LT:
			case GT:
			case BANG:
			case QUESTION:
			case AND:
			case SUB:
			case EQUAL:
			case OR:
			case DIV:
			case ADD:
			case MUL:
			case MOD:
			case CARET:
			case TILDE:
			case Operator_head_other:
				enterOuterAlt(_localctx, 1);
				{
				setState(2274);
				operator_head();
				}
				break;
			case Operator_following_character:
				enterOuterAlt(_localctx, 2);
				{
				setState(2275);
				match(Operator_following_character);
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

	public static class Operator_headContext extends ParserRuleContext {
		public TerminalNode DIV() { return getToken(TswiftParser.DIV, 0); }
		public TerminalNode EQUAL() { return getToken(TswiftParser.EQUAL, 0); }
		public TerminalNode SUB() { return getToken(TswiftParser.SUB, 0); }
		public TerminalNode ADD() { return getToken(TswiftParser.ADD, 0); }
		public TerminalNode BANG() { return getToken(TswiftParser.BANG, 0); }
		public TerminalNode MUL() { return getToken(TswiftParser.MUL, 0); }
		public TerminalNode MOD() { return getToken(TswiftParser.MOD, 0); }
		public TerminalNode AND() { return getToken(TswiftParser.AND, 0); }
		public TerminalNode OR() { return getToken(TswiftParser.OR, 0); }
		public TerminalNode LT() { return getToken(TswiftParser.LT, 0); }
		public TerminalNode GT() { return getToken(TswiftParser.GT, 0); }
		public TerminalNode CARET() { return getToken(TswiftParser.CARET, 0); }
		public TerminalNode TILDE() { return getToken(TswiftParser.TILDE, 0); }
		public TerminalNode QUESTION() { return getToken(TswiftParser.QUESTION, 0); }
		public TerminalNode Operator_head_other() { return getToken(TswiftParser.Operator_head_other, 0); }
		public Operator_headContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operator_head; }
	}

	public final Operator_headContext operator_head() throws RecognitionException {
		Operator_headContext _localctx = new Operator_headContext(_ctx, getState());
		enterRule(_localctx, 460, RULE_operator_head);
		int _la;
		try {
			setState(2280);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LT:
			case GT:
			case BANG:
			case QUESTION:
			case AND:
			case SUB:
			case EQUAL:
			case OR:
			case DIV:
			case ADD:
			case MUL:
			case MOD:
			case CARET:
			case TILDE:
				enterOuterAlt(_localctx, 1);
				{
				setState(2278);
				_la = _input.LA(1);
				if ( !(((((_la - 137)) & ~0x3f) == 0 && ((1L << (_la - 137)) & ((1L << (LT - 137)) | (1L << (GT - 137)) | (1L << (BANG - 137)) | (1L << (QUESTION - 137)) | (1L << (AND - 137)) | (1L << (SUB - 137)) | (1L << (EQUAL - 137)) | (1L << (OR - 137)) | (1L << (DIV - 137)) | (1L << (ADD - 137)) | (1L << (MUL - 137)) | (1L << (MOD - 137)) | (1L << (CARET - 137)) | (1L << (TILDE - 137)))) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case Operator_head_other:
				enterOuterAlt(_localctx, 2);
				{
				setState(2279);
				match(Operator_head_other);
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

	public static class Dot_operator_headContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(TswiftParser.DOT, 0); }
		public Dot_operator_headContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dot_operator_head; }
	}

	public final Dot_operator_headContext dot_operator_head() throws RecognitionException {
		Dot_operator_headContext _localctx = new Dot_operator_headContext(_ctx, getState());
		enterRule(_localctx, 462, RULE_dot_operator_head);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2282);
			match(DOT);
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

	public static class Dot_operator_characterContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(TswiftParser.DOT, 0); }
		public Operator_characterContext operator_character() {
			return getRuleContext(Operator_characterContext.class,0);
		}
		public Dot_operator_characterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dot_operator_character; }
	}

	public final Dot_operator_characterContext dot_operator_character() throws RecognitionException {
		Dot_operator_characterContext _localctx = new Dot_operator_characterContext(_ctx, getState());
		enterRule(_localctx, 464, RULE_dot_operator_character);
		try {
			setState(2286);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case DOT:
				enterOuterAlt(_localctx, 1);
				{
				setState(2284);
				match(DOT);
				}
				break;
			case LT:
			case GT:
			case BANG:
			case QUESTION:
			case AND:
			case SUB:
			case EQUAL:
			case OR:
			case DIV:
			case ADD:
			case MUL:
			case MOD:
			case CARET:
			case TILDE:
			case Operator_head_other:
			case Operator_following_character:
				enterOuterAlt(_localctx, 2);
				{
				setState(2285);
				operator_character();
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

	public static class LiteralContext extends ParserRuleContext {
		public Numeric_literalContext numeric_literal() {
			return getRuleContext(Numeric_literalContext.class,0);
		}
		public String_literalContext string_literal() {
			return getRuleContext(String_literalContext.class,0);
		}
		public Boolean_literalContext boolean_literal() {
			return getRuleContext(Boolean_literalContext.class,0);
		}
		public Nil_literalContext nil_literal() {
			return getRuleContext(Nil_literalContext.class,0);
		}
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 466, RULE_literal);
		try {
			setState(2292);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,255,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(2288);
				numeric_literal();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(2289);
				string_literal();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(2290);
				boolean_literal();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(2291);
				nil_literal();
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

	public static class Numeric_literalContext extends ParserRuleContext {
		public Integer_literalContext integer_literal() {
			return getRuleContext(Integer_literalContext.class,0);
		}
		public Negate_prefix_operatorContext negate_prefix_operator() {
			return getRuleContext(Negate_prefix_operatorContext.class,0);
		}
		public TerminalNode Floating_point_literal() { return getToken(TswiftParser.Floating_point_literal, 0); }
		public Numeric_literalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_numeric_literal; }
	}

	public final Numeric_literalContext numeric_literal() throws RecognitionException {
		Numeric_literalContext _localctx = new Numeric_literalContext(_ctx, getState());
		enterRule(_localctx, 468, RULE_numeric_literal);
		try {
			setState(2302);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,258,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(2295);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,256,_ctx) ) {
				case 1:
					{
					setState(2294);
					negate_prefix_operator();
					}
					break;
				}
				setState(2297);
				integer_literal();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(2299);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,257,_ctx) ) {
				case 1:
					{
					setState(2298);
					negate_prefix_operator();
					}
					break;
				}
				setState(2301);
				match(Floating_point_literal);
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

	public static class Boolean_literalContext extends ParserRuleContext {
		public Boolean_literalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_boolean_literal; }
	}

	public final Boolean_literalContext boolean_literal() throws RecognitionException {
		Boolean_literalContext _localctx = new Boolean_literalContext(_ctx, getState());
		enterRule(_localctx, 470, RULE_boolean_literal);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2304);
			_la = _input.LA(1);
			if ( !(_la==T__117 || _la==T__122) ) {
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

	public static class Nil_literalContext extends ParserRuleContext {
		public Nil_literalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_nil_literal; }
	}

	public final Nil_literalContext nil_literal() throws RecognitionException {
		Nil_literalContext _localctx = new Nil_literalContext(_ctx, getState());
		enterRule(_localctx, 472, RULE_nil_literal);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2306);
			match(T__119);
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

	public static class Integer_literalContext extends ParserRuleContext {
		public TerminalNode Binary_literal() { return getToken(TswiftParser.Binary_literal, 0); }
		public TerminalNode Octal_literal() { return getToken(TswiftParser.Octal_literal, 0); }
		public TerminalNode Decimal_literal() { return getToken(TswiftParser.Decimal_literal, 0); }
		public TerminalNode Pure_decimal_digits() { return getToken(TswiftParser.Pure_decimal_digits, 0); }
		public TerminalNode Hexadecimal_literal() { return getToken(TswiftParser.Hexadecimal_literal, 0); }
		public Integer_literalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_integer_literal; }
	}

	public final Integer_literalContext integer_literal() throws RecognitionException {
		Integer_literalContext _localctx = new Integer_literalContext(_ctx, getState());
		enterRule(_localctx, 474, RULE_integer_literal);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2308);
			_la = _input.LA(1);
			if ( !(((((_la - 156)) & ~0x3f) == 0 && ((1L << (_la - 156)) & ((1L << (Binary_literal - 156)) | (1L << (Octal_literal - 156)) | (1L << (Decimal_literal - 156)) | (1L << (Pure_decimal_digits - 156)) | (1L << (Hexadecimal_literal - 156)))) != 0)) ) {
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

	public static class String_literalContext extends ParserRuleContext {
		public TerminalNode Static_string_literal() { return getToken(TswiftParser.Static_string_literal, 0); }
		public TerminalNode Interpolated_string_literal() { return getToken(TswiftParser.Interpolated_string_literal, 0); }
		public String_literalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_string_literal; }
	}

	public final String_literalContext string_literal() throws RecognitionException {
		String_literalContext _localctx = new String_literalContext(_ctx, getState());
		enterRule(_localctx, 476, RULE_string_literal);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2310);
			_la = _input.LA(1);
			if ( !(_la==Static_string_literal || _la==Interpolated_string_literal) ) {
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 127:
			return pattern_sempred((PatternContext)_localctx, predIndex);
		case 143:
			return any_punctuation_for_balanced_token_sempred((Any_punctuation_for_balanced_tokenContext)_localctx, predIndex);
		case 185:
			return postfix_expresion_sempred((Postfix_expresionContext)_localctx, predIndex);
		case 193:
			return type__sempred((Type_Context)_localctx, predIndex);
		case 217:
			return assignment_operator_sempred((Assignment_operatorContext)_localctx, predIndex);
		case 218:
			return negate_prefix_operator_sempred((Negate_prefix_operatorContext)_localctx, predIndex);
		case 219:
			return compilation_condition_AND_sempred((Compilation_condition_ANDContext)_localctx, predIndex);
		case 220:
			return compilation_condition_OR_sempred((Compilation_condition_ORContext)_localctx, predIndex);
		case 221:
			return compilation_condition_GE_sempred((Compilation_condition_GEContext)_localctx, predIndex);
		case 222:
			return arrow_operator_sempred((Arrow_operatorContext)_localctx, predIndex);
		case 223:
			return range_operator_sempred((Range_operatorContext)_localctx, predIndex);
		case 224:
			return same_type_equals_sempred((Same_type_equalsContext)_localctx, predIndex);
		case 225:
			return binary_operator_sempred((Binary_operatorContext)_localctx, predIndex);
		case 226:
			return prefix_operator_sempred((Prefix_operatorContext)_localctx, predIndex);
		case 227:
			return postfix_operator_sempred((Postfix_operatorContext)_localctx, predIndex);
		case 228:
			return operator__sempred((Operator_Context)_localctx, predIndex);
		}
		return true;
	}
	private boolean pattern_sempred(PatternContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean any_punctuation_for_balanced_token_sempred(Any_punctuation_for_balanced_tokenContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return SwiftSupport.isPrefixOp(_input);
		case 2:
			return SwiftSupport.isPostfixOp(_input);
		}
		return true;
	}
	private boolean postfix_expresion_sempred(Postfix_expresionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 3:
			return precpred(_ctx, 12);
		case 4:
			return precpred(_ctx, 11);
		case 5:
			return precpred(_ctx, 10);
		case 6:
			return precpred(_ctx, 9);
		case 7:
			return precpred(_ctx, 8);
		case 8:
			return precpred(_ctx, 7);
		case 9:
			return precpred(_ctx, 6);
		case 10:
			return precpred(_ctx, 5);
		case 11:
			return precpred(_ctx, 4);
		case 12:
			return precpred(_ctx, 3);
		case 13:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean type__sempred(Type_Context _localctx, int predIndex) {
		switch (predIndex) {
		case 14:
			return precpred(_ctx, 7);
		case 15:
			return precpred(_ctx, 6);
		case 16:
			return precpred(_ctx, 4);
		case 17:
			return precpred(_ctx, 3);
		}
		return true;
	}
	private boolean assignment_operator_sempred(Assignment_operatorContext _localctx, int predIndex) {
		switch (predIndex) {
		case 18:
			return SwiftSupport.isBinaryOp(_input);
		}
		return true;
	}
	private boolean negate_prefix_operator_sempred(Negate_prefix_operatorContext _localctx, int predIndex) {
		switch (predIndex) {
		case 19:
			return SwiftSupport.isPrefixOp(_input);
		}
		return true;
	}
	private boolean compilation_condition_AND_sempred(Compilation_condition_ANDContext _localctx, int predIndex) {
		switch (predIndex) {
		case 20:
			return SwiftSupport.isOperator(_input,"&&");
		}
		return true;
	}
	private boolean compilation_condition_OR_sempred(Compilation_condition_ORContext _localctx, int predIndex) {
		switch (predIndex) {
		case 21:
			return SwiftSupport.isOperator(_input,"||");
		}
		return true;
	}
	private boolean compilation_condition_GE_sempred(Compilation_condition_GEContext _localctx, int predIndex) {
		switch (predIndex) {
		case 22:
			return SwiftSupport.isOperator(_input,">=");
		}
		return true;
	}
	private boolean arrow_operator_sempred(Arrow_operatorContext _localctx, int predIndex) {
		switch (predIndex) {
		case 23:
			return SwiftSupport.isOperator(_input,"->");
		}
		return true;
	}
	private boolean range_operator_sempred(Range_operatorContext _localctx, int predIndex) {
		switch (predIndex) {
		case 24:
			return SwiftSupport.isOperator(_input,"...");
		}
		return true;
	}
	private boolean same_type_equals_sempred(Same_type_equalsContext _localctx, int predIndex) {
		switch (predIndex) {
		case 25:
			return SwiftSupport.isOperator(_input,"==");
		}
		return true;
	}
	private boolean binary_operator_sempred(Binary_operatorContext _localctx, int predIndex) {
		switch (predIndex) {
		case 26:
			return SwiftSupport.isBinaryOp(_input);
		}
		return true;
	}
	private boolean prefix_operator_sempred(Prefix_operatorContext _localctx, int predIndex) {
		switch (predIndex) {
		case 27:
			return SwiftSupport.isPrefixOp(_input);
		}
		return true;
	}
	private boolean postfix_operator_sempred(Postfix_operatorContext _localctx, int predIndex) {
		switch (predIndex) {
		case 28:
			return SwiftSupport.isPostfixOp(_input);
		}
		return true;
	}
	private boolean operator__sempred(Operator_Context _localctx, int predIndex) {
		switch (predIndex) {
		case 29:
			return _input.get(_input.index()-1).getType()!=WS;
		case 30:
			return _input.get(_input.index()-1).getType()!=WS;
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\u00a8\u090b\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4I"+
		"\tI\4J\tJ\4K\tK\4L\tL\4M\tM\4N\tN\4O\tO\4P\tP\4Q\tQ\4R\tR\4S\tS\4T\tT"+
		"\4U\tU\4V\tV\4W\tW\4X\tX\4Y\tY\4Z\tZ\4[\t[\4\\\t\\\4]\t]\4^\t^\4_\t_\4"+
		"`\t`\4a\ta\4b\tb\4c\tc\4d\td\4e\te\4f\tf\4g\tg\4h\th\4i\ti\4j\tj\4k\t"+
		"k\4l\tl\4m\tm\4n\tn\4o\to\4p\tp\4q\tq\4r\tr\4s\ts\4t\tt\4u\tu\4v\tv\4"+
		"w\tw\4x\tx\4y\ty\4z\tz\4{\t{\4|\t|\4}\t}\4~\t~\4\177\t\177\4\u0080\t\u0080"+
		"\4\u0081\t\u0081\4\u0082\t\u0082\4\u0083\t\u0083\4\u0084\t\u0084\4\u0085"+
		"\t\u0085\4\u0086\t\u0086\4\u0087\t\u0087\4\u0088\t\u0088\4\u0089\t\u0089"+
		"\4\u008a\t\u008a\4\u008b\t\u008b\4\u008c\t\u008c\4\u008d\t\u008d\4\u008e"+
		"\t\u008e\4\u008f\t\u008f\4\u0090\t\u0090\4\u0091\t\u0091\4\u0092\t\u0092"+
		"\4\u0093\t\u0093\4\u0094\t\u0094\4\u0095\t\u0095\4\u0096\t\u0096\4\u0097"+
		"\t\u0097\4\u0098\t\u0098\4\u0099\t\u0099\4\u009a\t\u009a\4\u009b\t\u009b"+
		"\4\u009c\t\u009c\4\u009d\t\u009d\4\u009e\t\u009e\4\u009f\t\u009f\4\u00a0"+
		"\t\u00a0\4\u00a1\t\u00a1\4\u00a2\t\u00a2\4\u00a3\t\u00a3\4\u00a4\t\u00a4"+
		"\4\u00a5\t\u00a5\4\u00a6\t\u00a6\4\u00a7\t\u00a7\4\u00a8\t\u00a8\4\u00a9"+
		"\t\u00a9\4\u00aa\t\u00aa\4\u00ab\t\u00ab\4\u00ac\t\u00ac\4\u00ad\t\u00ad"+
		"\4\u00ae\t\u00ae\4\u00af\t\u00af\4\u00b0\t\u00b0\4\u00b1\t\u00b1\4\u00b2"+
		"\t\u00b2\4\u00b3\t\u00b3\4\u00b4\t\u00b4\4\u00b5\t\u00b5\4\u00b6\t\u00b6"+
		"\4\u00b7\t\u00b7\4\u00b8\t\u00b8\4\u00b9\t\u00b9\4\u00ba\t\u00ba\4\u00bb"+
		"\t\u00bb\4\u00bc\t\u00bc\4\u00bd\t\u00bd\4\u00be\t\u00be\4\u00bf\t\u00bf"+
		"\4\u00c0\t\u00c0\4\u00c1\t\u00c1\4\u00c2\t\u00c2\4\u00c3\t\u00c3\4\u00c4"+
		"\t\u00c4\4\u00c5\t\u00c5\4\u00c6\t\u00c6\4\u00c7\t\u00c7\4\u00c8\t\u00c8"+
		"\4\u00c9\t\u00c9\4\u00ca\t\u00ca\4\u00cb\t\u00cb\4\u00cc\t\u00cc\4\u00cd"+
		"\t\u00cd\4\u00ce\t\u00ce\4\u00cf\t\u00cf\4\u00d0\t\u00d0\4\u00d1\t\u00d1"+
		"\4\u00d2\t\u00d2\4\u00d3\t\u00d3\4\u00d4\t\u00d4\4\u00d5\t\u00d5\4\u00d6"+
		"\t\u00d6\4\u00d7\t\u00d7\4\u00d8\t\u00d8\4\u00d9\t\u00d9\4\u00da\t\u00da"+
		"\4\u00db\t\u00db\4\u00dc\t\u00dc\4\u00dd\t\u00dd\4\u00de\t\u00de\4\u00df"+
		"\t\u00df\4\u00e0\t\u00e0\4\u00e1\t\u00e1\4\u00e2\t\u00e2\4\u00e3\t\u00e3"+
		"\4\u00e4\t\u00e4\4\u00e5\t\u00e5\4\u00e6\t\u00e6\4\u00e7\t\u00e7\4\u00e8"+
		"\t\u00e8\4\u00e9\t\u00e9\4\u00ea\t\u00ea\4\u00eb\t\u00eb\4\u00ec\t\u00ec"+
		"\4\u00ed\t\u00ed\4\u00ee\t\u00ee\4\u00ef\t\u00ef\4\u00f0\t\u00f0\3\2\3"+
		"\2\3\2\3\3\7\3\u01e5\n\3\f\3\16\3\u01e8\13\3\3\4\3\4\3\4\3\4\3\4\5\4\u01ef"+
		"\n\4\3\5\3\5\5\5\u01f3\n\5\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\b"+
		"\3\b\3\b\5\b\u0202\n\b\3\t\3\t\3\t\3\t\5\t\u0208\n\t\3\n\3\n\3\n\3\n\5"+
		"\n\u020e\n\n\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\5\f\u0219\n\f\3"+
		"\f\3\f\3\r\3\r\5\r\u021f\n\r\3\16\3\16\3\16\3\16\3\16\3\16\5\16\u0227"+
		"\n\16\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\21\3\21\3\21\5\21\u0233\n\21"+
		"\3\22\3\22\3\23\3\23\3\24\3\24\5\24\u023b\n\24\3\25\3\25\3\25\3\25\3\26"+
		"\3\26\3\26\7\26\u0244\n\26\f\26\16\26\u0247\13\26\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\5\27\u0252\n\27\3\30\3\30\3\30\3\31\3\31\3\31"+
		"\7\31\u025a\n\31\f\31\16\31\u025d\13\31\3\32\3\32\5\32\u0261\n\32\3\33"+
		"\3\33\3\33\3\33\3\33\3\33\3\33\3\33\5\33\u026b\n\33\3\34\3\34\3\34\3\34"+
		"\3\35\3\35\3\35\3\35\3\36\3\36\3\36\7\36\u0278\n\36\f\36\16\36\u027b\13"+
		"\36\3\37\3\37\3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \5 \u028d\n \3"+
		"!\6!\u0290\n!\r!\16!\u0291\3\"\3\"\5\"\u0296\n\"\3\"\3\"\3#\3#\3$\3$\3"+
		"$\7$\u029f\n$\f$\16$\u02a2\13$\3%\3%\5%\u02a6\n%\3&\3&\3&\3\'\3\'\3\'"+
		"\7\'\u02ae\n\'\f\'\16\'\u02b1\13\'\3(\3(\5(\u02b5\n(\3)\3)\3)\3*\3*\3"+
		"*\3*\3*\3*\3*\3*\3*\3*\3*\3*\3*\3*\3*\3*\3*\3*\3*\5*\u02cd\n*\3*\3*\3"+
		"*\3*\3*\3*\3*\5*\u02d6\n*\3*\3*\3*\3*\3*\5*\u02dd\n*\3+\5+\u02e0\n+\3"+
		"+\5+\u02e3\n+\3+\3+\3,\3,\3-\3-\3-\5-\u02ec\n-\3-\3-\3-\3-\3-\3-\3-\5"+
		"-\u02f5\n-\3.\5.\u02f8\n.\3.\5.\u02fb\n.\3.\3.\3.\3/\5/\u0301\n/\3/\5"+
		"/\u0304\n/\3/\3/\5/\u0308\n/\3/\3/\3\60\3\60\3\60\3\60\3\61\3\61\3\61"+
		"\5\61\u0313\n\61\3\61\3\61\3\61\3\61\3\61\3\61\3\61\5\61\u031c\n\61\3"+
		"\62\5\62\u031f\n\62\3\62\5\62\u0322\n\62\3\62\3\62\3\63\5\63\u0327\n\63"+
		"\3\63\5\63\u032a\n\63\3\63\3\63\3\64\3\64\3\64\5\64\u0331\n\64\3\64\3"+
		"\64\3\64\3\64\3\64\3\64\3\64\5\64\u033a\n\64\3\65\5\65\u033d\n\65\3\65"+
		"\3\65\5\65\u0341\n\65\3\65\3\65\3\66\5\66\u0346\n\66\3\66\3\66\5\66\u034a"+
		"\n\66\3\66\3\66\3\67\3\67\3\67\5\67\u0351\n\67\3\67\3\67\5\67\u0355\n"+
		"\67\3\67\5\67\u0358\n\67\38\58\u035b\n8\38\58\u035e\n8\38\38\39\39\59"+
		"\u0364\n9\3:\3:\5:\u0368\n:\3:\5:\u036b\n:\3:\3:\3:\5:\u0370\n:\5:\u0372"+
		"\n:\3;\3;\5;\u0376\n;\3;\3;\3<\3<\3=\3=\3=\3=\3=\3=\5=\u0382\n=\3>\3>"+
		"\3>\7>\u0387\n>\f>\16>\u038a\13>\3?\5?\u038d\n?\3?\3?\3?\5?\u0392\n?\3"+
		"?\5?\u0395\n?\3?\3?\3?\3?\5?\u039b\n?\3?\3?\3?\3?\5?\u03a1\n?\3@\3@\3"+
		"A\3A\3B\3B\3B\3C\5C\u03ab\nC\3C\5C\u03ae\nC\3C\3C\5C\u03b2\nC\3C\5C\u03b5"+
		"\nC\3C\5C\u03b8\nC\3D\5D\u03bb\nD\3D\3D\3D\5D\u03c0\nD\3D\5D\u03c3\nD"+
		"\3D\5D\u03c6\nD\3D\3D\5D\u03ca\nD\3D\3D\3E\3E\5E\u03d0\nE\3F\3F\5F\u03d4"+
		"\nF\3G\5G\u03d7\nG\3G\5G\u03da\nG\3G\3G\3G\3H\3H\3H\3H\3H\5H\u03e4\nH"+
		"\3I\3I\5I\u03e8\nI\3J\3J\3K\3K\3L\3L\3L\5L\u03f1\nL\3L\3L\5L\u03f5\nL"+
		"\3L\3L\3L\3L\3M\3M\5M\u03fd\nM\3N\3N\5N\u0401\nN\3O\5O\u0404\nO\3O\3O"+
		"\3O\3P\3P\3P\3P\3P\5P\u040e\nP\3Q\3Q\5Q\u0412\nQ\3R\3R\3R\3S\3S\3S\5S"+
		"\u041a\nS\3T\5T\u041d\nT\3T\5T\u0420\nT\3T\3T\3T\5T\u0425\nT\3T\5T\u0428"+
		"\nT\3T\5T\u042b\nT\3T\3T\3U\3U\3V\3V\7V\u0433\nV\fV\16V\u0436\13V\3V\3"+
		"V\3W\3W\3X\5X\u043d\nX\3X\5X\u0440\nX\3X\5X\u0443\nX\3X\3X\3X\5X\u0448"+
		"\nX\3X\5X\u044b\nX\3X\5X\u044e\nX\3X\3X\3X\5X\u0453\nX\3X\5X\u0456\nX"+
		"\3X\3X\5X\u045a\nX\3X\3X\3X\5X\u045f\nX\3X\5X\u0462\nX\3X\5X\u0465\nX"+
		"\3X\3X\5X\u0469\nX\3Y\3Y\3Z\3Z\7Z\u046f\nZ\fZ\16Z\u0472\13Z\3Z\3Z\3[\3"+
		"[\3\\\5\\\u0479\n\\\3\\\5\\\u047c\n\\\3\\\3\\\3\\\5\\\u0481\n\\\3\\\3"+
		"\\\3]\3]\3^\3^\7^\u0489\n^\f^\16^\u048c\13^\3^\3^\3_\3_\3`\3`\3`\3`\3"+
		"`\5`\u0497\n`\3a\3a\3a\3a\3a\3b\3b\3b\5b\u04a1\nb\3b\3b\5b\u04a5\nb\3"+
		"c\3c\5c\u04a9\nc\3c\3c\5c\u04ad\nc\3c\5c\u04b0\nc\3c\3c\5c\u04b4\nc\3"+
		"c\3c\3c\5c\u04b9\nc\5c\u04bb\nc\3d\3d\3d\3d\3e\5e\u04c2\ne\3e\5e\u04c5"+
		"\ne\3e\3e\5e\u04c9\ne\3f\3f\5f\u04cd\nf\3f\3f\5f\u04d1\nf\3f\5f\u04d4"+
		"\nf\3f\3f\3f\3f\5f\u04da\nf\3f\3f\3f\5f\u04df\nf\3f\3f\5f\u04e3\nf\3g"+
		"\5g\u04e6\ng\3g\5g\u04e9\ng\3g\3g\5g\u04ed\ng\3g\5g\u04f0\ng\3g\3g\3g"+
		"\5g\u04f5\ng\3g\5g\u04f8\ng\3g\3g\5g\u04fc\ng\3h\3h\3i\5i\u0501\ni\3i"+
		"\3i\3i\3j\5j\u0507\nj\3j\5j\u050a\nj\3j\3j\3j\5j\u050f\nj\3j\3j\3j\5j"+
		"\u0514\nj\3j\5j\u0517\nj\3j\3j\3j\3j\3j\5j\u051e\nj\3k\3k\7k\u0522\nk"+
		"\fk\16k\u0525\13k\3k\3k\3l\3l\3m\3m\3m\3m\3m\3m\3m\3m\3m\3m\3m\3m\5m\u0537"+
		"\nm\3n\5n\u053a\nn\3n\5n\u053d\nn\3n\3n\3n\3o\3o\5o\u0544\no\3o\3o\3p"+
		"\3p\3p\5p\u054b\np\3q\3q\3q\3q\3r\3r\3r\3r\3s\3s\3s\3s\5s\u0559\ns\3t"+
		"\3t\3t\3u\3u\3u\3u\7u\u0562\nu\fu\16u\u0565\13u\3u\3u\3v\3v\3v\5v\u056c"+
		"\nv\3w\3w\3w\3w\3w\3w\5w\u0574\nw\3x\3x\3x\3x\3y\3y\3y\3y\3z\3z\3{\3{"+
		"\3{\7{\u0583\n{\f{\16{\u0586\13{\3|\3|\3}\3}\3}\3}\3}\3}\3}\3}\3}\3}\3"+
		"}\3}\3}\3}\3}\3}\3}\3}\3}\3}\3}\3}\3}\3}\5}\u05a2\n}\3~\6~\u05a5\n~\r"+
		"~\16~\u05a6\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177"+
		"\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177"+
		"\3\177\3\177\3\177\5\177\u05c2\n\177\3\u0080\3\u0080\3\u0081\3\u0081\3"+
		"\u0081\5\u0081\u05c9\n\u0081\3\u0081\3\u0081\5\u0081\u05cd\n\u0081\3\u0081"+
		"\3\u0081\3\u0081\5\u0081\u05d2\n\u0081\3\u0081\3\u0081\3\u0081\3\u0081"+
		"\3\u0081\5\u0081\u05d9\n\u0081\3\u0081\3\u0081\3\u0081\7\u0081\u05de\n"+
		"\u0081\f\u0081\16\u0081\u05e1\13\u0081\3\u0082\3\u0082\3\u0083\3\u0083"+
		"\3\u0084\3\u0084\3\u0084\3\u0084\5\u0084\u05eb\n\u0084\3\u0085\3\u0085"+
		"\5\u0085\u05ef\n\u0085\3\u0085\3\u0085\3\u0086\3\u0086\3\u0086\7\u0086"+
		"\u05f6\n\u0086\f\u0086\16\u0086\u05f9\13\u0086\3\u0087\3\u0087\3\u0088"+
		"\5\u0088\u05fe\n\u0088\3\u0088\3\u0088\3\u0088\5\u0088\u0603\n\u0088\3"+
		"\u0089\3\u0089\3\u0089\3\u008a\3\u008a\3\u008b\3\u008b\3\u008b\5\u008b"+
		"\u060d\n\u008b\3\u008c\3\u008c\3\u008d\3\u008d\3\u008d\3\u008d\3\u008e"+
		"\6\u008e\u0616\n\u008e\r\u008e\16\u008e\u0617\3\u008f\7\u008f\u061b\n"+
		"\u008f\f\u008f\16\u008f\u061e\13\u008f\3\u0090\3\u0090\3\u0090\3\u0090"+
		"\3\u0090\3\u0090\3\u0090\3\u0090\3\u0090\3\u0090\3\u0090\3\u0090\3\u0090"+
		"\3\u0090\3\u0090\3\u0090\5\u0090\u0630\n\u0090\3\u0091\3\u0091\3\u0091"+
		"\3\u0091\3\u0091\3\u0091\5\u0091\u0638\n\u0091\3\u0092\5\u0092\u063b\n"+
		"\u0092\3\u0092\3\u0092\5\u0092\u063f\n\u0092\3\u0093\3\u0093\3\u0093\7"+
		"\u0093\u0644\n\u0093\f\u0093\16\u0093\u0647\13\u0093\3\u0094\3\u0094\3"+
		"\u0094\3\u0094\3\u0094\5\u0094\u064e\n\u0094\3\u0095\3\u0095\3\u0095\3"+
		"\u0096\3\u0096\3\u0096\3\u0096\3\u0096\5\u0096\u0658\n\u0096\3\u0097\3"+
		"\u0097\3\u0097\3\u0097\3\u0097\5\u0097\u065f\n\u0097\3\u0097\3\u0097\3"+
		"\u0097\3\u0097\5\u0097\u0665\n\u0097\3\u0097\3\u0097\3\u0097\5\u0097\u066a"+
		"\n\u0097\3\u0098\6\u0098\u066d\n\u0098\r\u0098\16\u0098\u066e\3\u0099"+
		"\3\u0099\5\u0099\u0673\n\u0099\3\u0099\3\u0099\3\u0099\3\u009a\3\u009a"+
		"\3\u009a\3\u009a\3\u009a\3\u009a\3\u009a\3\u009a\3\u009a\3\u009a\5\u009a"+
		"\u0682\n\u009a\3\u009b\3\u009b\5\u009b\u0686\n\u009b\3\u009b\3\u009b\3"+
		"\u009b\3\u009b\3\u009b\3\u009b\3\u009b\3\u009b\3\u009b\3\u009b\5\u009b"+
		"\u0692\n\u009b\3\u009c\3\u009c\3\u009c\3\u009c\3\u009c\3\u009c\3\u009c"+
		"\3\u009c\5\u009c\u069c\n\u009c\3\u009d\3\u009d\5\u009d\u06a0\n\u009d\3"+
		"\u009d\3\u009d\3\u009e\3\u009e\5\u009e\u06a6\n\u009e\3\u009e\3\u009e\3"+
		"\u009e\3\u009e\5\u009e\u06ac\n\u009e\3\u009f\3\u009f\3\u00a0\3\u00a0\3"+
		"\u00a0\3\u00a0\3\u00a0\3\u00a0\3\u00a0\5\u00a0\u06b7\n\u00a0\3\u00a1\3"+
		"\u00a1\5\u00a1\u06bb\n\u00a1\3\u00a1\3\u00a1\3\u00a1\3\u00a1\5\u00a1\u06c1"+
		"\n\u00a1\3\u00a2\3\u00a2\3\u00a2\3\u00a2\3\u00a3\3\u00a3\3\u00a3\3\u00a3"+
		"\3\u00a3\3\u00a3\3\u00a3\3\u00a3\3\u00a3\3\u00a3\3\u00a3\3\u00a3\3\u00a3"+
		"\3\u00a3\3\u00a3\3\u00a3\3\u00a3\3\u00a3\3\u00a3\3\u00a3\3\u00a3\3\u00a3"+
		"\3\u00a3\3\u00a3\3\u00a3\3\u00a3\3\u00a3\3\u00a3\3\u00a3\3\u00a3\3\u00a3"+
		"\3\u00a3\3\u00a3\5\u00a3\u06e8\n\u00a3\3\u00a4\3\u00a4\3\u00a4\3\u00a4"+
		"\3\u00a4\3\u00a4\3\u00a4\3\u00a4\3\u00a4\3\u00a4\3\u00a4\3\u00a4\3\u00a4"+
		"\3\u00a4\3\u00a4\3\u00a4\3\u00a4\3\u00a4\3\u00a4\5\u00a4\u06fd\n\u00a4"+
		"\3\u00a5\3\u00a5\3\u00a5\5\u00a5\u0702\n\u00a5\3\u00a6\3\u00a6\3\u00a6"+
		"\3\u00a6\3\u00a7\3\u00a7\3\u00a7\3\u00a7\3\u00a7\3\u00a8\3\u00a8\3\u00a8"+
		"\3\u00a8\3\u00a9\3\u00a9\5\u00a9\u0713\n\u00a9\3\u00a9\5\u00a9\u0716\n"+
		"\u00a9\3\u00a9\3\u00a9\3\u00aa\5\u00aa\u071b\n\u00aa\3\u00aa\3\u00aa\5"+
		"\u00aa\u071f\n\u00aa\3\u00aa\5\u00aa\u0722\n\u00aa\3\u00aa\3\u00aa\3\u00aa"+
		"\3\u00aa\3\u00aa\5\u00aa\u0729\n\u00aa\3\u00ab\3\u00ab\3\u00ab\3\u00ab"+
		"\3\u00ab\3\u00ab\3\u00ab\5\u00ab\u0732\n\u00ab\3\u00ac\3\u00ac\3\u00ac"+
		"\7\u00ac\u0737\n\u00ac\f\u00ac\16\u00ac\u073a\13\u00ac\3\u00ad\3\u00ad"+
		"\3\u00ad\7\u00ad\u073f\n\u00ad\f\u00ad\16\u00ad\u0742\13\u00ad\3\u00ae"+
		"\3\u00ae\5\u00ae\u0746\n\u00ae\3\u00ae\3\u00ae\3\u00ae\3\u00ae\5\u00ae"+
		"\u074c\n\u00ae\3\u00af\3\u00af\3\u00b0\3\u00b0\3\u00b0\3\u00b0\3\u00b1"+
		"\3\u00b1\3\u00b1\7\u00b1\u0757\n\u00b1\f\u00b1\16\u00b1\u075a\13\u00b1"+
		"\3\u00b2\5\u00b2\u075d\n\u00b2\3\u00b2\3\u00b2\3\u00b3\3\u00b3\3\u00b4"+
		"\3\u00b4\3\u00b4\3\u00b5\3\u00b5\3\u00b5\3\u00b5\3\u00b6\3\u00b6\3\u00b6"+
		"\3\u00b6\3\u00b6\3\u00b6\6\u00b6\u0770\n\u00b6\r\u00b6\16\u00b6\u0771"+
		"\3\u00b6\3\u00b6\5\u00b6\u0776\n\u00b6\3\u00b7\3\u00b7\3\u00b7\3\u00b7"+
		"\3\u00b7\5\u00b7\u077d\n\u00b7\3\u00b8\3\u00b8\3\u00b9\3\u00b9\3\u00b9"+
		"\3\u00b9\3\u00b9\3\u00b9\3\u00b9\3\u00b9\3\u00b9\3\u00b9\3\u00b9\3\u00b9"+
		"\3\u00b9\3\u00b9\3\u00b9\3\u00b9\3\u00b9\5\u00b9\u0792\n\u00b9\3\u00ba"+
		"\3\u00ba\3\u00ba\3\u00ba\3\u00ba\3\u00bb\3\u00bb\3\u00bb\5\u00bb\u079c"+
		"\n\u00bb\3\u00bb\3\u00bb\3\u00bb\3\u00bb\3\u00bb\3\u00bb\5\u00bb\u07a4"+
		"\n\u00bb\3\u00bb\3\u00bb\3\u00bb\3\u00bb\3\u00bb\3\u00bb\3\u00bb\3\u00bb"+
		"\3\u00bb\3\u00bb\3\u00bb\3\u00bb\3\u00bb\3\u00bb\3\u00bb\3\u00bb\3\u00bb"+
		"\3\u00bb\5\u00bb\u07b8\n\u00bb\3\u00bb\3\u00bb\3\u00bb\3\u00bb\3\u00bb"+
		"\3\u00bb\3\u00bb\3\u00bb\3\u00bb\3\u00bb\3\u00bb\3\u00bb\3\u00bb\3\u00bb"+
		"\3\u00bb\3\u00bb\3\u00bb\3\u00bb\3\u00bb\3\u00bb\7\u00bb\u07ce\n\u00bb"+
		"\f\u00bb\16\u00bb\u07d1\13\u00bb\3\u00bc\3\u00bc\3\u00bc\3\u00bc\3\u00bc"+
		"\3\u00bc\5\u00bc\u07d9\n\u00bc\3\u00bd\3\u00bd\3\u00bd\7\u00bd\u07de\n"+
		"\u00bd\f\u00bd\16\u00bd\u07e1\13\u00bd\3\u00be\3\u00be\3\u00be\3\u00be"+
		"\3\u00be\3\u00be\3\u00be\3\u00be\3\u00be\3\u00be\5\u00be\u07ed\n\u00be"+
		"\3\u00bf\3\u00bf\3\u00c0\3\u00c0\7\u00c0\u07f3\n\u00c0\f\u00c0\16\u00c0"+
		"\u07f6\13\u00c0\3\u00c1\3\u00c1\3\u00c1\3\u00c2\3\u00c2\3\u00c2\3\u00c2"+
		"\3\u00c2\3\u00c2\3\u00c2\3\u00c3\3\u00c3\3\u00c3\3\u00c3\3\u00c3\3\u00c3"+
		"\3\u00c3\3\u00c3\3\u00c3\5\u00c3\u080b\n\u00c3\3\u00c3\3\u00c3\3\u00c3"+
		"\3\u00c3\3\u00c3\3\u00c3\3\u00c3\3\u00c3\3\u00c3\3\u00c3\7\u00c3\u0817"+
		"\n\u00c3\f\u00c3\16\u00c3\u081a\13\u00c3\3\u00c4\3\u00c4\5\u00c4\u081e"+
		"\n\u00c4\3\u00c4\5\u00c4\u0821\n\u00c4\3\u00c4\3\u00c4\3\u00c5\3\u00c5"+
		"\5\u00c5\u0827\n\u00c5\3\u00c5\3\u00c5\5\u00c5\u082b\n\u00c5\3\u00c6\3"+
		"\u00c6\3\u00c7\3\u00c7\5\u00c7\u0831\n\u00c7\3\u00c7\3\u00c7\3\u00c8\3"+
		"\u00c8\3\u00c8\3\u00c8\3\u00c8\5\u00c8\u083a\n\u00c8\3\u00c9\3\u00c9\3"+
		"\u00c9\3\u00c9\5\u00c9\u0840\n\u00c9\3\u00ca\3\u00ca\3\u00cb\5\u00cb\u0845"+
		"\n\u00cb\3\u00cb\3\u00cb\5\u00cb\u0849\n\u00cb\3\u00cb\3\u00cb\3\u00cb"+
		"\3\u00cb\5\u00cb\u084f\n\u00cb\3\u00cb\3\u00cb\3\u00cb\3\u00cb\3\u00cb"+
		"\5\u00cb\u0856\n\u00cb\3\u00cc\3\u00cc\3\u00cc\3\u00cc\3\u00cc\5\u00cc"+
		"\u085d\n\u00cc\3\u00cc\3\u00cc\5\u00cc\u0861\n\u00cc\3\u00cd\3\u00cd\3"+
		"\u00cd\3\u00cd\3\u00cd\5\u00cd\u0868\n\u00cd\3\u00ce\5\u00ce\u086b\n\u00ce"+
		"\3\u00ce\5\u00ce\u086e\n\u00ce\3\u00ce\3\u00ce\3\u00ce\3\u00ce\5\u00ce"+
		"\u0874\n\u00ce\3\u00cf\3\u00cf\3\u00d0\3\u00d0\3\u00d0\3\u00d0\3\u00d1"+
		"\3\u00d1\3\u00d1\3\u00d1\3\u00d1\3\u00d1\3\u00d2\3\u00d2\3\u00d2\6\u00d2"+
		"\u0885\n\u00d2\r\u00d2\16\u00d2\u0886\3\u00d3\3\u00d3\3\u00d4\3\u00d4"+
		"\3\u00d4\3\u00d4\3\u00d4\3\u00d4\3\u00d4\3\u00d4\3\u00d4\5\u00d4\u0894"+
		"\n\u00d4\3\u00d5\3\u00d5\3\u00d5\3\u00d5\3\u00d5\5\u00d5\u089b\n\u00d5"+
		"\3\u00d6\3\u00d6\3\u00d7\3\u00d7\5\u00d7\u08a1\n\u00d7\3\u00d8\3\u00d8"+
		"\5\u00d8\u08a5\n\u00d8\3\u00d9\3\u00d9\3\u00da\3\u00da\3\u00db\3\u00db"+
		"\3\u00db\3\u00dc\3\u00dc\3\u00dc\3\u00dd\3\u00dd\3\u00dd\3\u00dd\3\u00de"+
		"\3\u00de\3\u00de\3\u00de\3\u00df\3\u00df\3\u00df\3\u00df\3\u00e0\3\u00e0"+
		"\3\u00e0\3\u00e0\3\u00e1\3\u00e1\3\u00e1\3\u00e1\3\u00e1\3\u00e2\3\u00e2"+
		"\3\u00e2\3\u00e2\3\u00e3\3\u00e3\3\u00e3\3\u00e4\3\u00e4\3\u00e4\3\u00e5"+
		"\3\u00e5\3\u00e5\3\u00e6\3\u00e6\3\u00e6\7\u00e6\u08d6\n\u00e6\f\u00e6"+
		"\16\u00e6\u08d9\13\u00e6\3\u00e6\3\u00e6\3\u00e6\7\u00e6\u08de\n\u00e6"+
		"\f\u00e6\16\u00e6\u08e1\13\u00e6\5\u00e6\u08e3\n\u00e6\3\u00e7\3\u00e7"+
		"\5\u00e7\u08e7\n\u00e7\3\u00e8\3\u00e8\5\u00e8\u08eb\n\u00e8\3\u00e9\3"+
		"\u00e9\3\u00ea\3\u00ea\5\u00ea\u08f1\n\u00ea\3\u00eb\3\u00eb\3\u00eb\3"+
		"\u00eb\5\u00eb\u08f7\n\u00eb\3\u00ec\5\u00ec\u08fa\n\u00ec\3\u00ec\3\u00ec"+
		"\5\u00ec\u08fe\n\u00ec\3\u00ec\5\u00ec\u0901\n\u00ec\3\u00ed\3\u00ed\3"+
		"\u00ee\3\u00ee\3\u00ef\3\u00ef\3\u00f0\3\u00f0\3\u00f0\2\5\u0100\u0174"+
		"\u0184\u00f1\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\62\64"+
		"\668:<>@BDFHJLNPRTVXZ\\^`bdfhjlnprtvxz|~\u0080\u0082\u0084\u0086\u0088"+
		"\u008a\u008c\u008e\u0090\u0092\u0094\u0096\u0098\u009a\u009c\u009e\u00a0"+
		"\u00a2\u00a4\u00a6\u00a8\u00aa\u00ac\u00ae\u00b0\u00b2\u00b4\u00b6\u00b8"+
		"\u00ba\u00bc\u00be\u00c0\u00c2\u00c4\u00c6\u00c8\u00ca\u00cc\u00ce\u00d0"+
		"\u00d2\u00d4\u00d6\u00d8\u00da\u00dc\u00de\u00e0\u00e2\u00e4\u00e6\u00e8"+
		"\u00ea\u00ec\u00ee\u00f0\u00f2\u00f4\u00f6\u00f8\u00fa\u00fc\u00fe\u0100"+
		"\u0102\u0104\u0106\u0108\u010a\u010c\u010e\u0110\u0112\u0114\u0116\u0118"+
		"\u011a\u011c\u011e\u0120\u0122\u0124\u0126\u0128\u012a\u012c\u012e\u0130"+
		"\u0132\u0134\u0136\u0138\u013a\u013c\u013e\u0140\u0142\u0144\u0146\u0148"+
		"\u014a\u014c\u014e\u0150\u0152\u0154\u0156\u0158\u015a\u015c\u015e\u0160"+
		"\u0162\u0164\u0166\u0168\u016a\u016c\u016e\u0170\u0172\u0174\u0176\u0178"+
		"\u017a\u017c\u017e\u0180\u0182\u0184\u0186\u0188\u018a\u018c\u018e\u0190"+
		"\u0192\u0194\u0196\u0198\u019a\u019c\u019e\u01a0\u01a2\u01a4\u01a6\u01a8"+
		"\u01aa\u01ac\u01ae\u01b0\u01b2\u01b4\u01b6\u01b8\u01ba\u01bc\u01be\u01c0"+
		"\u01c2\u01c4\u01c6\u01c8\u01ca\u01cc\u01ce\u01d0\u01d2\u01d4\u01d6\u01d8"+
		"\u01da\u01dc\u01de\2\r\3\2\20\26\3\2.\60\3\2AB\7\2EF\u0081\u0081\u0088"+
		"\u008a\u008f\u0090\u0093\u0093\5\288;;XY\r\2\30\33\36\37%%\'(*\668;@B"+
		"NQSS^ads\13\2\3\24\26\26\30DGGNQSSUW^bd}\5\2\u008b\u008c\u008e\u008f\u0091"+
		"\u009a\4\2xx}}\3\2\u009e\u00a2\3\2\u00a4\u00a5\2\u0996\2\u01e0\3\2\2\2"+
		"\4\u01e6\3\2\2\2\6\u01ee\3\2\2\2\b\u01f2\3\2\2\2\n\u01f4\3\2\2\2\f\u01fa"+
		"\3\2\2\2\16\u0201\3\2\2\2\20\u0203\3\2\2\2\22\u020d\3\2\2\2\24\u020f\3"+
		"\2\2\2\26\u0214\3\2\2\2\30\u021c\3\2\2\2\32\u0226\3\2\2\2\34\u0228\3\2"+
		"\2\2\36\u022c\3\2\2\2 \u0232\3\2\2\2\"\u0234\3\2\2\2$\u0236\3\2\2\2&\u0238"+
		"\3\2\2\2(\u023c\3\2\2\2*\u0240\3\2\2\2,\u0251\3\2\2\2.\u0253\3\2\2\2\60"+
		"\u0256\3\2\2\2\62\u0260\3\2\2\2\64\u026a\3\2\2\2\66\u026c\3\2\2\28\u0270"+
		"\3\2\2\2:\u0274\3\2\2\2<\u027c\3\2\2\2>\u028c\3\2\2\2@\u028f\3\2\2\2B"+
		"\u0293\3\2\2\2D\u0299\3\2\2\2F\u029b\3\2\2\2H\u02a5\3\2\2\2J\u02a7\3\2"+
		"\2\2L\u02aa\3\2\2\2N\u02b2\3\2\2\2P\u02b6\3\2\2\2R\u02dc\3\2\2\2T\u02df"+
		"\3\2\2\2V\u02e6\3\2\2\2X\u02f4\3\2\2\2Z\u02f7\3\2\2\2\\\u0300\3\2\2\2"+
		"^\u030b\3\2\2\2`\u031b\3\2\2\2b\u031e\3\2\2\2d\u0326\3\2\2\2f\u0339\3"+
		"\2\2\2h\u033c\3\2\2\2j\u0345\3\2\2\2l\u034d\3\2\2\2n\u035a\3\2\2\2p\u0363"+
		"\3\2\2\2r\u0371\3\2\2\2t\u0373\3\2\2\2v\u0379\3\2\2\2x\u0381\3\2\2\2z"+
		"\u0383\3\2\2\2|\u03a0\3\2\2\2~\u03a2\3\2\2\2\u0080\u03a4\3\2\2\2\u0082"+
		"\u03a6\3\2\2\2\u0084\u03b7\3\2\2\2\u0086\u03ba\3\2\2\2\u0088\u03cd\3\2"+
		"\2\2\u008a\u03d3\3\2\2\2\u008c\u03d6\3\2\2\2\u008e\u03e3\3\2\2\2\u0090"+
		"\u03e5\3\2\2\2\u0092\u03e9\3\2\2\2\u0094\u03eb\3\2\2\2\u0096\u03ed\3\2"+
		"\2\2\u0098\u03fa\3\2\2\2\u009a\u0400\3\2\2\2\u009c\u0403\3\2\2\2\u009e"+
		"\u040d\3\2\2\2\u00a0\u040f\3\2\2\2\u00a2\u0413\3\2\2\2\u00a4\u0419\3\2"+
		"\2\2\u00a6\u041c\3\2\2\2\u00a8\u042e\3\2\2\2\u00aa\u0430\3\2\2\2\u00ac"+
		"\u0439\3\2\2\2\u00ae\u0468\3\2\2\2\u00b0\u046a\3\2\2\2\u00b2\u046c\3\2"+
		"\2\2\u00b4\u0475\3\2\2\2\u00b6\u0478\3\2\2\2\u00b8\u0484\3\2\2\2\u00ba"+
		"\u0486\3\2\2\2\u00bc\u048f\3\2\2\2\u00be\u0496\3\2\2\2\u00c0\u0498\3\2"+
		"\2\2\u00c2\u049d\3\2\2\2\u00c4\u04ba\3\2\2\2\u00c6\u04bc\3\2\2\2\u00c8"+
		"\u04c1\3\2\2\2\u00ca\u04e2\3\2\2\2\u00cc\u04fb\3\2\2\2\u00ce\u04fd\3\2"+
		"\2\2\u00d0\u0500\3\2\2\2\u00d2\u051d\3\2\2\2\u00d4\u051f\3\2\2\2\u00d6"+
		"\u0528\3\2\2\2\u00d8\u0536\3\2\2\2\u00da\u0539\3\2\2\2\u00dc\u0541\3\2"+
		"\2\2\u00de\u054a\3\2\2\2\u00e0\u054c\3\2\2\2\u00e2\u0550\3\2\2\2\u00e4"+
		"\u0554\3\2\2\2\u00e6\u055a\3\2\2\2\u00e8\u055d\3\2\2\2\u00ea\u056b\3\2"+
		"\2\2\u00ec\u0573\3\2\2\2\u00ee\u0575\3\2\2\2\u00f0\u0579\3\2\2\2\u00f2"+
		"\u057d\3\2\2\2\u00f4\u057f\3\2\2\2\u00f6\u0587\3\2\2\2\u00f8\u05a1\3\2"+
		"\2\2\u00fa\u05a4\3\2\2\2\u00fc\u05c1\3\2\2\2\u00fe\u05c3\3\2\2\2\u0100"+
		"\u05d8\3\2\2\2\u0102\u05e2\3\2\2\2\u0104\u05e4\3\2\2\2\u0106\u05ea\3\2"+
		"\2\2\u0108\u05ec\3\2\2\2\u010a\u05f2\3\2\2\2\u010c\u05fa\3\2\2\2\u010e"+
		"\u05fd\3\2\2\2\u0110\u0604\3\2\2\2\u0112\u0607\3\2\2\2\u0114\u0609\3\2"+
		"\2\2\u0116\u060e\3\2\2\2\u0118\u0610\3\2\2\2\u011a\u0615\3\2\2\2\u011c"+
		"\u061c\3\2\2\2\u011e\u062f\3\2\2\2\u0120\u0637\3\2\2\2\u0122\u063a\3\2"+
		"\2\2\u0124\u0640\3\2\2\2\u0126\u064d\3\2\2\2\u0128\u064f\3\2\2\2\u012a"+
		"\u0657\3\2\2\2\u012c\u0669\3\2\2\2\u012e\u066c\3\2\2\2\u0130\u0670\3\2"+
		"\2\2\u0132\u0681\3\2\2\2\u0134\u0691\3\2\2\2\u0136\u069b\3\2\2\2\u0138"+
		"\u069d\3\2\2\2\u013a\u06ab\3\2\2\2\u013c\u06ad\3\2\2\2\u013e\u06b6\3\2"+
		"\2\2\u0140\u06c0\3\2\2\2\u0142\u06c2\3\2\2\2\u0144\u06e7\3\2\2\2\u0146"+
		"\u06fc\3\2\2\2\u0148\u0701\3\2\2\2\u014a\u0703\3\2\2\2\u014c\u0707\3\2"+
		"\2\2\u014e\u070c\3\2\2\2\u0150\u0710\3\2\2\2\u0152\u0728\3\2\2\2\u0154"+
		"\u0731\3\2\2\2\u0156\u0733\3\2\2\2\u0158\u073b\3\2\2\2\u015a\u074b\3\2"+
		"\2\2\u015c\u074d\3\2\2\2\u015e\u074f\3\2\2\2\u0160\u0753\3\2\2\2\u0162"+
		"\u075c\3\2\2\2\u0164\u0760\3\2\2\2\u0166\u0762\3\2\2\2\u0168\u0765\3\2"+
		"\2\2\u016a\u0775\3\2\2\2\u016c\u077c\3\2\2\2\u016e\u077e\3\2\2\2\u0170"+
		"\u0791\3\2\2\2\u0172\u0793\3\2\2\2\u0174\u079b\3\2\2\2\u0176\u07d8\3\2"+
		"\2\2\u0178\u07da\3\2\2\2\u017a\u07ec\3\2\2\2\u017c\u07ee\3\2\2\2\u017e"+
		"\u07f0\3\2\2\2\u0180\u07f7\3\2\2\2\u0182\u07fa\3\2\2\2\u0184\u080a\3\2"+
		"\2\2\u0186\u081b\3\2\2\2\u0188\u0824\3\2\2\2\u018a\u082c\3\2\2\2\u018c"+
		"\u082e\3\2\2\2\u018e\u0839\3\2\2\2\u0190\u083f\3\2\2\2\u0192\u0841\3\2"+
		"\2\2\u0194\u0855\3\2\2\2\u0196\u0860\3\2\2\2\u0198\u0867\3\2\2\2\u019a"+
		"\u0873\3\2\2\2\u019c\u0875\3\2\2\2\u019e\u0877\3\2\2\2\u01a0\u087b\3\2"+
		"\2\2\u01a2\u0881\3\2\2\2\u01a4\u0888\3\2\2\2\u01a6\u0893\3\2\2\2\u01a8"+
		"\u089a\3\2\2\2\u01aa\u089c\3\2\2\2\u01ac\u08a0\3\2\2\2\u01ae\u08a4\3\2"+
		"\2\2\u01b0\u08a6\3\2\2\2\u01b2\u08a8\3\2\2\2\u01b4\u08aa\3\2\2\2\u01b6"+
		"\u08ad\3\2\2\2\u01b8\u08b0\3\2\2\2\u01ba\u08b4\3\2\2\2\u01bc\u08b8\3\2"+
		"\2\2\u01be\u08bc\3\2\2\2\u01c0\u08c0\3\2\2\2\u01c2\u08c5\3\2\2\2\u01c4"+
		"\u08c9\3\2\2\2\u01c6\u08cc\3\2\2\2\u01c8\u08cf\3\2\2\2\u01ca\u08e2\3\2"+
		"\2\2\u01cc\u08e6\3\2\2\2\u01ce\u08ea\3\2\2\2\u01d0\u08ec\3\2\2\2\u01d2"+
		"\u08f0\3\2\2\2\u01d4\u08f6\3\2\2\2\u01d6\u0900\3\2\2\2\u01d8\u0902\3\2"+
		"\2\2\u01da\u0904\3\2\2\2\u01dc\u0906\3\2\2\2\u01de\u0908\3\2\2\2\u01e0"+
		"\u01e1\5\4\3\2\u01e1\u01e2\7\2\2\3\u01e2\3\3\2\2\2\u01e3\u01e5\5\6\4\2"+
		"\u01e4\u01e3\3\2\2\2\u01e5\u01e8\3\2\2\2\u01e6\u01e4\3\2\2\2\u01e6\u01e7"+
		"\3\2\2\2\u01e7\5\3\2\2\2\u01e8\u01e6\3\2\2\2\u01e9\u01ef\5\u0122\u0092"+
		"\2\u01ea\u01ef\5> \2\u01eb\u01ef\5\b\5\2\u01ec\u01ef\5\16\b\2\u01ed\u01ef"+
		"\5 \21\2\u01ee\u01e9\3\2\2\2\u01ee\u01ea\3\2\2\2\u01ee\u01eb\3\2\2\2\u01ee"+
		"\u01ec\3\2\2\2\u01ee\u01ed\3\2\2\2\u01ef\7\3\2\2\2\u01f0\u01f3\5\n\6\2"+
		"\u01f1\u01f3\5\f\7\2\u01f2\u01f0\3\2\2\2\u01f2\u01f1\3\2\2\2\u01f3\t\3"+
		"\2\2\2\u01f4\u01f5\7\3\2\2\u01f5\u01f6\5\u0100\u0081\2\u01f6\u01f7\7\4"+
		"\2\2\u01f7\u01f8\5\u0122\u0092\2\u01f8\u01f9\5B\"\2\u01f9\13\3\2\2\2\u01fa"+
		"\u01fb\7\5\2\2\u01fb\u01fc\5\u0122\u0092\2\u01fc\u01fd\5B\"\2\u01fd\r"+
		"\3\2\2\2\u01fe\u0202\5\20\t\2\u01ff\u0202\5\24\13\2\u0200\u0202\5\26\f"+
		"\2\u0201\u01fe\3\2\2\2\u0201\u01ff\3\2\2\2\u0201\u0200\3\2\2\2\u0202\17"+
		"\3\2\2\2\u0203\u0204\7\6\2\2\u0204\u0205\5\u0122\u0092\2\u0205\u0207\5"+
		"B\"\2\u0206\u0208\5\22\n\2\u0207\u0206\3\2\2\2\u0207\u0208\3\2\2\2\u0208"+
		"\21\3\2\2\2\u0209\u020a\7\7\2\2\u020a\u020e\5B\"\2\u020b\u020c\7\7\2\2"+
		"\u020c\u020e\5\20\t\2\u020d\u0209\3\2\2\2\u020d\u020b\3\2\2\2\u020e\23"+
		"\3\2\2\2\u020f\u0210\7\b\2\2\u0210\u0211\5\u0122\u0092\2\u0211\u0212\7"+
		"\7\2\2\u0212\u0213\5B\"\2\u0213\25\3\2\2\2\u0214\u0215\7\t\2\2\u0215\u0216"+
		"\5\u0122\u0092\2\u0216\u0218\7\u0082\2\2\u0217\u0219\5\30\r\2\u0218\u0217"+
		"\3\2\2\2\u0218\u0219\3\2\2\2\u0219\u021a\3\2\2\2\u021a\u021b\7\u0085\2"+
		"\2\u021b\27\3\2\2\2\u021c\u021e\5\32\16\2\u021d\u021f\5\30\r\2\u021e\u021d"+
		"\3\2\2\2\u021e\u021f\3\2\2\2\u021f\31\3\2\2\2\u0220\u0221\5\34\17\2\u0221"+
		"\u0222\5\4\3\2\u0222\u0227\3\2\2\2\u0223\u0224\5\36\20\2\u0224\u0225\5"+
		"\4\3\2\u0225\u0227\3\2\2\2\u0226\u0220\3\2\2\2\u0226\u0223\3\2\2\2\u0227"+
		"\33\3\2\2\2\u0228\u0229\7\n\2\2\u0229\u022a\5\u0122\u0092\2\u022a\u022b"+
		"\7\u0089\2\2\u022b\35\3\2\2\2\u022c\u022d\7\13\2\2\u022d\u022e\7\u0089"+
		"\2\2\u022e\37\3\2\2\2\u022f\u0233\5\"\22\2\u0230\u0233\5$\23\2\u0231\u0233"+
		"\5&\24\2\u0232\u022f\3\2\2\2\u0232\u0230\3\2\2\2\u0232\u0231\3\2\2\2\u0233"+
		"!\3\2\2\2\u0234\u0235\7\f\2\2\u0235#\3\2\2\2\u0236\u0237\7\r\2\2\u0237"+
		"%\3\2\2\2\u0238\u023a\7\16\2\2\u0239\u023b\5\u0122\u0092\2\u023a\u0239"+
		"\3\2\2\2\u023a\u023b\3\2\2\2\u023b\'\3\2\2\2\u023c\u023d\7\u008b\2\2\u023d"+
		"\u023e\5*\26\2\u023e\u023f\7\u008c\2\2\u023f)\3\2\2\2\u0240\u0245\5,\27"+
		"\2\u0241\u0242\7\u0088\2\2\u0242\u0244\5,\27\2\u0243\u0241\3\2\2\2\u0244"+
		"\u0247\3\2\2\2\u0245\u0243\3\2\2\2\u0245\u0246\3\2\2\2\u0246+\3\2\2\2"+
		"\u0247\u0245\3\2\2\2\u0248\u0252\5\u018a\u00c6\2\u0249\u024a\5\u018a\u00c6"+
		"\2\u024a\u024b\7\u0089\2\2\u024b\u024c\5\u0188\u00c5\2\u024c\u0252\3\2"+
		"\2\2\u024d\u024e\5\u018a\u00c6\2\u024e\u024f\7\u0089\2\2\u024f\u0250\5"+
		"\u01a2\u00d2\2\u0250\u0252\3\2\2\2\u0251\u0248\3\2\2\2\u0251\u0249\3\2"+
		"\2\2\u0251\u024d\3\2\2\2\u0252-\3\2\2\2\u0253\u0254\7\17\2\2\u0254\u0255"+
		"\5\60\31\2\u0255/\3\2\2\2\u0256\u025b\5\62\32\2\u0257\u0258\7\u0088\2"+
		"\2\u0258\u025a\5\62\32\2\u0259\u0257\3\2\2\2\u025a\u025d\3\2\2\2\u025b"+
		"\u0259\3\2\2\2\u025b\u025c\3\2\2\2\u025c\61\3\2\2\2\u025d\u025b\3\2\2"+
		"\2\u025e\u0261\5\64\33\2\u025f\u0261\5\66\34\2\u0260\u025e\3\2\2\2\u0260"+
		"\u025f\3\2\2\2\u0261\63\3\2\2\2\u0262\u0263\5\u0188\u00c5\2\u0263\u0264"+
		"\7\u0089\2\2\u0264\u0265\5\u0188\u00c5\2\u0265\u026b\3\2\2\2\u0266\u0267"+
		"\5\u0188\u00c5\2\u0267\u0268\7\u0089\2\2\u0268\u0269\5\u01a2\u00d2\2\u0269"+
		"\u026b\3\2\2\2\u026a\u0262\3\2\2\2\u026a\u0266\3\2\2\2\u026b\65\3\2\2"+
		"\2\u026c\u026d\5\u0188\u00c5\2\u026d\u026e\5\u01c2\u00e2\2\u026e\u026f"+
		"\5\u0184\u00c3\2\u026f\67\3\2\2\2\u0270\u0271\7\u008b\2\2\u0271\u0272"+
		"\5:\36\2\u0272\u0273\7\u008c\2\2\u02739\3\2\2\2\u0274\u0279\5<\37\2\u0275"+
		"\u0276\7\u0088\2\2\u0276\u0278\5<\37\2\u0277\u0275\3\2\2\2\u0278\u027b"+
		"\3\2\2\2\u0279\u0277\3\2\2\2\u0279\u027a\3\2\2\2\u027a;\3\2\2\2\u027b"+
		"\u0279\3\2\2\2\u027c\u027d\5\u0184\u00c3\2\u027d=\3\2\2\2\u027e\u028d"+
		"\5J&\2\u027f\u028d\5R*\2\u0280\u028d\5l\67\2\u0281\u028d\5\u0084C\2\u0282"+
		"\u028d\5\u00a6T\2\u0283\u028d\5\u00aeX\2\u0284\u028d\5\u00b6\\\2\u0285"+
		"\u028d\5\u00caf\2\u0286\u028d\5\u00d0i\2\u0287\u028d\5\u00d2j\2\u0288"+
		"\u028d\5\u00d8m\2\u0289\u028d\5\u00dep\2\u028a\u028d\5\u00dep\2\u028b"+
		"\u028d\5\u00e8u\2\u028c\u027e\3\2\2\2\u028c\u027f\3\2\2\2\u028c\u0280"+
		"\3\2\2\2\u028c\u0281\3\2\2\2\u028c\u0282\3\2\2\2\u028c\u0283\3\2\2\2\u028c"+
		"\u0284\3\2\2\2\u028c\u0285\3\2\2\2\u028c\u0286\3\2\2\2\u028c\u0287\3\2"+
		"\2\2\u028c\u0288\3\2\2\2\u028c\u0289\3\2\2\2\u028c\u028a\3\2\2\2\u028c"+
		"\u028b\3\2\2\2\u028d?\3\2\2\2\u028e\u0290\5> \2\u028f\u028e\3\2\2\2\u0290"+
		"\u0291\3\2\2\2\u0291\u028f\3\2\2\2\u0291\u0292\3\2\2\2\u0292A\3\2\2\2"+
		"\u0293\u0295\7\u0082\2\2\u0294\u0296\5\4\3\2\u0295\u0294\3\2\2\2\u0295"+
		"\u0296\3\2\2\2\u0296\u0297\3\2\2\2\u0297\u0298\7\u0085\2\2\u0298C\3\2"+
		"\2\2\u0299\u029a\t\2\2\2\u029aE\3\2\2\2\u029b\u02a0\5H%\2\u029c\u029d"+
		"\7\u0081\2\2\u029d\u029f\5H%\2\u029e\u029c\3\2\2\2\u029f\u02a2\3\2\2\2"+
		"\u02a0\u029e\3\2\2\2\u02a0\u02a1\3\2\2\2\u02a1G\3\2\2\2\u02a2\u02a0\3"+
		"\2\2\2\u02a3\u02a6\5\u01ac\u00d7\2\u02a4\u02a6\5\u01ca\u00e6\2\u02a5\u02a3"+
		"\3\2\2\2\u02a5\u02a4\3\2\2\2\u02a6I\3\2\2\2\u02a7\u02a8\7\27\2\2\u02a8"+
		"\u02a9\5L\'\2\u02a9K\3\2\2\2\u02aa\u02af\5N(\2\u02ab\u02ac\7\u0088\2\2"+
		"\u02ac\u02ae\5N(\2\u02ad\u02ab\3\2\2\2\u02ae\u02b1\3\2\2\2\u02af\u02ad"+
		"\3\2\2\2\u02af\u02b0\3\2\2\2\u02b0M\3\2\2\2\u02b1\u02af\3\2\2\2\u02b2"+
		"\u02b4\5\u0100\u0081\2\u02b3\u02b5\5P)\2\u02b4\u02b3\3\2\2\2\u02b4\u02b5"+
		"\3\2\2\2\u02b5O\3\2\2\2\u02b6\u02b7\5\u01b4\u00db\2\u02b7\u02b8\5\u0122"+
		"\u0092\2\u02b8Q\3\2\2\2\u02b9\u02ba\5T+\2\u02ba\u02bb\5V,\2\u02bb\u02bc"+
		"\5\u0186\u00c4\2\u02bc\u02bd\5B\"\2\u02bd\u02dd\3\2\2\2\u02be\u02bf\5"+
		"T+\2\u02bf\u02c0\5V,\2\u02c0\u02c1\5\u0186\u00c4\2\u02c1\u02c2\5X-\2\u02c2"+
		"\u02dd\3\2\2\2\u02c3\u02c4\5T+\2\u02c4\u02c5\5V,\2\u02c5\u02c6\5\u0186"+
		"\u00c4\2\u02c6\u02c7\5`\61\2\u02c7\u02dd\3\2\2\2\u02c8\u02c9\5T+\2\u02c9"+
		"\u02ca\5V,\2\u02ca\u02cc\5\u0186\u00c4\2\u02cb\u02cd\5P)\2\u02cc\u02cb"+
		"\3\2\2\2\u02cc\u02cd\3\2\2\2\u02cd\u02ce\3\2\2\2\u02ce\u02cf\5f\64\2\u02cf"+
		"\u02dd\3\2\2\2\u02d0\u02d1\5T+\2\u02d1\u02d2\5V,\2\u02d2\u02d3\5\u0186"+
		"\u00c4\2\u02d3\u02d5\5\u0186\u00c4\2\u02d4\u02d6\5P)\2\u02d5\u02d4\3\2"+
		"\2\2\u02d5\u02d6\3\2\2\2\u02d6\u02d7\3\2\2\2\u02d7\u02d8\5f\64\2\u02d8"+
		"\u02dd\3\2\2\2\u02d9\u02da\5T+\2\u02da\u02db\5L\'\2\u02db\u02dd\3\2\2"+
		"\2\u02dc\u02b9\3\2\2\2\u02dc\u02be\3\2\2\2\u02dc\u02c3\3\2\2\2\u02dc\u02c8"+
		"\3\2\2\2\u02dc\u02d0\3\2\2\2\u02dc\u02d9\3\2\2\2\u02ddS\3\2\2\2\u02de"+
		"\u02e0\5\u011a\u008e\2\u02df\u02de\3\2\2\2\u02df\u02e0\3\2\2\2\u02e0\u02e2"+
		"\3\2\2\2\u02e1\u02e3\5\u00fa~\2\u02e2\u02e1\3\2\2\2\u02e2\u02e3\3\2\2"+
		"\2\u02e3\u02e4\3\2\2\2\u02e4\u02e5\7\25\2\2\u02e5U\3\2\2\2\u02e6\u02e7"+
		"\5\u01ac\u00d7\2\u02e7W\3\2\2\2\u02e8\u02e9\7\u0082\2\2\u02e9\u02eb\5"+
		"Z.\2\u02ea\u02ec\5\\/\2\u02eb\u02ea\3\2\2\2\u02eb\u02ec\3\2\2\2\u02ec"+
		"\u02ed\3\2\2\2\u02ed\u02ee\7\u0085\2\2\u02ee\u02f5\3\2\2\2\u02ef\u02f0"+
		"\7\u0082\2\2\u02f0\u02f1\5\\/\2\u02f1\u02f2\5Z.\2\u02f2\u02f3\7\u0085"+
		"\2\2\u02f3\u02f5\3\2\2\2\u02f4\u02e8\3\2\2\2\u02f4\u02ef\3\2\2\2\u02f5"+
		"Y\3\2\2\2\u02f6\u02f8\5\u011a\u008e\2\u02f7\u02f6\3\2\2\2\u02f7\u02f8"+
		"\3\2\2\2\u02f8\u02fa\3\2\2\2\u02f9\u02fb\5\u00fe\u0080\2\u02fa\u02f9\3"+
		"\2\2\2\u02fa\u02fb\3\2\2\2\u02fb\u02fc\3\2\2\2\u02fc\u02fd\7\30\2\2\u02fd"+
		"\u02fe\5B\"\2\u02fe[\3\2\2\2\u02ff\u0301\5\u011a\u008e\2\u0300\u02ff\3"+
		"\2\2\2\u0300\u0301\3\2\2\2\u0301\u0303\3\2\2\2\u0302\u0304\5\u00fe\u0080"+
		"\2\u0303\u0302\3\2\2\2\u0303\u0304\3\2\2\2\u0304\u0305\3\2\2\2\u0305\u0307"+
		"\7\31\2\2\u0306\u0308\5^\60\2\u0307\u0306\3\2\2\2\u0307\u0308\3\2\2\2"+
		"\u0308\u0309\3\2\2\2\u0309\u030a\5B\"\2\u030a]\3\2\2\2\u030b\u030c\7\u0083"+
		"\2\2\u030c\u030d\5\u01ac\u00d7\2\u030d\u030e\7\u0086\2\2\u030e_\3\2\2"+
		"\2\u030f\u0310\7\u0082\2\2\u0310\u0312\5b\62\2\u0311\u0313\5d\63\2\u0312"+
		"\u0311\3\2\2\2\u0312\u0313\3\2\2\2\u0313\u0314\3\2\2\2\u0314\u0315\7\u0085"+
		"\2\2\u0315\u031c\3\2\2\2\u0316\u0317\7\u0082\2\2\u0317\u0318\5d\63\2\u0318"+
		"\u0319\5b\62\2\u0319\u031a\7\u0085\2\2\u031a\u031c\3\2\2\2\u031b\u030f"+
		"\3\2\2\2\u031b\u0316\3\2\2\2\u031ca\3\2\2\2\u031d\u031f\5\u011a\u008e"+
		"\2\u031e\u031d\3\2\2\2\u031e\u031f\3\2\2\2\u031f\u0321\3\2\2\2\u0320\u0322"+
		"\5\u00fe\u0080\2\u0321\u0320\3\2\2\2\u0321\u0322\3\2\2\2\u0322\u0323\3"+
		"\2\2\2\u0323\u0324\7\30\2\2\u0324c\3\2\2\2\u0325\u0327\5\u011a\u008e\2"+
		"\u0326\u0325\3\2\2\2\u0326\u0327\3\2\2\2\u0327\u0329\3\2\2\2\u0328\u032a"+
		"\5\u00fe\u0080\2\u0329\u0328\3\2\2\2\u0329\u032a\3\2\2\2\u032a\u032b\3"+
		"\2\2\2\u032b\u032c\7\31\2\2\u032ce\3\2\2\2\u032d\u032e\7\u0082\2\2\u032e"+
		"\u0330\5h\65\2\u032f\u0331\5j\66\2\u0330\u032f\3\2\2\2\u0330\u0331\3\2"+
		"\2\2\u0331\u0332\3\2\2\2\u0332\u0333\7\u0085\2\2\u0333\u033a\3\2\2\2\u0334"+
		"\u0335\7\u0082\2\2\u0335\u0336\5j\66\2\u0336\u0337\5h\65\2\u0337\u0338"+
		"\7\u0085\2\2\u0338\u033a\3\2\2\2\u0339\u032d\3\2\2\2\u0339\u0334\3\2\2"+
		"\2\u033ag\3\2\2\2\u033b\u033d\5\u011a\u008e\2\u033c\u033b\3\2\2\2\u033c"+
		"\u033d\3\2\2\2\u033d\u033e\3\2\2\2\u033e\u0340\7\32\2\2\u033f\u0341\5"+
		"^\60\2\u0340\u033f\3\2\2\2\u0340\u0341\3\2\2\2\u0341\u0342\3\2\2\2\u0342"+
		"\u0343\5B\"\2\u0343i\3\2\2\2\u0344\u0346\5\u011a\u008e\2\u0345\u0344\3"+
		"\2\2\2\u0345\u0346\3\2\2\2\u0346\u0347\3\2\2\2\u0347\u0349\7\33\2\2\u0348"+
		"\u034a\5^\60\2\u0349\u0348\3\2\2\2\u0349\u034a\3\2\2\2\u034a\u034b\3\2"+
		"\2\2\u034b\u034c\5B\"\2\u034ck\3\2\2\2\u034d\u034e\5n8\2\u034e\u0350\5"+
		"p9\2\u034f\u0351\5(\25\2\u0350\u034f\3\2\2\2\u0350\u0351\3\2\2\2\u0351"+
		"\u0352\3\2\2\2\u0352\u0354\5r:\2\u0353\u0355\5.\30\2\u0354\u0353\3\2\2"+
		"\2\u0354\u0355\3\2\2\2\u0355\u0357\3\2\2\2\u0356\u0358\5v<\2\u0357\u0356"+
		"\3\2\2\2\u0357\u0358\3\2\2\2\u0358m\3\2\2\2\u0359\u035b\5\u011a\u008e"+
		"\2\u035a\u0359\3\2\2\2\u035a\u035b\3\2\2\2\u035b\u035d\3\2\2\2\u035c\u035e"+
		"\5\u00fa~\2\u035d\u035c\3\2\2\2\u035d\u035e\3\2\2\2\u035e\u035f\3\2\2"+
		"\2\u035f\u0360\7\26\2\2\u0360o\3\2\2\2\u0361\u0364\5\u01ac\u00d7\2\u0362"+
		"\u0364\5\u01ca\u00e6\2\u0363\u0361\3\2\2\2\u0363\u0362\3\2\2\2\u0364q"+
		"\3\2\2\2\u0365\u0367\5x=\2\u0366\u0368\7\34\2\2\u0367\u0366\3\2\2\2\u0367"+
		"\u0368\3\2\2\2\u0368\u036a\3\2\2\2\u0369\u036b\5t;\2\u036a\u0369\3\2\2"+
		"\2\u036a\u036b\3\2\2\2\u036b\u0372\3\2\2\2\u036c\u036d\5x=\2\u036d\u036f"+
		"\7\35\2\2\u036e\u0370\5t;\2\u036f\u036e\3\2\2\2\u036f\u0370\3\2\2\2\u0370"+
		"\u0372\3\2\2\2\u0371\u0365\3\2\2\2\u0371\u036c\3\2\2\2\u0372s\3\2\2\2"+
		"\u0373\u0375\5\u01be\u00e0\2\u0374\u0376\5\u011a\u008e\2\u0375\u0374\3"+
		"\2\2\2\u0375\u0376\3\2\2\2\u0376\u0377\3\2\2\2\u0377\u0378\5\u0184\u00c3"+
		"\2\u0378u\3\2\2\2\u0379\u037a\5B\"\2\u037aw\3\2\2\2\u037b\u037c\7\u0083"+
		"\2\2\u037c\u0382\7\u0086\2\2\u037d\u037e\7\u0083\2\2\u037e\u037f\5z>\2"+
		"\u037f\u0380\7\u0086\2\2\u0380\u0382\3\2\2\2\u0381\u037b\3\2\2\2\u0381"+
		"\u037d\3\2\2\2\u0382y\3\2\2\2\u0383\u0388\5|?\2\u0384\u0385\7\u0088\2"+
		"\2\u0385\u0387\5|?\2\u0386\u0384\3\2\2\2\u0387\u038a\3\2\2\2\u0388\u0386"+
		"\3\2\2\2\u0388\u0389\3\2\2\2\u0389{\3\2\2\2\u038a\u0388\3\2\2\2\u038b"+
		"\u038d\5~@\2\u038c\u038b\3\2\2\2\u038c\u038d\3\2\2\2\u038d\u038e\3\2\2"+
		"\2\u038e\u038f\5\u0080A\2\u038f\u0391\5\u0186\u00c4\2\u0390\u0392\5\u0082"+
		"B\2\u0391\u0390\3\2\2\2\u0391\u0392\3\2\2\2\u0392\u03a1\3\2\2\2\u0393"+
		"\u0395\5~@\2\u0394\u0393\3\2\2\2\u0394\u0395\3\2\2\2\u0395\u0396\3\2\2"+
		"\2\u0396\u0397\5\u0080A\2\u0397\u0398\5\u0186\u00c4\2\u0398\u03a1\3\2"+
		"\2\2\u0399\u039b\5~@\2\u039a\u0399\3\2\2\2\u039a\u039b\3\2\2\2\u039b\u039c"+
		"\3\2\2\2\u039c\u039d\5\u0080A\2\u039d\u039e\5\u0186\u00c4\2\u039e\u039f"+
		"\5\u01c0\u00e1\2\u039f\u03a1\3\2\2\2\u03a0\u038c\3\2\2\2\u03a0\u0394\3"+
		"\2\2\2\u03a0\u039a\3\2\2\2\u03a1}\3\2\2\2\u03a2\u03a3\5\u01ae\u00d8\2"+
		"\u03a3\177\3\2\2\2\u03a4\u03a5\5\u01ae\u00d8\2\u03a5\u0081\3\2\2\2\u03a6"+
		"\u03a7\5\u01b4\u00db\2\u03a7\u03a8\5\u0122\u0092\2\u03a8\u0083\3\2\2\2"+
		"\u03a9\u03ab\5\u011a\u008e\2\u03aa\u03a9\3\2\2\2\u03aa\u03ab\3\2\2\2\u03ab"+
		"\u03ad\3\2\2\2\u03ac\u03ae\5\u00fc\177\2\u03ad\u03ac\3\2\2\2\u03ad\u03ae"+
		"\3\2\2\2\u03ae\u03af\3\2\2\2\u03af\u03b8\5\u0086D\2\u03b0\u03b2\5\u011a"+
		"\u008e\2\u03b1\u03b0\3\2\2\2\u03b1\u03b2\3\2\2\2\u03b2\u03b4\3\2\2\2\u03b3"+
		"\u03b5\5\u00fc\177\2\u03b4\u03b3\3\2\2\2\u03b4\u03b5\3\2\2\2\u03b5\u03b6"+
		"\3\2\2\2\u03b6\u03b8\5\u0096L\2\u03b7\u03aa\3\2\2\2\u03b7\u03b1\3\2\2"+
		"\2\u03b8\u0085\3\2\2\2\u03b9\u03bb\7\36\2\2\u03ba\u03b9\3\2\2\2\u03ba"+
		"\u03bb\3\2\2\2\u03bb\u03bc\3\2\2\2\u03bc\u03bd\7\23\2\2\u03bd\u03bf\5"+
		"\u0092J\2\u03be\u03c0\5(\25\2\u03bf\u03be\3\2\2\2\u03bf\u03c0\3\2\2\2"+
		"\u03c0\u03c2\3\2\2\2\u03c1\u03c3\5\u01a6\u00d4\2\u03c2\u03c1\3\2\2\2\u03c2"+
		"\u03c3\3\2\2\2\u03c3\u03c5\3\2\2\2\u03c4\u03c6\5.\30\2\u03c5\u03c4\3\2"+
		"\2\2\u03c5\u03c6\3\2\2\2\u03c6\u03c7\3\2\2\2\u03c7\u03c9\7\u0082\2\2\u03c8"+
		"\u03ca\5\u0088E\2\u03c9\u03c8\3\2\2\2\u03c9\u03ca\3\2\2\2\u03ca\u03cb"+
		"\3\2\2\2\u03cb\u03cc\7\u0085\2\2\u03cc\u0087\3\2\2\2\u03cd\u03cf\5\u008a"+
		"F\2\u03ce\u03d0\5\u0088E\2\u03cf\u03ce\3\2\2\2\u03cf\u03d0\3\2\2\2\u03d0"+
		"\u0089\3\2\2\2\u03d1\u03d4\5> \2\u03d2\u03d4\5\u008cG\2\u03d3\u03d1\3"+
		"\2\2\2\u03d3\u03d2\3\2\2\2\u03d4\u008b\3\2\2\2\u03d5\u03d7\5\u011a\u008e"+
		"\2\u03d6\u03d5\3\2\2\2\u03d6\u03d7\3\2\2\2\u03d7\u03d9\3\2\2\2\u03d8\u03da"+
		"\7\36\2\2\u03d9\u03d8\3\2\2\2\u03d9\u03da\3\2\2\2\u03da\u03db\3\2\2\2"+
		"\u03db\u03dc\7\n\2\2\u03dc\u03dd\5\u008eH\2\u03dd\u008d\3\2\2\2\u03de"+
		"\u03e4\5\u0090I\2\u03df\u03e0\5\u0090I\2\u03e0\u03e1\7\u0088\2\2\u03e1"+
		"\u03e2\5\u008eH\2\u03e2\u03e4\3\2\2\2\u03e3\u03de\3\2\2\2\u03e3\u03df"+
		"\3\2\2\2\u03e4\u008f\3\2\2\2\u03e5\u03e7\5\u0094K\2\u03e6\u03e8\5\u018c"+
		"\u00c7\2\u03e7\u03e6\3\2\2\2\u03e7\u03e8\3\2\2\2\u03e8\u0091\3\2\2\2\u03e9"+
		"\u03ea\5\u01ac\u00d7\2\u03ea\u0093\3\2\2\2\u03eb\u03ec\5\u01ac\u00d7\2"+
		"\u03ec\u0095\3\2\2\2\u03ed\u03ee\7\23\2\2\u03ee\u03f0\5\u0092J\2\u03ef"+
		"\u03f1\5(\25\2\u03f0\u03ef\3\2\2\2\u03f0\u03f1\3\2\2\2\u03f1\u03f2\3\2"+
		"\2\2\u03f2\u03f4\5\u01a6\u00d4\2\u03f3\u03f5\5.\30\2\u03f4\u03f3\3\2\2"+
		"\2\u03f4\u03f5\3\2\2\2\u03f5\u03f6\3\2\2\2\u03f6\u03f7\7\u0082\2\2\u03f7"+
		"\u03f8\5\u0098M\2\u03f8\u03f9\7\u0085\2\2\u03f9\u0097\3\2\2\2\u03fa\u03fc"+
		"\5\u009aN\2\u03fb\u03fd\5\u0098M\2\u03fc\u03fb\3\2\2\2\u03fc\u03fd\3\2"+
		"\2\2\u03fd\u0099\3\2\2\2\u03fe\u0401\5> \2\u03ff\u0401\5\u009cO\2\u0400"+
		"\u03fe\3\2\2\2\u0400\u03ff\3\2\2\2\u0401\u009b\3\2\2\2\u0402\u0404\5\u011a"+
		"\u008e\2\u0403\u0402\3\2\2\2\u0403\u0404\3\2\2\2\u0404\u0405\3\2\2\2\u0405"+
		"\u0406\7\n\2\2\u0406\u0407\5\u009eP\2\u0407\u009d\3\2\2\2\u0408\u040e"+
		"\5\u00a0Q\2\u0409\u040a\5\u00a0Q\2\u040a\u040b\7\u0088\2\2\u040b\u040c"+
		"\5\u009eP\2\u040c\u040e\3\2\2\2\u040d\u0408\3\2\2\2\u040d\u0409\3\2\2"+
		"\2\u040e\u009f\3\2\2\2\u040f\u0411\5\u0094K\2\u0410\u0412\5\u00a2R\2\u0411"+
		"\u0410\3\2\2\2\u0411\u0412\3\2\2\2\u0412\u00a1\3\2\2\2\u0413\u0414\5\u01b4"+
		"\u00db\2\u0414\u0415\5\u00a4S\2\u0415\u00a3\3\2\2\2\u0416\u041a\5\u01d6"+
		"\u00ec\2\u0417\u041a\7\u00a4\2\2\u0418\u041a\5\u01d8\u00ed\2\u0419\u0416"+
		"\3\2\2\2\u0419\u0417\3\2\2\2\u0419\u0418\3\2\2\2\u041a\u00a5\3\2\2\2\u041b"+
		"\u041d\5\u011a\u008e\2\u041c\u041b\3\2\2\2\u041c\u041d\3\2\2\2\u041d\u041f"+
		"\3\2\2\2\u041e\u0420\5\u00fc\177\2\u041f\u041e\3\2\2\2\u041f\u0420\3\2"+
		"\2\2\u0420\u0421\3\2\2\2\u0421\u0422\7\21\2\2\u0422\u0424\5\u00a8U\2\u0423"+
		"\u0425\5(\25\2\u0424\u0423\3\2\2\2\u0424\u0425\3\2\2\2\u0425\u0427\3\2"+
		"\2\2\u0426\u0428\5\u01a6\u00d4\2\u0427\u0426\3\2\2\2\u0427\u0428\3\2\2"+
		"\2\u0428\u042a\3\2\2\2\u0429\u042b\5.\30\2\u042a\u0429\3\2\2\2\u042a\u042b"+
		"\3\2\2\2\u042b\u042c\3\2\2\2\u042c\u042d\5\u00aaV\2\u042d\u00a7\3\2\2"+
		"\2\u042e\u042f\5\u01ac\u00d7\2\u042f\u00a9\3\2\2\2\u0430\u0434\7\u0082"+
		"\2\2\u0431\u0433\5\u00acW\2\u0432\u0431\3\2\2\2\u0433\u0436\3\2\2\2\u0434"+
		"\u0432\3\2\2\2\u0434\u0435\3\2\2\2\u0435\u0437\3\2\2\2\u0436\u0434\3\2"+
		"\2\2\u0437\u0438\7\u0085\2\2\u0438\u00ab\3\2\2\2\u0439\u043a\5> \2\u043a"+
		"\u00ad\3\2\2\2\u043b\u043d\5\u011a\u008e\2\u043c\u043b\3\2\2\2\u043c\u043d"+
		"\3\2\2\2\u043d\u043f\3\2\2\2\u043e\u0440\5\u00fc\177\2\u043f\u043e\3\2"+
		"\2\2\u043f\u0440\3\2\2\2\u0440\u0442\3\2\2\2\u0441\u0443\7\37\2\2\u0442"+
		"\u0441\3\2\2\2\u0442\u0443\3\2\2\2\u0443\u0444\3\2\2\2\u0444\u0445\7\22"+
		"\2\2\u0445\u0447\5\u00b0Y\2\u0446\u0448\5(\25\2\u0447\u0446\3\2\2\2\u0447"+
		"\u0448\3\2\2\2\u0448\u044a\3\2\2\2\u0449\u044b\5\u01a6\u00d4\2\u044a\u0449"+
		"\3\2\2\2\u044a\u044b\3\2\2\2\u044b\u044d\3\2\2\2\u044c\u044e\5.\30\2\u044d"+
		"\u044c\3\2\2\2\u044d\u044e\3\2\2\2\u044e\u044f\3\2\2\2\u044f\u0450\5\u00b2"+
		"Z\2\u0450\u0469\3\2\2\2\u0451\u0453\5\u011a\u008e\2\u0452\u0451\3\2\2"+
		"\2\u0452\u0453\3\2\2\2\u0453\u0455\3\2\2\2\u0454\u0456\5\u00fc\177\2\u0455"+
		"\u0454\3\2\2\2\u0455\u0456\3\2\2\2\u0456\u0457\3\2\2\2\u0457\u0459\7\37"+
		"\2\2\u0458\u045a\5\u00fc\177\2\u0459\u0458\3\2\2\2\u0459\u045a\3\2\2\2"+
		"\u045a\u045b\3\2\2\2\u045b\u045c\7\22\2\2\u045c\u045e\5\u00b0Y\2\u045d"+
		"\u045f\5(\25\2\u045e\u045d\3\2\2\2\u045e\u045f\3\2\2\2\u045f\u0461\3\2"+
		"\2\2\u0460\u0462\5\u01a6\u00d4\2\u0461\u0460\3\2\2\2\u0461\u0462\3\2\2"+
		"\2\u0462\u0464\3\2\2\2\u0463\u0465\5.\30\2\u0464\u0463\3\2\2\2\u0464\u0465"+
		"\3\2\2\2\u0465\u0466\3\2\2\2\u0466\u0467\5\u00b2Z\2\u0467\u0469\3\2\2"+
		"\2\u0468\u043c\3\2\2\2\u0468\u0452\3\2\2\2\u0469\u00af\3\2\2\2\u046a\u046b"+
		"\5\u01ac\u00d7\2\u046b\u00b1\3\2\2\2\u046c\u0470\7\u0082\2\2\u046d\u046f"+
		"\5\u00b4[\2\u046e\u046d\3\2\2\2\u046f\u0472\3\2\2\2\u0470\u046e\3\2\2"+
		"\2\u0470\u0471\3\2\2\2\u0471\u0473\3\2\2\2\u0472\u0470\3\2\2\2\u0473\u0474"+
		"\7\u0085\2\2\u0474\u00b3\3\2\2\2\u0475\u0476\5> \2\u0476\u00b5\3\2\2\2"+
		"\u0477\u0479\5\u011a\u008e\2\u0478\u0477\3\2\2\2\u0478\u0479\3\2\2\2\u0479"+
		"\u047b\3\2\2\2\u047a\u047c\5\u00fc\177\2\u047b\u047a\3\2\2\2\u047b\u047c"+
		"\3\2\2\2\u047c\u047d\3\2\2\2\u047d\u047e\7\24\2\2\u047e\u0480\5\u00b8"+
		"]\2\u047f\u0481\5\u01a6\u00d4\2\u0480\u047f\3\2\2\2\u0480\u0481\3\2\2"+
		"\2\u0481\u0482\3\2\2\2\u0482\u0483\5\u00ba^\2\u0483\u00b7\3\2\2\2\u0484"+
		"\u0485\5\u01ac\u00d7\2\u0485\u00b9\3\2\2\2\u0486\u048a\7\u0082\2\2\u0487"+
		"\u0489\5\u00bc_\2\u0488\u0487\3\2\2\2\u0489\u048c\3\2\2\2\u048a\u0488"+
		"\3\2\2\2\u048a\u048b\3\2\2\2\u048b\u048d\3\2\2\2\u048c\u048a\3\2\2\2\u048d"+
		"\u048e\7\u0085\2\2\u048e\u00bb\3\2\2\2\u048f\u0490\5\u00be`\2\u0490\u00bd"+
		"\3\2\2\2\u0491\u0497\5\u00c0a\2\u0492\u0497\5\u00c2b\2\u0493\u0497\5\u00c4"+
		"c\2\u0494\u0497\5\u00c6d\2\u0495\u0497\5\u00c8e\2\u0496\u0491\3\2\2\2"+
		"\u0496\u0492\3\2\2\2\u0496\u0493\3\2\2\2\u0496\u0494\3\2\2\2\u0496\u0495"+
		"\3\2\2\2\u0497\u00bf\3\2\2\2\u0498\u0499\5T+\2\u0499\u049a\5V,\2\u049a"+
		"\u049b\5\u0186\u00c4\2\u049b\u049c\5`\61\2\u049c\u00c1\3\2\2\2\u049d\u049e"+
		"\5n8\2\u049e\u04a0\5p9\2\u049f\u04a1\5(\25\2\u04a0\u049f\3\2\2\2\u04a0"+
		"\u04a1\3\2\2\2\u04a1\u04a2\3\2\2\2\u04a2\u04a4\5r:\2\u04a3\u04a5\5.\30"+
		"\2\u04a4\u04a3\3\2\2\2\u04a4\u04a5\3\2\2\2\u04a5\u00c3\3\2\2\2\u04a6\u04a8"+
		"\5\u00ccg\2\u04a7\u04a9\5(\25\2\u04a8\u04a7\3\2\2\2\u04a8\u04a9\3\2\2"+
		"\2\u04a9\u04aa\3\2\2\2\u04aa\u04ac\5x=\2\u04ab\u04ad\7\34\2\2\u04ac\u04ab"+
		"\3\2\2\2\u04ac\u04ad\3\2\2\2\u04ad\u04af\3\2\2\2\u04ae\u04b0\5.\30\2\u04af"+
		"\u04ae\3\2\2\2\u04af\u04b0\3\2\2\2\u04b0\u04bb\3\2\2\2\u04b1\u04b3\5\u00cc"+
		"g\2\u04b2\u04b4\5(\25\2\u04b3\u04b2\3\2\2\2\u04b3\u04b4\3\2\2\2\u04b4"+
		"\u04b5\3\2\2\2\u04b5\u04b6\5x=\2\u04b6\u04b8\7\35\2\2\u04b7\u04b9\5.\30"+
		"\2\u04b8\u04b7\3\2\2\2\u04b8\u04b9\3\2\2\2\u04b9\u04bb\3\2\2\2\u04ba\u04a6"+
		"\3\2\2\2\u04ba\u04b1\3\2\2\2\u04bb\u00c5\3\2\2\2\u04bc\u04bd\5\u00dan"+
		"\2\u04bd\u04be\5\u00dco\2\u04be\u04bf\5`\61\2\u04bf\u00c7\3\2\2\2\u04c0"+
		"\u04c2\5\u011a\u008e\2\u04c1\u04c0\3\2\2\2\u04c1\u04c2\3\2\2\2\u04c2\u04c4"+
		"\3\2\2\2\u04c3\u04c5\5\u00fc\177\2\u04c4\u04c3\3\2\2\2\u04c4\u04c5\3\2"+
		"\2\2\u04c5\u04c6\3\2\2\2\u04c6\u04c8\7 \2\2\u04c7\u04c9\5\u01a6\u00d4"+
		"\2\u04c8\u04c7\3\2\2\2\u04c8\u04c9\3\2\2\2\u04c9\u00c9\3\2\2\2\u04ca\u04cc"+
		"\5\u00ccg\2\u04cb\u04cd\5(\25\2\u04cc\u04cb\3\2\2\2\u04cc\u04cd\3\2\2"+
		"\2\u04cd\u04ce\3\2\2\2\u04ce\u04d0\5x=\2\u04cf\u04d1\7\34\2\2\u04d0\u04cf"+
		"\3\2\2\2\u04d0\u04d1\3\2\2\2\u04d1\u04d3\3\2\2\2\u04d2\u04d4\5.\30\2\u04d3"+
		"\u04d2\3\2\2\2\u04d3\u04d4\3\2\2\2\u04d4\u04d5\3\2\2\2\u04d5\u04d6\5\u00ce"+
		"h\2\u04d6\u04e3\3\2\2\2\u04d7\u04d9\5\u00ccg\2\u04d8\u04da\5(\25\2\u04d9"+
		"\u04d8\3\2\2\2\u04d9\u04da\3\2\2\2\u04da\u04db\3\2\2\2\u04db\u04dc\5x"+
		"=\2\u04dc\u04de\7\35\2\2\u04dd\u04df\5.\30\2\u04de\u04dd\3\2\2\2\u04de"+
		"\u04df\3\2\2\2\u04df\u04e0\3\2\2\2\u04e0\u04e1\5\u00ceh\2\u04e1\u04e3"+
		"\3\2\2\2\u04e2\u04ca\3\2\2\2\u04e2\u04d7\3\2\2\2\u04e3\u00cb\3\2\2\2\u04e4"+
		"\u04e6\5\u011a\u008e\2\u04e5\u04e4\3\2\2\2\u04e5\u04e6\3\2\2\2\u04e6\u04e8"+
		"\3\2\2\2\u04e7\u04e9\5\u00fa~\2\u04e8\u04e7\3\2\2\2\u04e8\u04e9\3\2\2"+
		"\2\u04e9\u04ea\3\2\2\2\u04ea\u04fc\7!\2\2\u04eb\u04ed\5\u011a\u008e\2"+
		"\u04ec\u04eb\3\2\2\2\u04ec\u04ed\3\2\2\2\u04ed\u04ef\3\2\2\2\u04ee\u04f0"+
		"\5\u00fa~\2\u04ef\u04ee\3\2\2\2\u04ef\u04f0\3\2\2\2\u04f0\u04f1\3\2\2"+
		"\2\u04f1\u04f2\7!\2\2\u04f2\u04fc\7\u008f\2\2\u04f3\u04f5\5\u011a\u008e"+
		"\2\u04f4\u04f3\3\2\2\2\u04f4\u04f5\3\2\2\2\u04f5\u04f7\3\2\2\2\u04f6\u04f8"+
		"\5\u00fa~\2\u04f7\u04f6\3\2\2\2\u04f7\u04f8\3\2\2\2\u04f8\u04f9\3\2\2"+
		"\2\u04f9\u04fa\7!\2\2\u04fa\u04fc\7\u008e\2\2\u04fb\u04e5\3\2\2\2\u04fb"+
		"\u04ec\3\2\2\2\u04fb\u04f4\3\2\2\2\u04fc\u00cd\3\2\2\2\u04fd\u04fe\5B"+
		"\"\2\u04fe\u00cf\3\2\2\2\u04ff\u0501\5\u011a\u008e\2\u0500\u04ff\3\2\2"+
		"\2\u0500\u0501\3\2\2\2\u0501\u0502\3\2\2\2\u0502\u0503\7\"\2\2\u0503\u0504"+
		"\5B\"\2\u0504\u00d1\3\2\2\2\u0505\u0507\5\u011a\u008e\2\u0506\u0505\3"+
		"\2\2\2\u0506\u0507\3\2\2\2\u0507\u0509\3\2\2\2\u0508\u050a\5\u00fc\177"+
		"\2\u0509\u0508\3\2\2\2\u0509\u050a\3\2\2\2\u050a\u050b\3\2\2\2\u050b\u050c"+
		"\7#\2\2\u050c\u050e\5\u0188\u00c5\2\u050d\u050f\5\u01a6\u00d4\2\u050e"+
		"\u050d\3\2\2\2\u050e\u050f\3\2\2\2\u050f\u0510\3\2\2\2\u0510\u0511\5\u00d4"+
		"k\2\u0511\u051e\3\2\2\2\u0512\u0514\5\u011a\u008e\2\u0513\u0512\3\2\2"+
		"\2\u0513\u0514\3\2\2\2\u0514\u0516\3\2\2\2\u0515\u0517\5\u00fc\177\2\u0516"+
		"\u0515\3\2\2\2\u0516\u0517\3\2\2\2\u0517\u0518\3\2\2\2\u0518\u0519\7#"+
		"\2\2\u0519\u051a\5\u0188\u00c5\2\u051a\u051b\5.\30\2\u051b\u051c\5\u00d4"+
		"k\2\u051c\u051e\3\2\2\2\u051d\u0506\3\2\2\2\u051d\u0513\3\2\2\2\u051e"+
		"\u00d3\3\2\2\2\u051f\u0523\7\u0082\2\2\u0520\u0522\5\u00d6l\2\u0521\u0520"+
		"\3\2\2\2\u0522\u0525\3\2\2\2\u0523\u0521\3\2\2\2\u0523\u0524\3\2\2\2\u0524"+
		"\u0526\3\2\2\2\u0525\u0523\3\2\2\2\u0526\u0527\7\u0085\2\2\u0527\u00d5"+
		"\3\2\2\2\u0528\u0529\5> \2\u0529\u00d7\3\2\2\2\u052a\u052b\5\u00dan\2"+
		"\u052b\u052c\5\u00dco\2\u052c\u052d\5B\"\2\u052d\u0537\3\2\2\2\u052e\u052f"+
		"\5\u00dan\2\u052f\u0530\5\u00dco\2\u0530\u0531\5X-\2\u0531\u0537\3\2\2"+
		"\2\u0532\u0533\5\u00dan\2\u0533\u0534\5\u00dco\2\u0534\u0535\5`\61\2\u0535"+
		"\u0537\3\2\2\2\u0536\u052a\3\2\2\2\u0536\u052e\3\2\2\2\u0536\u0532\3\2"+
		"\2\2\u0537\u00d9\3\2\2\2\u0538\u053a\5\u011a\u008e\2\u0539\u0538\3\2\2"+
		"\2\u0539\u053a\3\2\2\2\u053a\u053c\3\2\2\2\u053b\u053d\5\u00fa~\2\u053c"+
		"\u053b\3\2\2\2\u053c\u053d\3\2\2\2\u053d\u053e\3\2\2\2\u053e\u053f\7$"+
		"\2\2\u053f\u0540\5x=\2\u0540\u00db\3\2\2\2\u0541\u0543\5\u01be\u00e0\2"+
		"\u0542\u0544\5\u011a\u008e\2\u0543\u0542\3\2\2\2\u0543\u0544\3\2\2\2\u0544"+
		"\u0545\3\2\2\2\u0545\u0546\5\u0184\u00c3\2\u0546\u00dd\3\2\2\2\u0547\u054b"+
		"\5\u00e0q\2\u0548\u054b\5\u00e2r\2\u0549\u054b\5\u00e4s\2\u054a\u0547"+
		"\3\2\2\2\u054a\u0548\3\2\2\2\u054a\u0549\3\2\2\2\u054b\u00df\3\2\2\2\u054c"+
		"\u054d\7%\2\2\u054d\u054e\7&\2\2\u054e\u054f\5\u01ca\u00e6\2\u054f\u00e1"+
		"\3\2\2\2\u0550\u0551\7\'\2\2\u0551\u0552\7&\2\2\u0552\u0553\5\u01ca\u00e6"+
		"\2\u0553\u00e3\3\2\2\2\u0554\u0555\7(\2\2\u0555\u0556\7&\2\2\u0556\u0558"+
		"\5\u01ca\u00e6\2\u0557\u0559\5\u00e6t\2\u0558\u0557\3\2\2\2\u0558\u0559"+
		"\3\2\2\2\u0559\u00e5\3\2\2\2\u055a\u055b\7\u0089\2\2\u055b\u055c\5\u00f6"+
		"|\2\u055c\u00e7\3\2\2\2\u055d\u055e\7)\2\2\u055e\u055f\5\u00f6|\2\u055f"+
		"\u0563\7\u0082\2\2\u0560\u0562\5\u00eav\2\u0561\u0560\3\2\2\2\u0562\u0565"+
		"\3\2\2\2\u0563\u0561\3\2\2\2\u0563\u0564\3\2\2\2\u0564\u0566\3\2\2\2\u0565"+
		"\u0563\3\2\2\2\u0566\u0567\7\u0085\2\2\u0567\u00e9\3\2\2\2\u0568\u056c"+
		"\5\u00ecw\2\u0569\u056c\5\u00eex\2\u056a\u056c\5\u00f0y\2\u056b\u0568"+
		"\3\2\2\2\u056b\u0569\3\2\2\2\u056b\u056a\3\2\2\2\u056c\u00eb\3\2\2\2\u056d"+
		"\u056e\7*\2\2\u056e\u056f\7\u0089\2\2\u056f\u0574\5\u00f4{\2\u0570\u0571"+
		"\7+\2\2\u0571\u0572\7\u0089\2\2\u0572\u0574\5\u00f4{\2\u0573\u056d\3\2"+
		"\2\2\u0573\u0570\3\2\2\2\u0574\u00ed\3\2\2\2\u0575\u0576\7,\2\2\u0576"+
		"\u0577\7\u0089\2\2\u0577\u0578\5\u01d8\u00ed\2\u0578\u00ef\3\2\2\2\u0579"+
		"\u057a\7-\2\2\u057a\u057b\7\u0089\2\2\u057b\u057c\5\u00f2z\2\u057c\u00f1"+
		"\3\2\2\2\u057d\u057e\t\3\2\2\u057e\u00f3\3\2\2\2\u057f\u0584\5\u00f6|"+
		"\2\u0580\u0581\7\u0088\2\2\u0581\u0583\5\u00f6|\2\u0582\u0580\3\2\2\2"+
		"\u0583\u0586\3\2\2\2\u0584\u0582\3\2\2\2\u0584\u0585\3\2\2\2\u0585\u00f5"+
		"\3\2\2\2\u0586\u0584\3\2\2\2\u0587\u0588\5\u01ac\u00d7\2\u0588\u00f7\3"+
		"\2\2\2\u0589\u05a2\7\22\2\2\u058a\u05a2\7\61\2\2\u058b\u05a2\7\62\2\2"+
		"\u058c\u05a2\7\37\2\2\u058d\u05a2\7(\2\2\u058e\u05a2\7\63\2\2\u058f\u05a2"+
		"\7\64\2\2\u0590\u05a2\7\65\2\2\u0591\u05a2\7\'\2\2\u0592\u05a2\7%\2\2"+
		"\u0593\u05a2\7\66\2\2\u0594\u05a2\7\67\2\2\u0595\u05a2\78\2\2\u0596\u0597"+
		"\78\2\2\u0597\u0598\7\u0083\2\2\u0598\u0599\79\2\2\u0599\u05a2\7\u0086"+
		"\2\2\u059a\u059b\78\2\2\u059b\u059c\7\u0083\2\2\u059c\u059d\7:\2\2\u059d"+
		"\u05a2\7\u0086\2\2\u059e\u05a2\7;\2\2\u059f\u05a2\5\u00fc\177\2\u05a0"+
		"\u05a2\5\u00fe\u0080\2\u05a1\u0589\3\2\2\2\u05a1\u058a\3\2\2\2\u05a1\u058b"+
		"\3\2\2\2\u05a1\u058c\3\2\2\2\u05a1\u058d\3\2\2\2\u05a1\u058e\3\2\2\2\u05a1"+
		"\u058f\3\2\2\2\u05a1\u0590\3\2\2\2\u05a1\u0591\3\2\2\2\u05a1\u0592\3\2"+
		"\2\2\u05a1\u0593\3\2\2\2\u05a1\u0594\3\2\2\2\u05a1\u0595\3\2\2\2\u05a1"+
		"\u0596\3\2\2\2\u05a1\u059a\3\2\2\2\u05a1\u059e\3\2\2\2\u05a1\u059f\3\2"+
		"\2\2\u05a1\u05a0\3\2\2\2\u05a2\u00f9\3\2\2\2\u05a3\u05a5\5\u00f8}\2\u05a4"+
		"\u05a3\3\2\2\2\u05a5\u05a6\3\2\2\2\u05a6\u05a4\3\2\2\2\u05a6\u05a7\3\2"+
		"\2\2\u05a7\u00fb\3\2\2\2\u05a8\u05c2\7<\2\2\u05a9\u05aa\7<\2\2\u05aa\u05ab"+
		"\7\u0083\2\2\u05ab\u05ac\7\31\2\2\u05ac\u05c2\7\u0086\2\2\u05ad\u05c2"+
		"\7=\2\2\u05ae\u05af\7=\2\2\u05af\u05b0\7\u0083\2\2\u05b0\u05b1\7\31\2"+
		"\2\u05b1\u05c2\7\u0086\2\2\u05b2\u05c2\7>\2\2\u05b3\u05b4\7>\2\2\u05b4"+
		"\u05b5\7\u0083\2\2\u05b5\u05b6\7\31\2\2\u05b6\u05c2\7\u0086\2\2\u05b7"+
		"\u05c2\7?\2\2\u05b8\u05b9\7?\2\2\u05b9\u05ba\7\u0083\2\2\u05ba\u05bb\7"+
		"\31\2\2\u05bb\u05c2\7\u0086\2\2\u05bc\u05c2\7@\2\2\u05bd\u05be\7@\2\2"+
		"\u05be\u05bf\7\u0083\2\2\u05bf\u05c0\7\31\2\2\u05c0\u05c2\7\u0086\2\2"+
		"\u05c1\u05a8\3\2\2\2\u05c1\u05a9\3\2\2\2\u05c1\u05ad\3\2\2\2\u05c1\u05ae"+
		"\3\2\2\2\u05c1\u05b2\3\2\2\2\u05c1\u05b3\3\2\2\2\u05c1\u05b7\3\2\2\2\u05c1"+
		"\u05b8\3\2\2\2\u05c1\u05bc\3\2\2\2\u05c1\u05bd\3\2\2\2\u05c2\u00fd\3\2"+
		"\2\2\u05c3\u05c4\t\4\2\2\u05c4\u00ff\3\2\2\2\u05c5\u05c6\b\u0081\1\2\u05c6"+
		"\u05c8\5\u0102\u0082\2\u05c7\u05c9\5\u0186\u00c4\2\u05c8\u05c7\3\2\2\2"+
		"\u05c8\u05c9\3\2\2\2\u05c9\u05d9\3\2\2\2\u05ca\u05cc\5\u0104\u0083\2\u05cb"+
		"\u05cd\5\u0186\u00c4\2\u05cc\u05cb\3\2\2\2\u05cc\u05cd\3\2\2\2\u05cd\u05d9"+
		"\3\2\2\2\u05ce\u05d9\5\u0106\u0084\2\u05cf\u05d1\5\u0108\u0085\2\u05d0"+
		"\u05d2\5\u0186\u00c4\2\u05d1\u05d0\3\2\2\2\u05d1\u05d2\3\2\2\2\u05d2\u05d9"+
		"\3\2\2\2\u05d3\u05d9\5\u010e\u0088\2\u05d4\u05d9\5\u0110\u0089\2\u05d5"+
		"\u05d6\7C\2\2\u05d6\u05d9\5\u0184\u00c3\2\u05d7\u05d9\5\u0112\u008a\2"+
		"\u05d8\u05c5\3\2\2\2\u05d8\u05ca\3\2\2\2\u05d8\u05ce\3\2\2\2\u05d8\u05cf"+
		"\3\2\2\2\u05d8\u05d3\3\2\2\2\u05d8\u05d4\3\2\2\2\u05d8\u05d5\3\2\2\2\u05d8"+
		"\u05d7\3\2\2\2\u05d9\u05df\3\2\2\2\u05da\u05db\f\4\2\2\u05db\u05dc\7D"+
		"\2\2\u05dc\u05de\5\u0184\u00c3\2\u05dd\u05da\3\2\2\2\u05de\u05e1\3\2\2"+
		"\2\u05df\u05dd\3\2\2\2\u05df\u05e0\3\2\2\2\u05e0\u0101\3\2\2\2\u05e1\u05df"+
		"\3\2\2\2\u05e2\u05e3\7\u008d\2\2\u05e3\u0103\3\2\2\2\u05e4\u05e5\5\u01ac"+
		"\u00d7\2\u05e5\u0105\3\2\2\2\u05e6\u05e7\7\25\2\2\u05e7\u05eb\5\u0100"+
		"\u0081\2\u05e8\u05e9\7\27\2\2\u05e9\u05eb\5\u0100\u0081\2\u05ea\u05e6"+
		"\3\2\2\2\u05ea\u05e8\3\2\2\2\u05eb\u0107\3\2\2\2\u05ec\u05ee\7\u0083\2"+
		"\2\u05ed\u05ef\5\u010a\u0086\2\u05ee\u05ed\3\2\2\2\u05ee\u05ef\3\2\2\2"+
		"\u05ef\u05f0\3\2\2\2\u05f0\u05f1\7\u0086\2\2\u05f1\u0109\3\2\2\2\u05f2"+
		"\u05f7\5\u010c\u0087\2\u05f3\u05f4\7\u0088\2\2\u05f4\u05f6\5\u010c\u0087"+
		"\2\u05f5\u05f3\3\2\2\2\u05f6\u05f9\3\2\2\2\u05f7\u05f5\3\2\2\2\u05f7\u05f8"+
		"\3\2\2\2\u05f8\u010b\3\2\2\2\u05f9\u05f7\3\2\2\2\u05fa\u05fb\5\u0100\u0081"+
		"\2\u05fb\u010d\3\2\2\2\u05fc\u05fe\5\u0188\u00c5\2\u05fd\u05fc\3\2\2\2"+
		"\u05fd\u05fe\3\2\2\2\u05fe\u05ff\3\2\2\2\u05ff\u0600\7\u0081\2\2\u0600"+
		"\u0602\5\u0094K\2\u0601\u0603\5\u0108\u0085\2\u0602\u0601\3\2\2\2\u0602"+
		"\u0603\3\2\2\2\u0603\u010f\3\2\2\2\u0604\u0605\5\u0104\u0083\2\u0605\u0606"+
		"\7\u008f\2\2\u0606\u0111\3\2\2\2\u0607\u0608\5\u0122\u0092\2\u0608\u0113"+
		"\3\2\2\2\u0609\u060a\7\u0090\2\2\u060a\u060c\5\u0116\u008c\2\u060b\u060d"+
		"\5\u0118\u008d\2\u060c\u060b\3\2\2\2\u060c\u060d\3\2\2\2\u060d\u0115\3"+
		"\2\2\2\u060e\u060f\5\u01ac\u00d7\2\u060f\u0117\3\2\2\2\u0610\u0611\7\u0083"+
		"\2\2\u0611\u0612\5\u011c\u008f\2\u0612\u0613\7\u0086\2\2\u0613\u0119\3"+
		"\2\2\2\u0614\u0616\5\u0114\u008b\2\u0615\u0614\3\2\2\2\u0616\u0617\3\2"+
		"\2\2\u0617\u0615\3\2\2\2\u0617\u0618\3\2\2\2\u0618\u011b\3\2\2\2\u0619"+
		"\u061b\5\u011e\u0090\2\u061a\u0619\3\2\2\2\u061b\u061e\3\2\2\2\u061c\u061a"+
		"\3\2\2\2\u061c\u061d\3\2\2\2\u061d\u011d\3\2\2\2\u061e\u061c\3\2\2\2\u061f"+
		"\u0620\7\u0083\2\2\u0620\u0621\5\u011c\u008f\2\u0621\u0622\7\u0086\2\2"+
		"\u0622\u0630\3\2\2\2\u0623\u0624\7\u0084\2\2\u0624\u0625\5\u011c\u008f"+
		"\2\u0625\u0626\7\u0087\2\2\u0626\u0630\3\2\2\2\u0627\u0628\7\u0082\2\2"+
		"\u0628\u0629\5\u011c\u008f\2\u0629\u062a\7\u0085\2\2\u062a\u0630\3\2\2"+
		"\2\u062b\u0630\5\u01ae\u00d8\2\u062c\u0630\5\u01d4\u00eb\2\u062d\u0630"+
		"\5\u01ca\u00e6\2\u062e\u0630\5\u0120\u0091\2\u062f\u061f\3\2\2\2\u062f"+
		"\u0623\3\2\2\2\u062f\u0627\3\2\2\2\u062f\u062b\3\2\2\2\u062f\u062c\3\2"+
		"\2\2\u062f\u062d\3\2\2\2\u062f\u062e\3\2\2\2\u0630\u011f\3\2\2\2\u0631"+
		"\u0638\t\5\2\2\u0632\u0638\5\u01be\u00e0\2\u0633\u0634\6\u0091\3\2\u0634"+
		"\u0638\7\u0091\2\2\u0635\u0636\6\u0091\4\2\u0636\u0638\7\u008e\2\2\u0637"+
		"\u0631\3\2\2\2\u0637\u0632\3\2\2\2\u0637\u0633\3\2\2\2\u0637\u0635\3\2"+
		"\2\2\u0638\u0121\3\2\2\2\u0639\u063b\5\u012a\u0096\2\u063a\u0639\3\2\2"+
		"\2\u063a\u063b\3\2\2\2\u063b\u063c\3\2\2\2\u063c\u063e\5\u0126\u0094\2"+
		"\u063d\u063f\5\u012e\u0098\2\u063e\u063d\3\2\2\2\u063e\u063f\3\2\2\2\u063f"+
		"\u0123\3\2\2\2\u0640\u0645\5\u0122\u0092\2\u0641\u0642\7\u0088\2\2\u0642"+
		"\u0644\5\u0122\u0092\2\u0643\u0641\3\2\2\2\u0644\u0647\3\2\2\2\u0645\u0643"+
		"\3\2\2\2\u0645\u0646\3\2\2\2\u0646\u0125\3\2\2\2\u0647\u0645\3\2\2\2\u0648"+
		"\u0649\5\u01c6\u00e4\2\u0649\u064a\5\u0174\u00bb\2\u064a\u064e\3\2\2\2"+
		"\u064b\u064e\5\u0174\u00bb\2\u064c\u064e\5\u0128\u0095\2\u064d\u0648\3"+
		"\2\2\2\u064d\u064b\3\2\2\2\u064d\u064c\3\2\2\2\u064e\u0127\3\2\2\2\u064f"+
		"\u0650\7\u0091\2\2\u0650\u0651\5\u01ac\u00d7\2\u0651\u0129\3\2\2\2\u0652"+
		"\u0653\7G\2\2\u0653\u0658\7\u008f\2\2\u0654\u0655\7G\2\2\u0655\u0658\7"+
		"\u008e\2\2\u0656\u0658\7G\2\2\u0657\u0652\3\2\2\2\u0657\u0654\3\2\2\2"+
		"\u0657\u0656\3\2\2\2\u0658\u012b\3\2\2\2\u0659\u065a\5\u01c4\u00e3\2\u065a"+
		"\u065b\5\u0126\u0094\2\u065b\u066a\3\2\2\2\u065c\u065e\5\u01b4\u00db\2"+
		"\u065d\u065f\5\u012a\u0096\2\u065e\u065d\3\2\2\2\u065e\u065f\3\2\2\2\u065f"+
		"\u0660\3\2\2\2\u0660\u0661\5\u0126\u0094\2\u0661\u066a\3\2\2\2\u0662\u0664"+
		"\5\u0130\u0099\2\u0663\u0665\5\u012a\u0096\2\u0664\u0663\3\2\2\2\u0664"+
		"\u0665\3\2\2\2\u0665\u0666\3\2\2\2\u0666\u0667\5\u0126\u0094\2\u0667\u066a"+
		"\3\2\2\2\u0668\u066a\5\u0132\u009a\2\u0669\u0659\3\2\2\2\u0669\u065c\3"+
		"\2\2\2\u0669\u0662\3\2\2\2\u0669\u0668\3\2\2\2\u066a\u012d\3\2\2\2\u066b"+
		"\u066d\5\u012c\u0097\2\u066c\u066b\3\2\2\2\u066d\u066e\3\2\2\2\u066e\u066c"+
		"\3\2\2\2\u066e\u066f\3\2\2\2\u066f\u012f\3\2\2\2\u0670\u0672\7\u008f\2"+
		"\2\u0671\u0673\5\u012a\u0096\2\u0672\u0671\3\2\2\2\u0672\u0673\3\2\2\2"+
		"\u0673\u0674\3\2\2\2\u0674\u0675\5\u0122\u0092\2\u0675\u0676\7\u0089\2"+
		"\2\u0676\u0131\3\2\2\2\u0677\u0678\7C\2\2\u0678\u0682\5\u0184\u00c3\2"+
		"\u0679\u067a\7D\2\2\u067a\u0682\5\u0184\u00c3\2\u067b\u067c\7D\2\2\u067c"+
		"\u067d\7\u008f\2\2\u067d\u0682\5\u0184\u00c3\2\u067e\u067f\7D\2\2\u067f"+
		"\u0680\7\u008e\2\2\u0680\u0682\5\u0184\u00c3\2\u0681\u0677\3\2\2\2\u0681"+
		"\u0679\3\2\2\2\u0681\u067b\3\2\2\2\u0681\u067e\3\2\2\2\u0682\u0133\3\2"+
		"\2\2\u0683\u0685\5\u01ac\u00d7\2\u0684\u0686\58\35\2\u0685\u0684\3\2\2"+
		"\2\u0685\u0686\3\2\2\2\u0686\u0692\3\2\2\2\u0687\u0692\5\u0136\u009c\2"+
		"\u0688\u0692\5\u0146\u00a4\2\u0689\u0692\5\u0148\u00a5\2\u068a\u0692\5"+
		"\u0150\u00a9\2\u068b\u0692\5\u0168\u00b5\2\u068c\u0692\5\u016a\u00b6\2"+
		"\u068d\u0692\5\u0166\u00b4\2\u068e\u0692\5\u016e\u00b8\2\u068f\u0692\5"+
		"\u0170\u00b9\2\u0690\u0692\5\u0172\u00ba\2\u0691\u0683\3\2\2\2\u0691\u0687"+
		"\3\2\2\2\u0691\u0688\3\2\2\2\u0691\u0689\3\2\2\2\u0691\u068a\3\2\2\2\u0691"+
		"\u068b\3\2\2\2\u0691\u068c\3\2\2\2\u0691\u068d\3\2\2\2\u0691\u068e\3\2"+
		"\2\2\u0691\u068f\3\2\2\2\u0691\u0690\3\2\2\2\u0692\u0135\3\2\2\2\u0693"+
		"\u069c\5\u01d4\u00eb\2\u0694\u069c\5\u0138\u009d\2\u0695\u069c\5\u013e"+
		"\u00a0\2\u0696\u069c\7H\2\2\u0697\u069c\7I\2\2\u0698\u069c\7J\2\2\u0699"+
		"\u069c\7K\2\2\u069a\u069c\7L\2\2\u069b\u0693\3\2\2\2\u069b\u0694\3\2\2"+
		"\2\u069b\u0695\3\2\2\2\u069b\u0696\3\2\2\2\u069b\u0697\3\2\2\2\u069b\u0698"+
		"\3\2\2\2\u069b\u0699\3\2\2\2\u069b\u069a\3\2\2\2\u069c\u0137\3\2\2\2\u069d"+
		"\u069f\7\u0084\2\2\u069e\u06a0\5\u013a\u009e\2\u069f\u069e\3\2\2\2\u069f"+
		"\u06a0\3\2\2\2\u06a0\u06a1\3\2\2\2\u06a1\u06a2\7\u0087\2\2\u06a2\u0139"+
		"\3\2\2\2\u06a3\u06a5\5\u013c\u009f\2\u06a4\u06a6\7\u0088\2\2\u06a5\u06a4"+
		"\3\2\2\2\u06a5\u06a6\3\2\2\2\u06a6\u06ac\3\2\2\2\u06a7\u06a8\5\u013c\u009f"+
		"\2\u06a8\u06a9\7\u0088\2\2\u06a9\u06aa\5\u013a\u009e\2\u06aa\u06ac\3\2"+
		"\2\2\u06ab\u06a3\3\2\2\2\u06ab\u06a7\3\2\2\2\u06ac\u013b\3\2\2\2\u06ad"+
		"\u06ae\5\u0122\u0092\2\u06ae\u013d\3\2\2\2\u06af\u06b0\7\u0084\2\2\u06b0"+
		"\u06b1\5\u0140\u00a1\2\u06b1\u06b2\7\u0087\2\2\u06b2\u06b7\3\2\2\2\u06b3"+
		"\u06b4\7\u0084\2\2\u06b4\u06b5\7\u0089\2\2\u06b5\u06b7\7\u0087\2\2\u06b6"+
		"\u06af\3\2\2\2\u06b6\u06b3\3\2\2\2\u06b7\u013f\3\2\2\2\u06b8\u06ba\5\u0142"+
		"\u00a2\2\u06b9\u06bb\7\u0088\2\2\u06ba\u06b9\3\2\2\2\u06ba\u06bb\3\2\2"+
		"\2\u06bb\u06c1\3\2\2\2\u06bc\u06bd\5\u0142\u00a2\2\u06bd\u06be\7\u0088"+
		"\2\2\u06be\u06bf\5\u0140\u00a1\2\u06bf\u06c1\3\2\2\2\u06c0\u06b8\3\2\2"+
		"\2\u06c0\u06bc\3\2\2\2\u06c1\u0141\3\2\2\2\u06c2\u06c3\5\u0122\u0092\2"+
		"\u06c3\u06c4\7\u0089\2\2\u06c4\u06c5\5\u0122\u0092\2\u06c5\u0143\3\2\2"+
		"\2\u06c6\u06c7\7M\2\2\u06c7\u06c8\7\u0083\2\2\u06c8\u06c9\7N\2\2\u06c9"+
		"\u06ca\7\u0089\2\2\u06ca\u06cb\5\u0122\u0092\2\u06cb\u06cc\7\u0088\2\2"+
		"\u06cc\u06cd\7O\2\2\u06cd\u06ce\7\u0089\2\2\u06ce\u06cf\5\u0122\u0092"+
		"\2\u06cf\u06d0\7\u0088\2\2\u06d0\u06d1\7P\2\2\u06d1\u06d2\7\u0089\2\2"+
		"\u06d2\u06d3\5\u0122\u0092\2\u06d3\u06d4\7\u0088\2\2\u06d4\u06d5\7Q\2"+
		"\2\u06d5\u06d6\7\u0089\2\2\u06d6\u06d7\5\u0122\u0092\2\u06d7\u06d8\7\u0086"+
		"\2\2\u06d8\u06e8\3\2\2\2\u06d9\u06da\7R\2\2\u06da\u06db\7\u0083\2\2\u06db"+
		"\u06dc\7S\2\2\u06dc\u06dd\7\u0089\2\2\u06dd\u06de\5\u0122\u0092\2\u06de"+
		"\u06df\7\u0086\2\2\u06df\u06e8\3\2\2\2\u06e0\u06e1\7T\2\2\u06e1\u06e2"+
		"\7\u0083\2\2\u06e2\u06e3\7S\2\2\u06e3\u06e4\7\u0089\2\2\u06e4\u06e5\5"+
		"\u0122\u0092\2\u06e5\u06e6\7\u0086\2\2\u06e6\u06e8\3\2\2\2\u06e7\u06c6"+
		"\3\2\2\2\u06e7\u06d9\3\2\2\2\u06e7\u06e0\3\2\2\2\u06e8\u0145\3\2\2\2\u06e9"+
		"\u06fd\7U\2\2\u06ea\u06eb\7U\2\2\u06eb\u06ec\7\u0081\2\2\u06ec\u06fd\5"+
		"\u01ac\u00d7\2\u06ed\u06ee\7U\2\2\u06ee\u06ef\7\u0084\2\2\u06ef\u06f0"+
		"\5\u0124\u0093\2\u06f0\u06f1\7\u0087\2\2\u06f1\u06fd\3\2\2\2\u06f2\u06f3"+
		"\7U\2\2\u06f3\u06f4\7\u0081\2\2\u06f4\u06fd\7!\2\2\u06f5\u06fd\7V\2\2"+
		"\u06f6\u06f7\7V\2\2\u06f7\u06f8\7\u0081\2\2\u06f8\u06fd\5\u01ac\u00d7"+
		"\2\u06f9\u06fa\7V\2\2\u06fa\u06fb\7\u0081\2\2\u06fb\u06fd\7!\2\2\u06fc"+
		"\u06e9\3\2\2\2\u06fc\u06ea\3\2\2\2\u06fc\u06ed\3\2\2\2\u06fc\u06f2\3\2"+
		"\2\2\u06fc\u06f5\3\2\2\2\u06fc\u06f6\3\2\2\2\u06fc\u06f9\3\2\2\2\u06fd"+
		"\u0147\3\2\2\2\u06fe\u0702\5\u014a\u00a6\2\u06ff\u0702\5\u014c\u00a7\2"+
		"\u0700\u0702\5\u014e\u00a8\2\u0701\u06fe\3\2\2\2\u0701\u06ff\3\2\2\2\u0701"+
		"\u0700\3\2\2\2\u0702\u0149\3\2\2\2\u0703\u0704\7W\2\2\u0704\u0705\7\u0081"+
		"\2\2\u0705\u0706\5\u01ac\u00d7\2\u0706\u014b\3\2\2\2\u0707\u0708\7W\2"+
		"\2\u0708\u0709\7\u0084\2\2\u0709\u070a\5\u0122\u0092\2\u070a\u070b\7\u0087"+
		"\2\2\u070b\u014d\3\2\2\2\u070c\u070d\7W\2\2\u070d\u070e\7\u0081\2\2\u070e"+
		"\u070f\7!\2\2\u070f\u014f\3\2\2\2\u0710\u0712\7\u0082\2\2\u0711\u0713"+
		"\5\u0152\u00aa\2\u0712\u0711\3\2\2\2\u0712\u0713\3\2\2\2\u0713\u0715\3"+
		"\2\2\2\u0714\u0716\5\4\3\2\u0715\u0714\3\2\2\2\u0715\u0716\3\2\2\2\u0716"+
		"\u0717\3\2\2\2\u0717\u0718\7\u0085\2\2\u0718\u0151\3\2\2\2\u0719\u071b"+
		"\5\u015e\u00b0\2\u071a\u0719\3\2\2\2\u071a\u071b\3\2\2\2\u071b\u071c\3"+
		"\2\2\2\u071c\u071e\5\u0154\u00ab\2\u071d\u071f\7\34\2\2\u071e\u071d\3"+
		"\2\2\2\u071e\u071f\3\2\2\2\u071f\u0721\3\2\2\2\u0720\u0722\5t;\2\u0721"+
		"\u0720\3\2\2\2\u0721\u0722\3\2\2\2\u0722\u0723\3\2\2\2\u0723\u0724\7\4"+
		"\2\2\u0724\u0729\3\2\2\2\u0725\u0726\5\u015e\u00b0\2\u0726\u0727\7\4\2"+
		"\2\u0727\u0729\3\2\2\2\u0728\u071a\3\2\2\2\u0728\u0725\3\2\2\2\u0729\u0153"+
		"\3\2\2\2\u072a\u072b\7\u0083\2\2\u072b\u0732\7\u0086\2\2\u072c\u072d\7"+
		"\u0083\2\2\u072d\u072e\5\u0158\u00ad\2\u072e\u072f\7\u0086\2\2\u072f\u0732"+
		"\3\2\2\2\u0730\u0732\5\u0156\u00ac\2\u0731\u072a\3\2\2\2\u0731\u072c\3"+
		"\2\2\2\u0731\u0730\3\2\2\2\u0732\u0155\3\2\2\2\u0733\u0738\5\u01ac\u00d7"+
		"\2\u0734\u0735\7\u0088\2\2\u0735\u0737\5\u01ac\u00d7\2\u0736\u0734\3\2"+
		"\2\2\u0737\u073a\3\2\2\2\u0738\u0736\3\2\2\2\u0738\u0739\3\2\2\2\u0739"+
		"\u0157\3\2\2\2\u073a\u0738\3\2\2\2\u073b\u0740\5\u015a\u00ae\2\u073c\u073d"+
		"\7\u0088\2\2\u073d\u073f\5\u015a\u00ae\2\u073e\u073c\3\2\2\2\u073f\u0742"+
		"\3\2\2\2\u0740\u073e\3\2\2\2\u0740\u0741\3\2\2\2\u0741\u0159\3\2\2\2\u0742"+
		"\u0740\3\2\2\2\u0743\u0745\5\u015c\u00af\2\u0744\u0746\5\u0186\u00c4\2"+
		"\u0745\u0744\3\2\2\2\u0745\u0746\3\2\2\2\u0746\u074c\3\2\2\2\u0747\u0748"+
		"\5\u015c\u00af\2\u0748\u0749\5\u0186\u00c4\2\u0749\u074a\5\u01c0\u00e1"+
		"\2\u074a\u074c\3\2\2\2\u074b\u0743\3\2\2\2\u074b\u0747\3\2\2\2\u074c\u015b"+
		"\3\2\2\2\u074d\u074e\5\u01ae\u00d8\2\u074e\u015d\3\2\2\2\u074f\u0750\7"+
		"\u0084\2\2\u0750\u0751\5\u0160\u00b1\2\u0751\u0752\7\u0087\2\2\u0752\u015f"+
		"\3\2\2\2\u0753\u0758\5\u0162\u00b2\2\u0754\u0755\7\u0088\2\2\u0755\u0757"+
		"\5\u0162\u00b2\2\u0756\u0754\3\2\2\2\u0757\u075a\3\2\2\2\u0758\u0756\3"+
		"\2\2\2\u0758\u0759\3\2\2\2\u0759\u0161\3\2\2\2\u075a\u0758\3\2\2\2\u075b"+
		"\u075d\5\u0164\u00b3\2\u075c\u075b\3\2\2\2\u075c\u075d\3\2\2\2\u075d\u075e"+
		"\3\2\2\2\u075e\u075f\5\u0122\u0092\2\u075f\u0163\3\2\2\2\u0760\u0761\t"+
		"\6\2\2\u0761\u0165\3\2\2\2\u0762\u0763\7\u0081\2\2\u0763\u0764\5\u01ae"+
		"\u00d8\2\u0764\u0167\3\2\2\2\u0765\u0766\7\u0083\2\2\u0766\u0767\5\u0122"+
		"\u0092\2\u0767\u0768\7\u0086\2\2\u0768\u0169\3\2\2\2\u0769\u076a\7\u0083"+
		"\2\2\u076a\u0776\7\u0086\2\2\u076b\u076c\7\u0083\2\2\u076c\u076f\5\u016c"+
		"\u00b7\2\u076d\u076e\7\u0088\2\2\u076e\u0770\5\u016c\u00b7\2\u076f\u076d"+
		"\3\2\2\2\u0770\u0771\3\2\2\2\u0771\u076f\3\2\2\2\u0771\u0772\3\2\2\2\u0772"+
		"\u0773\3\2\2\2\u0773\u0774\7\u0086\2\2\u0774\u0776\3\2\2\2\u0775\u0769"+
		"\3\2\2\2\u0775\u076b\3\2\2\2\u0776\u016b\3\2\2\2\u0777\u077d\5\u0122\u0092"+
		"\2\u0778\u0779\5\u01ae\u00d8\2\u0779\u077a\7\u0089\2\2\u077a\u077b\5\u0122"+
		"\u0092\2\u077b\u077d\3\2\2\2\u077c\u0777\3\2\2\2\u077c\u0778\3\2\2\2\u077d"+
		"\u016d\3\2\2\2\u077e\u077f\7\u008d\2\2\u077f\u016f\3\2\2\2\u0780\u0781"+
		"\7Z\2\2\u0781\u0782\7\u0083\2\2\u0782\u0783\5\u0122\u0092\2\u0783\u0784"+
		"\7\u0086\2\2\u0784\u0792\3\2\2\2\u0785\u0786\7Z\2\2\u0786\u0787\7\u0083"+
		"\2\2\u0787\u0788\7[\2\2\u0788\u0789\5\u0122\u0092\2\u0789\u078a\7\u0086"+
		"\2\2\u078a\u0792\3\2\2\2\u078b\u078c\7Z\2\2\u078c\u078d\7\u0083\2\2\u078d"+
		"\u078e\7\\\2\2\u078e\u078f\5\u0122\u0092\2\u078f\u0790\7\u0086\2\2\u0790"+
		"\u0792\3\2\2\2\u0791\u0780\3\2\2\2\u0791\u0785\3\2\2\2\u0791\u078b\3\2"+
		"\2\2\u0792\u0171\3\2\2\2\u0793\u0794\7]\2\2\u0794\u0795\7\u0083\2\2\u0795"+
		"\u0796\5\u0122\u0092\2\u0796\u0797\7\u0086\2\2\u0797\u0173\3\2\2\2\u0798"+
		"\u0799\b\u00bb\1\2\u0799\u079c\5\u0134\u009b\2\u079a\u079c\5\u0182\u00c2"+
		"\2\u079b\u0798\3\2\2\2\u079b\u079a\3\2\2\2\u079c\u07cf\3\2\2\2\u079d\u079e"+
		"\f\16\2\2\u079e\u07ce\5\u01c8\u00e5\2\u079f\u07a0\f\r\2\2\u07a0\u07ce"+
		"\5\u0176\u00bc\2\u07a1\u07a3\f\f\2\2\u07a2\u07a4\5\u0176\u00bc\2\u07a3"+
		"\u07a2\3\2\2\2\u07a3\u07a4\3\2\2\2\u07a4\u07a5\3\2\2\2\u07a5\u07ce\5\u017c"+
		"\u00bf\2\u07a6\u07a7\f\13\2\2\u07a7\u07a8\7\u0081\2\2\u07a8\u07ce\7!\2"+
		"\2\u07a9\u07aa\f\n\2\2\u07aa\u07ab\7\u0081\2\2\u07ab\u07ac\7!\2\2\u07ac"+
		"\u07ad\7\u0083\2\2\u07ad\u07ae\5\u017e\u00c0\2\u07ae\u07af\7\u0086\2\2"+
		"\u07af\u07ce\3\2\2\2\u07b0\u07b1\f\t\2\2\u07b1\u07b2\7\u0081\2\2\u07b2"+
		"\u07ce\7\u00a1\2\2\u07b3\u07b4\f\b\2\2\u07b4\u07b5\7\u0081\2\2\u07b5\u07b7"+
		"\5\u01ac\u00d7\2\u07b6\u07b8\58\35\2\u07b7\u07b6\3\2\2\2\u07b7\u07b8\3"+
		"\2\2\2\u07b8\u07ce\3\2\2\2\u07b9\u07ba\f\7\2\2\u07ba\u07bb\7\u0081\2\2"+
		"\u07bb\u07bc\5\u01ac\u00d7\2\u07bc\u07bd\7\u0083\2\2\u07bd\u07be\5\u017e"+
		"\u00c0\2\u07be\u07bf\7\u0086\2\2\u07bf\u07ce\3\2\2\2\u07c0\u07c1\f\6\2"+
		"\2\u07c1\u07c2\7\u0083\2\2\u07c2\u07c3\5\u017e\u00c0\2\u07c3\u07c4\7\u0086"+
		"\2\2\u07c4\u07ce\3\2\2\2\u07c5\u07c6\f\5\2\2\u07c6\u07c7\7\u0081\2\2\u07c7"+
		"\u07ce\7U\2\2\u07c8\u07c9\f\3\2\2\u07c9\u07ca\7\u0084\2\2\u07ca\u07cb"+
		"\5\u0124\u0093\2\u07cb\u07cc\7\u0087\2\2\u07cc\u07ce\3\2\2\2\u07cd\u079d"+
		"\3\2\2\2\u07cd\u079f\3\2\2\2\u07cd\u07a1\3\2\2\2\u07cd\u07a6\3\2\2\2\u07cd"+
		"\u07a9\3\2\2\2\u07cd\u07b0\3\2\2\2\u07cd\u07b3\3\2\2\2\u07cd\u07b9\3\2"+
		"\2\2\u07cd\u07c0\3\2\2\2\u07cd\u07c5\3\2\2\2\u07cd\u07c8\3\2\2\2\u07ce"+
		"\u07d1\3\2\2\2\u07cf\u07cd\3\2\2\2\u07cf\u07d0\3\2\2\2\u07d0\u0175\3\2"+
		"\2\2\u07d1\u07cf\3\2\2\2\u07d2\u07d3\7\u0083\2\2\u07d3\u07d9\7\u0086\2"+
		"\2\u07d4\u07d5\7\u0083\2\2\u07d5\u07d6\5\u0178\u00bd\2\u07d6\u07d7\7\u0086"+
		"\2\2\u07d7\u07d9\3\2\2\2\u07d8\u07d2\3\2\2\2\u07d8\u07d4\3\2\2\2\u07d9"+
		"\u0177\3\2\2\2\u07da\u07df\5\u017a\u00be\2\u07db\u07dc\7\u0088\2\2\u07dc"+
		"\u07de\5\u017a\u00be\2\u07dd\u07db\3\2\2\2\u07de\u07e1\3\2\2\2\u07df\u07dd"+
		"\3\2\2\2\u07df\u07e0\3\2\2\2\u07e0\u0179\3\2\2\2\u07e1\u07df\3\2\2\2\u07e2"+
		"\u07ed\5\u0122\u0092\2\u07e3\u07e4\5\u01ae\u00d8\2\u07e4\u07e5\7\u0089"+
		"\2\2\u07e5\u07e6\5\u0122\u0092\2\u07e6\u07ed\3\2\2\2\u07e7\u07ed\5\u01ca"+
		"\u00e6\2\u07e8\u07e9\5\u01ae\u00d8\2\u07e9\u07ea\7\u0089\2\2\u07ea\u07eb"+
		"\5\u01ca\u00e6\2\u07eb\u07ed\3\2\2\2\u07ec\u07e2\3\2\2\2\u07ec\u07e3\3"+
		"\2\2\2\u07ec\u07e7\3\2\2\2\u07ec\u07e8\3\2\2\2\u07ed\u017b\3\2\2\2\u07ee"+
		"\u07ef\5\u0150\u00a9\2\u07ef\u017d\3\2\2\2\u07f0\u07f4\5\u0180\u00c1\2"+
		"\u07f1\u07f3\5\u0180\u00c1\2\u07f2\u07f1\3\2\2\2\u07f3\u07f6\3\2\2\2\u07f4"+
		"\u07f2\3\2\2\2\u07f4\u07f5\3\2\2\2\u07f5\u017f\3\2\2\2\u07f6\u07f4\3\2"+
		"\2\2\u07f7\u07f8\5\u01ae\u00d8\2\u07f8\u07f9\7\u0089\2\2\u07f9\u0181\3"+
		"\2\2\2\u07fa\u07fb\7^\2\2\u07fb\u07fc\7\u0083\2\2\u07fc\u07fd\7_\2\2\u07fd"+
		"\u07fe\7\u0089\2\2\u07fe\u07ff\5\u0122\u0092\2\u07ff\u0800\7\u0086\2\2"+
		"\u0800\u0183\3\2\2\2\u0801\u0802\b\u00c3\1\2\u0802\u080b\5\u019e\u00d0"+
		"\2\u0803\u080b\5\u01a0\u00d1\2\u0804\u080b\5\u0194\u00cb\2\u0805\u080b"+
		"\5\u0188\u00c5\2\u0806\u080b\5\u018c\u00c7\2\u0807\u080b\5\u01a2\u00d2"+
		"\2\u0808\u080b\7b\2\2\u0809\u080b\7V\2\2\u080a\u0801\3\2\2\2\u080a\u0803"+
		"\3\2\2\2\u080a\u0804\3\2\2\2\u080a\u0805\3\2\2\2\u080a\u0806\3\2\2\2\u080a"+
		"\u0807\3\2\2\2\u080a\u0808\3\2\2\2\u080a\u0809\3\2\2\2\u080b\u0818\3\2"+
		"\2\2\u080c\u080d\f\t\2\2\u080d\u0817\7\u008f\2\2\u080e\u080f\f\b\2\2\u080f"+
		"\u0817\7\u008e\2\2\u0810\u0811\f\6\2\2\u0811\u0812\7\u0081\2\2\u0812\u0817"+
		"\7`\2\2\u0813\u0814\f\5\2\2\u0814\u0815\7\u0081\2\2\u0815\u0817\7a\2\2"+
		"\u0816\u080c\3\2\2\2\u0816\u080e\3\2\2\2\u0816\u0810\3\2\2\2\u0816\u0813"+
		"\3\2\2\2\u0817\u081a\3\2\2\2\u0818\u0816\3\2\2\2\u0818\u0819\3\2\2\2\u0819"+
		"\u0185\3\2\2\2\u081a\u0818\3\2\2\2\u081b\u081d\7\u0089\2\2\u081c\u081e"+
		"\5\u011a\u008e\2\u081d\u081c\3\2\2\2\u081d\u081e\3\2\2\2\u081e\u0820\3"+
		"\2\2\2\u081f\u0821\7c\2\2\u0820\u081f\3\2\2\2\u0820\u0821\3\2\2\2\u0821"+
		"\u0822\3\2\2\2\u0822\u0823\5\u0184\u00c3\2\u0823\u0187\3\2\2\2\u0824\u0826"+
		"\5\u018a\u00c6\2\u0825\u0827\58\35\2\u0826\u0825\3\2\2\2\u0826\u0827\3"+
		"\2\2\2\u0827\u082a\3\2\2\2\u0828\u0829\7\u0081\2\2\u0829\u082b\5\u0188"+
		"\u00c5\2\u082a\u0828\3\2\2\2\u082a\u082b\3\2\2\2\u082b\u0189\3\2\2\2\u082c"+
		"\u082d\5\u01ac\u00d7\2\u082d\u018b\3\2\2\2\u082e\u0830\7\u0083\2\2\u082f"+
		"\u0831\5\u018e\u00c8\2\u0830\u082f\3\2\2\2\u0830\u0831\3\2\2\2\u0831\u0832"+
		"\3\2\2\2\u0832\u0833\7\u0086\2\2\u0833\u018d\3\2\2\2\u0834\u083a\5\u0190"+
		"\u00c9\2\u0835\u0836\5\u0190\u00c9\2\u0836\u0837\7\u0088\2\2\u0837\u0838"+
		"\5\u018e\u00c8\2\u0838\u083a\3\2\2\2\u0839\u0834\3\2\2\2\u0839\u0835\3"+
		"\2\2\2\u083a\u018f\3\2\2\2\u083b\u083c\5\u0192\u00ca\2\u083c\u083d\5\u0186"+
		"\u00c4\2\u083d\u0840\3\2\2\2\u083e\u0840\5\u0184\u00c3\2\u083f\u083b\3"+
		"\2\2\2\u083f\u083e\3\2\2\2\u0840\u0191\3\2\2\2\u0841\u0842\5\u01ae\u00d8"+
		"\2\u0842\u0193\3\2\2\2\u0843\u0845\5\u011a\u008e\2\u0844\u0843\3\2\2\2"+
		"\u0844\u0845\3\2\2\2\u0845\u0846\3\2\2\2\u0846\u0848\5\u0196\u00cc\2\u0847"+
		"\u0849\7\34\2\2\u0848\u0847\3\2\2\2\u0848\u0849\3\2\2\2\u0849\u084a\3"+
		"\2\2\2\u084a\u084b\5\u01be\u00e0\2\u084b\u084c\5\u0184\u00c3\2\u084c\u0856"+
		"\3\2\2\2\u084d\u084f\5\u011a\u008e\2\u084e\u084d\3\2\2\2\u084e\u084f\3"+
		"\2\2\2\u084f\u0850\3\2\2\2\u0850\u0851\5\u0196\u00cc\2\u0851\u0852\7\35"+
		"\2\2\u0852\u0853\5\u01be\u00e0\2\u0853\u0854\5\u0184\u00c3\2\u0854\u0856"+
		"\3\2\2\2\u0855\u0844\3\2\2\2\u0855\u084e\3\2\2\2\u0856\u0195\3\2\2\2\u0857"+
		"\u0858\7\u0083\2\2\u0858\u0861\7\u0086\2\2\u0859\u085a\7\u0083\2\2\u085a"+
		"\u085c\5\u0198\u00cd\2\u085b\u085d\5\u01c0\u00e1\2\u085c\u085b\3\2\2\2"+
		"\u085c\u085d\3\2\2\2\u085d\u085e\3\2\2\2\u085e\u085f\7\u0086\2\2\u085f"+
		"\u0861\3\2\2\2\u0860\u0857\3\2\2\2\u0860\u0859\3\2\2\2\u0861\u0197\3\2"+
		"\2\2\u0862\u0868\5\u019a\u00ce\2\u0863\u0864\5\u019a\u00ce\2\u0864\u0865"+
		"\7\u0088\2\2\u0865\u0866\5\u0198\u00cd\2\u0866\u0868\3\2\2\2\u0867\u0862"+
		"\3\2\2\2\u0867\u0863\3\2\2\2\u0868\u0199\3\2\2\2\u0869\u086b\5\u011a\u008e"+
		"\2\u086a\u0869\3\2\2\2\u086a\u086b\3\2\2\2\u086b\u086d\3\2\2\2\u086c\u086e"+
		"\7c\2\2\u086d\u086c\3\2\2\2\u086d\u086e\3\2\2\2\u086e\u086f\3\2\2\2\u086f"+
		"\u0874\5\u0184\u00c3\2\u0870\u0871\5\u019c\u00cf\2\u0871\u0872\5\u0186"+
		"\u00c4\2\u0872\u0874\3\2\2\2\u0873\u086a\3\2\2\2\u0873\u0870\3\2\2\2\u0874"+
		"\u019b\3\2\2\2\u0875\u0876\5\u01ae\u00d8\2\u0876\u019d\3\2\2\2\u0877\u0878"+
		"\7\u0084\2\2\u0878\u0879\5\u0184\u00c3\2\u0879\u087a\7\u0087\2\2\u087a"+
		"\u019f\3\2\2\2\u087b\u087c\7\u0084\2\2\u087c\u087d\5\u0184\u00c3\2\u087d"+
		"\u087e\7\u0089\2\2\u087e\u087f\5\u0184\u00c3\2\u087f\u0880\7\u0087\2\2"+
		"\u0880\u01a1\3\2\2\2\u0881\u0884\5\u01a4\u00d3\2\u0882\u0883\7\u0091\2"+
		"\2\u0883\u0885\5\u01a4\u00d3\2\u0884\u0882\3\2\2\2\u0885\u0886\3\2\2\2"+
		"\u0886\u0884\3\2\2\2\u0886\u0887\3\2\2\2\u0887\u01a3\3\2\2\2\u0888\u0889"+
		"\5\u0188\u00c5\2\u0889\u01a5\3\2\2\2\u088a\u088b\7\u0089\2\2\u088b\u088c"+
		"\5\u01aa\u00d6\2\u088c\u088d\7\u0088\2\2\u088d\u088e\5\u01a8\u00d5\2\u088e"+
		"\u0894\3\2\2\2\u088f\u0890\7\u0089\2\2\u0890\u0894\5\u01aa\u00d6\2\u0891"+
		"\u0892\7\u0089\2\2\u0892\u0894\5\u01a8\u00d5\2\u0893\u088a\3\2\2\2\u0893"+
		"\u088f\3\2\2\2\u0893\u0891\3\2\2\2\u0894\u01a7\3\2\2\2\u0895\u089b\5\u0188"+
		"\u00c5\2\u0896\u0897\5\u0188\u00c5\2\u0897\u0898\7\u0088\2\2\u0898\u0899"+
		"\5\u01a8\u00d5\2\u0899\u089b\3\2\2\2\u089a\u0895\3\2\2\2\u089a\u0896\3"+
		"\2\2\2\u089b\u01a9\3\2\2\2\u089c\u089d\7\22\2\2\u089d\u01ab\3\2\2\2\u089e"+
		"\u08a1\7\u0080\2\2\u089f\u08a1\5\u01b0\u00d9\2\u08a0\u089e\3\2\2\2\u08a0"+
		"\u089f\3\2\2\2\u08a1\u01ad\3\2\2\2\u08a2\u08a5\7\u0080\2\2\u08a3\u08a5"+
		"\5\u01b2\u00da\2\u08a4\u08a2\3\2\2\2\u08a4\u08a3\3\2\2\2\u08a5\u01af\3"+
		"\2\2\2\u08a6\u08a7\t\7\2\2\u08a7\u01b1\3\2\2\2\u08a8\u08a9\t\b\2\2\u08a9"+
		"\u01b3\3\2\2\2\u08aa\u08ab\6\u00db\24\2\u08ab\u08ac\7\u0093\2\2\u08ac"+
		"\u01b5\3\2\2\2\u08ad\u08ae\6\u00dc\25\2\u08ae\u08af\7\u0092\2\2\u08af"+
		"\u01b7\3\2\2\2\u08b0\u08b1\6\u00dd\26\2\u08b1\u08b2\7\u0091\2\2\u08b2"+
		"\u08b3\7\u0091\2\2\u08b3\u01b9\3\2\2\2\u08b4\u08b5\6\u00de\27\2\u08b5"+
		"\u08b6\7\u0094\2\2\u08b6\u08b7\7\u0094\2\2\u08b7\u01bb\3\2\2\2\u08b8\u08b9"+
		"\6\u00df\30\2\u08b9\u08ba\7\u008c\2\2\u08ba\u08bb\7\u0093\2\2\u08bb\u01bd"+
		"\3\2\2\2\u08bc\u08bd\6\u00e0\31\2\u08bd\u08be\7\u0092\2\2\u08be\u08bf"+
		"\7\u008c\2\2\u08bf\u01bf\3\2\2\2\u08c0\u08c1\6\u00e1\32\2\u08c1\u08c2"+
		"\7\u0081\2\2\u08c2\u08c3\7\u0081\2\2\u08c3\u08c4\7\u0081\2\2\u08c4\u01c1"+
		"\3\2\2\2\u08c5\u08c6\6\u00e2\33\2\u08c6\u08c7\7\u0093\2\2\u08c7\u08c8"+
		"\7\u0093\2\2\u08c8\u01c3\3\2\2\2\u08c9\u08ca\6\u00e3\34\2\u08ca\u08cb"+
		"\5\u01ca\u00e6\2\u08cb\u01c5\3\2\2\2\u08cc\u08cd\6\u00e4\35\2\u08cd\u08ce"+
		"\5\u01ca\u00e6\2\u08ce\u01c7\3\2\2\2\u08cf\u08d0\6\u00e5\36\2\u08d0\u08d1"+
		"\5\u01ca\u00e6\2\u08d1\u01c9\3\2\2\2\u08d2\u08d7\5\u01ce\u00e8\2\u08d3"+
		"\u08d4\6\u00e6\37\2\u08d4\u08d6\5\u01cc\u00e7\2\u08d5\u08d3\3\2\2\2\u08d6"+
		"\u08d9\3\2\2\2\u08d7\u08d5\3\2\2\2\u08d7\u08d8\3\2\2\2\u08d8\u08e3\3\2"+
		"\2\2\u08d9\u08d7\3\2\2\2\u08da\u08df\5\u01d0\u00e9\2\u08db\u08dc\6\u00e6"+
		" \2\u08dc\u08de\5\u01d2\u00ea\2\u08dd\u08db\3\2\2\2\u08de\u08e1\3\2\2"+
		"\2\u08df\u08dd\3\2\2\2\u08df\u08e0\3\2\2\2\u08e0\u08e3\3\2\2\2\u08e1\u08df"+
		"\3\2\2\2\u08e2\u08d2\3\2\2\2\u08e2\u08da\3\2\2\2\u08e3\u01cb\3\2\2\2\u08e4"+
		"\u08e7\5\u01ce\u00e8\2\u08e5\u08e7\7\u009c\2\2\u08e6\u08e4\3\2\2\2\u08e6"+
		"\u08e5\3\2\2\2\u08e7\u01cd\3\2\2\2\u08e8\u08eb\t\t\2\2\u08e9\u08eb\7\u009b"+
		"\2\2\u08ea\u08e8\3\2\2\2\u08ea\u08e9\3\2\2\2\u08eb\u01cf\3\2\2\2\u08ec"+
		"\u08ed\7\u0081\2\2\u08ed\u01d1\3\2\2\2\u08ee\u08f1\7\u0081\2\2\u08ef\u08f1"+
		"\5\u01cc\u00e7\2\u08f0\u08ee\3\2\2\2\u08f0\u08ef\3\2\2\2\u08f1\u01d3\3"+
		"\2\2\2\u08f2\u08f7\5\u01d6\u00ec\2\u08f3\u08f7\5\u01de\u00f0\2\u08f4\u08f7"+
		"\5\u01d8\u00ed\2\u08f5\u08f7\5\u01da\u00ee\2\u08f6\u08f2\3\2\2\2\u08f6"+
		"\u08f3\3\2\2\2\u08f6\u08f4\3\2\2\2\u08f6\u08f5\3\2\2\2\u08f7\u01d5\3\2"+
		"\2\2\u08f8\u08fa\5\u01b6\u00dc\2\u08f9\u08f8\3\2\2\2\u08f9\u08fa\3\2\2"+
		"\2\u08fa\u08fb\3\2\2\2\u08fb\u0901\5\u01dc\u00ef\2\u08fc\u08fe\5\u01b6"+
		"\u00dc\2\u08fd\u08fc\3\2\2\2\u08fd\u08fe\3\2\2\2\u08fe\u08ff\3\2\2\2\u08ff"+
		"\u0901\7\u00a3\2\2\u0900\u08f9\3\2\2\2\u0900\u08fd\3\2\2\2\u0901\u01d7"+
		"\3\2\2\2\u0902\u0903\t\n\2\2\u0903\u01d9\3\2\2\2\u0904\u0905\7z\2\2\u0905"+
		"\u01db\3\2\2\2\u0906\u0907\t\13\2\2\u0907\u01dd\3\2\2\2\u0908\u0909\t"+
		"\f\2\2\u0909\u01df\3\2\2\2\u0105\u01e6\u01ee\u01f2\u0201\u0207\u020d\u0218"+
		"\u021e\u0226\u0232\u023a\u0245\u0251\u025b\u0260\u026a\u0279\u028c\u0291"+
		"\u0295\u02a0\u02a5\u02af\u02b4\u02cc\u02d5\u02dc\u02df\u02e2\u02eb\u02f4"+
		"\u02f7\u02fa\u0300\u0303\u0307\u0312\u031b\u031e\u0321\u0326\u0329\u0330"+
		"\u0339\u033c\u0340\u0345\u0349\u0350\u0354\u0357\u035a\u035d\u0363\u0367"+
		"\u036a\u036f\u0371\u0375\u0381\u0388\u038c\u0391\u0394\u039a\u03a0\u03aa"+
		"\u03ad\u03b1\u03b4\u03b7\u03ba\u03bf\u03c2\u03c5\u03c9\u03cf\u03d3\u03d6"+
		"\u03d9\u03e3\u03e7\u03f0\u03f4\u03fc\u0400\u0403\u040d\u0411\u0419\u041c"+
		"\u041f\u0424\u0427\u042a\u0434\u043c\u043f\u0442\u0447\u044a\u044d\u0452"+
		"\u0455\u0459\u045e\u0461\u0464\u0468\u0470\u0478\u047b\u0480\u048a\u0496"+
		"\u04a0\u04a4\u04a8\u04ac\u04af\u04b3\u04b8\u04ba\u04c1\u04c4\u04c8\u04cc"+
		"\u04d0\u04d3\u04d9\u04de\u04e2\u04e5\u04e8\u04ec\u04ef\u04f4\u04f7\u04fb"+
		"\u0500\u0506\u0509\u050e\u0513\u0516\u051d\u0523\u0536\u0539\u053c\u0543"+
		"\u054a\u0558\u0563\u056b\u0573\u0584\u05a1\u05a6\u05c1\u05c8\u05cc\u05d1"+
		"\u05d8\u05df\u05ea\u05ee\u05f7\u05fd\u0602\u060c\u0617\u061c\u062f\u0637"+
		"\u063a\u063e\u0645\u064d\u0657\u065e\u0664\u0669\u066e\u0672\u0681\u0685"+
		"\u0691\u069b\u069f\u06a5\u06ab\u06b6\u06ba\u06c0\u06e7\u06fc\u0701\u0712"+
		"\u0715\u071a\u071e\u0721\u0728\u0731\u0738\u0740\u0745\u074b\u0758\u075c"+
		"\u0771\u0775\u077c\u0791\u079b\u07a3\u07b7\u07cd\u07cf\u07d8\u07df\u07ec"+
		"\u07f4\u080a\u0816\u0818\u081d\u0820\u0826\u082a\u0830\u0839\u083f\u0844"+
		"\u0848\u084e\u0855\u085c\u0860\u0867\u086a\u086d\u0873\u0886\u0893\u089a"+
		"\u08a0\u08a4\u08d7\u08df\u08e2\u08e6\u08ea\u08f0\u08f6\u08f9\u08fd\u0900";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}