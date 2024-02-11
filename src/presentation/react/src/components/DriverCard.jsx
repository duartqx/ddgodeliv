import React from "react";
import Card from "./Card";

/** @param {{ driver: import("../services/driver/driver").Driver }} props */
export default function DriverCard({ driver }) {
  const parts = [
    { label: "", value: driver.user.email },
    { label: "", value: driver.license_id },
  ];
  return <Card title={driver.user.name} parts={parts} />;
}
