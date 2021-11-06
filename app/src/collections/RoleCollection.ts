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

   async getFromServer(): Promise<Array<Role>> {
      try {
         const resp = await fetch(serverUrl + '/roles')
         try {
            const data = await resp.json()
            data.map(item => this._roles.push(new Role(item)))

         } catch (err) {
            alert(err)
         }

      } catch (err) {
         alert(err)
      }

      console.log(this._roles[0])
      return this._roles
   }

   get roles() { return this._roles }
}