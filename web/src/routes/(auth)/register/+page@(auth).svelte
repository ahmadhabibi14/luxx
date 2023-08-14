<script lang="ts">
  import { fade } from "svelte/transition";
  import Head from "$lib/partials/head.svelte";

  import PopupSuccess from "$lib/components/popup_success.svelte";

  interface registerIn {
    email: string;
    password: string;
    username: string;
    fullname: string;
  };
  interface registerOut {
    token: string;
    user_id: string;
    message: string;
  };
  interface registerError {
    message: string;
  };

  let payload: registerIn = {
    email: "",
    password: "",
    username: "",
    fullname: "",
  };
  let response: registerOut = {
    token: "",
    user_id: "",
    message: ""
  };
  let errorResp: registerError = {
    message: ""
  };
  
  $: uncaughtError = "";
  $: isLoading = false;
  $: isError = false;
  $: isSuccess = false;
  
  async function handleRegister( e: Event | SubmitEvent ): Promise<void> {
    isLoading = true;
    if (payload.email === "" || payload.fullname === "" || payload.username === "" || payload.password === "") {
      uncaughtError = "Please fill in all fields";
      isLoading = false;
      return;
    }
    const requestBody = JSON.stringify(payload);
    const actionURL: string = "http://localhost:1414/api/auth/register";
    try {
      const resp: Response = await fetch(actionURL, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: requestBody,
      });

      if (resp.ok) {
        const creds: Promise<any> = await resp.json();
        response = JSON.parse(await creds);
        isSuccess = true;
        setTimeout((): void => {
          isSuccess = false;
        }, 1500);
        window.location.href = "/"
        isLoading = false;
        return;
      } else {
        isError = true as boolean
        const errorData: Promise<any> = await resp.json();
        errorResp = JSON.parse(await errorData);
        isLoading = false;
        return;
      }
    } catch (error) {
      uncaughtError = error as string;
      isLoading = false;
      return;
    }
  }
</script>

<svelte:head>
  <Head
    title="LuXX - Register"
    description="Join to LuXX"
  />
</svelte:head>

<div class="min-h-screen w-full flex flex-row text-slate-900">
  {#if isSuccess}
    <div transition:fade class="delay-700 absolute top-0 right-0 h-full w-[55%] bg-white flex justify-center items-center">
      <PopupSuccess message={response.message}/>
    </div>
  {/if}
  <div class="basis-[45%] flex justify-center items-center">
    <img src="/icons/main/girl_stand_laptop.svg" alt="" class="w-[500px] h-auto">
  </div>
  <div class="basis-[55%] bg-white flex justify-center items-center">
    <section class="flex flex-col h-fit w-[55%]">
      <h1 class="font-bold text-center text-4xl">Create Your Account</h1>
      <p class="text-center text-sm text-slate-700 mt-1.5">Welcome! Please enter your details</p>
      <form
        on:submit|preventDefault={handleRegister}
        class="flex flex-col space-y-3 mt-6">
        <label for="fullname">
          <input
            class="w-full border-slate-400/80 border rounded-xl bg-transparent focus:border-slate-950 caret-slate-950 py-2 px-4 outline-1 focus:outline-slate-950"
            type="text"
            id="fullname"
            bind:value={payload.fullname}
            placeholder="Full Name"
          />
        </label>
        <label for="username">
          <input
            class="w-full border-slate-400/80 border rounded-xl bg-transparent focus:border-slate-950 caret-slate-950 py-2 px-4 outline-1 focus:outline-slate-950"
            type="text"
            id="username"
            bind:value={payload.username}
            placeholder="UserName"
          />
        </label>
        <label for="email">
          <input
            class="w-full border-slate-400/80 border rounded-xl bg-transparent focus:border-slate-950 caret-slate-950 py-2 px-4 outline-1 focus:outline-slate-950"
            type="text"
            id="email"
            bind:value={payload.email}
            placeholder="Your Email"
          />
        </label>
        <label for="password">
          <input
            class="w-full border-slate-400/80 border rounded-xl bg-transparent focus:border-slate-950 caret-slate-950 py-2 px-4 outline-1 focus:outline-slate-950"
            type="password"
            id="password"
            bind:value={payload.password}
            placeholder="Password"
          />
        </label>
        <button
          class="w-full flex justify-center items-center cursor-pointer py-2.5 bg-slate-950 hover:bg-slate-900 hover:shadow-md text-slate-100 rounded-xl"
          type="submit"
          >
          {#if isLoading}
            <svg class="w-[20px] h-auto animate-spin" fill="none" viewBox="0 0 24 24">
              <path d="M2.19995 14.02C3.12995 18.58 7.15995 22 12 22C16.82 22 20.8399 18.59 21.7899 14.05" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
              <path d="M21.81 10.06C20.91 5.46 16.86 2 12 2C7.16995 2 3.13995 5.43001 2.19995 9.98001" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
              <path d="M12 13.5C12.8284 13.5 13.5 12.8284 13.5 12C13.5 11.1716 12.8284 10.5 12 10.5C11.1716 10.5 10.5 11.1716 10.5 12C10.5 12.8284 11.1716 13.5 12 13.5Z" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>
          {:else}        
            <p>Create Account</p>
          {/if}
        </button>
      </form>
      <div class="flex flex-row space-x-2 items-center my-4">
        <span class="grow h-px bg-slate-400"></span>
        <p class="w-fit text-slate-700 text-sm">OR</p>
        <span class="grow h-px bg-slate-400"></span>
      </div>
      <a href="/" class="py-2.5 flex flex-row space-x-2 justify-center items-center bg-slate-100 border-slate-200 border hover:bg-slate-200/80 rounded-xl">
        <img src="/icons/brands/google.svg" alt="" class="w-[25px] h-auto"/>
        <p>Continue with Google</p>
      </a>
      <a href="/" class="mt-3 py-2.5 flex flex-row space-x-2 justify-center items-center bg-slate-100 border-slate-200 border hover:bg-slate-200/80 rounded-xl">
        <img src="/icons/brands/github.svg" alt="" class="w-[25px] h-auto"/>
        <p>Continue with Github</p>
      </a>
      <div class="mt-5 text-center text-sm">
        <p>ALready have account? <a href="/login" class="text-sky-600 hover:underline">Login</a></p>
      </div>
    </section>
  </div>
</div>
