grammar Tswift;

// Tokens a ignorar
COMENTARIO_COMPUESTO : '/*' .*? '*/' -> skip;

COMENTARIO_SIMPLE : '//' ~[\r\n]* -> skip;

inicio
   : instrucciones EOF
   ;

instrucciones
   : instruccion*
   ;

instruccion
   : expresion
   | declaracion
   | loop_statement
   | branch_statement
   | control_transfer_statement
   ;

loop_statement 
   : for_in_statement
   | while_statement
   ;

// GRAMMAR OF A FOR_IN STATEMENT

for_in_statement : 'for' pattern 'in' expresion code_block ; //cambiar pattern a ID y falta rango

// GRAMMAR OF A WHILE STATEMENT

while_statement : 'while' expresion code_block ;

// GRAMMAR OF A BRANCH STATEMENT

branch_statement : if_statement
 | guard_statement
 | switch_statement
 ;

// GRAMMAR OF AN IF STATEMENT

if_statement : 'if' expresion code_block else_clause? ;
else_clause : 'else' code_block | 'else' if_statement  ;

// GRAMMAR OF A GUARD STATEMENT

guard_statement : 'guard' expresion 'else' code_block ;

// GRAMMAR OF A SWITCH STATEMENT

switch_statement : 'switch' expresion '{' switch_cases? '}'  ;
switch_cases : switch_case switch_cases? ;
switch_case : case_label instrucciones | default_label instrucciones  ;
case_label : 'case' expresion ':' ;
default_label : 'default' ':' ;

// GRAMMAR OF A CONTROL TRANSFER STATEMENT

control_transfer_statement 
   : break_statement
   | continue_statement
   | return_statement
   ;

// GRAMMAR OF A BREAK STATEMENT

break_statement : 'break';

// GRAMMAR OF A CONTINUE STATEMENT

continue_statement : 'continue';

// GRAMMAR OF A RETURN STATEMENT

return_statement : 'return' expresion? ;

// Generic Parameters and Arguments

// GRAMMAR OF A GENERIC PARAMETER CLAUSE

generic_parameter_clause : '<' generic_parameter_list '>'  ;
generic_parameter_list : generic_parameter (',' generic_parameter)*  ;
generic_parameter
 : type_name
 | type_name ':' type_identifier
 | type_name ':' protocol_composition_type
 ;

generic_where_clause : 'where' requirement_list ;
requirement_list : requirement (',' requirement)*  ;
requirement : conformance_requirement | same_type_requirement  ;

conformance_requirement : type_identifier ':' type_identifier | type_identifier ':' protocol_composition_type  ;
same_type_requirement : type_identifier same_type_equals type_  ;

// GRAMMAR OF A GENERIC ARGUMENT CLAUSE

generic_argument_clause : '<' generic_argument_list '>'  ;
generic_argument_list : generic_argument (',' generic_argument)* ;
generic_argument : type_ ;

// context-sensitive. Allow < as pre, post, or binary op
//lt : {_input.LT(1).getText().equals("<")}? operator_ ;
//gt : {_input.LT(1).getText().equals(">")}? operator_ ;
// declaracions

// GRAMMAR OF A declaracion

declaracion
   : constant_declaracion
   | variable_declaracion
   | function_declaracion
   | enum_declaracion
   | struct_declaracion
   | class_declaracion
   | protocol_declaracion
   | initializer_declaracion
   | deinitializer_declaracion
   | extension_declaracion
   | subscript_declaracion
   | operator_declaracion
   | operator_declaracion
   | precedence_group_declaracion
   ;

declaracions : declaracion+ ;

// GRAMMAR OF A CODE BLOCK

code_block : '{' instrucciones? '}' ;

// GRAMMAR OF AN IMPORT declaracion

import_kind : 'typealias' | 'struct' | 'class' | 'enum' | 'protocol' | 'var' | 'func'  ;
import_path : import_path_identifier ('.' import_path_identifier)*  ;
import_path_identifier : declaracion_identifier | operator_  ;

// GRAMMAR OF A CONSTANT declaracion

constant_declaracion : 'let' pattern_initializer_list  ;
pattern_initializer_list : pattern_initializer (',' pattern_initializer)* ;

/** rule is ambiguous. can match "var x = 1" with x as pattern
 *  OR with x as expresion_pattern.
 *  ANTLR resolves in favor or first choice: pattern is x, 1 is initializer.
 */
pattern_initializer : pattern initializer? ;
initializer : assignment_operator expresion  ;

// GRAMMAR OF A VARIABLE declaracion

variable_declaracion
 : variable_declaracion_head variable_name type_annotation code_block
 | variable_declaracion_head variable_name type_annotation getter_setter_block
 | variable_declaracion_head variable_name type_annotation getter_setter_keyword_block
 | variable_declaracion_head variable_name type_annotation initializer? willSet_didSet_block
 | variable_declaracion_head variable_name type_annotation type_annotation initializer? willSet_didSet_block
 | variable_declaracion_head pattern_initializer_list
 ;

variable_declaracion_head : attributes? declaracion_modifiers? 'var'  ;
variable_name : declaracion_identifier  ;

getter_setter_block : '{' getter_clause setter_clause?'}'  | '{' setter_clause getter_clause '}'  ;
getter_clause : attributes? mutation_modifier? 'get' code_block  ;
setter_clause : attributes? mutation_modifier? 'set' setter_name? code_block  ;
setter_name : '(' declaracion_identifier ')'  ;

getter_setter_keyword_block : '{' getter_keyword_clause setter_keyword_clause?'}' | '{' setter_keyword_clause getter_keyword_clause '}'  ;
getter_keyword_clause : attributes? mutation_modifier? 'get'  ;
setter_keyword_clause : attributes? mutation_modifier? 'set'  ;

willSet_didSet_block : '{' willSet_clause didSet_clause?'}' | '{' didSet_clause willSet_clause '}'  ;
willSet_clause : attributes? 'willSet' setter_name? code_block  ;
didSet_clause : attributes? 'didSet' setter_name? code_block  ;

// GRAMMAR OF A FUNCTION declaracion
function_declaracion : function_head function_name generic_parameter_clause? function_signature generic_where_clause? function_body? ;
 
function_head : attributes? declaracion_modifiers? 'func' ;

function_name : declaracion_identifier | operator_ ;

function_signature
 : parameter_clause 'throws'? function_result?
 | parameter_clause 'rethrows' function_result?
 ;
 
function_result : arrow_operator attributes? type_ ;

function_body : code_block ;

parameter_clause : '(' ')' |  '(' parameter_list ')'  ;
parameter_list : parameter (',' parameter)*  ;

parameter
 : external_parameter_name? local_parameter_name type_annotation default_argument_clause?
 | external_parameter_name? local_parameter_name type_annotation
 | external_parameter_name? local_parameter_name type_annotation range_operator
 ;
external_parameter_name : label_identifier ; // TODO: Check that deleting " | '_'" doesn't break anything
local_parameter_name : label_identifier ; // TODO: Check that deleting " | '_'" doesn't break anything
default_argument_clause : assignment_operator expresion ;

// GRAMMAR OF AN ENUMERATION declaracion

enum_declaracion : attributes? access_level_modifier? union_style_enum | attributes? access_level_modifier? raw_value_style_enum  ;

union_style_enum : 'indirect'? 'enum' enum_name generic_parameter_clause? type_inheritance_clause? generic_where_clause? '{' union_style_enum_members?'}' ;

union_style_enum_members : union_style_enum_member union_style_enum_members? ;

union_style_enum_member
 : declaracion
 | union_style_enum_case_clause
 ;

union_style_enum_case_clause : attributes? 'indirect'? 'case' union_style_enum_case_list  ;

union_style_enum_case_list : union_style_enum_case | union_style_enum_case ',' union_style_enum_case_list  ;

union_style_enum_case : enum_case_name tuple_type? ;

enum_name : declaracion_identifier  ;

enum_case_name : declaracion_identifier  ;

raw_value_style_enum : 'enum' enum_name generic_parameter_clause? type_inheritance_clause generic_where_clause? '{' raw_value_style_enum_members '}' ;

raw_value_style_enum_members : raw_value_style_enum_member raw_value_style_enum_members? ;

raw_value_style_enum_member
 : declaracion
 | raw_value_style_enum_case_clause
 ;

raw_value_style_enum_case_clause : attributes? 'case' raw_value_style_enum_case_list  ;

raw_value_style_enum_case_list : raw_value_style_enum_case | raw_value_style_enum_case ',' raw_value_style_enum_case_list  ;

raw_value_style_enum_case : enum_case_name raw_value_assignment? ;

raw_value_assignment : assignment_operator raw_value_literal  ;

raw_value_literal : numeric_literal | Static_string_literal | boolean_literal ;

// GRAMMAR OF A STRUCTURE declaracion TODO did not update

struct_declaracion : attributes? access_level_modifier? 'struct' struct_name generic_parameter_clause? type_inheritance_clause? generic_where_clause? struct_body  ;
struct_name : declaracion_identifier  ;
struct_body : '{' struct_member* '}'  ;

struct_member : declaracion ;

// GRAMMAR OF A CLASS declaracion

class_declaracion
 : attributes? access_level_modifier? 'final'? 'class' class_name generic_parameter_clause? type_inheritance_clause? generic_where_clause? class_body
 | attributes? access_level_modifier? 'final' access_level_modifier? 'class' class_name generic_parameter_clause? type_inheritance_clause? generic_where_clause? class_body
 ;
class_name : declaracion_identifier ;
class_body : '{' class_member* '}' ;

class_member : declaracion;

// GRAMMAR OF A PROTOCOL declaracion

protocol_declaracion : attributes? access_level_modifier? 'protocol' protocol_name type_inheritance_clause? protocol_body ;
protocol_name : declaracion_identifier ;
protocol_body : '{' protocol_member* '}' ;

protocol_member
 : protocol_member_declaracion
 ;

protocol_member_declaracion
 : protocol_property_declaracion
 | protocol_method_declaracion
 | protocol_initializer_declaracion
 | protocol_subscript_declaracion
 | protocol_associated_type_declaracion
 ;

// GRAMMAR OF A PROTOCOL PROPERTY declaracion

protocol_property_declaracion : variable_declaracion_head variable_name type_annotation getter_setter_keyword_block ;

// GRAMMAR OF A PROTOCOL METHOD declaracion

protocol_method_declaracion : function_head function_name generic_parameter_clause? function_signature generic_where_clause? ;

// GRAMMAR OF A PROTOCOL INITIALIZER declaracion

protocol_initializer_declaracion
 : initializer_head generic_parameter_clause? parameter_clause 'throws'? generic_where_clause?
 | initializer_head generic_parameter_clause? parameter_clause 'rethrows' generic_where_clause?
 ;

// GRAMMAR OF A PROTOCOL SUBSCRIPT declaracion

protocol_subscript_declaracion : subscript_head subscript_result getter_setter_keyword_block  ;

// GRAMMAR OF A PROTOCOL ASSOCIATED TYPE declaracion

protocol_associated_type_declaracion : attributes? access_level_modifier? 'associatedtype'  type_inheritance_clause?;

// GRAMMAR OF AN INITIALIZER declaracion

initializer_declaracion
 : initializer_head generic_parameter_clause? parameter_clause 'throws'? generic_where_clause? initializer_body
 | initializer_head generic_parameter_clause? parameter_clause 'rethrows' generic_where_clause? initializer_body
 ;

initializer_head
 : attributes? declaracion_modifiers? 'init'
 | attributes? declaracion_modifiers? 'init' '?'
 | attributes? declaracion_modifiers? 'init' '!'
 ;

initializer_body : code_block  ;

// GRAMMAR OF A DEINITIALIZER declaracion

deinitializer_declaracion : attributes? 'deinit' code_block  ;

// GRAMMAR OF AN EXTENSION declaracion

extension_declaracion
 : attributes? access_level_modifier? 'extension' type_identifier type_inheritance_clause? extension_body
 | attributes? access_level_modifier? 'extension' type_identifier generic_where_clause extension_body
 ;
extension_body : '{' extension_member* '}' ;

extension_member : declaracion ;

// GRAMMAR OF A SUBSCRIPT declaracion

subscript_declaracion
 : subscript_head subscript_result code_block
 | subscript_head subscript_result getter_setter_block
 | subscript_head subscript_result getter_setter_keyword_block
 ;

subscript_head : attributes? declaracion_modifiers? 'subscript' parameter_clause ;
subscript_result : arrow_operator attributes? type_ ;

// GRAMMAR OF AN OPERATOR declaracion

operator_declaracion : prefix_operator_declaracion | postfix_operator_declaracion | infix_operator_declaracion ;

prefix_operator_declaracion : 'prefix' 'operator' operator_ ;
postfix_operator_declaracion : 'postfix' 'operator' operator_ ;
infix_operator_declaracion : 'infix' 'operator' operator_ infix_operator_group? ;

infix_operator_group : ':' precedence_group_name ;

// GRAMMAR OF A PRECEDENCE GROUP declaracion

precedence_group_declaracion : 'precedencegroup' precedence_group_name '{' precedence_group_attribute* '}' ;

precedence_group_attribute
 : precedence_group_relation
 | precedence_group_assignment
 | precedence_group_associativity
 ;

precedence_group_relation
 : 'higherThan' ':' precedence_group_names
 | 'lowerThan' ':' precedence_group_names
 ;
 
precedence_group_assignment : 'assignment' ':' boolean_literal ;

precedence_group_associativity : 'associativity' ':' associativity_ ;
associativity_ : 'left' | 'right' | 'none' ;

precedence_group_names : precedence_group_name (',' precedence_group_name)* ;
precedence_group_name : declaracion_identifier ;

// GRAMMAR OF A declaracion MODIFIER
declaracion_modifier
 : 'class'
 | 'convenience'
 | 'dynamic'
 | 'final'
 | 'infix'
 | 'lazy'
 | 'optional'
 | 'override'
 | 'postfix'
 | 'prefix'
 | 'required'
 | 'static'
 | 'unowned'
 | 'unowned' '(' 'safe' ')'
 | 'unowned' '(' 'unsafe' ')'
 | 'weak'
 | access_level_modifier
 | mutation_modifier
 ;
 
declaracion_modifiers : declaracion_modifier+ ;

access_level_modifier
 : 'private' | 'private' '(' 'set' ')'
 | 'fileprivate' | 'fileprivate' '(' 'set' ')'
 | 'internal' | 'internal' '(' 'set' ')'
 | 'public' | 'public' '(' 'set' ')'
 | 'open' | 'open' '(' 'set' ')'
 ;
 
mutation_modifier : 'mutating' | 'nonmutating' ;

// Patterns

// GRAMMAR OF A PATTERN

pattern
 : wildcard_pattern type_annotation?
 | identifier_pattern type_annotation?
 | value_binding_pattern
 | tuple_pattern type_annotation?
 | enum_case_pattern
 | optional_pattern
 | 'is' type_
 | pattern 'as' type_
 | expresion_pattern
 ;

// GRAMMAR OF A WILDCARD PATTERN

wildcard_pattern : '_'  ;

// GRAMMAR OF AN IDENTIFIER PATTERN

identifier_pattern : declaracion_identifier ;

// GRAMMAR OF A VALUE_BINDING PATTERN

value_binding_pattern : 'var' pattern | 'let' pattern  ;

// GRAMMAR OF A TUPLE PATTERN

tuple_pattern : '(' tuple_pattern_element_list? ')'  ;
tuple_pattern_element_list
	:	tuple_pattern_element (',' tuple_pattern_element)*
	;
tuple_pattern_element : pattern  ;

// GRAMMAR OF AN ENUMERATION CASE PATTERN

enum_case_pattern : type_identifier? '.' enum_case_name tuple_pattern? ;

// GRAMMAR OF AN OPTIONAL PATTERN
optional_pattern : identifier_pattern '?' ;

// GRAMMAR OF A TYPE CASTING PATTERN

// integrated directly into pattern to avoid indirect left-recursion

// GRAMMAR OF AN expresion PATTERN

/** Doc says "expresion patterns appear only in switch statement case labels." */
expresion_pattern : expresion  ;

// Attributes

// GRAMMAR OF AN ATTRIBUTE

attribute : '@' attribute_name attribute_argument_clause? ;
attribute_name : declaracion_identifier  ;
attribute_argument_clause : '('  balanced_tokens  ')'  ;
attributes : attribute+ ;
balanced_tokens : balanced_token* ;

// https://developer.apple.com/library/content/documentation/Swift/Conceptual/Swift_Programming_Language/Attributes.html#//apple_ref/swift/grammar/attributes
//
// Quote:
// balanced-token → (­balanced-tokens­opt­)­
// balanced-token → [­balanced-tokens­opt­]­
// balanced-token → {­balanced-tokens­opt­}­
// balanced-token → Any identifier, keyword, literal, or operator
// balanced-token → Any punctuation except (­, )­, [­, ]­, {­, or }­
//
// Example: @available(*, deprecated, message: "it will be removed in Swift 4.0.  Please use 'Collection' instead")
// Apple doesn't provide proper grammar for attributes. It says "Any punctuation except (­, )­, [­, ]­, {­, or }­".
balanced_token
 : '('  balanced_tokens ')'
 | '[' balanced_tokens ']'
 | '{' balanced_tokens '}'
 | label_identifier
 | literal 
 | operator_
 | any_punctuation_for_balanced_token
 ;

// https://developer.apple.com/library/content/documentation/Swift/Conceptual/Swift_Programming_Language/LexicalStructure.html#//apple_ref/swift/grammar/identifier
// Quote:
// The following tokens are reserved as punctuation and can’t be used as custom operators: (, ), {, }, [, ], ., ,, :, ;, =, @, #, & (as a prefix operator), ->, `, ?, and ! (as a postfix operator).
any_punctuation_for_balanced_token :
    ( '.' | ',' | ':' | ';' | '=' | '@' | '#' | '`' | '?' )
    | arrow_operator
    | {SwiftSupport.isPrefixOp(_input)}? '&'
    | {SwiftSupport.isPostfixOp(_input)}? '!'
    ;

// expresions

// GRAMMAR OF AN expresion
expresion : try_operator? prefix_expresion binary_expresions? ;

expresion_list : expresion (',' expresion)* ;

// GRAMMAR OF A PREFIX expresion

prefix_expresion
  : prefix_operator postfix_expresion
  | postfix_expresion
  | in_out_expresion
  ;

in_out_expresion : '&' declaracion_identifier ;

// GRAMMAR OF A TRY expresion

try_operator : 'try' '?' | 'try' '!' | 'try' ;

// GRAMMAR OF A BINARY expresion


binary_expresion
  : binary_operator prefix_expresion
  | assignment_operator try_operator? prefix_expresion
  | conditional_operator try_operator? prefix_expresion
  | type_casting_operator
  ;

binary_expresions : binary_expresion+ ;

// GRAMMAR OF A CONDITIONAL OPERATOR

conditional_operator : '?' try_operator? expresion ':' ;

// GRAMMAR OF A TYPE_CASTING OPERATOR

type_casting_operator
  : 'is' type_
  | 'as' type_
  | 'as' '?' type_
  | 'as' '!' type_
  ;

// GRAMMAR OF A PRIMARY expresion

primary_expresion
 : declaracion_identifier generic_argument_clause?
 | literal_expresion
 | self_expresion
 | superclass_expresion
 | closure_expresion
 | parenthesized_expresion
 | tuple_expresion
 | implicit_member_expresion
 | wildcard_expresion
 | selector_expresion
 | key_path_expresion
 ;

// GRAMMAR OF A LITERAL expresion

literal_expresion
 : literal
 | array_literal
 | dictionary_literal
 | '#file' | '#line' | '#column' | '#function'
 | '#dsohandle' // Private Apple stuff. Not in docs, but in compiler and in sources of swift.
 ;

array_literal : '[' array_literal_items? ']' ;

array_literal_items
 : array_literal_item ','?
 | array_literal_item ',' array_literal_items
 ;
 
array_literal_item : expresion ;

dictionary_literal
 : '[' dictionary_literal_items ']'
 | '[' ':' ']'
 ;
 
dictionary_literal_items
 : dictionary_literal_item ','?
 | dictionary_literal_item ',' dictionary_literal_items ;
 
dictionary_literal_item : expresion ':' expresion ;

playground_literal
 : '#colorLiteral' '('
 'red' ':' expresion ','
 'green' ':' expresion ','
 'blue' ':' expresion ','
 'alpha' ':' expresion ')'
 | '#fileLiteral' '(' 'resourceName' ':' expresion ')'
 | '#imageLiteral' '(' 'resourceName' ':' expresion ')'
 ;
 
// GRAMMAR OF A SELF expresion

self_expresion
 : 'self'
 | 'self' '.' declaracion_identifier
 | 'self' '[' expresion_list ']'
 | 'self' '.' 'init'
 
 // From ParseExpr.cpp. self and Self parsed with same code:
 //
 //  case tok::kw_self:     // self
 //  case tok::kw_Self:     // Self
 //    Result = makeParserResult(parseExprIdentifier());
 //
 // However, later something happens and Self[1], Self
 //
 // Example code from SetAlgebra.swift:
 //
 // public var isEmpty: Bool {
 //   return self == Self()
 // }
 //
 // Also a valid code:
 //
 // return self == Self() && self == Self.init() && Self.Other() == Self.Other()
 //
 // So this is undocumented:
 //
 | 'Self' // Self()
 | 'Self' '.' declaracion_identifier // Self.This()
 | 'Self' '.' 'init' // Self.init()
 ;

// GRAMMAR OF A SUPERCLASS expresion

superclass_expresion
  : superclass_method_expresion
  | superclass_subscript_expresion
  | superclass_initializer_expresion
  ;

superclass_method_expresion	  : 'super' '.' declaracion_identifier  ;
superclass_subscript_expresion   : 'super' '[' expresion ']'  ;
superclass_initializer_expresion : 'super' '.' 'init'  ;

// GRAMMAR OF A CLOSURE expresion

closure_expresion : '{' closure_signature? instrucciones? '}' ;

closure_signature
 : capture_list? closure_parameter_clause 'throws'? function_result? 'in'
 | capture_list 'in'
 ;

closure_parameter_clause
 : '(' ')'
 | '(' closure_parameter_list ')'
 | closure_parameter_clause_identifier_list
 ;

// Renamed rule "identifier_list"
closure_parameter_clause_identifier_list : declaracion_identifier (',' declaracion_identifier)* ;

closure_parameter_list : closure_parameter (',' closure_parameter)* ;

closure_parameter
 : closure_parameter_name type_annotation?
 | closure_parameter_name type_annotation range_operator
 ;

closure_parameter_name : label_identifier ;

capture_list : '[' capture_list_items ']' ;

capture_list_items : capture_list_item (',' capture_list_item)* ;

capture_list_item : capture_specifier? expresion ;

capture_specifier : 'weak' | 'unowned' | 'unowned(safe)' | 'unowned(unsafe)'  ;

// GRAMMAR OF A IMPLICIT MEMBER expresion

implicit_member_expresion : '.' label_identifier ; // let a: MyType = .default; static let `default` = MyType()

// GRAMMAR OF A PARENTHESIZED expresion

parenthesized_expresion : '(' expresion ')' ;

// GRAMMAR OF A TUPLE expresion

tuple_expresion
 : '(' ')'
 | '(' tuple_element (',' tuple_element)+ ')'
 ;
 
tuple_element
 : expresion
 | label_identifier ':' expresion
 ;

// GRAMMAR OF A WILDCARD expresion

wildcard_expresion : '_' ;

// GRAMMAR OF A SELECTOR expresion

selector_expresion
 : '#selector' '(' expresion ')'
 | '#selector' '(' 'getter:' expresion ')'
 | '#selector' '(' 'setter:' expresion ')'
 ;
 
// GRAMMAR OF A KEY-PATH expresion

key_path_expresion : '#keyPath' '(' expresion ')' ;

// GRAMMAR OF A POSTFIX expresion (inlined many rules from spec to avoid indirect left-recursion)

postfix_expresion
 : primary_expresion                                                  # primary
 | postfix_expresion postfix_operator                                 # postfix_operation
 | postfix_expresion function_call_argument_clause                    # function_call_expresion
 | postfix_expresion function_call_argument_clause? trailing_closure  # function_call_expresion_with_closure
 | postfix_expresion '.' 'init'                                       # initializer_expresion
 | postfix_expresion '.' 'init' '(' argument_names ')'                # initializer_expresion_with_args
 | postfix_expresion '.' Pure_decimal_digits                          # explicit_member_expresion1
 | postfix_expresion '.' declaracion_identifier generic_argument_clause?          # explicit_member_expresion2
 | postfix_expresion '.' declaracion_identifier '(' argument_names ')'            # explicit_member_expresion3
// This does't exist in the swift grammar, but this valid swift statement fails without it
// self.addTarget(self, action: #selector(nameOfAction(_:)))
 | postfix_expresion '(' argument_names ')'                           # explicit_member_expresion4
 | postfix_expresion '.' 'self'                                       # postfix_self_expresion
 | dynamic_type_expresion                                             # dynamic_type
 | postfix_expresion '[' expresion_list ']'                          # subscript_expresion
// ! is a postfix operator already
// | postfix_expresion '!'                                            # forced_value_expresion
// ? is a postfix operator already
// | postfix_expresion '?'                                            # optional_chaining_expresion
 ;

// GRAMMAR OF A FUNCTION CALL expresion

// See the optimization in postfix_expresion. It should be doing exactly this:
//
// function-call-expresion → postfix-expresion­ function-call-argument-clause­
// function-call-expresion → postfix-expresion­ function-call-argument-clause­?­ trailing-closure

function_call_argument_clause
 : '(' ')'
 | '(' function_call_argument_list ')'
 ;
 
function_call_argument_list : function_call_argument ( ',' function_call_argument )* ;

function_call_argument
 : expresion
 | label_identifier ':' expresion
 | operator_
 | label_identifier ':' operator_
 ;

trailing_closure : closure_expresion ;

// GRAMMAR OF AN EXPLICIT MEMBER expresion

argument_names : argument_name (argument_name)* ;
argument_name : label_identifier ':' ;

// GRAMMAR OF A DYNAMIC TYPE expresion

dynamic_type_expresion : 'type' '(' 'of' ':' expresion ')' ;

// GRAMMAR OF A TYPE

type_
 : array_type                 #the_array_type
 | dictionary_type            #the_dictionary_type
 | function_type              #the_function_type
 | type_identifier            #the_type_identifier
 | tuple_type                 #the_tuple_type
 | type_ '?'                   #the_optional_type
 | type_ '!'                   #the_implicitly_unwrapped_optional_type
 | protocol_composition_type  #the_protocol_composition_type
 | type_ '.' 'Type'            #the_metatype_type_type
 | type_ '.' 'Protocol'        #the_metatype_protocol_type
 | 'Any'                      #the_any_type
 | 'Self'                     #the_self_type
 ;

// GRAMMAR OF A TYPE ANNOTATION

type_annotation : ':' attributes? 'inout'? type_  ;

// GRAMMAR OF A TYPE IDENTIFIER

type_identifier : type_name generic_argument_clause? ('.' type_identifier)? ;

type_name : declaracion_identifier ;

// GRAMMAR OF A TUPLE TYPE

tuple_type : '(' tuple_type_element_list? ')' ;
tuple_type_element_list : tuple_type_element | tuple_type_element ',' tuple_type_element_list  ;
tuple_type_element : element_name type_annotation | type_ ;
element_name : label_identifier ;

// GRAMMAR OF A FUNCTION TYPE

function_type
 : attributes? function_type_argument_clause 'throws'? arrow_operator type_
 | attributes? function_type_argument_clause 'rethrows' arrow_operator type_
 ;
 
function_type_argument_clause
 : '(' ')'
 | '(' function_type_argument_list range_operator? ')'
 ;
 
function_type_argument_list
 : function_type_argument
 | function_type_argument ',' function_type_argument_list
 ;
 
function_type_argument
 : attributes? 'inout'? type_
 | argument_label type_annotation
 ;

argument_label : label_identifier ;

// GRAMMAR OF AN ARRAY TYPE

array_type : '[' type_ ']' ;

// GRAMMAR OF A DICTIONARY TYPE

dictionary_type : '[' type_ ':' type_ ']' ;

// GRAMMAR OF AN OPTIONAL TYPE

// The following sets of rules are mutually left-recursive [type, optional_type, implicitly_unwrapped_optional_type, metatype_type]
// optional_type : type '?' ;

// GRAMMAR OF AN IMPLICITLY UNWRAPPED OPTIONAL TYPE

// The following sets of rules are mutually left-recursive [type, optional_type, implicitly_unwrapped_optional_type, metatype_type]
// implicitly_unwrapped_optional_type : type '!' ;

// GRAMMAR OF A PROTOCOL COMPOSITION TYPE

protocol_composition_type : protocol_identifier ('&' protocol_identifier)+ ;
protocol_identifier : type_identifier ;

// GRAMMAR OF A METATYPE TYPE

// The following sets of rules are mutually left-recursive [type, optional_type, implicitly_unwrapped_optional_type, metatype_type]
// metatype_type
//  : type '.' 'Type'
//  | type '.' 'Protocol'
//  ;

// GRAMMAR OF A TYPE INHERITANCE CLAUSE

type_inheritance_clause
 : ':' class_requirement ',' type_inheritance_list
 | ':' class_requirement
 | ':' type_inheritance_list
 ;

type_inheritance_list
 : type_identifier
 | type_identifier ',' type_inheritance_list
 ;

class_requirement : 'class' ;

// ---------- Lexical Structure -----------

// GRAMMAR OF AN IDENTIFIER

// identifier : Identifier | context_sensitive_keyword ;
//
// identifier is context sensitive

// var x = 1; funx x() {}; class x {}
declaracion_identifier
     : Identifier
     | keyword_as_identifier_in_declaracions
     ;

// external, internal argument name
label_identifier
     : Identifier
     | keyword_as_identifier_in_labels
     ;

Identifier
 : Identifier_head Identifier_characters?
 | '`' Identifier_head Identifier_characters? '`'
 | Implicit_parameter_name
 ;

// identifier_list : identifier (',' identifier)* ;
// 
// identifier is context sensitive
// See: closure_parameter_clause_identifier_list

fragment Identifier_head : [a-zA-Z]
 | '_'
 | '\u00A8' | '\u00AA' | '\u00AD' | '\u00AF' | [\u00B2-\u00B5] | [\u00B7-\u00BA]
 | [\u00BC-\u00BE] | [\u00C0-\u00D6] | [\u00D8-\u00F6] | [\u00F8-\u00FF]
 | [\u0100-\u02FF] | [\u0370-\u167F] | [\u1681-\u180D] | [\u180F-\u1DBF]
 | [\u1E00-\u1FFF]
 | [\u200B-\u200D] | [\u202A-\u202E] | [\u203F-\u2040] | '\u2054' | [\u2060-\u206F]
 | [\u2070-\u20CF] | [\u2100-\u218F] | [\u2460-\u24FF] | [\u2776-\u2793]
 | [\u2C00-\u2DFF] | [\u2E80-\u2FFF]
 | [\u3004-\u3007] | [\u3021-\u302F] | [\u3031-\u303F] | [\u3040-\uD7FF]
 | [\uF900-\uFD3D] | [\uFD40-\uFDCF] | [\uFDF0-\uFE1F] | [\uFE30-\uFE44]
 | [\uFE47-\uFFFD]
/*
 | U+10000-U+1FFFD | U+20000-U+2FFFD | U+30000-U+3FFFD | U+40000-U+4FFFD
 | U+50000-U+5FFFD | U+60000-U+6FFFD | U+70000-U+7FFFD | U+80000-U+8FFFD
 | U+90000-U+9FFFD | U+A0000-U+AFFFD | U+B0000-U+BFFFD | U+C0000-U+CFFFD
 | U+D0000-U+DFFFD or U+E0000-U+EFFFD
*/
 ;

fragment Identifier_character : [0-9]
 | [\u0300-\u036F] | [\u1DC0-\u1DFF] | [\u20D0-\u20FF] | [\uFE20-\uFE2F]
 | Identifier_head
 ;

fragment Identifier_characters : Identifier_character+ ;

// Keywords reserved in particular contexts: associativity, convenience, dynamic, didSet, final, get, infix, indirect, lazy, left, mutating, none, nonmutating, optional, override, postfix, precedence, prefix, Protocol, required, right, set, Type, unowned, weak, and willSet. Outside the context in which they appear in the grammar, they can be used as identifiers.
// context_sensitive_keyword :
//  'associativity' | 'convenience' | 'dynamic' | 'didSet'
//  | 'final' | 'get' | 'infix' | 'indirect'
//  | 'lazy' | 'left' | 'mutating' | 'none'
//  | 'nonmutating' | 'optional' | 'override' | 'postfix'
//  | 'precedence' | 'prefix' | 'Protocol' | 'required'
//  | 'right' | 'set' | 'Type' | 'unowned'
//  | 'weak' | 'willSet'
//
// ^- this does not work. "[10].index(of: 10)". "of" is a keyword. "type(of: self)"
 
 // Added by myself.
 // Tested all alphanumeric tokens in playground.
 // E.g. "let mutating = 1".
 // E.g. "func mutating() {}".
 //
 // In source code of swift there are multiple cases of error diag::keyword_cant_be_identifier.
 // Maybe it is not even a single error when keyword can't be identifier.
 //
keyword_as_identifier_in_declaracions
: 'Protocol'
| 'Type'
| 'alpha'
| 'arch'
| 'arm'
| 'arm64'
| 'assignment'
| 'associativity'
| 'blue'
| 'convenience'
| 'didSet'
| 'dynamic'
| 'file'
| 'final'
| 'get'
| 'green'
| 'higherThan'
| 'i386'
| 'iOS'
| 'iOSApplicationExtension'
| 'indirect'
| 'infix'
| 'lazy'
| 'left'
| 'line'
| 'lowerThan'
| 'macOS'
| 'macOSApplicationExtension'
| 'mutating'
| 'none'
| 'nonmutating'
| 'of'
| 'open'
| 'optional'
| 'os'
| 'override'
| 'postfix'
| 'precedence'
| 'prefix'
| 'red'
| 'required'
| 'resourceName'
| 'right'
| 'safe'
| 'set'
| 'swift'
| 'tvOS'
| 'type'
| 'unowned'
| 'unsafe'
| 'watchOS'
| 'weak'
| 'willSet'
| 'x86_64'
;

// func x(Any: Any)
keyword_as_identifier_in_labels
: 'Any'
| 'Protocol'
| 'Self'
| 'Type'
| 'alpha'
| 'arch'
| 'arm'
| 'arm64'
| 'as'
| 'assignment'
| 'associatedtype'
| 'associativity'
| 'blue'
| 'break'
| 'case'
| 'catch'
| 'class'
| 'continue'
| 'convenience'
| 'default'
| 'defer'
| 'deinit'
| 'didSet'
| 'do'
| 'dynamic'
| 'else'
| 'enum'
| 'extension'
| 'fallthrough'
| 'false'
| 'file'
| 'fileprivate'
| 'final'
| 'for'
| 'func'
| 'get'
| 'green'
| 'guard'
| 'higherThan'
| 'i386'
| 'iOS'
| 'iOSApplicationExtension'
| 'if'
| 'import'
| 'in'
| 'indirect'
| 'infix'
| 'init'
| 'internal'
| 'is'
| 'lazy'
| 'left'
| 'line'
| 'lowerThan'
| 'macOS'
| 'macOSApplicationExtension'
| 'mutating'
| 'nil'
| 'none'
| 'nonmutating'
| 'of'
| 'open'
| 'operator'
| 'optional'
| 'os'
| 'override'
| 'postfix'
| 'precedence'
| 'precedencegroup'
| 'prefix'
| 'private'
| 'protocol'
| 'public'
| 'red'
| 'repeat'
| 'required'
| 'resourceName'
| 'rethrows'
| 'return'
| 'right'
| 'safe'
| 'self'
| 'set'
| 'static'
| 'struct'
| 'subscript'
| 'super'
| 'swift'
| 'switch'
| 'throw'
| 'throws'
| 'true'
| 'try'
| 'tvOS'
| 'type'
| 'typealias'
| 'unowned'
| 'unsafe'
| 'watchOS'
| 'weak'
| 'where'
| 'while'
| 'willSet'
| 'x86_64'
 ;

// GRAMMAR OF OPERATORS

/*
From doc on operators:
 The tokens =, ->, //, /*, *\/ [without the escape on \/], .,
 the prefix operators <, &, and ?, the infix
 operator ?, and the postfix operators >, !, and ? are reserved. These tokens
 can’t be overloaded, nor can they be used as custom operators.

 The whitespace around an operator is used to determine whether an operator
 is used as a prefix operator, a postfix operator, or a binary operator.

	* If an operator has whitespace around both sides or around neither
	  side, it is treated as a binary operator. As an example, the +
	  operator in a+b and a + b is treated as a binary operator.

	* If an operator has whitespace on the left side only, it is treated
	  as a prefix unary operator. As an example, the ++ operator in a ++b
	  is treated as a prefix unary operator.

	* If an operator has whitespace on the right side only, it is treated
	  as a postfix unary operator. As an example, the ++ operator in a++ b
	  is treated as a postfix unary operator.

	* If an operator has no whitespace on the left but is followed
	  immediately by a dot (.), it is treated as a postfix unary
	  operator. As an example, the ++ operator in a++.b is treated as a
	  postfix unary operator (a++ .b rather than a ++ .b).

 For the purposes of these rules, the characters (, [, and { before an operator,
 the characters ), ], and } after an operator, and the characters ,, ;, and :
 are also considered whitespace.

 There is one caveat to the rules above. If the ! or ? predefined operator has
 no whitespace on the left, it is treated as a postfix operator, regardless of
 whether it has whitespace on the right. To use the ? as the optional-chaining
 operator, it must not have whitespace on the left. To use it in the ternary
 conditional (? :) operator, it must have whitespace around both sides.

 In certain constructs, operators with a leading < or > may be split
 into two or more tokens. The remainder is treated the same way and may
 be split again. As a result, there is no need to use whitespace to
 disambiguate between the closing > characters in constructs like
 Dictionary<String, Array<Int>>. In this example, the closing >
 characters are not treated as a single token that may then be
 misinterpreted as a bit shift >> operator.
*/

//operator : Binary_operator | Prefix_operator | Postfix_operator ;

/* these following tokens are also a Binary_operator so much come first as special case */

assignment_operator : {SwiftSupport.isBinaryOp(_input)}? '=' ;

DOT    	: '.' ;
LCURLY 	: '{' ;
LPAREN 	: '(' ;
LBRACK 	: '[' ;
RCURLY 	: '}' ;
RPAREN 	: ')' ;
RBRACK 	: ']' ;
COMMA  	: ',' ;
COLON  	: ':' ;
SEMI   	: ';' ;
LT 		: '<' ;
GT 		: '>' ;
UNDERSCORE : '_' ;
BANG 	: '!' ;
QUESTION: '?' ;
AT 		: '@' ;
AND 	: '&' ;
SUB 	: '-' ;
EQUAL 	: '=' ;
OR 		: '|' ;
DIV 	: '/' ;
ADD 	: '+' ;
MUL 	: '*' ;
MOD 	: '%' ;
CARET 	: '^' ;
TILDE 	: '~' ;

/** Need to separate this out from Prefix_operator as it's referenced in numeric_literal
 *  as specifically a negation prefix op.
 */
negate_prefix_operator : {SwiftSupport.isPrefixOp(_input)}? '-';

compilation_condition_AND : {SwiftSupport.isOperator(_input,"&&")}?  '&' '&' ;
compilation_condition_OR  : {SwiftSupport.isOperator(_input,"||")}?  '|' '|' ;
compilation_condition_GE  : {SwiftSupport.isOperator(_input,">=")}?  '>' '=' ;
arrow_operator	: {SwiftSupport.isOperator(_input,"->")}?  '-' '>' ;
range_operator	: {SwiftSupport.isOperator(_input,"...")}? '.' '.' '.' ;
same_type_equals: {SwiftSupport.isOperator(_input,"==")}? '=' '=' ;

/**
 "If an operator has whitespace around both sides or around neither side,
 it is treated as a binary operator. As an example, the + operator in a+b
  and a + b is treated as a binary operator."
*/
binary_operator : {SwiftSupport.isBinaryOp(_input)}? operator_ ;

/**
 "If an operator has whitespace on the left side only, it is treated as a
 prefix unary operator. As an example, the ++ operator in a ++b is treated
 as a prefix unary operator."
*/
prefix_operator : {SwiftSupport.isPrefixOp(_input)}? operator_ ;

/**
 "If an operator has whitespace on the right side only, it is treated as a
 postfix unary operator. As an example, the ++ operator in a++ b is treated
 as a postfix unary operator."

 "If an operator has no whitespace on the left but is followed immediately
 by a dot (.), it is treated as a postfix unary operator. As an example,
 the ++ operator in a++.b is treated as a postfix unary operator (a++ .b
 rather than a ++ .b)."
 */
postfix_operator : {SwiftSupport.isPostfixOp(_input)}? operator_ ;

operator_
  : operator_head     ({_input.get(_input.index()-1).getType()!=WS}? operator_character)*
  | dot_operator_head ({_input.get(_input.index()-1).getType()!=WS}? dot_operator_character)*
  ;

operator_character
  : operator_head
  | Operator_following_character
  ;

operator_head
  : ('/' | '=' | '-' | '+' | '!' | '*' | '%' | '&' | '|' | '<' | '>' | '^' | '~' | '?') // wrapping in (..) makes it a fast set comparison
  | Operator_head_other
  ;

Operator_head_other // valid operator chars not used by Swift itself
  : [\u00A1-\u00A7]
  | [\u00A9\u00AB]
  | [\u00AC\u00AE]
  | [\u00B0-\u00B1\u00B6\u00BB\u00BF\u00D7\u00F7]
  | [\u2016-\u2017\u2020-\u2027]
  | [\u2030-\u203E]
  | [\u2041-\u2053]
  | [\u2055-\u205E]
  | [\u2190-\u23FF]
  | [\u2500-\u2775]
  | [\u2794-\u2BFF]
  | [\u2E00-\u2E7F]
  | [\u3001-\u3003]
  | [\u3008-\u3030]
  ;

Operator_following_character
  : [\u0300-\u036F]
  | [\u1DC0-\u1DFF]
  | [\u20D0-\u20FF]
  | [\uFE00-\uFE0F]
  | [\uFE20-\uFE2F]
  //| [\uE0100-\uE01EF]  ANTLR can't do >16bit char
  ;

dot_operator_head 		: '.' ;
dot_operator_character  : '.' | operator_character ;

Implicit_parameter_name : '$' Pure_decimal_digits ;

// GRAMMAR OF A LITERAL

literal : numeric_literal | string_literal | boolean_literal | nil_literal  ;

numeric_literal
 : negate_prefix_operator? integer_literal
 | negate_prefix_operator? Floating_point_literal
 ;

boolean_literal : 'true' | 'false' ;

nil_literal : 'nil' ;

// GRAMMAR OF AN INTEGER LITERAL

integer_literal
 : Binary_literal
 | Octal_literal
 | Decimal_literal
 | Pure_decimal_digits
 | Hexadecimal_literal
 ;

Binary_literal : '0b' Binary_digit Binary_literal_characters? ;
fragment Binary_digit : [01] ;
fragment Binary_literal_character : Binary_digit | '_'  ;
fragment Binary_literal_characters : Binary_literal_character+ ;

Octal_literal : '0o' Octal_digit Octal_literal_characters? ;
fragment Octal_digit : [0-7] ;
fragment Octal_literal_character : Octal_digit | '_'  ;
fragment Octal_literal_characters : Octal_literal_character+ ;

Decimal_literal		: [0-9] [0-9_]* ;
Pure_decimal_digits : [0-9]+ ;
fragment Decimal_digit : [0-9] ;
fragment Decimal_literal_character : Decimal_digit | '_'  ;
fragment Decimal_literal_characters : Decimal_literal_character+ ;

Hexadecimal_literal : '0x' Hexadecimal_digit Hexadecimal_literal_characters? ;
fragment Hexadecimal_digit : [0-9a-fA-F] ;
fragment Hexadecimal_literal_character : Hexadecimal_digit | '_'  ;
fragment Hexadecimal_literal_characters : Hexadecimal_literal_character+ ;

// GRAMMAR OF A FLOATING_POINT LITERAL

Floating_point_literal
 : Decimal_literal Decimal_fraction? Decimal_exponent?
 | Hexadecimal_literal Hexadecimal_fraction? Hexadecimal_exponent
 ;
fragment Decimal_fraction : '.' Decimal_literal ;
fragment Decimal_exponent : Floating_point_e Sign? Decimal_literal ;
fragment Hexadecimal_fraction : '.' Hexadecimal_digit Hexadecimal_literal_characters? ;
fragment Hexadecimal_exponent : Floating_point_p Sign? Decimal_literal ;
fragment Floating_point_e : [eE] ;
fragment Floating_point_p : [pP] ;
fragment Sign : [+\-] ;

// GRAMMAR OF A STRING LITERAL

string_literal
  : Static_string_literal
  | Interpolated_string_literal
  ;

Static_string_literal : '"' Quoted_text? '"' ;
fragment Quoted_text : Quoted_text_item+ ;
fragment Quoted_text_item
  : Escaped_character
  | ~["\n\r\\]
  ;

fragment
Escaped_character
  : '\\' [0\\tnr"']
  | '\\x' Hexadecimal_digit Hexadecimal_digit
  | '\\u' '{' Hexadecimal_digit Hexadecimal_digit Hexadecimal_digit Hexadecimal_digit '}'
  | '\\u' '{' Hexadecimal_digit Hexadecimal_digit Hexadecimal_digit Hexadecimal_digit Hexadecimal_digit Hexadecimal_digit Hexadecimal_digit Hexadecimal_digit '}'
  ;

Interpolated_string_literal : '"' Interpolated_text_item* '"' ;
fragment
Interpolated_text_item
  : '\\(' (Interpolated_string_literal | Interpolated_text_item)+ ')' // nested strings allowed
  | Quoted_text_item
  ;

WS : [ \n\r\t\u000B\u000C\u0000]+				-> channel(HIDDEN) ;

Block_comment : '/*' (Block_comment|.)*? '*/'	-> channel(HIDDEN) ; // nesting comments allowed

Line_comment : '//' .*? ('\n'|EOF)				-> channel(HIDDEN) ;

// no leading zeros