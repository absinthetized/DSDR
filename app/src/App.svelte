<script lang="ts">
   import RoleComponent from "./components/RoleComponent.svelte";
   import { RoleSearchService } from "./services/RoleSearchService";
   import RoleSearchBar from "./components/RoleSearchBarComponent.svelte";
   import FilterComponent from "./components/FilterComponent.svelte";
   import { Navbar, NavbarBrand } from "sveltestrap";  
   import { FilterService } from "./services/FilterService";

   const roleSearchEngine = new RoleSearchService()
   const filter = new FilterService()
   let searchTerm = ""
   let roles = []

   async function handleSearch(event: CustomEvent) {
      searchTerm = event.detail.searchFor;
      roles = [] //forces reinit of array and fixes repaint of badges for GA/BETA/...
      roles = await roleSearchEngine.handleSearch(searchTerm)
      // console.log("found so many matches: " + roleSearchEngine.filteredRoles.length)
   }
</script>

<div class="container-fluid gx-5">
   <Navbar color="light" light expand="md">
      <NavbarBrand href="/">Welcome to the DSDR!</NavbarBrand>
   </Navbar>

   <RoleSearchBar on:searchMessage={handleSearch} />
   <FilterComponent 
      bind:doAlpha={filter.doAlpha}
      bind:doBeta={filter.doBeta}
      bind:doDeprec={filter.doDeprec}
      bind:doMinPerc={filter.doMinPerc}
   ></FilterComponent>
   <br />

   {#each roles as role}
      {#if (
            (role.stage === 'ALPHA' && filter.doAlpha) || 
            (role.stage === 'BETA' && filter.doBeta) ||
            (role.stage === 'DEPRECATED' && filter.doDeprec) ||
            (role.stage === 'GA')
           ) && (role.perc_match*100) >= filter.doMinPerc
      }
      <RoleComponent
         name={role.name}
         description={role.description}
         title={role.title}
         stage={role.stage}
         includedPermissions={role.includedPermissions}
         searchedBy={role.matchedBy}
         matches = {role.matches}
         id = {role.id}
         perc_matches ={(role.perc_match * 100).toFixed(2)}
      />
      {/if}
   {:else}
      <h2 style="text-align: center;">Perform you search!</h2>
   {/each}
</div>
