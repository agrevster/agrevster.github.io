<script>
    import portrait from "./assets/portrait.png"
    import Divider from "./lib/Divider.svelte";
    import About from "./lib/about me/About.svelte";
</script>

<main>
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="anonymous">
    <script src="https://kit.fontawesome.com/30df65de2e.js" crossorigin="anonymous"></script>
    <link href="https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@200;400;800&display=swap" rel="stylesheet">
    <div class = "card">
        <img id = "portrait" class = "no-select" src = "{portrait}" alt="" title="An AI generated lighthouse picture">
        <p id = "nametag" >Hello, I'm Alex
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
        font-size:1.0rem;
    }
</style>