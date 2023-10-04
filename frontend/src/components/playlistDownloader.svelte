<script lang="ts">
    import { GetPlaylist } from "../../wailsjs/go/main/App.js";

    let playlist: any;
    let url = "";

    function getPlaylist() {
        if (url.trim() === "") return;
        GetPlaylist(url)
        .then((value) => playlist = value);
    }
</script>

<div class="input-group">
    <label for="url" class="label">Playlist URL</label>
    <input name="url" id="url" class="input" type="text" bind:value={url} placeholder="https://youtube.com/playlist">
</div>

<button class="button primary-button mt-3" on:click={getPlaylist}>Get playlist</button>

{#if playlist}
    <h2 class="mt-3">{ playlist.Title }</h2>

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