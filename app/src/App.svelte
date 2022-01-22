<script lang='ts'>
import { Navbar, NavbarBrand, Nav, NavLink } from "sveltestrap"
import Router from 'svelte-spa-router'
import {wrap} from 'svelte-spa-router/wrap'

const routes = {
    '/': wrap({ asyncComponent: () => import("pages/Home.svelte") }),
    '/search': wrap({ asyncComponent: () => import("pages/MainPage.svelte") }),

    // Catch-all
    '*': wrap({ asyncComponent: () => import("pages/NotFound.svelte") }),
}
</script>

<div class="container-fluid gx-5">
   <Navbar color="light" light expand="md">
      <NavbarBrand href="/">Welcome to the DSDR!</NavbarBrand>
      <Nav>
         <NavLink href="#/">Home</NavLink>
         <NavLink href="#/search">Search</NavLink>
         <NavLink disabled href="#/reindex">Reindex DB</NavLink>
         <NavLink disabled href="#/logout">Logout</NavLink>
       </Nav>
   </Navbar>
   <Router {routes}/>
</div>