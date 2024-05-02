### Equipe: 
* Bernardo Mascarenhas Magalhães
* Douglas da Silva Ferreira
* Filipe Oliveira Saldanha da Gama
* Guilherme Duran Duran Gea
* Marcio Moreira do Nascimento Filho

## Como Utilizar o Programa
Para iniciar os testes é necessário rodar a main, em seguida deve-se utilizar o Postman para fazer as requisições.<br><br>

### Adicionar Produto
Com a main rodando, a primeira requisição será a de Adicionar Produto que utilizará um método POST através da URL http://localhost:8080/produto.<br>
Para adicionar um produto terá que passar no body (raw) da requisição um json como neste modelo.<br>
Caso o cadastro seja bem sucedido, como retorno receberá a mensagem "Produto cadastrado com sucesso!".<br>
Caso o haja algum erro no momento de cadastro, o retorno será a mensagem "Erro ao decodificar dados do produto".<br><br>

### Listar Produtos
Após adicionar um produto é possivel Listar Produtos enviando uma requisição que utilizará um método GET através da URL http://localhost:8080/produtos.<br>
Caso haja algum produto cadastrado, como retorno receberá uma lista com os produtos cadastrados em formato json.
Caso nenhum produto tenha sido cadastrado, o retorno será "Nenhum produto encontrado"

### Obter Produto Por Id
Além de poder Listar Produtos, pode-se Obter um Produto Por meio de seu Id usando um método get através da url  http://localhost:8080/produto{id}.<br>
Caso haja o usuário esteja tentando obeter um id nao cadastrado no sistema, o retorno sera uma mensagem "Produto não encontrado".<br>
Ou caso o id passado não seja correspondente como por exemplo tentar acessar por um character em vez de um inteiro retornará ID inválido.<br>

### Remover Produto Por Id
Assim como a funcao obter produtos, a funcão remover por meio de um id usa o método delete pela mesma url, ou seja,http://localhost:8080/produto{id}. <br>
Caso haja o usuário esteja tentando remover um id nao cadastrado no sistema, o retorno sera uma mensagem "Produto não encontrado".<br>
Ou caso o id passado não seja correspondente como por exemplo tentar acessar por um character em vez de um inteiro retornará ID inválido.<br>

### Adicionar Pedido
Para adicinar um pedido usa-se um método post através da url http://localhost:8080/pedido.<br>
Para adicionar um produto terá que passar no body (raw) da requisição um json como neste modelo.<br>
Caso o cadastro seja bem sucedido, como retorno receberá a mensagem "Pedido cadastrado com sucesso!".<br>
Caso o haja algum erro no momento de cadastro, o retorno será a mensagem "Erro ao decodificar dados do pedido".<br><br>
### Obter Pedidos
Para obter os pedidos, usamos um método get através da url  http://localhost:8080/pedidos
Caso haja algum pedido cadastrado, como retorno receberá uma lista com os pedidos cadastrados em formato json.
Caso nenhum produto tenha sido cadastrado, o retorno será "Nenhum pedido encontrado"


### Abrir Loja

### Fechar Loja

### Obter Métricas

## Fluxo de teste
1 - Adicione os produtos 1 e 2 disponíveis aqui.<br>
2 - Liste todos os produtos.<br>
3 - Busque o produto 1 pelo seu id.<br>
4 - Remova o produto 2 passando seu id.<br>
5 - Adicione o produto 3 disponível aqui.<br>
6 - Liste todos os produtos.<br>
7 - Adicione o pedido 1 disponível aqui.<br>
(...)
