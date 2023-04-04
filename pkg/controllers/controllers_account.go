package controllers

import (
	"Template/pkg/models"
	"Template/pkg/models/errors"
	"Template/pkg/models/response"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func LoginAPI(c *fiber.Ctx) error {
	user_accounts := &models.User_Accounts{}

	if parsErr := c.BodyParser(user_accounts); parsErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "201",
			Message: "fail",
			Data:    parsErr.Error(),
		})
	}

	jsonReq, err := json.Marshal(user_accounts)
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "fail",
			Data:    err.Error(),
		})
	}

	req, err := http.NewRequest("POST", "http://127.0.0.1:8000/api/public/v1/dashboards/login-authentication", bytes.NewBuffer(jsonReq))
	// req.Header.Add("Authorization", "Bearer eme")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("charset", "utf-8")
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "fail",
			Data:    err.Error(),
		})
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "fail",
			Data:    err.Error(),
		})
	}

	defer resp.Body.Close()

	// read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "fail",
			Data:    err.Error(),
		})
	}

	result := json.RawMessage(body)
	mapResult := make(map[string]interface{})
	if unmarErr := json.Unmarshal(result, &mapResult); unmarErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data: errors.ErrorModel{
				Message:   "unmarshal error",
				IsSuccess: false,
				Error:     unmarErr,
			},
		})
	}

	var data models.Login_Data
	data.Username = user_accounts.Username
	data.Token = mapResult["data"].(string)

	// send the response back to the client
	return c.JSON(response.ResponseModel{
		RetCode: "200",
		Message: "success",
		Data:    data,
	})
}
