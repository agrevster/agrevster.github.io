<script lang="js">
    import AboutExperience from "./AboutExperience.svelte";
    import {onMount} from "svelte";

    export let tab

    function redirectToUniSite() {
        window.open("https://uni.edu", '_blank');
    }

    async function getProjects() {
		const res = await fetch("http://localhost:8080/projects");
		const text = await res.json();

		if (res.ok) {
			return text;
		} else {
			throw new Error(text);
		}
	}

    let projectsPromise = getProjects();

</script>

<div class = "about-content">
    {#if tab === "about"}
        <p>I am a college freshmen at the University of <b on:click = {redirectToUniSite} role = "button" tabindex = {0}
                                                           on:keydown = {null}><span
                class = "uni-purple bold"> Northern</span> <span class = "uni-gold bold">Iowa</span></b> who enjoys
            solving computer problems and developing software.</p>
    {/if}

    {#if tab === "experience"}
        <div class = "experience">
            {#await projectsPromise}
                <p>Loading...</p>
            {:then projects}
                {#each projects as project}
                    <AboutExperience data={project}/>
                    <br>
                {/each}
            {:catch err}
                <p class="error">ERROR: {err}</p>
            {/await}
        </div>
    {/if}
</div>

<style>
    .about-content {
        font-family: "JetBrains Mono", serif;
        word-wrap: break-word;
        white-space: normal;
        font-size: 1vw;
        width: 25vw;
        color: var(--white_wash);
    }

    .uni-purple {
        color: #500778;
    }

    .uni-gold {
        color: #FFB500;
    }

    b {
        cursor: pointer;
    }
</style>