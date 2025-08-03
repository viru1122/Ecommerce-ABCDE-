import axios from 'axios';

const API_URL = 'http://localhost:8080';

export const register = (username, password) => 
  axios.post(`${API_URL}/users`, { username, password });

export const login = (username, password) => 
  axios.post(`${API_URL}/users/login`, { username, password });

export const getItems = () => 
  axios.get(`${API_URL}/items`);

export const addToCart = (itemId, token) => 
  axios.post(`${API_URL}/carts`, { itemId }, { headers: { Authorization: token } });

export const getCarts = () => 
  axios.get(`${API_URL}/carts`);

export const createOrder = (cartId, token) => 
  axios.post(`${API_URL}/orders`, { cartId }, { headers: { Authorization: token } });

export const getOrders = () => 
  axios.get(`${API_URL}/orders`);