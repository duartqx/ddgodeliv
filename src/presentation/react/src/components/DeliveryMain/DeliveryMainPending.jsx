import React from "react";
import DeliveryMainHeader from "./DeliveryMainHeader"

/**
 * @param {{
 * delivery: import("../../services/deliveries/deliveries").Delivery
 * }}
 * props */
export default function DeliveryMainPending({ delivery }) {
  return (
    delivery && (
      <>
        <div
          className="d-flex flex-column mx-4 flex-grow-1"
          style={{ maxHeight: "calc(100vh - 4rem)" }}
        >
          <DeliveryMainHeader delivery={delivery} />
        </div>
      </>
    )
  );
}
