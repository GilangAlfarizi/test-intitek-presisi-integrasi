import { useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import { Space, Table } from 'antd';

export default function CreateProduct() {
	const [name, setName] = useState("");
	const [sku, setSku] = useState("");
	const [quantity, setQuantity] = useState(0);
	const [location, setLocation] = useState("");
	const [status, setStatus] = useState("");

	const navigate = useNavigate();

	const handleSubmit = async (e) => {
		e.preventDefault();
		try {
			const response = await axios.post("http://localhost:8080/api/product", {
				name,
				sku,
				quantity,
				location,
				status,
			});
			if (response.status === 200) {
				navigate("/dashboard"); // Redirect to the dashboard after successful creation
			}
		} catch (error) {
			console.error("Error creating product:", error);
		}
	};

	return (
		<div className="min-h-screen flex items-center justify-center bg-gray-100">
			<div className="bg-white shadow-lg p-6 rounded-lg w-96">
				<h1 className="text-xl font-bold text-center mb-4">Create Product</h1>
				<form onSubmit={handleSubmit}>
					<div className="mb-4">
						<label className="block text-gray-700">Name</label>
						<input
							type="text"
							value={name}
							onChange={(e) => setName(e.target.value)}
							className="border p-2 w-full rounded-md"
							required
						/>
					</div>
					<div className="mb-4">
						<label className="block text-gray-700">SKU</label>
						<input
							type="text"
							value={sku}
							onChange={(e) => setSku(e.target.value)}
							className="border p-2 w-full rounded-md"
							required
						/>
					</div>
					<div className="mb-4">
						<label className="block text-gray-700">Quantity</label>
						<input
							type="number"
							value={quantity}
							onChange={(e) => setQuantity(parseInt(e.target.value, 10))}
							className="border p-2 w-full rounded-md"
							required
						/>
					</div>
					<div className="mb-4">
						<label className="block text-gray-700">Location</label>
						<input
							type="text"
							value={location}
							onChange={(e) => setLocation(e.target.value)}
							className="border p-2 w-full rounded-md"
							required
						/>
					</div>
					<div className="mb-4">
						<label className="block text-gray-700">Status</label>
						<select
							value={status}
							onChange={(e) => setStatus(e.target.value)}
							className="border p-2 w-full rounded-md"
							required>
							<option value="">Select Status</option>
							<option value="active">Active</option>
							<option value="inactive">Inactive</option>
						</select>
					</div>
					<a href="/">
						<button
							on
							type="submit"
							className="w-full px-4 py-2 bg-green-500 text-white rounded-md">
							Create
						</button>
					</a>
				</form>
			</div>
		</div>

		
	);
}
