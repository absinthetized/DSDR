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
      //console.log('search term: ' + this._searchTerm);
      this.rolesC.roles.forEach(role => role.resetMatches())

      if (searchString === "") {
         this.filteredRoles = [...this.rolesC.roles]

         return
      }

      let searchTerms = searchString.split(" ")

      /* accumulate all results */

      this.filteredRoles = searchTerms.reduce((accumulator, term): Role[] => {
         // console.log("accumulator is:")
         // console.log(accumulator)
         // console.log("seach term is: " + term)

         return accumulator.concat(this.searchSingleTerm(term))
      }, [])

      /* purge doubles with a Set */

      let tmpSet = new Set(this.filteredRoles)
      this.filteredRoles = Array.from(tmpSet.values())

      // order by matching percentage
      this.filteredRoles.sort( (first: Role, second: Role) => { return second.perc_match - first.perc_match })
   }

   searchSingleTerm(searchTerm: string): Role[] {
      let term = new RegExp(searchTerm)
      let filteredRoles = []

      filteredRoles = this.rolesC.roles.filter(role => {
         //console.log(role)
         if (role.includedPermissions === undefined)
            return null

         // let see if we match this search term
         let matchingPerms = role.includedPermissions.filter(perm => {
            return term.test(perm) ? true : false
         })

         //add the number of matches for this search term
         role.matches += matchingPerms.length
         if (role.matches > 0)
            role.matchedBy.push(searchTerm)

         // compute percentage of permissions matching against the search term
         if (role.includedPermissions.length > 0)
            role.perc_match = role.matches / role.includedPermissions.length * 100
         else
            role.perc_match = 0

         // if we match no permission skip this role
         return (matchingPerms.length > 0) ? role : null
      })

      if (filteredRoles.length === 0) {
         alert("No result found!")
      }

      return filteredRoles
   }
}