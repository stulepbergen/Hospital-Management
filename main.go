package main

import (
	"log"
	"net/http"

	"github.com/stulepbergen/Hospital-Management/controllers"
	"github.com/stulepbergen/Hospital-Management/controllers/admin"
	"github.com/stulepbergen/Hospital-Management/controllers/doctor"
	"github.com/stulepbergen/Hospital-Management/controllers/patient"

	"github.com/gorilla/mux"
)

func main() {
	rout := mux.NewRouter()
	fs := http.FileServer(http.Dir("./hospitalWeb/main"))
	rout.Handle("./hospitalWeb/main/", fs)

	rout.HandleFunc("/", controllers.Main)

	rout.HandleFunc("/doctors-sign-up", doctor.SignUpDoctor)
	rout.HandleFunc("/doctors-login", doctor.LoginDoctor)
	rout.HandleFunc("/doctors-auth", doctor.AuthDoctor)
	rout.HandleFunc("/doctors-save", doctor.UpdateDoctor)
	rout.HandleFunc("/doctors-main-page", doctor.MainPageDoctor)
	rout.HandleFunc("/doctors-patients-page", doctor.PatientsPageDoctor)

	rout.HandleFunc("/patients-sign-up", patient.SignUpPatient)
	rout.HandleFunc("/patients-login", patient.LoginPatient)
	rout.HandleFunc("/patients-auth", patient.AuthPatient)
	rout.HandleFunc("/patients-save", patient.UpdatePatient)
	rout.HandleFunc("/patients-main-page", patient.MainPagePatient)
	rout.HandleFunc("/patients-medcard-page", patient.MedCardPagePatient)
	rout.HandleFunc("/patients-medcard-save", patient.UpdatePatientMedCard)
	rout.HandleFunc("/patients-mydoctor-page", patient.MyDoctorPagePatient)

	rout.HandleFunc("/admins-sign-up", admin.SignUpAdmin)
	rout.HandleFunc("/admins-login", admin.LoginAdmin)
	rout.HandleFunc("/admins-auth", admin.AuthAdmin)
	rout.HandleFunc("/admins-main-page", admin.MainPageAdmin)
	rout.HandleFunc("/admins-save", admin.UpdateAdmin)
	rout.HandleFunc("/admins-doctors-page", admin.DoctorsPageAdmin)
	rout.HandleFunc("/admins-patients-page", admin.PatientsPageAdmin)
	rout.HandleFunc("/admin-delete-patient", admin.DeletePatient)
	rout.HandleFunc("/admin-delete-doctor", admin.DeleteDoctor)

	log.Fatal(http.ListenAndServe(":9090", rout))
}
