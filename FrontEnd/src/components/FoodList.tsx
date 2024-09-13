"use client";
import React, { useState } from "react";
import FoodItem from "./FoodItem";
import CurrentOrderPopupList from "./CurrentOrderPopupList";
import OrderHistoryPopup from "./OrderHistoryPopup";

type FoodCounts = {
  [key: string]: number;
};

export default function FoodList() {
  const [totalAmount, setTotalAmount] = useState(0);
  const [foodCounts, setFoodCounts] = useState<FoodCounts>({
    Bocchi: 0,
    Kita: 0,
    Ryo: 0,
    Nijika: 0,
  });
  const [isOrderListPopupOpen, setIsOrderListPopupOpen] = useState(false);
  const [isOrderHistoryPopupOpen, setIsOrderHistoryPopupOpen] = useState(false);
  const [orderHistory, setOrderHistory] = useState<any[]>([]);
  const [notificationVisible, setNotificationVisible] = useState(false);

  const handleAmountChange = (foodName: string, delta: number) => {
    setTotalAmount(totalAmount + delta);
    setFoodCounts((prevCounts) => ({
      ...prevCounts,
      [foodName]: prevCounts[foodName] + delta,
    }));
  };

  const handleRemoveItem = (foodName: string) => {
    setFoodCounts((prevCounts) => ({
      ...prevCounts,
      [foodName]: 0,
    }));
  };

  const handleSubmitOrder = async () => {
    const orderData = {
      FoodList: foodCounts,
    };

    try {
      const token = localStorage.getItem("token");

      const response = await fetch("http://localhost:8080/api/v1/orders", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(orderData),
      });

      if (!response.ok) {
        throw new Error("Failed to submit order");
      }
      const result = await response.json();
      console.log("Order submitted:", result);

      // Show notification
      setNotificationVisible(true);

      // Hide notification after 1 second
      setTimeout(() => {
        setNotificationVisible(false);
      }, 1000);

      // Reset food counts and total amount after successful order submission
      const resetCounts = Object.keys(foodCounts).reduce((acc, key) => {
        acc[key] = 0;
        return acc;
      }, {} as FoodCounts);

      setFoodCounts(resetCounts);
      setTotalAmount(0);
    } catch (error) {
      console.error("Error submitting order:", error);
    }
  };

  const handleGetOrderHistory = async (token: string) => {
    try {
      const response = await fetch("http://localhost:8080/api/v1/auth/orders", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        throw new Error("Failed to fetch order history.");
      }

      const orders = await response.json();
      return orders;
    } catch (e) {
      console.error("Error fetching order history:", e);
      return [];
    }
  };

  const handleShowOrderHistory = async () => {
    const token = localStorage.getItem("token");

    if (!token) {
      console.error("No token found.");
      return;
    }

    const orders = await handleGetOrderHistory(token);
    setOrderHistory(orders);
    setIsOrderHistoryPopupOpen(true);
  };

  return (
    <div className="relative flex flex-col justify-center items-center bg-gray-400/60 backdrop-blur-md rounded-lg px-6 py-6 border border-gray-300 shadow-md">
      {Object.entries(foodCounts).map(([foodName]) => (
        <FoodItem
          key={foodName}
          foodName={foodName}
          imgSrc={`/img/${foodName.toLowerCase()}.jpg`}
          amount={foodCounts[foodName]}
          onAmountChange={(delta) => handleAmountChange(foodName, delta)}
        />
      ))}

      <div className="flex flex-col items-center">
        <div>
          <button
            className="w-[150px] h-[40px] bg-green-500 rounded-lg mt-3 mr-2 text-white font-bold transition-transform transform hover:scale-105 active:scale-95 focus:outline-none"
            onClick={handleSubmitOrder}
          >
            Submit Order
          </button>
          <button
            className="w-[150px] h-[40px] bg-sky-500 rounded-lg mt-3 ml-2 text-white font-bold transition-transform transform hover:scale-105 active:scale-95 focus:outline-none"
            onClick={() => setIsOrderListPopupOpen(true)}
          >
            Order List
          </button>
        </div>
        <button
          className="block w-full h-[30px] bg-slate-400 rounded-lg mt-3 mx-2 text-white font-bold transition-transform transform hover:scale-105 active:scale-95 focus:outline-none"
          onClick={handleShowOrderHistory}
        >
          Order History
        </button>
      </div>

      {isOrderListPopupOpen && (
        <CurrentOrderPopupList
          foodCounts={foodCounts}
          onClose={() => setIsOrderListPopupOpen(false)}
          onRemoveItem={handleRemoveItem}
        />
      )}

      {isOrderHistoryPopupOpen && (
        <OrderHistoryPopup
          orders={orderHistory} // Pass fetched orders to the popup
          onClose={() => setIsOrderHistoryPopupOpen(false)}
        />
      )}

      {notificationVisible && (
        <div className="fixed top-4 left-1/2 transform -translate-x-1/2 bg-green-500 text-white px-4 py-2 rounded-lg shadow-lg">
          Order submitted successfully!
        </div>
      )}
    </div>
  );
}
