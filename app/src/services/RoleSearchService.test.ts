import { RoleSearchService } from "./RoleSearchService";

/* redirect allerts to console */
global.alert = console.log

test("hello-world-test", async () => {
   const roleSearchEngine = new RoleSearchService()
   const data = await roleSearchEngine.handleSearch("pippo")   
   expect(data).toStrictEqual([])
} )