<script>
	import * as m from '$lib/paraglide/messages.js';
  import { enhance } from '$app/forms';
	let { data } = $props();

  let createMemberInputRef

  const onCreateMember = () => {
    createMemberInputRef.focus()
  }
</script>

<h1>{m.new_draw()}</h1>

<h2>Members</h2>

<form method="POST" action="?/create_member" use:enhance>
  <input type="hidden" name="uuid" value={data.draw.uuid} />
  <input type="text" bind:this={createMemberInputRef} name="name" placeholder="Member name" required />
  <button type="submit">Add member</button>
</form>

<ul>
  {#each data.draw.members as member}
    <li>
      {member.name}
      <form method="POST" action="?/delete_member" use:enhance>
        <input type="hidden" name="uuid" value={member.uuid} />
        <button type="submit">Remove</button>
      </form>
    </li>
  {/each}
</ul>
