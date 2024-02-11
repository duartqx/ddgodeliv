import httpClient from "../client";

/**
 * @typedef {{
 *   id: number
 *   license_id: string
 *   user: {
 *     id: number
 *     email: string
 *     name: string
 *   }
 *   company: {
 *     id: number
 *     owner_id: number
 *     name: string
 *   }
 * }} Driver
 */

/** @returns {Promise<Driver[]>} */
async function companyDrivers() {
  try {
    const res = await httpClient().get("/driver");
    return res.data || [];
  } catch (e) {
    console.log(e);
    return [];
  }
}

export { companyDrivers };
