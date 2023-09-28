<script lang="ts">
    export let items: any[];
    
    import { createSelect, melt } from '@melt-ui/svelte';

    export let {
    elements: { trigger, menu, option, label },
    states: { selectedLabel, open, selected },
    helpers: { isSelected },
    } = createSelect({
        forceVisible: true,
        positioning: {
        placement: 'bottom',
        fitViewport: true,
        sameWidth: true,
        },
    });
</script>

    <!-- svelte-ignore a11y-label-has-associated-control -->
    <label class="label" use:melt={$label}>Quality</label>
    <button class="outline-button select-trigger"
    use:melt={$trigger}
    on:m-keydown={(e) => {
    e.preventDefault(); // Cancel default builder behabiour

    const { key } = e.detail.originalEvent;

    if (!['ArrowDown', 'ArrowUp', 'Space', 'Enter'].includes(key)) return;

    const allOptions = items;
    const index = allOptions.indexOf(`${$selectedLabel}`);

    if (key === 'ArrowDown') {
      const nextIndex = index + 1;
      const nextOption = allOptions[nextIndex] || allOptions[0];
      selected.set({ value: nextOption, label: nextOption });
    } else if (key === 'ArrowUp') {
      const prevIndex = index - 1;
      const prevOption =
        allOptions[prevIndex] || allOptions[allOptions.length - 1];
      selected.set({ value: prevOption, label: prevOption });
    } else {
      open.set(true);
    }
  }}
  aria-label="Quality"
>
  {$selectedLabel || 'Select quality'}
  <svg xmlns="http://www.w3.org/2000/svg" class="ml-2" width="16" height="16" fill="currentColor" viewBox="0 0 256 256"><path d="M216.49,104.49l-80,80a12,12,0,0,1-17,0l-80-80a12,12,0,0,1,17-17L128,159l71.51-71.52a12,12,0,0,1,17,17Z"></path></svg>  </button>
    {#if $open}
    <div class="select-content"
        use:melt={$menu}
    >
        {#each items as item}
            <div
            use:melt={$option({ value: item, label: item })}
            class="select-item"
            >
            <div class="{$isSelected(item) ? 'block' : 'hidden'}">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" viewBox="0 0 256 256"><path d="M232.49,80.49l-128,128a12,12,0,0,1-17,0l-56-56a12,12,0,1,1,17-17L96,183,215.51,63.51a12,12,0,0,1,17,17Z"></path></svg>            </div>
            {item}
            </div>
        {/each}
    </div>
    {/if}

<style>
    .outline-button {
        border: 1px solid hsl(var(--input)); 
        background-color: inherit;
        color: inherit;
    }

    .outline-button:hover {
        background-color: hsl(var(--accent));
    }

    .block {
        display: block;
    }

    .hidden {
        display: none;
    }

    .ml-2 {
        margin-left: .5rem;
    }

    .content-start {
        justify-content: start;
    }

    .select-trigger {
        display: inline-flex;
        height: 2.5rem;
        padding: 0 1rem;
        align-items: center;
        justify-content: space-between;
        border-radius: 4.8px;
        font-size: .875rem;
        line-height: 1.25rem;
        font-weight: 500;
        font-family: inherit;
        transition: background-color 150ms ease;
        cursor: pointer;
        width: 11.25rem;
    }

    .select-content {
        border-radius: 4.8px;
        border: 1px solid hsl(var(--input));
        background-color: hsl(var(--background));
        padding: .25rem;
    }

    .input-group {
        display: grid;
        gap: .5rem;
        width: 100%;
    }
</style>