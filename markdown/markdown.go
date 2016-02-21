// Package markdown contains funtionality to take WordPress extraded data into
// MarkDown files.
package markdown

import (
	"bytes"
	"errors"
	"html/template"

	"github.com/xllora/WPtoMD/convert"
)

// FrontMatterFormat list the possible front matter formats.
type FrontMatterFormat uint8

const (
	// JSON formated front matter.
	JSON FrontMatterFormat = iota
	// TOML formated front matter.
	TOML
	// YAML formated front matter.
	YAML
)

var (
	frontMatterJSON *template.Template
	frontMatterTOML *template.Template
	frontMatterYAML *template.Template

	jsonTemplate = ``

	tomlTemplate = `+++
title = "{{.Title}}"
description = "{{.Description}}"
tag = [{{.Tags}}]
date = "{{.PublicationDate}}"
categories = []
slug = "{{.PostName}}"
+++`

	yamlTemplate = ``
)

func init() {
	frontMatterJSON = template.New("json")
	if _, err := frontMatterJSON.Parse(jsonTemplate); err != nil {
		panic(err)
	}
	frontMatterTOML = template.New("TOML")
	if _, err := frontMatterTOML.Parse(tomlTemplate); err != nil {
		panic(err)
	}
	frontMatterYAML = template.New("YAML")
	if _, err := frontMatterYAML.Parse(yamlTemplate); err != nil {
		panic(err)
	}
}

// ToFrontMatter adds the front matter to the provided buffer.
func ToFrontMatter(buff *bytes.Buffer, tfm FrontMatterFormat, wpe *convert.Item) error {
	var tmpl *template.Template
	switch tfm {
	case JSON:
		tmpl = frontMatterJSON
	case TOML:
		tmpl = frontMatterTOML
	case YAML:
		tmpl = frontMatterYAML
	default:
		return errors.New("unknown front matter type")
	}
	return tmpl.Execute(buff, wpe)
}
