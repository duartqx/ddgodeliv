import httpClient from "../client";
import * as cache from "../cache";

/**
 * @typedef {{
 *     id: number
 *     name: string
 *     manufacturer: string
 *     year: number
 *     transmission: string
 *     type: string
 *   }} VehicleModel
 */

/** @returns {string} */
function getVehicleModelsCacheKey() {
  const authData = JSON.parse(localStorage.getItem("auth") || "{}");
  return `cachedVehicleModels__${authData?.token}`;
}

/** @returns {Promise<VehicleModel[]>} */
async function getVehicleModels() {
  const cacheKey = getVehicleModelsCacheKey();
  try {
    const cachedVehicleModels = await cache.getFromCache(cacheKey);
    if (cachedVehicleModels?.length) {
      return cachedVehicleModels;
    }

    const res = await httpClient().get("/api/vehicle/model");

    if (res.data) {
      cache.setToCache(cacheKey, res.data);
    }

    return res.data || [];
  } catch (e) {
    console.log(e);
    return [];
  }
}

export { getVehicleModels };
