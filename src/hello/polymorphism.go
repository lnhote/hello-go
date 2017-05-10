package main

type Animal interface {
	Say()
	GetName() string
	Introduce()
}
type Pet struct {
	Name string
	Inst interface{}
}

func (this *Pet) Say() {
	println("my name is ", this.Inst.(Animal).GetName())

	// This is not gonna work.
	println("my name is ", this.GetName())
}
func (this *Pet) GetName() string {
	return "pet"
}
func (this *Pet) Introduce() {
	println("I am a pet.")
}

type Dog struct {
	Pet
}

func (this *Dog) GetName() string {
	return this.Name
}
func (this *Dog) Introduce() {
	println("Introduce: I am a dog.")
}

type Cat struct {
	Pet
}

func (this *Cat) GetName() string {
	return this.Name
}
func (this *Cat) Introduce() {
	println("Introduce: I am a cat.")
}

func AnimalGreet(animal Animal) {
	animal.Introduce()
	animal.Say()
}
func NewCat() *Cat {
	this := &Cat{}
	this.Inst = this
	this.Name = "tomcat"
	return this
}
func NewDog() *Dog {
	this := &Dog{}
	this.Inst = this
	this.Name = "snoopy"
	return this
}

func main() {
	cat := NewCat()
	dog := NewDog()
	AnimalGreet(cat)
	AnimalGreet(dog)
}
