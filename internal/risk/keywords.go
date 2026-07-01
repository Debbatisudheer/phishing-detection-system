package risk

var SubjectWeights = map[string]int{

	"urgent":             30,
	"security alert":     25,
	"account suspended":  30,
}

var BodyWeights = map[string]int{

	"login":          20,
	"verify account": 25,
	"click here":     20,
	"password":       15,
	"bank":           20,
}