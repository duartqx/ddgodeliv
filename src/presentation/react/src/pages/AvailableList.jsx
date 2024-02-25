import React, { useContext, useEffect, useState } from "react";
import SideBarList from "../components/SideBarList/SideBarList";
import DeliveryCard from "../components/DeliveryCard";
import DeliveryMainPending from "../components/DeliveryMain/DeliveryMainPending";
import { TitleContext } from "../middlewares/TitleContext";
import * as deliveryService from "../services/deliveries/deliveries";

export default function AvailableList() {
  const [pendingDeliveries, setPendingDeliveries] = useState(
    /** @type {deliveryService.Delivery[]} */ ([]),
  );
  const [filter, setFilter] = useState("");
  const [selectedDelivery, setSelectedDelivery] = useState(0);
  const { setTitle } = useContext(TitleContext);

  useEffect(() => {
    setTitle("Available");
    deliveryService
      .getPendingDeliveries()
      .then((deliveries) => setPendingDeliveries(deliveries));
  }, []);

  const filteredDeliveries = filter
    ? pendingDeliveries.filter((d) =>
        d.loadout.toLowerCase().includes(filter.toLowerCase()),
      )
    : pendingDeliveries;

  const handleFilterOutDelivery = (delivery) =>
    setPendingDeliveries(
      pendingDeliveries.filter((d) => d.id !== delivery.id),
    );

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
      <div className="d-flex flex-grow-1">
        <SideBarList
          listing={deliveriesCards}
          filterValue={filter}
          filterOnChangeHandler={(e) => setFilter(e.target.value)}
        />
        {filteredDeliveries[selectedDelivery] && (
          <DeliveryMainPending
            delivery={filteredDeliveries[selectedDelivery]}
            handleFilterOutDelivery={handleFilterOutDelivery}
          />
        )}
      </div>
    </>
  );
}
