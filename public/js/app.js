const endpoint = "/api/"



$( (evt) => {
  
  $("#form-patrimonio").submit( (e) => {
    e.preventDefault();
    incluir(e)
  })

  if ( ! localStorage.getItem("_token_patrimonio") ) {

    alert("Por favor, realize o login para ")
    
    const formLogin = `<div class="card" style="display: block;margin: 50px auto;width: 450px;">
                            <div class="card-header">
                                 <div class="card-title">
                                     Autenticação Requerida
                                  </div>
                            </div>
                            <div class="card-body">
                                <input type="text" class="form-control" id="email" name="email" placeholder="email">
                                <input type="password" class="form-control" id="senha" email="senha" placeholder="senha">
                               
                                <div class="card-text">
                                   <div id="retorno" class="alert alert-info></div>
                                 </div>
                            </div>
                            <div class="card-footer">
                               <button  onclick="login();" class="btn btn-primary" id="btn-login">Entrar</button>
                            </div>
                         </div>`

    $("body").html(formLogin)


  } else {

     const json = JSON.parse( atob(localStorage.getItem("_token_patrimonio") ) )
     localStorage.setItem("email", json.email) 
     $("#userinfo").text(json.email)
     
  }

  listar()
})


const login =  (e) =>  {

  const email  = $("#email").val()
  const senha = $("#senha").val()
  let isValid = false

  const credentials = JSON.stringify({"email":email, "senha":senha})

  fetch('/api/app/login', {
      method: 'POST', 
      headers: {'Content-type': 'Application/json'},
      body: credentials
  })
    .then ( resp => resp.json() )
    .then ( resp => { 
    

     if (resp.Data ){
        
        localStorage.setItem("_token_patrimonio", btoa(credentials) )

        isValid = true
          $("#retorno").text("Ok, autenticado") 
          location.href="/api/app"
       } else {

        $("#retorno").text("Acesso negado")
        localStorage.removeItem("__token_patrimonio")
        localStorage.removeItem("email")
     }
  
  })

}

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
     location.reload()
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
     location.reload()
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
     location.reload()
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