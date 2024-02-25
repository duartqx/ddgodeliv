import React, { useContext, useEffect, useState } from "react";
import SideBarList from "../components/SideBarList/SideBarList";
import DeliveryCard from "../components/DeliveryCard";
import DeliveryMainPending from "../components/DeliveryMain/DeliveryMainPending";
import { TitleContext } from "../middlewares/TitleContext";
import PendingDeliveriesContextProvider, {
  PendingDeliveriesContext,
} from "../components/DeliveryMain/PendingDeliveriesContext";

function AvailableListWithContext() {
  const [filter, setFilter] = useState("");
  const [selectedDelivery, setSelectedDelivery] = useState(0);
  const { setTitle } = useContext(TitleContext);
  const { getFiltered } = useContext(PendingDeliveriesContext);

  useEffect(() => {
    setTitle("Available");
  }, []);

  const filteredDeliveries = getFiltered(filter);

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
          />
        )}
      </div>
    </>
  );
}

export default function AvailableList() {
  return (
    <PendingDeliveriesContextProvider>
      <AvailableListWithContext />
    </PendingDeliveriesContextProvider>
  );
}
