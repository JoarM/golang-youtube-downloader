<script lang="ts">
    import { createTooltip, melt } from '@melt-ui/svelte';
    import { fade } from 'svelte/transition';
    
    const {
        elements: { trigger, content, arrow },
        states: { open },
    } = createTooltip({
        positioning: {
            placement: 'right',
        },
        openDelay: 500,
        closeOnPointerDown: false,
        forceVisible: true,
    });
</script>

<button type="button" use:melt={$trigger} class="trigger" aria-label="tooltip">
    <slot />
</button>

{#if $open}
    <div 
    use:melt={$content}
    transition:fade={{ duration: 100 }}
    class="content"
    >
        <div use:melt={$arrow} />
        <slot name="tooltip"  />
    </div>
{/if}

<style>
    .trigger {
        background-color: transparent;
        border: none;
        outline: none;
        color: hsl(var(--foreground));
        height: 1.25rem;
        width: 1.25rem;
        display: flex;
        align-items: center;
        justify-content: center;
        border-radius: 100%;
        transition: background-color 150ms ease,
        color 150ms ease;
    }

    .content {
        background-color: hsl(var(--accent));
        color: hsl(var(--accent-foreground));
        margin-right: 1rem;
        z-index: 10;
        border-radius: 4.8px;
    }

    .content:last-child {
        padding-inline: .5rem;
        padding-block: .25rem;
        font-size: .75rem;
        font-weight: 500;
    }
</style>