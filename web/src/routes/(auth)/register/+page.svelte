<script lang="ts">
  import type { ActionData } from "./$types";
  import Head from "$lib/partials/head.svelte";

  export let form: ActionData;
</script>

<svelte:head>
  <Head title="Register" description="LuXX - Register" />
</svelte:head>

<div>
  {#if form?.error}
    <p>{form.error}</p>
  {/if}
  <form method="POST" class="flex flex-col">
    <label>
      Full Name
      <input name="fullname" type="text" />
    </label>
    <label>
      username
      <input name="username" type="text" />
    </label>
    <label>
      Email
      <input name="email" type="email" />
    </label>
    <label>
      Password
      <input name="password" type="password" />
    </label>
    <button type="submit">Create Account</button>
  </form>
</div>
