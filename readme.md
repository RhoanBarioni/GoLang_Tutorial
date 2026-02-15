mod go init NAME

    Fazer isso, tu vai criar o go.mod com o nome fornecido no NAME. Se for criar módulos ou serviços para o main
    

VAR :=

    Isso faz com que a VAR seja auto declarada pelo Go
    Só pode ser usada dentro de funções

    nome := "Golang" // O Go entende que 'nome' é do tipo string
    idade := 10      // O Go entende que 'idade' é do tipo int

    ou tu escreve assim

    var nome string = "Golang"
    var idade int = 10

    := declara e atribui (cria uma variável nova).
    = atribui apenas (atualiza uma variável existente).


fmt.Sprintf("Hi, %v. Welcome!", VAR)

    Vai printar no console um texto + a variavél declarada após e vírgula.
    Usar o %v é como usar um := ele vai pegar o tipo da var automaticamente


    1. Os Mais Usados

        %v: O valor no formato padrão (muito versátil, serve para quase tudo).

        %+v: Quando usado com structs, ele mostra também o nome dos campos.

        %T: Mostra o tipo da variável (ex: string, int, main.User).

    2. Strings e Caracteres

        %s: Para strings simples.

        %q: Coloca a string entre aspas duplas (útil para debug).

        %c: Para um único caractere (runas).

    3. Números Inteiros

        %d: Inteiro na base 10 (decimal padrão).

        %b: Representação em binário.

        %x: Representação em hexadecimal.

    4. Números Decimais (Float)

        %f: Ponto flutuante padrão (ex: 1.234567).

        %.2f: Ponto flutuante com apenas 2 casas decimais.

        %e: Notação científica.

    5. Outros

        %t: Para booleanos (true ou false).

        %p: Para endereços de memória (ponteiros).


return formats[rand.Intn(len(formats))]

    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    tamanho = len(formats)
    indiceAleatorio := rand.Intn(tamanho)
    mensagem = formats[indiceAleatorio]
    return mensagem


slice

    []Tipo

    []int
    []string
    []float64

    é praticamente um array