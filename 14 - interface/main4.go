package main

import "fmt"

type (
	Notifier interface {
		Send(message string) error
	}
	EmailNotifier struct {
		Email string
	}
	SMSNotifier struct {
		Phone string
	}
	SlackNotifier struct {
		Channel string
	}
	TelgramNotifier struct {
		ChatID string
	}
)

func (e *EmailNotifier) Send(message string) error {
	fmt.Printf("Sending email to %s: %s\n", e.Email, message)
	return nil
}

func (s *SMSNotifier) Send(message string) error {
	fmt.Printf("Sending SMS to %s: %s\n", s.Phone, message)
	return nil
}

func (s *SlackNotifier) Send(message string) error {
	fmt.Printf("Sending Slack to %s: %s\n", s.Channel, message)
	return nil
}

func (t *TelgramNotifier) Send(message string) error {
	fmt.Printf("Sending Telgram to %s: %s\n", t.ChatID, message)
	return nil
}

func SendNotification(n Notifier, message string) {
	err := n.Send(message)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	fmt.Println("Message sent successfully")
}

func main() {

	notifiers := []Notifier{
		&EmailNotifier{Email: "user1@example.com"},
		&SMSNotifier{Phone: "1234567890"},
		&EmailNotifier{Email: "user2@example.com"},
	}
	message := "Hello, this is a test message"

	for _, notifier := range notifiers {
		SendNotification(notifier, message)
	}
}
