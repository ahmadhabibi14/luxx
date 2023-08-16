import { redirect } from "@sveltejs/kit";
import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = async ({cookies}) => {
  // get session_token from cookies
  const session_token = cookies.get("session_token");

  // If there is a session_token, redirect to the home page
  if (session_token) {
    throw redirect(302, "/");
  }
}