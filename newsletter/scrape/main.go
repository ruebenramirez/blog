package main

import (
  "fmt"
  "log"
  "net/http"
  "net/url"
  "strings"

  "github.com/PuerkitoBio/goquery"
  "github.com/yosssi/gohtml"
)

var DEBUG bool = false

func GrabPage(url string) *goquery.Document {
  // Request the HTML page.
  res, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
    log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
  }

  // Load the HTML document
  doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {
    log.Fatal(err)
  }

  if DEBUG {
    //print out the page
    fmt.Printf(">>>status code: %d\n", res.StatusCode)
    pageHtml, err := doc.Html()
    if err != nil {
      log.Fatal(err)
    }
    fmt.Printf(">>>page html: %s\n", gohtml.Format(pageHtml))
  }

  return doc
}

func GrabPostTitleFromPage(doc *goquery.Document, titleSelector string) string {
  blogTitle := doc.Find(titleSelector).Text()
  if DEBUG {
    fmt.Printf(">>>post title: %s\n", blogTitle)
  }

  return blogTitle
}

// return the root url domain from provided postUrl
func GetRootUrl(postUrl string) string {
  url, err := url.Parse(postUrl)
  if err != nil {
    log.Fatal(err)
  }
  url.Path = "/"

  if DEBUG {
    fmt.Printf(">>>root url: %s\n", url.String())
  }

  return url.String()
}

// replace root src and hrefs with FQDN refs
func ReplaceRootRefs(html string, postUrl string) string {
  rootReplacementUrl := GetRootUrl(postUrl)
  html = strings.Replace(html, "=\"/", "=\"" + rootReplacementUrl, -1)

  return html
}

func GrabPostFromPage(doc *goquery.Document, postUrl string, postSelector string) string {
  // grab post html content
  blogContent, err := doc.Find(postSelector).Html()
  if err != nil {
    log.Fatal(err)
  }

  // replace `/` resource refs with FQDN
  blogContent = ReplaceRootRefs(blogContent, postUrl)

  if DEBUG {
    fmt.Printf(">>>blog post: %s\n", gohtml.Format(blogContent))
  }

  return blogContent
}

func ScrapePost(postUrl string, titleSelector string, postSelector string) (string, string) {
  doc := GrabPage(postUrl)

  title := GrabPostTitleFromPage(doc, titleSelector)
  postHtml := GrabPostFromPage(doc, postUrl, postSelector)
  return title, postHtml
}

func main() {
  // where can we find the post title and content?
  titleSelector := "div#content div.container h2"
  postSelector := "div#content div.container"


  // // example usage
  // postUrl := "https://blog.ruebenramirez.com/posts/2023-10-06-gear-update/"
  // title, postHtml := ScrapePost(postUrl, titleSelector, postSelector)
  // fmt.Printf(">>>blog post title: %s\n", gohtml.Format(title))
  // fmt.Printf(">>>blog post html: %s\n", gohtml.Format(postHtml))

  // example usage
  // postUrl := "https://blog.ruebenramirez.com/posts/2023-10-21-round-2-with-covid/"
  // title, postHtml := ScrapePost(postUrl, titleSelector, postSelector)
  // fmt.Printf(">>>blog post title: %s\n", gohtml.Format(title))
  // fmt.Printf(">>>blog post html: %s\n", gohtml.Format(postHtml))

  // example usage
  postUrl := "https://blog.ruebenramirez.com/posts/2023-11-05-celebratin-60-pounds-down/"
  title, postHtml := ScrapePost(postUrl, titleSelector, postSelector)
  fmt.Printf(">>>blog post title: %s\n", gohtml.Format(title))
  fmt.Printf(">>>blog post html: %s\n", gohtml.Format(postHtml))

  // example usage
  // postUrl := "https://blog.ruebenramirez.com/posts/2023-10-30-good-news-from-biopsy/"
  // title, postHtml := ScrapePost(postUrl, titleSelector, postSelector)
  // fmt.Printf(">>>blog post title: %s\n", gohtml.Format(title))
  // fmt.Printf(">>>blog post html: %s\n", gohtml.Format(postHtml))
}


