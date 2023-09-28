<script lang="js">

    export let icon
    export let title
    export let href = null
    export let store = null

    function onClick() {
        if (href != null) {
            window.open(href)
        } else {
            store.set(title)
        }
    }


</script>
<i class="group-picker-option fa-solid {icon}" on:click={onClick} on:keydown={null} role="button"
   tabindex="{0}">
    <br>
    <span class="group-picker-option-label">{title}</span>
</i>

<style>
    .group-picker-option-label {
        font-size: .5em;
    }

    .group-picker-option {
        font-size: 1.5em;
        padding: .5em;
        -webkit-user-select: none;
        -ms-user-select: none;
        user-select: none;
        cursor: pointer;
    }

    @keyframes fade-on-hover {
        from {
            background-color: var(--deep_ocean);
        }
        to {
            background-color: var(--morning_sea);
        }
    }

    .group-picker-option:hover {
        background: var(--morning_sea);
        border-radius: .15em;

        animation-name: fade-on-hover;
        animation-duration: 1s;
    }


    @media (max-width: 500px) {
        .group-picker-option{
            font-size: 2.5vw;
        }
    }
</style>