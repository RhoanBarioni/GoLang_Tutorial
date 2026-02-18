go mod init NAME

    Fazer isso, tu vai criar o go.mod com o nome fornecido no NAME. Se for criar módulos ou serviços para o main
    Pode usar sites para importar os modules -> example.com/greetings


go mod tidy

    Go vai buscar possiveis packages faltantando e vai resolver dando import no go.mod é como um fix this shit


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


Diferença entre Println e Sprintln

    Println já mostra logo de cara no console, enquanto o Sprintln tu vai ter que declarar um var recebendo o valor dela
    e usar o Println para mostrar esse Sprintln

    O Sprintln armazena o valor na memória e faz formatação de string


fmt.Sprintf("Hi, %v. Welcome!", VAR)

    Vai printar no console um texto + a variavél declarada após e vírgula.
    Usar o %v é como usar um := ele vai pegar o tipo da var automaticamente
    Usar o \n para dar a quebra de linha


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
        

import

    o padrão é usar apenar import "NOME"
    porém se for usar outras libs, usar com 
    import(
        "NOME1"
        "NOME2"
    )


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