import httpClient from "../client";

/**
 * @typedef {{
 *   id: number
 *   email: string
 *   name: string
 * }} SignUpUser
 */

/**
 * @typedef {{
 *   id: number
 *   email: string
 *   name: string
 *   driver: {
 *     driver_id: number
 *     company_id: number
 *   }
 * }} User
 */

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

async function logout() {
  try {
    const res = await httpClient().delete("/logout");

    if (res.status >= 200 && res.status <= 299) {
      localStorage.removeItem("auth");
      localStorage.removeItem("user");
    }
  } catch (e) {
    console.log(e);
  }
}

/** @returns {Promise<SignUpUser>} */
async function register({ name, email, password }) {
  try {
    const res = await httpClient().post("/user", {
      name,
      email,
      password,
    });

    /** @type {User} */
    const user = res.data;

    return user;
  } catch (e) {
    console.log(e);
    return /** @type {User} */ ({});
  }
}

/** @returns {Promise<User>} */
async function getUser() {
  /** @type {User} */
  var user = JSON.parse(localStorage.getItem("user") || "{}");

  if (!user?.id) {
    user = await httpClient().get("/user");
    localStorage.setItem("user", JSON.stringify(user));
  }

  return user;
}
export { login, logout, getUser, register };
