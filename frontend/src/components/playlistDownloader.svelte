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
{/if}

<style>
    h2 {
        font-size: 1.5rem;
        line-height: 1.875rem;
        letter-spacing: -.00625em;
        font-weight: 500;
        border-bottom: 1px solid hsl(var(--border));
        padding-bottom: .5rem;
    }
</style>