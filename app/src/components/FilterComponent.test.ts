/**
 * @jest-environment jsdom
 */

// NOTE: NEVER REMOVE THE PREVIOUS COMMENT!!!

import '@testing-library/jest-dom'
import {fireEvent, render} from '@testing-library/svelte'
import FilterComponent from './FilterComponent.svelte'

test('filterComponent-render-test', () => {
  const {getByPlaceholderText, getByText} = render(FilterComponent, {doAlpha: true, doBeta:true, doDeprec: true, doMinPerc: 0})

  fireEvent.change(getByPlaceholderText('0%') as HTMLInputElement, {target: {value: '1'}})
  
  expect(getByText('cut off matches under')).toBeInTheDocument()
  expect(getByPlaceholderText('0%')).not.toHaveClass('is-invalid')
})

test('filterComponent-validator-render-test', () => {
  const {getByPlaceholderText, getByText} = render(FilterComponent, {doAlpha: true, doBeta:true, doDeprec: true, doMinPerc: -1})
  
  fireEvent.change(getByPlaceholderText('0%') as HTMLInputElement, {target: {value: '-1'}})
  
  expect(getByText('cut off matches under')).toBeInTheDocument()
  expect(getByPlaceholderText('0%')).toHaveClass('is-invalid')
})