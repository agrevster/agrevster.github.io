<script>
import Icon from "@iconify/svelte";

export let stars

let gold = "#fcba03"

</script>

<div class="skill-stars no-select">
    {#each {length: 3} as _, i}
        {#if stars > i}
            <Icon icon="ic:baseline-star" color="{gold}" />
        {/if}
        {#if stars <= i}
           <Icon icon="ic:baseline-star" color="" />
        {/if}
    {/each}
</div>


<style>
    .skill-stars{
        margin-left: 2em;
        display: flex;
        transform: scale(2);
    }
</style>