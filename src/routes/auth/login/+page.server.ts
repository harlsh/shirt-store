import { fail, redirect } from '@sveltejs/kit';

export const actions = {
	login: async ({ cookies, request, url }) => {
		const data = await request.formData();
		const email = data.get('email');
		const password = data.get('password');


		if (!email) return fail(400, { email, message: 'You forgot to enter your email.' });

		if (!password) return fail(400, { email, message: 'You forgot to enter your password.' });

		const response = await fetch('http://localhost:8080/login', {
            method: "POST",
            headers:{
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ email, password})
        });

        const user = await response.json();

        if(response.status != 200) return fail(400, {email, message: user.error})

		cookies.set('session', user.AuthKey, { path: '/', httpOnly:true, sameSite: 'strict', maxAge: 60 * 60 * 24 * 30})

		// throw redirect(303, '/');
	}
};
