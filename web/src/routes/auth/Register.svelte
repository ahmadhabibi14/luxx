<script>
	import { onMount } from "svelte";
	import { navigate } from "svelte-routing";
	import Head from "../../lib/partials/head.svelte";
  import { getCookie } from "../../lib/utils/helper";

	onMount(() => {
		const auth = getCookie("auth");
		if (auth) {
			navigate("/");
		}
	});

	let errorMessage = "";
	let email, password, username, fullname;

	async function handleRegister() {
		const resp = await fetch("http://localhost:1414/api/auth/register", {
			method: "POST",
			headers: {
				"Content-Type": "application/json"
			},
			body: JSON.stringify({
				email: email,
				password: password,
				username: username,
				fullname: fullname
			})
		});
    if (resp) {
			console.log(resp)
      const auth = getCookie("auth");
      if (auth) {
        navigate("/login");
        return;
      }
    }
		return;
	}
</script>

<Head title="Register" description="LuXX - Register" />

<main class="bg-slate-100 h-screen w-full">
	<div
		class="px-5 md:px-10 flex flex-row justify-center items-center h-full w-full gap-28 text-slate-950"
	>
		<div class="md:block hidden">
			<img src="/icons/main/laptop.svg" alt="" class="w-[420px] h-auto" />
		</div>
		<div class="w-[450px] bg-white p-6 rounded-xl drop-shadow-lg flex flex-col gap-3">
			<header>
				<h1 class="font-bold text-center text-4xl">Create Your Account</h1>
				<p class="text-center text-sm text-slate-700 mt-1.5">Welcome! Please enter your details</p>
			</header>
			{#if errorMessage}
				<div
					class="w-full h-fit p-2 rounded-md flex flex-row justify-center gap-2 items-center bg-red-500 text-white"
				>
					<svg viewBox="0 0 24 24" class="w-[18px] h-auto fill-current">
						<path
							d="M12 22.75C6.07 22.75 1.25 17.93 1.25 12C1.25 6.07 6.07 1.25 12 1.25C17.93 1.25 22.75 6.07 22.75 12C22.75 17.93 17.93 22.75 12 22.75ZM12 2.75C6.9 2.75 2.75 6.9 2.75 12C2.75 17.1 6.9 21.25 12 21.25C17.1 21.25 21.25 17.1 21.25 12C21.25 6.9 17.1 2.75 12 2.75Z"
						/>
						<path
							d="M4.89984 19.7499C4.70984 19.7499 4.51984 19.6799 4.36984 19.5299C4.07984 19.2399 4.07984 18.7599 4.36984 18.4699L18.3698 4.46994C18.6598 4.17994 19.1398 4.17994 19.4298 4.46994C19.7198 4.75994 19.7198 5.23994 19.4298 5.52994L5.42984 19.5299C5.27984 19.6799 5.08984 19.7499 4.89984 19.7499Z"
						/>
					</svg>
					<span>{errorMessage}</span>
				</div>
			{/if}
			<div class="flex flex-col gap-4">
				<label for="fullname">
					<input
						bind:value={fullname}
						class="w-full border-slate-400/80 border rounded-md bg-transparent focus:border-slate-950 caret-slate-950 py-2 px-4 outline-1 focus:outline-slate-950"
						type="text"
						id="fullname"
						placeholder="Full Name"
					/>
				</label>
				<label for="username">
					<input
						bind:value={username}
						class="w-full border-slate-400/80 border rounded-md bg-transparent focus:border-slate-950 caret-slate-950 py-2 px-4 outline-1 focus:outline-slate-950"
						type="text"
						id="username"
						placeholder="UserName"
					/>
				</label>
				<label for="email">
					<input
						bind:value={email}
						class="w-full border-slate-400/80 border rounded-md bg-transparent focus:border-slate-950 caret-slate-950 py-2 px-4 outline-1 focus:outline-slate-950"
						type="email"
						id="email"
						placeholder="Your Email"
					/>
				</label>
				<label for="password">
					<input
						bind:value={password}
						class="w-full border-slate-400/80 border rounded-md bg-transparent focus:border-slate-950 caret-slate-950 py-2 px-4 outline-1 focus:outline-slate-950"
						type="password"
						id="password"
						placeholder="Password"
					/>
				</label>
				<button
					on:click={handleRegister}
					class="cursor-pointer py-2.5 bg-slate-950 hover:bg-slate-900 hover:shadow-md text-slate-100 rounded-md"
				>
					Create Account
				</button>
			</div>
		</div>
	</div>
</main>
