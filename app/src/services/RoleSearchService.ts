import { RoleCollection } from "../collections/RoleCollection";
import type { Role } from "../models/role"

export class RoleSearchService {
   private rolesC: RoleCollection

   constructor() {
      this.rolesC = new RoleCollection();
   }

   async handleSearch(searchString: string): Promise<Array<Role>> {
      // stub for server side search - does nothing for now and passes to client side search
      let roles = await this.rolesC.search(searchString)

      if (roles.length === 0) {
         alert("No result found!")
         return []
      }

      // order by matching
      roles.sort( (first: Role, second: Role) => {
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

      return roles
   }
}