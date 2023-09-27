<script lang="ts">
  import {DownloadVideo} from '../wailsjs/go/main/App.js'

  let url;
</script>

<main>
  <div class="input-box" id="input">
    <input autocomplete="off" bind:value={url} class="input" id="name" type="text"/>
    <button on:click={() => DownloadVideo(url)}>download</button>
  </div>
</main>

<style>
  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }
  button {
    display: inline-flex;
    height: 2.5rem;
    padding: 0 1rem;
    align-items: center;
    justify-content: center;
    border-radius: 4.8px;
    font-size: .825rem;
    font-weight: 500;
    border: none;
    background-color: hsl(var(--primary));
    color: hsl(var(--primary-foreground));
    transition: background-color 150ms ease;
  }

  button:hover {
    background-color: hsl(var(--primary) / .9);
  }
</style>
