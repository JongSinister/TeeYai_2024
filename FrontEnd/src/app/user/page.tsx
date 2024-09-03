import styles from "../page.module.css";
import FoodList from "@/components/FoodList";

export default function UserPage() {
    return (
        <div className={styles.mainuser}>
            <h1 className="text-2xl font-extrabold mt-[20px] mb-[10px]">Food List</h1>
            <FoodList/>
        </div>
    );
}