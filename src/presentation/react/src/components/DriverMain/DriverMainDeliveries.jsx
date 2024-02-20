import React, { useEffect, useState } from "react";
import * as deliveryService from "../../services/deliveries/deliveries";
import DottedLink from "../DottedLink";
import { DeliveryStatus } from "../../domains/deliveries/status";
import useWidthHeight from "../../middlewares/useWidthHeight";

/** @param {{ margin: { marginLeft?: string, marginRight?: string } }} props */
function MiddleLine({ margin = { marginLeft: "1rem" } }) {
  return (
    <div
      className="flex-grow-1"
      style={{ borderBottom: "1px solid #f0f2f7", ...margin }}
    ></div>
  );
}

/** @param {{ delivery: import("../../services/deliveries/deliveries").Delivery }} props */
function DeliveryCard({ delivery }) {
  return (
    delivery && (
      <div className="py-3">
        <div className="row">
          <div className="col-5 fw-bold">ID: {delivery.id}</div>
          <div className="col-2 fw-light text-center">
            {delivery.sender.name}
          </div>
          <div className="col-5 fw-light text-end">
            {delivery.weight}kg | {DeliveryStatus[delivery.status]}
          </div>
        </div>
        <div className="row">
          <div className="col-5 fw-medium">{delivery.origin}</div>
          <i className="col-2 bi bi-arrow-right text-center"></i>
          <div className="col-5 fw-medium text-end">{delivery.destination}</div>
        </div>
      </div>
    )
  );
}

/** @param {{ delivery: import("../../services/deliveries/deliveries").Delivery }} props */
function CurrentDelivery({ delivery }) {
  return (
    <>
      <div className="d-flex align-items-center my-3">
        <div className="fw-medium">Current</div>
        <MiddleLine />
      </div>

      <DeliveryCard delivery={delivery} />
      <img
        className="img-fluid img-thumbnail"
        style={{
          objectFit: "cover",
          height: "16rem",
          width: "100%",
        }}
        src={
          delivery?.id
            ? "https://www.ergosum.co/wp-content/uploads/2017/06/Google-Maps-Route.png"
            : "https://www.drodd.com/images14/white2.jpg"
        }
        alt=""
      />
    </>
  );
}

/** @param {{ driver: import("../../services/driver/driver").Driver }} props */
export default function DriverMainDeliveries({ driver }) {
  const [driverDeliveries, setDriverDeliveries] = useState(
    /** @type {import("../../services/deliveries/deliveries").Delivery[]} */ ([])
  );
  const [currentDelivery, setCurrentDelivery] = useState(
    /** @type {import("../../services/deliveries/deliveries").Delivery} */
    (null)
  );
  const { isSmallWindow } = useWidthHeight();

  useEffect(() => {
    if (isSmallWindow()) {
      deliveryService
        .getOtherDeliveriesByDriverId(driver.id)
        .then((deliveries) => {
          setDriverDeliveries(deliveries);
        });
    }

    deliveryService
      .getCurrentByDriverId(driver.id)
      .then((delivery) => setCurrentDelivery(delivery));
  }, [driver]);

  return (
    <div className="flex-grow-1 p-2 d-flex flex-column">
      {!isSmallWindow() && (
        <div className="d-flex align-items-center flex-grow-1 fw-bold my-3">
          Deliveries
        </div>
      )}
      <CurrentDelivery delivery={currentDelivery} />
      <div className="my-3 d-flex flex-column justify-content-center flex-grow-1">
        {isSmallWindow() && (
          <div className="d-flex align-items-center flex-grow-1 mb-auto">
            <div>History</div>
            <MiddleLine />
          </div>
        )}
        {driverDeliveries
          .filter((delivery) => delivery.id !== currentDelivery?.id)
          .splice(0, 3)
          .map((delivery) => (
            <DeliveryCard key={delivery.id} delivery={delivery} />
          ))}
      </div>
      <div className="d-flex align-items-center flex-grow-1 mb-auto">
        <MiddleLine margin={{ marginRight: "1rem" }} />
        <DottedLink to="#" label={isSmallWindow() ? "See more" : "History"} />
      </div>
    </div>
  );
}
