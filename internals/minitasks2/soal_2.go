package minitask2

import "fmt"

type Person struct {
    Name    string
    Address string
    Phone   string
}

// Method Getter
func (p *Person) GetName() string {
    return p.Name
}

func (p *Person) GetAddress() string {
    return p.Address
}

func (p *Person) GetPhone() string {
    return p.Phone
}

// Method Greet
func (p *Person) Greet() string {
    return "Hello, my name is " + p.Name
}

// Method Setter
func (p *Person) SetName(newName string) {
    p.Name = newName
}

// Method Print 
func (p *Person) Print() string {
    return fmt.Sprintf("Name: %s, Address: %s, Phone: %s", p.Name, p.Address, p.Phone)
}

// Fungsi Constructor 
func NewPerson(name, address, phone string) *Person {
    return &Person{
        Name:    name,
        Address: address,
        Phone:   phone,
    }
}
