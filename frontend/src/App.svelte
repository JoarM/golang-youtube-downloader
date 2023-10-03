<script lang="ts">
    import Select from './components/select.svelte';
    import type { Readable } from 'svelte/store';
    import { DownloadAudio, DownloadByQuality, DownloadVideo } from "../wailsjs/go/main/App.js"
    import Tooltip from './components/ui/tooltip.svelte';
    import Loader from './components/ui/loader.svelte';

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
        if (!parse()) {
            return;
        }
        
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

    function parse(): boolean {
        if (url.trim() === "") {
            response = "Please enter a youtube link";
            return false;
        } 

        if (filename.trim() === "") {
            response = "Please enter a filename";
            return false;
        }

        return true;
    }
</script>

<nav>
  <h1>Youtube downloader</h1>
</nav>
<main>
    <div class="input-group">
        <div class="label-group">
            <label class="label" for="url">URL</label>
            <Tooltip>
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" viewBox="0 0 256 256"><path d="M128,20A108,108,0,1,0,236,128,108.12,108.12,0,0,0,128,20Zm0,192a84,84,0,1,1,84-84A84.09,84.09,0,0,1,128,212Zm-12-80V80a12,12,0,0,1,24,0v52a12,12,0,0,1-24,0Zm28,40a16,16,0,1,1-16-16A16,16,0,0,1,144,172Z"></path></svg>
                <p slot="tooltip">
                    Link or id to the video you want to download
                </p>
            </Tooltip>
        </div>
        
        <input autocomplete="off" id="url" bind:value={url} class="input" placeholder="https://youtube.com"/>
    </div>

    <div class="input-group mt-3">
        <div class="label-group">
            <label class="label" for="url">Filename</label>
            <Tooltip>
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" viewBox="0 0 256 256"><path d="M240.26,186.1,152.81,34.23h0a28.74,28.74,0,0,0-49.62,0L15.74,186.1a27.45,27.45,0,0,0,0,27.71A28.31,28.31,0,0,0,40.55,228h174.9a28.31,28.31,0,0,0,24.79-14.19A27.45,27.45,0,0,0,240.26,186.1Zm-20.8,15.7a4.46,4.46,0,0,1-4,2.2H40.55a4.46,4.46,0,0,1-4-2.2,3.56,3.56,0,0,1,0-3.73L124,46.2a4.77,4.77,0,0,1,8,0l87.44,151.87A3.56,3.56,0,0,1,219.46,201.8ZM116,136V104a12,12,0,0,1,24,0v32a12,12,0,0,1-24,0Zm28,40a16,16,0,1,1-16-16A16,16,0,0,1,144,176Z"></path></svg>
                <p slot="tooltip">
                    Name of the output file be careful not to overwrite other files
                </p>
            </Tooltip>
        </div>
        
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
    
    <button class="button primary-button mt-3" on:click={download}>
        {#if downloading}
            <Loader />
        {/if}
        <span>
            { downloading ? `Downloading ${format}` : `Download ${format}` }
        </span>
    </button>
    <p class="message mt-1">{ response }</p>

    <span class="version">v 1.0</span>
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

    .mt-1 {
        margin-top: .25rem;
    }

    .version {
        position: absolute;
        bottom: .5rem;
        right: .5rem;
        font-size: .75rem;
        font-weight: 300;
    }

    .label-group {
        display: flex;
        gap: .25rem;
        align-items: center;
        justify-content: start;
    }
</style>
