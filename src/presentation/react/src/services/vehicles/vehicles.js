import httpClient from "../client";
import * as cache from "../cache";

/**
 * @typedef {{
 *   id: number
 *   license_id: string
 *   model: import("./vehicleModels").VehicleModel
 *   company: {
 *     id: number
 *     owner_id: number
 *     name: string
 *   }
 * }} Vehicle
 */

/** @returns {string} */
function getVehiclesCacheKey() {
  const authData = JSON.parse(localStorage.getItem("auth") || "{}");
  return `cachedVehicles__${authData?.token}`;
}

/** @returns {Promise<Vehicle[]>} */
async function companyVehicles() {
  const cacheKey = getVehiclesCacheKey();
  try {
    const cachedVehicles = await cache.getFromCache(cacheKey);
    if (cachedVehicles?.length) {
      return cachedVehicles;
    }

    const res = await httpClient().get("/api/vehicle");

    if (res.data) {
      cache.setToCache(cacheKey, res.data);
    }

    return res.data || [];
  } catch (e) {
    console.log(e);
    return [];
  }
}

/**
 * @param {{ license: string, model: number }}
 * @returns {Promise<Vehicle>}
 */
async function createVehicle({ license, model }) {
  try {
    const res = await httpClient().post("/api/vehicle", {
      license_id: license,
      model_id: model,
    });

    if (res.data) {
      cache.invalidateCache(getVehiclesCacheKey());
      return res.data;
    }
  } catch (e) {
    console.log(e);
    return {};
  }
}

/**
 * @param {number} id
 * @returns {Promise<boolean>}
 */
async function deleteVehicle({ id }) {
  try {
    await httpClient().delete(`/api/vehicle/${id}`);
    cache.invalidateCache(getVehiclesCacheKey());
    return true;
  } catch (e) {
    console.log(e);
    return false;
  }
}

export { companyVehicles, createVehicle, deleteVehicle };
