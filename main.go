package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/anikhasibul/democracy/submit"
	"github.com/anikhasibul/queue"
)

func main() {
	q := queue.New(20)
	fmt.Println("Started....")
	for i := 0; 1 < 2000; i++ {
		q.Add()
		go func() {
			defer q.Done()
			email := genEmail()
			password := genPassword()
			projectID := genProjectID()
			submit.NewVote(email, password, projectID)
		}()
	}
	q.Wait()
}

func genEmail() string {
	randStr := strconv.FormatInt(time.Now().UnixNano(), 36)
	return randStr + "@gmail.com"
}

func genPassword() string {
	return "12345678"
}

func genProjectID() string {
	return os.Args[1]
}
