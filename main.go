package main

import (
	"fmt"
	"github.com/thexxiv/urbandictionary-go/urbandictionary"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	app(args)
}

func app(args []string) {
	if len(args) < 2 {
		fmt.Print("Usage: urban \"term\" [page number]\n")
		fmt.Print("           (Get list of definitions by page number)\n")
		fmt.Print("   or  urban -R\n")
		fmt.Print("           (Get list of random definitions)\n")
		fmt.Print("   or  urban -ID [id]\n")
		fmt.Print("           (Get the definition from given id)\n")
		fmt.Print("   or  urban -TIP \"term\"\n")
		fmt.Print("           (Get a tooltip based on given term)")
	} else {
		if args[1] == "-R" {
			random, err := urbandictionary.Random()
			if err != nil {
				fmt.Print("Error: Something went wrong")
			} else {
				writeFromSlice(random)
			}
		} else if args[1] == "-ID" {
			if len(args) < 3 {
				fmt.Print("Error: ID is missing")
			} else {
				num, err := strconv.Atoi(args[2])
				if err != nil {
					fmt.Print("Error: Something went wrong")
				} else {
					id, err := urbandictionary.DefinitionById(int64(num))
					if err != nil {
						fmt.Print("Error: Something went wrong")
					} else {
						write(id)
					}
				}
			}
		} else if args[1] == "-TIP" {
			if len(args) < 3 {
				fmt.Print("Error: Term is missing")
			} else {
				tip, err := urbandictionary.ToolTip(url.QueryEscape(args[2]))
				if err != nil {
					fmt.Print("Error: Something went wrong")
				} else {
					fmt.Print(strings.Replace(strings.Replace(tip,"<b>", "", -1),"</b>", "", -1))
				}
			}
		} else {
			if len(args) < 3 {
				fmt.Print("Error: Page number is missing")
			} else {
				num, err := strconv.Atoi(args[2])
				if err != nil {
					fmt.Print("Error: Something went wrong")
				} else {
					urban := urbandictionary.New(url.QueryEscape(args[1]), int32(num))
					data, err := urban.Data()
					if err != nil {
						fmt.Print("Error: Something went wrong")
					} else {
						writeFromSlice(data)
					}
				}
			}
		}
	}
}

func writeFromSlice(u []urbandictionary.Response) {
	for _, v := range u {
		fmt.Printf("Word: %s\n", v.Word)
		fmt.Printf("ID: %d\n", v.Defid)
		fmt.Printf("Definition: %s\n", v.Definition)
		fmt.Printf("Example: %s\n", v.Example)
		fmt.Printf("Likes: %d\n", v.ThumbsUp)
		fmt.Printf("Dislikes: %d\n", v.ThumbsDown)
		fmt.Printf("Date: %s\n", v.WrittenOn)
		fmt.Printf("Author: %s\n\n", v.Author)
	}
}

func write(r urbandictionary.Response) {
	fmt.Printf("Word: %s\n", r.Word)
	fmt.Printf("ID: %d\n", r.Defid)
	fmt.Printf("Definition: %s\n", r.Definition)
	fmt.Printf("Example: %s\n", r.Example)
	fmt.Printf("Likes: %d\n", r.ThumbsUp)
	fmt.Printf("Dislikes: %d\n", r.ThumbsDown)
	fmt.Printf("Date: %s\n", r.WrittenOn)
	fmt.Printf("Author: %s\n\n", r.Author)
}