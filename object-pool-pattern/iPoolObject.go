package main

type iPoolObject interface {
	//This is any id which can be used to compare two different pool objects
	getID() string
}
