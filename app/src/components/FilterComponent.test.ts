/**
 * @jest-environment jsdom
 */

// NOTE: NEVER REMOVE THE PREVIOUS COMMENT!!!

import '@testing-library/jest-dom'
import {render} from '@testing-library/svelte'
import FilterComponent from './FilterComponent.svelte'

test('basic-filter-comp-test', () => {
  const {getByText} = render(FilterComponent, {doAlpha: true, doBeta:true, doDeprec: true, doMinPerc: 0})
  expect(getByText('cut off matches under')).toBeInTheDocument()
})
