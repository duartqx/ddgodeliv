import React from "react";
import DriverMainDriverInfo from "./DriverMainDriverInfo";
import DriverMainDeliveries from "./DriverMainDeliveries";
import DriverMainStatistics from "./DriverMainStatistics";

export default function DriverMainCarded({ driver }) {
  return (
    <>
      <DriverMainDriverInfo driver={driver} />
      <div className="row">
        <div className="col-sm-6">
          <DriverMainDeliveries driver={driver} />
        </div>
        <div className="col-sm-6">
          <DriverMainStatistics driver={driver} />
        </div>
      </div>
    </>
  );
}
