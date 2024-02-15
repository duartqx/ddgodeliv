import React from "react";
import Card from "../Card";
import { DeliveryStatus } from "../../../domains/deliveries/status";

/** @param {{
 * onClickHandler: () => void
 * driver: import("../../../services/driver/driver").Driver
 * selected: boolean
 * }} props */
export default function DriverCard({ onClickHandler, driver, selected }) {
  const parts = [
    { label: "Status:", value: DeliveryStatus[driver.status], border: false },
  ];
  return (
    <Card
      selected={selected}
      onClickHandler={onClickHandler}
      title={driver.user.name}
      parts={parts}
    />
  );
}
