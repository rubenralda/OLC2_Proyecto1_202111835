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
    | struct_declaracion
    | asignar_atributos ';'?
    ;

// GRAMATICA PARA UNA FUNCION

function_declaracion 
    : 'func' Identificador '(' lista_parametros? ')' ( '->' tipos)? code_block
    ;

lista_parametros
    : declaracion_parametro (',' declaracion_parametro)*
    ;

declaracion_parametro
    : (Identificador)? Identificador ':' refencia = 'inout'? tipos #declaracion_parametro_simple
    | (Identificador)? Identificador ':' refencia = 'inout'? '[' tipos ']' #declaracion_parametro_vector
    ;

// GRAMATICA BLOQUE DE CODIGO

code_block
    : '{' instrucciones '}'
    ;

// GRAMATICA PARA DECLARAR UN STRUCT

struct_declaracion
    : 'struct' Identificador '{' lista_atributos* '}'
    ;

lista_atributos 
    : tipo = ('let'| 'var') Identificador (':' tipos)? ('=' expresion)? ';'? #declarar_atributo
    | m='mutating'? function_declaracion #declarar_funcion_sc
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
   | llamadas_funciones ';'?
   | asignar_atributos ';'?
   ;

// GRAMATICA PARA DECLARAR TIPOS DE VARIBLE

declaracion
   : constant_declaracion
   | variable_declaracion
   | array_declaracion
   | matriz_declaracion
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
    | funcion_append
    | funcion_removeLast
    | funcion_removeat
    | funcion_int
    | funcion_float
    | funcion_string
    | llamada_normal
    | llamada_metodos
    ;

// GRAMATICA PARA LLAMADA DE FUNCIONES

llamada_normal
    : Identificador '(' lista_argumentos? ')'
    ;

lista_argumentos
    : declaracion_argumento (',' declaracion_argumento)*
    ;

declaracion_argumento
    : (Identificador ':')? r='&'? expresion
    ;

// GRAMATICA PARA LLAMADA DE METODOS

llamada_metodos
    : Identificador '.' Identificador '(' lista_argumentos? ')'
    ;

// GRAMATICA PARA ATRIBUTOS

atributos // hay que arreglar los atributos 
    : Identificador '.' 'isEmpty' #atributos_vector_empty
    | Identificador '.' 'count' #atributos_vector_count
    | ide_atributo ('.' ide_atributo)+ #atributos_generales
    ;

ide_atributo
    : Identificador #ide_atributo_simple
    | Identificador '[' expresion ']' #ide_atributo_vector
    ;
// GRAMATICA PARA ASIGNAR ATRIBUTOS

asignar_atributos
    : ide_atributo ('.' ide_atributo)+ '=' expresion #asignar_atributos_normal
    | ide_atributo ('.' ide_atributo)+ '+=' expresion #incre_atributos_normal
    | ide_atributo ('.' ide_atributo)+ '-=' expresion #decre_atributos_normal
    ;

// GRAMATICA PARA DECLARAR MATRICES

matriz_declaracion
    : 'var' Identificador ( ':' tipo_matriz )? '=' definicion_matriz
    ;

tipo_matriz
    : '[' tipo_matriz ']'
    | '[' tipos ']'
    ;

definicion_matriz
    : lista_valores_matriz
    | simple_matriz
    ;

lista_valores_matriz: '[' elementos_matriz ']';

elementos_matriz: elemento_matriz (',' elemento_matriz)*;

elemento_matriz: lista_valores_matriz | expresion;

simple_matriz
    : tipo_matriz '(' 'repeating' ':' simple_matriz ',' 'count' ':' Int ')'
    | tipo_matriz '(' 'repeating' ':' expresion ',' 'count' ':' Int ')'
    ;

// GRAMMAR OF A FOR_IN STATEMENT

for_in_statement : 'for' Identificador 'in' (expresion|rango) code_block ;

rango
    : expresion '...' expresion
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
    | Identificador
    ;

//declaracin de vectores

array_declaracion
    : 'var' Identificador ':' '[' tipos ']' definicion_vector #array_comun
    | 'var' Identificador '=' '[' tipos ']' '(' ')' #array_vacio
    ;

definicion_vector
    : '=' '[' lista_expresiones ']'
    | '=' '[' ']'
    | '=' Identificador
    ;

lista_expresiones : expresion (',' expresion)*
    ;

// GRAMATICA FUNCION PRINT

funcion_print 
    : 'print' '(' lista_expresiones ')'
    ;

// GRAMATICA FUNCION APPEND

funcion_append
    : Identificador '.' 'append' '(' expresion ')'
    ;

// GRAMATICA FUNCION REMOVELAST

funcion_removeLast
    : Identificador '.' 'removeLast' '(' ')'
    ;

// GRAMATICA FUNCION REMOVEAT

funcion_removeat
    : Identificador '.' 'remove' '(' 'at' ':' expresion ')'
    ;

// GRAMATICA FUNCION INT

funcion_int
    : 'Int' '(' expresion ')'
    ;

// GRAMATICA FUNCION FLOAT

funcion_float
    : 'Float' '(' expresion ')'
    ;

// GRAMATICA FUNCION STRING

funcion_string
    : 'String' '(' expresion ')'
    ;

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
    | atributos #expresion_atributos
    | Identificador '[' expresion ']' #expresion_vector
    | llamadas_funciones #expresion_llamada
    | Identificador '(' l_duble ')' #expresion_struct_dupla
    | Identificador #expresion_id
    | '(' expresion ')' #expresion_paren
    | op='!' expresion #expresion_nega
    | '-' expresion #expresion_unario
    | expresion op=('/'|'%'|'*') expresion #expresion_arit
    | expresion op=('+'|'-') expresion #expresion_arit
    | expresion op=('<'|'<='|'>='|'>'|'=='|'!=') expresion #expresion_compa
    | expresion op=('&&'|'||') expresion #expresion_rela
    ;

l_duble
    : Identificador ':' expresion (',' Identificador ':' expresion)*;

primitivos 
    : Caracter #primitivo_caracter 
    | Int #primitivo_int
    | Float #primitivo_float
    | String #primitivo_string
    | Bool #primitivo_bool
    | 'nil' #primitivo_nulo
    ;


Int : [0-9]+
    ;

Float : [0-9]+ ('.' [0-9]+)?
    ;

Caracter : '\'' Quoted_text_item '\'';

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
 : [a-zA-Z_][a-zA-Z0-9_]*
 ;