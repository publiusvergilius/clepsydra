
<script lang="ts">
  import { Greet, CreateQuartum, GetQuartaByDies } from '../wailsjs/go/main/App.js'
  let isOpenModal = false
  let resultText: string = "Please enter your name below 👇"
  let quarta = [];

  type Dies = {
    id: number,
    date: string
  }
  
  export let currentDies: Dies = {id: 0, date: ""};

 const res =  GetQuartaByDies(currentDies?.id).then(result => quarta = JSON.parse(result))
 console.log(res)

  let name: string
  let titulum: string
  let pars: number
  let hora: string

  function greet(): void {
    console.log('all quarta', quarta) 
    Greet(name).then(result => resultText = result)
  }



  function createQuartum(): void {
    const jsonStr = JSON.stringify({titulum, pars, hora, dies_id: currentDies?.id})
    CreateQuartum(jsonStr).then(result => console.log(result))
  }
</script>

<main>
  <div class="layout">
    <div class="box">
      <h1 class="primary_text">Mane</h1>
      <div class="box-list-container">
        <div class="box-list">
          {#each quarta as quartum}
            {#if quartum?.pars === 1}
            <li>
              <h3>{quartum?.titulum}</h3>
              <h4 class="hora">{quartum?.hora}</h4>
            </li>
            {/if}
          {/each}
        </div>
      </div>
    </div>
    <div class="box">
      <h1 class="primary_text">Meridies</h1>
        <div class="box-list-container">
          <div class="box-list">
          {#each quarta as quartum}
            {#if quartum?.pars === 2}
            <li>
              <h3>{quartum?.titulum}</h3>
              <h4>{quartum?.hora}</h4>
            </li>
            {/if}
          {/each}
        </div>
      </div>

    </div>
    <div class="box">
      <h1 class="primary_text">Vesper</h1>
      <div class="box-list-container">
        <div class="box-list">
          {#each quarta as quartum}
            {#if quartum?.pars === 3}
            <li>
              <h3>{quartum?.titulum}</h3>
              <h4>{quartum?.hora}</h4>
            </li>
            {/if}
          {/each}
        </div>
      </div>

    </div>
    <div class="box">
      <h1 class="primary_text">Nox</h1>
      <div class="box-list-container">
        <div class="box-list">
          {#each quarta as quartum}
            {#if quartum?.pars === 4}
            <li>
              <h3>{quartum?.titulum}</h3>
              <h4>{quartum?.hora}</h4>
            </li>
            {/if}
          {/each}
        </div>
      </div>
    </div>
  </div>
  <div>

  <button class="add-button" on:click={() => (isOpenModal = !isOpenModal)}>+</button>

  {#if isOpenModal}
  <div class="modal" >
  <button class="close" on:click={() => (isOpenModal = false)}>x</button>
  <form class="form">
      <label for="titulum">Titulum:</label>
      <input autocomplete="off" bind:value={titulum} class="input" id="name" type="text"/>
      <label for="pars">Pars:</label>
      <input autocomplete="off" bind:value={pars} class="input" id="name" type="number"/>
      <label for="hora">Hora:</label>
      <input autocomplete="off" bind:value={hora} class="input" id="name" type="text"/>

      <button on:click={() => {greet(); createQuartum();}}>Criar</button>
    </form>
  </div>
  {/if}
  <hr/>
</main>

<style>

  .box {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    margin: 1.5rem auto;
  }

	.layout {
  display: flex;
  flex-wrap: wrap;
  height: 100vh;
  width: 100vw;
  max-height: 100vh;
  max-width: 100vw;
  flex-direction: row;
}

.box {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 2rem;
  color: white;
  background-size: cover;
  background-position: center;
}

.box-list-container {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
}

.box:nth-child(1) {
  background-image: url("./assets/imgs/mane-field-with-rising-sun-1889-vincent-van-gogh.jpg");
}
.box:nth-child(2) {
  background-image: url("./assets/imgs/tarde-Vincent-van-gogh-evening-landscape-with-rising-moon.jpeg");
}
.box:nth-child(3) {
  background-image: url("./assets/imgs/nox-paris.jpg");
}
.box:nth-child(4) {
  background-image: url("./assets/imgs/nox_starry-night.webp");
}

/* Responsive: Stack vertically when screen width is small */
@media (max-width: 768px) {
  .layout {
    flex-direction: column;
    display: none;
    place-content: center;
  }

  .box {
    height: 25vh !important; /* Each takes 25% of the screen height */
  }
}

.primary_text {
  font-weight: bold;
  font-family: Arial, Helvetica, sans-serif;
  font-size: 3rem;
}

.add-button {
  position: fixed;
  top: 1rem;
  right: 1rem;
  background-color: rgba(255, 255, 255, 0.5);
  color: black;
  padding: 1rem 1.5rem;
  font-size: 1.5rem;
  font-weight: bold;
  border: none;
  cursor: pointer;
  border-radius: 20px;
}

.modal {
  display: flex;
  position: absolute;
  z-index: 999;
  align-items: center;
  justify-content: center;
  left: 0;
  top: 0;
  bottom: 0;
  right: 0;
  min-width: 300px;
  min-height: 300px;
  width: 100%;
  height: 100%;
  overflow: auto;
  background-color: rgb(0, 0, 0);
  background-color: rgba(0, 0, 0, 0.4);
}

.form {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background-color: white;
  padding: 2rem;
  border-radius: 20px;
  gap: 1rem;
  width: 500px;
  color: black;
}

.close {
  position: relative;
  top: -210px;
  right: -630px;
  background-color: rgba(255, 255, 255, 0.5);
  color: black;
  padding: 1rem 1.5rem;
  font-size: 1.5rem;
  font-weight: bold;
  border: none;
  cursor: pointer;
  border-radius: 20px;
}

</style>
