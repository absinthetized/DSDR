/**
 * @jest-environment jsdom
 */

// NOTE: NEVER REMOVE THE PREVIOUS COMMENT!!!

import '@testing-library/jest-dom'
import {render} from '@testing-library/svelte'
import RoleSearchBarComponent from './RoleSearchBarComponent.svelte'

test('searchbar-render-test', () => {
  const {getByText} = render(RoleSearchBarComponent)
  const mock = jest.fn()
  //component.$on('searchMessage', mock)
  
  expect(getByText('Enter a partial permission name + ENTER')).toBeVisible()
})
