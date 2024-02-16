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
 *  status: ?boolean
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
      localStorage.clear();
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
    try {
      const res = await httpClient().get("/user");
      if (res.data?.id) {
        user = res.data;
        localStorage.setItem("user", JSON.stringify(user));
      }
    } catch (e) {
      console.log(e);
    }
  }

  return user;
}

/** @returns {AuthResponse} */
function getAuth() {
  return JSON.parse(localStorage.getItem("auth") || "{}");
}

export { login, logout, register, getUser, getAuth };
