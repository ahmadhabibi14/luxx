<script>
  import Head from "$lib/partials/head.svelte";
  import { user } from "../stores";
  export let data;

  $: User = data?.authedUser;
  $: {
    const newUser = data?.authedUser;
    user.set(newUser);
  }
</script>

<svelte:head>
  <Head title="LuXX" description="Social app" />
</svelte:head>

<section>
  <h1>Welcome {User?.fullname ?? ""}</h1>
  <p>{User?.email ?? ""}</p>
  <p>{User?.username ?? ""}</p>
  <p>{User?.join_at ?? ""}</p>
</section>
