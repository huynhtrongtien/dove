package http_parser

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/beego/beego/validation"
	"github.com/gin-gonic/gin"
)

func BindAndValid(c *gin.Context, form interface{}, skipValids ...string) error {
	err := c.ShouldBind(form)
	if err != nil {
		return err
	}

	return Valid(c, form, skipValids...)
}

func BindJSONAndValid(c *gin.Context, form interface{}, skipValids ...string) error {
	err := c.ShouldBindJSON(form)
	if err != nil {
		return err
	}

	return Valid(c, form, skipValids...)
}

func BindAndValidUri(c *gin.Context, form interface{}, skipValids ...string) error {

	err := c.BindUri(form)

	if err != nil {
		return err
	}

	return Valid(c, form, skipValids...)
}

func GetJsonAndValid(c *gin.Context, form interface{}, skipValids ...string) (map[string]interface{}, error) {
	jsonData := map[string]interface{}{}
	data, _ := io.ReadAll(c.Request.Body)

	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, form); err != nil {
		return nil, err
	}

	valid := validation.Validation{RequiredFirst: true}
	for _, validName := range skipValids {
		valid.CanSkipAlso(validName)
	}
	check, err := valid.Valid(form)
	if err != nil {
		return nil, err
	}
	if !check {
		// MarkErrors(valid.Errors)
		return nil, fmt.Errorf("validate failed")
	}

	return jsonData, nil
}

func Valid(c *gin.Context, form interface{}, skipValids ...string) error {
	valid := validation.Validation{RequiredFirst: true}
	for _, validName := range skipValids {
		valid.CanSkipAlso(validName)
	}

	check, err := valid.RecursiveValid(form)
	if err != nil {
		return err
	}
	if !check {
		return fmt.Errorf("validate failed")
	}

	return nil
}

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		fmt.Printf("Error %s - %s \n", err.Key, err.Message)
	}
}
