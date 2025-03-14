import { createBrowserRouter } from "react-router-dom";
import Dashboard from "./pages/Dashboard.jsx";
import CreateProduct from "./pages/CreateProduct.jsx";
import UpdateProduct from "./pages/UpdateProduct.jsx";
import App from "./App.jsx";

const router = createBrowserRouter([
	{
		path: "/",
		element: <App />,
	},
	{
		path: "/dashboard",
		element: <Dashboard />,
	},
	{
		path: "/create-product",
		element: <CreateProduct />,
	},
	{
		path: "/update-product/:id",
		element: <UpdateProduct />,
	},
]);

export default router;
