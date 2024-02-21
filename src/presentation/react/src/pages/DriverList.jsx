import React, { useContext, useEffect, useState } from "react";
import SideBarList from "../components/SideBarList/SideBarList";
import DriverCard from "../components/DriverCard";
import DriverCardForm from "../components/CardForm/Driver/DriverCardForm";
import DriverMain from "../components/DriverMain/DriverMain";
import * as driverService from "../services/driver/driver";
import { TitleContext } from "../middlewares/TitleContext";

export default function DriverList() {
  const [drivers, setDrivers] = useState(
    /** @type {import("../services/driver/driver").Driver[]} */ ([]),
  );
  const [filterDriver, setFilterDriver] = useState("");
  const [selectedDriver, setSelectedDriver] = useState(0);
  const { setTitle } = useContext(TitleContext)

  useEffect(() => {

    setTitle("Drivers")

    driverService.companyDrivers().then((drivers) => setDrivers(drivers));
  }, []);

  const deleteClickHandler = async (/** @type {number} */ id) => {
    if (await driverService.deleteDriver({ id })) {
      setDrivers(drivers.filter((d) => d.id !== id));
      setSelectedDriver(0);
    }
  };

  const filteredDrivers = filterDriver
    ? drivers.filter((d) =>
        d.user.name.toLowerCase().includes(filterDriver.toLowerCase()),
      )
    : drivers;

  const driversCards = filteredDrivers.map((d, idx) => (
    <DriverCard
      onClickHandler={() => setSelectedDriver(idx)}
      selected={selectedDriver === idx}
      driver={d}
      key={`driver__${d.id}__${d.user.email}`}
    />
  ));

  return (
    <>
      <div className="d-flex flex-grow-1">
        <DriverCardForm
        appendDriver={(driver) => setDrivers(drivers.concat(driver))}
        />
        <SideBarList
        listing={driversCards}
        filterValue={filterDriver}
        filterOnChangeHandler={(e) => setFilterDriver(e.target.value)}
        />
        {filteredDrivers[selectedDriver] && (
          <DriverMain
            driver={filteredDrivers[selectedDriver]}
            deleteHandler={deleteClickHandler}
          />
        )}
      </div>
    </>
  );
}
