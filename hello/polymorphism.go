package main

// golang don't realy support base class/inheritance + polymorphism
// this demo shows 2 way to do it

type Animal interface {
	// we will play inheritance on this function
	Say()

	// we will do polymorphism on this function
	GetName() string

	// golang's polymorphism can do this
	Introduce()
}
type EatFunc func() string

type Pet struct {
	Name string

	// method 1: the class way
	//      Need a struct pointer to the real obj
	//      Need a Interface which contains the function(e.g. GetName)
	Inst interface{}

	// method 2:
	//      Function/closure pointer
	//      Don't need interface, use pointer
	Eat EatFunc
}

func (pet *Pet) Say() {
	// No, no, this is not gonna work, although it's beautiful.
	println("my name is", pet.GetName())

	// method 1:
	println("my name is", pet.Inst.(Animal).GetName())

	// method 2:
	println("I eat", pet.Eat())
}
func (pet *Pet) GetName() string {
	return "pet"
}
func (pet *Pet) Introduce() {
	println("I am a pet.")
}

type Dog struct {
	Pet
}

func (dog *Dog) GetName() string {
	return dog.Name
}
func (dog *Dog) Introduce() {
	println("Introduce: I am a dog.")
}

type Cat struct {
	Pet
}

func (cat *Cat) GetName() string {
	return cat.Name
}
func (cat *Cat) Introduce() {
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
	this.Eat = func() string {
		return "fish, rat."
	}
	return this
}
func NewDog() *Dog {
	this := &Dog{}
	this.Inst = this
	this.Name = "snoopy"
	this.Eat = func() string {
		return "bone, meat."
	}
	return this
}

func main() {
	cat, dog := NewCat(), NewDog()
	AnimalGreet(cat)
	AnimalGreet(dog)
}
