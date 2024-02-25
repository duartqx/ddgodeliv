import httpClient from "../client";
import * as cache from "../cache";

/**
 * @typedef {{
 *  id: number
 *  loadout: string
 *  weight: number
 *  origin: string
 *  destination: string
 *  created_at: string
 *  deadline: string
 *  status: number
 *  driver: import("../driver/driver").Driver
 *  sender: import("../auth/auth").User
 * }} Delivery
 */

/**
 * @param {string} prefix
 * @param {number} driverId
 * @returns {string}
 **/
function getCacheKey(prefix, driverId) {
  const authData = JSON.parse(localStorage.getItem("auth") || "{}");
  if (driverId) {
    return `${prefix}__${driverId}__${authData?.token}`;
  }
  return `${prefix}__${authData?.token}`;
}

/**
 * @param {number} driverId
 * @returns {Promise<Delivery[]>}
 **/
async function getOtherDeliveriesByDriverId(driverId) {
  try {
    const cachedDeliveries = await cache.getFromCache(
      getCacheKey("cachedDriverDeliveries", driverId),
    );
    if (cachedDeliveries) {
      // Can be an empty array
      return cachedDeliveries;
    }

    const res = await httpClient().get(
      `/delivery/company/driver/${driverId}/`,
    );

    const deliveries = res.data || [];

    cache.setToCache(
      getCacheKey("cachedDriverDeliveries", driverId),
      deliveries,
    );

    return deliveries;
  } catch (e) {
    return [];
  }
}

/**
 * @param {number} driverId
 * @returns {Promise<Delivery | null>}
 **/
async function getCurrentByDriverId(driverId) {
  try {
    const cachedDelivery = await cache.getFromCache(
      getCacheKey("cachedDriverCurrentDelivery", driverId),
    );
    if (cachedDelivery) {
      // Can be an empty array
      return cachedDelivery;
    }

    const res = await httpClient().get(
      `/delivery/company/driver/${driverId}/current/`,
    );

    const delivery = res.data || null;

    cache.setToCache(
      getCacheKey("cachedDriverCurrentDelivery", driverId),
      delivery,
    );

    return delivery;
  } catch (e) {
    return null;
  }
}

/** @returns {Promise<Delivery[]>} */
async function getByCompanyId() {
  try {
    const cachedDeliveries = await cache.getFromCache(
      getCacheKey("cachedCompanyDeliveries"),
    );
    if (cachedDeliveries) {
      // Can be an empty array
      return cachedDeliveries;
    }

    const res = await httpClient().get(`/delivery/company/`);

    const deliveries = res.data || [];
    cache.setToCache(getCacheKey("cachedCompanyDeliveries"), deliveries);
    return deliveries;
  } catch (e) {
    return [];
  }
}

/** @returns {Promise<Delivery[]>} */
async function getPendingDeliveries() {
  try {
    const cachedDeliveries = await cache.getFromCache(
      getCacheKey("cachedPendingDeliveries"),
    );
    if (cachedDeliveries) {
      // Can be an empty array
      return cachedDeliveries;
    }

    const res = await httpClient().get(`/delivery/pending/`);

    const deliveries = res.data || [];
    cache.setToCache(getCacheKey("cachedPendingDeliveries"), deliveries);
    return deliveries;
  } catch (e) {
    return [];
  }
}

/** @returns {Promise<boolean>} */
async function assignDriver({ deliveryId, driverId }) {
  try {
    const res = await httpClient().patch(`/delivery/${deliveryId}/assign/`, {
      driver_id: Number(driverId),
    });

    cache.invalidateCache(
      getCacheKey("cachedPendingDeliveries"),
      getCacheKey("cachedCompanyDeliveries"),
      getCacheKey("cachedDriver"),
    );
    return res.status === 200;
  } catch (e) {
    console.log(e);
    return false;
  }
}

export {
  assignDriver,
  getOtherDeliveriesByDriverId,
  getCurrentByDriverId,
  getByCompanyId,
  getPendingDeliveries,
};
