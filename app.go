package main

import (
	"context"
	"io"
	"os"
	"os/exec"
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

func (a *App) DownloadVideo(url string, filename string) string {
	client := youtube.Client{}

	video, err := client.GetVideo(url)
	if err != nil {
		return "Couldnt get video"
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		return "Error getting stream"
	}
	defer stream.Close()

	basepath, err := os.UserHomeDir()
	if err != nil {
		return "Couldnt find home directory"
	}

	filepath := filepath.Join(basepath, "Downloads", filename+".mp4")

	file, err := os.Create(filepath)
	if err != nil {
		return "Failed to create file"
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		return "Failed to copy to file"
	}

	return "Downloaded at: " + filepath
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
		return "Couldnt find home directory"
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

	return "Downloaded at: " + filepath
}

func (a *App) DownloadByQuality(url string, filename string, quality string) string {
	if err := checkFFMPEG(); err != nil {
		return "FFMPEG required"
	}

	basepath, err := os.UserHomeDir()
	if err != nil {
		return "Couldnt find home directory"
	}

	audioFile, errString, err := DownloadTempAudio(url)
	if err != nil {
		return errString
	}
	defer os.Remove(audioFile.Name())

	videoFile, errString, err := DownloadTempVideo(quality, url)
	if err != nil {
		return errString
	}
	defer os.Remove(videoFile.Name())

	filepath := filepath.Join(basepath, "Downloads", filename+".mp4")

	ffmpegVersionCmd := exec.Command("ffmpeg", "-y",
		"-i", audioFile.Name(),
		"-i", videoFile.Name(),
		"-shortest", // Finish encoding when the shortest input stream ends
		filepath,
		"-loglevel", "warning",
	)
	ffmpegVersionCmd.Stderr = os.Stderr
	ffmpegVersionCmd.Stdout = os.Stdout

	err = ffmpegVersionCmd.Run()
	if err != nil {
		return "FFMPEG failed"
	}

	return "Downloaded at: " + filepath
}

func DownloadTempAudio(url string) (*os.File, string, error) {
	youtube := youtube.Client{}
	video, err := youtube.GetVideo(url)
	if err != nil {
		return nil, "Failed to get audio", err
	}

	formats := video.Formats.Type("audio")
	stream, _, err := youtube.GetStream(video, &formats[0])
	if err != nil {
		return nil, "Failed to get audio stream", err
	}
	defer stream.Close()

	file, err := os.CreateTemp("", "youtube.mp3")
	if err != nil {
		return nil, "Failed to create audio file", err
	}

	_, err = io.Copy(file, stream)
	if err != nil {
		return nil, "Failed to copy audio to file", err
	}

	return file, "", nil
}

func DownloadTempVideo(quality string, url string) (*os.File, string, error) {
	youtube := youtube.Client{}
	video, err := youtube.GetVideo(url)
	if err != nil {
		return nil, "Failed to get video", err
	}

	format := video.Formats.FindByQuality(quality)
	stream, _, err := youtube.GetStream(video, format)
	if err != nil {
		return nil, "Failed to get video stream", err
	}
	defer stream.Close()

	file, err := os.CreateTemp("", "youtube.mp4")
	if err != nil {
		return nil, "Failed to create video file", err
	}

	_, err = io.Copy(file, stream)
	if err != nil {
		return nil, "Failed to copy video to file", err
	}

	return file, "", err
}

func checkFFMPEG() error {
	if err := exec.Command("ffmpeg", "-version").Run(); err != nil {
		return err
	}

	return nil
}

func (a *App) GetPlaylist(url string) *youtube.Playlist {
	client := youtube.Client{}
	playlist, err := client.GetPlaylist(url)
	if err != nil {
		return nil
	}
	return playlist
}
