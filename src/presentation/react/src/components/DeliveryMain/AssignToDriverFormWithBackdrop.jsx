import React from "react";
import AssignToDriverForm from "./AssignToDriverForm";

/**
 * @param {{
 *  delivery: import("../../services/deliveries/deliveries").Delivery
 *  handleBackdropClick: () => void
 * }}
 */
export default function AssignToDriverFormWithBackdrop({
  delivery,
  handleBackdropClick,
}) {
  return (
    <div
      style={{
        position: "absolute",
        top: "55vh",
        right: "50vw",
      }}
    >
      <AssignToDriverForm
        delivery={delivery}
        dissmissForm={handleBackdropClick}
      />
      <div
        onClick={handleBackdropClick}
        style={{
          top: 0,
          left: 0,
          position: "fixed",
          background: `linear-gradient(
                      rgba(255, 255, 255, 0.75) 0%,
                      rgba(255, 255, 255, 0.98) 25%,
                      rgba(255, 255, 255, 0.98) 75%,
                      rgba(255, 255, 255, 0.75) 100%
                    )`,
          height: "100vh",
          width: "100vw",
        }}
      ></div>
    </div>
  );
}
