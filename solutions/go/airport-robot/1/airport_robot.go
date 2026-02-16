package airportrobot

type Greeter interface {
 	LanguageName() string   
	Greet(visitorName string) string
}

// SayHello accepts a visitor's name and a Greeter implementation
// and returns the appropriate greeting
func SayHello(name string, greeter Greeter) string {
    return greeter.Greet(name)
}

// Italian implements Greeter for Italian
type Italian struct{}

func (ig Italian) LanguageName() string {
    return "Italian"
}

func (ig Italian) Greet(name string) string {
    return "I can speak " + ig.LanguageName() + ": Ciao " + name + "!"
}

// Portuguese implements Greeter for Portuguese
type Portuguese struct{}

func (pg Portuguese) LanguageName() string {
    return "Portuguese"
}

func (pg Portuguese) Greet(name string) string {
    return "I can speak " + pg.LanguageName() + ": Olá " + name + "!"
}