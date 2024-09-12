"use client";
import { useState, useEffect } from "react";
import OrderCard from "./OrderCard";

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

export default function OrderList() {
  const [orders, setOrders] = useState<Order[]>([]);
  const [error, setError] = useState("");

  const fetchOrders = async () => {
    try {
      const token = localStorage.getItem("token");

      const response = await fetch("http://localhost:8080/api/v1/orders", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        throw new Error("Failed to fetch orders.");
      }

      const data: Order[] = await response.json();
      setOrders(data);
    } catch (error) {
      setError(
        "An error occurred while fetching orders. Please try again later."
      );
    }
  };

  useEffect(() => {
    fetchOrders();
  }, []);

  const handleDelete = (orderId: string) => {
    setOrders((prevOrders) =>
      prevOrders.filter((order) => order.OrderID !== orderId)
    );
  };

  return (
    <div className="overflow-x-auto w-[80%] bg-slate-300">
      {error && <div>{error}</div>}
      <ul className="flex space-x-4 p-4 w-max">
        {orders
          .map((order) => (
            <li key={order.OrderID} className="flex-shrink-0">
              <OrderCard order={order} onDelete={handleDelete} />
            </li>
          ))}
      </ul>
    </div>
  );
}
