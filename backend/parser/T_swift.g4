grammar T_swift;

// Tokens a ignorar
WS : [ \n\r\t\u000B\u000C\u0000]+				-> channel(HIDDEN) ;

Block_comment : '/*' (Block_comment|.)*? '*/'	-> channel(HIDDEN) ; // nesting comments allowed

Line_comment : '//' .*? ('\n'|EOF)				-> channel(HIDDEN) ;

inicio
   : instrucciones_globales EOF
   ;

// GRAMATICA PARA SENTENCIAS QUE SE DECLARAN SOLO GLOBALES

instrucciones_globales: intruccion_global* ;

intruccion_global
    : declaracion ';'?
    | loop_statement ';'?
    | branch_statement ';'?
    | asignacion ';'?
    | llamadas_funciones ';'?
    | function_declaracion
     //| struct_declaracion
    ;

// GRAMATICA PARA UNA FUNCION

function_declaracion 
    : 'func' Identificador '(' lista_parametros? ')' ( '->' tipos)? code_block
    ;

lista_parametros
    : ',' (Identificador | '_')? ':' 'inout'? tipos lista_parametros
    | (Identificador | '_')? ':' 'inout'? tipos
    ;

// GRAMATICA BLOQUE DE CODIGO

code_block
    : '{' instrucciones '}'
    ;

// GRAMATICA PARA INSTRUCCIONES EN UN BLOQUE DE FUNCION

instrucciones 
   : instruccion*
   ;

instruccion
   : declaracion ';'?
   | loop_statement ';'?
   | branch_statement ';'?
   | control_transfer_statement ';'?
   | asignacion ';'?
   | llamadas_funciones
   ;

// GRAMATICA PARA DECLARAR TIPOS DE VARIBLE

declaracion
   : constant_declaracion
   | variable_declaracion
   //| matriz_declaracion
   | array_declaracion
   ;

// GRAMATICA PARA SENTENCIAS DE CICLOS

loop_statement 
   : for_in_statement
   | while_statement
   ;

// GRAMMAR OF A BRANCH STATEMENT

branch_statement : if_statement #declarar_if
 | guard_statement #declarar_guard
 | switch_statement #declarar_switch
 ;

// GRAMMAR OF A CONTROL TRANSFER STATEMENT

control_transfer_statement 
   : break_statement
   | continue_statement
   | return_statement
   ;

// GRAMATICA PARA LAS LLAMADAS DE FUNCIONES NATIVAS Y NUESTRAS

llamadas_funciones
    : funcion_print
    ;

// GRAMMAR OF A FOR_IN STATEMENT

for_in_statement : 'for' Identificador 'in' (expresion|rango) code_block ; //cambiar pattern a ID y falta rango

rango
    : Int '...' Int
    ;

// GRAMMAR OF A WHILE STATEMENT

while_statement : 'while' expresion code_block ;

// GRAMMAR OF AN IF STATEMENT

if_statement 
    : 'if' expresion code_block #if_simple
    | 'if' expresion code_block 'else' code_block #else_final
    | 'if' expresion code_block 'else' if_statement #siguiente_if
    ;

// GRAMMAR OF A GUARD STATEMENT

guard_statement : 'guard' expresion 'else' '{' instrucciones ('continue'|'break'|'return') '}';

// GRAMMAR OF A SWITCH STATEMENT

switch_statement 
    : 'switch' expresion '{' switch_case* default_case? '}'  
    ;

switch_case 
    : 'case' expresion ':' instrucciones 'break'?
    ;

default_case
    : 'default' ':' instrucciones 'break'?
    ;

// GRAMMAR OF A BREAK STATEMENT

break_statement : 'break';

// GRAMMAR OF A CONTINUE STATEMENT

continue_statement : 'continue';

// GRAMMAR OF A RETURN STATEMENT

return_statement : 'return' expresion? ;

constant_declaracion : 'let' Identificador anotacion_tipo? '=' expresion
    ;

variable_declaracion : 'var' Identificador anotacion_tipo '?'
    | 'var' Identificador anotacion_tipo? '=' expresion
    ;

anotacion_tipo : ':' tipos ;

tipos 
    : 'String'
    | 'Int'
    | 'Float'
    | 'Bool'
    | 'Character'
    ;

//declaracin de vectores

array_declaracion
    : 'var' Identificador ':' '[' tipos ']' definicion_vector
    ;

definicion_vector
    : '=' '[' lista_expresiones ']'
    | '=' '[' ']'
    | '=' Identificador
    ;

lista_expresiones : expresion (',' expresion)*
    ;

// funcion print
funcion_print : 'print' '(' expresion ')';

// gramatica de asignacion

asignacion
    : Identificador '=' expresion #asignacion_normal
    | Identificador '+=' expresion #asignacion_incremento
    | Identificador '-=' expresion #asignacion_decremento
    | Identificador '[' expresion ']' '=' expresion #asignacion_vector
    | Identificador '[' expresion ']' '+=' expresion #asignacion_incremento_vector
    | Identificador '[' expresion ']' '-=' expresion #asignacion_decremento_vector
    ;

expresion //agregar llamada de una funcion, de struct, matriz y atributos
    : primitivos #valor_primitivo
    | Identificador '[' expresion ']' #expresion_vector
    | Identificador #expresion_id
    | '(' expresion ')' #expresion_paren
    | op='!' expresion #expresion_nega
    | '-' expresion #expresion_unario
    | expresion op=('/'|'%'|'*') expresion #expresion_arit
    | expresion op=('+'|'-') expresion #expresion_arit
    | expresion op=('<'|'<='|'>='|'>'|'=='|'!=') expresion #expresion_compa
    | expresion op=('&&'|'||') expresion #expresion_rela
    ;

primitivos 
    : Caracter #primitivo_caracter 
    | Int #primitivo_int
    | Float #primitivo_float
    | String #primitivo_string
    | Bool #primitivo_bool
    ;


Int : [0-9]+
    ;

Float : [0-9]+ ('.' [0-9]+)?
    ;

Caracter : '"' Quoted_text_item '"';

String : '"' Quoted_text? '"'
    ;

fragment Quoted_text : Quoted_text_item+ ;
fragment Quoted_text_item
  : '\\' [0\\tnr"']
  | ~["\n\r\\]
  ;

Bool : 'true' 
    | 'false'
    ;

// gramatica para identificador
Identificador
 : Identifier_head Identifier_characters?
 ;

fragment Identifier_head : [a-zA-Z]
 | '_'
 ;

fragment Identifier_character : [0-9]
 | Identifier_head
 ;

fragment Identifier_characters : Identifier_character+ ;