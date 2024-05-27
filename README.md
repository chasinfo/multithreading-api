## Utilizar dois recursos: Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.
 
#### As duas requisições serão feitas simultaneamente para as seguintes APIs:
 
https://brasilapi.com.br/api/cep/v1/01153000 + cep
 
http://viacep.com.br/ws/" + cep + "/json/
 
#### Os requisitos para este desafio são:

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

#### :warning: <span style="color:red">Nos requisitos do desafio não diz como seria a requisição, se é através de um client API, ou através de linha de comando. Implementei como uma requisição através de linha de comando.</span> :warning:

#### Sobre o sistema 
Go Versão: 1.21

#### Utilização 

baixar o projeto
```shel
git clone https://github.com/chasinfo/multithreading-api.git
```

Executar o comando informando o CEP

Ex. go run main.go <numero do cep>
```shel
cd multithreading-api
go run main.go 72156209
```