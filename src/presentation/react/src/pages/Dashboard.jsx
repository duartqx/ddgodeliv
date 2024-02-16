import React, { useState } from "react";
import DriverList from "../components/SideBarList/Driver/DriverList";
import VehiclesList from "../components/SideBarList/Vehicle/VehiclesList";

/**
 * @typedef {{
 *  selected: "drivers"|"vehicles"|"pending"|"taken"
 *  isSelected: (field: string) => boolean
 * }} SideBar
 */

export default function Dashboard() {
    return <main className="d-flex"> </main>;
}
