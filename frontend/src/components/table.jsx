import { Button, Table, Space, message, Popconfirm } from "antd";
import { useNavigate } from "react-router-dom";

export default function DashboardTable({ products, deleteProduct }) {
	const navigate = useNavigate();

	const confirm = (e) => {
		console.log(e);
		message.success("Click on Yes");
	};
	const cancel = (e) => {
		console.log(e);
		message.error("Click on No");
	};

	const columns = [
		{
			title: "Name",
			dataIndex: "Name",
			key: "Name",
		},
		{
			title: "Sku",
			dataIndex: "Sku",
			key: "Sku",
		},
		{
			title: "Quantity",
			dataIndex: "Quantity",
			key: "Quantity",
		},
		{
			title: "Location",
			dataIndex: "Location",
			key: "Location",
		},
		{
			title: "Status",
			dataIndex: "Status",
			key: "Status",
		},
		{
			title: "Action",
			key: "action",
			render: (text, record) => (
				<Space size="middle">
					<Button
						color="gold"
						variant="solid"
						onClick={() => navigate(`/update-product/${record.ID}`)}>
						Update
					</Button>
					<Popconfirm
						title="Delete the task"
						description="Are you sure to delete this task?"
						onConfirm={confirm}
						onCancel={cancel}
						okText="Yes"
						cancelText="No">
						<Button danger onClick={() => deleteProduct(record.ID)}>
							Delete
						</Button>
					</Popconfirm>
				</Space>
			),
		},
	];

	return <Table columns={columns} dataSource={products} />;
}
