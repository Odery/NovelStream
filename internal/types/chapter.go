package types

import "time"

// Chapter represents a single chapter of a novel, including its content.
type Chapter struct {
    Number      int       `json:"number"`      // Chapter number (for ordering)
    ImageLinks  []string  `json:"imageLinks"`  // URLs to images contained in the chapter, if any
    PublishDate time.Time `json:"publishDate"` // Publish date of the chapter
}