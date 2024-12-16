<script>
	import * as m from '$lib/paraglide/messages.js';
  import * as api from '$lib/api';
	let { data } = $props();

  const onCreateMember = (event) => {
    api.createDrawMember(data.draw.uuid, { name: event.target.name.value })
      .then(() => {
        event.target.reset()
        event.target.name.focus()
        api.getDraw(fetch, data.draw.uuid)
          .then(draw => {
            data = { draw }
          })
      })
  }

  const onRemoveMember = (event) => {
    api.deleteMember(event.target.uuid.value)
      .then(() => {
        api.getDraw(data.draw.uuid)
          .then(draw => {
            data = { draw }
          })
      })
  }

  const onPatchConfig = (event) => {
    api.patchDrawConfig(data.draw.uuid, { config: {
      chained: event.target.chained.checked
    }})
      .then(() => {
        api.getDraw(data.draw.uuid)
          .then(draw => {
            data = { draw }
          })
      })
  }
</script>

<h1 class="text-3xl">{m.draw_create()}</h1>
<h3 class="text-2xl">{m.draw_config()}</h3>

<form onsubmit={onPatchConfig}>
  <label id="chained">
    {m.draw_config_chained()}
    <input type="checkbox" name="chained" checked={data.draw.config.chained} />
  </label>
  <button type="submit">{m.form_submit()}</button>
</form>

<h3 class="text-2xl">{m.member_title()}</h3>
<form onsubmit={onCreateMember}>
  <input type="hidden" name="uuid" value={data.draw.uuid} />
  <input type="text" name="name" placeholder={m.member_name()} required />
  <button type="submit">{m.member_create()}</button>
</form>

<hr class="m-2">

<ul class="space-y-2">
  {#each data.draw.members as member}
    <li class="flex items-center gap-2">
      {member.name}
      <form onsubmit={onRemoveMember}>
        <input type="hidden" name="uuid" value={member.uuid} />
        <button type="submit">{m.member_delete()}</button>
      </form>
    </li>
  {/each}
</ul>
