<script lang="ts">
   import RoleComponent from "./components/RoleComponent.svelte";
   import { RoleSearchService } from "./services/RoleSearchService";
   import RoleSearchBar from "./components/RoleSearchBar.svelte";
   import { Navbar, NavbarBrand } from "sveltestrap";  

   const roleSearchEngine = new RoleSearchService()
   let searchTerm = ""
   let roles = []

   function handleSearch(event: CustomEvent) {
      searchTerm = event.detail.searchFor;
      roleSearchEngine.handleSearch(searchTerm)
      // console.log("found so many matches: " + roleSearchEngine.filteredRoles.length)

      roles = roleSearchEngine.filteredRoles
   }
</script>

<div class="container-fluid gx-5">
   <Navbar color="light" light expand="md">
      <NavbarBrand href="/">Welcome to the DSDR!</NavbarBrand>
   </Navbar>

   <RoleSearchBar on:searchMessage={handleSearch} />
   <br />

   {#each roles as role}
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
   {:else}
      <h2 style="text-align: center;">Perform you search!</h2>
   {/each}
</div>
