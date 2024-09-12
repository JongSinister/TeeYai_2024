import React from "react";

type OrderHistoryPopupProps = {
  orders: Array<{ [key: string]: number }>;
  onClose: () => void;
};

const OrderHistoryPopup: React.FC<OrderHistoryPopupProps> = ({ orders, onClose }) => {
  return (
    <div className="fixed inset-0 flex justify-center items-center bg-black bg-opacity-50 z-50">
      <div className="relative bg-white rounded-lg p-4 w-11/12 md:w-1/2 lg:w-1/3">
        <h2 className="text-lg font-bold mb-4">Order History</h2>
        <button
          className="absolute top-2 right-2 bg-red-500 text-white rounded p-1 w-10 h-10 flex justify-center items-center font-bold"
          onClick={onClose}
        >
          X
        </button>
        {orders.length === 0 ? (
          <p>No orders found.</p>
        ) : (
          orders.map((order, index) => (
            <div key={index} className="mb-4">
              <h3 className="text-md font-semibold">Order {index + 1}</h3>
              {Object.entries(order).map(([foodItem, quantity]) => (
                quantity > 0 && (
                  <p key={foodItem}>
                    {foodItem}: {quantity}
                  </p>
                )
              ))}
            </div>
          ))
        )}
      </div>
    </div>
  );
};

export default OrderHistoryPopup;
