package patient

import (
	"HospitalManagement/database"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DefaultPatientService PatientService
	patientDB             = database.Database()
	CurrentPatient        PatientAuth
	patientInfo           PatientInfo
	patientMedCard        PatientMedCard
	patientMyDoctor       PatientMyDoctor
)

type PatientInfo struct {
	id          int
	FirstName   string
	LastName    string
	DateBirth   string
	PhoneNumber string
	Address     string
}

type PatientAuth struct {
	id       int
	login    string
	password string
}

type PatientMedCard struct {
	id              int
	BloodGroup      string
	Diagnoz         string
	DateOfArrival   string
	DateOfDischarge string
}

type PatientMyDoctor struct {
	FirstName string
	LastName  string
	Position  string
}

type PatientService struct {
}

func Error(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func (PatientService) AddPatient(newPatient PatientAuth) error {
	_, e := patientDB.Exec(`INSERT INTO PatientAuth (login, password) VALUES (?, ?)`, newPatient.login, newPatient.password)
	Error(e)

	patients, err := patientDB.Query("SELECT id, login, password FROM PatientAuth")
	if err != nil {
		panic(err.Error())
	}
	for patients.Next() {
		var patient PatientAuth
		err = patients.Scan(&patient.id, &patient.login, &patient.password)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(patient.id, patient.login, patient.password)
	}

	fmt.Println("Add patient")

	return nil
}

func (PatientService) HasPatient(patient PatientAuth) bool {
	d, e := patientDB.Query(`SELECT COUNT(id) FROM PatientAuth WHERE login=? AND password=?`, patient.login, patient.password)
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

func (PatientService) GetPatientInfo(patient PatientAuth) *PatientInfo {
	d1, e := patientDB.Query(`SELECT id FROM PatientAuth WHERE login=?`, patient.login)
	Error(e)

	d1.Next()
	var id int
	d1.Scan(&id)

	d, e := patientDB.Query(`SELECT COUNT(id) FROM Patients WHERE id=?`, id)
	Error(e)

	d.Next()
	var count int
	d.Scan(&count)

	if count == 0 {
		return &PatientInfo{
			FirstName:   "",
			LastName:    "",
			DateBirth:   "",
			PhoneNumber: "",
			Address:     "",
		}
	}

	d2, e := patientDB.Query(`SELECT firstname, lastName, birthDate, phoneNumber, address FROM Patients Where id=?`, id)
	Error(e)

	d2.Next()
	d2.Scan(&patientInfo.FirstName, &patientInfo.LastName, &patientInfo.DateBirth, &patientInfo.PhoneNumber, &patientInfo.Address)

	return &patientInfo
}

func (PatientService) GetMedCardInfo(patient PatientAuth) *PatientMedCard {
	d1, e := patientDB.Query(`SELECT id FROM PatientAuth WHERE login=?`, patient.login)
	Error(e)

	d1.Next()
	var id int
	d1.Scan(&id)

	d, e := patientDB.Query(`SELECT COUNT(idPatient) FROM PatientMedCard WHERE idPatient=?`, id)
	Error(e)

	d.Next()
	var count int
	d.Scan(&count)

	if count == 0 {
		return &PatientMedCard{
			BloodGroup:      "",
			Diagnoz:         "",
			DateOfArrival:   "",
			DateOfDischarge: "",
		}
	}

	d2, e := patientDB.Query(`SELECT bloodGroup, diagnoz, dateOfArrival, dateOfDischarge FROM PatientMedCard Where idPatient=?`, id)
	Error(e)

	d2.Next()
	d2.Scan(&patientMedCard.BloodGroup, &patientMedCard.Diagnoz, &patientMedCard.DateOfArrival, &patientMedCard.DateOfDischarge)

	fmt.Println(patientMedCard)

	fmt.Println("MedCard Info!")
	return &patientMedCard
}

func (PatientService) GetPatientMyDoctor(patient PatientAuth) *PatientMyDoctor {
	d1, e := patientDB.Query(`SELECT id FROM PatientAuth WHERE login=?`, patient.login)
	Error(e)

	d1.Next()
	var id int
	d1.Scan(&id)
	fmt.Println(id)

	d2, e := patientDB.Query(`SELECT diagnoz FROM PatientMedCard Where idPatient=?`, id)
	Error(e)

	d2.Next()
	var diagnoz string
	d2.Scan(&diagnoz)
	fmt.Println(diagnoz)

	d3, e := patientDB.Query(`SELECT doctorPosition FROM Diagnoz WHERE diagnoz=?`, diagnoz)
	Error(e)

	d3.Next()
	var positionDoctor string
	d3.Scan(&positionDoctor)
	fmt.Println(positionDoctor)

	d4, e := patientDB.Query(`SELECT idDoctor FROM PositionDoctor WHERE position=?`, positionDoctor)
	Error(e)

	d4.Next()
	var idDoctor int
	d4.Scan(&idDoctor)

	d5, e := patientDB.Query(`SELECT firstName, lastName FROM Doctors WHERE id=?`, idDoctor)
	Error(e)

	d5.Next()
	d5.Scan(&patientMyDoctor.FirstName, &patientMyDoctor.LastName)
	patientMyDoctor.Position = positionDoctor
	fmt.Println(patientMyDoctor.FirstName, patientMyDoctor.LastName, patientMyDoctor.Position)
	fmt.Println("MyDoctor Info!")
	return &patientMyDoctor
}

func (PatientService) ChangePatientInfo(patient PatientAuth, info PatientInfo) {
	d1, e := patientDB.Query(`SELECT id FROM PatientAuth WHERE login=?`, patient.login)
	Error(e)

	d1.Next()
	var id int
	d1.Scan(&id)

	d2, e := patientDB.Query(`SELECT COUNT(id) FROM Patients WHERE id=?`, id)
	Error(e)

	d2.Next()
	var count int
	d2.Scan(&count)

	if count == 0 {
		_, e = patientDB.Exec(`INSERT INTO Patients (firstName, lastName, birthDate, phoneNumber, address) VALUES(?, ?, ?, ?, ?)`,
			info.FirstName, info.LastName, info.DateBirth, info.PhoneNumber, info.Address)
		Error(e)
		return
	}

	_, e = patientDB.Exec(`UPDATE Patients SET firstname=?, lastName=?, birthDate=?, phoneNumber=?, address=? WHERE id=?`,
		info.FirstName, info.LastName, info.DateBirth, info.PhoneNumber, info.Address, id)
	Error(e)

}

func (PatientService) ChangePatientMedCard(patient PatientAuth, medcard *PatientMedCard) {
	d1, e := patientDB.Query(`SELECT id FROM PatientAuth WHERE login=?`, patient.login)
	Error(e)

	d1.Next()
	var id int
	d1.Scan(&id)

	d2, e := patientDB.Query(`SELECT COUNT(idPatient) FROM PatientMedCard WHERE idPatient=?`, id)
	Error(e)
	medcard.id = id

	d2.Next()
	var count int
	d2.Scan(&count)

	fmt.Println("id: ", id, " count: ", count)
	if count == 0 {
		fmt.Println(medcard)
		_, e = patientDB.Exec(`INSERT INTO PatientMedCard (idPatient, bloodGroup, diagnoz, dateOfArrival, dateOfDischarge) VALUES(?, ?, ?, ?, ?)`,
			medcard.id, medcard.BloodGroup, medcard.Diagnoz, medcard.DateOfArrival, medcard.DateOfDischarge)
		Error(e)
		fmt.Println("INSERTED PATIENTMEDCARD")
		return
	}

	_, e = patientDB.Exec(`UPDATE PatientMedCard SET bloodGroup=?, diagnoz=?, dateOfArrival=?, dateOfDischarge=? WHERE idPatient=?`,
		medcard.BloodGroup, medcard.Diagnoz, medcard.DateOfArrival, medcard.DateOfDischarge, id)
	Error(e)

}
