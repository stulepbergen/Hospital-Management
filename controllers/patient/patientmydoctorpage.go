package patient

import (
	"HospitalManagement/database"
	"html/template"
	"net/http"
)

func MyDoctorPage(w http.ResponseWriter, r *http.Request, filename string) {
	t, _ := template.ParseFiles(filename)
	t.Execute(w, DefaultPatientService.GetPatientMyDoctor(CurrentPatient))
}

func MyDoctorPagePatient(w http.ResponseWriter, r *http.Request) {
	MyDoctorPage(w, r, "././hospitalWeb/patient-page/patient-mydoctor-page.html")
	_ = database.Database()
}
