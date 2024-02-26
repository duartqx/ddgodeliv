async function invalidateCache(/** @type {string[]} */ ...keys) {
  for (const key of keys) {
    localStorage.removeItem(key);
  }
}

/** @returns {Promise<any | null>} */
async function getFromCache(/** @type {string} */ key) {
  const cached = JSON.parse(localStorage.getItem(key) || "{}");

  if (cached?.expiresAt && new Date(cached?.expiresAt) > new Date()) {
    return cached.values;
  }
  return null;
}

/**
 * @param {string} key
 * @param {any[]} values
 */
async function setToCache(key, values) {
  const newCached = {
    values: values,
    expiresAt: new Date().setMinutes(new Date().getMinutes() + 5),
  };
  localStorage.setItem(key, JSON.stringify(newCached));
}

export { invalidateCache, getFromCache, setToCache };
