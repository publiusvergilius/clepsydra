<script lang="ts">
  import {Greet, GetAllQuarta,CreateQuartum} from '../wailsjs/go/main/App.js'

  let resultText: string = "Please enter your name below 👇"

  let quarta = []

  GetAllQuarta().then(result => quarta = JSON.parse(result))

  console.log('quarta', quarta)

  let name: string
  let titulum: string
  let pars: number
  let hora: string
  let dies_id: number

  function greet(): void {
    console.log('all quarta', quarta) 
    Greet(name).then(result => resultText = result)
  }

  function createQuartum(): void {
    const jsonStr = JSON.stringify({titulum, pars, hora, dies_id})
    console.log(jsonStr)
    CreateQuartum(jsonStr).then(result => console.log(result))
  }
</script>

<main>
  <div>
    <h2>Deveres</h2>
   {#each quarta as quartum}
    <li>{quartum?.titulum}</li>
  {/each}
  </div>
  <hr/>
    
  <div class="box" id="box">
      <label for="titulum">Titulum:</label>
      <input autocomplete="off" bind:value={titulum} class="input" id="name" type="text"/>
      <label for="pars">Pars:</label>
      <input autocomplete="off" bind:value={pars} class="input" id="name" type="number"/>
      <label for="hora">Hora:</label>
      <input autocomplete="off" bind:value={hora} class="input" id="name" type="text"/>
      <label for="dies_id">Dies_id:</label>
      <input autocomplete="off" bind:value={dies_id} class="input" id="name" type="number"/>
    </div>
    <button class="btn" on:click={() => {greet(); createQuartum();}}>Criar</button>
</main>

<style>

  .box {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    margin: 1.5rem auto;
  }
  
</style>
