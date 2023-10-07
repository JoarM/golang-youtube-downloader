<script lang="ts">
    import PlaylistDownloader from "./components/playlistDownloader.svelte";
    import VideoDownloader from "./components/videoDownloader.svelte";
    import { page } from "./lib/state";
</script>

<nav>
    <h1>YT downloader</h1>
    <div>
        <button on:click={() => page.set("video")} aria-current="{$page === "video" ? "true" : "false"}">Video</button>
        <button on:click={() => page.set("playlist")} aria-current="{$page === "playlist" ? "true" : "false"}">Playlist</button>
    </div>
</nav>

<main>
    {#if $page === "video"}
        <VideoDownloader />
    {:else}
        <PlaylistDownloader />
    {/if}
</main>

<span class="version">v 1.0</span>

<style>
    nav {
        padding: 1rem;
        border-bottom: 1px solid hsl(var(--border));
        display: flex;
        gap: .75rem;
        align-items: center;
        justify-content: start;
    }

    h1 {
        font-size: 1.25rem;
        font-weight: 700;
        letter-spacing: -.05;
    }

    main {
        display: flex;
        flex-direction: column;
        align-items: start;
        justify-content: start;
        max-width: 520px;
        margin-inline: auto;
        padding-inline: 1rem;
        margin-top: 1.5rem;
    }

    div {
        border-radius: 4.8px;
        border: 1px solid hsl(var(--border));
        padding: .25rem;
        display: flex;
        align-items: center;
        justify-content: center;
        gap: .25rem;
    }

    button {
        display: inline-flex;
        height: 2.25rem;
        padding: 0 .75rem;
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
        background-color: hsl(var(--background));
        color: hsl(var(--foreground));
    }

    button[aria-current="true"] {
        background-color: hsl(var(--accent));
        color: hsl(var(--accent-foreground));
    }

    button:hover {
        background-color: hsl(var(--secondary));
        color: hsl(var(--secondary-foreground));
    }

    .version {
        position: fixed;
        bottom: .5rem;
        right: .5rem;
        font-size: .75rem;
        font-weight: 300;
    }
</style>
