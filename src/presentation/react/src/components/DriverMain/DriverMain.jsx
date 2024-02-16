import React from "react";
import DriverMainHeader from "./Header/DriverMainHeader";
import DriverMainCard from "./DriverMainCard";

/**
 * @param {{
 * driver: import("../../services/driver/driver").Driver
 * deleteHandler: (id: number) => void
 * }}
 * props */
export default function DriverMain({ driver, deleteHandler }) {
    return (
        driver && (
            <>
                <div className="d-flex flex-column flex-grow-1">
                    <div className="p-5">
                        <DriverMainHeader
                            driver={driver}
                            deleteHandler={deleteHandler}
                        />
                        <DriverMainCard driver={driver} />
                    </div>
                </div>
            </>
        )
    );
}
