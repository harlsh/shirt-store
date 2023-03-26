<script>
    import { enhance } from '$app/forms';
    import { fly } from 'svelte/transition';
	import Spinner from 'svelte-spinner';

	export let form;

	let loading = false;
</script>

<form 
method="POST"
	action="?/register"
	data-sveltekit-keepfocus
	use:enhance={() => {
		loading = true;

		return async ({ update }) => {
			await update();
			loading = false;
		};
	}}>
    	{#if form?.message}
		<p in:fly>{form?.message}</p>
	{/if}
	<label for="email">Email <input type="email" name="email" minlength="8" value={form?.email ?? ''}/></label>

	<label for="first-name">First Name <input type="text" name="first-name" /></label>

	<label for="last-name">Last Name <input type="text" name="last-name" /></label>

	<label for="password">Password <input type="password" name="password" /> </label>

	<label for="password-conformation"
		>Reenter Password <input type="password" name="password-confirmation" /></label
	>

	<button disabled={loading}>
		{#if loading}
			<Spinner />
		{:else}
			Register
		{/if}
	</button>
</form>

<style>
	form {
		width: 40%;
		height: 80%;
		display: flex;
		flex-direction: column;
		text-align: right;
		margin-right: 0;
		/* align-items: center; */
	}
</style>
