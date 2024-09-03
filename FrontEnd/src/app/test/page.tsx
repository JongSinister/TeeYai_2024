"use client";
import React, { useState } from "react";

type Order = {
  OrderID: string;
  UserID: string;
  FoodList: {
    [key: string]: number; // Allows for dynamic food names
  };
  CreatedAt: string;
};

export default function TestPage() {
  const [orders, setOrders] = useState<Order[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  // Form state
  const [userID, setUserID] = useState<string>("");
  const [foodList, setFoodList] = useState<{ [key: string]: number }>({
    Bocchi: 0,
    Kita: 0,
    Ryo: 0,
    Nijika: 0,
  });

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setUserID(e.target.value);
  };

  const handleFoodChange = (foodName: string, value: number) => {
    setFoodList(prevFoodList => ({
      ...prevFoodList,
      [foodName]: value,
    }));
  };

  const fetchOrders = async () => {
    setIsLoading(true);
    setError(null);
    try {
      const response = await fetch("http://localhost:8080/api/v1/orders");
      if (!response.ok) {
        throw new Error("Failed to fetch orders");
      }
      const data = await response.json();
      setOrders(data);
    } catch (error) {
      setError("Error fetching orders");
      console.error("Error fetching orders:", error);
    } finally {
      setIsLoading(false);
    }
  };

  const handleSubmitOrder = async (e: React.FormEvent) => {
    e.preventDefault();
    const orderData = {
      UserID: userID,
      FoodList: foodList,
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
      // Clear form after successful submission
      setUserID("");
      setFoodList({
        Bocchi: 0,
        Kita: 0,
        Ryo: 0,
        Nijika: 0,
      });
    } catch (error) {
      setError("Error submitting order");
      console.error("Error submitting order:", error);
    }
  };

  return (
    <div className="flex flex-col items-center p-4">
      <h1 className="text-2xl mb-4">Orders</h1>
      <button
        className="w-[200px] h-[50px] bg-blue-500 rounded-lg text-white font-bold mb-4"
        onClick={fetchOrders}
      >
        Fetch Orders
      </button>

      <h2 className="text-xl mb-4">Submit New Order</h2>
      <form onSubmit={handleSubmitOrder} className="flex flex-col items-center mb-4">
        <label className="mb-2">
          User ID:
          <input
            type="text"
            value={userID}
            onChange={handleInputChange}
            className="ml-2 p-1 border rounded"
            required
          />
        </label>
        {Object.keys(foodList).map(foodName => (
          <label key={foodName} className="mb-2">
            {foodName}:
            <input
              type="number"
              value={foodList[foodName]}
              onChange={(e) => handleFoodChange(foodName, Number(e.target.value))}
              className="ml-2 p-1 border rounded"
              min="0"
              required
            />
          </label>
        ))}
        <button
          type="submit"
          className="w-[200px] h-[50px] bg-green-500 rounded-lg text-white font-bold"
        >
          Submit Order
        </button>
      </form>

      {isLoading && <p className="text-blue-500">Loading...</p>}
      {error && <p className="text-red-500">{error}</p>}

      <ul>
        {orders.map((order) => (
          <li key={order.OrderID} className="mb-2 p-2 border rounded">
            <h2 className="text-xl">Order ID: {order.OrderID}</h2>
            <p>User ID: {order.UserID}</p>
            <p>Created At: {new Date(order.CreatedAt).toLocaleString()}</p>
            <h3 className="font-bold mt-2">Food List:</h3>
            <ul>
              {Object.entries(order.FoodList).map(([foodName, quantity]) => (
                <li key={foodName}>
                  {foodName}: {quantity}
                </li>
              ))}
            </ul>
          </li>
        ))}
      </ul>
    </div>
  );
}
