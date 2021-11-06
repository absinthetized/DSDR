import { RoleCollection } from "../collections/RoleCollection";
import type { Role } from "../models/role"

export class RoleSearchService {
   public filteredRoles: Array<Role> = [];

   private rolesC: RoleCollection

   constructor() {
      this.rolesC = new RoleCollection();
   }

   async handleSearch(searchString: string) {
      // stub for server side seach - does nothing for now and passes to client side search
      let roles = await this.rolesC.search(searchString)

      // init stats
      // AMEND this.rolesC.roles.forEach(role => role.resetMatches())
      roles.forEach(role => role.resetMatches())

      // nothing to do here, just return the whole roles set
      if (searchString === "") {
         //this.filteredRoles = [...this.rolesC.roles]
         this.filteredRoles = [...roles]
         return
      }

      // get serch terms from search bar, each separated by space
      let searchTerms = searchString.split(" ")
      // sort search so that, if not all the criteria are matched we can order 
      // alphabetically by first maching term later
      searchTerms.sort() 
      
      // scan roles for each search term, this manipulates the roles values
      searchTerms.forEach(term => this.searchSingleTerm(roles, term))

      // extract only those roles whose match the search terms
      // AMEND this.filteredRoles = this.rolesC.roles.filter( role => {
      this.filteredRoles = roles.filter( role => {
         // compute percentage of permissions matching against the search term
         try {
            if (role.includedPermissions.length > 0)
               role.perc_match += role.matches / role.includedPermissions.length
         
        } catch {
           console.log(role)
        }

         return (role.perc_match > 0) ? role : null
      })

      if (this.filteredRoles.length === 0) {
         alert("No result found!")
         return
      }

      // order by matching
      this.filteredRoles.sort( (first: Role, second: Role) => {
         // sort by number of matches
         let criterium = second.matchedBy.length - first.matchedBy.length 
         if (criterium != 0)
            return criterium

         // if 2 roles matches the same number of terms sort alphabetically
         // by first term letter (we have sorted the terms so this doesn't
         // depent on user input)
         if (second.matchedBy[0][0] > first.matchedBy[0][0]) {
            criterium = -1
            return criterium
         } else if (second.matchedBy[0][0] < first.matchedBy[0][0]) {
            criterium = 1
            return criterium
         }

         // if 2 roles match the same (number of) keywork(s), sort by percentage of
         // permissions that match the terms
         criterium = second.perc_match - first.perc_match
         return criterium
      })
   }

   searchSingleTerm(roles: Array<Role>, searchTerm: string) {
      let term = new RegExp(searchTerm)
      
      // AMEND this.rolesC.roles.forEach(role => {
      roles.forEach(role => {
         //console.log(role)
         if (role.includedPermissions === undefined)
            return

         // let see if we match this search term
         let matchingPerms = role.includedPermissions.filter(perm => {
            return term.test(perm) ? true : false
         })

         //add the number of matches for this search term
         if (matchingPerms.length > 0) {
            role.matches += matchingPerms.length
            role.matchedBy.push(searchTerm)
         }
      })
   }
}