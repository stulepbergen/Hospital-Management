package admin

import (
	"html/template"
	"net/http"
)

func DeletePatient(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("patientId")

	_, e := adminDB.Exec(`DELETE FROM PatientMedCard WHERE idPatient=?`, id)
	Error(e)

	_, e = adminDB.Exec(`DELETE FROM PatientAuth WHERE id=?`, id)
	Error(e)

	_, e = adminDB.Exec(`DELETE FROM Patients WHERE id=?`, id)
	Error(e)

	http.Redirect(w, r, "/admins-patients-page", 307)
}

func (AdminService) GetPatients() Patients {
	d, e := adminDB.Query(`SELECT id, firstName, lastName, phoneNumber FROM Patients`)
	Error(e)

	var patients Patients
	for d.Next() {
		var patient Patient
		d.Scan(&patient.Id, &patient.FirstName, &patient.LastName, &patient.PhoneNumber)
		d1, e := adminDB.Query(`SELECT diagnoz FROM PatientMedCard WHERE idPatient=?`, patient.Id)
		Error(e)
		d1.Next()
		d1.Scan(&patient.Diagnoz)
		patients.Patients = append(patients.Patients, patient)
	}
	return patients
}

func PatientsPageAdmin(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("././hospitalWeb/admin-page/admin-patients-page.html")
	t.Execute(w, DefaultAdminService.GetPatients())
}
