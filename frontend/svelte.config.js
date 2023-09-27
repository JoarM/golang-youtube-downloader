import {preprocessMeltUI} from "@melt-ui/pp";
import sequence from "svelte-sequential-preprocessor";
import sveltePreprocess from 'svelte-preprocess';
export default {
  // Consult https://github.com/sveltejs/svelte-preprocess
  // for more information about preprocessors
  preprocess: sequence([sveltePreprocess(), preprocessMeltUI()])
};
