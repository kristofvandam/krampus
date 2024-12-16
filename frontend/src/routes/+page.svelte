<script lang="ts">
	import * as m from '$lib/paraglide/messages.js'
  import { goto } from '$app/navigation'
  import { i18n } from '$lib/i18n'
  import * as api from '$lib/api'
  import { page } from '$app/stores'

  async function newDraw() {
    const draw = await api.createDraw()
    goto(i18n.resolveRoute(`/draw/${draw.uuid}`))
  }
</script>

<div class="@container">
  <header>
    <h2>{m.welcome()}</h2>
  </header>

  <button on:click={() => newDraw()}>{m.draw_create()}</button>
</div>
