package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/kkdai/youtube/v2"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) DownloadVideo(url string, filename string) string {
	client := youtube.Client{}

	video, err := client.GetVideo(url)
	if err != nil {
		return ""
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		return ""
	}
	defer stream.Close()

	basepath, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	filepath := filepath.Join(basepath, "Downloads", filename+".mp4")

	file, err := os.Create(filepath)
	if err != nil {
		return ""
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		return ""
	}

	return "Donloaded at: " + filepath
}

func (a *App) DownloadAudio(url string, filename string) string {
	client := youtube.Client{}

	video, err := client.GetVideo(url)
	if err != nil {
		return "Couldnt get video"
	}

	formats := video.Formats.Type("audio") // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		return "Error getting stream"
	}
	defer stream.Close()

	basepath, err := os.UserHomeDir()
	if err != nil {
		return "Couldnt get home directory"
	}

	filepath := filepath.Join(basepath, "Downloads", filename+".mp3")

	file, err := os.Create(filepath)
	if err != nil {
		return "Failed to create file"
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		return "Failed to copy to file"
	}

	return "Donloaded at: " + filepath
}
