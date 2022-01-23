/**
 * @jest-environment jsdom
 */

import '@testing-library/jest-dom'
import { render, fireEvent } from '@testing-library/svelte'
import App from './App.svelte'

test('home-page-render-test', async () => {
   const { getByText } = render(App)

   expect(getByText('Welcome aboard')).toBeVisible()
   expect(getByText('Welcome to the DSDR!')).toBeVisible()
})
