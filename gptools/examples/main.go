package main

import (
	"context"
)

type Examples struct{}

// example on how to return a transcript from a video file
func (m *Examples) GptoolsTranscript(
	ctx context.Context,
	// free to use movie https://mango.blender.org/about/
	//+default="http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/TearsOfSteel.mp4"
	url string,
) (string, error) {
	video := dag.HTTP(url)
	return dag.Gptools().Transcript(video).Contents(ctx)
}

// example on how to get a transcript for a video and
// return it as a *Directory to use in other functions
func (m *Examples) GptoolsTranscript_Directory(
	ctx context.Context,
	// free to use movie https://mango.blender.org/about/
	//+default="http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/TearsOfSteel.mp4"
	url string,
) *Directory {
	video := dag.HTTP(url)
	return dag.Directory().WithFile("video-transcript.txt",
		dag.Gptools().Transcript(video))
}
