package rest

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"test-fizz-buzz/dto"
	"test-fizz-buzz/pkg"
)

var (
	m = make(map[string]int)
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

	var params []string
	params = append(params, req.FirstString, req.SecondString, strconv.Itoa(req.FirstNumber), strconv.Itoa(req.SecondNumber), strconv.Itoa(req.Limit))
	res := strings.Join(params, ";")

	val, ok := m[res]
	if ok {
		m[res] = val + 1
	} else {
		m[res] = 1
	}
}
