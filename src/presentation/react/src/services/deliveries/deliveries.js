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
 * @param {number} driverId
 * @returns {string}
 **/
function getDriverDeliveriesCacheKey(driverId) {
  const authData = JSON.parse(localStorage.getItem("auth") || "{}");
  return `cachedDriverDeliveries__${driverId}__${authData?.token}`;
}

/**
 * @param {number} driverId
 * @returns {string}
 **/
function getDriverCurrentDeliveryCacheKey(driverId) {
  const authData = JSON.parse(localStorage.getItem("auth") || "{}");
  return `cachedDriverCurrentDelivery__${driverId}__${authData?.token}`;
}

/** @returns {string} */
function getCompanyDeliveriesCacheKey() {
  const authData = JSON.parse(localStorage.getItem("auth") || "{}");
  return `cachedCompanyDeliveries__${authData?.token}`;
}

/**
 * @param {number} driverId
 * @returns {Promise<Delivery[]>}
 **/
async function getOtherDeliveriesByDriverId(driverId) {
  try {
    const cachedDeliveries = await cache.getFromCache(
      getDriverDeliveriesCacheKey(driverId)
    );
    if (cachedDeliveries) {
      // Can be an empty array
      return cachedDeliveries;
    }

    const res = await httpClient().get(`/delivery/company/driver/${driverId}/`);

    const deliveries = res.data || [];

    cache.setToCache(getDriverDeliveriesCacheKey(driverId), deliveries);

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
      getDriverCurrentDeliveryCacheKey(driverId)
    );
    if (cachedDelivery) {
      // Can be an empty array
      return cachedDelivery;
    }

    const res = await httpClient().get(
      `/delivery/company/driver/${driverId}/current/`
    );

    const delivery = res.data || null;

    cache.setToCache(getDriverCurrentDeliveryCacheKey(driverId), delivery);

    return delivery;
  } catch (e) {
    return null;
  }
}

/** @returns {Promise<Delivery[]>} */
async function getByCompanyId() {
  try {
    const cachedDeliveries = await cache.getFromCache(
      getCompanyDeliveriesCacheKey()
    );
    if (cachedDeliveries) {
      // Can be an empty array
      return cachedDeliveries;
    }

    const res = await httpClient().get(`/delivery/company/`);

    const deliveries = res.data || [];
    cache.setToCache(getCompanyDeliveriesCacheKey(), deliveries);
    return deliveries;
  } catch (e) {
    return [];
  }
}

export { getOtherDeliveriesByDriverId, getCurrentByDriverId, getByCompanyId };
