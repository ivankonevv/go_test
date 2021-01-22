package main

import (
	"fmt"
	"net/http"
	"html/template"
)
type User struct {
	Name string
	Age uint16
	Money int16
	AvgGrades, Happiness float64
	Hobbies []string
}
func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money " +
		"equal: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}
func homePage(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, 100,4.3, 0.9, []string{
		"Football", "Skate"}}
	//bob.setNewName("Alex")
	//fmt.Fprintf(w, bob.getAllInfo())
	tmpl, _ := template.ParseFiles("templates/homePage.html")
    tmpl.Execute(w, bob)
}

func contactsPage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "hello")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contactsPage)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}