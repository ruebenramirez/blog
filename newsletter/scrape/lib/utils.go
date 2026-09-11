package lib

import(
  "log"
  "net/url"
  "strings"
)

// return the root url domain from provided postUrl
func getRootUrl(postUrl string) string {
  url, err := url.Parse(postUrl)
  if err != nil {
    log.Fatal(err)
  }
  url.Path = "/"

  return url.String()
}

// replace root src and hrefs with FQDN refs
func replaceRootRefs(html string, postUrl string) string {
  rootReplacementUrl := getRootUrl(postUrl)
  html = strings.Replace(html, "=\"/", "=\"" + rootReplacementUrl, -1)

  return html
}
