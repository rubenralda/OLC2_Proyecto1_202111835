const lineNumbers = document.querySelector('.line-numbers')
const btnEjecutar = document.getElementById("btn");
const codigo = document.getElementById("codigo");//textarea
const lineNumbersConsole = document.querySelector('.line-numbers-consola')
const consola = document.getElementById("textConsola");//textarea
const btnArbol = document.getElementById("arbolAst");
const btnTabla = document.getElementById("tablaSimbolos");
const btnAbrir = document.getElementById("abriArchivo");
const inputFile = document.createElement('input');
const btnErrores = document.getElementById('reporteErrores');
//const listArchivos = document.getElementById("listDeArchivos");
inputFile.type = 'file';
inputFile.accept = '.swift';
inputFile.style.display = 'none';
document.body.appendChild(inputFile);
let errores = null

document.addEventListener("DOMContentLoaded", (e) => {
    e.preventDefault;
    if ( localStorage.getItem("texto") != null) {
        codigo.value = localStorage.getItem("texto")
        let dispararEvento = new Event("keyup")
        codigo.dispatchEvent(dispararEvento)
    }
    e.stopPropagation;
});

codigo.addEventListener("input", (e) =>{
    e.preventDefault;
    localStorage.setItem("texto", codigo.value)
    e.stopPropagation;
});

codigo.addEventListener("keyup", event => {
    const numberOfLines = event.target.value.split("\n").length;
    lineNumbers.innerHTML = Array(numberOfLines)
    .fill('<span></span>')
    .join('')
});

codigo.addEventListener("keydown", function(event) {
    if (event.key === "Tab") {
      event.preventDefault();  // Evita que el foco cambie al siguiente elemento
      const tabCharacter = "    "; // 4 espacios para simular un tabulador
      const start = this.selectionStart;
      const end = this.selectionEnd;
      this.value = this.value.substring(0, start) + tabCharacter + this.value.substring(end);
      this.selectionStart = this.selectionEnd = start + tabCharacter.length;
    }
});

consola.addEventListener("keyup", event => {
    const numberOfLines = event.target.value.split("\n").length;
    lineNumbersConsole.innerHTML = Array(numberOfLines)
    .fill('<span></span>')
    .join('')
});

btnEjecutar.addEventListener("click", eventbtnEjecutar);
async function eventbtnEjecutar() {
    const ruta = `http://localhost:3000/ejecutar`;
    let bodyJson = {
        Codigo : codigo.value
    }
    const respuesta = await fetch(ruta,{
        method: 'POST', // or 'PUT'
        body: JSON.stringify(bodyJson), // data can be `string` or {object}!
        headers:{
        'Content-Type': 'application/json'
        }
    })
    .then((res)=> res.json())
    .then((data) => {
        return data
    })
    console.log(respuesta)
    if (respuesta.Err) {
        alert("Ocurrio un error")
        errores = respuesta.Reporte_error
        return
    }
    consola.value = respuesta.Salida
    let dispararEvento = new Event("keyup")
    consola.dispatchEvent(dispararEvento)
}

btnArbol.addEventListener("click", obtenerArbol);
async function obtenerArbol() {
    const ruta = `http://localhost:3000/ast`;
    let bodyJson = {
        codigo : codigo.value
    }
    const respuesta = await fetch(ruta,{
        method: 'POST', // or 'PUT'
        body: JSON.stringify(bodyJson), // data can be `string` or {object}!
        headers:{
        'Content-Type': 'application/json'
        }
    })
    .then((res)=> res.json())
    .then((data) => {
        return data
    })
    localStorage.setItem("reporte", respuesta.arbol)
    //window.location.href = "./html/reporte.html";
    window.open(`./html/reporte.html`, "_blank");
    console.log(respuesta)
}

btnTabla.addEventListener("click", obtenerTabla);
async function obtenerTabla() {
    const ruta = `http://localhost:3000/simbolos`;
    const respuesta = await fetch(ruta,{
        method: 'GET', // or 'PUT'
    })
    .then((res)=> res.json())
    .then((data) => {
        return data
    })
    localStorage.setItem("reporte", respuesta.tabla)
    window.open(`./html/reporte.html`, "_blank");
    console.log(respuesta)
    //window.location.href = "./html/reporte.html"
}

inputFile.addEventListener('change', () => {
    const file = inputFile.files[0];
    const reader = new FileReader();
    reader.readAsText(file);
    reader.onload = () => {
      const fileContent = reader.result;
      codigo.value = fileContent
      let dispararEvento = new Event("keyup")
      codigo.dispatchEvent(dispararEvento)
    };
});

btnAbrir.addEventListener("click", (e) =>{
    e.preventDefault();
    inputFile.click();
    e.stopPropagation();
});

btnErrores.addEventListener("click", (e) =>{
    e.preventDefault();
    if (errores == null) {
        alert("No hay errores")
        return
    }
    let cuerpo = "digraph erres {\n"
                + "    node [shape=plaintext]\n"
                + "\n"
                + "    tbl [\n"
                + "        label=<\n"
                + "            <table border=\"0\" cellborder=\"1\" cellspacing=\"0\">"
                + "<tr><td bgcolor=\"/rdylgn11/5:/rdylgn11/5\"><b>#</b></td>"
                + "<td bgcolor=\"/rdylgn11/5:/rdylgn11/5\"><b>Tipo de error</b></td>"
                + "<td bgcolor=\"/rdylgn11/5:/rdylgn11/5\"><b>Descripcion</b></td>"
                + "<td bgcolor=\"/rdylgn11/5:/rdylgn11/5\"><b>Linea</b></td>"
                + "<td bgcolor=\"/rdylgn11/5:/rdylgn11/5\"><b>Columna</b></td>"
                + "</tr>";
    for (let index = 0; index < errores.length; index++) {
        cuerpo += `<tr>
        <td>${index + 1}</td>
        <td>${errores[index].tipo}</td>
        <td>${errores[index].mensaje}</td>
        <td>${errores[index].linea}</td>
        <td>${errores[index].columna}</td>
        </tr>`;
    }
    cuerpo += "</table>\n"
                + "        >\n"
                + "    ];\n"
                + "}";
    localStorage.setItem("reporte", cuerpo)
    window.open(`./html/reporte.html`, "_blank");
    e.stopPropagation();
});