import React from "react";
import DriverMainDriverInfo from "./DriverMainDriverInfo";
import DriverMainDeliveries from "./DriverMainDeliveries";
import DriverMainStatistics from "./DriverMainStatistics";

export default function DriverMainCarded({ driver }) {
  return (
    <div className="d-flex flex-column" style={{ maxHeight: "87vh" }}>
      <DriverMainDriverInfo driver={driver} />
      <div className="row" style={{ height: "60vh" }}>
        <div className="col-sm-6">
          <DriverMainDeliveries driver={driver} />
        </div>
        <div className="col-sm-6">
          <DriverMainStatistics driver={driver} />
        </div>
      </div>
    </div>
  );
}
