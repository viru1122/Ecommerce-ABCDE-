import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { toast, ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { addToCart, createOrder, getCarts, getItems, getOrders } from '../api';

function ItemList() {
  const [items, setItems] = useState([]);
  const navigate = useNavigate();
  const token = localStorage.getItem('token');

  useEffect(() => {
    if (!token) {
      navigate('/');
      return;
    }
    getItems().then((response) => setItems(response.data));
  }, [navigate, token]);

  const handleAddToCart = async (itemId) => {
    try {
      await addToCart(itemId, token);
      toast.success('Item added to cart');
    } catch (error) {
      toast.error('Failed to add item to cart');
    }
  };

  const handleViewCart = async () => {
    try {
      const response = await getCarts();
      const userCart = response.data.find((c) => c.items.length > 0);
      if (userCart) {
        const cartItems = userCart.items.map((item) => `Cart ID: ${userCart.id}, Item ID: ${item.id}`);
        window.alert(cartItems.join('\n'));
      } else {
        window.alert('Cart is empty');
      }
    } catch (error) {
      window.alert('Failed to fetch cart');
    }
  };

  const handleViewOrders = async () => {
    try {
      const response = await getOrders();
      const orderIds = response.data.map((order) => `Order ID: ${order.id}`);
      window.alert(orderIds.join('\n') || 'No orders found');
    } catch (error) {
      window.alert('Failed to fetch orders');
    }
  };

  const handleCheckout = async () => {
    try {
      const response = await getCarts();
      const userCart = response.data.find((c) => c.items.length > 0);
      if (!userCart) {
        toast.error('Cart is empty');
        return;
      }
      await createOrder(userCart.id, token);
      toast.success('Order successful');
    } catch (error) {
      toast.error('Failed to create order');
    }
  };

  return (
    <div style={{ padding: '20px' }}>
      <ToastContainer />
      <h2>Items</h2>
      <div style={{ marginBottom: '20px' }}>
        <button onClick={handleCheckout} style={{ marginRight: '10px', padding: '10px 20px' }}>
          Checkout
        </button>
        <button onClick={handleViewCart} style={{ marginRight: '10px', padding: '10px 20px' }}>
          Cart
        </button>
        <button onClick={handleViewOrders} style={{ padding: '10px 20px' }}>
          Order History
        </button>
      </div>
      <ul style={{ listStyle: 'none', padding: 0 }}>
        {items.map((item) => (
          <li
            key={item.id}
            onClick={() => handleAddToCart(item.id)}
            style={{ cursor: 'pointer', padding: '10px', border: '1px solid #ccc', marginBottom: '5px' }}
          >
            {item.name} - ${item.price}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default ItemList;