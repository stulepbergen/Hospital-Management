package admin

import (
	"HospitalManagement/database"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DefaultAdminService AdminService
	adminDB             = database.Database()
	CurrentAdmin        AdminAuth
	adminInfo           AdminInfo
)

type Doctor struct {
	Id          int
	FirstName   string
	LastName    string
	PhoneNumber string
	Position    string
}

type Doctors struct {
	Doctors []Doctor
}

type Patient struct {
	Id          int
	FirstName   string
	LastName    string
	PhoneNumber string
	Diagnoz     string
}

type Patients struct {
	Patients []Patient
}

type AdminService struct{}

type AdminAuth struct {
	id       int
	Login    string
	Password string
}

type AdminInfo struct {
	id          int
	FirstName   string
	LastName    string
	DateBirth   string
	PhoneNumber string
	Address     string
}

func (AdminService) AddAdmin(newAdmin AdminAuth) error {
	fmt.Println(newAdmin)
	_, e := adminDB.Exec(`INSERT INTO AdminAuth (login, password) VALUES (?, ?)`, newAdmin.Login, newAdmin.Password)
	Error(e)

	admins, err := adminDB.Query("SELECT id, login, password FROM AdminAuth")
	if err != nil {
		panic(err.Error())
	}

	for admins.Next() {
		var admin AdminAuth
		err = admins.Scan(&admin.id, &admin.Login, &admin.Password)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(admin.id, admin.Login, admin.Password)
	}

	fmt.Println("Add admin")

	return nil
}

func (AdminService) GetAdminInfo(admin AdminAuth) *AdminInfo {
	d1, e := adminDB.Query(`SELECT id FROM AdminAuth WHERE login=?`, admin.Login)
	Error(e)

	d1.Next()
	var id int
	d1.Scan(&id)

	d, e := adminDB.Query(`SELECT COUNT(id) FROM Admins WHERE id=?`, id)
	Error(e)

	d.Next()
	var count int
	d.Scan(&count)

	if count == 0 {
		return &AdminInfo{
			FirstName:   "",
			LastName:    "",
			DateBirth:   "",
			PhoneNumber: "",
			Address:     "",
		}
	}

	d2, e := adminDB.Query(`SELECT firstname, lastName, birthDate, phoneNumber, address FROM Admins Where id=?`, id)
	Error(e)

	d2.Next()
	d2.Scan(&adminInfo.FirstName, &adminInfo.LastName, &adminInfo.DateBirth, &adminInfo.PhoneNumber, &adminInfo.Address)

	return &adminInfo
}

func (AdminService) ChangeAdminInfo(admin AdminAuth, info AdminInfo) {
	d1, e := adminDB.Query(`SELECT id FROM AdminAuth WHERE login=?`, admin.Login)
	Error(e)

	d1.Next()
	var id int
	d1.Scan(&id)

	d2, e := adminDB.Query(`SELECT COUNT(id) FROM Admins WHERE id=?`, id)
	Error(e)

	d2.Next()
	var count int
	d2.Scan(&count)

	if count == 0 {
		_, e = adminDB.Exec(`INSERT INTO Admins (firstName, lastName, birthDate, phoneNumber, address) VALUES(?, ?, ?, ?, ?)`,
			info.FirstName, info.LastName, info.DateBirth, info.PhoneNumber, info.Address)
		Error(e)
		return
	}

	_, e = adminDB.Exec(`UPDATE Admins SET firstname=?, lastName=?, birthDate=?, phoneNumber=?, address=? WHERE id=?`,
		info.FirstName, info.LastName, info.DateBirth, info.PhoneNumber, info.Address, id)
	Error(e)

}

func (AdminService) HasAdmin(admin AdminAuth) bool {
	d, e := adminDB.Query(`SELECT COUNT(id) FROM AdminAuth WHERE login=? AND password=?`, admin.Login, admin.Password)
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
