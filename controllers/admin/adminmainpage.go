package admin

import (
	"HospitalManagement/database"
	"fmt"
	"html/template"
	"net/http"
)

func Error(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func Login(w http.ResponseWriter, r *http.Request, filename string) {
	t, err := template.ParseFiles(filename)
	Error(err)

	err = t.Execute(w, filename)
	Error(err)
}

func LoginAdmin(w http.ResponseWriter, r *http.Request) {
	fmt.Print("admins login")
	filename := "././hospitalWeb/admin-login/index.html"
	Login(w, r, filename)
}

func SignUp(w http.ResponseWriter, r *http.Request, filename string, role string, err error) {
	if err != nil {
		fmt.Println(err)
		t, _ := template.ParseFiles(filename)

		t.ExecuteTemplate(w, filename, "New "+role+" Sign Up failure!")
		http.Redirect(w, r, "/", 307)
		return
	}
	t, _ := template.ParseFiles(filename)
	t.ExecuteTemplate(w, filename, "New "+role+" Sign Up Success!")
	fmt.Println("Redirecting")
	http.Redirect(w, r, "/admins-main-page", 307)
}

func SignUpAdmin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SignUpAdmins")
	newAdmin := getAdmin(r)
	fmt.Println("login: " + newAdmin.Login)
	fmt.Println("password: " + newAdmin.Password)
	has := DefaultAdminService.HasAdmin(newAdmin)

	if has {
		fmt.Println("You have already registered!")
		http.Redirect(w, r, "/admins-login", 307)
		return
	}

	err := DefaultAdminService.AddAdmin(newAdmin)
	Error(err)

	SignUp(w, r, "././hospitalWeb/admin-login/index.html", "Admin", err)
}

func getAdmin(r *http.Request) AdminAuth {
	login := r.FormValue("login")
	password := r.FormValue("password")

	return AdminAuth{
		id:       0,
		Login:    login,
		Password: password,
	}
}

func PageAdmin(w http.ResponseWriter, r *http.Request, filename string) {
	t, _ := template.ParseFiles(filename)
	t.Execute(w, DefaultAdminService.GetAdminInfo(CurrentAdmin))
}

func AuthAdmin(w http.ResponseWriter, r *http.Request) {
	admin := getAdmin(r)
	hasAdmin := DefaultAdminService.HasAdmin(admin)

	if !hasAdmin {
		http.Redirect(w, r, "/admins-login", 307)
		return
	}

	CurrentAdmin = admin
	http.Redirect(w, r, "/admins-main-page", 307)
}

func UpdateAdmin(w http.ResponseWriter, r *http.Request) {
	fname := r.FormValue("firstName")
	lname := r.FormValue("lastName")
	dateBirth := r.FormValue("dateBirth")
	phoneNumber := r.FormValue("phoneNumber")
	address := r.FormValue("address")
	fmt.Println("save " + fname + " " + lname + " " + dateBirth + " " + phoneNumber + " " + address)

	DefaultAdminService.ChangeAdminInfo(CurrentAdmin, AdminInfo{
		FirstName:   fname,
		LastName:    lname,
		DateBirth:   dateBirth,
		PhoneNumber: phoneNumber,
		Address:     address,
	})

	http.Redirect(w, r, "/admins-main-page", 307)
}

func MainPageAdmin(w http.ResponseWriter, r *http.Request) {
	PageAdmin(w, r, "././hospitalWeb/admin-page/admin-main-page.html")
	_ = database.Database()
}
