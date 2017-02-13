package services

import (
	"errors"
	"io"
	"regexp"

	"fmt"

	"os"

	"github.com/rylio/ytdl"
)

//Play function will be called from handlers as an all in one function.
func Play(t string, dest io.Writer, cache bool) error {
	//If caching is enabled

	if cache {
		//Check if it is a valid link
		if checkTrackLink(t) {
			//Grab the video info from youtube
			vi, err := ytdl.GetVideoInfo(t)
			if err != nil {
				return err
			}

			//Search by Video ID, and if it does not exist, download
			fm := vi.Formats.Best(ytdl.FormatResolutionKey)[0]
			ext := fm.Extension
			fmt.Println(ext)
			/*
				f, err := ioutil.ReadFile("cache/" + vi.ID + "." + ext)
				if os.IsNotExist(err) {
					//Get extension
					//Download video
					err = downloadYTCache(vi, fm, ext)

				} else {

				}
			*/

		} else if checkPlaylistLink(t) {

		} else {
			return errors.New("not a valid youtube link")
		}
	}
	return nil

}

func DownloadYTCache(vi *ytdl.VideoInfo, fm ytdl.Format, e string) error {

	f, err := os.Create("cache/" + vi.ID + "." + e)
	defer f.Close()
	if err != nil {
		return err
	}
	err = vi.Download(fm, f)
	if err != nil {
		return err
	}
	f.Sync()

	return nil
}

func checkTrackLink(t string) bool {
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

func checkPlaylistLink(t string) bool {
	m, err := regexp.MatchString(`https?:\/\/www\.youtube\.com\/playlist\?list=(?P<id>[\w-]+)`, t)
	if err != nil {
		return false
	}
	return m
}
