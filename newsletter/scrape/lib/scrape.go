package lib

import (
  "log"
  "net/http"

  "github.com/PuerkitoBio/goquery"
)

func grabPage(url string) *goquery.Document {
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

  return doc
}

func grabPostFromPage(doc *goquery.Document, postUrl string, postSelector string) string {
  // grab post html content
  blogContent, err := doc.Find(postSelector).Html()
  if err != nil {
    log.Fatal(err)
  }

  // replace `/` resource refs with FQDN
  blogContent = replaceRootRefs(blogContent, postUrl)

  return blogContent
}

func grabPostTitleFromPage(doc *goquery.Document, titleSelector string) string {
  blogTitle := doc.Find(titleSelector).Text()

  return blogTitle
}

func ScrapePost(postUrl string, titleSelector string, postSelector string) (string, string) {
  doc := grabPage(postUrl)

  title := grabPostTitleFromPage(doc, titleSelector)
  postHtml := grabPostFromPage(doc, postUrl, postSelector)
  return title, postHtml
}
