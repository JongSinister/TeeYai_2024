"use client";
import { useState, useEffect } from "react";

type FoodList = {
  [foodName: string]: number;
};

type Order = {
  OrderID: string;
  UserID: string;
  UserName: string;
  FoodList: FoodList;
  CreatedAt: string;
};

interface OrderCardProps {
  order: Order;
  onDelete: (orderId: string) => void;
}

export default function OrderCard({ order, onDelete }: OrderCardProps) {
  const [isOpen, setIsOpen] = useState(false);
  const [selectedItems, setSelectedItems] = useState<Set<string>>(new Set());
  const [allChecked, setAllChecked] = useState(false);

  const togglePopup = () => {
    setIsOpen(!isOpen);
  };

  const handleCheckboxChange = (foodName: string) => {
    setSelectedItems((prevSelectedItems) => {
      const newSelectedItems = new Set(prevSelectedItems);
      if (newSelectedItems.has(foodName)) {
        newSelectedItems.delete(foodName);
      } else {
        newSelectedItems.add(foodName);
      }
      return newSelectedItems;
    });
  };

  useEffect(() => {
    const totalItems = Object.keys(order.FoodList).length;
    const checkedItems = selectedItems.size;
    setAllChecked(totalItems === checkedItems);
  }, [selectedItems, order.FoodList]);

  const handleFinish = async () => {
    try {
      const token = localStorage.getItem("token");

      const response = await fetch(
        `http://localhost:8080/api/v1/orders/${order.OrderID}`,
        {
          method: "DELETE",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
          },
        }
      );

      if (!response.ok) {
        throw new Error("Failed to delete order.");
      }

      onDelete(order.OrderID);
      togglePopup();
    } catch (error) {
      console.error("An error occurred while deleting the order:", error);
    }
  };

  const formattedDate = new Date(order.CreatedAt).toLocaleDateString(undefined, {
    year: "numeric",
    month: "numeric",
    day: "numeric",
  });
  const formattedTime = new Date(order.CreatedAt).toLocaleTimeString(undefined, {
    hour: "2-digit",
    minute: "2-digit",
  });

  return (
    <div className="bg-white rounded-md p-4 shadow-md w-[250px] flex flex-col items-center relative">
      <div className="text-sm font-semibold">{order.UserName}</div>
      <div className="text-xs text-center mb-2">
        OrderDate:{formattedTime} , {formattedDate}
      </div>
      {allChecked && (
        <div className="absolute top-1 left-1 w-2 h-2 bg-green-500 rounded-full"></div>
      )}
      <button
        className="p-2 bg-blue-500 text-white rounded-md"
        onClick={togglePopup}
      >
        View Order
      </button>

      {isOpen && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center z-50">
          <div className="bg-white p-6 rounded-md shadow-md w-[300px] relative flex flex-col">
            <button
              className="absolute top-2 right-2 p-2 bg-red-500 text-white rounded-md"
              onClick={togglePopup}
            >
              Close
            </button>
            <h2 className="text-xl font-bold mb-4">Order Details</h2>
            <ul className="flex-1 overflow-y-auto">
              {Object.entries(order.FoodList).map(([foodName, quantity]) => (
                <li key={foodName} className="flex items-center">
                  <input
                    type="checkbox"
                    checked={selectedItems.has(foodName)}
                    onChange={() => handleCheckboxChange(foodName)}
                    className="mr-2"
                  />
                  {foodName}: {quantity}
                </li>
              ))}
            </ul>
            <button
              className="absolute bottom-2 right-2 p-2 bg-green-500 text-white rounded-md"
              onClick={handleFinish}
            >
              Finish
            </button>
          </div>
        </div>
      )}
    </div>
  );
}
