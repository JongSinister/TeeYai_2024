"use client";
import React, { useState } from "react";
import Image from "next/image";

export default function FoodItem({
  foodName,
  imgSrc,
  amount,
  onAmountChange,
}: {
  foodName: string;
  imgSrc: string;
  amount: number;
  onAmountChange: (delta: number) => void;
}) {
  const [count, setCount] = useState(amount);

  const increaseCount = () => {
    setCount(prevCount => prevCount + 1);
  };

  const decreaseCount = () => {
    if (count > 0) {
      setCount(prevCount => prevCount - 1);
    }
  };

  const handleAddToOrder = () => {
    onAmountChange(count);
    setCount(0);
  };

  return (
    <div className="w-[400px] h-[150px] rounded-lg bg-white flex items-center my-4 shadow-lg">
      <div className="flex items-center w-full px-4">
        <Image src={imgSrc} alt={foodName} width={100} height={100} />
        <div className="text-base mx-10 w-[50px] flex justify-center items-center">
          <p>{foodName}</p>
        </div>
        <div className="flex flex-col">
          <div className="flex items-center mt-2">
            <button
              className="w-[30px] h-[30px] bg-blue-300 rounded-md mx-2 font-semibold transition-transform transform hover:scale-110 active:scale-90 focus:outline-none"
              onClick={decreaseCount}
            >
              -
            </button>
            <p className="mx-1 w-[10px] text-center">{count}</p>
            <button
              className="w-[30px] h-[30px] bg-blue-300 rounded-md mx-2 font-semibold transition-transform transform hover:scale-110 active:scale-90 focus:outline-none"
              onClick={increaseCount}
            >
              +
            </button>
          </div>
          <button
            className="w-[100px] h-[30px] bg-blue-300 rounded-md mt-2 mx-auto font-semibold px-1 text-sm transition-transform transform hover:scale-105 active:scale-95 focus:outline-none"
            onClick={handleAddToOrder}
          >
            Add to Order
          </button>
        </div>
      </div>
    </div>
  );
}
