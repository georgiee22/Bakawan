package controllers

import (
	"Template/pkg/models"
	"Template/pkg/models/response"
	"Template/pkg/utils/go-utils/database"
	goUtils "Template/pkg/utils/go-utils/fiber"
	"Template/pkg/utils/go-utils/passwordHashing"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func HelloWorld(c *fiber.Ctx) error {
	c.Set("Authorization", "hehe")
	return c.SendString("HELLO MA'AM MAIC")
}

// get view login
func ReportsLogin(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}

// add functionality, if first login ask to change password else continue to home
func ReportsLoginAuth(c *fiber.Ctx) error {
	user_accounts := &models.User_Accounts{}

	if parsErr := c.BodyParser(user_accounts); parsErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "201",
			Message: "fail",
			Data:    parsErr.Error(),
		})
	}

	var dbpass string
	var userid string
	var ispasschange string

	err := database.DBConn.Debug().Raw("SELECT password, user_id, is_pass_change FROM user_accounts WHERE username=$1", user_accounts.Username).Row().Scan(&dbpass, &userid, &ispasschange)
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "invalid username or password",
			Data:    err.Error(),
		})
	}

	if !passwordHashing.CheckPasswordHash(user_accounts.Password, dbpass) {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "fail",
			Data:    "invalid password",
		})
	}

	claims := fiber.Map{
		"username": user_accounts.Username,
		"userid":   userid,
	}

	token, err := goUtils.GenerateJWTSignedString(claims)
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "fail",
			Data:    err.Error(),
		})
	}

	//c.ClearCookie()

	//Set the JWT token in a cookie
	// cookie := new(fiber.Cookie)
	// cookie.Name = "token"
	// cookie.Value = token
	// cookie.Expires = time.Now().Add(time.Minute * 5) // Expires in 15 mins
	// cookie.HTTPOnly = true                           // Cannot be accessed by JavaScript
	// cookie.Secure = false                            // Only transmitted over HTTPS
	// c.Cookie(cookie)

	// cookie := new(fiber.Cookie)
	// cookie.Name = "userid"
	// cookie.Value = userid
	// cookie.Expires = time.Now().Add(time.Minute * 5) // Expires in 15 mins
	// cookie.HTTPOnly = true                           // Cannot be accessed by JavaScript
	// cookie.Secure = false                            // Only transmitted over HTTPS
	// c.Cookie(cookie)
	// c.Locals("userid", userid)

	c.Append("Authorization", "Bearer "+token)
	// c.Append("Access-Control-Allow-Credentials", "true")
	// c.Append("Access-Control-Allow-Origin", "http://localhost:49909")

	// check if user wants to change password
	if ispasschange != "1" {
		return c.JSON(response.ResponseModel{
			RetCode: "100",
			Message: "continue",
			Data:    "change password",
		})
	}

	return c.JSON(response.ResponseModel{
		RetCode: "200",
		Message: "success",
		Data:    user_accounts.Username,
	})
}

func testHeader(w http.ResponseWriter, r *http.Request) error {
	user_accounts := &models.User_Accounts{}
	c := *&fiber.Ctx{}

	if parsErr := c.BodyParser(user_accounts); parsErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "201",
			Message: "fail",
			Data:    parsErr.Error(),
		})
	}

	var dbpass string
	var userid string
	var ispasschange string

	err := database.DBConn.Debug().Raw("SELECT password, user_id, is_pass_change FROM user_accounts WHERE username=$1", user_accounts.Username).Row().Scan(&dbpass, &userid, &ispasschange)
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "invalid username or password",
			Data:    err.Error(),
		})
	}

	if !passwordHashing.CheckPasswordHash(user_accounts.Password, dbpass) {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "fail",
			Data:    "invalid password",
		})
	}

	claims := fiber.Map{
		"username": user_accounts.Username,
		"userid":   userid,
	}

	token, err := goUtils.GenerateJWTSignedString(claims)
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "fail",
			Data:    err.Error(),
		})
	}

	//c.ClearCookie()

	//Set the JWT token in a cookie
	// cookie := new(fiber.Cookie)
	// cookie.Name = "token"
	// cookie.Value = token
	// cookie.Expires = time.Now().Add(time.Minute * 5) // Expires in 15 mins
	// cookie.HTTPOnly = true                           // Cannot be accessed by JavaScript
	// cookie.Secure = false                            // Only transmitted over HTTPS
	// c.Cookie(cookie)

	// cookie := new(fiber.Cookie)
	// cookie.Name = "userid"
	// cookie.Value = userid
	// cookie.Expires = time.Now().Add(time.Minute * 5) // Expires in 15 mins
	// cookie.HTTPOnly = true                           // Cannot be accessed by JavaScript
	// cookie.Secure = false                            // Only transmitted over HTTPS
	// c.Cookie(cookie)
	// c.Locals("userid", userid)

	c.Set("Authorization", "Bearer "+token)
	w.Header().Set("Authorization", "Bearer ")

	// check if user wants to change password
	if ispasschange != "1" {
		return c.JSON(response.ResponseModel{
			RetCode: "100",
			Message: "continue",
			Data:    "change password",
		})
	}

	return c.JSON(response.ResponseModel{
		RetCode: "200",
		Message: "success",
		Data:    user_accounts.Username,
	})
}

// GateToken := fmt.Sprintf("bearer %s", decryptedGateToken)
// 		// fmt.Println("PATH:", path)
// 		req, reqErr := http.NewRequest(routeResponse.Method, path, bytes.NewBuffer(jsonReq))
// 		req.Header.Add("Authorization", GateToken)
// 		req.Header.Add("Content-Type", "application/json")
// 		req.Header.Add("charset", "utf-8")

// 		// fmt.Println("HTTP REQUEST TLS STATUS:", req)
// 		if reqErr != nil {
// 			middleware.SystemLoggerError(c.Path(), "http request error", deviceId, reqErr)
// 			return c.JSON(response.ResponseModel{
// 				RetCode: "400",
// 				Message: "Error",
// 				Result: errors.ErrorModel{
// 					Message:   "Error in http request",
// 					IsSuccess: false,
// 					Error:     reqErr,
// 					Data:      reqErr.Error(),
// 				},
// 			})
// 		}

// 		client := &http.Client{}
// 		resp, clientErr := client.Do(req)
// 		// fmt.Println("HTTP RESPONSE STATUS:", resp.Status)
// 		if clientErr != nil {
// 			middleware.SystemLoggerError(c.Path(), "client response error", deviceId, clientErr)
// 			return c.JSON(response.ResponseModel{
// 				RetCode: "400",
// 				Message: "Error",
// 				Result: errors.ErrorModel{
// 					Message:   "Error in client response",
// 					IsSuccess: false,
// 					Error:     clientErr,
// 					Data:      clientErr.Error(),
// 				},
// 			})
// 		}

// 		defer resp.Body.Close()
// 		body, readErr := ioutil.ReadAll(resp.Body)
// 		if readErr != nil {
// 			middleware.SystemLoggerError(c.Path(), "read error", deviceId, readErr)
// 			return c.JSON(response.ResponseModel{
// 				RetCode: "400",
// 				Message: "Error",
// 				Result: errors.ErrorModel{
// 					Message:   "Error in reading http request",
// 					IsSuccess: false,
// 					Error:     readErr,
// 					Data:      readErr.Error(),
// 				},
// 			})
// 		}

// 		result := json.RawMessage(body)
// 		mapResult := make(map[string]interface{})
// 		if unmarErr := json.Unmarshal(result, &mapResult); unmarErr != nil {
// 			middleware.SystemLoggerError(c.Path(), "unmarshal error", deviceId, unmarErr)
// 			return c.JSON(response.ResponseModel{
// 				RetCode: "400",
// 				Message: "Error",
// 				Result: errors.ErrorModel{
// 					Message:   "unmarshal error",
// 					IsSuccess: false,
// 					Error:     unmarErr,
// 					Data:      unmarErr.Error(),
// 				},
// 			})
// 		}

// 		middleware.SystemLoggerAPI(path, req, c.Path(), resp, mapResult, c.IP()+":"+c.Port(), deviceId)
// 		fmt.Println("GATE ROUTE DONE")
// 		return c.JSON(response.ResponseModel{
// 			RetCode: "200",
// 			Message: "success",
// 			Result:  mapResult,
// 		})
// 	}

// get view change password, send user name
func ChangePasswordView(c *fiber.Ctx) error {
	// send user data to render
	// change into
	return c.Render("changepass", fiber.Map{
		"Title": "Login",
	})
}

// post change password
func ChangePassword(c *fiber.Ctx) error {
	change_pass := &models.Change_Pass{}
	// userid := c.Cookies("userid")

	// body parser, parses data submitted
	if parsErr := c.BodyParser(change_pass); parsErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "201",
			Message: "fail",
			Data:    parsErr.Error(),
		})
	}

	if change_pass.Password != change_pass.Confirm_password {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "fail",
			Data:    "password is not the same",
		})
	}

	// hash random password
	hashedPassword, err := passwordHashing.HashPassword(change_pass.Confirm_password)
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "password hashing error",
			Data:    err.Error(),
		})
	}

	err = database.DBConn.Exec("UPDATE user_accounts SET password = ?, is_pass_change = ? WHERE user_id = ?", hashedPassword, 1, change_pass.User_id).Error
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "203",
			Message: "query error",
			Data:    err.Error(),
		})
	}

	return c.JSON(response.ResponseModel{
		RetCode: "200",
		Message: "success",
		Data:    "password successfuly modified",
	})
}

// get list dashboards
func ListDashboards(c *fiber.Ctx) error {
	dashboard_list := []models.Dashboard_Apps_Tags_View{}
	id := c.Params("id")
	result := []models.Dashboards{}

	err := database.DBConn.Raw("SELECT DISTINCT dashboard_id, dashboard_title FROM dashboard_apps_tags WHERE app_id = $1", id).Find(&result, &dashboard_list).Error
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "203",
			Message: "query error",
			Data:    err.Error(),
		})
	}

	return c.JSON(response.ResponseModel{
		RetCode: "200",
		Message: "success",
		Data:    result,
	})
}

// get specific dashboard
func ViewDashboard(c *fiber.Ctx) error {
	dashboard := models.Dashboards{}
	id := c.Params("id")

	err := database.DBConn.Raw("SELECT * FROM dashboards WHERE dashboard_id = $1", id).Find(&dashboard).Error
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "203",
			Message: "query error",
			Data:    err.Error(),
		})
	}

	return c.JSON(response.ResponseModel{
		RetCode: "200",
		Message: "success",
		Data:    dashboard,
	})
}

// add functionality to create account with specified category
// post account creation with random password
func CreateReportsAccount(c *fiber.Ctx) error {
	// point to models
	user_accounts := &models.User_Accounts{}

	// body parser, parses data submitted
	if parsErr := c.BodyParser(user_accounts); parsErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "201",
			Message: "fail",
			Data:    parsErr.Error(),
		})
	}

	// query to check if username already exist
	var checker bool
	err := database.DBConn.Raw("SELECT EXISTS(SELECT 1 FROM user_accounts WHERE username = $1)", user_accounts.Username).Row().Scan(&checker)
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "203",
			Message: "query error",
			Data:    err.Error(),
		})
	}

	// check if username exist
	if checker {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "fail",
			Data:    "username already exist",
		})
	}

	// generate password
	random_password := goUtils.GenerateRandomPassword()

	// hash random password
	hashedPassword, err := passwordHashing.HashPassword(random_password)
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "password hashing error",
			Data:    err.Error(),
		})
	}

	// Insert new user with hashed password
	var lastId int
	err = database.DBConn.Raw("INSERT INTO user_accounts (username, password) VALUES (?, ?) RETURNING user_id",
		user_accounts.Username, hashedPassword).Scan(&lastId).Error
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "203",
			Message: "query error",
			Data:    err.Error(),
		})
	}

	// Insert User Details

	// return message
	return c.JSON(response.ResponseModel{
		RetCode: "200",
		Message: "success",
		Data:    random_password,
	})
}

// include every dashboard to automatically have admin tag....
// post create dashboard link
func CreateDashbord(c *fiber.Ctx) error {
	report_looker := &models.Report_Looker{}

	if parsErr := c.BodyParser(report_looker); parsErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "201",
			Message: "fail",
			Data:    parsErr.Error(),
		})
	}

	var checker bool
	err := database.DBConn.Raw("SELECT EXISTS(SELECT 1 FROM report_looker WHERE report_link = $1)", report_looker.Report_link).Row().Scan(&checker)
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "203",
			Message: "query error",
			Data:    err.Error(),
		})
	}

	// check if report exist
	if checker {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "fail",
			Data:    "report already exist",
		})
	}

	// Insert new report
	var lastId int
	err = database.DBConn.Raw("INSERT INTO report_looker (report_link) VALUES (?) RETURNING report_id",
		report_looker.Report_link).Scan(&lastId).Error
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "203",
			Message: "query error",
			Data:    err.Error(),
		})
	}

	// return message
	return c.JSON(response.ResponseModel{
		RetCode: "200",
		Message: "success",
		Data:    "report successfully added",
	})
}

// get send list of reports
func ListReports(c *fiber.Ctx) error {
	// userid := c.Cookies("userid")
	// report_accounts := &models.Report_Accounts{}
	report_link := []models.Report_Looker{}

	// reports contain categories and locations and etc.
	// table for user authorization
	// list categories or location for access
	// use checkmarks(true or false)

	// api

	err := database.DBConn.Raw("SELECT report_id, report_title FROM report_looker").Find(&report_link).Error
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "203",
			Message: "query error",
			Data:    err.Error(),
		})
	}

	return c.Render("home", fiber.Map{
		"Report_Looker": report_link,
	})
}

// get send specific report
func ViewReport(c *fiber.Ctx) error {
	report_id := models.Report_Looker{}
	id := c.Params("id")

	err := database.DBConn.Raw("SELECT * FROM report_looker WHERE report_id = $1", id).Find(&report_id).Error
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "203",
			Message: "query error",
			Data:    err.Error(),
		})
	}

	return c.Render("selectreport", fiber.Map{
		"Report_Looker": report_id,
	})
}

// get list of categories
