import React from "react";
import Card from "./Card";

/** @param {{ driver: import("../services/driver/driver").Driver }} props */
export default function DriverCard({ driver }) {
  const parts = [
    { label: "", value: driver.user.email, border: true },
    { label: "", value: driver.license_id, border: false },
  ];
  return <Card title={driver.user.name} parts={parts} />;
}
