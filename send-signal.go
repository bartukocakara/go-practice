package main

func workerRoutine(ch chan struct{}) {
	// Receive message from main program.
	<-ch
	println("Signal Received")

	// Send a message to the main program.
	close(ch)
}

func main() {
	//Create channel
	ch := make(chan struct{})

	//define workerRoutine
	go workerRoutine(ch)

	// Send signal to worker goroutine
	ch <- struct{}{}

	// Receive a message from the workerRoutine.
	<-ch
	println("Signal Received")

}