package global

var (
	Operate  = make(chan int, 8)
	DockerID = make(chan string, 8)
)
