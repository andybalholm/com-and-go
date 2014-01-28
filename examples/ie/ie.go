// This example program opens Internet Explorer and does a Google search
// for "golang". It is based on github.com/mattn/go-ole/example/ie.
//
// It leaks memory, since Release isn't called on anything.
package main

import (
	"code.google.com/p/com-and-go/v2"
	"log"
)

func main() {
	com.CoInitializeEx(nil, 0)

	ie, err := com.NewIDispatch("InternetExplorer.Application.1")
	if err != nil {
		log.Fatal(err)
	}

	ie.Call("Navigate", "http://www.google.com")
	ie.Put("Visible", true)

	for ie.Get("Busy") == true {
	}

	doc := ie.Get("document").(*com.IDispatch)
	ec := doc.Call("getElementsByName", "q").(*com.IDispatch)
	q := ec.Call("item", 0).(*com.IDispatch)
	q.Put("value", "golang")

	ec = doc.Call("getElementsByName", "gbqf").(*com.IDispatch)
	form := ec.Call("item", 0).(*com.IDispatch)

	form.Call("submit")
}
