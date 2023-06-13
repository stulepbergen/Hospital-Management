package admin

import (
	"html/template"
	"net/http"
)

func DeleteDoctor(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("doctorId")

	_, e := adminDB.Exec(`DELETE FROM PositionDoctor WHERE idDoctor=?`, id)
	Error(e)

	_, e = adminDB.Exec(`DELETE FROM DoctorAuth WHERE id=?`, id)
	Error(e)

	_, e = adminDB.Exec(`DELETE FROM Doctors WHERE id=?`, id)
	Error(e)

	http.Redirect(w, r, "/admins-doctors-page", 307)
}

func (AdminService) GetDoctors() Doctors {
	d, e := adminDB.Query(`SELECT id, firstName, lastName, phoneNumber FROM Doctors`)
	Error(e)

	var doctors Doctors
	for d.Next() {
		var doctor Doctor
		d.Scan(&doctor.Id, &doctor.FirstName, &doctor.LastName, &doctor.PhoneNumber)
		d1, e := adminDB.Query(`SELECT position FROM PositionDoctor WHERE idDoctor=?`, doctor.Id)
		Error(e)
		d1.Next()
		d1.Scan(&doctor.Position)
		doctors.Doctors = append(doctors.Doctors, doctor)
	}
	return doctors
}

func DoctorsPageAdmin(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("././hospitalWeb/admin-page/admin-doctors-page.html")
	t.Execute(w, DefaultAdminService.GetDoctors())
}
