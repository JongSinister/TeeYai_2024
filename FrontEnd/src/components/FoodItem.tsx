"use client";
import React, { useState } from "react";
import Image from "next/image";

export default function FoodItem({
  foodName,
  imgSrc,
}: {
  foodName: string;
  imgSrc: string;
}) {
  const [count, setCount] = useState(0); // Initialize state with count set to 0

  const increaseCount = () => setCount(count + 1);
  const decreaseCount = () => {
    if (count > 0) {
      // Prevent count from going below 0
      setCount(count - 1);
    }
  };

  return (
    <div className="w-[400px] h-[150px] rounded-lg bg-white flex items-center my-4 shadow-lg">
      <div className="flex items-center w-full px-4">
        <Image src={imgSrc} alt="Product Picture" width={100} height={100} />
        <div className="text-base mx-10 w-[50px] flex justify-center items-center">
          <p>{foodName}</p>
        </div>
        <div className="flex flex-col">
          <div className="flex items-center mt-2">
            <button
              className="w-[30px] h-[30px] bg-blue-300 rounded-md mx-2 font-semibold"
              onClick={decreaseCount}
            >
              -
            </button>
            <p className="mx-1">{count}</p>
            <button
              className="w-[30px] h-[30px] bg-blue-300 rounded-md mx-2 font-semibold"
              onClick={increaseCount}
            >
              +
            </button>
          </div>
          <button className="w-[100px] h-[30px] bg-blue-300 rounded-md mt-2 mx-auto font-semibold px-1 text-sm">
            Add to Order
          </button>
        </div>
      </div>
    </div>
  );
}
