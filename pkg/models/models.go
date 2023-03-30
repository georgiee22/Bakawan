package models

type Accounts struct {
	Account_id int    `json:"accountid" gorm:"primaryKey"`
	First_name string `json:"firstname" validate:"required"`
	Last_name  string `json:"lastname" validate:"required"`
	Username   string `json:"username" validate:"required"`
	Password   string `json:"password" validate:"required"`
}

type Contacts struct {
	Id         int    `json:"id"`
	Account_id int    `json:"accountid"`
	Email      string `json:"email"  validate:"omitempty,email"`
	Contact    string `json:"contact"  validate:"omitempty,number,max=11"`
}

type Address struct {
	Id          int    `json:"id"`
	Account_id  int    `json:"accountid"`
	House_no    string `json:"houseno"`
	Street      string `json:"street"`
	Subdivision string `json:"subdivision"`
	Barangay    string `json:"barangay"`
	City        string `json:"city"`
	Province    string `json:"province"`
	Country     string `json:"country"`
	Zip_code    string `json:"zipcode"`
}

// change to user_accounts (delete)
type Report_Accounts struct {
	User_id        int    `json:"userid"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Is_pass_change int    `json:"ispasschange"`
}

// new model
type User_Accounts struct {
	User_id        int    `json:"userid"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Is_pass_change int    `json:"ispasschange"`
}

type Change_Pass struct {
	User_id          int    `json:"userid"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Confirm_password string `json:"confirmpassword"`
}

// change to dashboards (delete)
type Report_Looker struct {
	Report_id    int    `json:"reportid"`
	Report_title string `json:"reporttitle"`
	Report_link  string `json:"reportlink"`
}

// new model
type Dashboards struct {
	Dashboard_id    int    `json:"dashboardid"`
	Dashboard_title string `json:"dashboardtitle"`
	Dashboard_link  string `json:"dashboardlink"`
}

// new model
type Apps struct {
	App_id    int    `json:"appid"`
	App_title string `json:"apptitle"`
}

type Access_Dashboard_Apps struct {
	Access_dashboard_apps_id int `json:"accessdashboardappsid"`
	Dashboard_id             int `json:"dashboardid"`
	App_id                   int `json:"appid"`
}

// dashboard apps and tags view
type Dashboard_Apps_Tags_View struct {
	Dashboard_id    int    `json:"dashboardid"`
	Dashboard_title string `json:"dashboardtitle"`
	Dashboard_link  string `json:"dashboardlink"`
	App_id          int    `json:"appid"`
	App_title       string `json:"apptitle"`
	Tag_id          int    `json:"tagid"`
	Tag_title       string `json:"tagtitle"`
}

type Account_Apps_Tags_View struct {
	User_id int `json:"userid"`
	App_id  int `json:"appid"`
	Tag_id  int `json:"tagid"`
}
