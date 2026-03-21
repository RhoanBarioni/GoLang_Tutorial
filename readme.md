# 📌 Estrutura de Projeto em Go

## 🧠 Conceitos principais

- **Package (pacote):** cada pasta é um pacote  
- **main.go:** ponto de entrada da aplicação  
- **go.mod:** define o módulo do projeto  
- **Separação de responsabilidades:** cada parte do código tem um papel claro  

---

## 🗂️ Estrutura recomendada

```text
meu-projeto/
├── go.mod
├── cmd/
│   └── app/
│       └── main.go
│
├── internal/
│   ├── models/   → estruturas de dados (structs)
│   ├── service/  → regras de negócio (o que o sistema faz)
│   └── utils/    → funções auxiliares e validações
````

---

## 🧰 UTILS

Ferramentas, função auxiliar, ajuda técnica → validar um número

---

## ⚙️ SERVICE

Decisões, regra do projeto, o que o sistema vai fazer → aprovar um aluno

---

# 📌 Descrição das pastas

## 🔹 cmd/

Contém os pontos de entrada da aplicação.

Exemplo:

```text
cmd/api/main.go
cmd/cli/main.go
```

---

## 🔹 internal/

Código interno da aplicação.

Não pode ser importado por outros projetos
Usado para proteger a lógica do sistema

---

## 🔹 models/

Define as estruturas de dados (structs).

Responsável por:

* Representar entidades (User, Product, etc.)
* Mapear dados (JSON, banco, etc.)

Exemplo:

```go
type User struct {
    ID    int
    Name  string
    Email string
}
```

---

## 🔹 service/

Contém as regras de negócio.

Responsável por:

* Processamento
* Validações
* Lógica da aplicação

Exemplo:

```go
func CreateUser(u models.User) error {
    if u.Email == "" {
        return errors.New("email obrigatório")
    }
    return nil
}
```

---

## 🔹 utils/

Funções auxiliares reutilizáveis.

Exemplos:

* Formatação
* Validações simples
* Helpers genéricos

⚠️ Evitar colocar regras de negócio aqui.

---

# 🔄 Fluxo da aplicação

main → service → models → utils

---

# ⚙️ go mod init NAME

Fazer isso cria o go.mod com o nome fornecido no NAME.
Se for criar módulos ou serviços para o main, pode usar sites para importar modules → example.com/greetings

---

# ⚙️ go mod tidy

Go vai buscar possíveis packages faltando e vai resolver, atualizando o go.mod.
É como um “fix dependencies”.

---

# 🧪 VAR :=

Isso faz com que a variável seja auto declarada pelo Go.
Só pode ser usada dentro de funções.

```go
nome := "Golang" // Go infere string
idade := 10      // Go infere int
```

ou:

```go
var nome string = "Golang"
var idade int = 10
```

* `:=` declara e atribui (cria variável nova)
* `=` apenas atualiza variável existente

---

# 🖨️ Diferença entre Println e Sprintln

* `Println` mostra direto no console
* `Sprintln` armazena a string formatada na memória

---

# 🧾 fmt.Sprintf

```go
fmt.Sprintf("Hi, %v. Welcome!", VAR)
```

Vai juntar texto + variável.

* `%v` pega o valor automaticamente
* `\n` quebra linha

---

## 1. Os mais usados

* `%v` → valor padrão
* `%+v` → struct com campos
* `%T` → tipo

---

## 2. Strings e caracteres

* `%s` → string
* `%q` → string com aspas
* `%c` → caractere

---

## 3. Inteiros

* `%d` → decimal
* `%b` → binário
* `%x` → hexadecimal

---

## 4. Float

* `%f` → float
* `%.2f` → 2 casas decimais
* `%e` → notação científica

---

## 5. Outros

* `%t` → boolean
* `%p` → ponteiros

---

# 📥 import

```go
import "NOME"
```

ou:

```go
import (
    "NOME1"
    "NOME2"
)
```

---

# 🎲 return formats[rand.Intn(len(formats))]

```go
formats := []string{
    "Hi, %v. Welcome!",
    "Great to see you, %v!",
    "Hail, %v! Well met!",
}

tamanho := len(formats)
indiceAleatorio := rand.Intn(tamanho)
mensagem := formats[indiceAleatorio]
return mensagem
```

---

# 📦 slice

```go
[]int
[]string
[]float64
```

É praticamente um array.

---

# ⌨️ fmt.Scan

Usado para capturar entrada do usuário.

```go
fmt.Scan(&VAR)
```