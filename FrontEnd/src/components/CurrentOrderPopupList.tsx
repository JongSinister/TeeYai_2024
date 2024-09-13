import React from 'react';

type FoodCounts = {
  [key: string]: number;
};

type CurrentOrderPopupListProps = {
  foodCounts: FoodCounts;
  onClose: () => void;
  onRemoveItem: (foodName: string) => void;
};

export default function CurrentOrderPopupList({
  foodCounts,
  onClose,
  onRemoveItem,
}: CurrentOrderPopupListProps) {
  const hasOrders = Object.values(foodCounts).some(count => count > 0);

  return (
    <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50">
      <div className="relative bg-white p-6 rounded-lg shadow-lg max-w-md w-full">
        <button
          className="absolute top-4 right-4 w-10 h-10 flex items-center justify-center bg-red-600 text-white rounded-md shadow-md transition-transform transform hover:scale-110 hover:bg-red-700 focus:outline-none"
          onClick={onClose}
        >
          &times;
        </button>
        <h2 className="text-2xl font-semibold mb-4">Current Orders</h2>
        {hasOrders ? (
          <ul className="list-disc pl-5">
            {Object.entries(foodCounts)
              .filter(([_, count]) => count > 0)
              .map(([foodName, count]) => (
                <li key={foodName} className="mb-3 flex items-center justify-between border-b border-gray-200 pb-2">
                  <span className="text-lg">{foodName}: {count}</span>
                  <button
                    className="text-red-500 hover:text-red-700 font-semibold"
                    onClick={() => onRemoveItem(foodName)}
                  >
                    Delete
                  </button>
                </li>
              ))}
          </ul>
        ) : (
          <p className="text-gray-600">You haven't ordered any food yet.</p>
        )}
      </div>
    </div>
  );
}
