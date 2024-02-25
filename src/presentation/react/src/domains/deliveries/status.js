const StatusDisplay = {
  0: "Pending",
  1: "Assigned",
  2: "InTransit",
  3: "Late",
  4: "Completed",
};

/**
 * @param {number} status
 * @returns {string}
 */
function getStatusDisplay(status) {
  return StatusDisplay[status];
}

/**
 * @param {number} status
 * @returns {boolean}
 */
function isPending(status) {
  return status === 0;
}

export { getStatusDisplay, isPending };
