package main

import (
	"fmt"
	"retriever/mock"
	"retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

// RetrieverPoster
// 使用者设置接口组合
type RetrieverPoster interface {
	Retriever
	Poster
}

const (
	url = "https://www.imooc.com"
)

func download(r Retriever) string {
	return r.Get(url)
}

func post(p Poster) {
	p.Post(
		url,
		map[string]string{
			"name":   "asura",
			"course": "golang",
		})
}

func session(rp RetrieverPoster) string {
	rp.Post(url, map[string]string{
		"Contents": "another faked imooc.com",
	})
	return rp.Get(url)
}

func inspect(r Retriever) {
	fmt.Printf("1-----------%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Printf("2-----------%T %v\n", r, r)
		fmt.Println("Contents: ", v)
	case *real.Retriever:
		fmt.Printf("3-----------%T %v\n", r, r)
		fmt.Println("UserAgent: ", v.UserAgent)
	}
	fmt.Println("-----------------------")
}

func main() {
	var r Retriever
	// 指针接收者
	retriever := &mock.Retriever{
		Contents: "mock data",
	}
	r = retriever
	inspect(r)
	r = &mock.Retriever{
		Contents: "mock data",
	}
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		Timeout:   time.Minute,
	}
	inspect(r)

	// Type assertion 类型断言
	if realRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(realRetriever.Contents)
	} else {
		fmt.Println("r is not mock Retriever")
	}
	fmt.Println("-----------------------")
	fmt.Println("Try a session")
	fmt.Println(session(retriever))

}
