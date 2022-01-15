<script lang="ts">
   import { InputGroup, InputGroupText, Input } from "sveltestrap";
   import { createEventDispatcher } from "svelte";

   let searchString: string = "";
   const dispatch = createEventDispatcher();

   function dispatchSearch() {
      dispatch('searchMessage', {'searchFor': searchString});
   }

   document.onkeydown = function (e) {
      if (e.key === "Enter") {
         console.log('pippo')
         dispatchSearch();
      }
   };
</script>

<div>
   <InputGroup>
      <InputGroupText>Enter a partial permission name + ENTER</InputGroupText>
      <Input
         placeholder="e.g.: 'loggi' will search for all permissions matching the 'loggi' substring. no regex for now sorry!"
         bind:value={searchString}
         style="margin-bottom: 0px;"
      />
   </InputGroup>
</div>
