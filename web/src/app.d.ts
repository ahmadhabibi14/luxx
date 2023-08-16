import type { User } from "./types/user";

declare global {
	namespace App {
		interface Error {
			code:			number,
			message:	string
		}
		interface Locals {
			authedUser: User | undefined
		}
		// interface PageData {}
		// interface Platform {}
	}
}

export {};