package config

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/signal"
	goUser "os/user"
	"path/filepath"
	"syscall"

	"github.com/Terry-Mao/goconf"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
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

func getPassword() string {
	// Get the initial state of the terminal.
	initialTermState, e1 := terminal.GetState(syscall.Stdin)
	if e1 != nil {
		log.Panic(e1)
	}

	// Restore it in the event of an interrupt.
	// CITATION: Konstantin Shaposhnikov - https://groups.google.com/forum/#!topic/golang-nuts/kTVAbtee9UA
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		<-c
		_ = terminal.Restore(syscall.Stdin, initialTermState)
		os.Exit(1)
	}()

	// Now get the password.
	p, err := terminal.ReadPassword(syscall.Stdin)
	fmt.Println("")
	if err != nil {
		log.Panic(err)
	}

	// Stop looking for ^C on the channel.
	signal.Stop(c)

	// Return the password as a string.
	return string(p)
}

func GetUserInfo() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Set user information.")

	fmt.Print("User ID : ")
	scanner.Scan()
	User.UserId = scanner.Text()
	/*TODO: Check alphanumeric user id*/

	fmt.Print("User PW : ")
	// User.UserPasswd = getPassword()
	md := sha256.Sum256([]byte(getPassword()))
	mdStr := hex.EncodeToString(md[:])
	hashPW, err := bcrypt.GenerateFromPassword([]byte(mdStr), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Password hash error: %v", err)
	}
	User.UserPasswd = string(hashPW)
}

func SaveTokenString(tokenString string) {
	User.Token = tokenString
	conf := usrConf.Get("user")
	conf.Add("token", tokenString)
	if err := usrConf.Save(userConfLocation); err != nil {
		fmt.Println("Token save failed")
	}
}
