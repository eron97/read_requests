package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type MeuObjeto struct {
	Campo1 string `json:"campo1"`
	Campo2 int    `json:"campo2"`
}

func manipularRequisicao(w http.ResponseWriter, r *http.Request) {

	// Acessando cabeçalhos
	cabecalhos := r.Header
	fmt.Println("Cabeçalhos:")
	for chave, valor := range cabecalhos {
		fmt.Printf("%s: %s\n", chave, valor)
	}

	// Acessando a linha de requisição
	metodo := r.Method
	uri := r.RequestURI
	fmt.Printf("Método: %s, URI: %s\n", metodo, uri)

	// Verifica se o método é POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Acessando o corpo da requisição
	corpo, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusInternalServerError)
		return
	}

	// Imprimindo o array de bytes no terminal
	fmt.Printf("Corpo da Requisição (Array de Bytes): %v\n", corpo)

	// Convertendo o corpo para string e imprimindo no terminal
	corpoString := string(corpo)
	fmt.Printf("Corpo da Requisição: %s\n", corpoString)

	/*
		// Decodifica o corpo da requisição JSON em uma estrutura Go
		var meuObjeto MeuObjeto
		err := json.NewDecoder(r.Body).Decode(&meuObjeto)
		if err != nil {
			http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
			return
		}

		// Agora você pode acessar os campos da estrutura
		fmt.Printf("Campo1: %s\n", meuObjeto.Campo1)
		fmt.Printf("Campo2: %d\n", meuObjeto.Campo2)
	*/

	// Faça algo com os dados...

	// Responda ao cliente
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Requisição processada com sucesso")
}

func main() {
	// Configura um manipulador para a rota "/exemplo"
	http.HandleFunc("/exemplo", manipularRequisicao)

	// Inicia o servidor na porta 8080
	fmt.Println("Servidor ouvindo em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

/* SAÍDA GERADA

Servidor ouvindo em http://localhost:8080
Cabeçalhos:
Content-Type: [application/json]
User-Agent: [PostmanRuntime/7.35.0]
Postman-Token: [855bc491-7ff6-4195-bf46-ca1b6f21fe0d]
Accept-Encoding: [gzip, deflate, br]
Connection: [keep-alive]
Content-Length: [69]
Método: POST, URI: /exemplo
Corpo da Requisição (Array de Bytes): [123 13 10 32 32 32 32 34 67 97 109 112 111 49 34 58 32 34 101 120 101 109 112 108 111 32 100 101 32 112 114 101 101 110 99 104 105 109 101 110 116 111 32 49 34 44 13 10 32 32 32 32 34 67 97 109 112 111 50 34 58 32 49 48 48 48 13 10 125]
Corpo da Requisição: {
	"Campo1": "exemplo de preenchimento 1",
	"Campo2": 1000
}

*/

/* OBS

123: Representa o caractere '{', que é o início de um objeto JSON.
13 10: Representam os caracteres de quebra de linha (CR e LF), indicando uma nova linha no formato CRLF (carriage return e line feed).
32: Representa o caractere de espaço.
34: Representa o caractere '"'.
Outros números: Representam os códigos ASCII dos caracteres no corpo do JSON.



Content-Type: [application/json]
Indica o tipo de mídia do corpo da requisição, neste caso, application/json, significando que o corpo da requisição está no formato JSON.

User-Agent: [PostmanRuntime/7.35.0]
Identifica o agente do usuário que fez a requisição, neste caso, é o PostmanRuntime na versão 7.35.0.

Postman-Token: [dee97ab6-e66a-4b6a-bbd3-ba67278b46ff]
Um token específico do Postman para identificar e rastrear a requisição.

Accept-Encoding: [gzip, deflate, br]
Indica os métodos de codificação que o cliente aceita para a resposta. Neste caso, aceita gzip, deflate e br (Brotli).

*/
