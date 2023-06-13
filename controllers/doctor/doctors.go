package doctor

import (
	"HospitalManagement/database"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DefaultDoctorService DoctorService
	doctorDB             = database.Database()
	CurrentDoctor        DoctorAuth
	doctorInfo           DoctorInfo
)

type DoctorInfo struct {
	id          int
	FirstName   string
	LastName    string
	DateBirth   string
	PhoneNumber string
	Address     string
	Position    string
}

type DoctorAuth struct {
	id       int
	login    string
	password string
}

type DoctorService struct {
}

type Patient struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Diagnoz     string
}

type Patients struct {
	Patients []Patient
}

func Error(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func (DoctorService) AddDoctor(newDoctor DoctorAuth) error {
	_, e := doctorDB.Exec(`INSERT INTO DoctorAuth (login, password) VALUES (?, ?)`, newDoctor.login, newDoctor.password)
	Error(e)

	doctors, err := doctorDB.Query("SELECT id, login, password FROM DoctorAuth")
	if err != nil {
		panic(err.Error())
	}
	for doctors.Next() {
		var doctor DoctorAuth
		err = doctors.Scan(&doctor.id, &doctor.login, &doctor.password)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(doctor.id, doctor.login, doctor.password)
	}

	fmt.Println("Add doctor")

	return nil
}

func (DoctorService) HasDoctor(doctor DoctorAuth) bool {
	d, e := doctorDB.Query(`SELECT COUNT(id) FROM DoctorAuth WHERE login=? AND password=?`, doctor.login, doctor.password)
	Error(e)

	var count int
	d.Next()
	d.Scan(&count)
	fmt.Println("count ", count)

	has := false

	if count != 0 {
		has = true
	}

	return has
}

func (DoctorService) GetDoctorInfo(doctor DoctorAuth) *DoctorInfo {
	d1, e := doctorDB.Query(`SELECT id FROM DoctorAuth WHERE login=?`, doctor.login)
	Error(e)

	d1.Next()
	var id int
	d1.Scan(&id)

	d2, e := doctorDB.Query(`SELECT firstname, lastName, birthDate, phoneNumber, address FROM Doctors Where id=?`, id)
	Error(e)

	d, e := doctorDB.Query(`SELECT position FROM PositionDoctor WHERE idDoctor=?`, id)
	d.Next()
	d.Scan(&doctorInfo.Position)

	d2.Next()
	d2.Scan(&doctorInfo.FirstName, &doctorInfo.LastName, &doctorInfo.DateBirth, &doctorInfo.PhoneNumber, &doctorInfo.Address)

	return &doctorInfo
}

func (DoctorService) ChangeDoctorInfo(doctor DoctorAuth, info DoctorInfo) {
	d1, e := doctorDB.Query(`SELECT id FROM DoctorAuth WHERE login=?`, doctor.login)
	Error(e)

	d1.Next()
	var id int
	d1.Scan(&id)

	d2, e := doctorDB.Query(`SELECT COUNT(id) FROM Doctors WHERE id=?`, id)
	Error(e)

	d2.Next()
	var count int
	d2.Scan(&count)

	if count == 0 {
		_, e = doctorDB.Exec(`INSERT INTO Doctors (firstname, lastName, address, phoneNumber, birthDate) VALUES(?, ?, ?, ?, ?)`,
			info.FirstName, info.LastName, info.Address, info.PhoneNumber, info.DateBirth)
		Error(e)
		_, e = doctorDB.Exec(`INSERT INTO PositionDoctor(idDoctor, position) VALUES(?, ?)`, id, info.Position)
		Error(e)
		return
	}

	_, e = doctorDB.Exec(`UPDATE Doctors SET firstname=?, lastName=?, birthDate=?, phoneNumber=?, address=? WHERE id=?`,
		info.FirstName, info.LastName, info.DateBirth, info.PhoneNumber, info.Address, id)
	Error(e)
	_, e = doctorDB.Exec(`UPDATE PositionDoctor SET idDoctor=?, position=?`, id, info.Position)
	Error(e)

}

func (DoctorService) GetPatients(doctor DoctorAuth) Patients {
	d1, e := doctorDB.Query(`SELECT id FROM DoctorAuth WHERE login=?`, doctor.login)
	Error(e)

	d1.Next()
	var id int
	d1.Scan(&id)

	d2, e := doctorDB.Query(`SELECT position FROM PositionDoctor WHERE idDoctor=?`, id)
	Error(e)
	d2.Next()
	var position string
	d2.Scan(&position)

	d3, e := doctorDB.Query(`SELECT diagnoz FROM Diagnoz WHERE doctorPosition=?`, position)
	Error(e)

	d3.Next()
	var diagnoz string
	d3.Scan(&diagnoz)

	d4, e := doctorDB.Query(`SELECT idPatient FROM PatientMedCard WHERE diagnoz=?`, diagnoz)
	Error(e)

	var patients Patients
	for d4.Next() {
		var idPatient int
		d4.Scan(&idPatient)

		d4, e := doctorDB.Query(`SELECT firstName, lastName, phoneNumber FROM Patients WHERE id=?`, idPatient)
		Error(e)

		d4.Next()
		var patient Patient
		d4.Scan(&patient.FirstName, &patient.LastName, &patient.PhoneNumber)
		patient.Diagnoz = diagnoz
		patients.Patients = append(patients.Patients, patient)
	}
	return patients
}
