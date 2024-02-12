import React from "react";
import Card from "./Card";

/** @param {{
 * deleteHandler: Function
 * driver: import("../services/driver/driver").Driver
 * }} props */
export default function DriverCard({ deleteHandler, driver }) {
  const parts = [
    { label: "", value: driver.user.email, border: true },
    { label: "", value: driver.license_id, border: false },
  ];
  return (
    <Card
      title={driver.user.name}
      parts={parts}
      deleteHandler={deleteHandler}
    />
  );
}
