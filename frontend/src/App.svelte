<script lang="ts">
    import VideoDownloader from "./components/videoDownloader.svelte";

    let page: "video" | "playlist" = "video";
</script>

<nav>
    <h1>YT downloader</h1>
    <div>
        <button on:click={() => page = "video"}>Video</button>
        <button on:click={() => page = "playlist"}>Playlist</button>
    </div>
</nav>

<main>
    {#if page === "video"}
        <VideoDownloader />
    {:else}
        <p>Playlist</p>
    {/if}
</main>

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
        max-width: 376px;
        margin-inline: auto;
        padding-inline: 1rem;
        margin-top: 1.5rem;
    }
</style>
