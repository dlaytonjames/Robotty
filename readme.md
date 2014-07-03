Robotty is a robots.txt parser for Go language
--------------------------------------------------

Usage:

import the package

```
import "robot"
```

```go
// parse a robots.txt string
decision := robot.FromString("User-agent: * \n Disallow: /css/ \n Disallow: /cgi-bin/")
decision.IsAllowed("http://site.com/css", "*")     // returns false

// or parse a robots.txt http.Response
resp, err := http.Get("http://google.com.ng/robots.txt")
decision, _ := robot.FromResponse(resp)
decision.IsAllowed("http://google.com.ng/bleh/bleh", "*")
```


- This library follows the exact matching pattern on this page https://developers.google.com/webmasters/control-crawl-index/docs/robots_txt
- It will first parse the first group that matches the user agent specified in Decision.IsAllowed and only fallback to the default group "*" if the specified user agent is not found
- Only the first matching directive is accepted.

