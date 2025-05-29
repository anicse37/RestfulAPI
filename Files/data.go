package files

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type User struct {
	Id         int
	First_name string
	Last_name  string
	Email      string
	Gender     string
	Ip_address string
}
type Gtech []User

type UserData struct {
	JsonFile *json.Encoder
	Users    Gtech
}

func (e *UserData) GetAllUserData() Gtech {
	return e.Users
}

func GetDataFromJSON() (*os.File, error) {
	file, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return file, nil
}
func Setup() *UserData {

	File, err := GetDataFromJSON()
	if err != nil {
		fmt.Printf("could not open file %v, %v", File.Name(), err)
	}
	var emp Gtech
	err = json.NewDecoder(File).Decode(&emp)

	if err != nil {
		fmt.Print("rhnen")
	}
	return &UserData{
		JsonFile: json.NewEncoder(&Tape{File: File}),
		Users:    emp,
	}
}

type Tape struct {
	File *os.File
}

func (t *Tape) Write(data []byte) (n int, err error) {
	t.File.Truncate(0)
	t.File.Seek(0, io.SeekStart)
	t.File.Write(data)
	return
}
