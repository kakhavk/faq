package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	// TODO: change fixed faqPath with dynamic path

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprint(w, "<!DOCTYPE html><html><head><meta charset=\"utf-8\" /><title></title></head><body></body></html>")
		renderTemplate(w, "index.html")
	})

	http.HandleFunc("/android", func(w http.ResponseWriter, r *http.Request) {
		// file, err := os.Open("faq/android.txt")
		// if err != nil {
		// 	fmt.Println("Error openinf file!!!")
		// }
		// defer file.Close()

		// const maxSz = 1024

		// // create buffer
		// b := make([]byte, maxSz)
		// fmt.Fprint(w, "<!DOCTYPE html><html><head><meta charset=\"utf-8\" /><title>Android</title></head><body>")
		// for {
		// 	// read content to buffer
		// 	readTotal, err := file.Read(b)
		// 	if err != nil {
		// 		if err != io.EOF {
		// 			fmt.Println(err)
		// 		}
		// 		break
		// 	}
		// 	fmt.Fprint(w, string(b[:readTotal])) // pring content from buffer
		// }
		// fmt.Fprint(w, "</body></html>")
		// faqDir, err := os.Getwd()
		faqDir := "include"
		faqPath := faqDir + "/android.txt"

		content, err := ioutil.ReadFile(faqPath)
		if err != nil {
			fmt.Println("Err")
		}

		fmt.Fprint(w, "<!DOCTYPE html><html><head><meta charset=\"utf-8\" /><title>Android</title></head><body>")
		fmt.Fprint(w, string(content))
		fmt.Fprint(w, "</body></html>")

	})

	http.HandleFunc("/bootstrap", func(w http.ResponseWriter, r *http.Request) {
		faqDir := "include"
		faqPath := faqDir + "/bootstrap.txt"

		file, err := os.Open(faqPath)
		if err != nil {
			fmt.Println("Error openinf file!!!")
		}
		defer file.Close()

		const maxSz = 1024

		// create buffer
		b := make([]byte, maxSz)
		fmt.Fprint(w, "<!DOCTYPE html><html><head><meta charset=\"utf-8\" /><title>Bootstrap</title></head><body>")
		for {
			// read content to buffer
			readTotal, err := file.Read(b)
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			}
			fmt.Fprint(w, string(b[:readTotal])) // pring content from buffer
		}
		fmt.Fprint(w, "</body></html>")
	})

	http.HandleFunc("/linux", func(w http.ResponseWriter, r *http.Request) {
		faqDir := "include"
		faqPath := faqDir + "/linux.txt"

		file, err := os.Open(faqPath)
		if err != nil {
			fmt.Println("Error openinf file!!!")
		}
		defer file.Close()

		const maxSz = 1024

		// create buffer
		b := make([]byte, maxSz)
		fmt.Fprint(w, "<!DOCTYPE html><html><head><meta charset=\"utf-8\" /><title>Linux</title></head><body>")
		for {
			// read content to buffer
			readTotal, err := file.Read(b)
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			}
			fmt.Fprint(w, string(b[:readTotal])) // pring content from buffer
		}
		fmt.Fprint(w, "</body></html>")
	})

	http.HandleFunc("/php", func(w http.ResponseWriter, r *http.Request) {
		faqDir := "include"
		faqPath := faqDir + "/php.txt"

		file, err := os.Open(faqPath)
		if err != nil {
			fmt.Println("Error openinf file!!!")
		}
		defer file.Close()

		const maxSz = 1024

		// create buffer
		b := make([]byte, maxSz)
		fmt.Fprint(w, "<!DOCTYPE html><html><head><meta charset=\"utf-8\" /><title>PHP</title></head><body>")
		for {
			// read content to buffer
			readTotal, err := file.Read(b)
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			}
			fmt.Fprint(w, string(b[:readTotal])) // pring content from buffer
		}
		fmt.Fprint(w, "</body></html>")
	})

	http.HandleFunc("/windows", func(w http.ResponseWriter, r *http.Request) {
		faqDir := "include"
		faqPath := faqDir + "/windows.txt"

		file, err := os.Open(faqPath)
		if err != nil {
			fmt.Println("Error openinf file!!!")
		}
		defer file.Close()

		const maxSz = 1024

		// create buffer
		b := make([]byte, maxSz)
		fmt.Fprint(w, "<!DOCTYPE html><html><head><meta charset=\"utf-8\" /><title>Windows</title></head><body>")
		for {
			// read content to buffer
			readTotal, err := file.Read(b)
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			}
			fmt.Fprint(w, string(b[:readTotal])) // pring content from buffer
		}
		fmt.Fprint(w, "</body></html>")
	})

	http.ListenAndServe(":9991", nil)

}

func renderTemplate(w http.ResponseWriter, page string) {
	var htmlData string
	parsedTemplate, _ := template.New("").ParseFiles("templates/"+page, "templates/base.html")

	err := parsedTemplate.ExecuteTemplate(w, "base", htmlData)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}

}
