package ui

import (
	"bufio"
	"fmt"
	"github.com/agusnoceto/modak-challenge/internal/model"
	"net/mail"
	"os"
	"strings"
)

const (
	Banner = `██████╗  █████╗ ████████╗███████╗    ██╗     ██╗███╗   ███╗██╗████████╗███████╗██████╗     ███████╗███████╗██████╗ ██╗   ██╗██╗ ██████╗███████╗
██╔══██╗██╔══██╗╚══██╔══╝██╔════╝    ██║     ██║████╗ ████║██║╚══██╔══╝██╔════╝██╔══██╗    ██╔════╝██╔════╝██╔══██╗██║   ██║██║██╔════╝██╔════╝
██████╔╝███████║   ██║   █████╗      ██║     ██║██╔████╔██║██║   ██║   █████╗  ██║  ██║    ███████╗█████╗  ██████╔╝██║   ██║██║██║     █████╗  
██╔══██╗██╔══██║   ██║   ██╔══╝      ██║     ██║██║╚██╔╝██║██║   ██║   ██╔══╝  ██║  ██║    ╚════██║██╔══╝  ██╔══██╗╚██╗ ██╔╝██║██║     ██╔══╝  
██║  ██║██║  ██║   ██║   ███████╗    ███████╗██║██║ ╚═╝ ██║██║   ██║   ███████╗██████╔╝    ███████║███████╗██║  ██║ ╚████╔╝ ██║╚██████╗███████╗
╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ╚══════╝    ╚══════╝╚═╝╚═╝     ╚═╝╚═╝   ╚═╝   ╚══════╝╚═════╝     ╚══════╝╚══════╝╚═╝  ╚═╝  ╚═══╝  ╚═╝ ╚═════╝╚══════╝
                                                                                                                                               `

	WelcomeMessage = "Welcome to Modak's code challenge solution"

	Instructions = `This program allows you to send messages of different types to the desired user. Please follow the instruction listed. For more information please refer to the readme file`

	EnterUserMail    = "Please enter the user e-mail: "
	EnterMessageKey  = "Please enter the message type to send (%s, %s, %s): "
	EnterMessage     = "Please enter the message to send to the user: "
	Again            = "Do you want to send another message ? [y/n]: "
	GoodBye          = "Good bye!"
	MessageDelimiter = "------------------------------------------------------------------------------------------------"
)

func PrintWelcomeMessage() {
	fmt.Println()
	fmt.Println(Banner)
	fmt.Println()
	fmt.Println(WelcomeMessage)
	fmt.Println()
	fmt.Println(Instructions)
}

func PrintGoodBye() {
	fmt.Println(GoodBye)
}

func PrintDelimiter() {
	fmt.Println(MessageDelimiter)
}

func ReadValues() (model.MessageKey, string, string) {
	key := readMessageType()
	email := readEmail()
	msg := readMessage()
	fmt.Println()
	return key, email, msg
}

func readMessageType() model.MessageKey {
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Println()
		fmt.Printf(EnterMessageKey, model.MessageKeyStatus, model.MessageKeyNews, model.MessageKeyMarketing)

		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}

		key := scanner.Text()
		if !isValidKey(key) {
			fmt.Println("Error. Please enter a valid message type")
			continue
		}
		return model.MessageKey(key)
	}
	return ""
}

func isValidKey(key string) bool {
	switch model.MessageKey(key) {
	case model.MessageKeyStatus, model.MessageKeyNews, model.MessageKeyMarketing:
		return true
	default:
		return false
	}
}

func readEmail() string {
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Println()
		fmt.Print(EnterUserMail)

		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		email := scanner.Text()
		if !isValidEmail(email) {
			fmt.Println("Error. Please enter a valid user e-mail")
			continue
		}
		return email
	}
	return ""
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func readMessage() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println()
	fmt.Print(EnterMessage)

	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	msg := scanner.Text()
	return msg
}

// readYesOrNo returns true if the user enters 'y' or 'Y'.
// false if he/she enters 'n', 'N'. No others characters allowed.
func readYesOrNo(msg string) bool {
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Println()
		fmt.Print(msg)

		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}
		input := strings.ToLower(scanner.Text())
		if len(input) != 1 || input != "y" && input != "n" {
			fmt.Println("Error: Only ['y', 'n', 'Y, 'N'] are allowed.")
			continue
		}
		return input == "y"
	}
	return false
}

func SendAnotherMessage() bool {
	return readYesOrNo(Again)
}
