package controllers

import (
	"context"
	"errors"
	"os"
	"strings"

	validatedoc "github.com/klassmann/cpfcnpj"

	"github.com/gin-gonic/gin"
	"github.com/marceloamoreno87/benefit/pkg/grpc"
	benefitsv "github.com/marceloamoreno87/benefit/proto"
)

type Benefit struct {
	Doc string `form:"doc" binding:"required" query:"doc"`
}

func (b *Benefit) ValidateBenefit() error {
	chunks := strings.Split(b.Doc, ",")
	for _, c := range chunks {
		if c == "" {
			return errors.New("Doc is required")
		}
		doc := validatedoc.NewCPF(c)
		if !doc.IsValid() {
			return errors.New("Doc is invalid")
		}
	}
	return nil
}

func NewBenefit() (*Benefit, error) {
	benefit := &Benefit{}
	return benefit, nil
}

// GetBenfitByDoc             godoc
// @Summary      Get benefits by document (CPF)
// @Description  Return all benefits of the user by document (CPF) from http://extratoclube.com.br.
// @Tags         Get Benefits
// @Param        doc  query      string  true  "Type the cpf of the user who wants to get the benefits"
// @Produce      json
// @Router       /benefit [get]
func GetBenfitByDoc(c *gin.Context) {
	benefit, err := NewBenefit()
	c.ShouldBind(&benefit)
	err = benefit.ValidateBenefit()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"status": false, "error": err.Error()})
		return
	}
	b := benefitsv.NewBenefitServiceClient(grpc.GetConnGrpc(os.Getenv("BENEFIT_GRPC_URL")))
	resp, err := b.GetBenefit(context.Background(), &benefitsv.NewBenefit{Doc: benefit.Doc})
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"status": false, "error": err})
		return
	}

	c.JSON(200, resp)
}
