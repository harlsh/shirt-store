import { fail, redirect } from '@sveltejs/kit';

export const actions = {
	login: async ({ cookies, request, url }) => {
		const data = await request.formData();
		const email = data.get('email');
		const password = data.get('password');

		await new Promise((fulfil) => setTimeout(fulfil, 1000));

		if (!email) return fail(400, { email, message: 'You forgot to enter your email.' });

		if (!password) return fail(400, { email, message: 'You forgot to enter your password.' });

		throw redirect(303, '/');
	}
};
