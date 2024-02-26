import httpClient from "../client";
import * as cache from "../cache";

/**
 * @typedef {{
 *   id: number
 *   license_id: string
 *   status: number
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

/** @returns {string} */
function getDriverCacheKey() {
  const authData = JSON.parse(localStorage.getItem("auth") || "{}");
  return `cachedDriver__${authData?.token}`;
}

/** @returns {Promise<Driver[]>} */
async function companyDrivers() {
  try {
    const cachedDrivers = await cache.getFromCache(getDriverCacheKey());
    if (cachedDrivers?.length) {
      return cachedDrivers;
    }

    const res = await httpClient().get("/api/driver");

    if (res.data) {
      cache.setToCache(getDriverCacheKey(), res.data);
    }

    return res.data || [];
  } catch (e) {
    console.log(e);
    return [];
  }
}

/** @returns {Promise<Driver>} */
async function createDriver({ name, email, license }) {
  try {
    const res = await httpClient().post("/api/driver", {
      license_id: license,
      user: {
        name: name,
        email: email,
      },
    });

    if (res.data) {
      cache.invalidateCache(getDriverCacheKey());
      return res.data;
    }
  } catch (e) {
    console.log(e);
  }
  return /** @type {Driver} */ ({});
}

/**
 * @param {{id: number}} params
 * @returns {Promise<boolean>}
 */
async function deleteDriver({ id }) {
  try {
    await httpClient().delete(`/api/driver/${id}`);
    cache.invalidateCache(getDriverCacheKey());
    return true;
  } catch (e) {
    console.log(e);
    return false;
  }
}

export { companyDrivers, createDriver, deleteDriver };
