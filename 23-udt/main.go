package main

import "fmt"

func main() {

	p1 := Person{Name: "JP", Email: "Jp@Gmail.com", Contact: "9618558500", Status: "active",
		Address: struct {
			Line1   string
			City    string
			State   string
			Country string
			PinCode string
			Status  string
		}{Line1: "Street1", City: "Hyd", State: "Telangagna", Country: "India", PinCode: "501231", Status: "true"},
		SocialMedia: struct {
			Social map[string]string
		}{Social: map[string]string{"linkedin": "linkedin.com/jp", "twitte": "x.com/jpalapartjhi"}}}

	// map1 := map[string]string{"A": "1", "B": "2"}
	// map2 := make(map[string]string)
	// map2["A"] = "1"
	// map2["B"] = "2"

	// creating variable directly from a struct type
	var data struct {
		Location struct {
			Lag string
			Lan string
		}
		Errors  []error
		Records struct {
			Count int
			Valid int
		}
	}
	fmt.Println(data)

	// var data1 Data
	// fmt.Println(data1)
}

// type Data struct {
// 	Location struct {
// 		Lag string
// 		Lan string
// 	}
// 	Errors  []error
// 	Records struct {
// 		Count int
// 		Valid int
// 	}
// }

type Person struct {
	Name, Email, Contact, Status string

	Address struct {
		Line1, City, State, Country, PinCode, Status string
	}
	SocialMedia struct {
		Social map[string]string
	}
}

func (p *Person) PrintSM() error {
	if p.SocialMedia.Social != nil {
		for k, v := range p.SocialMedia.Social {
			fmt.Println("Account Type:", k, "Account Name:", v)
		}
	} else {
		return fmt.Errorf("nil map in sa object")
	}
	return nil

}

// type Users struct {
// 	Model // id, created_at,deleted_at,updated_at
// 	Name
// 	Email
// }
// // u1.ID

// type SocialMedia struct {
// 	Social map[string]string
// }

// type Address struct {
// 	Line1, City, State, Country, PinCode, Status string
// }
