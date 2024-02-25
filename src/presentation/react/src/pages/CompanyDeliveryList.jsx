import React, { useContext, useEffect, useState } from "react";
import SideBarList from "../components/SideBarList/SideBarList";
import * as deliveryService from "../services/deliveries/deliveries";
import DeliveryCard from "../components/DeliveryCard";
import { TitleContext } from "../middlewares/TitleContext";

export default function CompanyDeliveryList() {
  const [deliveries, setDeliveries] = useState(
    /** @type {import("../services/deliveries/deliveries").Delivery[]} */ ([]),
  );
  const [filterDeliveries, setFilterDeliveries] = useState("");
  const [selectedDelivery, setSelectedDelivery] = useState(0);
  const { setTitle } = useContext(TitleContext);

  useEffect(() => {
    setTitle("Company Deliveries");

    deliveryService
      .getByCompanyId()
      .then((deliveries) => setDeliveries(deliveries));
  }, []);

  const filteredDeliveries = filterDeliveries
    ? deliveries.filter((d) =>
        d.loadout.toLowerCase().includes(filterDeliveries.toLowerCase()),
      )
    : deliveries;

  const deliveriesCards = filteredDeliveries.map((d, idx) => (
    <DeliveryCard
      delivery={d}
      selected={selectedDelivery === idx}
      key={`delivery__${d.id}__${d.loadout.replaceAll(" ", "")}`}
      onClickHandler={() => setSelectedDelivery(idx)}
    />
  ));

  return (
    <>
      <SideBarList
        listing={deliveriesCards}
        filterValue={filterDeliveries}
        filterOnChangeHandler={(e) => setFilterDeliveries(e.target.value)}
      />
    </>
  );
}
