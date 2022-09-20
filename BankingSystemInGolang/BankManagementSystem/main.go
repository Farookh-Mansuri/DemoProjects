package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

type BankUser struct {
	Id       int
	Name     string
	Password string
	LoggedAt time.Time
}
type autoInc struct {
	sync.Mutex
	id int 
}

func (a *autoInc) ID() (id int) {
	a.Lock()
	defer a.Unlock()

	id = a.id
	a.id++
	return
}

var ai autoInc // global instance

var choice string
var userName string
var password string

// Login
func login_account() {
	// user := BankUser{}
	// user.Id = 1122
	// user.Name = "Peter"
	// user.Password = "Parker"
	// user.LoggedAt = time.Now()
	// user := []BankUser{
	// 	{Id: 2, Name: "john", Password: "1234"},
	// 	{Id: 3, Name: "kim", Password: "abc"},
	// 	{Id: 4, Name: "david", Password: "0011"},
	// }
	//...................................
	//Writing struct type to a JSON file
	//...................................
	// content, err := json.Marshal(user)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = ioutil.WriteFile("bankuser.json", content, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//...................................
	//Reading into struct type from a JSON file
	content, err := json.Marshal("bankuser.json")
	if err != nil {
		fmt.Println(err)
	}
	//...................................
	content, err = ioutil.ReadFile("bankuser.json")
	if err != nil {
		log.Fatal(err)
	}
	validateUser := BankUser{}
	err = json.Unmarshal(content, &validateUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Id:%d, Name:%s, Password:%s, LoggedAt:%v", validateUser.Id, validateUser.Name, validateUser.Password, validateUser.LoggedAt)
	fmt.Println("Please Enter User Name: ")
	fmt.Scan(&userName)
	fmt.Println("Please Enter Password: ")
	fmt.Scan(&password)
	// user := BankUser{Name: userId, Password: password}

	// f, err := os.Open("bankuser.json")
	// if err != nil {
	// 	fmt.Println("User Validation Error!")
	// }

	// var userData BankUser
	// err = json.NewDecoder(f).Decode(&userData)
	// if err != nil {
	// 	fmt.Println("Error in file")
	// }

	if validateUser.Name == userName && validateUser.Password == password {
		fmt.Println("Welcome Mr.", userName)
		fmt.Println("\n\n Please select below options from the menu:")
		fmt.Print("Press 1 to Deposit account\n")
		fmt.Print("Press 2 to Withdraw money\n")
		fmt.Print("Press 3 to Display Account \n")
		fmt.Print("Press 4 to Quit Account \n")
		fmt.Print("Enter your choice:\n")
		fmt.Scan(&choice)
	} else {
		fmt.Println("\n\n Incorrect UserId or Password \n ")
	}
	switch {
	case choice == "1":
		deposit_money()
	case choice == "2":
		withdraw_money()
	case choice == "3":
		display_account()
	case choice == "4":
		os.Exit(3)
	default:
		fmt.Println("Invalid")
	}

}

// Method to open the account
func open_account() {
	fmt.Println("Please enter userNam:")
	fmt.Scan(&userName)
	fmt.Println("Please enter password:")
	fmt.Scan(&password)
	newUser := []BankUser{
		{Id: ai.ID(), Name: userName, Password: password},
	}

	file, err := os.OpenFile("bankuser.json", os.O_RDWR, 0644)
		defer file.Close()

		//read file and unmarshall json file to slice of users
		b, err := ioutil.ReadAll(file)
		
		err = json.Unmarshal(b, &BankUser.Users)
		
		max := 0

		//generation of id(last id at the json file+1)
		for _, usr := range alUsrs.Users {
			if usr.Id > max {
				max = usr.Id
			}
		}
		id := max + 1
		newUser.Id = id

		//appending newUser to slice of all Users and rewrite json file
		alUsrs.Users = append(alUsrs.Users, newUser)
		newUserBytes, err := json.MarshalIndent(&alUsrs.Users, "", " ")
		outil.WriteFile("list.json", newUserBytes, 0666)
	//...................................
	//Writing struct type to a JSON file
	//...................................
	content, err := json.Marshal(newUser)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("bankuser.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Account Created..", newUser)
	fmt.Println("Please Login Your Accounr..")
	login_account()
}

// Method to deposit the account
func deposit_money() {
	//int amount
	amount := 9500
	balance := 100
	fmt.Println("Enter how much money you want to deposit: ", amount)
	balance += amount
	fmt.Println("\n Available Balance: ", balance)
}

// Method to display the account
func display_account() {
	json_data, err := json.Marshal("bankuser.json")
	if err != nil {

		log.Fatal(err)
	}
	fmt.Println(string(json_data))
}

// Method to withdraw the account
func withdraw_money() {
	//float amount
	amount := 3200
	balance := 1000
	fmt.Println("Enter how much money you want to withdraw: ", amount)
	balance -= amount
	fmt.Println("\n Available balance: ", balance)
}

func main() {

	var choice int
	fmt.Println("Welcome to John ! \n\n Please select an option from the menu:")
	fmt.Print("Press 1 to Login account \n")
	fmt.Print("Press 2 to Open account \n")
	fmt.Print("Enter your choice:\n")
	for {
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Invalid choice: err:", err)
			continue
		}
		switch {
		case choice == 1:
			login_account()
		case choice == 2:
			open_account()
		case choice == 3:
			os.Exit(3)
		default:
			fmt.Println("Invalid")
		}
	}
}
