// Package convert contains the machinery to read a WordPress export file
// and return a collection of data structures containing the exported data.
package convert

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// WPExport contains the exported data.
type WPExport struct {
	Channel Channel `xml:"channel"`
}

// Channel contains the blog information.
type Channel struct {
	Title       string     `xml:"title"`
	Link        string     `xml:"link"`
	Description string     `xml:"description"`
	PubDate     string     `xml:"pubDate"`
	Language    string     `xml:"language"`
	Version     string     `xml:"wxr_version"`
	SiteURL     string     `xml:"base_site_url"`
	BlogURL     string     `xml:"base_blog_url"`
	Authors     []Author   `xml:"author"`
	Categories  []Category `xml:"category"`
	Tags        []Tag      `xml:"tag"`
	Generator   string     `xml:"generator"`
	Items       []Item     `xml:"item"`
}

// Author contains the author information.
type Author struct {
	ID          string `xml:"author_id"`
	Login       string `xml:"author_login"`
	Email       string `xml:"author_email"`
	DisplayName string `xml:"author_display_name"`
	FirstName   string `xml:"author_first_name"`
	LastName    string `xml:"author_last_name"`
}

// Category contains the available categories information.
type Category struct {
	ID       string `xml:"term_id"`
	NiceName string `xml:"category_nicename"`
	Parent   string `xml:"category_parent"`
	Name     string `xml:"cat_name"`
}

// Tag contins the avilable information for a given tag.
type Tag struct {
	ID   string `xml:"term_id"`
	Slug string `xml:"tag_slug"`
	Name string `xml:"tag_name"`
}

// Item contains the blog entries.
type Item struct {
	ID                string           `xml:"post_id"`
	Title             string           `xml:"title"`
	Link              string           `xml:"link"`
	PublicationDate   string           `xml:"pubDate"`
	Creator           string           `xml:"creator"`
	GUID              string           `xml:"guid"`
	Description       string           `xml:"description"`
	PostDate          string           `xml:"post_date"`
	PostDateGMT       string           `xml:"post_date_gmt"`
	CommentStatus     string           `xml:"comment_status"`
	PingStatus        string           `xml:"ping_status"`
	Status            string           `xml:"status"`
	MenuOrder         string           `xml:"menu_order"`
	PostName          string           `xml:"post_name"`
	PostParentID      string           `xml:"post_parent"`
	PostType          string           `xml:"post_type"`
	PostPassword      string           `xml:"post_password"`
	IsSticky          string           `xml:"is_sticky"`
	AttachmentURL     string           `xml:"attachment_url"`
	CategoriesAndTags []CategoryAndTag `xml:"category"`
	PostMetas         []PostMeta       `xml:"postmeta"`
}

// CategoryAndTag contains post categories and tags.
type CategoryAndTag struct {
	Value    string `xml:",chardata"`
	Domain   string `xml:"domain,attr"`
	NiceName string `xml:"nicename,attr"`
}

// Tags returns the tags avaialble.
func (i Item) Tags() string {
	var tags []string
	for _, t := range i.CategoriesAndTags {
		if t.Domain == "post_tag" {
			tags = append(tags, fmt.Sprintf("\"%s\"", strings.TrimSpace(t.Value)))
		}
	}
	if tags == nil {
		return ""
	}
	return fmt.Sprintf(" %s ", strings.Join(tags, ", "))
}

// Categories returns the tags avaialble.
func (i Item) Categories(prelude, pre, sep, post, ending string) string {
	var cats []string
	for _, t := range i.CategoriesAndTags {
		if t.Domain == "category" {
			cats = append(cats, fmt.Sprintf("\"%s\"", strings.TrimSpace(t.Value)))
		}
	}
	if cats == nil {
		return ""
	}
	return fmt.Sprintf("%s%s%s%s",
		prelude,
		pre,
		strings.Join(cats, strings.Join([]string{sep, post, pre}, "")),
		ending)
}

// PostMeta contains post meta info in form of key value pairs.
type PostMeta struct {
	Key   string `xml:"meta_key"`
	Value string `xml:"meta_value"`
}

// ToMarkDown converts a XML export string into the corresponding WPExport type.
func ToMarkDown(xmlBytes []byte) (*WPExport, error) {
	wpe := &WPExport{}
	if err := xml.Unmarshal(xmlBytes, wpe); err != nil {
		return nil, err
	}
	return wpe, nil
}
