package config

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/Terry-Mao/goconf"
	"golang.org/x/crypto/bcrypt"
	"hcc/clarinet/lib/passwordUtil"
	"log"
	"os"
	goUser "os/user"
	"path/filepath"
)

type ClarinetConfig struct {
	PiccoloConfig *goconf.Section
	UserConfig    *goconf.Section
}

var configLocation = "/etc/hcc/clarinet/clarinet.conf"
var userConfLocation string

func setUserConfFilePath() {
	curUser, _ := goUser.Current()
	userConfLocation = curUser.HomeDir + "/.hcc/clarinet/user.conf"
}

func createConfFile() error {
	if err := os.MkdirAll(filepath.Dir(userConfLocation), 0770); err != nil {
		return err
	}
	if _, err = os.Create(userConfLocation); err != nil {
		return err
	}
	conf := usrConf.Add("user")
	conf.Add("token", "")
	if err = usrConf.Save(userConfLocation); err != nil {
		return err
	}

	return nil
}

func GetUserInfo(userID *string, userPassword *string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Set user information.")

	fmt.Print("User ID : ")
	scanner.Scan()
	*userID = scanner.Text()
	/*TODO: Check alphanumeric user id*/

	fmt.Print("User PW : ")
	// User.UserPasswd = getPassword()
	md := sha256.Sum256([]byte(passwordUtil.GetPassword()))
	mdStr := hex.EncodeToString(md[:])
	hashPW, err := bcrypt.GenerateFromPassword([]byte(mdStr), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Password hash error: %v", err)
	}
	*userPassword = string(hashPW)
}

func SaveTokenString(tokenString string) {
	User.Token = tokenString
	conf := usrConf.Get("user")
	conf.Add("token", tokenString)
	if err := usrConf.Save(userConfLocation); err != nil {
		fmt.Println("Token save failed")
	}
}

func RemoveTokenString() {

	if err := os.Remove(userConfLocation); err != nil {
		fmt.Println("Failed : Can not find user location")
	} else {
		fmt.Println("Succeed : User logged out")
	}
}
