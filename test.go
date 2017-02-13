package main

import (
	"fmt"

	"github.com/kotorialpaca/yukimikubot/services"
	"github.com/rylio/ytdl"
)

func main() {
	vi, err := ytdl.GetVideoInfo("https://www.youtube.com/watch?v=gN24W_psMpE")
	fm1 := vi.Formats.Best(ytdl.FormatAudioBitrateKey)[0]
	//fm2 := vi.Formats.Worst(ytdl.FormatResolutionKey)[0]

	ext1 := vi.Formats.Best(ytdl.FormatAudioBitrateKey)[0].Extension
	//ext2 := vi.Formats.Worst(ytdl.FormatResolutionKey)[0].Extension
	err = services.DownloadYTCache(vi, fm1, ext1)

	if err != nil {
		fmt.Println("error dling vid bruh, ", err)
	}

}
