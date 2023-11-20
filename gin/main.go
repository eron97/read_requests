package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.POST("/exemplo", func(c *gin.Context) {

		// Acessando cabeçalhos da requisição
		contentType := c.GetHeader("Content-Type")
		userAgent := c.GetHeader("User-Agent")
		token := c.GetHeader("Postman-Token")
		accepetEncoding := c.GetHeader("Accept-Encoding")

		fmt.Printf("Content-Type: %s\n", contentType)
		fmt.Printf("User-Agent: %s\n", userAgent)
		fmt.Printf("Content-Type: %s\n", token)
		fmt.Printf("Content-Type: %s\n", accepetEncoding)

		requestLine := c.Request.Method + " " + c.Request.RequestURI
		fmt.Printf("Linha de Requisição: %s\n", requestLine)

		metodo := c.Request.Method
		fmt.Printf("Método: %s\n", metodo)

		// Acessando o corpo da requisição como array de bytes
		corpo, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao ler o corpo da requisição"})
			return
		}

		// Imprimindo o array de bytes do corpo
		fmt.Printf("Corpo da Requisição (Array de Bytes): %v\n", corpo)

		// Convertendo o array de bytes para string e imprimindo
		corpoString := string(corpo)
		fmt.Printf("Corpo da Requisição (String): %s\n", corpoString)

		c.JSON(http.StatusOK, gin.H{
			"mensagem": "Requisição POST recebida com sucesso!",
		})
	})

	r.GET("/exemplo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mensagem": "GET ok",
		})
	})

	r.Run()

}

/* SAÍDA GERADA

Content-Type: application/json
User-Agent: PostmanRuntime/7.35.0
Content-Type: 91dd9d32-41b3-45c0-8dcf-ede26614a166
Content-Type: gzip, deflate, br
Linha de Requisição: POST /exemplo
Método: POST
Corpo da Requisição (Array de Bytes): [123 13 10 32 32 32 32 34 67 97 109 112 111 49 34 58 32 34 101 120 101 109 112 108 111 32 100 101 32 112 114 101 101 110 99 104 105 109 101 110 116 111 32 49 34 44 13 10 32 32 32 32 34 67 97 109 112 111 50 34 58 32 49 48 48 48 44 13 10 32 32 32 32 34 67 97 109 112 111 50 51 34 58 32 49 48 48 48 13 10 125]
Corpo da Requisição (String): {
    "Campo1": "exemplo de preenchimento 1",
    "Campo2": 1000,
    "Campo23": 1000
}

*/
