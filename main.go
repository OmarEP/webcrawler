package main

import "fmt"

func main() {
	html := `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`
	url := "https://blog.boot.dev"

	fmt.Println(getURLsFromHTML(html, url))
}
