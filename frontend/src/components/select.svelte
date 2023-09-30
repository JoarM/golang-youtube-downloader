<script lang="ts">
    export let items: any[];
    
    import { createSelect, melt } from '@melt-ui/svelte';
    import Tooltip from './ui/tooltip.svelte';

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
    <div class="label-group">
        <label class="label" use:melt={$label}>Quality</label>

        <Tooltip>
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" viewBox="0 0 256 256"><path d="M128,20A108,108,0,1,0,236,128,108.12,108.12,0,0,0,128,20Zm0,192a84,84,0,1,1,84-84A84.09,84.09,0,0,1,128,212Zm-12-80V80a12,12,0,0,1,24,0v52a12,12,0,0,1-24,0Zm28,40a16,16,0,1,1-16-16A16,16,0,0,1,144,172Z"></path></svg>
            <p slot="tooltip">Qualitys other than default require FFMPEG, note that these downloads take longer do to re-encoding</p>
        </Tooltip>
    </div>
    
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
                <span class="{$isSelected(item) ? 'block' : 'hidden'} select-check">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" viewBox="0 0 256 256"><path d="M232.49,80.49l-128,128a12,12,0,0,1-17,0l-56-56a12,12,0,1,1,17-17L96,183,215.51,63.51a12,12,0,0,1,17,17Z"></path></svg>            
                </span>
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

    .block {
        opacity: 100;
    }

    .hidden {
        opacity: 0;
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
        outline: none;
    }

    .select-trigger:focus-visible {
        border: 1px solid hsl(var(--primary-foreground));
    }

    .select-content {
        border-radius: 4.8px;
        border: 1px solid hsl(var(--input));
        background-color: hsl(var(--background));
        padding: .25rem;
    }

    .select-item {
        padding: .25rem;
        padding-left: 2rem;
        padding-right: .5rem;
        font-size: .875rem;
        line-height: 1.25rem;
        cursor: default;
        border-radius: 4.8px;
        display: flex;
        align-items: center;
        justify-content: start;
        gap: .5rem;
        position: relative;
    }

    .select-item:hover {
        background-color: hsl(var(--secondary));
        color: hsl(var(--secondary-foreground));
    }

    .select-check {
        position: absolute;
        top: 50%;
        transform: translateY(-50%);
        left: .5rem;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .label {
        font-size: .875rem;
        line-height: 1;
        font-weight: 500;
    }

    .label-group {
        display: flex;
        gap: .25rem;
        align-items: center;
        justify-content: start;
    }
</style>