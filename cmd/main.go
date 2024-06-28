package main

func main() {
	println("Hello World")

	app, err := NewApp().App("localhost:8088")
	if err != nil {
		return
	}

	err = app.Run()
	if err != nil {
		return
	}
}
