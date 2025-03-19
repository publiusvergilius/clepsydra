<script lang="ts">
    import { CreateDies } from '../wailsjs/go/main/App.js'
    export let diei = [];
    export let refetch;

    function optimisticUpdate () {
        const options = { day: "2-digit", year: "numeric", month: "2-digit" };
        const newDate = new Date().toLocaleDateString('en-US', options);
        const json = JSON.stringify({dies: newDate, format: "02/02/2021"});
        CreateDies(json).then(result => console.log(result));
        diei = [...diei, { date: newDate }];
        refetch();
    }
</script>

<main>
    <button class="add-button" on:click={optimisticUpdate}>+</button>
    <div class="layout">
        {#if (diei.length === 0)}
            <h2>NULLA DIES SCRIPTA.</h2>
        {/if}
        {#if (diei.length > 0)}
            {#each diei as dies}
            <div class="box">
            <h1 class="primary_text">{dies?.date}</h1>
            </div>
            {/each}
        {/if}
    </div> 
</main>

<style>
    .add-button {
        position: fixed;
        bottom: 30px;
        right: 30px;
        padding: 0.9rem 2.1rem;
        text-align: center;
        border-radius: 50%;
        background-color: #f1f1f1;
        color: #333;
        font-size: 4rem;
        border: none;
        cursor: pointer;
    }

</style>