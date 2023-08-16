import type { Actions } from "@sveltejs/kit";

export const actions: Actions = {
  default: async (event) => {
    const formData = await event.request.formData();
    const email = String(formData.get("email"));
    const password = formData.get("password");
    const body = await JSON.stringify({
      email,
      password,
    });
    console.log(body);
  }
};