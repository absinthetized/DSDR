/**
 * @jest-environment jsdom
 */

// NOTE: NEVER REMOVE THE PREVIOUS COMMENT!!!

import '@testing-library/jest-dom'
import {render} from '@testing-library/svelte'
import RoleComponent from './RoleComponent.svelte'

test('role-ga-render-test', () => {
  const level='GA'
  const {getByText} = render(RoleComponent, {name: 'pippo', description: 'pippo',
   title: 'pippo', stage: level, includedPermissions: ['super','hero'], searchedBy: ['sup'],
   matches: 1, id: 666, perc_matches: 0.15 })
  
  expect(getByText(level)).toBeVisible()
})

test('role-beta-render-test', () => {
   const level='BETA'
   const {getByText} = render(RoleComponent, {name: 'pippo', description: 'pippo',
    title: 'pippo', stage: level, includedPermissions: ['super','hero'], searchedBy: ['sup'],
    matches: 1, id: 666, perc_matches: 0.15 })
   
   expect(getByText(level)).toBeVisible()
 })

 test('role-beta-render-test', () => {
   const level='ALPHA'
   const {getByText} = render(RoleComponent, {name: 'pippo', description: 'pippo',
    title: 'pippo', stage: level, includedPermissions: ['super','hero'], searchedBy: ['sup'],
    matches: 1, id: 666, perc_matches: 0.15 })
   
   expect(getByText(level)).toBeVisible()
 })