// Package model - Content managed by Project Forge, see [projectforge.md] for details.
package model

import (
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

var LinkFieldDescs = util.FieldDescs{
	{Key: "title", Title: "Title", Description: "The title of the link"},
	{Key: "icon", Title: "Icon", Description: "The icon of this link"},
	{Key: "url", Title: "URL", Description: "The href of this link"},
}

type Link struct {
	Title string   `json:"title,omitempty"`
	URL   string   `json:"url"`
	Icon  string   `json:"icon,omitempty"`
	Tags  []string `json:"tags,omitempty"`
}

type Links []*Link

func (l Links) WithTags(includeEmpty bool, tags ...string) Links {
	return lo.Filter(l, func(x *Link, _ int) bool {
		return (includeEmpty && len(x.Tags) == 0) || (len(lo.Intersect(tags, x.Tags)) > 0)
	})
}