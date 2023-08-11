<script lang="ts">
  type JSONvalue = 
    | string
    | number
    | boolean

  interface JSONObject {
    [x: string]: JSONvalue;
  }
  interface registerIn {
    email: string;
    password: string;
    username: string;
    fullname: string;
  }
  let payload: registerIn = {
    email: '',
    password: '',
    username: '',
    fullname: '',
  };
  let response: JSONObject = {};
  let errorResp: JSONObject = {};
  
  async function handleRegister( e: Event | SubmitEvent ): Promise<void> {
    const requestBody = JSON.stringify(payload);
    const targetForm = e.target as HTMLFormElement;
    const actionURL: string = targetForm.action;
    try {
      const resp: Response = await fetch(actionURL, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: requestBody,
      });

      if (resp.ok) {
        const creds: Promise<any> = await resp.json();
        response = JSON.parse(await creds);
        console.log(response)
      } else {
        const errorData: Promise<any> = await resp.json();
        errorResp = JSON.parse(await errorData);
        console.error(errorData)
      }
    } catch (error) {
      console.error(error);
    }
  }
</script>

<div class="min-h-screen w-full flex flex-row text-slate-900">
  <div class="basis-[45%] flex justify-center items-center">
    <img src="/assets/icons/main/girl_stand_laptop.svg" alt="" class="w-[500px] h-auto">
  </div>
  <div class="basis-[55%] bg-white flex justify-center items-center">
    <section class="flex flex-col h-fit w-[55%]">
      <h1 class="font-bold text-center text-4xl">Create Your Account</h1>
      <p class="text-center text-sm text-slate-700 mt-1.5">Welcome! Please enter your details</p>
      <form
        action="/api/auth/register"
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
        <button class="w-full py-2.5 bg-slate-950 hover:bg-slate-900 hover:shadow-md text-slate-100 rounded-xl" type="submit">
          Create Account
        </button>
      </form>
      <div class="flex flex-row space-x-2 items-center my-4">
        <span class="grow h-px bg-slate-400"></span>
        <p class="w-fit text-slate-700 text-sm">OR</p>
        <span class="grow h-px bg-slate-400"></span>
      </div>
      <a href="/" class="py-2.5 flex flex-row space-x-2 justify-center items-center bg-slate-100 border-slate-200 border hover:bg-slate-200/80 rounded-xl">
        <img src="/assets/icons/brands/google.svg" alt="" class="w-[25px] h-auto"/>
        <p>Continue with Google</p>
      </a>
      <a href="/" class="mt-3 py-2.5 flex flex-row space-x-2 justify-center items-center bg-slate-100 border-slate-200 border hover:bg-slate-200/80 rounded-xl">
        <img src="/assets/icons/brands/github.svg" alt="" class="w-[25px] h-auto"/>
        <p>Continue with Github</p>
      </a>
      <div class="mt-5 text-center text-sm">
        <p>ALready have account? <a href="/login" class="text-sky-600 hover:underline">Login</a></p>
      </div>
    </section>
  </div>
</div>
