import httpClient from "./client";

/**
 * @typedef {{
 *  token: ?string
 *  expiresAt: ?string
 *  status: ?string
 * }} AuthResponse
 */

/** @returns {Promise<AuthResponse>} */
async function login({ email, password }) {
  try {
    const res = await httpClient().post("/login", {
      email,
      password,
    });

    /** @type {AuthResponse} */
    const data = res.data;

    if (data && data.token) {
      localStorage.setItem("auth", JSON.stringify(res.data));
    }
    return data;
  } catch (e) {
    console.log(e);
    return /** @type {AuthResponse} */ ({});
  }
}

export { login };
