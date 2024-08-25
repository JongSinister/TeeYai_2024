import styles from "../page.module.css";
import FoodList from "@/components/FoodList";

export default function UserPage() {
    return (
        <div className={styles.main}>
            <FoodList/>
        </div>
    );
}