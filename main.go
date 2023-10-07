package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("Welcome to", conferenceName, "booking application.")
	fmt.Println("We have total of", conferenceTickets, "tickets &", remainingTickets, "are still left!!!")
	fmt.Println("Get your tickets here to attend.")

	fmt.Println(conferenceName)
}