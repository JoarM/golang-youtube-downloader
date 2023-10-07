import { writable } from "svelte/store";
import { DownloadAudio, DownloadByQuality, DownloadVideo, GetPlaylist, CreateFolder } from "../../wailsjs/go/main/App.js"

export const page = writable<"video" | "playlist">("video")

interface videoInterface {
    url: string;
    format: "video" | "audio";
    filename: string;
    downloading: boolean;
    status: string;
}

export const video = writable<videoInterface>({
    downloading: false,
    status: "",
    url: "",
    format: "video",
    filename: "",
});

interface playlistInterface {
    url: string;
    fetching: boolean;
    folder: string;
    playlist: any;
    format: "video" | "audio";
    downloading: boolean;
    status: string;
}

export const playlist = writable<playlistInterface>({
    url: "",
    fetching: false,
    folder: "",
    playlist: null,
    format: "video",
    downloading: false,
    status: "",
});

export async function download (url: string, filename: string, format: string, quality: string) {
    if (!parse(url, filename)) {
        return;
    }

    video.update((prev) =>  {
        return {
            ...prev,
            downloading: true,
        }
    });
    
    if (format === "audio") {
        const res = await DownloadAudio(url, filename);
        video.update((prev) =>  {
            return {
                ...prev,
                downloading: false,
                status: res,
            }
        });
        return;
    }

    if (quality === "Default" || quality === "") {
        const res = await DownloadVideo(url, filename);
        video.update((prev) =>  {
            return {
                ...prev,
                downloading: false,
                status: res,
            }
        });
        return;
    }
    const res = await DownloadByQuality(url, filename, quality);
    video.update((prev) =>  {
        return {
            ...prev,
            downloading: false,
            status: res,
        }
    });
}

export function fetchPlaylist(url: string) {
    if (url.trim() === "") return;
    
    playlist.update((prev) => {
        return {
            ...prev,
            fetching: true,
        }
    });

    GetPlaylist(url)
    .then((value) => {
        if (value === null) {
            playlist.update((prev) => {
                return {
                    ...prev,
                    fetching: false,
                    status: "Failed to get playlist, make sure the playlist isnt private",
                }
            });
            return;
        }
        playlist.update((prev) => {
            return {
                ...prev,
                playlist: value,
                fetching: false,
            }
        });
    });
}

export function getIndividual(url: string) {
    video.update((prev) => {
        return {
            ...prev,
            url: `https://www.youtube.com/watch?v=${url}`,
            filename: "",
        }
    });
    page.set("video");
}

export async function getAll(folder: string, data: any, format: "video" | "audio") {
    playlist.update((prev) => {
        return {
            ...prev,
            downloading: true,
        }
    });

    const videos = data.Videos;

    const res = await CreateFolder(folder);
    if (res != "") {
        playlist.update((prev) => {
            return {
                ...prev,
                downloading: false,
                status: res,
            }
        });
        return;
    }

    if (format === "video") {
        for (let i = 0; i < videos.length; i++) {
            const res = await DownloadVideo(videos[i].ID as string, `${folder}/video${i + 1}`);
            playlist.update((prev) => {
                return {
                    ...prev,
                    status: res,
                }
            });
        }
    } else {
        for (let i = 0; i < videos.length; i++) {
            const res = await DownloadAudio(videos[i].ID as string, `${folder}/video${i + 1}`);
            playlist.update((prev) => {
                return {
                    ...prev,
                    status: res,
                }
            });
        }
    }

    playlist.update((prev) => {
        return {
            ...prev,
            downloading: false,
            status: "Download complete!"
        }
    });
}

function parse(url: string, filename: string): boolean {
    if (url.trim() === "") {
        video.update((prev) => {
            return {
                ...prev,
                status: "Please enter a youtube link"
            }
        });
        return false;
    } 
    if (filename.trim() === "") {
        video.update((prev) => {
            return {
                ...prev,
                status: "Please enter a filename",
            }
        });
        return false;
    }
    return true;
}