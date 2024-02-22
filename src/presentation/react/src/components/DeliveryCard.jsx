import React from "react";
import Card from "./SideBarList/Card";

/**
 * @param {{
 *  onClickHandler: () => void,
 *  selected: boolean
 *  delivery: import("../services/deliveries/deliveries").Delivery
 * }} props
 */
export default function DeliveryCard({ onClickHandler, selected, delivery }) {
  /** @type {import("./SideBarList/Card").CardPartProps[]} */
  const parts = [
    { label: "Weight:", value: delivery.weight, border: true },
    { label: "Origin:", value: delivery.origin, border: true },
    { label: "Destination:", value: delivery.destination, border: false },
  ];
  return (
    <Card
      selected={selected}
      title={delivery.loadout}
      parts={parts}
      onClickHandler={onClickHandler}
    />
  );
}
