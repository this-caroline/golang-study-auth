package main

import (
	"fmt"
	"golang-study-auth/users"
	"html/template"
	"net/http"
)

func getSignInPage(w http.ResponseWriter, r *http.Request) {
	templating(w, "signIn.html", nil)
}

func getSignUpPage(w http.ResponseWriter, r *http.Request) {
	templating(w, "signUp.html", nil)
}

func templating(w http.ResponseWriter, fileName string, data interface{}) {
	t, _ := template.ParseFiles(fileName)
	t.ExecuteTemplate(w, fileName, data)
}

func signInUser(w http.ResponseWriter, r *http.Request) {
	newUser := getUser(r)
	ok := users.DefaultUserService.VerifyUser(newUser)
	if !ok {
		fileName := "signIn.html"
		t, _ := template.ParseFiles(fileName)
		t.ExecuteTemplate(w, fileName, fmt.Sprintf("%s Sign-in Failure.", newUser.Email))
		return
	}
	fileName := "signIn.html"
	t, _ := template.ParseFiles(fileName)
	t.ExecuteTemplate(w, fileName, fmt.Sprintf("%s Sign-in Success.", newUser.Email))
	return
}

func signUpUser(w http.ResponseWriter, r *http.Request) {
	newUser := getUser(r)
	err := users.DefaultUserService.CreateUser(newUser)
	if err != nil {
		fileName := "signUp.html"
		t, _ := template.ParseFiles(fileName)
		t.ExecuteTemplate(w, fileName, fmt.Sprintf("%s Sign-up Failure.", newUser.Email))
		return
	}
	fileName := "signUp.html"
	t, _ := template.ParseFiles(fileName)
	t.ExecuteTemplate(w, fileName, fmt.Sprintf("%s Sign-up Success.", newUser.Email))
	return
}

func getUser(r *http.Request) users.User {
	email := r.FormValue("email")
	password := r.FormValue("password")
	return users.User{
		Email:    email,
		Password: password,
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/sign-in":
		signInUser(w, r)
	case "/sign-up":
		signUpUser(w, r)
	case "/sign-in-form":
		getSignInPage(w, r)
	case "/sign-up-form":
		getSignUpPage(w, r)
	}
}

func main() {
	http.HandleFunc("/", userHandler)
	http.ListenAndServe("localhost:5000", nil)
}
