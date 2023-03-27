import type { Handle } from '@sveltejs/kit'


export const handle: Handle = async ({ event, resolve }) => {
  // get cookies from browser
  const token = event.cookies.get('jwt')
  // console.log("token from the browser: ", token)
  if (!token) {
    // if there is no session load page as normal
    return await resolve(event)
  }

  const response = await fetch('http://localhost:8080/api/me', { 
    headers: {
      'Cookie': `jwt=${token}`
    }});
  const user = await response.json()
  console.log("user in hook: ", user.user)
  if(response.status == 200)
    event.locals.user = user.user;

  // load page as normal
  return await resolve(event)
}