<script lang="js">
    import Icon from '@iconify/svelte';
    import SkillStars from "./SkillStars.svelte";

    export let data

    let title = data.title
    let icon = data.icon
    let level = data.stars
</script>

<div class = "skill-container">
    <div class = "skill-icon no-select">
        <Icon icon="{icon}"/>
    </div>
    <p class = "skill-title">{title}</p>
    <p class = "skill-dash no-select">-</p>
    <SkillStars stars={level}/>
</div>
<style>
    .skill-container {
        display: flex;

        flex-direction: row;
        justify-content: center;
        flex-wrap: nowrap;
        align-items: center;
    }

    .skill-title {
        text-transform: capitalize;
        font-size: 1em;
    }

    .skill-dash {
        margin-left: .5em;
        font-size: 1.2em;
    }

    .skill-icon {
        transform: scale(2);
        margin-right: 1em;
    }


    @media (max-width: 500px) {
        .skill-icon {
            transform: scale(2.5)
        }

        .skill-title {
            font-size: 1.2vw;
        }
    }
</style>