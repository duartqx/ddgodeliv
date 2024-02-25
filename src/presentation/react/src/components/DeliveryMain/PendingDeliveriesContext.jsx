import React, { useState, useEffect } from "react";
import * as deliveryService from "../../services/deliveries/deliveries";

export const PendingDeliveriesContext = React.createContext({
  /**
   * @param {string} filter
   * @returns {deliveryService.Delivery[]}
   */
  getFiltered: (filter) => [],
  /**
   * @param {deliveryService.Delivery} delivery
   * @returns {void}
   */
  removeDelivery: (delivery) => {},
  /**
   * @returns {deliveryService.Delivery[]}
   */
  getPendingDeliveries: () => [],
  /**
   * @param {deliveryService.Delivery[]} deliveries
   * @returns {void}
   */
  setPendingDeliveries: (deliveries) => {},
});

export default function PendingDeliveriesContextProvider({ children }) {
  const [pendingDeliveries, setPendingDeliveries] = useState(
    /** @type {deliveryService.Delivery[]} */ ([])
  );

  useEffect(() => {
    deliveryService
      .getPendingDeliveries()
      .then((deliveries) => setPendingDeliveries(deliveries));
  }, []);

  const getPendingDeliveries = () => pendingDeliveries;

  const getFiltered = (/** @type {string} */ filter) => {
    return filter
      ? pendingDeliveries.filter((d) =>
          d.loadout.toLowerCase().includes(filter.toLowerCase())
        )
      : pendingDeliveries;
  };

  const removeDelivery = (/** @type {deliveryService.Delivery} */ delivery) => {
    setPendingDeliveries(pendingDeliveries.filter((d) => d.id !== delivery.id));
  };

  return (
    <PendingDeliveriesContext.Provider
      value={{
        getPendingDeliveries,
        setPendingDeliveries,
        getFiltered,
        removeDelivery,
      }}
    >
      {children}
    </PendingDeliveriesContext.Provider>
  );
}
