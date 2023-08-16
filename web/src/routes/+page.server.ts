import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ locals }) => {
  let authedUser = undefined;
  if (locals.authedUser) authedUser = locals.authedUser;
  return { authedUser }
}