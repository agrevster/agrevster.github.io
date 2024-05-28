<script lang="js">
    import AboutGroupPickerOption from "./AboutGroupPickerOption.svelte";
    import {writable} from "svelte/store";
    import AboutContent from "./AboutContent.svelte";

    let selectedTab = "about"


    const tabStore = writable("about")

    tabStore.subscribe(newValue => {
        selectedTab = newValue
    })
</script>
<div id="about-me">
    <div class="no-select" id="about-me-options">
        <AboutGroupPickerOption href={null} icon="fa-mug-hot" store={tabStore} title="about"/>
        <AboutGroupPickerOption href={null} icon="fa-laptop-code" store={tabStore} title="experience"/>
        <AboutGroupPickerOption href={null} icon="fa-passport" store={tabStore} title="skills"/>
        <AboutGroupPickerOption href="https://github.com/agrevster" icon="fa-github fa-brands" store="{tabStore}"
                                title="github"/>
    </div>

    <div id="about-me-content">
        <AboutContent tab="{selectedTab}"/>
    </div>

</div>
<br>
<style>
    #about-me-options {
        margin-bottom: 0em;
    }

    #about-me {
        background: var(--deep_ocean);
        padding: 1rem;
        border-radius: .3rem;
    }
</style>