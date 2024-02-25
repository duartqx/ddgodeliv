import React from "react";
import DeliveryMainHeader from "./DeliveryMainHeader";

/**
 * @param {{
 * delivery: import("../../services/deliveries/deliveries").Delivery
 * handleFilterOutDelivery: (delivery: import("../../services/deliveries/deliveries").Delivery) => void
 * }} props
 */
export default function DeliveryMainPending({
  delivery,
  handleFilterOutDelivery,
}) {
  return (
    delivery && (
      <>
        <div
          className="d-flex flex-column mx-4 flex-grow-1"
          style={{ maxHeight: "calc(100vh - 4rem)" }}
        >
          <DeliveryMainHeader
            delivery={delivery}
            handleFilterOutDelivery={handleFilterOutDelivery}
          />
        </div>
      </>
    )
  );
}
