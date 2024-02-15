import React from "react";
import Card from "../Card";

/** @param {{
 *  deleteHandler: () => void,
 *  vehicle: import("../../../services/vehicles/vehicles").Vehicle
 * }} props */
export default function VehicleCard({ deleteHandler, vehicle }) {
  const parts = [
    { label: "Year:", value: vehicle.model.year, border: true },
    { label: "Transmission:", value: vehicle.model.transmission, border: true },
    { label: "Type:", value: vehicle.model.type, border: true },
    { label: "Manufacturer:", value: vehicle.model.manufacturer, border: true },
    { label: "License:", value: vehicle.license_id, border: false },
  ];
  return (
    <Card
      title={vehicle.model.name}
      parts={parts}
      deleteHandler={deleteHandler}
    />
  );
}
