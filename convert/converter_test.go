package convert

import (
	"bytes"
	"reflect"
	"testing"
)

func TestXMLConversion(t *testing.T) {
	xmlInput := bytes.NewBufferString(`
    <?xml version="1.0" encoding="UTF-8" ?>
    <rss version="2.0"
    	xmlns:excerpt="http://wordpress.org/export/1.2/excerpt/"
    	xmlns:content="http://purl.org/rss/1.0/modules/content/"
    	xmlns:wfw="http://wellformedweb.org/CommentAPI/"
    	xmlns:dc="http://purl.org/dc/elements/1.1/"
    	xmlns:wp="http://wordpress.org/export/1.2/"
    >

    <channel>
    	<title>Some blog</title>
    	<link>http://www.some.url</link>
    	<description>Some description.</description>
    	<pubDate>Tue, 19 Jan 2016 04:34:13 +0000</pubDate>
    	<language>en-US</language>
    	<wp:wxr_version>1.2</wp:wxr_version>
    	<wp:base_site_url>http://www.some.url</wp:base_site_url>
    	<wp:base_blog_url>http://www.some.url</wp:base_blog_url>

    	<wp:author><wp:author_id>1</wp:author_id><wp:author_login>admin</wp:author_login><wp:author_email>admin@some.url</wp:author_email><wp:author_display_name><![CDATA[Some]]></wp:author_display_name><wp:author_first_name><![CDATA[Some]]></wp:author_first_name><wp:author_last_name><![CDATA[Author]]></wp:author_last_name></wp:author>

    	<wp:category><wp:term_id>1</wp:term_id><wp:category_nicename>activities</wp:category_nicename><wp:category_parent>parent</wp:category_parent><wp:cat_name><![CDATA[Activities]]></wp:cat_name></wp:category>
    	<wp:tag><wp:term_id>1</wp:term_id><wp:tag_slug>1</wp:tag_slug><wp:tag_name><![CDATA[tag]]></wp:tag_name></wp:tag>

    	<generator>http://wordpress.org/?v=4.3.2</generator>

    	<item>
    		<title>Visualizing content from metadata stores</title>
    		<link>http://www.some.url/path</link>
    		<pubDate>Sun, 15 Apr 2007 13:54:06 +0000</pubDate>
    		<dc:creator><![CDATA[admin]]></dc:creator>
    		<guid isPermaLink="false">http://www.some/url/path.pdf</guid>
    		<description>Some description</description>
    		<content:encoded><![CDATA[Visualizing content from metadata stores]]></content:encoded>
    		<excerpt:encoded><![CDATA[]]></excerpt:encoded>
    		<wp:post_id>8</wp:post_id>
    		<wp:post_date>2007-04-15 08:54:06</wp:post_date>
    		<wp:post_date_gmt>2007-04-15 13:54:06</wp:post_date_gmt>
    		<wp:comment_status>open</wp:comment_status>
    		<wp:ping_status>open</wp:ping_status>
    		<wp:post_name>visualizing-content-from-metadata-stores</wp:post_name>
    		<wp:status>inherit</wp:status>
    		<wp:post_parent>7</wp:post_parent>
    		<wp:menu_order>0</wp:menu_order>
    		<wp:post_type>attachment</wp:post_type>
    		<wp:post_password>1234</wp:post_password>
    		<wp:is_sticky>0</wp:is_sticky>
    		<wp:attachment_url>http://www.some.url/attachment.pdf</wp:attachment_url>
				<category domain="post_tag" nicename="tag"><![CDATA[Tag]]></category>
				<wp:postmeta>
    			<wp:meta_key>_wp_attached_file</wp:meta_key>
    			<wp:meta_value><![CDATA[/some/path/public_html/some/wp-content/uploads/2007/04/2007-04-13-file.pdf]]></wp:meta_value>
    		</wp:postmeta>
    	</item>
    </channel>
    </rss>
  `)
	want := &WPExport{
		Channel: Channel{
			Title:       "Some blog",
			Link:        "http://www.some.url",
			Description: "Some description.",
			PubDate:     "Tue, 19 Jan 2016 04:34:13 +0000",
			Language:    "en-US",
			Version:     "1.2",
			SiteURL:     "http://www.some.url",
			BlogURL:     "http://www.some.url",
			Authors: []Author{
				{
					ID:          "1",
					Login:       "admin",
					Email:       "admin@some.url",
					DisplayName: "Some",
					FirstName:   "Some",
					LastName:    "Author",
				},
			},
			Categories: []Category{
				{
					ID:       "1",
					NiceName: "activities",
					Parent:   "parent",
					Name:     "Activities",
				},
			},
			Tags: []Tag{
				{
					ID:   "1",
					Slug: "1",
					Name: "tag",
				},
			},
			Generator: "http://wordpress.org/?v=4.3.2",
			Items: []Item{
				{
					ID:              "8",
					Title:           "Visualizing content from metadata stores",
					Link:            "http://www.some.url/path",
					PublicationDate: "Sun, 15 Apr 2007 13:54:06 +0000",
					Creator:         "admin",
					GUID:            "http://www.some/url/path.pdf",
					Description:     "Some description",
					PostDate:        "2007-04-15 08:54:06",
					PostDateGMT:     "2007-04-15 13:54:06",
					CommentStatus:   "open",
					PingStatus:      "open",
					Status:          "inherit",
					MenuOrder:       "0",
					PostName:        "visualizing-content-from-metadata-stores",
					PostParentID:    "7",
					PostType:        "attachment",
					PostPassword:    "1234",
					IsSticky:        "0",
					AttachmentURL:   "http://www.some.url/attachment.pdf",
					CategoriesAndTags: []CategoryAndTag{
						{
							Value:    "Tag",
							Domain:   "post_tag",
							NiceName: "tag",
						},
					},
					PostMetas: []PostMeta{
						{
							Key:   "_wp_attached_file",
							Value: "/some/path/public_html/some/wp-content/uploads/2007/04/2007-04-13-file.pdf",
						},
					},
				},
			},
		},
	}
	got, err := ToMarkDown(xmlInput.Bytes())
	if err != nil {
		t.Fatalf("Failed to parse test export with error %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("convert.ToMarkDown failed to return the right data; got\n%#v\nwant\n%#v\n", got, want)
	}
}
