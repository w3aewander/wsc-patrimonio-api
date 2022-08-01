const endpoint = "/api/"


$( (evt) => {
  
  $("#form-patrimonio").submit( (e) => {
    e.preventDefault();
    incluir(e)
  })

  listar()
})

const novo = (e) => {

  $("#patrimonio").val("")
  $("#tipo").val("")
  $("#modelo").val("")
  $("#observacao").val("")
  $("#patrimonio").focus( ()=>{
     $(this).style('border-color', "#ffaff")
  })
  
}
  
const excluir = (e) => {


  let patrimonio = e.getAttribute("data-patrimonio")


  fetch("/api/patrimonio/delete", {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json, charset=utf-8",
    },
    body: JSON.stringify({
      "Id": patrimonio
    })
  })
  .then(response => {
     response.json().then ( r => console.log(r))
     novo()
     listar()
  })
  .catch(err => {
    console.error(err);
  });


}

 

const incluir =  (e) =>  {


  let patrimonio = e.target.patrimonio.value
  let tipo = e.target.tipo.value
  let modelo = e.target.modelo.value
  let observacao = e.target.observacao.value

  console.log(`${patrimonio};${tipo};${modelo};${observacao}\n`)

 if ( `/api/patrimonio/${patrimonio}/exists` ){

  let patrimonio = e.target.patrimonio.value
  let tipo = e.target.tipo.value
  let modelo = e.target.modelo.value
  let observacao = e.target.observacao.value

  console.log(`${patrimonio};${tipo};${modelo};${observacao}\n`)

  
  fetch("/api/patrimonio/update", {
    method: "PUT",
    headers: {
      "Content-Type": "application/json, charset=utf-8",
    },
    body: JSON.stringify({
      "Id": patrimonio,
      "Tipo": tipo,
      "Modelo": modelo,
      "Observacao": observacao
    })
  })
  .then(response => {
     response.json().then ( r => console.log(r))
     novo()
     listar()
  })
  .catch(err => {
    console.error(err);
  });


} else {


  fetch("/api/patrimonio/add", {
    method: "POST",
    headers: {
      "Content-Type": "application/json, charset=utf-8",
    },
    body: JSON.stringify({
      "Id": patrimonio,
      "Tipo": tipo,
      "Modelo": modelo,
      "Observacao": observacao
    })
  })
  .then(response => {
     response.json().then ( r => console.log(r))
     novo()
     listar()
  })
  .catch(err => {
    console.error(err);
  });


}



}


const editar = (e) => {


  const registro =  e.parentNode.parentNode
  
  $('#patrimonio').val(   $(registro).attr('data-patrimonio') )
  $('#tipo').val(  $(registro).attr('data-tipo')  )
  $('#modelo').val(  $(registro).attr('data-modelo')  )
  $('#observacao').val(  $(registro).attr('data-observacao') )


}

const listar = async (e) => {

  const resp = await fetch(`${endpoint}patrimonios`)
  const  json = await resp.json()


  let registro = 0

  json.forEach(el => {

   // if ( registro > 0 ) { usar apenas se o csv tiver titulo
      console.log(el)

      html = `<tr data-patrimonio="${el.id}" data-tipo="${el.tipo}" data-modelo="${el.modelo}" data-observacao="${el.observacao}">
        <td>${el.id}</td>
        <td>${el.tipo}</td>
        <td>${el.modelo}</td>
        <td>${el.observacao}</td> 
        <td>
          <button type="button" class="btn btn-primary btn-sm" data-patrimonio="${el.id}" onclick="editar(this)"><i class="fa fa-edit"></i></button>
          <button type="button" class="btn btn-danger btn-sm" data-patrimonio="${el.id}" onclick="excluir(this)"><i class="fa fa-trash"></i></button>
        </td> 
      </tr>`
   // }  

      $('#tb-body').append(html)
  
      registro++

  });
  
  

}