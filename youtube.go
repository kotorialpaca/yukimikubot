package yukimikubot

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"
)

type VideoInfo struct {
	// The video ID
	ID string `json:"id"`
	// The video title
	Title string `json:"title"`
	// The video description
	Description string `json:"description"`
	// The date the video was published
	DatePublished time.Time `json:"datePublished"`
	// Formats the video is available in
	Formats FormatList `json:"formats"`
	// List of keywords associated with the video
	Keywords []string `json:"keywords"`
	// Author of the video
	Author string `json:"author"`
	// Duration of the video
	Duration time.Duration

	htmlPlayerFile string
}

func downloadVideo(format Format, dest io.Writer) error {
	u, err := info.GetDownloadURL(format)
	if err != nil {
		return err
	}
	resp, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("Invalid status code: %d", resp.StatusCode)
	}
	_, err = io.Copy(dest, resp.Body)
	return err
}

func CheckTrackLink(t string) bool {
	if a, _ := regexp.MatchString(`https?:\/\/www.youtube.com\/watch\?v=(?P<id>[\w-]+)(?P<timestamp>\&t=\d*m?\d*s?)?`, t); a == false {
		if b, _ := regexp.MatchString(`https?:\/\/youtube.com\/watch\?v=(?P<id>[\w-]+)(?P<timestamp>\&t=\d*m?\d*s?)?`, t); b == false {
			if c, _ := regexp.MatchString(`https?:\/\/youtu.be\/(?P<id>[\w-]+)(?P<timestamp>\?t=\d*m?\d*s?)?`, t); c == false {
				if d, _ := regexp.MatchString(`https?:\/\/youtube.com\/v\/(?P<id>[\w-]+)(?P<timestamp>\?t=\d*m?\d*s?)?`, t); d == false {
					if e, _ := regexp.MatchString(`https?:\/\/www.youtube.com\/v\/(?P<id>[\w-]+)(?P<timestamp>\?t=\d*m?\d*s?)?`, t); e == false {
						return false
					}
				}
			}
		}
	}
	return true
}

func CheckPlaylistLink(t string) bool {
	m, err := regexp.MatchString(`https?:\/\/www\.youtube\.com\/playlist\?list=(?P<id>[\w-]+)`, t)
	if err != nil {
		return false
	}
	return m
}
