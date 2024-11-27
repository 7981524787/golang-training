package main

import "fmt"

func main() {

	p1 := Person{Name: "JP", Email: "Jp@Gmail.com", Contact: "9618558500", Address: Address{Line1: "Street1", City: "Hyd", State: "Telangagna", Country: "India", PinCode: "501231"}, SocialMedia: SocialMedia{Social: map[string]string{"linkedin": "linkedin.com/jp", "twitte": "x.com/jpalapartjhi"}}}

	fmt.Println(p1.Address.City)
	fmt.Println(p1)
	err := p1.SocialMedia.PrintSM()
	if err != nil {
		println(err.Error())
	}
}

type Person struct {
	Name, Email, Contact string
	Address              Address // composition
	SocialMedia          SocialMedia
}

type SocialMedia struct {
	Social map[string]string
}

func (sa *SocialMedia) PrintSM() error {
	if sa.Social != nil {
		for k, v := range sa.Social {
			fmt.Println("Account Type:", k, "Account Name:", v)
		}
	} else {
		return fmt.Errorf("nil map in sa object")
	}
	return nil

}

type Address struct {
	Line1, City, State, Country, PinCode string
}
