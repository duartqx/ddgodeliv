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

/** @returns {string} */
function getCacheKey() {
  const authData = JSON.parse(localStorage.getItem("auth") || "{}");
  return `cachedDriver__${authData?.token}`;
}

async function invalidateCache() {
  localStorage.removeItem(getCacheKey())
}

/** @returns {Promise<Driver[] | null>} */
async function getFromCache() {
  const cachedDrivers = JSON.parse(localStorage.getItem(getCacheKey()) || "{}");

  if (
    cachedDrivers?.expiresAt &&
    new Date(cachedDrivers?.expiresAt) > new Date()
  ) {
    return cachedDrivers.drivers;
  }
  return null;
}

/** @param {Driver[]} drivers */
async function setToCache(drivers) {
  const newCachedDrivers = {
    drivers: drivers,
    expiresAt: new Date().setMinutes(new Date().getMinutes() + 5),
  };
  localStorage.setItem(getCacheKey(), JSON.stringify(newCachedDrivers));
}

/** @returns {Promise<Driver[]>} */
async function companyDrivers() {
  try {
    const cachedDrivers = await getFromCache();
    if (cachedDrivers?.length) {
      return cachedDrivers;
    }

    const res = await httpClient().get("/driver");

    if (res.data) {
      setToCache(res.data);
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
    const res = await httpClient().post("/driver", {
      license_id: license,
      user: {
        name: name,
        email: email,
      },
    });

    if (res.data) {
      invalidateCache();
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
async function deleteDriver({id}) {
  try {
    await httpClient().delete(`/driver/${id}`)
    invalidateCache()
    return true
  } catch (e) {
    console.log(e)
    return false
  }
}

export { companyDrivers, createDriver, deleteDriver };
