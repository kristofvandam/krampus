<script lang="ts">
	import { i18n } from '$lib/i18n';
  import Nav from '$lib/Nav.svelte';
	import { ParaglideJS } from '@inlang/paraglide-sveltekit';
	import '../app.css';
	let { children } = $props();
</script>

<ParaglideJS {i18n}>
  <Nav />
	{@render children()}
</ParaglideJS>
