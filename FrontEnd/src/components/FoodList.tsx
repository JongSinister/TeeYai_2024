"use client";
import React, { useState } from "react";
import FoodItem from "./FoodItem";
import CurrentOrderPopupList from "./CurrentOrderPopupList";

type FoodCounts = {
  [key: string]: number; // Allows for dynamic food names
};

export default function FoodList() {
  const [totalAmount, setTotalAmount] = useState(0);
  const [foodCounts, setFoodCounts] = useState<FoodCounts>({
    Bocchi: 0,
    Kita: 0,
    Ryo: 0,
    Nijika: 0,
  });
  const [isPopupOpen, setIsPopupOpen] = useState(false);

  const handleAmountChange = (foodName: string, delta: number) => {
    setTotalAmount(totalAmount + delta);
    setFoodCounts((prevCounts) => ({
      ...prevCounts,
      [foodName]: prevCounts[foodName] + delta,
    }));
  };

  const handleSubmitOrder = async () => {
    const orderData = {
      FoodList: foodCounts,
      CreatedAt: new Date().toISOString(),
    };

    try {
      const response = await fetch("http://localhost:8080/api/v1/orders", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(orderData),
      });

      if (!response.ok) {
        throw new Error("Failed to submit order");
      }
      const result = await response.json();
      console.log("Order submitted:", result);

      /* Reset food counts and total amount after successful order submission */
      const resetCounts = Object.keys(foodCounts).reduce((acc, key) => {
        acc[key] = 0;
        return acc;
      }, {} as FoodCounts);

      setFoodCounts(resetCounts);
      setTotalAmount(0);
      /* Error handling for fetch request */
    } catch (error) {
      console.error("Error submitting order:", error);
    }
  };

  const handleRemoveItem = (foodName: string) => {
    setFoodCounts((prevCounts) => ({
      ...prevCounts,
      [foodName]: 0, // Set the item count to 0 to effectively remove it
    }));
  };

  return (
    <div className="flex flex-col justify-center items-center bg-indigo-800 rounded-lg px-4 py-4">
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
            onClick={() => setIsPopupOpen(true)}
          >
            Order List
          </button>
        </div>
        <button className="block w-full h-[30px] bg-slate-400 rounded-lg mt-3 mx-2 text-white font-bold transition-transform transform hover:scale-105 active:scale-95 focus:outline-none">
          Order History
        </button>
      </div>
      {isPopupOpen && (
        <CurrentOrderPopupList
          foodCounts={foodCounts}
          onClose={() => setIsPopupOpen(false)}
          onRemoveItem={handleRemoveItem}
        />
      )}
    </div>
  );
}
