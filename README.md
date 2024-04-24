
<p align="center">
  <img src="https://sancao.com.br/images/brasil-api.webp" alt="BrasilAPI"/>
</p>

<hr>
<br>

SDK desenvolvido em Golang, como o objetivo de facilitar a integração de sistemas em serviços, utilizando o [BrasilApi](https://brasilapi.com.br/).



### Importando para seu projeto Go

<br>

Em seu projeto, rode o comando:

<br>

~~~bash
    go get -u github.com/Marcelospegiorin/brasilapi-go-sdk
~~~

<br>

Após instalado, utilize dessa forma: 
~~~go
  import "https://github.com/Marcelospegiorin/brasilapi-go/{pacote}"
~~~


<br>

### Exemplos de uso

Exemplos de uso do SDK logo abaixo:

<br>

#### CEPs

Informações sobre CEP

~~~go
  import "github.com/Marcelospegiorin/brasilapi-go-sdk/cep"
~~~

**Consultando Cep:**

~~~go
  cepResult, err := cep.GetCep(01310100)
~~~

#### Bancos

**Consultando um banco especifico:**

É necessário saber o código do banco que deseja consultar, veja tabela de códigos aqui [Aqui](https://globalfinanceiro.com.br/codigos-ispb-e-compe-dos-bancos/). 

~~~go
  codigo := 260 // NUBANK S.A.

  bankResult, err := banks.GetOne(codigo)
~~~

<br>
