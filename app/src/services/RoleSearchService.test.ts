import { RoleCollection } from "../collections/RoleCollection";
import { Role } from "../models/role";
import { RoleSearchService } from "./RoleSearchService";

/* redirect allerts to console */
global.alert = console.log

/* setup mocking interfaces */

const roleC = new RoleCollection()
const mockedSearch = jest.fn() as jest.MockedFunction<typeof roleC.search>
roleC.search = mockedSearch

/* init the search engine */
const roleSearchEngine = new RoleSearchService(roleC)

test("empty-db-test", async () => {   
   mockedSearch.mockResolvedValueOnce([]) //empty response
   const data = await roleSearchEngine.handleSearch("")     
   
   expect(data).toStrictEqual([])
} )

test("sort-2-roles-with-different-number-of-matches-test", async () => {
   const role1 = new Role({})
   const role2 = new Role({})
   role1.matchedBy = ['a','b','c']
   role2.matchedBy = ['z','w']

   mockedSearch.mockResolvedValueOnce([role2, role1]) //swap wrt expected result
   const data = await roleSearchEngine.handleSearch("")     
   
   expect(data).toStrictEqual([role1, role2])
} )

test("sort-2-roles-alphabetically-test", async () => {
   const role1 = new Role({})
   const role2 = new Role({})
   role1.matchedBy = ['a','b']
   role2.matchedBy = ['z','w']

   mockedSearch.mockResolvedValueOnce([role2, role1]) //swap wrt expected result
   const data = await roleSearchEngine.handleSearch("")     
   
   expect(data).toStrictEqual([role1, role2])
} )

test("check-2-roles-alphabetically-already-sorted-test", async () => {
   const role1 = new Role({})
   const role2 = new Role({})
   role1.matchedBy = ['a','b']
   role2.matchedBy = ['z','w']

   mockedSearch.mockResolvedValueOnce([role1, role2]) //expected result
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

   mockedSearch.mockResolvedValueOnce([role2, role1]) //swap wrt expected result
   const data = await roleSearchEngine.handleSearch("")     
   
   expect(data).toStrictEqual([role1, role2])
} )
