package youtube

import "regexp"

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
