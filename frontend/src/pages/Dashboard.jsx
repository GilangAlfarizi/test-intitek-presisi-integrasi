import { useState, useEffect } from "react";
import { Button } from "antd";
import axios from "axios";
import DashboardTable from "../components/table"; // Import the DashboardTable component

export default function App() {
	const [products, setProducts] = useState([]);
	const [loading, setLoading] = useState(true);

	useEffect(() => {
		fetchProducts(); // Updated function name
	}, []);

	const fetchProducts = async () => {
		try {
			const response = await axios.get("http://localhost:8080/api/product");
			setProducts(response.data.payload); // Set the products state
		} catch (error) {
			console.error("Error fetching products:", error);
		} finally {
			setLoading(false);
		}
	};

	const deleteProduct = async (id) => {
		try {
			await axios.delete(`http://localhost:8080/api/product/${id}`);
			setProducts((prev) => prev.filter((product) => product.ID !== id)); // Remove the product from the list
		} catch (error) {
			console.error("Error deleting product:", error);
		}
	};

	if (loading) return <div className="text-center mt-10">Loading...</div>;

	return (
		<div className="min-h-screen flex items-center justify-center bg-gray-100">
			<div className="bg-white shadow-lg p-6 rounded-lg min-w-max">
				<h1 className="text-xl font-bold text-center mb-4">
					Product Dashboard
				</h1>
				<DashboardTable products={products} deleteProduct={deleteProduct} />
				<div className="mt-4 flex justify-end">
					<a href="/create-product">
						<Button type="primary" size="large" className="bg-blue-400">
							Create
						</Button>
					</a>
				</div>
			</div>
		</div>
	);
}
