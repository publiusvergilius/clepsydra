<script lang="ts">
    import { CreateDies } from '../wailsjs/go/main/App.js'
    export let diei = [];
    export let refetch: () => void;
    export let handleClick: (dies: { date: string }) => void;

    function optimisticUpdate () {
        const options = {
             day: "2-digit", 
             year: "numeric", 
             month: "2-digit" 
        } as Intl.DateTimeFormatOptions;

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
                <div class="item">
                    <span on:click={()=>handleClick(dies)}>{dies?.date}</span>
                </div>
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
    .item {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    }

    .item > span{
        background-color: rgba(255, 255, 255, 0.4);
        margin-bottom: 1rem;
        cursor: pointer;
        font-size: 2rem;
        font-weight: 600;
        padding: 1rem;
        border-radius: 0.5rem;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);color: #333;
    }

</style>