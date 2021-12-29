import { RoleCollection } from "./RoleCollection"

/* redirect allerts to console */
global.alert = console.log

/* init mocks */
const mockedFetch = jest.fn() as jest.MockedFunction<typeof fetch>
global.fetch = mockedFetch

test('failing-fetch-test', async () => {
   const roleC = new RoleCollection()
   mockedFetch.mockRejectedValueOnce('rejection massage correctly catched')
   
   const data = await roleC.search('')
   expect(data).toStrictEqual([])
})