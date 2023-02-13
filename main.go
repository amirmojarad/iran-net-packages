package main

import (
	"internet_service/transport/server/http"
	"log"
)

//func main() {
//	c := colly.NewCollector()
//	c.OnError(func(response *colly.Response, err error) {
//		fmt.Println(err)
//	})
//	result := make([]string, 0)
//	c.OnHTML("div", func(element *colly.HTMLElement) {
//		if element.Attr("class") == "product-box" {
//			res := strings.Split(element.Text, "\n")
//			result = append(result, res[4])
//			result = append(result, res[5])
//			result = append(result, res[14])
//
//			fmt.Printf("%+v\n", res)
//
//		}
//	})
//	c.OnScraped(func(response *colly.Response) {
//		fmt.Printf("on Scrapped %+v\n", result)
//	})
//	err := c.Visit("https://irancell.ir/o/1001/%D8%A7%DB%8C%D8%B1%D8%A7%D9%86%D8%B3%D9%84-%D8%A8%D8%B3%D8%AA%D9%87-%D9%87%D8%A7%DB%8C-%D8%A7%DB%8C%D9%86%D8%AA%D8%B1%D9%86%D8%AA-%D9%87%D9%85%D8%B1%D8%A7%D9%87")
//	if err != nil {
//		return
//	}
//}

func main() {
	server := http.Server()
	err := server.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
