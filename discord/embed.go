package discord

// Wrapper to make embeds less of a pain in the ass.
// Stolen from https://gist.github.com/Necroforger/8b0b70b1a69fa7828b8ad6387ebb3835

import (
	discord "github.com/chremoas/discord-gateway/proto"
)

//Embed ...
type Embed struct {
	*discord.MessageEmbed
}

// Constants for message embed character limits
const (
	EmbedLimitTitle       = 256
	EmbedLimitDescription = 2048
	EmbedLimitFieldValue  = 1024
	EmbedLimitFieldName   = 256
	EmbedLimitField       = 25
	EmbedLimitFooter      = 2048
	EmbedLimit            = 4000
)

//NewEmbed returns a new embed object
func NewEmbed() *Embed {
	return &Embed{&discord.MessageEmbed{}}
}

//SetTitle ...
func (e *Embed) SetTitle(name string) *Embed {
	e.Title = name
	return e
}

//SetDescription [desc]
func (e *Embed) SetDescription(description string) *Embed {
	if len(description) > EmbedLimitDescription {
		description = description[:EmbedLimitDescription]
	}
	e.Description = description
	return e
}

//AddField [name] [value]
func (e *Embed) AddField(name, value string) *Embed {
	if len(value) > EmbedLimitFieldValue {
		value = value[:EmbedLimitFieldValue]
	}

	if len(name) > EmbedLimitFieldName {
		name = name[:EmbedLimitFieldName]
	}

	e.Fields = append(e.Fields, &discord.MessageEmbedField{
		Name:  name,
		Value: value,
	})

	return e

}

//SetFooter [Text] [iconURL]
func (e *Embed) SetFooter(footer *discord.MessageEmbedFooter) *Embed {
	e.Footer = footer
	return e
}

//SetImage ...
func (e *Embed) SetImage(image *discord.MessageEmbedImage) *Embed {
	e.Image = image
	return e
}

//SetThumbnail ...
func (e *Embed) SetThumbnail(thumbnail *discord.MessageEmbedThumbnail) *Embed {
	e.Thumbnail = thumbnail
	return e
}

//SetAuthor ...
func (e *Embed) SetAuthor(author *discord.MessageEmbedAuthor) *Embed {
	e.Author = author
	return e
}

//SetURL ...
func (e *Embed) SetURL(URL string) *Embed {
	e.URL = URL
	return e
}

//SetColor ...
func (e *Embed) SetColor(clr int64) *Embed {
	e.Color = clr
	return e
}

//SetTimestamp ...
func (e *Embed) SetTimestamp(timestamp string) *Embed {
	e.Timestamp = timestamp
	return e
}

// InlineAllFields sets all fields in the embed to be inline
func (e *Embed) InlineAllFields() *Embed {
	for _, v := range e.Fields {
		v.Inline = true
	}
	return e
}

// Truncate truncates any embed value over the character limit.
func (e *Embed) Truncate() *Embed {
	e.TruncateDescription()
	e.TruncateFields()
	e.TruncateFooter()
	e.TruncateTitle()
	return e
}

// TruncateFields truncates fields that are too long
func (e *Embed) TruncateFields() *Embed {
	if len(e.Fields) > 25 {
		e.Fields = e.Fields[:EmbedLimitField]
	}

	for _, v := range e.Fields {

		if len(v.Name) > EmbedLimitFieldName {
			v.Name = v.Name[:EmbedLimitFieldName]
		}

		if len(v.Value) > EmbedLimitFieldValue {
			v.Value = v.Value[:EmbedLimitFieldValue]
		}

	}
	return e
}

// TruncateDescription ...
func (e *Embed) TruncateDescription() *Embed {
	if len(e.Description) > EmbedLimitDescription {
		e.Description = e.Description[:EmbedLimitDescription]
	}
	return e
}

// TruncateTitle ...
func (e *Embed) TruncateTitle() *Embed {
	if len(e.Title) > EmbedLimitTitle {
		e.Title = e.Title[:EmbedLimitTitle]
	}
	return e
}

// TruncateFooter ...
func (e *Embed) TruncateFooter() *Embed {
	if e.Footer != nil && len(e.Footer.Text) > EmbedLimitFooter {
		e.Footer.Text = e.Footer.Text[:EmbedLimitFooter]
	}
	return e
}
