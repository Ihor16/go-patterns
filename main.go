package main

import (
	"bufio"
	"fmt"
	"os"
	"patterns/composite"
	"patterns/factory"
	"patterns/observer"
	"patterns/singleton"
	"patterns/state"
	"patterns/strategy"
	"strings"
)

func main() {
	//callObserver()
	//callComposite()
	//callFactory()
	//callSingleton()
	//callStrategy()
	callState()
}

func callObserver() {
	printHeader("observer")
	customer := observer.Customer{}
	store := observer.NewStore()
	store.Register(&customer)
	store.UpdateStatus("store is open")
	store.UpdateStatus("store closes soon")
}

func callComposite() {
	printHeader("composite")
	post := composite.NewPost()
	page := composite.NewPage()
	t := []composite.Traverser{post, page}
	article := composite.NewArticle(t)
	article.Traverse()
}

func callFactory() {
	printHeader("factory")
	c := getCreator()
	car := c.Create()
	car.Drive()
}

func getCreator() factory.Creator {
	carType := "small"
	switch carType {
	case "small":
		return factory.NewSmallCarFactory()
	case "big":
		return factory.NewBigCarFactory()
	default:
		return factory.NewSmallCarFactory()
	}
}

func callSingleton() {
	printHeader("singleton")
	cs := singleton.NewCommentService()
	cs.GetCommentById()
	ps := singleton.NewProjectService()
	ps.GetProjectById()
	is := singleton.NewIssueService()
	is.GetIssueById()
}

func callStrategy() {
	printHeader("strategy")
	cs := strategy.NewCompressionService()
	cs.Compress(os.File{})
}

func callState() {
	r := bufio.NewReader(os.Stdin)
	var input string
	p := state.NewPhone()
	for {
		printPhoneOptions()
		input = readInput(r)
		if input == "3" {
			break
		}
		switch input {
		case "1":
			p.PressPower()
		case "2":
			p.PressHome()
		default:
			fmt.Println("error: invalid option")
		}
	}
	fmt.Println("Bye")
}

// readInput reads user input and returns it without the trailing newline character
func readInput(r *bufio.Reader) string {
	in, _ := r.ReadString('\n')
	in = strings.TrimRight(in, "\n")
	return in
}

func printPhoneOptions() {
	fmt.Println()
	fmt.Println("1: Power button")
	fmt.Println("2: Home button")
	fmt.Println("3: Throw into ocean")
}

func printHeader(t string) {
	fmt.Println()
	fmt.Println(fmt.Sprintf("=== %s ===", t))
}
