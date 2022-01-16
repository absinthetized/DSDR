/**
 * @jest-environment jsdom
 */

// NOTE: NEVER REMOVE THE PREVIOUS COMMENT!!!

import '@testing-library/jest-dom'
import { render, fireEvent } from '@testing-library/svelte'
import RoleComponent from './RoleComponent.svelte'

test('role-ga-render-test', async () => {
   const level = 'GA'
   const { getByText } = render(RoleComponent, {
      name: 'pippo', description: 'pippo',
      title: 'pippo', stage: level, includedPermissions: ['super', 'hero'], searchedBy: ['sup'],
      matches: 1, id: 666, perc_matches: 0.15
   })

   const button = getByText('Show permissions')
   await fireEvent.click(button)

   expect(getByText('super')).toBeVisible()
   expect(getByText('GA')).toBeVisible()
})
