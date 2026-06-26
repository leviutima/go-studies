# Go do Zero — Doc 01: Fundamentos e Primeiro Programa

> Caderno de estudos de Go para quem vem de TypeScript/JavaScript.

---

## Objetivo desta doc

Sair do nada até ter um programa Go rodando no terminal, entendendo cada arquivo, cada comando e cada linha — sem mágica escondida (tendo em vista a abstração em linguagens como JS/TS).

---

## O que é o Go (mentalidade vinda do JS/TS)

O JS/TS (com React, Express, Nest) abstrai muita coisa pra você ser produtivo rápido. Go faz o oposto: ele é **explícito**. Você vê o que normalmente fica escondido.

> Não é Go ser difícil, é Go ser **honesto** sobre o que sempre esteve lá embaixo.

---

## Comandos do Go (o "npm" do Go)

| Comando | Equivalente JS | O que faz |
|---|---|---|
| `go mod init <nome>` | `npm init` | Cria o `go.mod` (equivalente ao `package.json`) |
| `go mod tidy` | `npm install` | Instala/limpa dependências |
| `go run caminho/arquivo.go` | `node arquivo.js` | Roda o arquivo diretamente |
| `go build` | — | Gera o executável |
| `go get <pkg>` | `npm install <pkg>` | Adiciona uma dependência |
| `go fmt` | Prettier | Formata o código |

---

## Criando o projeto do ZERO

### Passo 1 — Criar a pasta e entrar nela

```bash
mkdir go-studies
cd go-studies
```

### Passo 2 — Inicializar o módulo

```bash
go mod init go-studies
```

Isso cria o `go.mod`. O arquivo gerado fica assim:

```
module go-studies
# em projetos reais usa-se o caminho do repo:
# github.com/leviutima/go-studies
# esse caminho vira o prefixo de todos os imports internos

go 1.26
# versão major.minor do Go
```

### Passo 3 — Criar a estrutura da lição

```bash
mkdir 01-hello
```

Dentro dela, crie o arquivo `01-hello/main.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("módulo de estudos sobre golang")
}
```

### Passo 4 — Rodar

```bash
go run ./01-hello
```

**Saída esperada:**
```
módulo de estudos sobre golang
```

---

## Dissecando o programa linha por linha

### `package main`

Todo arquivo Go pertence a um **pacote** (uma caixa de código relacionado).
O pacote `main` é especial: é o único que vira um executável (um programa que roda). Os outros pacotes são bibliotecas.

> Pensa no `package main` como: *"este é o arquivo que pode ser ligado"*.

---

### `import "fmt"`

Go não tem nada global como o `console` do JS. Tudo que vem de fora precisa ser importado.

- `fmt` (de *format*) = biblioteca padrão pra imprimir e formatar texto. É o "console" do Go.
- **Import não usado = erro de compilação.** Go te obriga a remover lixo.

---

### `func main() { ... }`

A função `main` é o **ponto de entrada**: a primeira coisa que roda.
Nome obrigatório (`main`), sem parâmetros, sem retorno.

> No JS isso fica escondido pelo runtime.

---

### `fmt.Println("...")`

- `fmt.Println` = o `console.log` do Go. (*Println* = "print line": imprime e pula linha.)
- Sintaxe `pacote.Função()` → chama função de um pacote com ponto, igual `console.log` no JS.

---

## Variações de impressão

```go
fmt.Print("sem pular linha")           // não pula linha no fim
fmt.Println("pula linha no fim")       // pula linha (= console.log)
fmt.Printf("Oi %s, %d anos\n", n, i)  // com formatação (placeholders)
```

### Placeholders do `Printf`

| Placeholder | Tipo | Exemplo |
|---|---|---|
| `%s` | string | `"Levi"` |
| `%d` | número inteiro | `20` |
| `%f` | número decimal | `3.14` |
| `%v` | qualquer valor (genérico) | qualquer coisa |
| `%t` | booleano | `true` |
| `\n` | quebra de linha | — |

---

## Comparação direta JS ↔ Go

```js
// JavaScript
console.log("Cave Mode bootando...")
```

```go
// Go
package main

import "fmt"

func main() {
    fmt.Println("Cave Mode bootando...")
}
```

O JS parece menor porque **esconde** o `package`, o `import` e o `main`. Eles existem no JS também — o runtime monta pra você. Go te faz escrever na mão: mais verboso, porém **nada acontece por mágica**.

---

## Detalhes de sintaxe do Go

- **Sem ponto-e-vírgula** no fim das linhas (o compilador adiciona sozinho).
- **Sem aspas simples** pra string — use aspas duplas `"texto"`.
  - Aspas simples `'a'` são pra um único caractere, chamado **rune**.
- O `{` da função **tem que ficar na mesma linha** que o `func` — colocar em baixo dá erro (regra do Go).
- **Variável declarada e não usada = erro.** Mesma coisa pra `import`.
- O Go formata o código por você: rode `go fmt ./...` e ele arruma a indentação.
