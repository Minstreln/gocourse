package intermediate

import (
	"fmt"
	"net/url"
)

func main() {

	// [schema://[userInfo@]host[:port][/path][?query][#fragment]

	rawUrl := "https://example.com:8080/path?query=param#fragment"

	parsedUrl, err := url.Parse(rawUrl)

	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Port:", parsedUrl.Port())
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("Raw Query:", parsedUrl.RawQuery)
	fmt.Println("fragment:", parsedUrl.Fragment)

	rawUrl1 := "https://example.com/path?name=john&age=30"

	parsedUrl1, err := url.Parse(rawUrl1)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
	}

	queryParams := parsedUrl1.Query()
	fmt.Println("Query Parameters:", queryParams)
	fmt.Println("Name:", queryParams.Get("name"))
	fmt.Println("Age:", queryParams.Get("age"))

	// building url
	baseUrl := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}

	query := baseUrl.Query()
	query.Set("name", "john")
	query.Set("age", "30")
	baseUrl.RawQuery = query.Encode()

	fmt.Println("Build url:", baseUrl.String())

	values := url.Values{}

	// add key value pairs to the value objects
	values.Add("name", "john")
	values.Add("age", "30")
	values.Add("city", "newyork")
	values.Add("country", "USA")

	// Encode the values to a URL-encoded query string
	encodedQuery := values.Encode()

	fmt.Println("Encoded Query String:", encodedQuery)

	// build a url
	baseUrl1 := "https://example.com/search"
	fullUrl := baseUrl1 + "?" + encodedQuery
	fmt.Println("Full URL:", fullUrl)

}
