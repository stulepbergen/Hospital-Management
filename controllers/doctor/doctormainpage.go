package doctor

import (
	"HospitalManagement/database"
	"fmt"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request, filename string) {
	t, err := template.ParseFiles(filename)
	Error(err)

	err = t.Execute(w, filename)
	Error(err)
}

func LoginDoctor(w http.ResponseWriter, r *http.Request) {
	fmt.Print("doctors login")
	filename := "././hospitalWeb/doctor-login/index.html"
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
	http.Redirect(w, r, "/doctors-main-page", 307)
}

func SignUpDoctor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SignUpDoctors")
	newDoctor := getDoctor(r)
	fmt.Println("login: " + newDoctor.login)
	fmt.Println("password: " + newDoctor.password)
	has := DefaultDoctorService.HasDoctor(newDoctor)

	if has {
		fmt.Println("You have already registered!")
		http.Redirect(w, r, "/doctors-login", 307)
		return
	}

	err := DefaultDoctorService.AddDoctor(newDoctor)
	Error(err)

	SignUp(w, r, "././hospitalWeb/doctor-login/index.html", "Doctor", err)
}

func getDoctor(r *http.Request) DoctorAuth {
	login := r.FormValue("login")
	password := r.FormValue("password")

	return DoctorAuth{
		id:       0,
		login:    login,
		password: password,
	}
}

func PageDoctor(w http.ResponseWriter, r *http.Request, filename string) {
	t, _ := template.ParseFiles(filename)
	t.Execute(w, DefaultDoctorService.GetDoctorInfo(CurrentDoctor))
}

func AuthDoctor(w http.ResponseWriter, r *http.Request) {
	doctor := getDoctor(r)
	hasDoctor := DefaultDoctorService.HasDoctor(doctor)

	if !hasDoctor {
		http.Redirect(w, r, "/doctors-login", 307)
		return
	}

	CurrentDoctor = doctor
	http.Redirect(w, r, "/doctors-main-page", 307)
}

func UpdateDoctor(w http.ResponseWriter, r *http.Request) {
	fname := r.FormValue("firstName")
	lname := r.FormValue("lastName")
	dateBirth := r.FormValue("dateBirth")
	phoneNumber := r.FormValue("phoneNumber")
	address := r.FormValue("address")
	position := r.FormValue("position")
	fmt.Println("save " + fname + " " + lname + " " + dateBirth + " " + phoneNumber + " " + address + " " + position)

	DefaultDoctorService.ChangeDoctorInfo(CurrentDoctor, DoctorInfo{
		FirstName:   fname,
		LastName:    lname,
		DateBirth:   dateBirth,
		PhoneNumber: phoneNumber,
		Address:     address,
		Position:    position,
	})

	http.Redirect(w, r, "/doctors-main-page", 307)
}

func MainPageDoctor(w http.ResponseWriter, r *http.Request) {
	PageDoctor(w, r, "././hospitalWeb/doctor-page/doctor-main-page.html")
	_ = database.Database()
}

func PatientsPageDoctor(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("././hospitalWeb/doctor-page/doctor-patients-page.html")
	t.Execute(w, DefaultDoctorService.GetPatients(CurrentDoctor))
}
