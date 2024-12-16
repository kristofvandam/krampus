<script lang="ts">
	import { i18n } from '$lib/i18n';
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

<div class="flex items-center w-full p-8">
  <div class="flex-grow">
    <ul class="flex space-x-4">
      <li><a href="/" class="text-blue-500">{m.nav_home()}</a></li>
    </ul>
  </div>

  <div>
    <select
      onchange={(e) => switchToLanguage(e.target.value)}
      class="border-none focus:ring-0"
      value={i18n.config.runtime.languageTag()}
    >
      {#each i18n.config.runtime.availableLanguageTags as lang}
        <option value={lang}>
          {flagMap[lang]} {lang}
        </option>
      {/each}
    </select>
  </div>
</div>
