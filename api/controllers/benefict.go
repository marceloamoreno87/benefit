package controllers

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	validatedoc "github.com/klassmann/cpfcnpj"
	benefitsv "github.com/marceloamoreno87/api/services/benefit"
	"google.golang.org/grpc"
)

type Benefit struct {
	Doc string `form:"doc" binding:"required" query:"doc"`
}

// GetBenfitByDoc             godoc
// @Summary      Get benefits by document (CPF)
// @Description  Return all benefits of the user by document (CPF) from http://extratoclube.com.br.
// @Tags         Get Benefits
// @Param        doc  query      string  true  "Type the cpf of the user who wants to get the benefits"
// @Produce      json
// @Router       /benefit [get]
func GetBenfitByDoc(c *gin.Context) {
	var benefit Benefit

	if err := c.ShouldBind(&benefit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "doc is required"})
		return
	}

	doc := validatedoc.NewCPF(benefit.Doc)

	if !doc.IsValid() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "doc is invalid"})
		return
	}

	conn, err := grpc.Dial(os.Getenv("BENEFIT_GRPC_URL"), grpc.WithInsecure())

	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()

	b := benefitsv.NewBenefitServiceClient(conn)

	resp, err := b.GetBenefit(context.Background(), &benefitsv.NewBenefit{Doc: benefit.Doc})
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"status": false, "error": err})
		return
	}

	c.JSON(200, resp)
}
