<script>
    import Icon from "@iconify/svelte";
    import AboutSkill from "./AboutSkill.svelte";
    import {expandedStore} from "../../state.js";

    export let data

    let title = data.title
    let icon = data.icon
    let skills = data.skills

    function toggleExpanded() {
        if ($expandedStore === title) expandedStore.set("")
        else expandedStore.set(title)
    }

    skills = skills.sort((a, b) => {
        if (a.title > b.title){
            return 1
        }
        if (a.title < b.title){
            return -1
        }
        return 0
    })
</script>

<div class="category">
     <div class="skill-category-icon no-select">
        <Icon icon="{icon}"/>
    </div>
    <p class="skill-category-title no-select">{title}</p>
    <div class="skill-category-expanded no-select" on:click={toggleExpanded} role="button" tabindex="0" on:keypress={toggleExpanded}>
        {#if title.toString() === $expandedStore}
            <Icon icon="quill:expand"/>
        {/if}
        {#if title.toString() !== $expandedStore}
            <Icon icon="quill:collapse"/>
        {/if}
    </div>
</div>
{#if title.toString() === $expandedStore}
    <div class="expanded-skills">
        {#each skills as skill}
            <AboutSkill data="{skill}"/>
        {/each}
    </div>
{/if}

<style>
    .category {
        display: flex;

        flex-direction: row;
        justify-content: center;
        flex-wrap: nowrap;
        align-items: center;
    }
    .skill-category-title{
        font-size: 1.2em;
        text-transform: capitalize;
    }
    .skill-category-icon{
        transform: scale(2.5);
        margin-right: .9em;
    }
    .skill-category-expanded{
        transform: scale(1.5);
        margin-left: .5em;
    }
    .skill-category-expanded:hover{
        animation-name: expand-grow;
        animation-duration: .3s;
        transform: scale(1.6);
    }

    @keyframes expand-grow {
        from {transform: scale(1.5)}
        to {transform: scale(1.8)}
    }

    .expanded-skills {
        margin-left: 1em;
        margin-right: 1em;
        border-style: ridge;
        border-width: .2em;
        border-radius: .2em;
        border-color: var(--morning_sea);
    }

</style>