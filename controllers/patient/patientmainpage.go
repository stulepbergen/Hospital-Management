package patient

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

func LoginPatient(w http.ResponseWriter, r *http.Request) {
	fmt.Print("patients login")
	filename := "././hospitalWeb/patient-login/index.html"
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
	http.Redirect(w, r, "/patients-main-page", 307)
}

func SignUpPatient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SignUpPatients")
	newPatient := getPatient(r)
	fmt.Println("login: " + newPatient.login)
	fmt.Println("password: " + newPatient.password)
	has := DefaultPatientService.HasPatient(newPatient)

	if has {
		fmt.Println("You have already registered!")
		http.Redirect(w, r, "/patients-login", 307)
		return
	}

	err := DefaultPatientService.AddPatient(newPatient)
	Error(err)

	SignUp(w, r, "././hospitalWeb/patient-login/index.html", "Patient", err)
}

func getPatient(r *http.Request) PatientAuth {
	login := r.FormValue("login")
	password := r.FormValue("password")

	return PatientAuth{
		id:       0,
		login:    login,
		password: password,
	}
}

func PagePatient(w http.ResponseWriter, r *http.Request, filename string) {
	t, _ := template.ParseFiles(filename)
	t.Execute(w, DefaultPatientService.GetPatientInfo(CurrentPatient))
}

func AuthPatient(w http.ResponseWriter, r *http.Request) {
	patient := getPatient(r)
	hasPatient := DefaultPatientService.HasPatient(patient)

	if !hasPatient {
		http.Redirect(w, r, "/patients-login", 307)
		return
	}

	CurrentPatient = patient
	http.Redirect(w, r, "/patients-main-page", 307)
}

func UpdatePatient(w http.ResponseWriter, r *http.Request) {
	fname := r.FormValue("firstName")
	lname := r.FormValue("lastName")
	dateBirth := r.FormValue("dateBirth")
	phoneNumber := r.FormValue("phoneNumber")
	address := r.FormValue("address")
	position := r.FormValue("position")
	fmt.Println("save " + fname + " " + lname + " " + dateBirth + " " + phoneNumber + " " + address + " " + position)

	DefaultPatientService.ChangePatientInfo(CurrentPatient, PatientInfo{
		FirstName:   fname,
		LastName:    lname,
		DateBirth:   dateBirth,
		PhoneNumber: phoneNumber,
		Address:     address,
	})

	http.Redirect(w, r, "/patients-main-page", 307)
}

func MainPagePatient(w http.ResponseWriter, r *http.Request) {
	PagePatient(w, r, "././hospitalWeb/patient-page/patient-main-page.html")
	_ = database.Database()
}
