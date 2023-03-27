package controllers

import (
	"Template/pkg/models"
	"Template/pkg/models/response"
	"Template/pkg/utils/go-utils/database"
	goUtils "Template/pkg/utils/go-utils/fiber"
	"Template/pkg/utils/go-utils/passwordHashing"
	"time"

	"github.com/gofiber/fiber/v2"
)

// get view login
func ReportsLogin(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}

// add functionality, if first login ask to change password else continue to home
func ReportsLoginAuth(c *fiber.Ctx) error {
	report_accounts := &models.Report_Accounts{}

	if parsErr := c.BodyParser(report_accounts); parsErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "201",
			Message: "fail",
			Data:    parsErr.Error(),
		})
	}

	var dbpass string
	var userid string
	var ispasschange string

	err := database.DBConn.Debug().Raw("SELECT password, user_id, is_pass_change FROM user_accounts WHERE username=$1", report_accounts.Username).Row().Scan(&dbpass, &userid, &ispasschange)
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "invalid username or password",
			Data:    err.Error(),
		})
	}

	if !passwordHashing.CheckPasswordHash(report_accounts.Password, dbpass) {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "fail",
			Data:    "invalid password",
		})
	}

	claims := fiber.Map{
		"username": report_accounts.Username,
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
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Minute * 5) // Expires in 15 mins
	cookie.HTTPOnly = true                           // Cannot be accessed by JavaScript
	cookie.Secure = false                            // Only transmitted over HTTPS
	c.Cookie(cookie)

	cookie = new(fiber.Cookie)
	cookie.Name = "userid"
	cookie.Value = userid
	cookie.Expires = time.Now().Add(time.Minute * 5) // Expires in 15 mins
	cookie.HTTPOnly = false                          // Cannot be accessed by JavaScript
	cookie.Secure = false                            // Only transmitted over HTTPS
	c.Cookie(cookie)

	// check if user wants to change password
	// if ispasschange != "1" {
	// 	return c.Redirect("/report/changepass")
	// }

	// return c.Redirect("./protected/listreports")

	return c.JSON(response.ResponseModel{
		RetCode: "200",
		Message: "success",
		Data:    "Login Successfully",
	})
}

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

	err = database.DBConn.Exec("UPDATE user_accounts SET password = ? WHERE user_id = ?", hashedPassword, 1).Error
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

// inserting report link
func CreateReportsLink(c *fiber.Ctx) error {
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

// verify if user is authorized to acces the report
func VerifyAuth2ndLayer(c *fiber.Ctx) error {
	report_id := c.Params("id")
	user_id := c.Cookies("userid")

	var looker_tags []string
	var user_tags []string

	// get report tags
	err := database.DBConn.Raw("SELECT tag_id FROM report_looker_tags WHERE report_id = $1", report_id).Find(&looker_tags).Error
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "203",
			Message: "query error",
			Data:    err.Error(),
		})
	}

	// get user tags
	err = database.DBConn.Raw("SELECT tag_id FROM report_accounts_tags WHERE account_id = $1", user_id).Find(&user_tags).Error
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "203",
			Message: "query error",
			Data:    err.Error(),
		})
	}

	// check if same tags exist between user tags and report tags
	strMap := make(map[string]bool)

	for _, str := range looker_tags {
		strMap[str] = true
	}

	for _, str := range user_tags {
		if strMap[str] {
			return c.Redirect("./viewreport/" + report_id)
		}
	}

	return c.JSON(response.ResponseModel{
		RetCode: "400",
		Message: "fail",
		Data:    "no access",
	})
}

// verify if user is authorized to acces the app specific reports
func VerifyAuth1stLayer(c *fiber.Ctx) error {
	report_id := c.Params("id")
	user_id := c.Cookies("userid")

	var looker_tags []string
	var user_tags []string

	// get specific category
	err := database.DBConn.Raw("SELECT category_id FROM report_looker_tags WHERE report_id = $1", report_id).Find(&looker_tags).Error
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "203",
			Message: "query error",
			Data:    err.Error(),
		})
	}

	// get user tags
	err = database.DBConn.Raw("SELECT tag_id FROM report_accounts_tags WHERE account_id = $1", user_id).Find(&user_tags).Error
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "203",
			Message: "query error",
			Data:    err.Error(),
		})
	}

	// check if same tags exist between user tags and report tags
	strMap := make(map[string]bool)

	for _, str := range looker_tags {
		strMap[str] = true
	}

	for _, str := range user_tags {
		if strMap[str] {
			return c.Redirect("./viewreport/" + report_id)
		}
	}

	return c.JSON(response.ResponseModel{
		RetCode: "400",
		Message: "fail",
		Data:    "no access",
	})
}

// get list of categories
