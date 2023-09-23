<script>
	export let user_data;
</script>

<div class="font-medium flex flex-col gap-3 relative">
	<div class="flex flex-col gap-3 p-2 bg-zinc-800 shadow rounded-xl absolute w-full h-[130px]">
		<textarea name="post" id="post" placeholder="What's your feeling..." class="post_input"
		></textarea>
		<div class="flex flex-row justify-end">
			<button class="bg-sky-500 hover:bg-sky-600 rounded-full text-white py-2 px-5"> Post </button>
		</div>
	</div>
	<div class="h-[800px] mt-[145px]">
		<p>Welcome {user_data?.full_name}</p>
	</div>
</div>
