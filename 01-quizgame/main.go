package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	duration := flag.Duration("time", 30*time.Second, "time limit for a quiz")
	filepath := flag.String("file", "problems.csv", "filepath")

	flag.Parse()
	p := Parser{
		fileName:  *filepath,
		timeLimit: *duration,
	}

	records, err := p.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}
	q := p.TokeniseTheQuiz(records)
	sc := &Score{}
	timer := time.NewTimer(p.timeLimit)

	an := make(chan int)
	for _, r := range q {
		go func() {
			var userInput int
			fmt.Printf("What is %v\n", r.question)
			_, err := fmt.Scan(&userInput)
			if err != nil {
				return
			}
			an <- userInput
		}()

		select {
		case ans := <-an:
			if ans == r.ans {
				AddRight(sc)
			} else {
				AddWrong(sc)
			}
		case <-timer.C:
			fmt.Println("time is up")
			fmt.Printf("your  score is %v / %v\n", sc.right, len(q))
			return
		}

	}
	fmt.Printf("your total score is %v / % v\n", sc.right, len(q))

}
func AddRight(s *Score) {
	s.right++
}
func AddWrong(s *Score) {
	s.wrong++
}
