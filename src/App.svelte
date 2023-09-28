<script>
    import portrait from "./assets/portrait.png"
    import Divider from "./lib/Divider.svelte";
    import About from "./lib/about me/About.svelte";
</script>

<main>
    <link href="https://fonts.googleapis.com" rel="preconnect">
    <link crossorigin="anonymous" href="https://fonts.gstatic.com" rel="preconnect">
    <script crossorigin="anonymous" src="https://kit.fontawesome.com/30df65de2e.js"></script>
    <link href="https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@200;400;800&display=swap" rel="stylesheet">
    <div class="card">
        <img alt="" class="no-select" id="portrait" src="{portrait}" title="An AI generated lighthouse picture">
        <p id = "nametag">Hello, I'm Alex
            <br>
            <span id = "description">[Developer and barbeque enthusiast]</span>
        </p>
        <Divider/>
        <About/>
    </div>
</main>

<style>
    #nametag {
        align-self: flex-start;
        font-family: 'JetBrains Mono', monospace;
        font-weight: 800;
        font-size: 1.5rem;
    }

    #portrait {
        margin-top: 1em;
        width: 10em;
        height: 10em;
        border-radius: 4.0em;
        box-shadow: .5em .5em var(--deep_ocean);
    }

    #description {
        font-family: 'JetBrains Mono', monospace;
        font-weight: 400;
        font-size: 1.0rem;
    }

    @media (max-width: 500px) {
        #description {
           font-size: 2.3vw;
        }

        #nametag {
            font-size: 5vw;
        }

        #portrait {
            width: 7em;
            height: 7em;
            border-radius: 2.5em;
        }
    }
</style>