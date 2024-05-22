### Equipe: 
* Bernardo Mascarenhas Magalhães
* Douglas da Silva Ferreira
* Filipe Oliveira Saldanha da Gama
* Guilherme Duran Duran Gea
* Marcio Moreira do Nascimento Filho

# Como Utilizar o Programa

### Para iniciar os testes é necessário rodar a main, em seguida deve-se utilizar o Postman para fazer as requisições.

### Todas as mensagens geradas pelo programa estão de acordo com a data atual e o horário de Brasilia.

## Adicionar Produto
Com a main rodando, a primeira requisição será a de Adicionar Produto que utilizará um método POST através da URL http://localhost:8080/produto.

Para adicionar um produto terá que passar no body (raw) da requisição um json como neste [modelo](https://github.com/GuiDDuran/API-McRonalds-ProjetoAP1/blob/main/Modelos%20Json/JSON-Produtos.txt).

Caso o cadastro seja bem sucedido, como retorno receberá a mensagem "Produto cadastrado com sucesso!".

Caso o haja algum erro no momento de cadastro, o retorno será a mensagem "Erro ao decodificar dados do produto".


## Listar Produtos
Após adicionar um produto é possivel Listar Produtos enviando uma requisição que utilizará um método GET através da URL http://localhost:8080/produtos.

Os produtos serão listados em ordem alfabética, ou seja, de acordo com o nome do produto.

Caso haja algum produto cadastrado, como retorno receberá uma lista com os produtos cadastrados em formato json e a mensagem "Listando produtos".

Caso nenhum produto tenha sido cadastrado, e o usuário tente listar-los, o retorno será "Não há produtos cadastrados".

## Obter Produto Por Id
Além de poder Listar Produtos, pode-se Obter um Produto Por meio de seu Id usando um método GET através da URL http://localhost:8080/produto{id}.

Caso o usuário tente obter um produto com um id cadastrado, ele receberá a mensagem "Produto encontrado" e no body da requisição receberá as informações previamente cadastradas sobre o produto em formato json.

Caso o usuário esteja tentando obeter um id nao cadastrado no sistema, o retorno sera uma mensagem "Produto não encontrado".

Ou caso o id passado não seja correspondente, como por exemplo, tentar acessar por um caracter em vez de um inteiro, retornará a mensagem "ID inválido".

## Remover Produto Por Id
Assim como a funcao obter produtos, a funcão remover por meio de um id usa o método DELETE pela mesma URL http://localhost:8080/produto{id}.

Caso o usuário tente obter um produto com um id cadastrado, ele receberá a mensagem "Produto removido com sucesso" e no body da requisição receberá a mensagem "Produto removido com sucesso".

Caso o usuário esteja tentando remover um id nao cadastrado no sistema, o retorno sera uma mensagem "Produto não encontrado".

Ou caso o id passado não seja correspondente, como por exemplo, tentar acessar por um caracter em vez de um inteiro, retornará a mensagem "ID inválido".

## Adicionar Pedido
Para adicinar um pedido usa-se um método POST através da URL http://localhost:8080/pedido.

Para adicionar um produto terá que passar no body (raw) da requisição um json como neste [modelo](https://github.com/GuiDDuran/API-McRonalds-ProjetoAP1/blob/main/Modelos%20Json/JSON-Pedidos.txt).

Caso o cadastro seja bem sucedido, como retorno receberá a mensagem "Pedido cadastrado com sucesso!".

Caso haja algum erro no momento de cadastro, o retorno será a mensagem "Erro ao decodificar dados do pedido".

Caso o usuário tente tente fazer um pedido sem adicionar nenhum produto, como retorno receberá a mensagem "A lista de produtos não pode ser vazia".

## Obter Pedidos
Para obter os pedidos, usamos um método GET através da URL http://localhost:8080/pedidos.

O usuário pode listar todos os pedidos pela ordem de seu id caso não passe um parâmetro opcional na url.

Ou se desejar, pode listar os pedidos de acordo com o valor total, de forma crescente, ou seja, no menor valor para o maior.

Para isso o usuário deve passar algum dos seguintes valores na url:
- http://localhost:8080/pedidos?ordenacao=insertionsort
- http://localhost:8080/pedidos?ordenacao=bubblesort
- http://localhost:8080/pedidos?ordenacao=quicksort

Caso haja algum pedido cadastrado, como retorno receberá uma lista com os pedidos cadastrados em formato json e a mensagem "Obtendo pedidos".

Caso nenhum pedido tenha sido cadastrado, ou todos os pedidos já tenham sido expedidos, o retorno será "Não há pedidos ativos"

## Abrir Loja
Para abrir a loja e fazer com que os pedidos previamente cadastrados sejam expedidos, será necessária uma requisição de método POST através da URL http://localhost:8080/abrir.

Assim que a loja for aberta, caso o usuário não passe nenhum parâmetro opcional na url, os pedidos serão expedidos com um intervalo pré definido de 30 segundos.

Caso o usuário queira alterar o tempo deste intervalo, basta passar um parâmetro opcional na url indicando a duração do intervalo, em segundos, que deseja. Veja o exemplo abaixo:
- http://localhost:8080/abrir?intervalo=60

Caso o intervalo tenha sido alterado, o usuário irá receber a mensagem "Intervalo de expedição definido para {valor do intervalo} segundos" e após isso a mensagem "Loja aberta".

Os pedidos serão expedidos até que a loja feche ou apareça a mensagem "Não há pedidos ativos."

## Fechar Loja
Para fechar a loja e encerrar a expedição dos pedidos, será necessária uma requisição de método POST através da URL http://localhost:8080/fechar.

Após o fechamento da loja, a mensagem "Loja fechada" será retornada e a expedição dos pedidos irá encerrar.

## Obter Métricas
Para obter as métricas será necessário enviar uma requisição de método GET através da URL http://localhost:8080/metricas.

Após o envio da requisição, será retornado no body todas as informações referentes ás métricas em formato json e a mensagem "Obtendo métricas".

As métricas retornadas serão as seguintes: 
- Quantidade total de produtos cadastrados.
- Quantidade de pedidos encerrados.
- Quantidade de pedidos em andamento.
- Faturamento total.
- Ticket médio dos pedidos
- Tempo de funcionamento do programa.

## Fluxo de teste
1 - Adicione os produtos 1 e 2 disponíveis [aqui](https://github.com/GuiDDuran/API-McRonalds-ProjetoAP1/blob/main/Modelos%20Json/JSON-Produtos.txt).<br>
2 - Liste todos os produtos. ([Resposta esperada](https://github.com/GuiDDuran/API-McRonalds-ProjetoAP1/blob/main/Modelos%20Json/JSON-ListaProdutos1.txt))<br>
3 - Busque o produto 1 pelo seu id. ([Resposta esperada](https://github.com/GuiDDuran/API-McRonalds-ProjetoAP1/blob/main/Modelos%20Json/JSON-BuscaProduto1.txt))<br>
4 - Remova o produto 2 passando seu id. (Resposta esperada: Produto removido com sucesso)<br>
5 - Adicione o produto 3 disponível [aqui](https://github.com/GuiDDuran/API-McRonalds-ProjetoAP1/blob/main/Modelos%20Json/JSON-Produtos.txt).<br>
6 - Liste todos os produtos. ([Resposta esperada](https://github.com/GuiDDuran/API-McRonalds-ProjetoAP1/blob/main/Modelos%20Json/JSON-ListaProdutos2.txt))<br>
7 - Adicione o pedido 1 disponível [aqui](https://github.com/GuiDDuran/API-McRonalds-ProjetoAP1/blob/main/Modelos%20Json/JSON-Pedidos.txt).<br>
8 - Obtenha os pedidos. ([Resposta esperada](https://github.com/GuiDDuran/API-McRonalds-ProjetoAP1/blob/main/Modelos%20Json/JSON-ObterPedidos1.txt))<br>
9 - Obtenha as métricas. ([Resposta esperada](https://github.com/GuiDDuran/API-McRonalds-ProjetoAP1/blob/main/Modelos%20Json/JSON-ObterMetricas1.txt))<br>
10 - Abra a loja e espere 30 segundos.<br>
11 - Obtenha os pedidos (Resposta esperada: Nenhum pedido encontrado).<br>
12 - Obtenha as métricas. ([Resposta esperada](https://github.com/GuiDDuran/API-McRonalds-ProjetoAP1/blob/main/Modelos%20Json/JSON-ObterMetricas2.txt))<br>
13 - Adicione o pedido 2 disponível [aqui](https://github.com/GuiDDuran/API-McRonalds-ProjetoAP1/blob/main/Modelos%20Json/JSON-Pedidos.txt).<br>
14 - Obtenha os pedidos. ([Resposta esperada](https://github.com/GuiDDuran/API-McRonalds-ProjetoAP1/blob/main/Modelos%20Json/JSON-ObterPedidos3.txt))<br>
15 - Obtenha as métricas. ([Resposta esperada](https://github.com/GuiDDuran/API-McRonalds-ProjetoAP1/blob/main/Modelos%20Json/JSON-ObterMetricas3.txt))<br>
16 - Feche a loja<br>
