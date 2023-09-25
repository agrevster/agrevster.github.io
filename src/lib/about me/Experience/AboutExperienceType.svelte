<script lang="js">
    export let experienceType
</script>


<span>[<span class="experience-type experience-type-{experienceType}">{experienceType}</span>]</span>


<style>
    span {
        font-family: Inter, sans-serif;
        font-size: .8em;
    }

    .experience-type {
        font-size: .85em;
    }

    .experience-type-project {
        color: darkorange;
    }

    .experience-type-job {
        color: seagreen;
    }

    .experience-type-other {
        color: cornflowerblue;
    }

    .experience-type-freelance {
        color: mediumpurple;
    }
</style>