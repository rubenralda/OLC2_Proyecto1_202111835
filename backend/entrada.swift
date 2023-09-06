
// constantes

// declaración de constantes
//Incorrecto, la constante debe tener un valor asignado
//let valor: String?
// correcto, declaración de una constante tipo Int con valor
let num1 = 10
var num10 = 10
//let num2:Int = 10.01 // Error: no se puede asignar un Float a un Int
let num3:Float = 10.2 // correcto
let num4 = "esto es una variable"; //correcto constante tipo String
let num5:Bool = true //correcto
//let .58 = 4; // debe ser error porque .58 no es un nombre válido
//let if = "10" // debe ser un error porque "if" es una palabra reservada
// error valor1 al ser una constante no puede cambiar su valor
//num1 = 20
num10 += 10
if num1 == num10 {
    print(num5)
} else if num5 {
    print("numero 5")
    var solo_segundo_if: Int = 5
    if false {
        solo_segundo_if += (num10 + num1)/2
    } else {
        //print("Sera el vacio")
    }
    print(solo_segundo_if)
} else {
    print("ninguna de antes")
}


let numero = 1
switch numero {
case 1:
    print("Uno")
    switch numero {
        case 1:
            print("Uno")
        case 2:
            print("Dos")
        case 3:
            print("Tres")
        default:
            print("Invalid day")
    }
case 2:
    print("Dos")
case 3:
    print("Tres")
default:
    print("Invalid day")
}
/* Salida esperada:
Dos
*/

var num = 10

while num != 0 {
    num -= 1
    print(num)
}
/* Salida esperada:
9
8
7
6
5
4
3
2
1
0
*/

// for que recorre un rango
for i in 1...5 {
    print(i)
}

// for que recorre una cadena
for c in "hola" {
    print(c)
    //c = "b"
}

var i = 0;
var j = i;
while i < 2 {
    if j == 0{
        i = 1;
        i += 1
        continue;
    }
    i += 1
}
print(i)
// i posee el valor de 2 al finalizar el ciclo

var m = 2
while (m <= 10) {
// la sentencia guard verifica si i es impar
    guard m % 2 == 0 else {
        m = m + 1
        continue
    }
    print(m)
    m = m + 1
}
/* Salida
2
4
6
8
10
*/


var p = 2
while(p <= 10) {
    guard p%3 == 0 else {
        if true {
            break; // salimos del while o del guard? en este caso se sale del while
        }
        print("no debe llega aqui")
        break //en este caso este sería el break que exige el guard
    }
    print("ni tampoco aqui")
}

//vector con valores
var vec1: [Int] = [10,20,30,40,50]
//vector vacío
var vec2: [Float] = []
//vector vacío
var vec3: [String] = [ ];
//se realiza una copia completa de vector
var copiaVec: [Int] = vec1
//print(5+-(-5.5)+10+5)
//Aceso a un elemento
let val: Int = vec1[3] + num1 - copiaVec[0] // val = 40

//asignación con []
vec1[1] = vec1[0]; //[10,10,30,40,50]
vec1[1] += vec1[0]; //[10,20,30,40,50]
vec1[1] -= vec1[0] + 5; //[10,5,30,40,50]

//imprime 0
print(vec2.count)
//inserta 100 al final
vec1.append(100) //[10,5,30,40,50,100]
//inserciones en vacíos
vec2.append(1.0) // [1.0]
vec3.append("cadena") // ["cadena"]
//elimina la primera posición
vec1.remove( at: 0); //[5,30,40,50,100]
//elimina la última posición
vec1.remove(at: vec1.count - 1); //[5,30,40,50]

struct StructArr {
    var datos: Float
}

struct CentroTuristico {
    var nombre: String
}

struct Carro {
    var placa: String
    var color: String
    var tipo: String
}

struct Personaje {
    var nombre: String
    var edad: Int
    var descripcion: String
    var carro: Carro
    var numeros: StructArr
}

print("*******************STRUCTS")

let centro1 = CentroTuristico(nombre: "Volcan de pacaya")
let centro2 = CentroTuristico(nombre: "Rio dulce")
let centro3 = CentroTuristico(nombre: "Laguna Luchoa")
let centro4 = CentroTuristico(nombre: "Playa Blanca")
let centro5 = CentroTuristico(nombre: "Antigua Guatemala")
let centro6 = CentroTuristico(nombre: "Lago de Atitlan")

print("El nombre del Centro turistico 1 es: ", centro1.nombre)
print("El nombre del Centro turistico 2 es: ", centro2.nombre)
print("El nombre del Centro turistico 3 es: ", centro3.nombre)
print("El nombre del Centro turistico 4 es: ", centro4.nombre)
print("El nombre del Centro turistico 5 es: ", centro5.nombre)
print("El nombre del Centro turistico 6 es: ", centro6.nombre)

print("*******************CREANDO STRUCTS COMPUESTO")

let newCarro = Carro(placa: "090PLO", color: "gris", tipo: "mecanico")
var nums = StructArr(datos: 0.0)

var p1 = Personaje(
    nombre: "Jose",
    edad: 18,
    descripcion: "No hace nada",
    carro: newCarro,
    numeros: nums
)

print("Persona nombre: ", p1.nombre, ", edad: ", p1.edad, ", carroTipo: ", p1.carro.tipo, ", numeros: ", p1.numeros.datos)

var nums2 = StructArr(datos: Float("23.43"))
p1.numeros = nums2

print("Persona nombre: ", p1.nombre, ", edad: ", p1.edad, ", carroTipo: ", p1.carro.tipo, ", nuevos numeros: ", p1.numeros.datos)
