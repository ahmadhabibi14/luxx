import type { User } from "./types/user";

declare global {
	namespace App {
		interface Error {
			code:			number,
			message:	string
		}
		// interface Locals {}
		// interface PageData {}
		// interface Platform {}
	}
}

export {};