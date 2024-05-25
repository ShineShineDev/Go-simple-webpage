package main
import (
	"fmt"
	"net/http" 
	"io/ioutil"
	"os"
    "encoding/json"
	)
	
func home(w http.ResponseWriter,r *http.Request){
	file, err := os.Open("home.spidey")
    if err != nil {
        http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
        return
    }
    defer file.Close()
	 content, err := ioutil.ReadFile("home.spidey")
    if err != nil {
        http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
        return
    }
	w.Write(content)
}

func loginForm(w http.ResponseWriter,r *http.Request){
	file, err := os.Open("login.spidey")
    if err != nil {
        http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
        return
    }
    defer file.Close()
	 content, err := ioutil.ReadFile("login.spidey")
    if err != nil {
        http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
        return
    }
	w.Write(content)
}


type LoginData struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}
func loginSubmit(w http.ResponseWriter,r *http.Request){
	if r.Method == http.MethodPost {
        err := r.ParseForm()
        if err != nil {
            http.Error(w, "Unable to parse form", http.StatusBadRequest)
            return
        }
        email := r.FormValue("email")
        password := r.FormValue("password")

        data := LoginData{
            Email:    email,
            Password: password,
        }
        jsonResponse, err := json.Marshal(data)
        if err != nil {
            http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        w.Write(jsonResponse)
    } else {
        http.ServeFile(w, r, "form.html")
    }
}

func main(){	
	http.HandleFunc("/home",home)
	http.HandleFunc("/",loginForm)
	http.HandleFunc("/login",loginForm)
    http.HandleFunc("/submit", loginSubmit)
	fmt.Println("Starting server at  http://localhost:7912/")
	http.ListenAndServe(":7912", nil)		
}
