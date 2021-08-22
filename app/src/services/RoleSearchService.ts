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
   }

   searchSingleTerm(searchTerm: string): Role[] {
      let term = new RegExp(searchTerm)
      let filteredRoles = []

      filteredRoles = this.rolesC.roles.filter(role => {
         //console.log(role)
         if (role.includedPermissions === undefined)
            return null

         let matchingPerms = role.includedPermissions.filter(perm => {
            return term.test(perm) ? true : false
         })

         role.matches += matchingPerms.length
         if (role.matches > 0)
            role.matchedBy.push(searchTerm)

         return (matchingPerms.length > 0) ? role : null
      })

      if (filteredRoles.length === 0) {
         alert("No result found!")
      }

      return filteredRoles
   }
}