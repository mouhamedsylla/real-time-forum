package auth

type Login struct{
	email string
	password string
}

type Register struct{
	nickname string
	age int
	gender string
	firstName string
	lastName string
	email string
	password string
}
