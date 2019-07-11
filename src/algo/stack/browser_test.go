package stack

import "testing"

func TestBrowser(t *testing.T) {
	browser := NewBrowser()
	browser.Open("http://ltinyho.top")
	browser.Open("http://www.google.com")
	browser.Open("http://www.baidu.com")
	browser.Back()
	browser.Back()
	browser.Back()
	browser.Forward()
}
