package cmd

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hcc/clarinet/lib/mysql"
	"hcc/clarinet/lib/passwordUtil"
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"

	"hcc/clarinet/action/graphql/mutationParser"
	"hcc/clarinet/action/graphql/queryParser"
	"hcc/clarinet/lib/config"
	"hcc/clarinet/model"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Commands for user",
	Long:  `Commands for create and update, delete users`,
	Args:  cobra.MinimumNArgs(1),
}

var mysqlAddress, mysqlUser string
var mysqlPort int64
var id, authentication, name, email string
var groupID int64
var changePassword bool

var addMaster = &cobra.Command{
	Use:   "add_master",
	Short: "Add a master account",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Enter MySQL Password : ")
		mysqlPassword := passwordUtil.GetPassword()
		fmt.Println("Connecting to MySQL...")
		err := mysql.Init(mysqlUser, mysqlPassword, mysqlAddress, mysqlPort)
		if err != nil {
			fmt.Println("[FAIL] :" + err.Error())
			return
		}

		fmt.Println("Setting password of master account...")
		fmt.Print("Enter Password : ")
		password1 := passwordUtil.GetPassword()
		fmt.Print("Confirm Password : ")
		password2 := passwordUtil.GetPassword()

		if password1 != password2 {
			fmt.Println("Passwords are mismatch.")
			return
		}

		sha256pass := sha256.Sum256([]byte(password1))
		password := hex.EncodeToString(sha256pass[:])

		fmt.Println("Creating master account .... ")

		user := model.User{
			ID:             "master",
			GroupID:        1, // master uses Group ID 1
			Authentication: "master",
			Name:           name,
			Email:          email,
		}

		sql := "insert into user(id, group_id, authentication, password, name, email, login_at, created_at) values (?, ?, ?, ?, ?, ?, now(), now())"
		stmt, err := mysql.Prepare(sql)
		if err != nil {
			panic(err)
		}
		defer func() {
			_ = stmt.Close()
		}()
		_, err = stmt.Exec(user.ID, user.GroupID, user.Authentication, password, user.Name, user.Email)
		if err != nil {
			fmt.Println("[ERROR]: " + err.Error())
			return
		}

		fmt.Println("[SUCCESS]")

		t := table.NewWriter()
		t.SetStyle(table.Style{
			Name: "clarinetTableStyle",
			Box:  table.StyleBoxLight,
			Format: table.FormatOptions{
				Header: text.FormatUpper,
			},
			Options: table.Options{
				DrawBorder:      true,
				SeparateColumns: true,
				SeparateFooter:  true,
				SeparateHeader:  true,
				SeparateRows:    false,
			},
		})
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"ID", user.ID})
		t.AppendRow([]interface{}{"Group ID", user.GroupID})
		t.AppendRow([]interface{}{"Name", user.Name})
		t.AppendRow([]interface{}{"E-Mail", user.Email})
		t.Render()
	},
}

var userSignUp = &cobra.Command{
	Use:     "signup",
	Short:   "User SignUp",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["id"] = id
		queryArgs["group_id"] = strconv.Itoa(int(groupID))
		queryArgs["authentication"] = authentication
		queryArgs["name"] = name
		queryArgs["email"] = email
		queryArgs["token"] = config.User.Token

		fmt.Print("Enter Password : ")
		password1 := passwordUtil.GetPassword()
		fmt.Print("Confirm Password : ")
		password2 := passwordUtil.GetPassword()

		if password1 != password2 {
			fmt.Println("Passwords are mismatch.")
			return
		}

		sha256pass := sha256.Sum256([]byte(password1))
		queryArgs["password"] = hex.EncodeToString(sha256pass[:])

		fmt.Println("Creating User .... ")

		data, err := mutationParser.SignUp(queryArgs)
		if err != nil {
			fmt.Println("[FAIL]")
			err.Println()
			return
		}

		userData := data.(model.User)
		if len(userData.Errors) > 0 {
			for _, hrr := range userData.Errors {
				hrr.Println()
			}
			return
		}

		fmt.Println("[SUCCESS]")

		t := table.NewWriter()
		t.SetStyle(table.Style{
			Name: "clarinetTableStyle",
			Box:  table.StyleBoxLight,
			Format: table.FormatOptions{
				Header: text.FormatUpper,
			},
			Options: table.Options{
				DrawBorder:      true,
				SeparateColumns: true,
				SeparateFooter:  true,
				SeparateHeader:  true,
				SeparateRows:    false,
			},
		})
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"ID", userData.ID})
		t.AppendRow([]interface{}{"Group ID", userData.GroupID})
		t.AppendRow([]interface{}{"Name", userData.Name})
		t.AppendRow([]interface{}{"E-Mail", userData.Email})
		t.Render()
	},
}

var userUpdate = &cobra.Command{
	Use:     "update",
	Short:   "Update User",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["id"] = id
		queryArgs["authentication"] = authentication
		queryArgs["name"] = name
		queryArgs["email"] = email
		queryArgs["token"] = config.User.Token

		if changePassword {
			fmt.Print("Enter Password : ")
			password1 := passwordUtil.GetPassword()
			fmt.Print("Confirm Password : ")
			password2 := passwordUtil.GetPassword()

			if password1 != password2 {
				fmt.Println("Passwords are mismatch.")
				return
			}

			sha256pass := sha256.Sum256([]byte(password1))
			queryArgs["password"] = hex.EncodeToString(sha256pass[:])
		}

		fmt.Println("Updating User .... ")

		data, err := mutationParser.UpdateUser(queryArgs)
		if err != nil {
			fmt.Println("[FAIL]")
			err.Println()
			return
		}

		userData := data.(model.User)
		if len(userData.Errors) > 0 {
			for _, hrr := range userData.Errors {
				hrr.Println()
			}
			return
		}

		fmt.Println("[SUCCESS]")

		t := table.NewWriter()
		t.SetStyle(table.Style{
			Name: "clarinetTableStyle",
			Box:  table.StyleBoxLight,
			Format: table.FormatOptions{
				Header: text.FormatUpper,
			},
			Options: table.Options{
				DrawBorder:      true,
				SeparateColumns: true,
				SeparateFooter:  true,
				SeparateHeader:  true,
				SeparateRows:    false,
			},
		})
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"ID", userData.ID})
		t.AppendRow([]interface{}{"Group ID", userData.GroupID})
		t.AppendRow([]interface{}{"Group Name", userData.GroupName})
		t.AppendRow([]interface{}{"Name", userData.Name})
		t.AppendRow([]interface{}{"E-Mail", userData.Email})
		t.AppendRow([]interface{}{"Login At", userData.LoginAt})
		t.AppendRow([]interface{}{"Created At", userData.CreatedAt})
		t.Render()
	},
}

var userUnregister = &cobra.Command{
	Use:     "unregister",
	Short:   "Unregister User",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["id"] = id
		queryArgs["token"] = config.User.Token

		fmt.Println("Deleting User .... ")

		data, err := mutationParser.Unregister(queryArgs)
		if err != nil {
			fmt.Println("[FAIL]")
			err.Println()
			return
		}

		userData := data.(model.User)
		if len(userData.Errors) > 0 {
			for _, hrr := range userData.Errors {
				hrr.Println()
			}
			return
		}

		fmt.Println("[SUCCESS]")

		t := table.NewWriter()
		t.SetStyle(table.Style{
			Name: "clarinetTableStyle",
			Box:  table.StyleBoxLight,
			Format: table.FormatOptions{
				Header: text.FormatUpper,
			},
			Options: table.Options{
				DrawBorder:      true,
				SeparateColumns: true,
				SeparateFooter:  true,
				SeparateHeader:  true,
				SeparateRows:    false,
			},
		})
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"ID", id})
		t.Render()
	},
}

var userList = &cobra.Command{
	Use:     "list",
	Short:   "Show user list",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["row"] = strconv.Itoa(row)
		queryArgs["page"] = strconv.Itoa(page)
		queryArgs["id"] = id
		queryArgs["group_id"] = strconv.Itoa(int(groupID))
		queryArgs["authentication"] = authentication
		queryArgs["name"] = name
		queryArgs["email"] = email
		queryArgs["token"] = config.User.Token

		data, err := queryParser.ListUser(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		userList := data.(model.UserList)
		if len(userList.Errors) > 0 {
			for _, hrr := range userList.Errors {
				hrr.Println()
			}
			return
		}

		t := table.NewWriter()
		t.SetStyle(table.Style{
			Name: "clarinetTableStyle",
			Box:  table.StyleBoxLight,
			Format: table.FormatOptions{
				Header: text.FormatUpper,
			},
			Options: table.Options{
				DrawBorder:      true,
				SeparateColumns: true,
				SeparateFooter:  true,
				SeparateHeader:  true,
				SeparateRows:    false,
			},
		})
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"No", "ID", "Group ID", "Group Name", "Authentication", "Name", "E-Mail",
			"Login At", "Created At"})

		for index, user := range userList.Users {
			t.AppendRow([]interface{}{index + 1, user.ID, user.GroupID, user.GroupName, user.Authentication,
				user.Name, user.Email, user.LoginAt, user.CreatedAt})
		}

		t.AppendFooter(table.Row{"Total", len(userList.Users)})
		t.Render()
	},
}

func ReadyUserCmd() {
	addMaster.Flags().StringVar(&mysqlAddress, "mysql_address", "", "MySQL Address")
	addMaster.Flags().Int64Var(&mysqlPort, "mysql_port", 3306, "MySQL Port Number")
	addMaster.Flags().StringVar(&mysqlUser, "mysql_user", "", "MySQL Username")
	addMaster.Flags().StringVar(&name, "master_name", "", "Name")
	addMaster.Flags().StringVar(&email, "master_email", "", "E-Mail")
	addMaster.MarkFlagRequired("mysql_address")
	addMaster.MarkFlagRequired("mysql_user")

	userSignUp.Flags().StringVar(&id, "id", "", "ID")
	userSignUp.Flags().Int64Var(&groupID, "group_id", 0, "Group ID")
	userSignUp.Flags().StringVar(&authentication, "authentication", "", "Authentication (admin/user)")
	userSignUp.Flags().StringVar(&name, "name", "", "Name")
	userSignUp.Flags().StringVar(&email, "email", "", "E-Mail")
	userSignUp.MarkFlagRequired("id")
	userSignUp.MarkFlagRequired("authentication")
	userSignUp.MarkFlagRequired("name")
	userSignUp.MarkFlagRequired("email")

	userUpdate.Flags().StringVar(&id, "id", "", "ID")
	userUpdate.Flags().StringVar(&authentication, "authentication", "", "Authentication (admin/user)")
	userUpdate.Flags().StringVar(&name, "name", "", "Name")
	userUpdate.Flags().StringVar(&email, "email", "", "E-Mail")
	userUpdate.Flags().BoolVar(&changePassword, "change_password", false, "Change Password (true/false)")
	userUpdate.MarkFlagRequired("id")

	userUnregister.Flags().StringVar(&id, "id", "", "ID")
	userUnregister.MarkFlagRequired("id")

	userList.Flags().IntVar(&row, "row", 0, "")
	userList.Flags().IntVar(&page, "page", 0, "")
	userList.Flags().StringVar(&id, "id", "", "ID")
	userList.Flags().Int64Var(&groupID, "group_id", 0, "Group ID")
	userList.Flags().StringVar(&authentication, "authentication", "", "Authentication (admin/user)")
	userList.Flags().StringVar(&name, "name", "", "Name")
	userList.Flags().StringVar(&email, "email", "", "E-Mail")

	userCmd.AddCommand(addMaster, userSignUp, userUpdate, userUnregister, userList)
}
