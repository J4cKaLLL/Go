package partyrobot

import (
	"fmt"
	"strconv"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
	panic("Please implement the Welcome function")
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return "Happy birthday " + name + "! " + "You are now " + strconv.Itoa(age) + " years old!"
	panic("Please implement the HappyBirthday function")
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor string, direction string, distance float64) string {
	return fmt.Sprintf("Welcome to my party, %v!\nYou have been assigned to table %03d. Your table is %v, exactly %.1f meters from here.\nYou will be sitting next to %v.", name, table, direction, distance, neighbor)
	panic("Please implement the AssignTable function")
}
