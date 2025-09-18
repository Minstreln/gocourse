package intermediate

import (
	"log"
	"os"
)

func main() {

	// log.Println("This is a log message.")
	// log.SetPrefix("INFO: ")
	// log.Println("This is another log message with a prefix.")

	// adding log flags to log messages
	// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// log.Println("This log message has a date.")

	infoLogger.Println("This is an info message.")
	warnLogger.Println("This is a warning message.")
	errorLogger.Println("This is an error message.")

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	defer file.Close()

	debugLogger := log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger1 := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger1 := log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger1 := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	debugLogger.Println("This is a debug message written to the file.")
	infoLogger1.Println("This is an info message written to the file.")
	warnLogger1.Println("This is a warning message written to the file.")
	errorLogger1.Println("This is an error message written to the file.")

}

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)
