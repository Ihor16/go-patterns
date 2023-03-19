package main

import (
	"fmt"
	"os"
	"patterns/composite"
	"patterns/factory"
	"patterns/observer"
	"patterns/singleton"
	"patterns/strategy"
)

func main() {
	//callObserver()
	//callComposite()
	//callFactory()
	//callSingleton()
	callStrategy()
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
	f := getCreator()
	car := f.Create()
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

func printHeader(t string) {
	fmt.Println()
	fmt.Println(fmt.Sprintf("=== %s ===", t))
}
