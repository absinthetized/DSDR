import { RoleCollection } from "../collections/RoleCollection";
import { Role } from "../models/role";
import { RoleSearchService } from "./RoleSearchService";

/* redirect allerts to console */
global.alert = console.log

test("empty-db-test", async () => {
   const emtpyResponse = 
      jest.fn<Promise<Array<Role>>, any>( _ => Promise.resolve([]) ) //empty response
   
   const roleC = new RoleCollection()
   roleC.search = emtpyResponse // inject mocked method
   const roleSearchEngine = new RoleSearchService(roleC)

   const data = await roleSearchEngine.handleSearch("")     
   
   expect(data).toStrictEqual([])
} )

test("sort-2-roles-with-different-number-of-matches-test", async () => {
   const role1 = new Role({})
   const role2 = new Role({})
   role1.matchedBy = ['a','b','c']
   role2.matchedBy = ['z','w']

   const twoRolesWithDifferentMatchesResponse = 
   jest.fn<Promise<Array<Role>>, any>( _ =>  Promise.resolve([role2, role1]) ) //swap wrt expected result
      
   const roleC = new RoleCollection()
   roleC.search = twoRolesWithDifferentMatchesResponse
   const roleSearchEngine = new RoleSearchService(roleC)

   const data = await roleSearchEngine.handleSearch("")     
   
   expect(data).toStrictEqual([role1, role2])
} )

test("sort-2-roles-alphabetically-test", async () => {
   const role1 = new Role({})
   const role2 = new Role({})
   role1.matchedBy = ['a','b']
   role2.matchedBy = ['z','w']

   const twoRolesWithDifferentMatchesResponse = 
   jest.fn<Promise<Array<Role>>, any>( _ =>  Promise.resolve([role2, role1]) ) //swap wrt expected result
      
   const roleC = new RoleCollection()
   roleC.search = twoRolesWithDifferentMatchesResponse
   const roleSearchEngine = new RoleSearchService(roleC)

   const data = await roleSearchEngine.handleSearch("")     
   
   expect(data).toStrictEqual([role1, role2])
} )

test("sort-2-roles-byperc-test", async () => {
   const role1 = new Role({})
   const role2 = new Role({})
   role1.matchedBy = ['a','b']
   role1.perc_match = 0.9
   role2.matchedBy = ['a','b']
   role2.perc_match = 0.2

   const twoRolesWithDifferentMatchesResponse = 
   jest.fn<Promise<Array<Role>>, any>( _ =>  Promise.resolve([role2, role1]) ) //swap wrt expected result
      
   const roleC = new RoleCollection()
   roleC.search = twoRolesWithDifferentMatchesResponse
   const roleSearchEngine = new RoleSearchService(roleC)

   const data = await roleSearchEngine.handleSearch("")     
   
   expect(data).toStrictEqual([role1, role2])
} )