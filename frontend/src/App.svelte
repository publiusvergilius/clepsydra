<script lang="ts">
  import { GetAllDies } from '../wailsjs/go/main/App.js'
  import Diei from "./Diei.svelte";
  let diei = [];
  function fetchAllDies () {
    GetAllDies().then(result => { 
      console.log('result', result) 
      diei = JSON.parse(result)
    })
  }
  fetchAllDies()

</script>

<div>
  <h1>Clepsydra</h1>
  <Diei diei={diei} refetch={fetchAllDies}/>
</div>