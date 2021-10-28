package main

import (
    "golang.org/x/crypto/ssh"
    "log"
    "os"
	"fmt"
	"github.com/fatih/color"
)

var (
	ipaddr string
	user string
	passwd string
	port string
)


func main() {
 	//Get Settings from User
	getSettings()
	color.Red("Dialing....")

	// Establish an SSH client connection
    client, err := ssh.Dial("tcp", ipaddr + ":" + port, &ssh.ClientConfig{
        User:            user,
        Auth:            []ssh.AuthMethod{ssh.Password(passwd)},
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    })
    if err != nil {
        log.Fatalf("SSH dial error: %s", err.Error())
    }

         // Establish a new session
    session, err := client.NewSession()
	if err != nil {
        log.Fatalf("new session error: %s", err.Error())
    }
    defer session.Close()

         session.Stdout = os.Stdout // Session output association to system standard output device
         session.Stderr = os.Stderr // Session error output association to system standard error output device
         session.Stdin = os.Stdin // Session input association to system standard input device
    modes := ssh.TerminalModes{
                 ssh.ECHO: 0, // Disabled Election (0 Disabled, 1 Start)
        ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
        ssh.TTY_OP_OSPEED: 14400, //output speed = 14.4kbaud
    }
    if err = session.RequestPty("xterm", 32, 160, modes); err != nil {
        log.Fatalf("request pty error: %s", err.Error())
    }
    if err = session.Shell(); err != nil {
        log.Fatalf("start shell error: %s", err.Error())
    }
    if err = session.Wait(); err != nil {
        log.Fatalf("return error: %s", err.Error())
    }
}

func getSettings() {
	fmt.Println("Host? (default localhost)")
	fmt.Scanln(&ipaddr)
	fmt.Println("Port? (default 22)")
	fmt.Scanln(&port)
	fmt.Println("Username? (default root)")
	fmt.Scanln(&user)
	fmt.Println("Password? (default toor)")
	fmt.Scanln(&passwd)

	if ipaddr == "" {
		ipaddr = "localhost"
	}
	if port == "" {
		port = "22"
	}
	if user == "" {
		user = "root"
	}
	if passwd == "" {
		passwd = "toor"
	}
}