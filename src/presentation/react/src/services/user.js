import httpClient from "./client";

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

export default async function getUser() {
  /** @type {User} */
  var user = JSON.parse(localStorage.getItem("user") || "{}");

  if (!user?.id) {
    user = await httpClient().get("/user");
    localStorage.setItem("user", JSON.stringify(user));
  }

  return user;
}
