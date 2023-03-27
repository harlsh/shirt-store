<script>
		import { page } from '$app/stores'

		// console.log($page)
</script>
<header>
	<h1>
		<strong> Shirt Store </strong>
	</h1>
	{#if !$page.data.user}
    <a href="/auth/login">Login</a>
    <a href="/auth/register">Register</a>
	{:else}
	<a href="about"> {$page.data.user.firstName}</a>
  	{/if}

	<p></p>
	<hr />
</header>

<style>
</style>
