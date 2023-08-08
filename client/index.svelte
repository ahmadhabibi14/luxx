<script lang="ts">
  import SideMenu from "_components/SideMenu.svelte";
  import Header from "_components/Header.svelte";
  import { users } from "types/user";

  let name: string = "Habi"
</script>

<section class="max-w-[1024px] mx-auto flex">
  <SideMenu />
  <div class="ml-[324px] w-[724px]">
    <Header />
    <main id="main_content" class="w-full overflow-y-scroll mt-[190px]">
      {#each users as user}
        <p>Name: {user.name}</p>
        <p>Email: {user.email}</p>
      {/each}
      <div class="w-full h-[5000px]"></div>
      <p>main</p>
    </main>
  </div>
</section>