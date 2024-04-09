package types

import "time"

// Novel represents the metadata and content of a web novel.
type Novel struct {
    Title       string    `json:"title"`       // Title of the novel
    Author      string    `json:"author"`      // Author of the novel
    Genre       []string  `json:"genre"`       // Genre(s) of the novel
    Summary     string    `json:"summary"`     // Brief summary of the novel
    CoverImage  string    `json:"coverImage"`  // URL to the cover image
    Chapters    []Chapter `json:"chapters"`    // List of chapters
    PublishDate time.Time `json:"publishDate"` // Original publish date of the novel
    UpdatedDate time.Time `json:"updatedDate"` // Date of the last update
    Status      string    `json:"status"`      // Status of the novel (e.g., Ongoing, Completed)
}