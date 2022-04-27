package rest

import (
	"github.com/gin-gonic/gin"
	"test-fizz-buzz/dto"
	"test-fizz-buzz/pkg"
)

func (s *Server) GenerateListOfStrings(c *gin.Context) {
	var req dto.Request
	if err := c.ShouldBind(&req); err != nil {
		BindJsonErr(c, err)
		return
	}

	if err := pkg.ValidateRequest(req); err != nil {
		BindJsonErr(c, err)
		return
	}
	result := FizzBuzz(req.FirstString, req.SecondString, req.FirstNumber, req.SecondNumber, req.Limit)
	ResponseData(c, result)
}
