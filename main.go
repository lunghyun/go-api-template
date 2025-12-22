package main

func main() {
	engin := MakeWebHandler()
	engin.Run(":8080")
}
