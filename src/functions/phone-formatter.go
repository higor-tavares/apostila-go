package main

import (
	"fmt"
	"regexp"
	"os"
)

type Phone string

func (phone Phone) format() string {
	match,_ := regexp.Match("\\d{11}", []byte(phone)) 
	if !match {
		fmt.Printf("%s doesn't is a valid phone number!\n", phone)
		os.Exit(1)
	}
	return fmt.Sprintf("(%s) %s-%s", string(phone)[:2], string(phone)[2:7], string(phone)[7:])
}

func main() {
	var phone1,phone2 Phone 
	phone1 = "88994382110"
	formated := phone1.format()
	fmt.Println(formated)
	phone2 = "88888"
	formated2 := phone2.format()
	fmt.Println(formated2)
}

