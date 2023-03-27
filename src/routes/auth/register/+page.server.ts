import { fail, redirect } from '@sveltejs/kit';

export const actions = {
	register: async ({ cookies, request, url }) => {
		const data = await request.formData();
		const email = data.get('email');
        const firstName = data.get('first-name');
        const lastName = data.get('last-name');
		const password = data.get('password');
        const passwordConfirm = data.get('password-confirmation');



		if (!email) return fail(400, { email, message: 'You forgot to enter your email.' });
		if (!password) return fail(400, { email, message: 'You forgot to enter your password.' });
        if (!firstName) return fail(400, { email, message: 'You forgot to enter your First Name.' });
        if (!lastName) return fail(400, { email, message: 'You forgot to enter your Last Name.' });
        if (!passwordConfirm) return fail(400, { email, message: 'You forgot to enter confirmation password.' });

        if( password.length < 8 ) return fail(400, { email, message: 'Password should atleast be 8 characters long' });

        if( password !== passwordConfirm) return fail(400, { email, message: 'Passwords dont match.' });

        //check in database if user is already present
        const user = {
            email, firstName, lastName, password
        }
        

        const response = await fetch('http://localhost:8080/api/register', {
            method: "POST",
            headers:{
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(user)
        });
        
        const responseJson = await response.json();

        if(response.status != 200) return fail(400, {email, message: responseJson.error})
        

		throw redirect(303, '/auth/login');
	}
};
