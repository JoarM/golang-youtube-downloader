<script lang="ts">
    import Loader from "./ui/loader.svelte";
    import { fetchPlaylist, playlist, getIndividual, getAll } from "../lib/state";
</script>

<div class="input-group">
    <label for="url" class="label">Playlist URL</label>
    <input name="url" id="url" class="input" type="text" bind:value={$playlist.url} placeholder="https://youtube.com/playlist">
</div>

<button class="button primary-button mt-3" aria-disabled={$playlist.fetching} on:click={() => fetchPlaylist($playlist.url)}>
    {#if $playlist.fetching}
        <Loader />
    {/if}
    <span>
        { $playlist.fetching ? "Getting playlist" : "Get playlist" }
    </span>
</button>

<p class="message mt-1">{ $playlist.status }</p>
 
{#if $playlist.playlist}
    <h2 class="mt-3">{ $playlist.playlist.Title }</h2>

    <div class="input-group mt-3">
        <label for="folder" class="label">Folder name</label>
        <input type="text" class="input" id="folder" name="folder" bind:value={$playlist.folder}>
    </div>

    <div class="input-group mt-3">
        <span class="label">Format</span>
        <div class="radio-group">
            <span>
                <input type="radio" id="video" name="format" value="video" class="sr-only" bind:group={$playlist.format}>
                <label for="video" class="button outline-button">Video</label>
            </span>
            <span>
                <input type="radio" id="audio" name="format" value="audio" class="sr-only" bind:group={$playlist.format}>
                <label for="audio" class="button outline-button">Audio</label>
            </span>
        </div>
    </div>

    <button class="button primary-button mt-3" aria-disabled={$playlist.downloading} on:click={() => getAll($playlist.folder, $playlist.playlist, $playlist.format)}>
        {#if $playlist.downloading}
            <Loader />
        {/if}
        <span>
            { $playlist.downloading ? "Downloading" : "Download all" }
        </span>
    </button>

    <ul>
        {#each $playlist.playlist.Videos as video}
            <li>
                <img src={video.Thumbnails[1].URL} alt="">
                <div class="video-group">
                    <h3>{ video.Title }</h3>
                    <button class="button-sm outline-button" on:click={() => getIndividual(video.ID)}>Download</button>
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