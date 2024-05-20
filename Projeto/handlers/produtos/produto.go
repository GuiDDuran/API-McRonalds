package produtos_handlers

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    "Projeto/modelos/produto"
    "github.com/gorilla/mux"
    "time"
)

func logMessage(message string) {
    var Horario = time.Now().Format("02/01/2006 15:04:05")
    fmt.Printf("%s - %s\n", Horario, message)
}

func AdicionarProduto(w http.ResponseWriter, r *http.Request) {
    var novoProduto produto.Produto

    err := json.NewDecoder(r.Body).Decode(&novoProduto)
    if err != nil {
        logMessage("Erro ao decodificar dados do produto")
        http.Error(w, "Erro ao decodificar dados do produto", http.StatusBadRequest)
        return
    }

    produto.LProdutos.Adicionar(novoProduto.Nome, novoProduto.Descricao, novoProduto.Valor)

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "text/html; charset=UTF-8")
    logMessage("Produto cadastrado com sucesso!")
    w.Write([]byte("Produto cadastrado com sucesso!"))
}

func ObterProduto(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        logMessage("ID inválido")
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    produto, err := produto.LProdutos.BuscarPorId(id)
    if err != nil {
        logMessage("Produto não encontrado")
        http.Error(w, "Produto não encontrado", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    logMessage("Produto encontrado")
    json.NewEncoder(w).Encode(produto)
}

func RemoverProduto(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        logMessage("ID inválido")
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    err = produto.LProdutos.Remover(id)
    if err != nil {
        logMessage("Produto não encontrado")
        http.Error(w, "Produto não encontrado", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    logMessage("Produto removido com sucesso")
    fmt.Fprintln(w, "Produto removido com sucesso")
}

func ListarProdutos(w http.ResponseWriter, r *http.Request) {
    todosOsProdutos := produto.LProdutos.ListarProdutos()

    if len(todosOsProdutos) == 0 {
        logMessage("Não há produtos cadastrados")
        http.Error(w, "Não há produtos cadastrados", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    logMessage("Listando produtos")
    json.NewEncoder(w).Encode(todosOsProdutos)
}
