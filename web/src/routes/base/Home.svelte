<script>
	export let user_data;
	let post_input;
	let post_input_row = 1;
	function dynamicPostInputHeight() {
		if (post_input.scrollHeight > post_input.clientHeight) {
			post_input_row += 1;
		}
	}
	function backSpace(event) {
		if (event.key === "Backspace") {
			if (post_input_row > post_input.value.split("\n").length) {
				post_input_row -= 1;
			}
		}
	}
</script>

<div class="font-medium flex flex-col gap-3 relative">
	<div class="flex flex-col gap-3 p-2 bg-zinc-800 shadow rounded-xl relative w-full min-h-[130px]">
		<textarea
			on:keydown={backSpace}
			on:input={dynamicPostInputHeight}
			bind:this={post_input}
			rows={post_input_row}
			name="post"
			id="post"
			placeholder="What's your feeling..."
			class="post_input"
		></textarea>
		<div class="flex flex-row justify-end">
			<button class="bg-sky-500 hover:bg-sky-600 rounded-full text-white py-2 px-5"> Post </button>
		</div>
	</div>
	<div class="h-[800px]">
		<p>Welcome {user_data?.full_name}</p>
	</div>
</div>
