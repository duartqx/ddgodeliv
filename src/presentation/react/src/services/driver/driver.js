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
    const authData = JSON.parse(localStorage.getItem("auth") || "{}");
    const cacheDriverKey = `cachedDriver__${authData?.token}`;

    const cachedDriver = JSON.parse(
      localStorage.getItem(cacheDriverKey) || "{}"
    );

    if (
      cachedDriver?.expiresAt &&
      (new Date(cachedDriver?.expiresAt) > new Date())
    ) {
      return cachedDriver.drivers;
    }

    const res = await httpClient().get("/driver");

    if (res.data) {
      const newCachedDrivers = {
        drivers: res.data,
        expiresAt: new Date().setMinutes(new Date().getMinutes() + 5),
      };
      localStorage.setItem(cacheDriverKey, JSON.stringify(newCachedDrivers));
    }

    return res.data || [];
  } catch (e) {
    console.log(e);
    return [];
  }
}

export { companyDrivers };
