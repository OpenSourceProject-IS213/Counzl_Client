package users

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/tidwall/buntdb"
	"log"
	"utilities/crypto"
)

const (
	userBucket = "./database/l_user.db"
)

type User struct {
	ID       string
	Password string
	Online   string
}

func CreateUser(id, username, password, online string) bool {
	hashed_pw, _ := crypto.HashPassword(password)
	user := &User{id, hashed_pw, online}

	u := string(user.serialize())
	db, err := buntdb.Open(userBucket)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(username, u, nil)
		return err
	})
	fmt.Println("New user created with username: '" + username + "'")
	return true
}

func FetchUser(username string) (string, string, string) {
	var pword string
	var id string
	var online string
	db, err := buntdb.Open(userBucket)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(username)
		if err != nil {
			return err
		}
		tmp := deserialize([]byte(val))
		pword = tmp.Password
		id = tmp.ID
		online = tmp.Online
		return nil
	})
	return pword, id, online
}

func PrintLocalUsers() {
	db, err := buntdb.Open(userBucket)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	erro := db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("", func(key, value string) bool {
			if key != "" {
				val := deserialize([]byte(value))
				fmt.Printf("Username: %s, ID: %s, Online: %s\n", key, val.ID, val.Online)
				return true
			} else {
				return false
			}
		})

		return err
	})
	fmt.Print(erro)
}

func (u *User) serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(u)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func deserialize(d []byte) *User {
	var user User

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&user)
	if err != nil {
		log.Panic(err)
	}
	return &user
}
