package katana

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

type HTMLElement struct {
	Name       string
	Text       string
	DOM        *goquery.Selection
	attributes []html.Attribute
}

func NewHTMLElementFromSelectionNode(s *goquery.Selection, n *html.Node) *HTMLElement {
	return &HTMLElement{
		Name:       n.Data,
		Text:       goquery.NewDocumentFromNode(n).Text(),
		DOM:        s,
		attributes: n.Attr,
	}
}

func (h *HTMLElement) Attr(k string) string {
	for _, a := range h.attributes {
		if a.Key == k {
			return a.Val
		}
	}
	return ""
}

func (h *HTMLElement) ChildText(goquerySelector string) string {
	return strings.TrimSpace(h.DOM.Find(goquerySelector).Text())
}

func (h *HTMLElement) ChildTexts(goquerySelector string) []string {
	var res []string
	h.DOM.Find(goquerySelector).Each(func(_ int, s *goquery.Selection) {

		res = append(res, strings.TrimSpace(s.Text()))
	})
	return res
}

func (h *HTMLElement) ChildAttr(goquerySelector, attrName string) string {
	if attr, ok := h.DOM.Find(goquerySelector).Attr(attrName); ok {
		return strings.TrimSpace(attr)
	}
	return ""
}

func (h *HTMLElement) ChildAttrs(goquerySelector, attrName string) []string {
	var res []string
	h.DOM.Find(goquerySelector).Each(func(_ int, s *goquery.Selection) {
		if attr, ok := s.Attr(attrName); ok {
			res = append(res, strings.TrimSpace(attr))
		}
	})
	return res
}
