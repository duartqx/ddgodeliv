import React, { useState, useEffect, useContext } from "react";
import CardForm from "../CardForm/CardForm";
import * as driverService from "../../services/driver/driver";
import * as deliveryService from "../../services/deliveries/deliveries";
import * as deliveryStatus from "../../domains/deliveries/status";
import { PendingDeliveriesContext } from "./PendingDeliveriesContext";

/**
 * @param {{
 *  delivery: import("../../services/deliveries/deliveries").Delivery
 *  dissmissForm: () => void
 * }} props
 */
export default function AssignToDriverForm({ delivery, dissmissForm }) {
  const [driverId, setDriverId] = useState(null);
  const [drivers, setDrivers] = useState(
    /** @type {driverService.Driver[]} */ ([])
  );
  const { removeDelivery } = useContext(PendingDeliveriesContext);

  useEffect(() => {
    driverService.companyDrivers().then((drivers) => setDrivers(drivers));
  }, []);

  const handleSubmit = async (/** @type {React.FormEvent} */ e) => {
    e.preventDefault();

    await deliveryService.assignDriver({
      deliveryId: delivery.id,
      driverId: driverId,
    });

    removeDelivery(delivery);

    if (dissmissForm) {
      dissmissForm();
    }
  };

  return (
    <>
      <CardForm
        handleSubmit={handleSubmit}
        title={`Assign Driver to ${delivery.loadout}`}
        inputs={[
          {
            label: "Driver",
            type: "select",
            options: drivers
              .filter((d) => deliveryStatus.isPending(d.status))
              .map((d) => ({
                value: d.id,
                label: `${d.id}: ${d.user.email}`,
              })),
            onChangeHandler: (e) => setDriverId(e.target.value),
          },
        ]}
      />
    </>
  );
}
