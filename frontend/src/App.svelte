<script lang="ts">
    import Select from './components/select.svelte';
    import type { Readable } from 'svelte/store';
    import { DownloadAudio, DownloadByQuality, DownloadVideo } from "../wailsjs/go/main/App.js"

    let url = "";
    let filename = ""
    let format = "video";
    let downloading = false;
    let quality: Readable<string>;
    let response: string = ""; 

    const qualitys = [
        "Default",
        "1080p",
        "720p",
        "480p",
        "360p",
        "240p",
        "144p"
    ];

    function download() {
        if (format === "audio") {
            downloadAudio();
            return;
        }

        if ($quality === "Default" || $quality === "") {
            downloadVideo();
            return;
        }
        downloadByQuality();
    }

    function downloadVideo() {
        downloading = true;
        DownloadVideo(url, filename)
        .then((res) => {
            response = res;
            downloading = false;
        });
    }

    function downloadAudio() {
        downloading = true;
        DownloadAudio(url, filename)
        .then((res) => {
            response = res;
            downloading = false;
        });
    }

    function downloadByQuality() {
        downloading = true;
        DownloadByQuality(url, filename, $quality)
        .then((res) => {
            response = res;
            downloading = false;
        });
    }
</script>

<nav>
  <h1>Youtube downloader</h1>
</nav>
<main>
    <div class="input-group">
        <label class="label" for="url">URL</label>
        <input autocomplete="off" id="url" bind:value={url} class="input" placeholder="https://youtube.com"/>
    </div>

    <div class="input-group mt-3">
        <label class="label" for="url">Filename</label>
        <input autocomplete="off" id="url" bind:value={filename} class="input" placeholder="funny-video"/>
    </div>
    
    <div class="input-group mt-3">
        <span class="label">Format</span>
        <div class="radio-group">
            <span>
                <input type="radio" id="video" name="format" value="video" class="sr-only" bind:group={format}>
                <label for="video" class="button outline-button">Video</label>
            </span>
            <span>
                <input type="radio" id="audio" name="format" value="audio" class="sr-only" bind:group={format}>
                <label for="audio" class="button outline-button">Audio</label>
            </span>
        </div>
    </div>

    <div class="qualitys" aria-expanded={format === "video" ? "true" : "false"}>
        <div class="input-group">
            <Select items={qualitys} bind:selectedLabel={quality} />
        </div>
    </div>
    
    
    <button class="button primary-button mt-3" on:click={download}>{ downloading ? `Downloading ${format}` : `Download ${format}` }</button>
    <p class="message mt-3">{ response }</p>
</main>

<style>
    nav {
        padding: 1rem;
        border-bottom: 1px solid hsl(var(--border));
    }

    h1 {
        font-size: 1.5rem;
        font-weight: 700;
    }

    main {
        display: flex;
        flex-direction: column;
        align-items: start;
        justify-content: start;
        max-width: 376px;
        margin-inline: auto;
        padding-inline: 1rem;
        margin-top: 1.5rem;
    }

    .label {
        font-size: .875rem;
        line-height: 1;
        font-weight: 500;
    }

    .sr-only {
        position: absolute;
        width: 1px;
        height: 1px;
        padding: 0;
        margin: -1px;
        overflow: hidden;
        clip: rect(0, 0, 0, 0);
        white-space: nowrap;
        border-width: 0;
    }

    .input-group {
        display: grid;
        gap: .5rem;
        width: 100%;
        overflow: hidden;
    }

    .radio-group {
        display: flex;
        gap: .5rem;
    }

    .button {
        display: inline-flex;
        height: 2.5rem;
        padding: 0 1rem;
        align-items: center;
        justify-content: center;
        border-radius: 4.8px;
        font-size: .875rem;
        line-height: 1.25rem;
        font-weight: 500;
        font-family: inherit;
        border: none;
        transition: background-color 150ms ease;
        cursor: pointer;
    }

    .primary-button {
        background-color: hsl(var(--primary));
        color: hsl(var(--primary-foreground));
    }

    .primary-button:hover {
        background-color: hsl(var(--primary) / .9);
    }

    .outline-button {
        border: 1px solid hsl(var(--input)); 
        background-color: inherit;
        color: inherit;
    }

    .outline-button:hover {
        background-color: hsl(var(--accent));
    }

    .input {
        width: 100%;
        height: 2.5rem;
        font-size: .875rem;
        line-height: 1.25rem;
        padding-block: .5rem;
        padding-inline: .75rem;
        border: 1px solid hsl(var(--input));
        border-radius: 4.8px;
        display: flex;
        background-color: inherit;
        color: inherit;
        outline: none;
    }

    .input:focus-visible {
        border: 1px solid hsl(var(--primary-foreground));
    }

    input:checked + * {
        border-color: hsl(var(--primary));
    }

    .message {
        font-size: .875rem;
        font-weight: 300;
    }

    .qualitys {
        display: grid;
        grid-template-rows: 0fr;
        transition: grid-template-rows 150ms ease;
        overflow: hidden;
        visibility: collapse;
    }

    .qualitys[aria-expanded="true"] {
        grid-template-rows: 1fr;
        visibility: visible;
        margin-top: .75rem;
    }

    .mt-3 {
        margin-top: .75rem;
    }
</style>
