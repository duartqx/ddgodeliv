import httpClient from "../client";

/**
 * @typedef {{
 *   id: number
 *   license_id: string
 *   model: {
 *     id: number
 *     name: string
 *     manufacturer: string
 *     year: number
 *     max_load: number
 *   }
 *   company: {
 *     id: number
 *     owner_id: number
 *     name: string
 *   }
 * }} Vehicle
 */

/** @returns {Promise<Vehicle[]>} */
async function companyVehicles() {
  try {
    const res = await httpClient().get("/vehicle");
    return res.data || [];
  } catch (e) {
    console.log(e);
    return [];
  }
}

export { companyVehicles };
