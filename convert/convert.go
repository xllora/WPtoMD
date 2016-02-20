// Package convert contains the machinery to read a WordPress export file
// and return a collection of data structures containing the exported data.
package convert

import "encoding/xml"

// WPExport contains the exported data.
type WPExport struct {
	Channel Channel `xml:"channel"`
}

// Channel contains the blog information.
type Channel struct {
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	PubDate     string   `xml:"pubDate"`
	Language    string   `xml:"language"`
	Version     string   `xml:"wxr_version"`
	SiteURL     string   `xml:"base_site_url"`
	BlogURL     string   `xml:"base_blog_url"`
	Authors     []Author `xmls:"author"`
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

// ToMarkDown converts a XML export string into the corresponding WPExport type.
func ToMarkDown(xmlBytes []byte) (*WPExport, error) {
	wpe := &WPExport{}
	if err := xml.Unmarshal(xmlBytes, wpe); err != nil {
		return nil, err
	}
	return wpe, nil
}
