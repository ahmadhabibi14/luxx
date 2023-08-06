// @ts-nocheck
const sveltePreprocess = require('svelte-preprocess');

module.exports = {
  preprocess: [
    sveltePreprocess()
  ],
  cache: false,
  compilerOptions: {
    preserveComments: true,
    css: 'external'
  },
};

