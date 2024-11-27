package main

import "fmt"

func main() {

	p1 := Person{Name: "JP", Email: "Jp@Gmail.com", Contact: "9618558500", Status: "active", Address: Address{Line1: "Street1", City: "Hyd", State: "Telangagna", Country: "India", PinCode: "501231", Status: "true"}, SocialMedia: SocialMedia{Social: map[string]string{"linkedin": "linkedin.com/jp", "twitte": "x.com/jpalapartjhi"}}}

	p1.City = "Hyderabad"
	fmt.Println(p1.Address.City)
	fmt.Println(p1)
	err := p1.SocialMedia.PrintSM()
	if err != nil {
		println(err.Error())
	}
	p1.PrintSM()
	fmt.Println(p1.State)
	fmt.Println(p1.Address.Status)
}

type Person struct {
	Name, Email, Contact, Status string
	Address                      // promoted fields
	SocialMedia                  // promoted fields
}

// type Users struct {
// 	Model // id, created_at,deleted_at,updated_at
// 	Name
// 	Email
// }
// // u1.ID

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
	Line1, City, State, Country, PinCode, Status string
}
