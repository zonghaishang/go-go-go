package main

import "log"

func main() {

	var test = "default"

	switch test {
	case "default":
		log.Print("default")
	case "other":
		log.Print("other")
	default:
		log.Print("case default")
	}

	var n = 1

	switch {
	case n < 2:
		log.Println("< 2")
		log.Println("test fallthrough")
		fallthrough
	case n < 6:
		log.Println("< 6")
	case n == 10:
		log.Println("== 10")
	default:
		log.Printf("default: %d", n)
	}

}
