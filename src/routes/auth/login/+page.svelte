<script lang="ts">
	import { fly, slide } from 'svelte/transition';
	import { enhance } from '$app/forms';
	import Spinner from 'svelte-spinner';

	export let form;

	let loading = false;
</script>

<form
	method="POST"
	action="?/login"
	data-sveltekit-keepfocus
	out:slide
	use:enhance={() => {
		loading = true;

		return async ({ update }) => {
			await update();
			loading = false;
		};
	}}
>
	{#if form?.message}
		<p in:fly>{form?.message}</p>
	{/if}

	<label>
		Email
		<input name="email" type="email" minlength="8" value={form?.email ?? ''} />
	</label>
	<label>
		Password
		<input name="password" type="password" minlength="8" />
	</label>
	<div class="btn">
		<button disabled={loading}>
			{#if loading}
				<Spinner />
			{:else}
				Log in
			{/if}
		</button>
		<button>
			<a href="/auth/register"> Register </a>
		</button>
	</div>
</form>

<style>
	form {
		width: 100%;
		height: 100%;
		display: flex;
		flex-direction: column;
		text-align: right;
		margin-right: 0;
		align-items: center;
		/* justify-content: center; */
		gap: 1rem;
		flex: 1;
	}

	/* input {
		float: right;
	} */
	input:invalid {
		background-color: #ffdddd;
	}

	.btn {
		display: inline-block;
		/* width: calc(50% - 4px);
	margin: 0 auto; */
	}
</style>
