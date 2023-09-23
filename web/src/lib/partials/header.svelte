<script>
	export let user_data;
</script>

<header
	class="flex flex-row justify-end items-center bg-zinc-900 border border-zinc-700 rounded-2xl shadow-sm text-zinc-100 h-16 w-full py-6 px-4"
>
	<div class="flex flex-row items-center gap-4 font-medium text-base cursor-pointer group">
		<span class="group-hover:text-sky-500">{user_data?.full_name}</span>
		<div class="w-10 h-auto rounded-full overflow-hidden">
			<img
				class="group-hover:grayscale w-full h-full object-cover"
				src={user_data?.avatar?.String === "/avatars/default.png"
					? user_data?.avatar?.String
					: user_data?.avatar?.String}
				alt="avatar"
			/>
		</div>
	</div>
</header>
