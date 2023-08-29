grammar T_swift;

// Tokens a ignorar
WS : [ \n\r\t\u000B\u000C\u0000]+				-> channel(HIDDEN) ;

Block_comment : '/*' (Block_comment|.)*? '*/'	-> channel(HIDDEN) ; // nesting comments allowed

Line_comment : '//' .*? ('\n'|EOF)				-> channel(HIDDEN) ;

inicio
   : instrucciones EOF
   ;

instrucciones
   : instruccion*
   ;

instruccion
   : declaracion ';'?
   | loop_statement ';'?
   | branch_statement ';'?
   | control_transfer_statement ';'?
   | funcion_print ';'?
   ;

loop_statement 
   : for_in_statement
   | while_statement
   ;

code_block
    : '{' instrucciones '}'
    ;

// GRAMMAR OF A FOR_IN STATEMENT

for_in_statement : 'for' Identificador 'in' expresion code_block ; //cambiar pattern a ID y falta rango

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

declaracion
   : constant_declaracion
   | variable_declaracion
   //| function_declaracion
   //| struct_declaracion
   | array_declaracion
   ;

constant_declaracion : 'let ' Identificador anotacion_tipo? '=' expresion
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
    ;

//declaracin de vectores

array_declaracion
    : 'var' Identificador ':' '[' tipos ']' definicion_vector
    ;

definicion_vector
    : '=' '[' lista_expresiones ']'
    | '=' '[' ']'
    | Identificador
    ;

lista_expresiones : expresion (',' expresion)*
    ;

// funcion print
funcion_print : 'print' '(' expresion ')';

expresion //agregar llamada de una funcion, de struct, vector y atributos
    : primitivos #valor_primitivo
    | Identificador #expresion_id
    | '-' expresion #expresion_unario
    | '(' expresion ')' #expresion_paren
    | right=expresion op=('*'|'/'|'%') left=expresion #expresion_arit
    | right=expresion op=('+'|'-') left=expresion #expresion_arit
    | right=expresion op=('=='|'!='|'>'|'<'|'<='|'>=') left=expresion #expresion_compa
    | op='!' left=expresion #expresion_nega
    | right=expresion op=('&&'|'||') left=expresion #expresion_rela
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