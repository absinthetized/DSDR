import { RoleCollection } from "../collections/RoleCollection";
import type { Role } from "../models/role"

export class RoleSearchService {
   public filteredRoles: Array<Role> = [];

   private rolesC: RoleCollection

   constructor() {
      this.rolesC = new RoleCollection();
      this.rolesC.getFromServer()
   }

   handleSearch(searchString: string) {
      // init stats
      this.rolesC.roles.forEach(role => role.resetMatches())

      // nothing to do here, just return the whole roles set
      if (searchString === "") {
         this.filteredRoles = [...this.rolesC.roles]
         return
      }

      // get serch terms from search bar, each separated by space
      let searchTerms = searchString.split(" ")

      // scan roles for each search term, this manipulates the roles values
      searchTerms.forEach(term => this.searchSingleTerm(term))

      // extract only those roles whose match the search terms
      this.filteredRoles = this.rolesC.roles.filter( role => {
         // compute percentage of permissions matching against the search term
         if (role.includedPermissions.length > 0)
            role.perc_match += role.matches / role.includedPermissions.length

         return (role.perc_match > 0) ? role : null
      })

      if (this.filteredRoles.length === 0) {
         alert("No result found!")
         return
      }

      // order by matching
      // here is a small trick: 
      // perc is always 0<=x<=1
      // matchedBy size is always >=1
      // by summing up the two I get ordering by number of matches first,
      // by match percentage after.
      //
      // don't be fooled by possibly undefined values:
      // 3.0 here is aways 2 matches with 100%
      // as one can't have 3 matches with 0% match
      this.filteredRoles.sort( (first: Role, second: Role) => { 
         return (second.perc_match + second.matchedBy.length) - (first.perc_match + first.matchedBy.length)
      })
   }

   searchSingleTerm(searchTerm: string) {
      let term = new RegExp(searchTerm)
      
      this.rolesC.roles.forEach(role => {
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