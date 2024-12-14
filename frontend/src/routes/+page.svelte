<script lang="ts">
	import * as m from '$lib/paraglide/messages.js';
  import { goto } from '$app/navigation';
  import * as api from '$lib/api';

  async function newDraw() {
    const draw = await api.createDraw();
    goto(`/draw/${draw.uuid}`);
  }
</script>

<div class="@container">
  <header>
    <h2>{m.welcome()}</h2>
  </header>

  <button on:click={() => newDraw()}>{m.new_draw()}</button>
</div>
