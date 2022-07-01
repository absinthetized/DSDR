/** @type {import('ts-jest/dist/types').InitialOptionsTsJest} */
module.exports = {
  coverageProvider: 'v8',
  preset: 'ts-jest',
  testEnvironment: 'node',
  // for coverage enforcing
  coverageThreshold: {
    global: {
      lines: 75
    }
  },
  // for svelte component test via testing library
  "transform": {
    "^.+\\.svelte$": [
      "svelte-jester",
      {
        "preprocess": "./svelte.config.mjs" //must use mjs extension or node explodes
      }
    ],
    "^.+\\.ts$": "ts-jest"
  }
};
