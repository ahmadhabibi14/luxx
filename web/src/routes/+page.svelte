<script>
  import Head from "$lib/partials/head.svelte";
  export let data;
</script>

<svelte:head>
  <Head title="LuXX" description="Social app" />
</svelte:head>

<section>
  {#if data.user}
    <h1>Welcome {data.user.fullname}</h1>
    <p>{data.user.email}</p>
    <p>{data.user.username}</p>
    <p>{data.user.join_at}</p>
  {/if}

  <h2>Hahahaha</h2>
</section>
