package patient

import (
	"HospitalManagement/database"
	"fmt"
	"html/template"
	"net/http"
)

func MedcardPage(w http.ResponseWriter, r *http.Request, filename string) {
	t, _ := template.ParseFiles(filename)
	t.Execute(w, DefaultPatientService.GetMedCardInfo(CurrentPatient))
}

func MedCardPagePatient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MedCardPage")
	MedcardPage(w, r, "././hospitalWeb/patient-page/patient-med-card.html")
	_ = database.Database()
}

func UpdatePatientMedCard(w http.ResponseWriter, r *http.Request) {
	blood := r.FormValue("bloodGroup")
	diagnoz := r.FormValue("diagnoz")
	dateOfArrival := r.FormValue("dateOfArrival")
	dateOfDischarge := r.FormValue("dateOfDischarge")

	DefaultPatientService.ChangePatientMedCard(CurrentPatient, &PatientMedCard{
		BloodGroup:      blood,
		Diagnoz:         diagnoz,
		DateOfArrival:   dateOfArrival,
		DateOfDischarge: dateOfDischarge,
	})

	http.Redirect(w, r, "/patients-medcard-page", 307)
}
