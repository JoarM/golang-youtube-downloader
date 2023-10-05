<script lang="ts">
    import { CreateDir, GetPlaylist } from "../../wailsjs/go/main/App.js";
    import Loader from "./ui/loader.svelte";

    let playlist: any;
    let url = "";
    let folder = "";
    let format: "video" | "audio" = "video";
    let getting = false;
    let downloading = false;
    let message = "";

    function getPlaylist() {
        if (url.trim() === "") return;
        if (downloading) return;
        getting = true;
        GetPlaylist(url)
        .then((value) => {
            playlist = value;
            getting = false;
        });
    }

    async function downloadPlaylist() {
        if (!playlist) return;
        if (getting) return;
        downloading = true;
        let res = await CreateDir(folder);
        if (res != "") {
            message = res; 
            downloading = false;
            return;
        }
        if (format === "video") {
            downloading = false;
        } else {

        }
    }
</script>

<div class="input-group">
    <label for="url" class="label">Playlist URL</label>
    <input name="url" id="url" class="input" type="text" bind:value={url} placeholder="https://youtube.com/playlist">
</div>

<button class="button primary-button mt-3" on:click={getPlaylist}>
    {#if getting}
        <Loader />
    {/if}
    <span>
        { getting ? "Getting playlist" : "Get playlist" }
    </span>
</button>

{#if playlist}
    <h2 class="mt-3">{ playlist.Title }</h2>

    <div class="input-group mt-3">
        <label for="folder" class="label">Folder name</label>
        <input type="text" class="input" id="folder" name="folder" bind:value={folder}>
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

    <button class="button primary-button mt-3" on:click={downloadPlaylist}>
        {#if downloading}
            <Loader />
        {/if}
        <span>
            { downloading ? "Downloading" : "Download all" }
        </span>
    </button>

    <p class="message mt-1">{ message }</p>
 
    <ul>
        {#each playlist.Videos as video}
            <li>
                <img src={video.Thumbnails[1].URL} alt="">
                <div class="video-group">
                    <h3>{ video.Title }</h3>
                    <button class="button-sm outline-button">Download</button>
                </div>
            </li>
        {/each}
    </ul>
{/if}

<style>
    h2 {
        font-size: 1.5rem;
        line-height: 1.875rem;
        letter-spacing: -.00625em;
        font-weight: 500;
        border-bottom: 1px solid hsl(var(--border));
        padding-bottom: .5rem;
        width: 100%;
    }

    ul {
        margin-block: 1rem;
        display: grid;
        gap: .5rem;
        list-style: none;
    }

    li {
        display: flex;
        align-items: start;
        justify-content: start;
        gap: .5rem;
    }

    img {
        width: 10rem;
        height: 5.625rem;
        border-radius: 4.8px;
        object-fit: cover;
    }

    .video-group {
        display: flex;
        justify-content: space-between;
        align-items: start;
        flex-direction: column;
        height: 5.625rem;
    }

    h3 {
        font-size: .875rem;
        line-height: 1.25rem;
        letter-spacing: 0em;
        font-weight: 500;
        display: -webkit-box;
        overflow: hidden;
        -webkit-box-orient: vertical;
        -webkit-line-clamp: 2;
    }
</style>