package main

func main() {
	engin := MakeWebHandler()
	if err := engin.Run(":8080"); err != nil {
		panic(err)
	}
}
