package main

import "fmt"

type User struct {
	Name  string
	Email string
	Sms   string
}

type UserNotifier interface {
	SendMessage(user *User, message string) error
}

type EmailNotifier struct {
}

func (notifier EmailNotifier) SendMessage(user *User, message string) error {
	_, err := fmt.Printf("sending email to  %s with content %s\n", user.Name, message)
	return err
}

type SmsNotifier struct {
}

func (notifier SmsNotifier) SendMessage(user *User, message string) error {
	_, err := fmt.Printf("sending SMS to  %s with content %s\n", user.Name, message)
	return err
}
func main() {
	user1 := &User{"yamuna", "yamuna.gopisetty@gmail.com", "hello yamuna"}

	fmt.Printf("welcome %s\n", user1.Name)

	user2 := &User{"sai", "sai.gopisetty@gmail.com", "hello sai"}

	fmt.Printf("welcome %s\n", user2.Name)

	var userNotifier UserNotifier

	userNotifier = EmailNotifier{}
	userNotifier.SendMessage(user1, "Interface all the way")

	userNotifier = SmsNotifier{}
	userNotifier.SendMessage(user2, "Interface all the way")
}
