<script lang="ts">
   import {
      Badge,
      Card,
      CardBody,
      CardFooter,
      CardHeader,
      CardSubtitle,
      CardText,
      CardTitle,
      Button,
      Collapse,
   } from "sveltestrap";

   export let title: string;
   export let description: string;
   export let name: string;
   export let stage: string;
   export let includedPermissions: Array<string> = [];
   export let searchedBy: Array<string> = [];
   export let matches: number;
   export let id: number;
   export let perc_matches: number;

   let level: string = "";

   if (stage === "GA") {
      level = "info";
   } else if (stage === "BETA") {
      level = "warning";
   } else if (stage === "ALPHA") {
      level = "danger";
   } else if (stage === "DEPRECATED") {
      level = "secondary";
   } else {
      level = "dark";
   }

   function permMatchesSearch(perm: string) {
      for (let i = 0; i < searchedBy.length; i++) {
         let term = new RegExp(searchedBy[i]);

         if (term.test(perm)) {
            return true;
         }  
      } 

      return false
   }
</script>

<Card class="mb-3">
   <CardHeader>
      <div class="row" style="vertical-align: middle;">
         <div class="col-2">
            <h2>
               <Badge body color={level} style="text-align: center;"
                  ><strong>{stage}</strong></Badge
               >
            </h2>
         </div>
         <div class="col" style="text-align: center;">
            <CardTitle>{title}: {name}</CardTitle>
         </div>
         <div class="col-2" style="text-align: right;">
            <h2>
               <Button color={level} outline id="toggler-{id}" class="mb-3"
                  >Show permissions</Button
               >
            </h2>
         </div>
      </div>
   </CardHeader>
   <CardBody>
      <CardText>
         <Collapse toggler="#toggler-{id}">
            <CardSubtitle>Included permissions:</CardSubtitle>
            <br />
            <div class="row">
               {#each includedPermissions as perm}
                  {#if permMatchesSearch(perm)}
                     <div
                        class="col-3"
                        style="border-color: black; border: 1px; border-radius: 0.5em; border-style: solid;"
                     >
                        <strong>{perm}</strong>
                     </div>
                  {:else}
                     <div class="col-3">{perm}</div>
                  {/if}
               {/each}
            </div>
         </Collapse>
      </CardText>
   </CardBody>
   <CardFooter>
      <div class="row">
         <div class="col">
            <strong>Description:</strong>
            {description}
         </div>
         <div class="col-3" style="text-align: right;">
            {#if matches > 0}
               matches <strong>{searchedBy}</strong> term(s) against <strong>{matches}</strong> of {includedPermissions.length}
               perms (<strong>{perc_matches}%</strong>)
            {/if}
         </div>
      </div>
   </CardFooter>
</Card>
