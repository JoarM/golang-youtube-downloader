<script lang="ts">
    import Select from './components/select.svelte';
    import type { Readable } from 'svelte/store';
    import { DownloadAudio, DownloadVideo } from "../wailsjs/go/main/App.js"

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

    // temp out of use
    function download() {
        if ($quality === "Default" || $quality === "") {
            downloading = true;
            DownloadVideo(url, filename).then((res) => {
                response = res;
                downloading = false;
            });
        }
    }

    function simpleDownload() {
        if (format === "video") {
            downloading = true;
            DownloadVideo(url, filename).then((res) => {
                response = res;
                downloading = false;
            });
        } else if (format === "audio") {
            downloading = true;
            DownloadAudio(url, filename).then((res) => {
                response = res;
                downloading = false;
            });
        }
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

    <div class="input-group">
        <label class="label" for="url">Filename</label>
        <input autocomplete="off" id="url" bind:value={filename} class="input" placeholder="funny-video"/>
    </div>
    
    <div class="input-group">
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

    <!-- <div class="input-group">
        <Select items={qualitys} bind:selectedLabel={quality} />
    </div> -->
    
    <button class="button primary-button" on:click={simpleDownload}>{ downloading ? `Downloading ${format}` : `Download ${format}` }</button>
    <p class="message">{ response }</p>
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
        gap: .75rem
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
    }

    input:checked + * {
        border-color: hsl(var(--primary));
    }

    .message {
        font-size: .875rem;
        font-weight: 300;
    }
</style>
