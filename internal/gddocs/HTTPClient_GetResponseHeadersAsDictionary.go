/*
{
	"content-length": 12,
	"Content-Type": "application/json; charset=UTF-8",
}
*/

package main

func HTTPClient_GetResponseHeadersAsDictionary() {
	headers := map[string]string{
		"Content-Length": "12",
		"Content-Type":   "application/json; charset=UTF-8",
	}
	_ = headers
}
