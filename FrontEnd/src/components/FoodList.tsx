import FoodItem from "./FoodItem";

export default function FoodList() {
  return (
    <div className="flex flex-col justify-center items-center bg-indigo-800 rounded-lg px-4">
      <FoodItem foodName="Bocchi" imgSrc="/img/bocchides.jpg" />
      <FoodItem foodName="Kita" imgSrc="/img/kitathisisbass.jpg" />
      <FoodItem foodName="Ryo" imgSrc="/img/ryothumbup.jpg" />
      <FoodItem foodName="Nijika" imgSrc="/img/nijikaleaf.jpg" />
    </div>
  );
}
