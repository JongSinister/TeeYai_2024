"use client";
import React, { useState } from "react";
import FoodItem from "./FoodItem";

type FoodCounts = {
  Bocchi: number;
  Kita: number;
  Ryo: number;
  Nijika: number;
};

export default function FoodList() {
  const [totalAmount, setTotalAmount] = useState(0);
  const [foodCounts, setFoodCounts] = useState<FoodCounts>({
    Bocchi: 0,
    Kita: 0,
    Ryo: 0,
    Nijika: 0,
  });

  const handleAmountChange = (foodName: keyof FoodCounts, delta: number) => {
    setTotalAmount(totalAmount + delta);
    setFoodCounts({
      ...foodCounts,
      [foodName]: foodCounts[foodName] + delta,
    });
  };

  return (
    <div className="flex flex-col justify-center items-center bg-indigo-800 rounded-lg px-4">
      <h1 className="text-white text-2xl mb-4">Total Amount: {totalAmount}</h1>
      <div className="text-white mb-4">
        {Object.entries(foodCounts).map(([foodName, count]) => (
          <p key={foodName}>
            {foodName}: {count}
          </p>
        ))}
      </div>
      <FoodItem foodName="Bocchi" imgSrc="/img/bocchides.jpg" amount={0} onAmountChange={(delta) => handleAmountChange('Bocchi', delta)} />
      <FoodItem foodName="Kita" imgSrc="/img/kitathisisbass.jpg" amount={0} onAmountChange={(delta) => handleAmountChange('Kita', delta)} />
      <FoodItem foodName="Ryo" imgSrc="/img/ryothumbup.jpg" amount={0} onAmountChange={(delta) => handleAmountChange('Ryo', delta)} />
      <FoodItem foodName="Nijika" imgSrc="/img/nijikaleaf.jpg" amount={0} onAmountChange={(delta) => handleAmountChange('Nijika', delta)} />
    </div>
  );
}
