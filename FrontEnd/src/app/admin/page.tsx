import styles from "../page.module.css";
import OrderList from "@/components/OrderList";

export default function Page() {
    return (
        <div className={styles.mainadmin}>
            <h1 className="text-2xl font-bold mt-[20px] mb-[10px]">Admin Page</h1>
            <OrderList />
        </div>
    );
}