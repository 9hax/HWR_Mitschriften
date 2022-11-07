package main

/* This must ge run with the Command Line flag GO111MODULE=off since it is not a proper module but requires global imports. */

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
)

type User struct {
	name, authHash, contact string
}

func main() {
	fmt.Println("Uebungsstunde 2 zu Go-Programmierung in Einführung in die Programmierung.")
	/* 4.11.2022, 08:45 Uhr. */

	users := make(map[string]User)

	users["pvierkorn"] = User{"Paul Friedrich Vierkorn", HashPassword("YouWishThisWasARealPassword"), "webmaster@activate-linux.org"}
	users["mmusterperson"] = User{"Maximilian Musterperson", HashPassword("ThisCouldHaveBeenGoodToKnow"), "max@musterperson.org"}
	users["areese"] = User{"Aamina Reese", HashPassword("IHaveCommittedToPublishing!"), "areese@outlook.zl"}
	users["ebaxter"] = User{"Euan Baxter", HashPassword("YouCouldHaveYourAdInserted!"), "ebaxter@freemail.ru"}
	users["bbassett"] = User{"Bob Bassett", HashPassword("ContactMeAtTheAboveAddress."), "bobbassett@dogs.to"}

	//fmt.Println(users)

	for {

		print("Login: ")
		var authUserId string
		fmt.Scanf("%s", &authUserId)

		user := users[authUserId]

		if checkUserPassword(user) {
			authenticate(user)
		} else {
			println("Returning to Login...")
		}
	}
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func checkUserPassword(user User) bool {
	for i := 0; i < 3; i++ {
		print("Password: ")
		password, _ := terminal.ReadPassword(0)
		println()
		if CheckPasswordHash(string(password), user.authHash) {
			return true
		}
		println("Password does not match.")
	}
	println("Error: 3 incorrect attempts.")
	return false
}

func authenticate(user User) {
	fmt.Printf("User %s authenticated!\n", user.name)
	os.Exit(0)
}
