import React from "react";
import { DeliveryStatus } from "../../domains/deliveries/status";
import RoundImage from "../RoundImage";

/** @param {{ driver: import("../../services/driver/driver").Driver }} props */
export default function DriverMainCard({ driver }) {
    return (
        <div className="card bg-body-tertiary">
            <div className="card-body d-flex">
                <RoundImage size="8rem" />
                <div className="p-5">
                    <div>{driver.user.email}</div>
                    <div>{DeliveryStatus[driver.status]}</div>
                </div>
            </div>
        </div>
    );
}
