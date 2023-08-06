<script lang="ts">
  let name: string = "Habi"
</script>

<main class="w-full h-screen flex justify-center bg-slate-100">
  <div class="w-96 h-fit bg-white border shadow-lg rounded-2xl p-6 mt-44 flex flex-col">
    <a
      class="text-white w-full py-3 rounded-lg bg-slate-900 hover:bg-slate-800 text-center"
      href="/api/oauth/github/login"
    >
      Login with Github
    </a>
  </div>
</main>