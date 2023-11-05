package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// var conferenceName string= "Go Conference"
// short hand
var conferenceName = "Go Conference"
//can't make short hand to const 
const conferenceTickets uint= 50

var remainingTickets uint= 50

// slice
// var bookings []string
// or
// var bookings =make([]map[string]string,0)
var bookings =make([]UserData,0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
	
}

var wg = sync.WaitGroup{

}

func main() {
	




	//%T is Type
	//%v is variable

	greetUsers()

	fmt.Printf("conferenceTickets is %T,remainingTickets is %T,conferenceName is %T\n",conferenceTickets,remainingTickets,conferenceName)

	// fmt.Printf("Welcome to %v booking application\n",conferenceName)
	// fmt.Printf("We have total of %v tickets and %v are still avaliable.\n",conferenceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// array
	// var bookings [50]string



	
		
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber :=helper.ValidateUserInput(firstName, lastName, email, userTickets,remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			
			bookTicket( userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket( userTickets, firstName, lastName, email)
			

			// array
			// bookings[0] = firstName + " " + lastName
	
			// slices
	
			// array
			// fmt.Printf("The whole array:%v\n" , bookings)
			// fmt.Printf("The first value: %v\n",bookings[0])
			// fmt.Printf("The type: %T\n",bookings)
			// fmt.Printf("Array length: %v\n",len(bookings))
	
			// slices
			fmt.Printf("The whole slice:%v\n" , bookings)
			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("slice type: %T\n", bookings)
			fmt.Printf("slice length: %v\n", len(bookings))
	
			firstNames :=getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
	
			 noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				// end program
				fmt.Println("Our conference is booked out.Comeback next year.")
			
			}
		} else {
			if !isValidName {
				println("Firstname or lastname you enter is too short,at least need to have 2 characters")
			}

			if !isValidEmail {
				println("Email address you enter is invalid,check for @")
			}

			if !isValidTicketNumber {
				println("number of tickets you entered is invalid")
			}
			
		}
		wg.Wait()
}

	// city := "London"

	// switch city {
	// 	case "New York":
	// 	// execute code for booking NewYork Tickets
	// 	case "Singapore", "Myanmar":
		
	// 		// execute code for booking Singapore and Myanmar Tickets
	// 	default:
	// 		fmt.Print("No valid city selected")
	// }



// else if userTickets == remainingTickets {
// 	// do something
// }


func greetUsers(){
	
	fmt.Printf("Welcome to %v from function\n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaliable.\n",conferenceTickets,remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		// var firstName =  names[0]
		// firstNames = append(firstNames, booking["Firstname"])
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames

}



func getUserInput() (string, string, string, uint) {
	var firstName string
		var lastName string
		var email string
		var userTickets uint
		
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your Email address:")
		fmt.Scan(&email)

		//& is pointer 

		// fmt.Println(remainingTickets)
		// fmt.Println(&remainingTickets)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		return firstName, lastName, email, userTickets
}

func bookTicket( userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// creat a map for user
	// var userData = make(map[string]string)
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets),10)

	bookings=append(bookings , userData)
	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets.You will receive a confirmation email at %v\n",
	firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets remaing for %v\n", remainingTickets,conferenceName)
	
}

func sendTicket(userTickets uint, firstName string, lastName string,email string) {
	time.Sleep(10 * time.Second)
	var ticket = 	fmt.Sprintf("%v tickets for %v %v\n",userTickets,firstName,lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("###############")

	wg.Done()
}