package basic

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}


func Hello()  {
	fmt.Println("Hello World!")
}


func HelloWithError(name string) (string, error) {
	// error 
	if name == "" {
		return "", errors.New("Empty Name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome to my channel", name)
	return message, nil
}


func HelloRandomMessage(name string) (string, error) {
	if name == "" {
        return name, errors.New("empty name")
	}
	
	formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
	}
	
    message := fmt.Sprintf(formats[rand.Intn(len(formats))], name)
    return message, nil
}


func HelloPeoples(names []string) (map[string]string, error) {
	messages := make(map[string]string)
    for _, name := range names {
        message, err := HelloRandomMessage(name)
        if err != nil {
            return nil, err
        }
        messages[name] = message
    }
    return messages, nil
}