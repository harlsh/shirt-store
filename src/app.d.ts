// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			user: {
				Email: string
				FirstName: string
				LastName: string
			}
		}
		// interface PageData {}
		// interface Platform {}
	}
}

export {};
