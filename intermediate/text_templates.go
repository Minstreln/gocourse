package intermediate

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)

func main() {

	// text templates
	// temp, err := template.New("example").Parse("Welcome {{.name}}! How are you doing?\n")
	// if err != nil {
	// 	panic(err)
	// }

	// // define data for the welcome message template
	// data := map[string]interface{}{
	// 	"name": "John",
	// }

	// err = temp.Execute(os.Stdout, data)
	// if err != nil {
	// 	panic(err)
	// }

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter you name: ")

	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// define named templates for different types of messages
	templates := map[string]string{
		"welcome":      "Welcome {{.name}}! We're glad you joined.",
		"notification": "{{.name}}, you have a new notification: {{.notify}}",
		"error":        "Oops! Something went wrong: {{.errMessage}}",
	}

	// paarse and store templates
	parsedTemplates := make(map[string]*template.Template)
	for name, temp := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(temp))
	}

	for {
		// show menu
		fmt.Println("\nMenu")
		fmt.Println("1. Join")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit")
		fmt.Println("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]interface{}
		var temp1 *template.Template

		switch choice {
		case "1":
			temp1 = parsedTemplates["welcome"]
			data = map[string]interface{}{
				"name": name,
			}
		case "2":
			fmt.Println("Enter your notification message: ")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			temp1 = parsedTemplates["notification"]
			data = map[string]interface{}{
				"name":   name,
				"notify": notification,
			}
		case "3":
			fmt.Println("Enter your error message: ")
			errorMessage, _ := reader.ReadString('\n')
			errorMessage = strings.TrimSpace(errorMessage)
			temp1 = parsedTemplates["error"]
			data = map[string]interface{}{
				"name":       name,
				"errMessage": errorMessage,
			}
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
			continue
		}

		// render and print the template to the console
		if err := temp1.Execute(os.Stdout, data); err != nil {
			fmt.Println("Error executing template:", err)
		}

	}
}
