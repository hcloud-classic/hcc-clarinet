package config

import (
	"bufio"
	"fmt"
	"github.com/Terry-Mao/goconf"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"os/signal"
	goUser "os/user"
	"path/filepath"
	"syscall"
)

var configLocation = "/etc/hcc/clarinet/clarinet.conf"
var userConfLocation string

type ClarinetConfig struct {
	FluteConfig  *goconf.Section
	HarpConfig   *goconf.Section
	ViolinConfig *goconf.Section
	UserConfig   *goconf.Section
}

func setUserConfFilePath() {
	curUser, _ := goUser.Current()
	userConfLocation = curUser.HomeDir + "/.hcc/clarinet/user.conf"

}

/*TODO: Check alphanumeric user id*/

func GetUserInfo() error {

	if err := os.MkdirAll(filepath.Dir(userConfLocation), 0770); err != nil {
		return err
	}
	if _, err = os.Create(userConfLocation); err != nil {
		return err
	}

	c := usrConf.Add("user")
	c.Remove("user_id")
	c.Remove("user_passwd")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Set user information.")

	fmt.Print("User ID : ")
	userID, _ := reader.ReadString('\n')

	fmt.Print("User PW : ")
	userPasswd := getPassword()

	fmt.Println(userPasswd)

	c.Add("user_id", userID)
	c.Add("user_passwd", userPasswd)

	if err = usrConf.Save(userConfLocation); err != nil {
		return err
	}

	usrConf.Parse(userConfLocation)

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
