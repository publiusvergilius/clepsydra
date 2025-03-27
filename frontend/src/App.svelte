<script lang="ts">
  import { GetAllDies} from '../wailsjs/go/main/App.js'
  import Diei from "./Diei.svelte";
  import Quartum from './Quartum.svelte';

  let diei = [];
  function fetchAllDies () {
    GetAllDies().then(result => { 
      console.log('result', result) 
      diei = JSON.parse(result)
    })
  }
  fetchAllDies()

  type Dies = {
    id: number,
    date: string,
  }
  let currentDies: Dies | null = null;

  function handleClick(dies: Dies | null) {
    currentDies = dies;
  }

</script>

<div>
  {#if currentDies === null}
    <Diei diei={diei} refetch={fetchAllDies} handleClick={handleClick}/>
  {/if}
  {#if currentDies !== null}
    <div class="back-button" on:click={() => handleClick(null)}>
      <span>{"<"}</span>
    </div>
    <Quartum currentDies={currentDies}/>
  {/if}
</div>


<style>
  .back-button{
    display: absolute;
    cursor: pointer;
    top: 0;
    left: 0;
    width: 200px;
    color: black;
    font-weight: bold;
    height: 50px;
    width: 50px;
    align-self: center;
    justify-content: center;
    border-radius: 50%;
    background-color: rgba(255, 255, 255, 0.3);
  }
</style>