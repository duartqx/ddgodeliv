import React from "react";
import Card from "./Card";

/** @param {{ vehicle: import("../services/vehicles/vehicles").Vehicle }} props */
export default function VehicleCard({ vehicle }) {
  const parts = [
    { label: "Year:", value: vehicle.model.year },
    { label: "Max Load:", value: vehicle.model.max_load },
    { label: "Manufacturer:", value: vehicle.model.manufacturer },
    { label: "License Number:", value: vehicle.license_id },
  ];
  return <Card title={vehicle.model.name} parts={parts} />;
}
