<script lang="ts">
	import { i18n } from '$lib/i18n';
	import { ParaglideJS } from '@inlang/paraglide-sveltekit';
	import '../app.css';
	let { children } = $props();

	import type { AvailableLanguageTag } from '$lib/paraglide/runtime';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import * as m from '$lib/paraglide/messages.js';

	function switchToLanguage(newLanguage: AvailableLanguageTag) {
		const canonicalPath = i18n.route($page.url.pathname);
		const localisedPath = i18n.resolveRoute(canonicalPath, newLanguage);
		goto(localisedPath);
	}

  let flagMap = {
    'en': '🇬🇧',
    'nl': '🇳🇱',
    'pt': '🇵🇹',
  }
</script>

<div class="flex items-center justify-end w-full p-8">
  <select onchange={(e) => switchToLanguage(e.target.value)} class="border-none focus:ring-0">
    {#each i18n.config.runtime.availableLanguageTags as lang}
      <option value={lang}>{flagMap[lang]} {lang}</option>
    {/each}
  </select>
</div>

<ParaglideJS {i18n}>
	{@render children()}
</ParaglideJS>
