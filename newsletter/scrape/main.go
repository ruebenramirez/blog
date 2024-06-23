package main

import (
  "fmt"

  "bopbot/newsletter/scrape/lib"
  "github.com/yosssi/gohtml"
)

func main() {
  // where can we find the post title and content?
  titleSelector := "div#content div.container h2"
  postSelector := "div#content div.container"


  // // example usage
  // postUrl := "https://blog.ruebenramirez.com/posts/2023-10-06-gear-update/"
  // title, postHtml := scrapePost(postUrl, titleSelector, postSelector)
  // fmt.Printf(">>>blog post title: %s\n", gohtml.Format(title))
  // fmt.Printf(">>>blog post html: %s\n", gohtml.Format(postHtml))

  // example usage
  // postUrl := "https://blog.ruebenramirez.com/posts/2023-10-21-round-2-with-covid/"
  // title, postHtml := scrapePost(postUrl, titleSelector, postSelector)
  // fmt.Printf(">>>blog post title: %s\n", gohtml.Format(title))
  // fmt.Printf(">>>blog post html: %s\n", gohtml.Format(postHtml))

  // example usage
  postUrl := "https://blog.ruebenramirez.com/posts/2023-11-05-celebratin-60-pounds-down/"
  title, postHtml := lib.ScrapePost(postUrl, titleSelector, postSelector)
  fmt.Printf(">>>blog post title: %s\n", gohtml.Format(title))
  fmt.Printf(">>>blog post html: %s\n", gohtml.Format(postHtml))

  // example usage
  // postUrl := "https://blog.ruebenramirez.com/posts/2023-10-30-good-news-from-biopsy/"
  // title, postHtml := scrapePost(postUrl, titleSelector, postSelector)
  // fmt.Printf(">>>blog post title: %s\n", gohtml.Format(title))
  // fmt.Printf(">>>blog post html: %s\n", gohtml.Format(postHtml))
}


