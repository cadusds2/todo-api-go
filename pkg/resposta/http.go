package resposta

import "github.com/gin-gonic/gin"

// Sucesso retorna resposta padrao de sucesso.
func Sucesso(ctx *gin.Context, status int, dados interface{}) {
	ctx.JSON(status, gin.H{"dados": dados})
}

// Erro retorna formato padrao de erro.
func Erro(ctx *gin.Context, status int, mensagem string) {
	ctx.AbortWithStatusJSON(status, gin.H{"erro": mensagem})
}
