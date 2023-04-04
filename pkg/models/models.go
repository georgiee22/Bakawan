package models

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

type Login_Data struct {
	Username string `json:"username"`
	Token    string `json:"token"`
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
