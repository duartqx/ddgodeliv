import React from "react";
import Card from "./Card";

/** @param {{ vehicle: import("../services/vehicles/vehicles").Vehicle }} props */
export default function VehicleCard({ vehicle }) {
  const parts = [
    { label: "Year:", value: vehicle.model.year, border: true },
    { label: "Max Load:", value: vehicle.model.max_load, border: true },
    { label: "Manufacturer:", value: vehicle.model.manufacturer, border: true },
    { label: "License Number:", value: vehicle.license_id, border: false },
  ];
  return <Card title={vehicle.model.name} parts={parts} />;
}
