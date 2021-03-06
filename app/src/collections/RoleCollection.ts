import { Role } from "../models/role";
import { serverUrl } from "../config";

/*
a class to retrieve and serve role models
*/
export class RoleCollection {
   private _roles: Array<Role>

   constructor() {
      this._roles = [] //init
   }

   // stub for server side search - cureently just returns all the roles as "getFromServer"
   async search(searchString: string): Promise<Array<Role>> {      
      this._roles = [] //re-init at each search

      try {
         const resp = await fetch(serverUrl + '/search?query=' + searchString)
         try {
            const data = await resp.json()
            data.map(item => this._roles.push(new Role(item)))

         } catch (err) {
            alert(err)
         }

      } catch (err) {
         alert(err)
      }

      return this._roles
   }
}