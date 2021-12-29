/** @type {import('ts-jest/dist/types').InitialOptionsTsJest} */
module.exports = {
  preset: 'ts-jest',
  testEnvironment: 'node',
  // for coverage enforcing
  coverageThreshold: {
    global: {
      lines: 75
    }
  },
  // for svelte component test via testing library
  transform: {
    "^.+\\.svelte$": "svelte-jester"
  }
};
