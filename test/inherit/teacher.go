package inherit


type People struct {
	Name string
	Age int
}

type Teacher struct {
	Classroom string
	People
}
