import { req } from "./axios";

/**
 * @typedef {{
 *  token: ?string
 *  expiresAt: ?string
 *  status: ?string
 *  user: ?{
 *    id: number
 *    email: string
 *    name: string
 *    driver: {
 *      driver_id: number
 *      company_id: number
 *    }
 *  }
 * }} AuthResponse
 */

/** @returns {Promise<AuthResponse>} */
async function login({ email, password }) {
  try {
    const res = await req.post("/login", {
      email,
      password,
    });

    /** @type {AuthResponse} */
    const data = res.data;

    console.log(data);

    if (data && data.token) {
      sessionStorage.setItem("auth", JSON.stringify(res.data));
    }
    return data;
  } catch (e) {
    console.log(e);
    return /** @type {AuthResponse} */ ({});
  }
}

export { login };
