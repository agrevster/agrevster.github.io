<script lang="js">
    import Icon from '@iconify/svelte';
    import AboutExperienceType from "./AboutExperienceType.svelte";

    export let data

    let dates = data.dates

    let title = data.title
    let type = data.type
    let description = data.description
    let technologies = data.technologies
    let technologiesIcons = data.technologiesIcons
    let link = data.link
</script>

<div class = "about-experience">
    <p class = "experience-title"><a href="{link}" target="_blank">{title}</a><br>
        <AboutExperienceType experienceType="{type}"/>
        <br><span class = "experience-description"> {description}</span></p>
    <div class="flex">
        <div class="vertical-flex">
            <p class="experience-dates">{dates}</p>
        </div>
        {#each technologies as technology, i}
            <div class="technologies no-select">
                <div title="{technology}" class="technologies-icon">
                    <Icon icon={technologiesIcons[i]}/>
                </div>
                <p class="technologies-label">{technology}</p>
            </div>
        {/each}

    </div>
</div>

<style>
    .flex {
        display: flex;
        flex-direction: row-reverse;
        justify-content: space-between;
    }

    .vertical-flex {
        display: flex;
        flex-direction: column;
    }

    .about-experience {
        background: var(--sail_blue);
        border-radius: .4em;
        padding: .2em;
    }

    .experience-title {
        text-transform: capitalize;
        font-size: 1.1em;
    }

    .experience-description {
        font-family: Inter, sans-serif;
        font-size: .8em;
    }

    .experience-dates {
        font-size: .6em;
    }

    .technologies {
        padding-left: .2em;

        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .technologies-label {
        margin-left: .5em;
        font-size: 0.6em;
    }

    .technologies-icon {
        transform: scale(1.5);
    }

    @media (max-width: 500px) {
        .experience-title {
            font-size: 2vw;
        }

        .technologies-label {
            font-size: 1vw;
        }

        .technologies-icon {
            transform: scale(2.5);
        }
    }
</style>