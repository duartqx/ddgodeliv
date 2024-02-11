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
    const authData = JSON.parse(localStorage.getItem("auth") || "{}");
    const cacheVehicleKey = `cachedVehicles__${authData?.token}`;

    const cachedVehicles = JSON.parse(
      localStorage.getItem(cacheVehicleKey) || "{}"
    );

    if (
      cachedVehicles?.expiresAt &&
      (new Date(cachedVehicles?.expiresAt) > new Date())
    ) {
      return cachedVehicles.vehicles;
    }

    const res = await httpClient().get("/vehicle");

    if (res.data) {
      const newCachedVehicles = {
        vehicles: res.data,
        expiresAt: new Date().setMinutes(new Date().getMinutes() + 5),
      };
      localStorage.setItem(cacheVehicleKey, JSON.stringify(newCachedVehicles));
    }

    return res.data || [];
  } catch (e) {
    console.log(e);
    return [];
  }
}

export { companyVehicles };
